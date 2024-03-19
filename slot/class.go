package slot

import (
	"context"

	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"github.com/snple/kirara/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClassService struct {
	ss *SlotService

	slots.UnimplementedClassServiceServer
}

func newClassService(ss *SlotService) *ClassService {
	return &ClassService{
		ss: ss,
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetClass().Create(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetClass().Update(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetClass().View(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetClass().Name(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetClass().Delete(ctx, in)
}

func (s *ClassService) List(ctx context.Context, in *slots.ClassListRequest) (*slots.ClassListResponse, error) {
	var err error
	var output slots.ClassListResponse

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

	request := &edges.ClassListRequest{
		Page: in.GetPage(),
		Tags: in.GetTags(),
		Type: in.GetType(),
	}

	reply, err := s.ss.Edge().GetClass().List(ctx, request)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ss.Edge().GetClass().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *ClassService) Pull(ctx context.Context, in *slots.ClassPullRequest) (*slots.ClassPullResponse, error) {
	var err error
	var output slots.ClassPullResponse

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

	request := &edges.ClassPullRequest{
		After: in.GetAfter(),
		Limit: in.GetLimit(),
		Type:  in.GetType(),
	}

	reply, err := s.ss.Edge().GetClass().Pull(ctx, request)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetClass().Sync(ctx, in)
}
