package gos7

import (
	"context"
	"time"

	"github.com/snple/kirara/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Conn) GetTagValue(ctx context.Context, in *pb.Id) (*pb.TagValue, error) {
	// var err error
	var output pb.TagValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag id")
		}
	}

	output.Id = in.GetId()

	value, err := s.readTag(in.GetId())
	if err != nil {
		return &output, status.Errorf(codes.FailedPrecondition, "readTag: %v", err)
	}

	output.Value = value
	output.Updated = time.Now().Unix()

	return &output, nil
}

func (s *Conn) SetTagValue(ctx context.Context, in *pb.TagValue) (*pb.TagValue, error) {
	// var err error
	var output pb.TagValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag id")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid value")
		}
	}

	output.Id = in.GetId()

	err := s.writeTag(in.GetId(), in.GetValue())
	if err != nil {
		return &output, status.Errorf(codes.FailedPrecondition, "writeTag: %v", err)
	}

	value, err := s.readTag(in.GetId())
	if err != nil {
		return &output, status.Errorf(codes.FailedPrecondition, "readTag: %v", err)
	}

	output.Value = value
	output.Updated = time.Now().Unix()

	return &output, nil
}
