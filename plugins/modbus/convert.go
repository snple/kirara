package modbus

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/snple/kirara/plugins/source"
	utilbytes "github.com/snple/kirara/util/bytes"
	"github.com/snple/kirara/util/datatype"
)

type TagConfig struct {
	Convert *Convert `json:"convert"`
}

type Convert struct {
	Type string `json:"type"`
	Set  string `json:"set"`
	Get  string `json:"get"`
}

func ParseTagConfig(params string) (TagConfig, error) {
	config := TagConfig{}

	err := json.Unmarshal([]byte(params), &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func convertBytesToValue(config *Config, tag *source.Tag, bytes []byte) (string, error) {
	dtype := datatype.DataType(tag.Raw.GetDataType())

	switch dtype {
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
			v := bytesToUint32(config.IsBIGEndian, config.IsHighWordFirst, bytes)
			return fmt.Sprintf("%v", int32(v)), nil
		}
	case datatype.DataTypeU32:
		if len(bytes) == 4 {
			v := bytesToUint32(config.IsBIGEndian, config.IsHighWordFirst, bytes)
			return fmt.Sprintf("%v", v), nil
		}
	case datatype.DataTypeI64:
		if len(bytes) == 8 {
			v := bytesToUint64(config.IsBIGEndian, config.IsHighWordFirst, bytes)
			return fmt.Sprintf("%v", int64(v)), nil
		}
	case datatype.DataTypeU64:
		if len(bytes) == 8 {
			v := bytesToUint64(config.IsBIGEndian, config.IsHighWordFirst, bytes)
			return fmt.Sprintf("%v", v), nil
		}
	case datatype.DataTypeF32:
		if len(bytes) == 4 {
			v := bytesToUint32(config.IsBIGEndian, config.IsHighWordFirst, bytes)
			return fmt.Sprintf("%v", float32(v)), nil
		}
	case datatype.DataTypeF64:
		if len(bytes) == 4 {
			v := bytesToUint64(config.IsBIGEndian, config.IsHighWordFirst, bytes)
			return fmt.Sprintf("%v", float64(v)), nil
		}
	}

	return "", fmt.Errorf("convertBytesToValue, tag: %+v, bytes: %v", tag.Raw, bytes)
}

func convertValueToBytes(config *Config, tag *source.Tag, value string) ([]byte, error) {
	dtype := datatype.DataType(tag.Raw.GetDataType())

	switch dtype {
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
		return uint32ToBytes(config.IsBIGEndian, config.IsHighWordFirst, uint32(value2)), nil
	case datatype.DataTypeU32:
		value2, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}
		return uint32ToBytes(config.IsBIGEndian, config.IsHighWordFirst, uint32(value2)), nil
	case datatype.DataTypeI64:
		value2, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}
		return uint64ToBytes(config.IsBIGEndian, config.IsHighWordFirst, uint64(value2)), nil
	case datatype.DataTypeU64:
		value2, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}
		return uint64ToBytes(config.IsBIGEndian, config.IsHighWordFirst, uint64(value2)), nil
	case datatype.DataTypeF32:
		value2, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}
		return uint32ToBytes(config.IsBIGEndian, config.IsHighWordFirst, uint32(value2)), nil
	case datatype.DataTypeF64:
		value2, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, fmt.Errorf("DecodeValue: %s", err)
		}
		return uint64ToBytes(config.IsBIGEndian, config.IsHighWordFirst, uint64(value2)), nil
	}

	return nil, fmt.Errorf("convertValueToBytes, tag: %+v, value: %v", tag.Raw, value)
}

func bytesToUint32(bigEndian bool, highWordFirst bool, in []byte) uint32 {
	var u32 uint32

	if bigEndian {
		if highWordFirst {
			u32 = binary.BigEndian.Uint32(in)
		} else {
			u32 = binary.BigEndian.Uint32([]byte{in[2], in[3], in[0], in[1]})
		}
	} else {
		if !highWordFirst {
			u32 = binary.LittleEndian.Uint32(in)
		} else {
			u32 = binary.LittleEndian.Uint32([]byte{in[2], in[3], in[0], in[1]})
		}
	}

	return u32
}

func uint32ToBytes(bigEndian bool, highWordFirst bool, in uint32) (out []byte) {
	out = make([]byte, 4)

	if bigEndian {
		binary.BigEndian.PutUint32(out, in)

		if !highWordFirst {
			out[0], out[1], out[2], out[3] = out[2], out[3], out[0], out[1]
		}
	} else {
		binary.LittleEndian.PutUint32(out, in)

		if highWordFirst {
			out[0], out[1], out[2], out[3] = out[2], out[3], out[0], out[1]
		}
	}

	return
}

func bytesToUint64(bigEndian bool, highWordFirst bool, in []byte) uint64 {
	var u64 uint64

	if bigEndian {
		if highWordFirst {
			u64 = binary.BigEndian.Uint64(in)
		} else {
			u64 = binary.BigEndian.Uint64(
				[]byte{in[6], in[7], in[4], in[5],
					in[2], in[3], in[0], in[1]})
		}
	} else {
		if !highWordFirst {
			u64 = binary.LittleEndian.Uint64(in)
		} else {
			u64 = binary.LittleEndian.Uint64(
				[]byte{in[6], in[7], in[4], in[5],
					in[2], in[3], in[0], in[1]})
		}
	}

	return u64
}

func uint64ToBytes(bigEndian bool, highWordFirst bool, in uint64) (out []byte) {
	out = make([]byte, 8)

	if bigEndian {
		binary.BigEndian.PutUint64(out, in)

		if !highWordFirst {
			out[0], out[1], out[2], out[3], out[4], out[5], out[6], out[7] =
				out[6], out[7], out[4], out[5], out[2], out[3], out[0], out[1]
		}
	} else {
		binary.LittleEndian.PutUint64(out, in)

		if highWordFirst {
			out[0], out[1], out[2], out[3], out[4], out[5], out[6], out[7] =
				out[6], out[7], out[4], out[5], out[2], out[3], out[0], out[1]
		}
	}

	return
}

func convertToFloat(o interface{}) (float64, bool) {
	if i, ok := o.(float64); ok {
		return i, true
	}
	if b, ok := o.(bool); ok {
		if b {
			return 1, true
		}
		return 0, true
	}

	v := reflect.ValueOf(o)
	for o != nil && v.Kind() == reflect.Ptr {
		v = v.Elem()
		if !v.IsValid() {
			return 0, false
		}
		o = v.Interface()
	}

	if o == false || o == nil || o == "false" || o == "FALSE" {
		return 0, true
	}
	if o == true || o == "true" || o == "TRUE" {
		return 1, true
	}

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint()), true
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	}
	if s, ok := o.(string); ok {
		f, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return f, true
		}
	}
	return 0, false
}
