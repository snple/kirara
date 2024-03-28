package node

import (
	"context"

	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/cores"
	"github.com/snple/kirara/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FnService struct {
	ns *NodeService

	nodes.UnimplementedFnServiceServer
}

func newFnService(ns *NodeService) *FnService {
	return &FnService{
		ns: ns,
	}
}

func (s *FnService) Create(ctx context.Context, in *pb.Fn) (*pb.Fn, error) {
	var output pb.Fn
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

	return s.ns.Core().GetFn().Create(ctx, in)
}

func (s *FnService) Update(ctx context.Context, in *pb.Fn) (*pb.Fn, error) {
	var output pb.Fn
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

	reply, err := s.ns.Core().GetFn().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetFn().Update(ctx, in)
}

func (s *FnService) View(ctx context.Context, in *pb.Id) (*pb.Fn, error) {
	var output pb.Fn
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

	reply, err := s.ns.Core().GetFn().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *FnService) Name(ctx context.Context, in *pb.Name) (*pb.Fn, error) {
	var output pb.Fn
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

	request := &cores.FnNameRequest{DeviceId: deviceID, Name: in.GetName()}

	reply, err := s.ns.Core().GetFn().Name(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *FnService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	reply, err := s.ns.Core().GetFn().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetFn().Delete(ctx, in)
}

func (s *FnService) List(ctx context.Context, in *nodes.FnListRequest) (*nodes.FnListResponse, error) {
	var err error
	var output nodes.FnListResponse

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

	request := &cores.FnListRequest{
		Page:     in.GetPage(),
		DeviceId: deviceID,
		Tags:     in.GetTags(),
		Type:     in.GetType(),
		Exec:     in.GetExec(),
	}

	reply, err := s.ns.Core().GetFn().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Fn = reply.GetFn()

	return &output, nil
}

func (s *FnService) Link(ctx context.Context, in *nodes.FnLinkRequest) (*pb.MyBool, error) {
	var output pb.MyBool
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

	reply, err := s.ns.Core().GetFn().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	request2 := &cores.FnLinkRequest{Id: in.GetId(), Status: in.GetStatus()}

	reply2, err := s.ns.Core().GetFn().Link(ctx, request2)
	if err != nil {
		return &output, err
	}

	return reply2, nil
}

func (s *FnService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Fn, error) {
	var output pb.Fn
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

	reply, err := s.ns.Core().GetFn().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *FnService) Pull(ctx context.Context, in *nodes.FnPullRequest) (*nodes.FnPullResponse, error) {
	var err error
	var output nodes.FnPullResponse

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

	request := &cores.FnPullRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		DeviceId: deviceID,
		Exec:     in.GetExec(),
	}

	reply, err := s.ns.Core().GetFn().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Fn = reply.GetFn()

	return &output, nil
}

func (s *FnService) Sync(ctx context.Context, in *pb.Fn) (*pb.MyBool, error) {
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

	return s.ns.Core().GetFn().Sync(ctx, in)
}
