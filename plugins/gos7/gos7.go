package gos7

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

type GoS7Slot struct {
	es *edge.EdgeService

	conns   map[string]*Conn
	updated int64
	lock    sync.RWMutex

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	dopts goS7Options
}

const (
	SOURCE = "GOS7"
)

func GoS7(es *edge.EdgeService, opts ...GoS7Option) (*GoS7Slot, error) {
	ctx, cancel := context.WithCancel(es.Context())

	gs := &GoS7Slot{
		es:     es,
		conns:  make(map[string]*Conn),
		ctx:    ctx,
		cancel: cancel,
		dopts:  defaultGoS7Options(),
	}

	for _, opt := range extraGoS7Options {
		opt.apply(&gs.dopts)
	}

	for _, opt := range opts {
		opt.apply(&gs.dopts)
	}

	return gs, nil
}

func (gs *GoS7Slot) Start() {
	gs.closeWG.Add(1)
	defer gs.closeWG.Done()

	gs.logger().Info("GoS7 slot started")

	go gs.waitDeviceUpdated()

	ticker := time.NewTicker(gs.dopts.tickerInterval)
	defer ticker.Stop()

	for {
		select {
		case <-gs.ctx.Done():
			return
		case <-ticker.C:
			err := gs.ticker()
			if err != nil {
				gs.logger().Sugar().Errorf("GoS7 ticker: %v", err)
			}
		}
	}
}

func (gs *GoS7Slot) Stop() {
	gs.cancel()
	gs.closeWG.Wait()
}

func (gs *GoS7Slot) ticker() error {
	request := edges.SourceListRequest{
		Page: &pb.Page{
			Limit: 1000,
		},
		Source: SOURCE,
	}

	reply, err := gs.es.GetSource().List(gs.ctx, &request)
	if err != nil {
		return err
	}

	for _, source := range reply.GetSource() {
		gs.startConn(source)
	}

	return nil
}

func (gs *GoS7Slot) waitDeviceUpdated() {
	gs.closeWG.Add(1)
	defer gs.closeWG.Done()

	notify := gs.es.GetSync().Notify(edge.NOTIFY)
	defer notify.Close()

	for {
		select {
		case <-gs.ctx.Done():
			return
		case <-notify.Wait():
			time.Sleep(time.Second)

			err := gs.checkUpdated()
			if err != nil {
				gs.logger().Sugar().Errorf("GoS7 checkUpdated: %v", err)
			}
		}
	}
}

func (gs *GoS7Slot) checkUpdated() error {
	if err := gs.checkSourceUpdated(); err != nil {
		return err
	}

	return gs.checkTagUpdated()
}

func (gs *GoS7Slot) checkSourceUpdated() error {
	sourceUpdated, err := gs.es.GetSync().GetSourceUpdated(gs.ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	updated := gs.getUpdated()

	if sourceUpdated.GetUpdated() <= updated {
		return nil
	}

	{
		after := updated
		limit := uint32(10)

		for {
			remotes, err := gs.es.GetSource().Pull(gs.ctx,
				&edges.SourcePullRequest{After: after, Limit: limit, Source: SOURCE})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetSource() {
				after = remote.GetUpdated()

				if remote.GetDeleted() > 0 {
					// delete
					gs.lock.Lock()
					if conn, ok := gs.conns[remote.GetId()]; ok {
						conn.stop()
					}
					gs.lock.Unlock()
				} else {
					gs.checkConn(remote)
				}
			}

			if len(remotes.GetSource()) < int(limit) {
				break
			}
		}
	}

	gs.setUpdated(sourceUpdated.GetUpdated())

	return nil
}

func (gs *GoS7Slot) getUpdated() int64 {
	gs.lock.RLock()
	defer gs.lock.RUnlock()

	return gs.updated
}

func (gs *GoS7Slot) setUpdated(updated int64) {
	gs.lock.Lock()
	defer gs.lock.Unlock()

	gs.updated = updated
}

func (gs *GoS7Slot) startConn(source *pb.Source) {
	if source.GetParams() == "" || source.GetStatus() != consts.ON {
		return
	}

	gs.lock.Lock()
	defer gs.lock.Unlock()

	if _, ok := gs.conns[source.GetId()]; ok {
		return
	}

	conn, err := newConn(gs, source)
	if err != nil {
		gs.logger().Sugar().Errorf("newConn error: %v", err)
		return
	}

	gs.conns[source.GetId()] = conn

	go func() {
		gs.closeWG.Add(1)
		defer gs.closeWG.Done()

		conn.start()

		gs.lock.Lock()
		delete(gs.conns, source.GetId())
		gs.lock.Unlock()

		conn.closeWG.Wait()
	}()
}

func (gs *GoS7Slot) checkConn(source *pb.Source) {
	gs.lock.Lock()
	defer gs.lock.Unlock()

	if conn, ok := gs.conns[source.GetId()]; ok {
		if source.GetStatus() == consts.ON && source.GetName() == conn.source.GetName() &&
			source.GetParams() == conn.source.GetParams() {
			return
		}

		conn.stop()
	}
}

func (gs *GoS7Slot) checkTagUpdated() error {
	gs.lock.RLock()
	defer gs.lock.RUnlock()

	for _, conn := range gs.conns {
		if err := conn.checkTagUpdated(); err != nil {
			gs.logger().Sugar().Errorf("GoS7 checkTagUpdated: %v", err)
		}
	}

	return nil
}

func (gs *GoS7Slot) logger() *zap.Logger {
	return gs.es.Logger()
}

type goS7Options struct {
	debug            bool
	tickerInterval   time.Duration
	readDataInterval time.Duration
}

func defaultGoS7Options() goS7Options {
	return goS7Options{
		debug:            false,
		tickerInterval:   60 * time.Second,
		readDataInterval: 30 * time.Second,
	}
}

type GoS7Option interface {
	apply(*goS7Options)
}

var extraGoS7Options []GoS7Option

type funcGoS7Option struct {
	f func(*goS7Options)
}

func (fdo *funcGoS7Option) apply(do *goS7Options) {
	fdo.f(do)
}

func newFuncGoS7Option(f func(*goS7Options)) *funcGoS7Option {
	return &funcGoS7Option{
		f: f,
	}
}

func WithDebug(debug bool) GoS7Option {
	return newFuncGoS7Option(func(o *goS7Options) {
		o.debug = debug
	})
}

func WithTickerInterval(d time.Duration) GoS7Option {
	return newFuncGoS7Option(func(o *goS7Options) {
		o.tickerInterval = d
	})
}

func WithReadDataInterval(d time.Duration) GoS7Option {
	return newFuncGoS7Option(func(o *goS7Options) {
		o.readDataInterval = d
	})
}
