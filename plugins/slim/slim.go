package slim

import (
	"context"
	"sync"
	"time"

	"github.com/snple/kirara/consts"
	"github.com/snple/kirara/edge"
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

type SlimSlot struct {
	es *edge.EdgeService

	procs        map[string]*Proc
	fnUpdated    int64
	logicUpdated int64
	lock         sync.RWMutex

	moduleDpds *ModuleDpds

	db *bbolt.DB

	store *Store

	moduleDevice  *ModuleDevice
	moduleSource  *ModuleSource
	moduleTag     *ModuleTag
	moduleConst   *ModuleConst
	moduleClass   *ModuleClass
	moduleAttr    *ModuleAttr
	moduleControl *ModuleControl

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	dopts slimOptions
}

const EXEC = "slim"

func Slim(es *edge.EdgeService, opts ...SlimOption) (*SlimSlot, error) {
	ctx, cancel := context.WithCancel(es.Context())

	ss := &SlimSlot{
		es:         es,
		procs:      make(map[string]*Proc),
		moduleDpds: NewModuleDpds(),
		ctx:        ctx,
		cancel:     cancel,
		dopts:      defaultSlimOptions(),
	}

	for _, opt := range extraSlimOptions {
		opt.apply(&ss.dopts)
	}

	for _, opt := range opts {
		opt.apply(&ss.dopts)
	}

	// open db
	db, err := bbolt.Open(ss.dopts.bbolt, 0600, &bbolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		return nil, err
	}
	ss.db = db

	store, err := NewStore(ss, ss.dopts.cacheTTL)
	if err != nil {
		return nil, err
	}
	ss.store = store

	ss.moduleDevice = NewModuleDevice(ss)
	ss.moduleSource = NewModuleSource(ss)
	ss.moduleTag = NewModuleTag(ss)
	ss.moduleConst = NewModuleConst(ss)
	ss.moduleClass = NewModuleClass(ss)
	ss.moduleAttr = NewModuleAttr(ss)
	ss.moduleControl = NewModuleControl(ss)

	return ss, nil
}

func (ss *SlimSlot) Start() {
	ss.closeWG.Add(1)
	defer ss.closeWG.Done()

	ss.logger().Info("Slim slot started")

	go ss.waitDeviceUpdated()

	err := ss.ticker()
	if err != nil {
		ss.logger().Sugar().Errorf("Slim ticker: %v", err)
	}

	ticker := time.NewTicker(ss.dopts.tickerInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ss.ctx.Done():
			return
		case <-ticker.C:
			err := ss.ticker()
			if err != nil {
				ss.logger().Sugar().Errorf("Slim ticker: %v", err)
			}
		}
	}
}

func (ss *SlimSlot) Stop() {
	ss.cancel()
	ss.closeWG.Wait()

	ss.db.Close()
}

func (ss *SlimSlot) ticker() error {
	request := edges.FnListRequest{
		Page: &pb.Page{
			Limit: 1000,
		},
		Exec: EXEC,
	}

	reply, err := ss.es.GetFn().List(ss.ctx, &request)
	if err != nil {
		return err
	}

	for _, fn := range reply.GetFn() {
		ss.startProc(fn)
	}

	return nil
}

func (ss *SlimSlot) waitDeviceUpdated() {
	ss.closeWG.Add(1)
	defer ss.closeWG.Done()

	notify := ss.es.GetSync().Notify(edge.NOTIFY)
	defer notify.Close()

	for {
		select {
		case <-ss.ctx.Done():
			return
		case <-notify.Wait():
			time.Sleep(time.Second)

			err := ss.checkUpdated()
			if err != nil {
				ss.logger().Sugar().Errorf("Slim checkConfigUpdated: %v", err)
			}
		}
	}
}

func (ss *SlimSlot) checkUpdated() error {
	if err := ss.checkFnUpdated(); err != nil {
		return err
	}

	return ss.checkLogicUpdated()
}

func (ss *SlimSlot) checkFnUpdated() error {
	fnUpdated, err := ss.es.GetSync().GetFnUpdated(ss.ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	updated := ss.getFnUpdated()

	if fnUpdated.GetUpdated() <= updated {
		return nil
	}

	{
		after := updated
		limit := uint32(10)

		for {
			remotes, err := ss.es.GetFn().Pull(ss.ctx,
				&edges.FnPullRequest{After: after, Limit: limit, Exec: EXEC})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetFn() {
				after = remote.GetUpdated()

				if remote.GetDeleted() > 0 {
					// delete
					ss.lock.Lock()
					if proc, ok := ss.procs[remote.GetId()]; ok {
						proc.stop()
					}
					ss.lock.Unlock()
				} else {
					ss.checkProc(remote)
				}
			}

			if len(remotes.GetFn()) < int(limit) {
				break
			}
		}
	}

	ss.setFnUpdated(fnUpdated.GetUpdated())

	return nil
}

func (ss *SlimSlot) checkLogicUpdated() error {
	logicUpdated, err := ss.es.GetSync().GetLogicUpdated(ss.ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	updated := ss.getLogicUpdated()

	if logicUpdated.GetUpdated() <= updated {
		return nil
	}

	{
		after := updated
		limit := uint32(10)

		for {
			remotes, err := ss.es.GetLogic().Pull(ss.ctx,
				&edges.LogicPullRequest{After: after, Limit: limit, Exec: EXEC})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetLogic() {
				after = remote.GetUpdated()

				ids := ss.moduleDpds.Take(remote.GetId())

				for id := range ids {
					ss.resetProc(id)
				}
			}

			if len(remotes.GetLogic()) < int(limit) {
				break
			}
		}
	}

	ss.setLogicUpdated(logicUpdated.GetUpdated())

	return nil
}

func (ss *SlimSlot) startProc(fn *pb.Fn) {
	if fn.GetMain() == "" || fn.GetStatus() != consts.ON {
		return
	}

	ss.lock.Lock()
	defer ss.lock.Unlock()

	if _, ok := ss.procs[fn.GetId()]; ok {
		return
	}

	proc, err := newProc(ss, fn)
	if err != nil {
		ss.logger().Sugar().Errorf("newProc error: %v", err)
		return
	}

	ss.procs[fn.GetId()] = proc

	go func() {
		ss.closeWG.Add(1)
		defer ss.closeWG.Done()

		proc.start()

		ss.lock.Lock()
		delete(ss.procs, fn.GetId())
		ss.lock.Unlock()
	}()
}

func (ss *SlimSlot) checkProc(fn *pb.Fn) {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	if proc, ok := ss.procs[fn.GetId()]; ok {
		if fn.GetStatus() == consts.ON && fn.GetName() == proc.fn.GetName() &&
			fn.GetMain() == proc.fn.GetMain() {
			if fn.GetDebug() != proc.getDebug() {
				proc.setDebug(fn.GetDebug())
			}

			return
		}

		proc.stop()
	}
}

func (ss *SlimSlot) resetProc(id string) {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	if proc, ok := ss.procs[id]; ok {
		proc.stop()
	}
}

func (ss *SlimSlot) getFnUpdated() int64 {
	ss.lock.RLock()
	defer ss.lock.RUnlock()

	return ss.fnUpdated
}

func (ss *SlimSlot) setFnUpdated(updated int64) {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	ss.fnUpdated = updated
}

func (ss *SlimSlot) getLogicUpdated() int64 {
	ss.lock.RLock()
	defer ss.lock.RUnlock()

	return ss.logicUpdated
}

func (ss *SlimSlot) setLogicUpdated(updated int64) {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	ss.logicUpdated = updated
}

func (ss *SlimSlot) logger() *zap.Logger {
	return ss.es.Logger()
}

type slimOptions struct {
	debug          bool
	tickerInterval time.Duration
	bbolt          string
	cacheTTL       time.Duration
}

func defaultSlimOptions() slimOptions {
	return slimOptions{
		debug:          false,
		tickerInterval: 60 * time.Second,
		bbolt:          "slim.db",
		cacheTTL:       time.Hour * 24 * 7,
	}
}

type SlimOption interface {
	apply(*slimOptions)
}

var extraSlimOptions []SlimOption

type funcSlimOption struct {
	f func(*slimOptions)
}

func (fdo *funcSlimOption) apply(do *slimOptions) {
	fdo.f(do)
}

func newFuncSlimOption(f func(*slimOptions)) *funcSlimOption {
	return &funcSlimOption{
		f: f,
	}
}

func WithDebug(debug bool) SlimOption {
	return newFuncSlimOption(func(o *slimOptions) {
		o.debug = debug
	})
}

func WithTickerInterval(d time.Duration) SlimOption {
	return newFuncSlimOption(func(o *slimOptions) {
		o.tickerInterval = d
	})
}

func WithBBolt(file string) SlimOption {
	return newFuncSlimOption(func(o *slimOptions) {
		o.bbolt = file
	})
}

func WithCacheTTL(d time.Duration) SlimOption {
	return newFuncSlimOption(func(o *slimOptions) {
		o.cacheTTL = d
	})
}
