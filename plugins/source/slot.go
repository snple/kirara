package source

import (
	"context"
	"sync"
	"time"

	"github.com/snple/kirara/consts"
	"github.com/snple/kirara/edge"
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"go.uber.org/zap"
)

type TagIOSlot struct {
	es *edge.EdgeService

	conns   map[string]*Conn
	updated int64
	lock    sync.RWMutex

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	dopts tagIOOptions
}

func GoS7(es *edge.EdgeService, opts ...TagIOOption) (*TagIOSlot, error) {
	ctx, cancel := context.WithCancel(es.Context())

	ts := &TagIOSlot{
		es:     es,
		conns:  make(map[string]*Conn),
		ctx:    ctx,
		cancel: cancel,
		dopts:  defaultTagIOOptions(),
	}

	for _, opt := range extratagIOOptions {
		opt.apply(&ts.dopts)
	}

	for _, opt := range opts {
		opt.apply(&ts.dopts)
	}

	return ts, nil
}

func (ts *TagIOSlot) Start() {
	ts.closeWG.Add(1)
	defer ts.closeWG.Done()

	ts.logger().Info("TagIO slot started")

	go ts.waitDeviceUpdated()

	ticker := time.NewTicker(ts.dopts.tickerInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ts.ctx.Done():
			return
		case <-ticker.C:
			err := ts.ticker()
			if err != nil {
				ts.logger().Sugar().Errorf("TagIO ticker: %v", err)
			}
		}
	}
}

func (ts *TagIOSlot) Stop() {
	ts.cancel()
	ts.closeWG.Wait()
}

func (ts *TagIOSlot) ticker() error {
	request := edges.SourceListRequest{
		Page: &pb.Page{
			Limit: 1000,
		},
		// Source: SOURCE,
	}

	reply, err := ts.es.GetSource().List(ts.ctx, &request)
	if err != nil {
		return err
	}

	for _, source := range reply.GetSource() {
		ts.startConn(source)
	}

	return nil
}

func (ts *TagIOSlot) waitDeviceUpdated() {
	ts.closeWG.Add(1)
	defer ts.closeWG.Done()

	notify := ts.es.GetSync().Notify(edge.NOTIFY)
	defer notify.Close()

	for {
		select {
		case <-ts.ctx.Done():
			return
		case <-notify.Wait():
			time.Sleep(time.Second)

			err := ts.checkUpdated()
			if err != nil {
				ts.logger().Sugar().Errorf("TagIO checkUpdated: %v", err)
			}
		}
	}
}

func (ts *TagIOSlot) checkUpdated() error {
	if err := ts.checkSourceUpdated(); err != nil {
		return err
	}

	return ts.checkTagUpdated()
}

func (ts *TagIOSlot) checkSourceUpdated() error {
	sourceUpdated, err := ts.es.GetSync().GetSourceUpdated(ts.ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	updated := ts.getUpdated()

	if sourceUpdated.GetUpdated() <= updated {
		return nil
	}

	{
		after := updated
		limit := uint32(10)

		for {
			remotes, err := ts.es.GetSource().Pull(ts.ctx,
				&edges.SourcePullRequest{After: after, Limit: limit /*Source: SOURCE*/})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetSource() {
				after = remote.GetUpdated()

				if remote.GetDeleted() > 0 {
					// delete
					ts.lock.Lock()
					if conn, ok := ts.conns[remote.GetId()]; ok {
						conn.stop()
					}
					ts.lock.Unlock()
				} else {
					ts.checkConn(remote)
				}
			}

			if len(remotes.GetSource()) < int(limit) {
				break
			}
		}
	}

	ts.setUpdated(sourceUpdated.GetUpdated())

	return nil
}

func (ts *TagIOSlot) getUpdated() int64 {
	ts.lock.RLock()
	defer ts.lock.RUnlock()

	return ts.updated
}

func (ts *TagIOSlot) setUpdated(updated int64) {
	ts.lock.Lock()
	defer ts.lock.Unlock()

	ts.updated = updated
}

func (ts *TagIOSlot) startConn(source *pb.Source) {
	if source.GetSource() == "" || source.GetParams() == "" || source.GetStatus() != consts.ON {
		return
	}

	if !HasAdapter(source.GetSource()) {
		return
	}

	ts.lock.Lock()
	defer ts.lock.Unlock()

	if _, ok := ts.conns[source.GetId()]; ok {
		return
	}

	conn, err := newConn(ts, source)
	if err != nil {
		ts.logger().Sugar().Errorf("newConn error: %v", err)
		return
	}

	ts.conns[source.GetId()] = conn

	go func() {
		ts.closeWG.Add(1)
		defer ts.closeWG.Done()

		conn.run()

		ts.lock.Lock()
		delete(ts.conns, source.GetId())
		ts.lock.Unlock()
	}()
}

func (ts *TagIOSlot) checkConn(source *pb.Source) {
	ts.lock.Lock()
	defer ts.lock.Unlock()

	if conn, ok := ts.conns[source.GetId()]; ok {
		if source.GetStatus() == consts.ON && source.GetName() == conn.source.GetName() &&
			source.GetParams() == conn.source.GetParams() {
			return
		}

		conn.stop()
	}
}

func (ts *TagIOSlot) checkTagUpdated() error {
	ts.lock.RLock()
	defer ts.lock.RUnlock()

	for _, conn := range ts.conns {
		if err := conn.checkTagUpdated(); err != nil {
			ts.logger().Sugar().Errorf("TagIO checkTagUpdated: %v", err)
		}
	}

	return nil
}

func (ts *TagIOSlot) logger() *zap.Logger {
	return ts.es.Logger()
}

type tagIOOptions struct {
	debug            bool
	tickerInterval   time.Duration
	readDataInterval time.Duration
}

func defaultTagIOOptions() tagIOOptions {
	return tagIOOptions{
		debug:            false,
		tickerInterval:   60 * time.Second,
		readDataInterval: 30 * time.Second,
	}
}

type TagIOOption interface {
	apply(*tagIOOptions)
}

var extratagIOOptions []TagIOOption

type funcTagIOOption struct {
	f func(*tagIOOptions)
}

func (fdo *funcTagIOOption) apply(do *tagIOOptions) {
	fdo.f(do)
}

func newFuncTagIOOption(f func(*tagIOOptions)) *funcTagIOOption {
	return &funcTagIOOption{
		f: f,
	}
}

func WithDebug(debug bool) TagIOOption {
	return newFuncTagIOOption(func(o *tagIOOptions) {
		o.debug = debug
	})
}

func WithTickerInterval(d time.Duration) TagIOOption {
	return newFuncTagIOOption(func(o *tagIOOptions) {
		o.tickerInterval = d
	})
}

func WithReadDataInterval(d time.Duration) TagIOOption {
	return newFuncTagIOOption(func(o *tagIOOptions) {
		o.readDataInterval = d
	})
}
