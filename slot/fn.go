package slot

import (
	"context"

	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"github.com/snple/kirara/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FnService struct {
	ss *SlotService

	slots.UnimplementedFnServiceServer
}

func newFnService(ss *SlotService) *FnService {
	return &FnService{
		ss: ss,
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetFn().Create(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetFn().Update(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetFn().View(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetFn().Name(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetFn().Delete(ctx, in)
}

func (s *FnService) List(ctx context.Context, in *slots.FnListRequest) (*slots.FnListResponse, error) {
	var err error
	var output slots.FnListResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &edges.FnListRequest{
		Page: in.GetPage(),
		Tags: in.GetTags(),
		Type: in.GetType(),
		Exec: in.GetExec(),
	}

	reply, err := s.ss.Edge().GetFn().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Fn = reply.GetFn()

	return &output, nil
}

func (s *FnService) Link(ctx context.Context, in *slots.FnLinkRequest) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request2 := &edges.FnLinkRequest{Id: in.GetId(), Status: in.GetStatus()}

	reply, err := s.ss.Edge().GetFn().Link(ctx, request2)
	if err != nil {
		return &output, err
	}

	return reply, nil
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ss.Edge().GetFn().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *FnService) Pull(ctx context.Context, in *slots.FnPullRequest) (*slots.FnPullResponse, error) {
	var err error
	var output slots.FnPullResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &edges.FnPullRequest{
		After: in.GetAfter(),
		Limit: in.GetLimit(),
		Exec:  in.GetExec(),
	}

	reply, err := s.ss.Edge().GetFn().Pull(ctx, request)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetFn().Sync(ctx, in)
}
