package api

import (
	"github.com/gin-gonic/gin"
	"github.com/snple/kirara/http/util"
	"github.com/snple/kirara/http/util/shiftime"
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/cores"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClassService struct {
	as *ApiService
}

func newClassService(as *ApiService) *ClassService {
	return &ClassService{
		as: as,
	}
}

func (s *ClassService) register(router gin.IRouter) {
	group := router.Group("/class")

	group.GET("/", s.list)

	group.GET("/:id", s.getById)

	group.GET("/name/:name", s.getByName)
	group.POST("/name", s.getByNames)
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

	request := &cores.ClassListRequest{
		Page:     page,
		DeviceId: params.DeviceId,
		Tags:     params.Tags,
		Type:     params.Type,
	}

	reply, err := s.as.Core().GetClass().List(ctx, request)
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

func (s *ClassService) getById(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.as.Core().GetClass().View(ctx, request)
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

func (s *ClassService) getByName(ctx *gin.Context) {
	name := ctx.Param("name")

	var params struct {
		DeviceId string `form:"device_id"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply, err := s.as.Core().GetClass().Name(ctx,
		&cores.ClassNameRequest{DeviceId: params.DeviceId, Name: name})
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

func (s *ClassService) getByNames(ctx *gin.Context) {
	var params struct {
		DeviceId string   `json:"device_id"`
		Name     []string `json:"name"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ret := make([]*pb.Class, 0, len(params.Name))

	for _, name := range params.Name {
		reply, err := s.as.Core().GetClass().Name(ctx,
			&cores.ClassNameRequest{DeviceId: params.DeviceId, Name: name})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.NotFound {
					continue
				}
			}

			ctx.JSON(util.Error(400, err.Error()))
			return
		}

		shiftime.Class(reply)

		ret = append(ret, reply)
	}

	ctx.JSON(util.Success(ret))
}
