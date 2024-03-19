package gos7

import (
	"errors"
	"fmt"

	"github.com/robinson/gos7"
	utilbytes "github.com/snple/kirara/util/bytes"
	"github.com/snple/kirara/util/datatype"
)

func (c *Conn) writeTag(id string, value string) (err error) {
	defer func() {
		if re := recover(); re != nil {
			err = errors.New(fmt.Sprint(re))
		}
	}()

	c.lock.Lock()
	defer c.lock.Unlock()

	if tag, ok := c.tags[id]; ok {
		client := gos7.NewClient(c.client)

		var writeBytes []byte
		if datatype.DataType(tag.raw.DataType) == datatype.DataTypeBool {
			bytes := make([]byte, tag.addr.Size)

			switch tag.addr.Area {
			case AreaPE:
				err = client.AGReadEB(tag.addr.Address, tag.addr.Size, bytes)
			case AreaPA:
				err = client.AGReadAB(tag.addr.Address, tag.addr.Size, bytes)
			case AreaMK:
				err = client.AGReadMB(tag.addr.Address, tag.addr.Size, bytes)
			case AreaDB:
				err = client.AGReadDB(tag.addr.DB, tag.addr.Address, tag.addr.Size, bytes)
			case AreaCT:
				err = client.AGReadCT(tag.addr.Address, tag.addr.Size, bytes)
			case AreaTM:
				err = client.AGReadTM(tag.addr.Address, tag.addr.Size, bytes)
			default:
				err = fmt.Errorf("unsupport area: %v", tag.addr.Area)
			}

			if err != nil {
				c.gs.logger().Sugar().Errorf("read error: id: %v, area: %v, db: %v, start: %v, size: %v, err: %v",
					id, tag.addr.Area, tag.addr.DB, tag.addr.Address, tag.addr.Size, err)
				return err
			}

			boolValue := false
			if value == "true" || value == "1" {
				boolValue = true
			}

			bytes, err = func(bytes []byte) ([]byte, error) {
				switch tag.addr.Format {
				case "B", "C", "X", "I", "DBB", "DBC", "DBX":
					if len(bytes) == 1 {
						bytes[0] = utilbytes.SetBitFromBites(bytes[0], tag.addr.Bit, boolValue)

						return bytes, nil
					}
				case "W", "DBW":
					if len(bytes) == 2 {
						i := tag.addr.Bit / 8
						bit := tag.addr.Bit % 8

						if i <= 1 {
							bytes[i] = utilbytes.SetBitFromBites(bytes[i], bit, boolValue)

							return bytes, nil
						}
					}
				case "D", "DBD", "DI", "DBDI", "REAL", "DBREAL":
					if len(bytes) == 4 {
						i := tag.addr.Bit / 8
						bit := tag.addr.Bit % 8

						if i <= 3 {
							bytes[i] = utilbytes.SetBitFromBites(bytes[i], bit, boolValue)

							return bytes, nil
						}
					}
				}

				return bytes, fmt.Errorf("set bit error when write, tag: %+v, bytes: %v", tag.raw, bytes)
			}(bytes)

			if err != nil {
				return err
			}

			writeBytes = bytes
		} else {
			bytes, err := convertValueToBytes(&c.config, tag, value)
			if err != nil {
				c.gs.logger().Sugar().Errorf("convertValueToBytes err: %s, tag: %+v", err, tag.raw)
				return err
			}

			if tag.addr.Size != len(bytes) {
				c.gs.logger().Sugar().Errorf("tag.addr.Size(%v) != len(bytes)(%v), tag: %+v, ", tag.addr.Size, len(bytes), tag.raw)
				return fmt.Errorf("tag.addr.Size(%v) != len(bytes)(%v), tag: %+v, ", tag.addr.Size, len(bytes), tag.raw)
			}

			writeBytes = bytes
		}

		if c.config.Debug {
			c.gs.logger().Sugar().Debugf("write tag: %v %v %v %v -> %v", tag.raw.Id, tag.raw.Name, tag.raw.Address, writeBytes, value)
		}

		if len(writeBytes) > 0 {
			switch tag.addr.Area {
			case AreaPE:
				err = client.AGWriteEB(tag.addr.Address, tag.addr.Size, writeBytes)
			case AreaPA:
				err = client.AGWriteAB(tag.addr.Address, tag.addr.Size, writeBytes)
			case AreaMK:
				err = client.AGWriteMB(tag.addr.Address, tag.addr.Size, writeBytes)
			case AreaDB:
				err = client.AGWriteDB(tag.addr.DB, tag.addr.Address, tag.addr.Size, writeBytes)
			case AreaCT:
				err = client.AGWriteCT(tag.addr.Address, tag.addr.Size, writeBytes)
			case AreaTM:
				err = client.AGWriteTM(tag.addr.Address, tag.addr.Size, writeBytes)
			default:
				err = fmt.Errorf("unsupport area: %v", tag.addr.Area)
			}

			if err != nil {
				c.gs.logger().Sugar().Errorf("write error: id: %v, area: %v, db: %v, start: %v, size: %v, err: %v",
					id, tag.addr.Area, tag.addr.DB, tag.addr.Address, tag.addr.Size, err)

				return err
			}
		}

		return nil
	}

	return fmt.Errorf("tag: %v is not found", id)
}
