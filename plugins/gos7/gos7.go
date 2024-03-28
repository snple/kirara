package gos7

import (
	"sync"

	"github.com/robinson/gos7"
	"github.com/snple/kirara/plugins/source"
)

func Connect(conn *source.Conn) (source.Adapter, error) {
	config, err := ParseConfig(conn.Source().GetParams())
	if err != nil {
		return nil, err
	}

	client := gos7.NewTCPClientHandler(config.Addr, config.Rank, config.Slot)
	err = client.Connect()
	if err != nil {
		conn.Logger().Sugar().Errorf("GoS7 source connect: %v, source: %v", err, conn.Source())
		return nil, err
	}

	s := &GoS7{
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

type GoS7 struct {
	conn   *source.Conn
	client *gos7.TCPClientHandler

	config Config

	lock sync.Mutex
}

func (s *GoS7) Close() error {
	s.client.Close()
	return nil
}

func (s *GoS7) Name() string {
	return "GoS7"
}

func (s *GoS7) ParseTag(tag *source.Tag) error {
	addr, err := ParseTagAddr(tag.Raw.GetAddress())
	if err != nil {
		return err
	}

	SetAddr(tag, addr)

	return nil
}

func (s *GoS7) ReadTags(tags map[string]source.Tag) error {
	return s.readTags(tags)
}

func (s *GoS7) ReadTag(tag source.Tag) (string, error) {
	return s.readTag(tag)
}

func (s *GoS7) WriteTag(tag source.Tag, value string) error {
	return s.writeTag(tag, value)
}
