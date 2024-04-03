package modbus

import (
	"strconv"
)

type Addr struct {
	Area string
	Name string
	Addr uint16
	Size uint16
	Bit  uint8
}

var type_size = map[string]uint16{
	"BOOL": 2,
	"I16":  2,
	"U16":  2,
	"I32":  4,
	"U32":  4,
	"I64":  8,
	"U64":  8,
	"F32":  4,
	"F64":  8,
	// "I8":   1,
	// "U8":   1,
}

func ParseTagAddr(addr string, dataType string) (*Addr, error) {
	var rs Addr
	rs.Name = addr
	var subaddress string
	rs.Area = addr[0:1]
	subaddress = addr[1:]
	rs.Addr, rs.Bit = ParseTagSubAddr(subaddress)
	rs.Size = type_size[dataType]

	if rs.Area == "0" || rs.Area == "1" {
		rs.Size = 1
	}

	return &rs, nil
}

func ParseTagSubAddr(subaddr string) (uint16, uint8) {
	bit_start := 0

	for i, s := range subaddr {
		if s == '.' {
			bit_start = i + 1
		}

	}
	address := uint16(0)
	bit := uint8(0)

	if bit_start == 0 {
		temp, _ := strconv.Atoi(subaddr)
		address = uint16(temp)
	} else {
		temp, _ := strconv.Atoi(subaddr[0 : bit_start-1])
		address = uint16(temp)
		temp2, _ := strconv.Atoi(subaddr[bit_start:])
		bit = uint8(temp2)

	}

	return address, bit
}

func unpackBits(bytes []byte, count uint16) []bool {
	ret := make([]bool, count)

	for i := 0; i < int(count); i++ {
		if (bytes[(uint32(i)/8)]>>(uint32(i)%8))&1 > 0 {
			ret[i] = true
		}
	}

	return ret
}
