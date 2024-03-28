package node

import (
	"github.com/snple/kirara/pb/cores"
	"github.com/snple/kirara/util/metadata"
	"github.com/snple/rgrpc"
)

type RgrpcService struct {
	ns *NodeService

	rgrpc.UnimplementedRgrpcServiceServer
}

var _ rgrpc.RgrpcServiceServer = (*RgrpcService)(nil)

func newRgrpcService(ns *NodeService) *RgrpcService {
	return &RgrpcService{ns: ns}
}

func (s *RgrpcService) OpenRgrpc(stream rgrpc.RgrpcService_OpenRgrpcServer) error {
	channel := rgrpc.NewClient(stream)
	defer channel.Close()

	deviceID, err := validateToken(stream.Context())
	if err != nil {
		return err
	}

	s.ns.Logger().Sugar().Infof("rgrpc connect success, id: %v, ip: %v",
		deviceID, metadata.GetPeerAddr(stream.Context()))

	client := cores.NewControlServiceClient(channel)

	s.ns.Core().GetControl().AddClient(deviceID, client)
	defer s.ns.Core().GetControl().RemoveClient(deviceID, client)

	defer s.ns.Logger().Sugar().Infof("rgrpc break connect, id: %v, ip: %v",
		deviceID, metadata.GetPeerAddr(stream.Context()))

	<-channel.Done()
	return channel.Err()
}
