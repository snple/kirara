package modbus

import (
	"errors"
	"fmt"
	"sort"

	"github.com/simonvetter/modbus"
	"github.com/snple/kirara/plugins/source"
)

func (s *Modbus) readTags(tags map[string]source.Tag) (err error) {
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
			var values []bool
			var buffer []byte
			var err error

			switch group.area {
			case "0":
				values, err = s.client.ReadCoils(group.addr, group.end-group.addr)
			case "1":
				values, err = s.client.ReadDiscreteInputs(group.addr, group.end-group.addr)
			case "3":
				buffer, err = s.client.ReadRawBytes(group.addr, group.end-group.addr, modbus.INPUT_REGISTER)
			case "4":
				buffer, err = s.client.ReadRawBytes(group.addr, group.end-group.addr, modbus.HOLDING_REGISTER)
			default:
				err = fmt.Errorf("unsupport area: %v", group.area)
			}

			if err != nil {
				s.conn.Logger().Sugar().Errorf("read error: area: %v, start: %v, err: %v",
					group.area, group.addr, err)

				return err
			}

			switch group.area {
			case "0", "1":
				for _, tag := range group.tagList {
					addr := GetAddr(tag)

					value := values[addr.Addr-group.addr]

					s.conn.SyncTagValue(tag.Raw.GetId(), fmt.Sprintf("%v", value))

					if s.config.Debug {
						s.conn.Logger().Sugar().Debugf("read tag: %v %v %v -> %v", tag.Raw.GetId(), tag.Raw.GetName(), tag.Raw.GetAddress(), value)
					}
				}
			case "3", "4":
				for _, tag := range group.tagList {
					addr := GetAddr(tag)

					start := addr.Addr*2 - group.addr*2
					end := start + addr.Size
					bytes := buffer[start:end]

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
	}

	return
}

func (s *Modbus) readTag(tag source.Tag) (value string, err error) {
	defer func() {
		if re := recover(); re != nil {
			err = errors.New(fmt.Sprint(re))
		}
	}()

	s.lock.Lock()
	defer s.lock.Unlock()

	addr := GetAddr(&tag)

	var ret bool
	var buffer []byte

	switch addr.Area {
	case "0":
		ret, err = s.client.ReadCoil(addr.Addr)
	case "1":
		ret, err = s.client.ReadDiscreteInput(addr.Addr)
	case "3":
		buffer, err = s.client.ReadRawBytes(addr.Addr, addr.Size, modbus.INPUT_REGISTER)
	case "4":
		buffer, err = s.client.ReadRawBytes(addr.Addr, addr.Size, modbus.HOLDING_REGISTER)
	default:
		err = fmt.Errorf("unsupport area: %v", addr.Area)
	}

	if err != nil {
		s.conn.Logger().Sugar().Errorf("read error: id: %v, area: %v, start: %v, size: %v, err: %v",
			tag.Raw.GetId(), addr.Area, addr.Addr, addr.Size, err)
		return "", err
	}

	switch addr.Area {
	case "0", "1":
		if s.config.Debug {
			s.conn.Logger().Sugar().Debugf("read tag: %v %v %v -> %v", tag.Raw.GetId(), tag.Raw.GetName(), tag.Raw.GetAddress(), value)
		}

		return fmt.Sprintf("%v", ret), nil
	case "3", "4":
		value, err := convertBytesToValue(&s.config, &tag, buffer)
		if err != nil {
			s.conn.Logger().Sugar().Error(err)
			return "", err
		}

		if s.config.Debug {
			s.conn.Logger().Sugar().Debugf("read tag: %v %v %v %v -> %v", tag.Raw.GetId(), tag.Raw.GetName(), tag.Raw.GetAddress(), buffer, value)
		}

		return value, nil
	}

	return "", nil
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

const block = 64

func groupTags(tags map[string]source.Tag) map[string][]readGroup {
	tags2 := make([]*source.Tag, 0, len(tags))

	for _, tag := range tags {
		tags2 = append(tags2, &tag)
	}

	sort.Sort(tagsSort(tags2))

	return group2(tags2)
}

type readGroup struct {
	area    string
	addr    uint16
	end     uint16
	tagList []*source.Tag
}

func newReadGroup(tag *source.Tag, addr *Addr) readGroup {
	group := readGroup{
		area:    addr.Area,
		addr:    addr.Addr,
		tagList: []*source.Tag{tag},
	}

	if addr.Area == "0" || addr.Area == "1" {
		group.end = addr.Addr + addr.Size
	} else {
		group.end = addr.Addr*2 + addr.Size
	}

	return group
}

func group2(tags []*source.Tag) map[string][]readGroup {
	groupMap := make(map[string][]readGroup)

	for _, tag := range tags {
		addr := GetAddr(tag)

		key := addr.Area

		if groups, ok := groupMap[key]; ok {
			if addr.Area == "0" || addr.Area == "1" {
				if (addr.Addr+addr.Size-groups[len(groups)-1].addr)*2 < block {
					groups[len(groups)-1].end = addr.Addr + addr.Size
					groups[len(groups)-1].tagList = append(groups[len(groups)-1].tagList, tag)
				} else {
					group := newReadGroup(tag, addr)
					groupMap[key] = append(groupMap[key], group)
				}
			} else {
				if (addr.Addr*2 + addr.Size - groups[len(groups)-1].addr) < block {
					groups[len(groups)-1].end = addr.Addr*2 + addr.Size - groups[len(groups)-1].addr
					groups[len(groups)-1].tagList = append(groups[len(groups)-1].tagList, tag)
				} else {
					group := newReadGroup(tag, addr)
					groupMap[key] = append(groupMap[key], group)
				}
			}
		} else {
			group := newReadGroup(tag, addr)

			groupMap[key] = []readGroup{group}
		}
	}

	return groupMap
}
