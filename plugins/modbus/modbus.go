package modbus

import (
	"errors"
	"sync"
	"time"

	"github.com/simonvetter/modbus"
	"github.com/snple/kirara/plugins/source"
)

func Connect(conn *source.Conn) (source.Adapter, error) {
	config, err := ParseConfig(conn.Source().GetParams())
	if err != nil {
		return nil, err
	}

	var client *modbus.ModbusClient
	if config.isTCP {
		client, err = modbus.NewClient(&modbus.ClientConfiguration{
			URL:     config.Addr,
			Timeout: time.Duration(config.Timeout) * time.Second,
		})
		if err != nil {
			return nil, err
		}
	} else if config.isRTU {
		panic("todo")
	} else {
		return nil, errors.New("only support tcp and rtu")
	}

	err = client.Open()
	if err != nil {
		return nil, err
	}

	s := &Modbus{
		conn:   conn,
		client: client,
		config: config,
	}

	return s, nil
}

func GetAddr(tag *source.Tag) *Addr {
	return tag.Metadata.(*Addr)
}

func SetAddr(tag *source.Tag, addr *Addr) {
	tag.Metadata = addr
}

type Modbus struct {
	conn   *source.Conn
	client *modbus.ModbusClient

	config Config

	lock sync.Mutex
}

func (s *Modbus) Close() error {
	s.client.Close()
	return nil
}

func (s *Modbus) Name() string {
	return "Modbus"
}

func (s *Modbus) ParseTag(tag *source.Tag) error {
	addr, err := ParseTagAddr(tag.Raw.GetAddress(), tag.Raw.GetDataType())
	if err != nil {
		return err
	}

	SetAddr(tag, addr)

	return nil
}

func (s *Modbus) ReadTags(tags map[string]source.Tag) error {
	return s.readTags(tags)
}

func (s *Modbus) ReadTag(tag source.Tag) (string, error) {
	return s.readTag(tag)
}

func (s *Modbus) WriteTag(tag source.Tag, value string) error {
	return s.writeTag(tag, value)
}
