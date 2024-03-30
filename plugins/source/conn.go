package source

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/snple/kirara/consts"
	"github.com/snple/kirara/edge"
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"github.com/snple/kirara/pb/slots"
)

type Conn struct {
	s      *SourceSlot
	source *pb.Source

	updated int64
	tags    map[string]Tag
	lock    sync.RWMutex

	valueCache *valueCache

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	slots.UnimplementedControlServiceServer

	adapter Adapter
}

func newConn(s *SourceSlot, source *pb.Source) (*Conn, error) {
	connect, has := GetAdapter(source.GetSource())
	if !has {
		return nil, fmt.Errorf("adapter %v not found", source.GetSource())
	}

	ctx, cancel := context.WithCancel(s.ctx)

	conn := &Conn{
		s:      s,
		source: source,
		ctx:    ctx,
		cancel: cancel,
	}

	adapter, err := connect(conn)
	if err != nil {
		s.logger().Sugar().Errorf("adapter %v connect error %v, source: %v", source.GetSource(), err, source)
		return nil, err
	}
	conn.adapter = adapter

	err = conn.init()
	if err != nil {
		conn.adapter.Close()
		return nil, err
	}

	return conn, nil
}

func (c *Conn) init() error {
	reply, err := c.s.es.GetSync().GetTagValueUpdated(c.ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}
	c.valueCache = newValueCache(reply.GetUpdated())
	c.tags = make(map[string]Tag)

	if err = c.fetchTags(); err != nil {
		c.s.logger().Sugar().Errorf("TagIO fetchTags: %v", err)
		return err
	}

	return nil
}

func (c *Conn) run() {
	defer c.closeWG.Wait()

	c.s.es.GetControl().AddSourceServer(c.source.GetId(), c)
	defer c.s.es.GetControl().RemoveSourceServer(c.source.GetId(), c)

	c.linkSource(consts.ON)
	defer c.linkSource(consts.OFF)

	readTicker := time.NewTicker(c.s.dopts.readDataInterval)
	defer readTicker.Stop()

	defer c.cancel()

	defer c.adapter.Close()

	go c.waitTagValueUpdated()
	go c.syncLinkStatus()

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-readTicker.C:
			c.lock.Lock()
			err := c.adapter.ReadTags(c.tags)
			c.lock.Unlock()

			if err != nil {
				c.s.logger().Sugar().Errorf("%v readTags: %v", c.adapter.Name(), err)
				return
			}
		}
	}
}

func (c *Conn) stop() {
	c.cancel()
	c.closeWG.Wait()
}

func (c *Conn) fetchTags() error {
	c.lock.Lock()
	defer c.lock.Unlock()

	request := edges.TagListRequest{
		Page: &pb.Page{
			Limit: 1000,
		},
		SourceId: c.source.Id,
	}

	reply, err := c.s.es.GetTag().List(c.ctx, &request)
	if err != nil {
		return err
	}

	for _, tag := range reply.GetTag() {
		if tag.GetAddress() != "" && tag.GetDataType() != "" && tag.GetStatus() == consts.ON {
			tag2 := Tag{Raw: tag}

			err = c.adapter.ParseTag(&tag2)
			if err != nil {
				c.s.logger().Sugar().Errorf("%v ParseTag: %v, err: %v", c.adapter.Name(), tag, err)
				continue
			}

			c.tags[tag.GetId()] = tag2
		}
	}

	return nil
}

func (c *Conn) checkTagUpdated() error {
	tagUpdated, err := c.s.es.GetSync().GetTagUpdated(c.ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	updated := c.getUpdated()

	if tagUpdated.GetUpdated() <= updated {
		return nil
	}

	{
		after := updated
		limit := uint32(10)

		for {
			remotes, err := c.s.es.GetTag().Pull(c.ctx,
				&edges.TagPullRequest{After: after, Limit: limit, SourceId: c.source.GetId()})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetTag() {
				after = remote.GetUpdated()

				c.lock.Lock()
				delete(c.tags, remote.GetId())

				if remote.GetDeleted() <= 0 && remote.GetAddress() != "" &&
					remote.GetDataType() != "" && remote.GetStatus() == consts.ON {
					tag2 := Tag{Raw: remote}

					err = c.adapter.ParseTag(&tag2)
					if err != nil {
						c.s.logger().Sugar().Errorf("%v ParseTag: %v, err: %v", c.adapter.Name(), remote, err)
						continue
					}

					c.tags[remote.GetId()] = tag2
				}
				c.lock.Unlock()
			}

			if len(remotes.GetTag()) < int(limit) {
				break
			}
		}
	}

	c.setUpdated(tagUpdated.GetUpdated())

	return nil
}

func (c *Conn) waitTagValueUpdated() {
	c.closeWG.Add(1)
	defer c.closeWG.Done()

	notify := c.s.es.GetSync().Notify(edge.NOTIFY_TVAL)
	defer notify.Close()

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-notify.Wait():
			time.Sleep(time.Second)

			err := c.checkTagValueUpdated()
			if err != nil {
				c.s.logger().Sugar().Errorf("checkTagValueUpdated: %v", err)
			}
		}
	}
}

func (c *Conn) checkTagValueUpdated() error {
	reply, err := c.s.es.GetSync().GetTagValueUpdated(c.ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	valueUpdated := reply.GetUpdated()
	localValueUpdated := c.valueCache.getUpdated()

	if valueUpdated <= localValueUpdated {
		return nil
	}

	{
		after := localValueUpdated
		limit := uint32(10)

		for {
			remotes, err := c.s.es.GetTag().PullValue(c.ctx,
				&edges.TagPullValueRequest{After: after, Limit: limit, SourceId: c.source.GetId()})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetTag() {
				after = remote.GetUpdated()

				if tag, ok := c.getTag(remote.GetId()); ok {
					if tag.Raw.GetAccess() != consts.ON {
						continue
					}

					if item, ok := c.valueCache.getTagValue(remote.GetId()); ok {
						if remote.GetValue() == item.value {
							continue
						}

						if remote.GetUpdated() <= item.updated {
							continue
						}
					}

					err := c.adapter.WriteTag(tag, remote.GetValue())
					if err != nil {
						c.s.logger().Sugar().Errorf("%v writeTag: %v", c.adapter.Name(), err)
					}
				}
			}

			if len(remotes.GetTag()) < int(limit) {
				break
			}
		}
	}

	c.valueCache.setUpdated(valueUpdated)

	return nil
}

func (c *Conn) getTag(id string) (Tag, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	tag, ok := c.tags[id]
	return tag, ok
}

func (c *Conn) getUpdated() int64 {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.updated
}

func (c *Conn) setUpdated(updated int64) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.updated = updated
}

func (c *Conn) syncLinkStatus() {
	c.closeWG.Add(1)
	defer c.closeWG.Done()

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-ticker.C:
			err := c.linkSource(consts.ON)
			if err != nil {
				c.s.logger().Sugar().Errorf("source link error: %v", err)
			}
		}
	}
}

func (c *Conn) linkSource(status int32) error {
	request := edges.SourceLinkRequest{Id: c.source.GetId(), Status: status}
	_, err := c.s.es.GetSource().Link(c.ctx, &request)
	return err
}
