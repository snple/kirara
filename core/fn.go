package core

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/snple/kirara/core/model"
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/cores"
	"github.com/snple/kirara/util"
	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FnService struct {
	cs *CoreService

	cores.UnimplementedFnServiceServer
}

func newFnService(cs *CoreService) *FnService {
	return &FnService{
		cs: cs,
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

		if len(in.GetDeviceId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Fn.DeviceID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Fn.Name")
		}
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "Fn.Name min 2 character")
		}

		err = s.cs.GetDB().NewSelect().Model(&model.Fn{}).Where("device_id = ?", in.GetDeviceId()).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			return &output, status.Error(codes.AlreadyExists, "Fn.Name must be unique")
		}
	}

	item := model.Fn{
		ID:       in.GetId(),
		DeviceID: in.GetDeviceId(),
		Name:     in.GetName(),
		Desc:     in.GetDesc(),
		Tags:     in.GetTags(),
		Type:     in.GetType(),
		Exec:     in.GetExec(),
		Main:     in.GetMain(),
		Config:   in.GetConfig(),
		Debug:    in.GetDebug(),
		Status:   in.GetStatus(),
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	if len(item.ID) == 0 {
		item.ID = util.RandomID()
	}

	_, err = s.cs.GetDB().NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Insert: %v", err)
	}

	if err = s.afterUpdate(ctx, &item); err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *FnService) Update(ctx context.Context, in *pb.Fn) (*pb.Fn, error) {
	var output pb.Fn
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Fn.ID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Fn.Name")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "Fn.Name min 2 character")
		}

		modelItem := model.Fn{}
		err = s.cs.GetDB().NewSelect().Model(&modelItem).Where("device_id = ?", in.GetDeviceId()).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			if modelItem.ID != item.ID {
				return &output, status.Error(codes.AlreadyExists, "Fn.Name must be unique")
			}
		}
	}

	item.Name = in.GetName()
	item.Desc = in.GetDesc()
	item.Tags = in.GetTags()
	item.Type = in.GetType()
	item.Exec = in.GetExec()
	item.Main = in.GetMain()
	item.Config = in.GetConfig()
	item.Status = in.GetStatus()
	item.Debug = in.GetDebug()
	item.Updated = time.Now()

	_, err = s.cs.GetDB().NewUpdate().Model(&item).WherePK().Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Update: %v", err)
	}

	if err = s.afterUpdate(ctx, &item); err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *FnService) View(ctx context.Context, in *pb.Id) (*pb.Fn, error) {
	var output pb.Fn
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Fn.ID")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *FnService) Name(ctx context.Context, in *cores.FnNameRequest) (*pb.Fn, error) {
	var output pb.Fn
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetDeviceId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid DeviceID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Fn.Name")
		}
	}

	item, err := s.viewByDeviceIDAndName(ctx, in.GetDeviceId(), in.GetName())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *FnService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Fn.ID")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	item.Updated = time.Now()
	item.Deleted = time.Now()

	_, err = s.cs.GetDB().NewUpdate().Model(&item).Column("updated", "deleted").WherePK().Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Delete: %v", err)
	}

	if err = s.afterDelete(ctx, &item); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *FnService) List(ctx context.Context, in *cores.FnListRequest) (*cores.FnListResponse, error) {
	var err error
	var output cores.FnListResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	defaultPage := pb.Page{
		Limit:  10,
		Offset: 0,
	}

	if in.GetPage() == nil {
		in.Page = &defaultPage
	}

	output.Page = in.GetPage()

	var items []model.Fn

	query := s.cs.GetDB().NewSelect().Model(&items)

	if len(in.GetDeviceId()) > 0 {
		query.Where("device_id = ?", in.GetDeviceId())
	}

	if len(in.GetExec()) > 0 {
		query.Where("exec = ?", in.GetExec())
	}

	if len(in.GetPage().GetSearch()) > 0 {
		search := fmt.Sprintf("%%%v%%", in.GetPage().GetSearch())

		query.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			q = q.Where(`"name" LIKE ?`, search).
				WhereOr(`"desc" LIKE ?`, search)

			return q
		})
	}

	if len(in.GetTags()) > 0 {
		tagsSplit := strings.Split(in.GetTags(), ",")

		if len(tagsSplit) == 1 {
			search := fmt.Sprintf("%%%v%%", tagsSplit[0])

			query = query.Where(`"tags" LIKE ?`, search)
		} else {
			query = query.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
				for i := 0; i < len(tagsSplit); i++ {
					search := fmt.Sprintf("%%%v%%", tagsSplit[i])

					q = q.WhereOr(`"tags" LIKE ?`, search)
				}

				return q
			})
		}
	}

	if len(in.GetType()) > 0 {
		query = query.Where(`type = ?`, in.GetType())
	}

	if len(in.GetPage().GetOrderBy()) > 0 && (in.GetPage().GetOrderBy() == "id" || in.GetPage().GetOrderBy() == "name" ||
		in.GetPage().GetOrderBy() == "created" || in.GetPage().GetOrderBy() == "updated") {
		query.Order(in.GetPage().GetOrderBy() + " " + in.GetPage().GetSort().String())
	} else {
		query.Order("id ASC")
	}

	count, err := query.Offset(int(in.GetPage().GetOffset())).Limit(int(in.GetPage().GetLimit())).ScanAndCount(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Query: %v", err)
	}

	output.Count = uint32(count)

	for i := 0; i < len(items); i++ {
		item := pb.Fn{}

		s.copyModelToOutput(&item, &items[i])

		output.Fn = append(output.Fn, &item)
	}

	return &output, nil
}

func (s *FnService) Link(ctx context.Context, in *cores.FnLinkRequest) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Fn.ID")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.cs.GetStatus().SetLink(item.ID, in.GetStatus())

	output.Bool = true

	return &output, nil
}

func (s *FnService) Clone(ctx context.Context, in *cores.CloneFnRequest) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Fn.ID")
		}
	}

	err = s.cs.getClone().fn(ctx, s.cs.GetDB(), in.GetId(), in.GetDeviceId())
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *FnService) ViewByID(ctx context.Context, id string) (model.Fn, error) {
	item := model.Fn{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Fn.ID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *FnService) viewByDeviceIDAndName(ctx context.Context, deviceID, name string) (model.Fn, error) {
	item := model.Fn{}

	err := s.cs.GetDB().NewSelect().Model(&item).Where("device_id = ?", deviceID).Where("name = ?", name).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, DeviceID: %v, Fn.Name: %v", err, deviceID, name)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *FnService) copyModelToOutput(output *pb.Fn, item *model.Fn) {
	output.Id = item.ID
	output.DeviceId = item.DeviceID
	output.Name = item.Name
	output.Desc = item.Desc
	output.Tags = item.Tags
	output.Type = item.Type
	output.Exec = item.Exec
	output.Main = item.Main
	output.Config = item.Config
	output.Link = s.cs.GetStatus().GetLink(item.ID)
	output.Status = item.Status
	output.Debug = item.Debug
	output.Created = item.Created.UnixMicro()
	output.Updated = item.Updated.UnixMicro()
	output.Deleted = item.Deleted.UnixMicro()
}

func (s *FnService) afterUpdate(ctx context.Context, item *model.Fn) error {
	var err error

	err = s.cs.GetSync().setDeviceUpdated(ctx, s.cs.GetDB(), item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *FnService) afterDelete(ctx context.Context, item *model.Fn) error {
	var err error

	err = s.cs.GetSync().setDeviceUpdated(ctx, s.cs.GetDB(), item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *FnService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Fn, error) {
	var output pb.Fn
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Fn.ID")
		}
	}

	item, err := s.viewWithDeleted(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *FnService) viewWithDeleted(ctx context.Context, id string) (model.Fn, error) {
	item := model.Fn{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().WhereAllWithDeleted().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Fn.ID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *FnService) Pull(ctx context.Context, in *cores.FnPullRequest) (*cores.FnPullResponse, error) {
	var err error
	var output cores.FnPullResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	var items []model.Fn

	query := s.cs.GetDB().NewSelect().Model(&items)

	if len(in.GetDeviceId()) > 0 {
		query.Where("device_id = ?", in.GetDeviceId())
	}

	if len(in.GetExec()) > 0 {
		query.Where("exec = ?", in.GetExec())
	}

	err = query.Where("updated > ?", time.UnixMicro(in.GetAfter())).WhereAllWithDeleted().Order("updated ASC").Limit(int(in.GetLimit())).Scan(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Query: %v", err)
	}

	for i := 0; i < len(items); i++ {
		item := pb.Fn{}

		s.copyModelToOutput(&item, &items[i])

		output.Fn = append(output.Fn, &item)
	}

	return &output, nil
}

func (s *FnService) Sync(ctx context.Context, in *pb.Fn) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Fn.ID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Fn.Name")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Fn.Updated")
		}
	}

	insert := false
	update := false

	item, err := s.viewWithDeleted(ctx, in.GetId())
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				insert = true
				goto SKIP
			}
		}

		return &output, err
	}

	update = true

SKIP:

	// insert
	if insert {
		// device validation
		{
			_, err = s.cs.GetDevice().viewWithDeleted(ctx, in.GetDeviceId())
			if err != nil {
				return &output, err
			}
		}

		// name validation
		{
			if len(in.GetName()) < 2 {
				return &output, status.Error(codes.InvalidArgument, "Fn.Name min 2 character")
			}

			err = s.cs.GetDB().NewSelect().Model(&model.Fn{}).Where("device_id = ?", in.GetDeviceId()).Where("name = ?", in.GetName()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				return &output, status.Error(codes.AlreadyExists, "Fn.Name must be unique")
			}
		}

		item := model.Fn{
			ID:       in.GetId(),
			DeviceID: in.GetDeviceId(),
			Name:     in.GetName(),
			Desc:     in.GetDesc(),
			Tags:     in.GetTags(),
			Type:     in.GetType(),
			Exec:     in.GetExec(),
			Main:     in.GetMain(),
			Config:   in.GetConfig(),
			Debug:    in.GetDebug(),
			Status:   in.GetStatus(),
			Created:  time.UnixMicro(in.GetCreated()),
			Updated:  time.UnixMicro(in.GetUpdated()),
		}

		_, err = s.cs.GetDB().NewInsert().Model(&item).Exec(ctx)
		if err != nil {
			return &output, status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	// update
	if update {
		if in.GetDeviceId() != item.DeviceID {
			return &output, status.Error(codes.NotFound, "Query: in.GetDeviceId() != item.DeviceID")
		}

		if in.GetUpdated() <= item.Updated.UnixMicro() {
			return &output, nil
		}

		// name validation
		{
			if len(in.GetName()) < 2 {
				return &output, status.Error(codes.InvalidArgument, "Fn.Name min 2 character")
			}

			modelItem := model.Fn{}
			err = s.cs.GetDB().NewSelect().Model(&modelItem).Where("device_id = ?", in.GetDeviceId()).Where("name = ?", in.GetName()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				if modelItem.ID != item.ID {
					return &output, status.Error(codes.AlreadyExists, "Fn.Name must be unique")
				}
			}
		}

		item.Name = in.GetName()
		item.Desc = in.GetDesc()
		item.Tags = in.GetTags()
		item.Type = in.GetType()
		item.Exec = in.GetExec()
		item.Main = in.GetMain()
		item.Config = in.GetConfig()
		item.Status = in.GetStatus()
		item.Debug = in.GetDebug()
		item.Updated = time.UnixMicro(in.GetUpdated())
		item.Deleted = time.UnixMicro(in.GetDeleted())

		_, err = s.cs.GetDB().NewUpdate().Model(&item).WherePK().WhereAllWithDeleted().Exec(ctx)
		if err != nil {
			return &output, status.Errorf(codes.Internal, "Update: %v", err)
		}
	}

	if err = s.afterUpdate(ctx, &item); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}
