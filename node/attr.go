package node

import (
	"context"

	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/cores"
	"github.com/snple/kirara/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AttrService struct {
	ns *NodeService

	nodes.UnimplementedAttrServiceServer
}

func newAttrService(ns *NodeService) *AttrService {
	return &AttrService{
		ns: ns,
	}
}

func (s *AttrService) Create(ctx context.Context, in *pb.Attr) (*pb.Attr, error) {
	var output pb.Attr

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	// class validation
	{
		deviceID, err := validateToken(ctx)
		if err != nil {
			return &output, err
		}

		request := &pb.Id{Id: in.GetClassId()}

		reply, err := s.ns.Core().GetClass().View(ctx, request)
		if err != nil {
			return &output, err
		}

		if reply.GetDeviceId() != deviceID {
			return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
		}
	}

	return s.ns.Core().GetAttr().Create(ctx, in)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetAttr().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetAttr().Update(ctx, in)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetAttr().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &cores.AttrNameRequest{DeviceId: deviceID, Name: in.GetName()}

	reply, err := s.ns.Core().GetAttr().Name(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetAttr().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetAttr().Delete(ctx, in)
}

func (s *AttrService) List(ctx context.Context, in *nodes.AttrListRequest) (*nodes.AttrListResponse, error) {
	var err error
	var output nodes.AttrListResponse

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

	request := &cores.AttrListRequest{
		Page:     in.GetPage(),
		DeviceId: deviceID,
		ClassId:  in.GetClassId(),
		Tags:     in.GetTags(),
		Type:     in.GetType(),
	}

	reply, err := s.ns.Core().GetAttr().List(ctx, request)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetAttr().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetAttr().GetValue(ctx, in)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetAttr().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetAttr().SetValue(ctx, in)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetAttr().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetAttr().SetValueForce(ctx, in)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetAttr().GetValueByName(ctx,
		&cores.AttrGetValueByNameRequest{DeviceId: deviceID, Name: in.GetName()})
	if err != nil {
		return &output, err
	}

	output.Id = reply.GetId()
	output.Name = reply.GetName()
	output.Value = reply.GetValue()
	output.Updated = reply.GetUpdated()

	return &output, nil
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ns.Core().GetAttr().SetValueByName(ctx,
		&cores.AttrNameValue{DeviceId: deviceID, Name: in.GetName(), Value: in.GetValue()})
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ns.Core().GetAttr().SetValueByNameForce(ctx,
		&cores.AttrNameValue{DeviceId: deviceID, Name: in.GetName(), Value: in.GetValue()})
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetAttr().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *AttrService) Pull(ctx context.Context, in *nodes.AttrPullRequest) (*nodes.AttrPullResponse, error) {
	var err error
	var output nodes.AttrPullResponse

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

	request := &cores.AttrPullRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		DeviceId: deviceID,
		ClassId:  in.GetClassId(),
		Type:     in.GetType(),
	}

	reply, err := s.ns.Core().GetAttr().Pull(ctx, request)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	in.DeviceId = deviceID

	return s.ns.Core().GetAttr().Sync(ctx, in)
}

func (s *AttrService) Upload(ctx context.Context, in *nodes.AttrValueUploadRequest) (*nodes.AttrValueUploadResponse, error) {
	var output nodes.AttrValueUploadResponse
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

	reply, err := s.ns.Core().GetAttr().Upload(ctx, &cores.AttrValueUploadRequest{
		Id:       in.GetId(),
		Content:  in.GetContent(),
		DeviceId: deviceID,
	})
	if err != nil {
		return &output, err
	}

	output.Id = reply.GetId()
	output.Message = reply.GetMessage()

	return &output, nil
}
