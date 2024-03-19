package gos7

import (
	"errors"
	"fmt"
	"time"

	"github.com/robinson/gos7"
	"github.com/snple/kirara/pb"
)

func (c *Conn) readTags(readGroups map[string][]readGroup) (err error) {
	defer func() {
		if re := recover(); re != nil {
			err = errors.New(fmt.Sprint(re))
		}
	}()

	for _, groups := range readGroups {
		for _, group := range groups {
			var buffer = make([]byte, group.end-group.start+1)
			var err error

			c.lock.Lock()
			client := gos7.NewClient(c.client)
			switch group.area {
			case AreaPE:
				err = client.AGReadEB(group.start, group.end-group.start+1, buffer)
			case AreaPA:
				err = client.AGReadAB(group.start, group.end-group.start+1, buffer)
			case AreaMK:
				err = client.AGReadMB(group.start, group.end-group.start+1, buffer)
			case AreaDB:
				err = client.AGReadDB(group.db, group.start, group.end-group.start+1, buffer)
			case AreaCT:
				err = client.AGReadCT(group.start, group.end-group.start+1, buffer)
			case AreaTM:
				err = client.AGReadTM(group.start, group.end-group.start+1, buffer)
			default:
				err = fmt.Errorf("unsupport area: %v", group.area)
			}
			c.lock.Unlock()

			if err != nil {
				c.gs.logger().Sugar().Errorf("read error: area: %v, db: %v, start: %v, size: %v, err: %v",
					group.area, group.db, group.start, len(buffer), err)

				return err
			}

			for _, tag := range group.tagList {
				bytes := buffer[tag.addr.Address-group.start : (tag.addr.Address-group.start)+tag.addr.Size]

				value, err := convertBytesToValue(&c.config, tag, bytes)
				if err != nil {
					c.gs.logger().Sugar().Error(err)
					continue
				}

				c.syncTagValue(tag.raw.GetId(), value)

				if c.config.Debug {
					c.gs.logger().Sugar().Debugf("read tag: %v %v %v %v -> %v", tag.raw.GetId(), tag.raw.GetName(), tag.raw.GetAddress(), bytes, value)
				}
			}
		}
	}

	return nil
}

func (c *Conn) syncTagValue(id string, value string) {
	t := time.Now().UnixMicro()

	if c.valueCache.setTagValue(id, value, t) {
		_, err := c.gs.es.GetTag().SyncValue(c.ctx, &pb.TagValue{Id: id, Value: value, Updated: t})
		if err != nil {
			c.gs.logger().Sugar().Errorf("GoS7 SyncValue: %v", err)
		}
	}
}

func (c *Conn) readTag(id string) (value string, err error) {
	defer func() {
		if re := recover(); re != nil {
			err = errors.New(fmt.Sprint(re))
		}
	}()

	c.lock.Lock()
	defer c.lock.Unlock()

	if tag, ok := c.tags[id]; ok {
		var buffer = make([]byte, tag.addr.Size)

		client := gos7.NewClient(c.client)

		switch tag.addr.Area {
		case AreaPE:
			err = client.AGReadEB(tag.addr.Address, tag.addr.Size, buffer)
		case AreaPA:
			err = client.AGReadAB(tag.addr.Address, tag.addr.Size, buffer)
		case AreaMK:
			err = client.AGReadMB(tag.addr.Address, tag.addr.Size, buffer)
		case AreaDB:
			err = client.AGReadDB(tag.addr.DB, tag.addr.Address, tag.addr.Size, buffer)
		case AreaCT:
			err = client.AGReadCT(tag.addr.Address, tag.addr.Size, buffer)
		case AreaTM:
			err = client.AGReadTM(tag.addr.Address, tag.addr.Size, buffer)
		default:
			err = fmt.Errorf("unsupport area: %v", tag.addr.Area)
		}

		if err != nil {
			c.gs.logger().Sugar().Errorf("read error: id: %v, area: %v, db: %v, start: %v, size: %v, err: %v",
				id, tag.addr.Area, tag.addr.DB, tag.addr.Address, tag.addr.Size, err)
			return "", err
		}

		value, err := convertBytesToValue(&c.config, tag, buffer)
		if err != nil {
			c.gs.logger().Sugar().Error(err)
			return "", err
		}

		if c.config.Debug {
			c.gs.logger().Sugar().Debugf("read tag: %v %v %v %v -> %v", tag.raw.GetId(), tag.raw.GetName(), tag.raw.GetAddress(), buffer, value)
		}

		return value, nil
	}

	return "", fmt.Errorf("tag: %v is not found", id)
}
