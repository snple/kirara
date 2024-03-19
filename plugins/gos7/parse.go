package gos7

import (
	"strconv"
	"strings"
)

type S7Addr struct {
	Addr    string
	Format  string
	Area    int
	DB      int
	Address int
	Size    int
	Bit     int
}

var typeSize = map[string]int{
	"B":      1,
	"DBB":    1,
	"C":      1,
	"DBC":    1,
	"X":      1,
	"DBX":    1,
	"I":      2,
	"DBI":    2,
	"W":      2,
	"DBW":    2,
	"D":      4,
	"DBD":    4,
	"DI":     4,
	"DBDI":   4,
	"REAL":   4,
	"DBREAL": 4,
}

const (
	AreaPE = 0x81
	AreaPA = 0x82
	AreaMK = 0x83
	AreaDB = 0x84
	AreaCT = 0x1C
	AreaTM = 0x1D
)

func ParseTagAddress(addr string) S7Addr {
	var rs S7Addr
	rs.Addr = addr

	var subaddress string
	// var DB_No = 0
	if strings.HasPrefix(addr, "DB") {
		rs.Area = AreaDB
		//Parse DBnn,xxnn or DBnn.xxnn
		block_end := -1
		for i, s := range addr {
			if s == '.' || s == ',' {
				block_end = i
				break
			}
		}

		rs.DB, _ = strconv.Atoi(addr[2:block_end])

		subaddress = addr[block_end+1:]

	} else if strings.HasPrefix(addr, "PI") || strings.HasPrefix(addr, "PE") {
		rs.Area = AreaPE
		subaddress = addr[2:]
		//PE area
	} else if strings.HasPrefix(addr, "PQ") || strings.HasPrefix(addr, "PA") {
		rs.Area = AreaPA
		subaddress = addr[2:]
		//PA area
	} else {
		if strings.HasPrefix(addr, "I") || strings.HasPrefix(addr, "E") {
			rs.Area = AreaPE
		} else if strings.HasPrefix(addr, "Q") || strings.HasPrefix(addr, "A") {
			rs.Area = AreaPA
		} else if strings.HasPrefix(addr, "M") || strings.HasPrefix(addr, "F") {
			rs.Area = AreaMK
		} else if strings.HasPrefix(addr, "C") || strings.HasPrefix(addr, "Z") {
			rs.Area = AreaCT
		} else if strings.HasPrefix(addr, "T") {
			rs.Area = AreaTM
		} else if strings.HasPrefix(addr, "V") {
			rs.Area = AreaDB
			rs.DB = 1
		}

		subaddress = addr[1:]
	}

	rs.Format, rs.Address, rs.Bit = ParseTagSubAddress(subaddress)
	rs.Size = typeSize[rs.Format]
	//rs.Cmd = [4]int{S7Area[rs.Area], DB_No, rs.Address, rs.Size}
	return rs
}

func ParseTagSubAddress(subaddr string) (string, int, int) {
	format := ""

	digitStart := -1
	bitStart := 0

	for i, s := range subaddr {
		if s >= 48 && s <= 57 && digitStart == -1 {
			digitStart = i
		}
		if s == '.' {
			bitStart = i + 1
		}

	}

	if digitStart == 0 {
		format = "X"
	} else {
		format = strings.ToUpper(subaddr[0:digitStart])
	}

	address := -1
	bit := -1

	if bitStart == 0 {
		address, _ = strconv.Atoi(subaddr[digitStart:])
	} else {
		address, _ = strconv.Atoi(subaddr[digitStart : bitStart-1])
		bit, _ = strconv.Atoi(subaddr[bitStart:])

	}

	return format, address, bit
}

/*
func Substr(str string, start int, end int) string {
	return string([]rune(str)[start:end])
}

func main() {
	s := "DB1,X0.1"
	res := ParseTagAddress("DB1,X0.1")
	fmt.Println(res)
	fmt.Println(Substr("DB1,X0.1", 2, 6))
	fmt.Println(s[2:3])
	fmt.Println(ParseTagSubAddress("111.0"))
	fmt.Println(ParseTagSubAddress("B0.2"))
	fmt.Println(ParseTagSubAddress("REAL23"))
	fmt.Println(ParseTagAddress("DB1,X0.1"))
	fmt.Println(ParseTagAddress("I0.1"))
	fmt.Println(ParseTagAddress("PIW12"))
	fmt.Println(ParseTagAddress("MD23"))
	fmt.Println(ParseTagAddress("M2.7"))
	fmt.Println(ParseTagAddress("MW0.14"))
	fmt.Println(ParseTagAddress("MREAL44"))
	fmt.Println(ParseTagAddress("DB1,REAL24"))
}
*/
