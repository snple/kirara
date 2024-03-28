package edge

import (
	"context"
	"time"

	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"github.com/snple/kirara/pb/nodes"
)

func (s *NodeService) syncRemoteToLocal(ctx context.Context) error {
	deviceUpdated, err := s.SyncServiceClient().GetDeviceUpdated(ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	deviceUpdated2, err := s.es.GetSync().getDeviceUpdatedRemoteToLocal(ctx)
	if err != nil {
		return err
	}

	if deviceUpdated.GetUpdated() <= deviceUpdated2.UnixMicro() {
		return nil
	}

	// device
	{
		remote, err := s.DeviceServiceClient().ViewWithDeleted(ctx, &pb.MyEmpty{})
		if err != nil {
			return err
		}

		_, err = s.es.GetDevice().Sync(ctx, remote)
		if err != nil {
			return err
		}
	}

	// slot
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.SlotServiceClient().Pull(ctx, &nodes.SlotPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetSlot() {
				_, err := s.es.GetSlot().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetSlot()) < int(limit) {
				break
			}
		}
	}

	// source
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.SourceServiceClient().Pull(ctx, &nodes.SourcePullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetSource() {
				_, err = s.es.GetSource().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetSource()) < int(limit) {
				break
			}
		}
	}

	// tag
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.TagServiceClient().Pull(ctx, &nodes.TagPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetTag() {
				_, err := s.es.GetTag().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetTag()) < int(limit) {
				break
			}
		}
	}

	// const
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.ConstServiceClient().Pull(ctx, &nodes.ConstPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetConst() {
				_, err := s.es.GetConst().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetConst()) < int(limit) {
				break
			}
		}
	}

	// class
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.ClassServiceClient().Pull(ctx, &nodes.ClassPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetClass() {
				_, err := s.es.GetClass().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetClass()) < int(limit) {
				break
			}
		}
	}

	// attr
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.AttrServiceClient().Pull(ctx, &nodes.AttrPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetAttr() {
				_, err := s.es.GetAttr().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetAttr()) < int(limit) {
				break
			}
		}
	}

	// logic
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.LogicServiceClient().Pull(ctx, &nodes.LogicPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetLogic() {
				_, err = s.es.GetLogic().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetLogic()) < int(limit) {
				break
			}
		}
	}

	// fn
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.FnServiceClient().Pull(ctx, &nodes.FnPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetFn() {
				_, err = s.es.GetFn().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetFn()) < int(limit) {
				break
			}
		}
	}

	return s.es.GetSync().setDeviceUpdatedRemoteToLocal(ctx, time.UnixMicro(deviceUpdated.GetUpdated()))
}

func (s *NodeService) syncLocalToRemote(ctx context.Context) error {
	deviceUpdated, err := s.es.GetSync().getDeviceUpdated(ctx)
	if err != nil {
		return err
	}

	deviceUpdated2, err := s.es.GetSync().getDeviceUpdatedLocalToRemote(ctx)
	if err != nil {
		return err
	}

	if deviceUpdated.UnixMicro() <= deviceUpdated2.UnixMicro() {
		return nil
	}

	// device
	{
		local, err := s.es.GetDevice().ViewWithDeleted(ctx, &pb.MyEmpty{})
		if err != nil {
			return err
		}

		_, err = s.DeviceServiceClient().Sync(ctx, local)
		if err != nil {
			return err
		}
	}

	// slot
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetSlot().Pull(ctx, &edges.SlotPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetSlot() {
				_, err = s.SlotServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetSlot()) < int(limit) {
				break
			}
		}
	}

	// source
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetSource().Pull(ctx, &edges.SourcePullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetSource() {
				_, err = s.SourceServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetSource()) < int(limit) {
				break
			}
		}
	}

	// tag
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetTag().Pull(ctx, &edges.TagPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetTag() {
				_, err = s.TagServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetTag()) < int(limit) {
				break
			}
		}
	}

	// const
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetConst().Pull(ctx, &edges.ConstPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetConst() {
				_, err = s.ConstServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetConst()) < int(limit) {
				break
			}
		}
	}

	// class
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetClass().Pull(ctx, &edges.ClassPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetClass() {
				_, err = s.ClassServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetClass()) < int(limit) {
				break
			}
		}
	}

	// attr
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetAttr().Pull(ctx, &edges.AttrPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetAttr() {
				_, err = s.AttrServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetAttr()) < int(limit) {
				break
			}
		}
	}

	// logic
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetLogic().Pull(ctx, &edges.LogicPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetLogic() {
				_, err = s.LogicServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetLogic()) < int(limit) {
				break
			}
		}
	}

	// fn
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetFn().Pull(ctx, &edges.FnPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetFn() {
				_, err = s.FnServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetFn()) < int(limit) {
				break
			}
		}
	}

	return s.es.GetSync().setDeviceUpdatedLocalToRemote(ctx, deviceUpdated)
}

func (s *NodeService) syncTagValueRemoteToLocal(ctx context.Context) error {
	tagValueUpdated, err := s.SyncServiceClient().GetTagValueUpdated(ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	tagValueUpdated2, err := s.es.GetSync().getTagValueUpdatedRemoteToLocal(ctx)
	if err != nil {
		return err
	}

	if tagValueUpdated.GetUpdated() <= tagValueUpdated2.UnixMicro() {
		return nil
	}

	after := tagValueUpdated2.UnixMicro()
	limit := uint32(100)

PULL:
	for {
		remotes, err := s.TagServiceClient().PullValue(ctx, &nodes.TagPullValueRequest{After: after, Limit: limit})
		if err != nil {
			return err
		}

		for _, remote := range remotes.GetTag() {
			if remote.GetUpdated() > tagValueUpdated.GetUpdated() {
				break PULL
			}

			_, err = s.es.GetTag().SyncValue(ctx,
				&pb.TagValue{Id: remote.GetId(), Value: remote.GetValue(), Updated: remote.GetUpdated()})
			if err != nil {
				s.es.Logger().Sugar().Errorf("SyncValue: %v", err)
				return err
			}

			after = remote.GetUpdated()
		}

		if len(remotes.GetTag()) < int(limit) {
			break
		}
	}

	return s.es.GetSync().setTagValueUpdatedRemoteToLocal(ctx, time.UnixMicro(tagValueUpdated.GetUpdated()))
}

func (s *NodeService) syncTagValueLocalToRemote(ctx context.Context) error {
	tagValueUpdated, err := s.es.GetSync().getTagValueUpdated(ctx)
	if err != nil {
		return err
	}

	tagValueUpdated2, err := s.es.GetSync().getTagValueUpdatedLocalToRemote(ctx)
	if err != nil {
		return err
	}

	if tagValueUpdated.UnixMicro() <= tagValueUpdated2.UnixMicro() {
		return nil
	}

	after := tagValueUpdated2.UnixMicro()
	limit := uint32(100)

PULL:
	for {
		locals, err := s.es.GetTag().PullValue(ctx, &edges.TagPullValueRequest{After: after, Limit: limit})
		if err != nil {
			return err
		}

		for _, local := range locals.GetTag() {
			if local.GetUpdated() > tagValueUpdated.UnixMicro() {
				break PULL
			}

			_, err = s.TagServiceClient().SyncValue(ctx,
				&pb.TagValue{Id: local.GetId(), Value: local.GetValue(), Updated: local.GetUpdated()})
			if err != nil {
				s.es.Logger().Sugar().Errorf("SyncValue: %v", err)
				return err
			}

			after = local.GetUpdated()
		}

		if len(locals.GetTag()) < int(limit) {
			break
		}
	}

	return s.es.GetSync().setTagValueUpdatedLocalToRemote(ctx, tagValueUpdated)
}
