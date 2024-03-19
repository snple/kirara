package gos7

import (
	"fmt"
	"strconv"

	"github.com/snple/kirara/util"
	utilbytes "github.com/snple/kirara/util/bytes"
	"github.com/snple/kirara/util/datatype"
)

func convertBytesToValue(config *Config, tag *Tag, bytes []byte) (string, error) {
	switch datatype.DataType(tag.raw.DataType) {
	case datatype.DataTypeI8:
		if len(bytes) == 1 {
			v := int8(bytes[0])
			return fmt.Sprintf("%v", v), nil
		}
	case datatype.DataTypeU8:
		if len(bytes) == 1 {
			v := uint8(bytes[0])
			return fmt.Sprintf("%v", v), nil
		}
	case datatype.DataTypeI16:
		if len(bytes) == 2 {
			v, _ := utilbytes.BytesToUInt16(bytes, config.IsBIGEndian)
			return fmt.Sprintf("%v", int16(v)), nil
		}
	case datatype.DataTypeU16:
		if len(bytes) == 2 {
			v, _ := utilbytes.BytesToUInt16(bytes, config.IsBIGEndian)
			return fmt.Sprintf("%v", v), nil
		}
	case datatype.DataTypeI32:
		if len(bytes) == 4 {
			v, _ := utilbytes.BytesToUInt32(bytes, config.IsBIGEndian)
			return fmt.Sprintf("%v", int32(v)), nil
		}
	case datatype.DataTypeU32:
		if len(bytes) == 4 {
			v, _ := utilbytes.BytesToUInt32(bytes, config.IsBIGEndian)
			return fmt.Sprintf("%v", v), nil
		}
	case datatype.DataTypeI64:
		if len(bytes) == 8 {
			v, _ := utilbytes.BytesToUInt64(bytes, config.IsBIGEndian)
			return fmt.Sprintf("%v", int64(v)), nil
		}
	case datatype.DataTypeU64:
		if len(bytes) == 8 {
			v, _ := utilbytes.BytesToUInt64(bytes, config.IsBIGEndian)
			return fmt.Sprintf("%v", v), nil
		}
	case datatype.DataTypeF32:
		if len(bytes) == 4 {
			v, _ := utilbytes.BytesToFloat32(bytes, config.IsBIGEndian)
			v = float32(util.Round(float64(v), 3))
			return fmt.Sprintf("%v", v), nil
		}
	case datatype.DataTypeF64:
		if len(bytes) == 8 {
			v, _ := utilbytes.BytesToFloat64(bytes, config.IsBIGEndian)
			v = util.Round(v, 3)
			return fmt.Sprintf("%v", v), nil
		}
	case datatype.DataTypeBool:
		switch tag.addr.Format {
		case "B", "C", "X", "I", "DBB", "DBC", "DBX":
			if len(bytes) == 1 {
				b := utilbytes.GetBitFromBites(bytes[0], tag.addr.Bit)
				return fmt.Sprintf("%v", b), nil
			}
		case "W", "DBW":
			if len(bytes) == 2 {
				i := tag.addr.Bit / 8
				bit := tag.addr.Bit % 8

				if i <= 1 {
					b := utilbytes.GetBitFromBites(bytes[i], bit)
					return fmt.Sprintf("%v", b), nil
				}
			}
		case "D", "DBD", "DI", "DBDI", "REAL", "DBREAL":
			if len(bytes) == 4 {
				i := tag.addr.Bit / 8
				bit := tag.addr.Bit % 8

				if i <= 3 {
					b := utilbytes.GetBitFromBites(bytes[i], bit)
					return fmt.Sprintf("%v", b), nil
				}
			}
		}
	}

	return "", fmt.Errorf("convertBytesToNson, tag: %+v, bytes: %v", tag.raw, bytes)
}

func convertValueToBytes(config *Config, tag *Tag, value string) ([]byte, error) {
	switch datatype.DataType(tag.raw.DataType) {
	case datatype.DataTypeI8:
		value2, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}

		return []byte{uint8(value2)}, nil
	case datatype.DataTypeU8:
		value2, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}

		return []byte{uint8(value2)}, nil
	case datatype.DataTypeI16:
		value2, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}

		return utilbytes.UInt16ToBytes(uint16(value2), config.IsBIGEndian), nil
	case datatype.DataTypeU16:
		value2, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}

		return utilbytes.UInt16ToBytes(uint16(value2), config.IsBIGEndian), nil
	case datatype.DataTypeI32:
		value2, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}

		return utilbytes.UInt32ToBytes(uint32(value2), config.IsBIGEndian), nil
	case datatype.DataTypeU32:
		value2, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}

		return utilbytes.UInt32ToBytes(uint32(value2), config.IsBIGEndian), nil
	case datatype.DataTypeI64:
		value2, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}

		return utilbytes.UInt64ToBytes(uint64(value2), config.IsBIGEndian), nil
	case datatype.DataTypeU64:
		value2, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}

		return utilbytes.UInt64ToBytes(uint64(value2), config.IsBIGEndian), nil
	case datatype.DataTypeF32:
		value2, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}

		return utilbytes.Float32ToBytes(float32(value2), config.IsBIGEndian), nil
	case datatype.DataTypeF64:
		value2, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}

		return utilbytes.Float64ToBytes(float64(value2), config.IsBIGEndian), nil
	}

	return nil, fmt.Errorf("convertValueToBytes, tag: %+v, value: %v", tag.raw, value)
}
