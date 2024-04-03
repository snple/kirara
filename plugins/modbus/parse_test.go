package modbus

import (
	"fmt"
	"testing"

	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/plugins/source"
	"github.com/stretchr/testify/assert"
)

func TestParseAddr(t *testing.T) {
	addr := "40033.4"

	mbaddr, err := ParseTagAddr(addr, "BOOL")

	fmt.Printf("%+v\n", mbaddr)
	assert.Exactly(t, err, nil)
	assert.Exactly(t, "4", mbaddr.Area)
	assert.Exactly(t, uint8(4), mbaddr.Bit)
	assert.Exactly(t, uint16(33), mbaddr.Addr)
	assert.Exactly(t, uint16(2), mbaddr.Size)
}

func TestHoldingGroup(t *testing.T) {
	tags := make(map[string]source.Tag, 0)

	{
		tag := source.Tag{
			Raw: &pb.Tag{Name: "tag1", DataType: "I32", Address: "40000"},
		}

		addr, err := ParseTagAddr(tag.Raw.GetAddress(), tag.Raw.GetDataType())
		assert.Exactly(t, err, nil)

		SetAddr(&tag, addr)

		// tags["1"] = tag
	}

	{
		tag := source.Tag{
			Raw: &pb.Tag{Name: "tag2", DataType: "I32", Address: "40004"},
		}

		addr, err := ParseTagAddr(tag.Raw.GetAddress(), tag.Raw.GetDataType())
		assert.Exactly(t, err, nil)

		SetAddr(&tag, addr)

		tags["2"] = tag
	}

	{
		tag := source.Tag{
			Raw: &pb.Tag{Name: "tag3", DataType: "I32", Address: "40008"},
		}

		addr, err := ParseTagAddr(tag.Raw.GetAddress(), tag.Raw.GetDataType())
		assert.Exactly(t, err, nil)

		SetAddr(&tag, addr)

		tags["3"] = tag
	}

	ret := groupTags(tags)

	group := ret["4"][0]
	assert.Exactly(t, uint16(4), group.addr)
	assert.Exactly(t, uint16(16), group.end)
}

func TestColisGroup(t *testing.T) {
	tags := make(map[string]source.Tag, 0)

	{
		tag := source.Tag{
			Raw: &pb.Tag{Name: "tag1", DataType: "BOOL", Address: "00000"},
		}

		addr, err := ParseTagAddr(tag.Raw.GetAddress(), tag.Raw.GetDataType())
		assert.Exactly(t, err, nil)

		SetAddr(&tag, addr)

		// tags["1"] = tag
		fmt.Printf("addr: %+v\n", addr)
	}

	{
		tag := source.Tag{
			Raw: &pb.Tag{Name: "tag2", DataType: "BOOL", Address: "00002"},
		}

		addr, err := ParseTagAddr(tag.Raw.GetAddress(), tag.Raw.GetDataType())
		assert.Exactly(t, err, nil)

		SetAddr(&tag, addr)

		tags["2"] = tag
		fmt.Printf("addr: %+v\n", addr)
	}

	{
		tag := source.Tag{
			Raw: &pb.Tag{Name: "tag3", DataType: "BOOL", Address: "00004"},
		}

		addr, err := ParseTagAddr(tag.Raw.GetAddress(), tag.Raw.GetDataType())
		assert.Exactly(t, err, nil)

		SetAddr(&tag, addr)

		tags["3"] = tag
		fmt.Printf("addr: %+v\n", addr)
	}

	ret := groupTags(tags)

	group := ret["0"][0]
	fmt.Printf("%+v\n", group)
	assert.Exactly(t, uint16(2), group.addr)
	assert.Exactly(t, uint16(5), group.end)
}
