package source

import (
	"context"
	"time"

	"github.com/snple/kirara/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Conn) GetTagValue(ctx context.Context, in *pb.Id) (*pb.TagValue, error) {
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

	if tag, ok := c.getTag(in.GetId()); ok {
		value, err := c.adapter.ReadTag(tag)
		if err != nil {
			return &output, status.Errorf(codes.FailedPrecondition, "Read Tag: %v", err)
		}

		output.Value = value
		output.Updated = time.Now().Unix()

		return &output, nil

	}

	return &output, status.Errorf(codes.FailedPrecondition, "tag: %v is not found", in.GetId())
}

func (c *Conn) SetTagValue(ctx context.Context, in *pb.TagValue) (*pb.TagValue, error) {
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

	if tag, ok := c.getTag(in.GetId()); ok {
		err := c.adapter.WriteTag(tag, in.GetValue())
		if err != nil {
			return &output, status.Errorf(codes.FailedPrecondition, "Write Tag: %v", err)
		}

		value, err := c.adapter.ReadTag(tag)
		if err != nil {
			return &output, status.Errorf(codes.FailedPrecondition, "Read Tag: %v", err)
		}

		output.Value = value
		output.Updated = time.Now().Unix()

		return &output, nil

	}

	return &output, status.Errorf(codes.FailedPrecondition, "tag: %v is not found", in.GetId())
}
