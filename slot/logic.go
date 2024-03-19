package slot

import (
	"context"

	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"github.com/snple/kirara/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LogicService struct {
	ss *SlotService

	slots.UnimplementedLogicServiceServer
}

func newLogicService(ss *SlotService) *LogicService {
	return &LogicService{
		ss: ss,
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetLogic().Create(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetLogic().Update(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetLogic().View(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetLogic().Name(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetLogic().Delete(ctx, in)
}

func (s *LogicService) List(ctx context.Context, in *slots.LogicListRequest) (*slots.LogicListResponse, error) {
	var err error
	var output slots.LogicListResponse

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

	request := &edges.LogicListRequest{
		Page: in.GetPage(),
		Tags: in.GetTags(),
		Type: in.GetType(),
		Exec: in.GetExec(),
	}

	reply, err := s.ss.Edge().GetLogic().List(ctx, request)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ss.Edge().GetLogic().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *LogicService) Pull(ctx context.Context, in *slots.LogicPullRequest) (*slots.LogicPullResponse, error) {
	var err error
	var output slots.LogicPullResponse

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

	request := &edges.LogicPullRequest{
		After: in.GetAfter(),
		Limit: in.GetLimit(),
		Exec:  in.GetExec(),
	}

	reply, err := s.ss.Edge().GetLogic().Pull(ctx, request)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetLogic().Sync(ctx, in)
}
