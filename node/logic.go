package node

import (
	"context"

	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/cores"
	"github.com/snple/kirara/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LogicService struct {
	ns *NodeService

	nodes.UnimplementedLogicServiceServer
}

func newLogicService(ns *NodeService) *LogicService {
	return &LogicService{
		ns: ns,
	}
}

func (s *LogicService) Create(ctx context.Context, in *pb.Logic) (*pb.Logic, error) {
	var output pb.Logic
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	in.DeviceId = deviceID

	return s.ns.Core().GetLogic().Create(ctx, in)
}

func (s *LogicService) Update(ctx context.Context, in *pb.Logic) (*pb.Logic, error) {
	var output pb.Logic
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetLogic().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetLogic().Update(ctx, in)
}

func (s *LogicService) View(ctx context.Context, in *pb.Id) (*pb.Logic, error) {
	var output pb.Logic
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetLogic().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *LogicService) Name(ctx context.Context, in *pb.Name) (*pb.Logic, error) {
	var output pb.Logic
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &cores.LogicNameRequest{DeviceId: deviceID, Name: in.GetName()}

	reply, err := s.ns.Core().GetLogic().Name(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *LogicService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetLogic().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetLogic().Delete(ctx, in)
}

func (s *LogicService) List(ctx context.Context, in *nodes.LogicListRequest) (*nodes.LogicListResponse, error) {
	var err error
	var output nodes.LogicListResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &cores.LogicListRequest{
		Page:     in.GetPage(),
		DeviceId: deviceID,
		Tags:     in.GetTags(),
		Type:     in.GetType(),
		Exec:     in.GetExec(),
	}

	reply, err := s.ns.Core().GetLogic().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Logic = reply.GetLogic()

	return &output, nil
}

func (s *LogicService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Logic, error) {
	var output pb.Logic
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetLogic().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *LogicService) Pull(ctx context.Context, in *nodes.LogicPullRequest) (*nodes.LogicPullResponse, error) {
	var err error
	var output nodes.LogicPullResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &cores.LogicPullRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		DeviceId: deviceID,
		Exec:     in.GetExec(),
	}

	reply, err := s.ns.Core().GetLogic().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Logic = reply.GetLogic()

	return &output, nil
}

func (s *LogicService) Sync(ctx context.Context, in *pb.Logic) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	in.DeviceId = deviceID

	return s.ns.Core().GetLogic().Sync(ctx, in)
}
