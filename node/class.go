package node

import (
	"context"

	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/cores"
	"github.com/snple/kirara/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClassService struct {
	ns *NodeService

	nodes.UnimplementedClassServiceServer
}

func newClassService(ns *NodeService) *ClassService {
	return &ClassService{
		ns: ns,
	}
}

func (s *ClassService) Create(ctx context.Context, in *pb.Class) (*pb.Class, error) {
	var output pb.Class
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

	return s.ns.Core().GetClass().Create(ctx, in)
}

func (s *ClassService) Update(ctx context.Context, in *pb.Class) (*pb.Class, error) {
	var output pb.Class
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

	reply, err := s.ns.Core().GetClass().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetClass().Update(ctx, in)
}

func (s *ClassService) View(ctx context.Context, in *pb.Id) (*pb.Class, error) {
	var output pb.Class
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

	reply, err := s.ns.Core().GetClass().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *ClassService) Name(ctx context.Context, in *pb.Name) (*pb.Class, error) {
	var output pb.Class
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

	request := &cores.ClassNameRequest{DeviceId: deviceID, Name: in.GetName()}

	reply, err := s.ns.Core().GetClass().Name(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *ClassService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	reply, err := s.ns.Core().GetClass().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetClass().Delete(ctx, in)
}

func (s *ClassService) List(ctx context.Context, in *nodes.ClassListRequest) (*nodes.ClassListResponse, error) {
	var err error
	var output nodes.ClassListResponse

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

	request := &cores.ClassListRequest{
		Page:     in.GetPage(),
		DeviceId: deviceID,
		Tags:     in.GetTags(),
		Type:     in.GetType(),
	}

	reply, err := s.ns.Core().GetClass().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Class = reply.GetClass()

	return &output, nil
}

func (s *ClassService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Class, error) {
	var output pb.Class
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

	reply, err := s.ns.Core().GetClass().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *ClassService) Pull(ctx context.Context, in *nodes.ClassPullRequest) (*nodes.ClassPullResponse, error) {
	var err error
	var output nodes.ClassPullResponse

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

	request := &cores.ClassPullRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		DeviceId: deviceID,
		Type:     in.GetType(),
	}

	reply, err := s.ns.Core().GetClass().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Class = reply.GetClass()

	return &output, nil
}

func (s *ClassService) Sync(ctx context.Context, in *pb.Class) (*pb.MyBool, error) {
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

	return s.ns.Core().GetClass().Sync(ctx, in)
}
