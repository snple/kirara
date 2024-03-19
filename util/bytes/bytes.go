package bytes

import (
	"encoding/binary"
	"fmt"
	"math"
)

func GetBitFromBites(b byte, bit int) bool {
	return (b & (1 << uint32(bit))) > 0
}

func SetBitFromBites(b byte, bit int, io bool) byte {
	if io {
		b = b | (1 << uint32(bit))
	} else {
		b = b &^ (1 << uint32(bit))
	}

	return b
}

// bytes to uint 16
func BytesToUInt16(b []byte, isBig bool) (uint16, error) {
	if len(b) != 2 {
		return 0, fmt.Errorf("[]byte(%v) len must == 2", b)
	}

	if isBig {
		return binary.BigEndian.Uint16(b), nil
	}

	return binary.LittleEndian.Uint16(b), nil
}

// bytes to uint 32
func BytesToUInt32(b []byte, isBig bool) (uint32, error) {
	if len(b) != 4 {
		return 0, fmt.Errorf("[]byte(%v) len must == 4", b)
	}

	if isBig {
		return binary.BigEndian.Uint32(b), nil
	}

	return binary.LittleEndian.Uint32(b), nil
}

// bytes to uint 64
func BytesToUInt64(b []byte, isBig bool) (uint64, error) {
	if len(b) != 8 {
		return 0, fmt.Errorf("[]byte(%v) len must == 8", b)
	}

	if isBig {
		return binary.BigEndian.Uint64(b), nil
	}

	return binary.LittleEndian.Uint64(b), nil
}

// bytes to float32
func BytesToFloat32(b []byte, isBig bool) (float32, error) {
	tmp, err := BytesToUInt32(b, isBig)
	if err != nil {
		return 0, err
	}

	return math.Float32frombits(tmp), nil
}

// bytes to float64
func BytesToFloat64(b []byte, isBig bool) (float64, error) {
	tmp, err := BytesToUInt64(b, isBig)
	if err != nil {
		return 0, err
	}

	return math.Float64frombits(tmp), nil
}

// int to 8 bytes
func UInt64ToBytes(i uint64, isBig bool) []byte {
	bytes := make([]byte, 8)
	if isBig {
		binary.BigEndian.PutUint64(bytes, i)
	} else {
		binary.LittleEndian.PutUint64(bytes, i)
	}
	return bytes
}

// int to 4 bytes
func UInt32ToBytes(i uint32, isBig bool) []byte {
	bytes := make([]byte, 4)
	if isBig {
		binary.BigEndian.PutUint32(bytes, i)
	} else {
		binary.LittleEndian.PutUint32(bytes, i)
	}
	return bytes
}

// int to 2 bytes
func UInt16ToBytes(i uint16, isBig bool) []byte {
	bytes := make([]byte, 2)
	if isBig {
		binary.BigEndian.PutUint16(bytes, i)
	} else {
		binary.LittleEndian.PutUint16(bytes, i)
	}
	return bytes
}

// Float64 to 8 bytes
func Float64ToBytes(f float64, isBig bool) []byte {
	tmp := math.Float64bits(f)
	return UInt64ToBytes(tmp, isBig)
}

// Float32 to 4 bytes
func Float32ToBytes(f float32, isBig bool) []byte {
	tmp := math.Float32bits(f)
	return UInt32ToBytes(tmp, isBig)
}
