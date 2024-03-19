package slot

import (
	"context"

	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"github.com/snple/kirara/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DataService struct {
	ss *SlotService

	slots.UnimplementedDataServiceServer
}

func newDataService(ss *SlotService) *DataService {
	return &DataService{
		ss: ss,
	}
}

func (s *DataService) Compile(ctx context.Context, in *slots.DataQueryRequest) (*pb.Message, error) {
	var err error
	var output pb.Message

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

	request := &edges.DataQueryRequest{
		Flux: in.GetFlux(),
		Vars: in.GetVars(),
	}

	return s.ss.es.GetData().Compile(ctx, request)
}

func (s *DataService) Query(in *slots.DataQueryRequest, stream slots.DataService_QueryServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(stream.Context())
	if err != nil {
		return err
	}

	request := &edges.DataQueryRequest{
		Flux: in.GetFlux(),
		Vars: in.GetVars(),
	}

	return s.ss.es.GetData().Query(request, stream)
}

func (s *DataService) QueryById(in *slots.DataQueryByIdRequest, stream slots.DataService_QueryByIdServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(stream.Context())
	if err != nil {
		return err
	}

	request := &edges.DataQueryByIdRequest{
		Id:   in.GetId(),
		Vars: in.GetVars(),
	}

	return s.ss.es.GetData().QueryById(request, stream)
}
