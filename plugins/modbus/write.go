package modbus

import (
	"errors"
	"fmt"

	"github.com/snple/kirara/plugins/source"
	"github.com/snple/kirara/util/datatype"
)

func (s *Modbus) writeTag(tag source.Tag, value string) (err error) {
	defer func() {
		if re := recover(); re != nil {
			err = errors.New(fmt.Sprint(re))
		}
	}()

	s.lock.Lock()
	defer s.lock.Unlock()

	addr := GetAddr(&tag)

	switch addr.Area {
	case "0":
		if tag.Raw.GetDataType() == string(datatype.DataTypeBool) {
			if s.config.Debug {
				s.conn.Logger().Sugar().Debugf("write tag: %v %v %v -> %v", tag.Raw.GetId(), tag.Raw.GetName(), tag.Raw.GetAddress(), value)
			}

			if value == "true" || value == "1" {
				return s.client.WriteCoil(addr.Addr, true)
			} else if value == "false" || value == "0" {
				return s.client.WriteCoil(addr.Addr, false)
			}
		}
	case "4":
		if tag.Raw.GetDataType() == string(datatype.DataTypeBool) ||
			tag.Raw.GetDataType() == string(datatype.DataTypeU8) ||
			tag.Raw.GetDataType() == string(datatype.DataTypeI8) ||
			tag.Raw.GetDataType() == string(datatype.DataTypeString) {
			// unsupport
		} else {
			bytes, err := convertValueToBytes(&s.config, &tag, value)
			if err != nil {
				s.conn.Logger().Sugar().Errorf("convertValueToBytes err: %s, tag: %+v", err, tag.Raw)
				return err
			}

			if int(addr.Size) != len(bytes) {
				s.conn.Logger().Sugar().Errorf("addr.Size(%v) != len(bytes)(%v), tag: %+v, ", addr.Size, len(bytes), tag.Raw)
				return fmt.Errorf("addr.Size(%v) != len(bytes)(%v), tag: %+v, ", addr.Size, len(bytes), tag.Raw)
			}

			if s.config.Debug {
				s.conn.Logger().Sugar().Debugf("write tag: %v %v %v %v -> %v", tag.Raw.GetId(), tag.Raw.GetName(), tag.Raw.GetAddress(), bytes, value)
			}

			return s.client.WriteRawBytes(addr.Addr, bytes)
		}
	}

	return nil
}
