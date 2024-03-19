package slot

import (
	"context"

	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"github.com/snple/kirara/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AttrService struct {
	ss *SlotService

	slots.UnimplementedAttrServiceServer
}

func newAttrService(ss *SlotService) *AttrService {
	return &AttrService{
		ss: ss,
	}
}

func (s *AttrService) Create(ctx context.Context, in *pb.Attr) (*pb.Attr, error) {
	var output pb.Attr
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

	return s.ss.Edge().GetAttr().Create(ctx, in)
}

func (s *AttrService) Update(ctx context.Context, in *pb.Attr) (*pb.Attr, error) {
	var output pb.Attr
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

	return s.ss.Edge().GetAttr().Update(ctx, in)
}

func (s *AttrService) View(ctx context.Context, in *pb.Id) (*pb.Attr, error) {
	var output pb.Attr
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

	return s.ss.Edge().GetAttr().View(ctx, in)
}

func (s *AttrService) Name(ctx context.Context, in *pb.Name) (*pb.Attr, error) {
	var output pb.Attr
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

	return s.ss.Edge().GetAttr().Name(ctx, in)
}

func (s *AttrService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetAttr().Delete(ctx, in)
}

func (s *AttrService) List(ctx context.Context, in *slots.AttrListRequest) (*slots.AttrListResponse, error) {
	var err error
	var output slots.AttrListResponse

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

	request := &edges.AttrListRequest{
		Page:    in.GetPage(),
		ClassId: in.GetClassId(),
		Tags:    in.GetTags(),
		Type:    in.GetType(),
	}

	reply, err := s.ss.Edge().GetAttr().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Attr = reply.GetAttr()

	return &output, nil
}

func (s *AttrService) GetValue(ctx context.Context, in *pb.Id) (*pb.AttrValue, error) {
	var err error
	var output pb.AttrValue

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

	return s.ss.Edge().GetAttr().GetValue(ctx, in)
}

func (s *AttrService) SetValue(ctx context.Context, in *pb.AttrValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetAttr().SetValue(ctx, in)
}

func (s *AttrService) SetValueForce(ctx context.Context, in *pb.AttrValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetAttr().SetValueForce(ctx, in)
}

func (s *AttrService) GetValueByName(ctx context.Context, in *pb.Name) (*pb.AttrNameValue, error) {
	var err error
	var output pb.AttrNameValue

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

	return s.ss.Edge().GetAttr().GetValueByName(ctx, in)
}

func (s *AttrService) SetValueByName(ctx context.Context, in *pb.AttrNameValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetAttr().SetValueByName(ctx, in)
}

func (s *AttrService) SetValueByNameForce(ctx context.Context, in *pb.AttrNameValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetAttr().SetValueByNameForce(ctx, in)
}

func (s *AttrService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Attr, error) {
	var output pb.Attr
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

	reply, err := s.ss.Edge().GetAttr().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *AttrService) Pull(ctx context.Context, in *slots.AttrPullRequest) (*slots.AttrPullResponse, error) {
	var err error
	var output slots.AttrPullResponse

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

	request := &edges.AttrPullRequest{
		After:   in.GetAfter(),
		Limit:   in.GetLimit(),
		ClassId: in.GetClassId(),
		Type:    in.GetType(),
	}

	reply, err := s.ss.Edge().GetAttr().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Attr = reply.GetAttr()

	return &output, nil
}

func (s *AttrService) Sync(ctx context.Context, in *pb.Attr) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetAttr().Sync(ctx, in)
}
