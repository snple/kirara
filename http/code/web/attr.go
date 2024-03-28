package web

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/snple/kirara/http/util"
	"github.com/snple/kirara/http/util/shiftime"
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/cores"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AttrService struct {
	ws *WebService
}

func newAttrService(ws *WebService) *AttrService {
	return &AttrService{
		ws: ws,
	}
}

func (s *AttrService) register(router gin.IRouter) {
	group := router.Group("/attr")

	group.Use(s.ws.GetAuth().MiddlewareFunc())

	group.GET("/", s.list)
	group.GET("/:id", s.get)
	group.POST("/", s.post)
	group.PATCH("/:id", s.patch)
	group.PATCH("/:id/status", s.status)
	group.DELETE("/:id", s.delete)
}

func (s *AttrService) list(ctx *gin.Context) {
	var params struct {
		util.Page `form:",inline"`
		DeviceId  string `form:"device_id"`
		ClassId   string `form:"class_id"`
		Tags      string `form:"tags"`
		Type      string `form:"type"`
	}

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	page := &pb.Page{
		Limit:   params.Limit,
		Offset:  params.Offset,
		Search:  params.Search,
		OrderBy: params.OrderBy,
		Sort:    pb.Page_ASC,
	}

	if params.Sort > 0 {
		page.Sort = pb.Page_DESC
	}

	request := &cores.AttrListRequest{
		Page:     page,
		DeviceId: params.DeviceId,
		ClassId:  params.ClassId,
		Tags:     params.Tags,
		Type:     params.Type,
	}

	reply, err := s.ws.Core().GetAttr().List(context.Background(), request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	items := reply.GetAttr()

	shiftime.Attrs(items)

	ctx.JSON(util.Success(gin.H{
		"items": items,
		"total": reply.GetCount(),
	}))
}

func (s *AttrService) get(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	var params struct {
		Group bool `form:"group"`
	}

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply, err := s.ws.Core().GetAttr().View(context.Background(), request)
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				ctx.JSON(util.Error(404, err.Error()))
				return
			}
		}

		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	shiftime.Attr(reply)

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *AttrService) post(ctx *gin.Context) {
	var params pb.Attr

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply, err := s.ws.Core().GetAttr().Create(context.Background(), &params)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	shiftime.Attr(reply)

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *AttrService) patch(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.ws.Core().GetAttr().View(context.Background(), request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	var params pb.Attr

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply.Name = params.Name
	reply.Desc = params.Desc
	reply.Tags = params.Tags
	reply.Type = params.Type
	reply.DataType = params.DataType
	reply.Value = params.Value
	reply.HValue = params.HValue
	reply.LValue = params.LValue
	reply.TagId = params.TagId
	reply.Config = params.Config
	reply.Status = params.Status
	reply.Access = params.Access
	reply.Save = params.Save

	reply2, err := s.ws.Core().GetAttr().Update(context.Background(), reply)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{"id": reply2.GetId()}))
}

func (s *AttrService) delete(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	_, err := s.ws.Core().GetAttr().Delete(context.Background(), request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{"id": ctx.Param("id")}))
}

func (s *AttrService) status(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.ws.Core().GetAttr().View(ctx, request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	var params struct {
		Status int32 `json:"status"`
	}

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply.Status = params.Status

	reply2, err := s.ws.Core().GetAttr().Update(ctx, reply)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{"id": reply2.GetId()}))
}
