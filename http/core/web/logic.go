package web

import (
	"github.com/gin-gonic/gin"
	"github.com/snple/kirara/http/util"
	"github.com/snple/kirara/http/util/shiftime"
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/cores"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LogicService struct {
	ws *WebService
}

func newLogicService(ws *WebService) *LogicService {
	return &LogicService{
		ws: ws,
	}
}

func (s *LogicService) register(router gin.IRouter) {
	group := router.Group("/logic")

	group.Use(s.ws.GetAuth().MiddlewareFunc())

	group.GET("/", s.list)
	group.GET("/:id", s.get)
	group.POST("/", s.post)
	group.PATCH("/:id", s.patch)
	group.PATCH("/:id/status", s.status)
	group.DELETE("/:id", s.delete)
}

func (s *LogicService) list(ctx *gin.Context) {
	var params struct {
		util.Page `form:",inline"`
		DeviceId  string `form:"device_id"`
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

	request := &cores.LogicListRequest{
		Page:     page,
		DeviceId: params.DeviceId,
		Tags:     params.Tags,
		Type:     params.Type,
	}

	reply, err := s.ws.Core().GetLogic().List(ctx, request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	items := reply.GetLogic()

	shiftime.Logics(items)

	ctx.JSON(util.Success(gin.H{
		"items": items,
		"total": reply.GetCount(),
	}))
}

func (s *LogicService) get(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.ws.Core().GetLogic().View(ctx, request)
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

	shiftime.Logic(reply)

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *LogicService) post(ctx *gin.Context) {
	var params pb.Logic

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply, err := s.ws.Core().GetLogic().Create(ctx, &params)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	shiftime.Logic(reply)

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *LogicService) patch(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.ws.Core().GetLogic().View(ctx, request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	var params pb.Logic

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply.Name = params.Name
	reply.Desc = params.Desc
	reply.Tags = params.Tags
	reply.Type = params.Type
	reply.Exec = params.Exec
	reply.Main = params.Main
	reply.Config = params.Config
	reply.Status = params.Status

	reply2, err := s.ws.Core().GetLogic().Update(ctx, reply)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{"id": reply2.GetId()}))
}

func (s *LogicService) delete(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	_, err := s.ws.Core().GetLogic().Delete(ctx, request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{"id": ctx.Param("id")}))
}

func (s *LogicService) status(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.ws.Core().GetLogic().View(ctx, request)
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

	reply2, err := s.ws.Core().GetLogic().Update(ctx, reply)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{"id": reply2.GetId()}))
}
