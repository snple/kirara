package web

import (
	"github.com/gin-gonic/gin"
	"github.com/snple/kirara/http/util"
	"github.com/snple/kirara/http/util/shiftime"
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClassService struct {
	ws *WebService
}

func newClassService(ws *WebService) *ClassService {
	return &ClassService{
		ws: ws,
	}
}

func (s *ClassService) register(router gin.IRouter) {
	group := router.Group("/class")

	group.Use(s.ws.GetAuth().MiddlewareFunc())

	group.GET("/", s.list)
	group.GET("/:id", s.get)
	group.POST("/", s.post)
	group.PATCH("/:id", s.patch)
	group.PATCH("/:id/status", s.status)
	group.DELETE("/:id", s.delete)
}

func (s *ClassService) list(ctx *gin.Context) {
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

	request := &edges.ClassListRequest{
		Page: page,
		Tags: params.Tags,
		Type: params.Type,
	}

	reply, err := s.ws.Edge().GetClass().List(ctx, request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	items := reply.GetClass()

	shiftime.Classes(items)

	ctx.JSON(util.Success(gin.H{
		"items": items,
		"total": reply.GetCount(),
	}))
}

func (s *ClassService) get(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.ws.Edge().GetClass().View(ctx, request)
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

	shiftime.Class(reply)

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *ClassService) post(ctx *gin.Context) {
	var params pb.Class

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply, err := s.ws.Edge().GetClass().Create(ctx, &params)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	shiftime.Class(reply)

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *ClassService) patch(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.ws.Edge().GetClass().View(ctx, request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	var params pb.Class
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply.Name = params.Name
	reply.Desc = params.Desc
	reply.Tags = params.Tags
	reply.Type = params.Type
	reply.Config = params.Config
	reply.Status = params.Status
	reply.Save = params.Save

	reply2, err := s.ws.Edge().GetClass().Update(ctx, reply)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{"id": reply2.GetId()}))
}

func (s *ClassService) delete(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	_, err := s.ws.Edge().GetClass().Delete(ctx, request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{"id": ctx.Param("id")}))
}

func (s *ClassService) status(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.ws.Edge().GetClass().View(ctx, request)
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

	reply2, err := s.ws.Edge().GetClass().Update(ctx, reply)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{"id": reply2.GetId()}))
}
