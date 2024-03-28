package edge

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/snple/kirara/db"
	"github.com/snple/kirara/model"
	"github.com/snple/kirara/pb/edges"
	"github.com/snple/types"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type EdgeService struct {
	db       *bun.DB
	badger   *BadgerService
	status   *StatusService
	sync     *SyncService
	device   *DeviceService
	slot     *SlotService
	source   *SourceService
	tag      *TagService
	constant *ConstService
	class    *ClassService
	attr     *AttrService
	logic    *LogicService
	fn       *FnService
	data     *DataService
	control  *ControlService
	save     types.Option[*SaveService]

	clone *cloneService

	auth *AuthService
	user *UserService

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	dopts edgeOptions
}

func Edge(db *bun.DB, opts ...EdgeOption) (*EdgeService, error) {
	ctx := context.Background()
	return EdgeContext(ctx, db, opts...)
}

func EdgeContext(ctx context.Context, db *bun.DB, opts ...EdgeOption) (*EdgeService, error) {
	ctx, cancel := context.WithCancel(ctx)

	if db == nil {
		panic("db == nil")
	}

	es := &EdgeService{
		db:     db,
		ctx:    ctx,
		cancel: cancel,
		dopts:  defaultEdgeOptions(),
	}

	for _, opt := range extraEdgeOptions {
		opt.apply(&es.dopts)
	}

	for _, opt := range opts {
		opt.apply(&es.dopts)
	}

	if err := es.dopts.check(); err != nil {
		return nil, err
	}

	badger, err := newBadgerService(es)
	if err != nil {
		return nil, err
	}
	es.badger = badger

	es.status = newStatusService(es)
	es.sync = newSyncService(es)
	es.device = newDeviceService(es)
	es.slot = newSlotService(es)
	es.source = newSourceService(es)
	es.tag = newTagService(es)
	es.constant = newConstService(es)
	es.class = newClassService(es)
	es.attr = newAttrService(es)
	es.logic = newLogicService(es)
	es.fn = newFnService(es)
	es.data = newDataService(es)

	if es.dopts.save {
		es.save = types.Some(newSaveService(es))
	}

	es.control = newControlService(es)

	es.clone = newCloneService(es)

	es.auth = newAuthService(es)
	es.user = newUserService(es)

	return es, nil
}

func (es *EdgeService) Start() {
	go func() {
		es.closeWG.Add(1)
		defer es.closeWG.Done()

		es.badger.start()
	}()

	if es.save.IsSome() {
		go func() {
			es.closeWG.Add(1)
			defer es.closeWG.Done()

			es.save.Unwrap().start()
		}()
	}

	if es.dopts.cache {
		go func() {
			es.closeWG.Add(1)
			defer es.closeWG.Done()

			es.cacheGC()
		}()
	}
}

func (es *EdgeService) Stop() {
	if es.save.IsSome() {
		es.save.Unwrap().stop()
	}

	es.badger.stop()

	es.cancel()
	es.closeWG.Wait()
	es.dopts.logger.Sync()
}

func (es *EdgeService) GetDB() *bun.DB {
	return es.db
}

func (es *EdgeService) GetBadgerDB() *badger.DB {
	return es.badger.GetDB()
}

func (es *EdgeService) GetInfluxDB() types.Option[*db.InfluxDB] {
	if es.dopts.influxdb != nil {
		return types.Some(es.dopts.influxdb)
	}

	return types.None[*db.InfluxDB]()
}

func (es *EdgeService) GetStatus() *StatusService {
	return es.status
}

func (es *EdgeService) GetSync() *SyncService {
	return es.sync
}

func (es *EdgeService) GetDevice() *DeviceService {
	return es.device
}

func (es *EdgeService) GetSlot() *SlotService {
	return es.slot
}

func (es *EdgeService) GetSource() *SourceService {
	return es.source
}

func (es *EdgeService) GetTag() *TagService {
	return es.tag
}

func (es *EdgeService) GetConst() *ConstService {
	return es.constant
}

func (es *EdgeService) GetClass() *ClassService {
	return es.class
}

func (es *EdgeService) GetAttr() *AttrService {
	return es.attr
}

func (es *EdgeService) GetLogic() *LogicService {
	return es.logic
}

func (es *EdgeService) GetFn() *FnService {
	return es.fn
}

func (es *EdgeService) GetData() *DataService {
	return es.data
}

func (es *EdgeService) GetSave() types.Option[*SaveService] {
	return es.save
}

func (es *EdgeService) GetControl() *ControlService {
	return es.control
}

func (es *EdgeService) getClone() *cloneService {
	return es.clone
}

func (es *EdgeService) GetAuth() *AuthService {
	return es.auth
}

func (es *EdgeService) GetUser() *UserService {
	return es.user
}

func (es *EdgeService) Context() context.Context {
	return es.ctx
}

func (es *EdgeService) Logger() *zap.Logger {
	return es.dopts.logger
}

func (es *EdgeService) cacheGC() {
	es.Logger().Sugar().Info("cache gc started")

	ticker := time.NewTicker(es.dopts.cacheGCTTL)
	defer ticker.Stop()

	for {
		select {
		case <-es.ctx.Done():
			return
		case <-ticker.C:
			{
				es.GetSource().GC()
				es.GetTag().GC()
				es.GetConst().GC()
				es.GetUser().GC()
			}
		}
	}
}

func (es *EdgeService) Register(server *grpc.Server) {
	edges.RegisterSyncServiceServer(server, es.sync)
	edges.RegisterDeviceServiceServer(server, es.device)
	edges.RegisterSlotServiceServer(server, es.slot)
	edges.RegisterSourceServiceServer(server, es.source)
	edges.RegisterTagServiceServer(server, es.tag)
	edges.RegisterConstServiceServer(server, es.constant)
	edges.RegisterClassServiceServer(server, es.class)
	edges.RegisterAttrServiceServer(server, es.attr)
	edges.RegisterLogicServiceServer(server, es.logic)
	edges.RegisterFnServiceServer(server, es.fn)
	edges.RegisterControlServiceServer(server, es.control)

	edges.RegisterAuthServiceServer(server, es.auth)
	edges.RegisterUserServiceServer(server, es.user)
}

func CreateSchema(db bun.IDB) error {
	models := []interface{}{
		(*model.Device)(nil),
		(*model.Slot)(nil),
		(*model.Source)(nil),
		(*model.Tag)(nil),
		(*model.Const)(nil),
		(*model.Class)(nil),
		(*model.Attr)(nil),
		(*model.Logic)(nil),
		(*model.Fn)(nil),
		(*model.User)(nil),
	}

	for _, model := range models {
		_, err := db.NewCreateTable().Model(model).IfNotExists().Exec(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

type edgeOptions struct {
	logger   *zap.Logger
	deviceID string
	secret   string

	BadgerOptions   badger.Options
	BadgerGCOptions BadgerGCOptions

	linkTTL      time.Duration
	cache        bool
	cacheTTL     time.Duration
	cacheGCTTL   time.Duration
	influxdb     *db.InfluxDB
	save         bool
	saveInterval time.Duration
}

type BadgerGCOptions struct {
	GC             time.Duration
	GCDiscardRatio float64
}

func defaultEdgeOptions() edgeOptions {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("zap.NewDevelopment(): %v", err)
	}

	return edgeOptions{
		logger:        logger,
		BadgerOptions: badger.DefaultOptions("").WithInMemory(true),
		BadgerGCOptions: BadgerGCOptions{
			GC:             time.Hour,
			GCDiscardRatio: 0.7,
		},
		linkTTL:      3 * time.Minute,
		cache:        true,
		cacheTTL:     3 * time.Second,
		cacheGCTTL:   3 * time.Hour,
		save:         false,
		saveInterval: time.Minute,
	}
}

func (o *edgeOptions) check() error {
	return nil
}

type EdgeOption interface {
	apply(*edgeOptions)
}

var extraEdgeOptions []EdgeOption

type funcEdgeOption struct {
	f func(*edgeOptions)
}

func (fdo *funcEdgeOption) apply(do *edgeOptions) {
	fdo.f(do)
}

func newFuncEdgeOption(f func(*edgeOptions)) *funcEdgeOption {
	return &funcEdgeOption{
		f: f,
	}
}

func WithLogger(logger *zap.Logger) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.logger = logger
	})
}

func WithDeviceID(id, secret string) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.deviceID = id
		o.secret = secret
	})
}

func WithBadger(options badger.Options) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.BadgerOptions = options
	})
}

func WithBadgerGC(options *BadgerGCOptions) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		if options.GC > 0 {
			o.BadgerGCOptions.GC = options.GC
		}

		if options.GCDiscardRatio > 0 {
			o.BadgerGCOptions.GCDiscardRatio = options.GCDiscardRatio
		}
	})
}

func WithLinkTTL(d time.Duration) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.linkTTL = d
	})
}

func WithCache(enable bool) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.cache = enable
	})
}

func WithCacheTTL(d time.Duration) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.cacheTTL = d
	})
}

func WithCacheGCTTL(d time.Duration) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.cacheGCTTL = d
	})
}

func WithInfluxDB(influxdb *db.InfluxDB) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.influxdb = influxdb
	})
}

func WithSave(enable bool) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.save = enable
	})
}

func WithSaveInterval(d time.Duration) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.saveInterval = d
	})
}
