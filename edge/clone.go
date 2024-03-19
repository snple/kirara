package edge

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/snple/kirara/model"
	"github.com/snple/kirara/util"
	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type cloneService struct {
	es *EdgeService
}

func newCloneService(es *EdgeService) *cloneService {
	return &cloneService{
		es: es,
	}
}

func (s *cloneService) slot(ctx context.Context, db bun.IDB, slotID string) error {
	var err error

	item := model.Slot{
		ID: slotID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setSlotUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) option(ctx context.Context, db bun.IDB, optionID string) error {
	var err error

	item := model.Option{
		ID: optionID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setOptionUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) source(ctx context.Context, db bun.IDB, sourceID string) error {
	var err error

	item := model.Source{
		ID: sourceID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	// clone tags
	{
		var tags []model.Tag

		err = db.NewSelect().Model(&tags).Where("source_id = ?", sourceID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, tag := range tags {
			tag.ID = util.RandomID()
			tag.SourceID = item.ID

			tag.Created = time.Now()
			tag.Updated = time.Now()

			_, err = db.NewInsert().Model(&tag).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setSourceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) tag(ctx context.Context, db bun.IDB, tagID, sourceID string) error {
	var err error

	item := model.Tag{
		ID: tagID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	// source validation
	if len(sourceID) > 0 {
		source := model.Source{
			ID: sourceID,
		}

		err = db.NewSelect().Model(&source).WherePK().Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return status.Error(codes.InvalidArgument, "Please supply valid Source.ID")
			}

			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		item.SourceID = source.ID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setTagUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) const_(ctx context.Context, db bun.IDB, constID string) error {
	var err error

	item := model.Const{
		ID: constID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setConstUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) class(ctx context.Context, db bun.IDB, classID string) error {
	var err error

	item := model.Class{
		ID: classID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	// clone attrs
	{
		var attrs []model.Attr

		err = db.NewSelect().Model(&attrs).Where("class_id = ?", classID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, attr := range attrs {
			attr.ID = util.RandomID()
			attr.ClassID = item.ID

			attr.Created = time.Now()
			attr.Updated = time.Now()

			_, err = db.NewInsert().Model(&attr).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setClassUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) attr(ctx context.Context, db bun.IDB, attrID, classID string) error {
	var err error

	item := model.Attr{
		ID: attrID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	// class validation
	if len(classID) > 0 {
		class := model.Class{
			ID: classID,
		}

		err = db.NewSelect().Model(&class).WherePK().Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return status.Error(codes.InvalidArgument, "Please supply valid ClassID")
			}

			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		item.ClassID = class.ID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setAttrUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) logic(ctx context.Context, db bun.IDB, logicID string) error {
	var err error

	item := model.Logic{
		ID: logicID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setLogicUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) fn(ctx context.Context, db bun.IDB, fnID string) error {
	var err error

	item := model.Fn{
		ID: fnID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setFnUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func randNameSuffix() string {
	buf := new(bytes.Buffer)

	random := rand.Uint32()
	binary.Write(buf, binary.BigEndian, random)

	return hex.EncodeToString(buf.Bytes())
}
