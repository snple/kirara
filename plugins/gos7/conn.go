package gos7

import (
	"context"
	"sync"
	"time"

	"github.com/robinson/gos7"
	"github.com/snple/kirara/consts"
	"github.com/snple/kirara/edge"
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"github.com/snple/kirara/pb/slots"
)

type Conn struct {
	gs     *GoS7Slot
	source *pb.Source
	config Config

	client     *gos7.TCPClientHandler
	updated    int64
	tags       map[string]*Tag
	grouped    bool
	readGroups map[string][]readGroup
	lock       sync.RWMutex

	valueCache *valueCache

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	slots.UnimplementedControlServiceServer
}

func newConn(gs *GoS7Slot, source *pb.Source) (*Conn, error) {
	config, err := ParseConfig(source.Params)
	if err != nil {
		return nil, err
	}

	client := gos7.NewTCPClientHandler(config.Addr, config.Rank, config.Slot)
	err = client.Connect()
	if err != nil {
		gs.logger().Sugar().Errorf("GoS7 source link: %v, source: %v", err, source)
		return nil, err
	}

	reply, err := gs.es.GetSync().GetTagValueUpdated(gs.ctx, &pb.MyEmpty{})
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(gs.ctx)

	conn := &Conn{
		gs:         gs,
		source:     source,
		config:     config,
		client:     client,
		tags:       make(map[string]*Tag),
		valueCache: newValueCache(reply.GetUpdated()),
		ctx:        ctx,
		cancel:     cancel,
	}

	if err = conn.fetchTags(); err != nil {
		gs.logger().Sugar().Errorf("GoS7 fetchTags: %v", err)
		return nil, err
	}

	go conn.waitTagValueUpdated()
	go conn.syncLinkStatus()

	return conn, nil
}

func (c *Conn) start() {
	c.closeWG.Add(1)
	defer c.closeWG.Done()

	c.gs.es.GetControl().AddSourceServer(c.source.GetId(), c)
	defer c.gs.es.GetControl().RemoveSourceServer(c.source.GetId(), c)

	c.linkSource(consts.ON)
	defer c.linkSource(consts.OFF)

	readTicker := time.NewTicker(c.gs.dopts.readDataInterval)
	defer readTicker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-readTicker.C:
			c.lock.Lock()
			grouped := c.grouped

			if !grouped {
				c.readGroups = groupTags(c.tags)
				c.grouped = true
			}

			readGroup := c.readGroups
			c.lock.Unlock()

			if err := c.readTags(readGroup); err != nil {
				c.gs.logger().Sugar().Errorf("GoS7 readTags: %v", err)
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

	reply, err := c.gs.es.GetTag().List(c.ctx, &request)
	if err != nil {
		return err
	}

	for _, tag := range reply.GetTag() {
		if tag.GetAddress() != "" && tag.GetDataType() != "" && tag.GetStatus() == consts.ON {
			tag2 := newTag(tag)
			tag2.parse()

			c.tags[tag.GetId()] = tag2
		}
	}

	c.grouped = false

	return nil
}

func (c *Conn) checkTagUpdated() error {
	tagUpdated, err := c.gs.es.GetSync().GetTagUpdated(c.ctx, &pb.MyEmpty{})
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
			remotes, err := c.gs.es.GetTag().Pull(c.ctx,
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

					tag2 := newTag(remote)
					tag2.parse()
					c.tags[remote.GetId()] = tag2
				}

				c.grouped = false
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

	notify := c.gs.es.GetSync().Notify(edge.NOTIFY_TVAL)
	defer notify.Close()

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-notify.Wait():
			time.Sleep(time.Second)

			err := c.checkTagValueUpdated()
			if err != nil {
				c.gs.logger().Sugar().Errorf("checkTagValueUpdated: %v", err)
			}
		}
	}
}

func (c *Conn) checkTagValueUpdated() error {
	reply, err := c.gs.es.GetSync().GetTagValueUpdated(c.ctx, &pb.MyEmpty{})
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
			remotes, err := c.gs.es.GetTag().PullValue(c.ctx,
				&edges.TagPullValueRequest{After: after, Limit: limit, SourceId: c.source.GetId()})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetTag() {
				after = remote.GetUpdated()

				if tag, ok := c.getTag(remote.GetId()); ok {
					if tag.raw.GetAccess() != consts.ON {
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

					err := c.writeTag(remote.GetId(), remote.GetValue())
					if err != nil {
						c.gs.logger().Sugar().Errorf("GoS7 writeTag: %v", err)
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

func (c *Conn) getTag(id string) (*Tag, bool) {
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
				c.gs.logger().Sugar().Errorf("source link error: %v", err)
			}
		}
	}
}

func (c *Conn) linkSource(status int32) error {
	request := edges.SourceLinkRequest{Id: c.source.GetId(), Status: status}
	_, err := c.gs.es.GetSource().Link(c.ctx, &request)
	return err
}
