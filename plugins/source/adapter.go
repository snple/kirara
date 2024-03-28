package source

import (
	"sync"
	"time"

	"github.com/snple/kirara/pb"
	"go.uber.org/zap"
)

var _adapters = make(map[string]Connect)
var _lock sync.RWMutex

type Connect = func(conn *Conn) (Adapter, error)

func SetAdapter(name string, connect Connect) {
	_lock.Lock()
	defer _lock.Unlock()

	_adapters[name] = connect
}

func HasAdapter(name string) bool {
	_lock.RLock()
	defer _lock.RUnlock()

	_, has := _adapters[name]
	return has
}

func GetAdapter(name string) (Connect, bool) {
	_lock.RLock()
	defer _lock.RUnlock()

	if fn, has := _adapters[name]; has {
		return fn, true
	}

	return nil, false
}

type Adapter interface {
	Close() error
	Name() string
	ParseTag(tag *Tag) error
	ReadTags(tags map[string]Tag) error
	ReadTag(tag Tag) (string, error)
	WriteTag(tag Tag, value string) error
}

type Tag struct {
	Raw      *pb.Tag
	Metadata any
}

func (c *Conn) Source() *pb.Source {
	return c.source
}

func (c *Conn) Logger() *zap.Logger {
	return c.s.logger()
}

func (c *Conn) SyncTagValue(id string, value string) {
	t := time.Now().UnixMicro()

	if c.valueCache.setTagValue(id, value, t) {
		_, err := c.s.es.GetTag().SyncValue(c.ctx, &pb.TagValue{Id: id, Value: value, Updated: t})
		if err != nil {
			c.s.logger().Sugar().Errorf("SyncTagValue: %v", err)
		}
	}
}
