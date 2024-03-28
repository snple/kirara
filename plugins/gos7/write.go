package gos7

import (
	"errors"
	"fmt"

	"github.com/robinson/gos7"
	"github.com/snple/kirara/plugins/source"
	utilbytes "github.com/snple/kirara/util/bytes"
	"github.com/snple/kirara/util/datatype"
)

func (s *GoS7) writeTag(tag source.Tag, value string) (err error) {
	defer func() {
		if re := recover(); re != nil {
			err = errors.New(fmt.Sprint(re))
		}
	}()

	s.lock.Lock()
	defer s.lock.Unlock()

	client := gos7.NewClient(s.client)

	addr := GetAddr(&tag)

	var writeBytes []byte
	if datatype.DataType(tag.Raw.GetDataType()) == datatype.DataTypeBool {
		bytes := make([]byte, addr.Size)

		switch addr.Area {
		case AreaPE:
			err = client.AGReadEB(addr.Address, addr.Size, bytes)
		case AreaPA:
			err = client.AGReadAB(addr.Address, addr.Size, bytes)
		case AreaMK:
			err = client.AGReadMB(addr.Address, addr.Size, bytes)
		case AreaDB:
			err = client.AGReadDB(addr.DB, addr.Address, addr.Size, bytes)
		case AreaCT:
			err = client.AGReadCT(addr.Address, addr.Size, bytes)
		case AreaTM:
			err = client.AGReadTM(addr.Address, addr.Size, bytes)
		default:
			err = fmt.Errorf("unsupport area: %v", addr.Area)
		}

		if err != nil {
			s.conn.Logger().Sugar().Errorf("read error: id: %v, area: %v, db: %v, start: %v, size: %v, err: %v",
				tag.Raw.GetId(), addr.Area, addr.DB, addr.Address, addr.Size, err)
			return err
		}

		boolValue := false
		if value == "true" || value == "1" {
			boolValue = true
		}

		bytes, err = func(bytes []byte) ([]byte, error) {
			switch addr.Format {
			case "B", "C", "X", "I", "DBB", "DBC", "DBX":
				if len(bytes) == 1 {
					bytes[0] = utilbytes.SetBitFromBites(bytes[0], addr.Bit, boolValue)

					return bytes, nil
				}
			case "W", "DBW":
				if len(bytes) == 2 {
					i := addr.Bit / 8
					bit := addr.Bit % 8

					if i <= 1 {
						bytes[i] = utilbytes.SetBitFromBites(bytes[i], bit, boolValue)

						return bytes, nil
					}
				}
			case "D", "DBD", "DI", "DBDI", "REAL", "DBREAL":
				if len(bytes) == 4 {
					i := addr.Bit / 8
					bit := addr.Bit % 8

					if i <= 3 {
						bytes[i] = utilbytes.SetBitFromBites(bytes[i], bit, boolValue)

						return bytes, nil
					}
				}
			}

			return bytes, fmt.Errorf("set bit error when write, tag: %+v, bytes: %v", tag.Raw, bytes)
		}(bytes)

		if err != nil {
			return err
		}

		writeBytes = bytes
	} else {
		bytes, err := convertValueToBytes(&s.config, &tag, value)
		if err != nil {
			s.conn.Logger().Sugar().Errorf("convertValueToBytes err: %s, tag: %+v", err, tag.Raw)
			return err
		}

		if addr.Size != len(bytes) {
			s.conn.Logger().Sugar().Errorf("addr.Size(%v) != len(bytes)(%v), tag: %+v, ", addr.Size, len(bytes), tag.Raw)
			return fmt.Errorf("addr.Size(%v) != len(bytes)(%v), tag: %+v, ", addr.Size, len(bytes), tag.Raw)
		}

		writeBytes = bytes
	}

	if s.config.Debug {
		s.conn.Logger().Sugar().Debugf("write tag: %v %v %v %v -> %v", tag.Raw.GetId(), tag.Raw.GetName(), tag.Raw.GetAddress(), writeBytes, value)
	}

	if len(writeBytes) > 0 {
		switch addr.Area {
		case AreaPE:
			err = client.AGWriteEB(addr.Address, addr.Size, writeBytes)
		case AreaPA:
			err = client.AGWriteAB(addr.Address, addr.Size, writeBytes)
		case AreaMK:
			err = client.AGWriteMB(addr.Address, addr.Size, writeBytes)
		case AreaDB:
			err = client.AGWriteDB(addr.DB, addr.Address, addr.Size, writeBytes)
		case AreaCT:
			err = client.AGWriteCT(addr.Address, addr.Size, writeBytes)
		case AreaTM:
			err = client.AGWriteTM(addr.Address, addr.Size, writeBytes)
		default:
			err = fmt.Errorf("unsupport area: %v", addr.Area)
		}

		if err != nil {
			s.conn.Logger().Sugar().Errorf("write error: id: %v, area: %v, db: %v, start: %v, size: %v, err: %v",
				tag.Raw.GetId(), addr.Area, addr.DB, addr.Address, addr.Size, err)

			return err
		}
	}

	return nil
}
