package api

import (
	"github.com/gin-gonic/gin"
	"github.com/snple/kirara/http/util"
	"github.com/snple/kirara/http/util/shiftime"
	"github.com/snple/kirara/pb"
	"github.com/snple/kirara/pb/edges"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AttrService struct {
	as *ApiService
}

func newAttrService(as *ApiService) *AttrService {
	return &AttrService{
		as: as,
	}
}

func (s *AttrService) register(router gin.IRouter) {
	group := router.Group("/attr")

	group.GET("/", s.list)

	group.GET("/:id", s.getById)
	group.GET("/:id/value", s.getValueById)
	group.PATCH("/:id/value", s.setValueById)
	group.PATCH("/:id/value_force", s.setValueByIdForce)

	group.GET("/name/:name", s.getByName)
	group.POST("/name", s.getByNames)

	group.POST("/get_value", s.getValueByNames)
	group.PATCH("/set_value", s.setValueByNames)
	group.PATCH("/set_value_force", s.setValueByNamesForce)
}

func (s *AttrService) list(ctx *gin.Context) {
	var params struct {
		util.Page `form:",inline"`
		Name      string `form:"name"`
		Tags      string `form:"tags"`
		Type      string `form:"type"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	class, err := s.as.Edge().GetClass().Name(ctx,
		&pb.Name{Name: params.Name})
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

	request := &edges.AttrListRequest{
		Page:    page,
		ClassId: class.Id,
		Tags:    params.Tags,
		Type:    params.Type,
	}

	reply, err := s.as.Edge().GetAttr().List(ctx, request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	items := reply.GetAttr()

	shiftime.Attrs(items)

	ctx.JSON(util.Success(gin.H{
		"class": class,
		"items": items,
		"total": reply.GetCount(),
	}))
}

func (s *AttrService) getById(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.as.Edge().GetAttr().View(ctx, request)
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

func (s *AttrService) getValueById(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.as.Edge().GetAttr().GetValue(ctx, request)
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

	shiftime.AttrValue(reply)

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *AttrService) setValueById(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	var params struct {
		Value string `json:"value"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply, err := s.as.Edge().GetAttr().SetValue(ctx,
		&pb.AttrValue{Id: request.Id, Value: params.Value})
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

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *AttrService) setValueByIdForce(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	var params struct {
		Value string `json:"value"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply, err := s.as.Edge().GetAttr().SetValueForce(ctx,
		&pb.AttrValue{Id: request.Id, Value: params.Value})
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

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *AttrService) getByName(ctx *gin.Context) {
	name := ctx.Param("name")

	reply, err := s.as.Edge().GetAttr().Name(ctx,
		&pb.Name{Name: name})
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

func (s *AttrService) getByNames(ctx *gin.Context) {
	var params struct {
		Name []string `json:"name"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ret := make([]*pb.Attr, 0, len(params.Name))

	for _, name := range params.Name {
		reply, err := s.as.Edge().GetAttr().Name(ctx,
			&pb.Name{Name: name})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.NotFound {
					continue
				}
			}

			ctx.JSON(util.Error(400, err.Error()))
			return
		}

		shiftime.Attr(reply)

		ret = append(ret, reply)
	}

	ctx.JSON(util.Success(ret))
}

func (s *AttrService) getValueByNames(ctx *gin.Context) {
	var params struct {
		Name []string `json:"name"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ret := make([]*pb.AttrNameValue, 0, len(params.Name))

	for _, name := range params.Name {
		reply, err := s.as.Edge().GetAttr().GetValueByName(ctx,
			&pb.Name{Name: name})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.NotFound {
					continue
				}
			}

			ctx.JSON(util.Error(400, err.Error()))
			return
		}

		shiftime.AttrNameValue(reply)

		ret = append(ret, reply)
	}

	ctx.JSON(util.Success(ret))
}

func (s *AttrService) setValueByNames(ctx *gin.Context) {
	var params struct {
		NameValue map[string]string `json:"name_value"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	errors := make(map[string]string)

	for name, value := range params.NameValue {
		_, err := s.as.Edge().GetAttr().SetValueByName(ctx,
			&pb.AttrNameValue{Name: name, Value: value})
		if err != nil {
			errors[name] = err.Error()
		}
	}

	if len(errors) > 0 {
		ctx.JSON(util.Success(gin.H{
			"ok":     false,
			"errors": errors,
		}))

		return
	}

	ctx.JSON(util.Success(gin.H{
		"ok": true,
	}))
}

func (s *AttrService) setValueByNamesForce(ctx *gin.Context) {
	var params struct {
		NameValue map[string]string `json:"name_value"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	errors := make(map[string]string)

	for name, value := range params.NameValue {
		_, err := s.as.Edge().GetAttr().SetValueByNameForce(ctx,
			&pb.AttrNameValue{Name: name, Value: value})
		if err != nil {
			errors[name] = err.Error()
		}
	}

	if len(errors) > 0 {
		ctx.JSON(util.Success(gin.H{
			"ok":     false,
			"errors": errors,
		}))

		return
	}

	ctx.JSON(util.Success(gin.H{
		"ok": true,
	}))
}
