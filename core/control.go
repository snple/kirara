package core

import (
	"context"
	"sync"

	"github.com/snple/kirara/consts"
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/cores"
	"github.com/snple/kirara/util/datatype"
	"github.com/snple/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ControlService struct {
	cs *CoreService

	clients map[string]*controlChannels
	lock    sync.RWMutex

	cores.UnimplementedControlServiceServer
}

func newControlService(cs *CoreService) *ControlService {
	return &ControlService{
		cs:      cs,
		clients: make(map[string]*controlChannels),
	}
}

func (s *ControlService) GetTagValue(ctx context.Context, in *pb.Id) (*pb.TagValue, error) {
	var err error
	var output pb.TagValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}
	}

	output.Id = in.GetId()

	// tag
	item, err := s.cs.GetTag().ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag.Status != ON")
	}

	// validation device and source
	{
		// device
		{
			device, err := s.cs.GetDevice().ViewByID(ctx, item.DeviceID)
			if err != nil {
				return &output, err
			}

			if device.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Device.Status != ON")
			}
		}

		// source
		{
			source, err := s.cs.GetSource().ViewByID(ctx, item.SourceID)
			if err != nil {
				return &output, err
			}

			if source.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Source.Status != ON")
			}
		}
	}

	if option := s.GetClient(item.DeviceID); option.IsSome() {
		client := option.Unwrap()

		return client.GetTagValue(ctx, in)
	}

	return &output, status.Errorf(codes.Unavailable, "Device not connect")
}

func (s *ControlService) SetTagValue(ctx context.Context, in *pb.TagValue) (*pb.TagValue, error) {
	var err error
	var output pb.TagValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value")
		}
	}

	// tag
	item, err := s.cs.GetTag().ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag.Status != ON")
	}

	if item.Access != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag.Access != ON")
	}

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	// validation device and source
	{
		// device
		{
			device, err := s.cs.GetDevice().ViewByID(ctx, item.DeviceID)
			if err != nil {
				return &output, err
			}

			if device.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Device.Status != ON")
			}
		}

		// source
		{
			source, err := s.cs.GetSource().ViewByID(ctx, item.SourceID)
			if err != nil {
				return &output, err
			}

			if source.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Source.Status != ON")
			}
		}
	}

	if option := s.GetClient(item.DeviceID); option.IsSome() {
		client := option.Unwrap()

		return client.SetTagValue(ctx, in)
	}

	return &output, status.Errorf(codes.Unavailable, "Device not connect")
}

func (s *ControlService) AddClient(deviceID string, client cores.ControlServiceClient) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if chans, ok := s.clients[deviceID]; ok {
		chans.add(client)
	} else {
		chans := &controlChannels{}
		chans.add(client)
		s.clients[deviceID] = chans
	}
}

func (s *ControlService) RemoveClient(deviceID string, client cores.ControlServiceClient) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if chans, ok := s.clients[deviceID]; ok {
		chans.remove(client)
	}
}

func (s *ControlService) GetClient(deviceID string) types.Option[cores.ControlServiceClient] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if chans, ok := s.clients[deviceID]; ok {
		return chans.pick()
	}

	return types.None[cores.ControlServiceClient]()
}

type controlChannels struct {
	cs []cores.ControlServiceClient
}

func (c *controlChannels) add(client cores.ControlServiceClient) {
	c.cs = append(c.cs, client)
}

func (c *controlChannels) remove(client cores.ControlServiceClient) {
	for i := range c.cs {
		if c.cs[i] == client {
			c.cs = append(c.cs[:i], c.cs[i+1:]...)
			break
		}
	}
}

func (c *controlChannels) pick() types.Option[cores.ControlServiceClient] {
	if c == nil {
		return types.None[cores.ControlServiceClient]()
	}

	if len(c.cs) == 0 {
		return types.None[cores.ControlServiceClient]()
	}

	return types.Some(c.cs[len(c.cs)-1])
}
