package gos7

import (
	"errors"
	"fmt"
	"sort"

	"github.com/robinson/gos7"
	"github.com/snple/kirara/plugins/source"
)

func (s *GoS7) readTags(tags map[string]source.Tag) (err error) {
	defer func() {
		if re := recover(); re != nil {
			err = errors.New(fmt.Sprint(re))
		}
	}()

	s.lock.Lock()
	defer s.lock.Unlock()

	readGroups := groupTags(tags)

	for _, groups := range readGroups {
		for _, group := range groups {
			var buffer = make([]byte, group.end-group.start+1)
			var err error

			client := gos7.NewClient(s.client)
			switch group.area {
			case AreaPE:
				err = client.AGReadEB(group.start, group.end-group.start, buffer)
			case AreaPA:
				err = client.AGReadAB(group.start, group.end-group.start, buffer)
			case AreaMK:
				err = client.AGReadMB(group.start, group.end-group.start, buffer)
			case AreaDB:
				err = client.AGReadDB(group.db, group.start, group.end-group.start, buffer)
			case AreaCT:
				err = client.AGReadCT(group.start, group.end-group.start, buffer)
			case AreaTM:
				err = client.AGReadTM(group.start, group.end-group.start, buffer)
			default:
				err = fmt.Errorf("unsupport area: %v", group.area)
			}

			if err != nil {
				s.conn.Logger().Sugar().Errorf("read error: area: %v, db: %v, start: %v, size: %v, err: %v",
					group.area, group.db, group.start, len(buffer), err)

				return err
			}

			for _, tag := range group.tagList {
				addr := GetAddr(tag)

				bytes := buffer[addr.Addr-group.start : (addr.Addr-group.start)+addr.Size]

				value, err := convertBytesToValue(&s.config, tag, bytes)
				if err != nil {
					s.conn.Logger().Sugar().Error(err)
					continue
				}

				s.conn.SyncTagValue(tag.Raw.GetId(), value)

				if s.config.Debug {
					s.conn.Logger().Sugar().Debugf("read tag: %v %v %v %v -> %v", tag.Raw.GetId(), tag.Raw.GetName(), tag.Raw.GetAddress(), bytes, value)
				}
			}
		}
	}

	return nil
}

func (s *GoS7) readTag(tag source.Tag) (value string, err error) {
	defer func() {
		if re := recover(); re != nil {
			err = errors.New(fmt.Sprint(re))
		}
	}()

	s.lock.Lock()
	defer s.lock.Unlock()

	addr := GetAddr(&tag)

	var buffer = make([]byte, addr.Size)

	client := gos7.NewClient(s.client)

	switch addr.Area {
	case AreaPE:
		err = client.AGReadEB(addr.Addr, addr.Size, buffer)
	case AreaPA:
		err = client.AGReadAB(addr.Addr, addr.Size, buffer)
	case AreaMK:
		err = client.AGReadMB(addr.Addr, addr.Size, buffer)
	case AreaDB:
		err = client.AGReadDB(addr.DB, addr.Addr, addr.Size, buffer)
	case AreaCT:
		err = client.AGReadCT(addr.Addr, addr.Size, buffer)
	case AreaTM:
		err = client.AGReadTM(addr.Addr, addr.Size, buffer)
	default:
		err = fmt.Errorf("unsupport area: %v", addr.Area)
	}

	if err != nil {
		s.conn.Logger().Sugar().Errorf("read error: id: %v, area: %v, db: %v, start: %v, size: %v, err: %v",
			tag.Raw.GetId(), addr.Area, addr.DB, addr.Addr, addr.Size, err)
		return "", err
	}

	value, err = convertBytesToValue(&s.config, &tag, buffer)
	if err != nil {
		s.conn.Logger().Sugar().Error(err)
		return "", err
	}

	if s.config.Debug {
		s.conn.Logger().Sugar().Debugf("read tag: %v %v %v %v -> %v", tag.Raw.GetId(), tag.Raw.GetName(), tag.Raw.GetAddress(), buffer, value)
	}

	return value, nil
}

func groupTags(tags map[string]source.Tag) map[string][]readGroup {
	tags2 := make([]*source.Tag, 0, len(tags))

	for _, tag := range tags {
		tags2 = append(tags2, &tag)
	}

	sort.Sort(tagsSort(tags2))

	return group(tags2)
}

// sort
type tagsSort []*source.Tag

func (s tagsSort) Len() int {
	return len(s)
}

func (s tagsSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s tagsSort) Less(i, j int) bool {
	return GetAddr(s[i]).Addr < GetAddr(s[j]).Addr
}

type readGroup struct {
	area    int
	db      int
	start   int
	end     int
	tagList []*source.Tag
}

func newReadGroup(tag *source.Tag, addr *Addr) readGroup {
	group := readGroup{
		area:    addr.Area,
		db:      addr.DB,
		start:   addr.Addr,
		end:     addr.Addr + addr.Size,
		tagList: []*source.Tag{tag},
	}

	return group
}

const block = 256

func group(tags []*source.Tag) map[string][]readGroup {
	groupMap := make(map[string][]readGroup)

	for _, tag := range tags {
		addr := GetAddr(tag)

		key := fmt.Sprintf("%v:%v", addr.Area, addr.DB)

		if groups, ok := groupMap[key]; ok {
			if (addr.Addr + addr.Size - groups[len(groups)-1].start) < block {
				groups[len(groups)-1].end = addr.Addr + addr.Size
				groups[len(groups)-1].tagList = append(groups[len(groups)-1].tagList, tag)
			} else {
				group := newReadGroup(tag, addr)

				groupMap[key] = append(groupMap[key], group)
			}
		} else {
			group := newReadGroup(tag, addr)

			groupMap[key] = []readGroup{group}
		}
	}

	return groupMap
}
