package api

import (
	"bytes"
	"context"
	"fmt"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snple/kirara/http/util"
	modutil "github.com/snple/kirara/util"
	"github.com/snple/kirara/util/flux"
)

type DataService struct {
	as *ApiService
}

func newDataService(as *ApiService) *DataService {
	return &DataService{
		as: as,
	}
}

func (s *DataService) register(router gin.IRouter) {
	group := router.Group("/data")

	group.POST("/compile", s.compile)
	group.POST("/query", s.query)
	group.GET("/query/:id", s.queryById)
}

func (s *DataService) compile(ctx *gin.Context) {
	var params struct {
		Flux string            `json:"flux"`
		Vars map[string]string `json:"vars"`
	}

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	if len(params.Flux) == 0 {
		ctx.JSON(util.Error(400, "Please supply valid flux"))
		return
	}

	influx := s.as.Edge().GetInfluxDB()
	if influx.IsNone() {
		ctx.JSON(util.Error(400, "Influxdb is not enable"))
		return
	}
	influxdb := influx.Unwrap()

	vars := map[string]string{
		"start":  "-1h",
		"stop":   "now()",
		"window": "10m",
		"fn":     "mean",
	}

	if len(params.Vars) > 0 {
		for key, val := range params.Vars {
			vars[key] = tryConvertTimeZone(val)
		}
	}

	vars["bucket"] = influxdb.Bucket()

	tmpl, err := template.New("").Parse(params.Flux)
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("Template parse: %v", err)))
		return
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, vars)
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("Template exec: %v", err)))
		return
	}

	ctx.JSON(util.Success(gin.H{
		"flux": buffer.String(),
	}))
}

func (s *DataService) query(ctx *gin.Context) {
	var params struct {
		Flux string            `json:"flux"`
		Vars map[string]string `json:"vars"`
	}

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	if len(params.Flux) == 0 {
		ctx.JSON(util.Error(400, "Please supply valid flux"))
		return
	}

	influx := s.as.Edge().GetInfluxDB()
	if influx.IsNone() {
		ctx.JSON(util.Error(400, "Influxdb is not enable"))
		return
	}
	influxdb := influx.Unwrap()

	vars := map[string]string{
		"start":  "-1h",
		"stop":   "now()",
		"window": "10m",
		"fn":     "mean",
	}

	if len(params.Vars) > 0 {
		for key, val := range params.Vars {
			vars[key] = tryConvertTimeZone(val)
		}
	}

	vars["bucket"] = influxdb.Bucket()

	tmpl, err := template.New("").Parse(params.Flux)
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("Template parse: %v", err)))
		return
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, vars)
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("Template exec: %v", err)))
		return
	}

	result, err := influxdb.QueryAPI().Query(context.Background(), buffer.String())
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("InfluxDB Query: %v", err)))
		return
	}

	rows := make([]map[string]interface{}, 0)

	for result.Next() {
		values := result.Record().Values()

		for key, val := range values {
			if t, ok := val.(time.Time); ok {
				values[key] = modutil.TimeFormat(t)
			}
		}

		rows = append(rows, values)
	}

	ctx.JSON(util.Success(gin.H{
		"flux": buffer.String(),
		"rows": rows,
	}))
}

func (s *DataService) queryById(ctx *gin.Context) {
	id := ctx.Param("id")

	if len(id) == 0 {
		ctx.JSON(util.Error(400, "Please supply valid id"))
		return
	}

	influx := s.as.Edge().GetInfluxDB()
	if influx.IsNone() {
		ctx.JSON(util.Error(400, "Influxdb is not enable"))
		return
	}
	influxdb := influx.Unwrap()

	vars2 := make(map[string]string, 0)

	if ctx.Request != nil {
		for key, value := range ctx.Request.URL.Query() {
			if len(value) > 0 {
				vars2[key] = value[0]
			}
		}
	}

	vars := map[string]string{
		"start":  "-1h",
		"stop":   "now()",
		"window": "10m",
		"fn":     "mean",
		"fill":   "none",
	}

	if len(vars2) > 0 {
		for key, val := range vars2 {
			vars[key] = tryConvertTimeZone(val)
		}
	}

	vars["id"] = id
	vars["bucket"] = influxdb.Bucket()

	script := flux.TemplateBasic

	if fill, ok := vars["fill"]; ok {
		switch fill {
		case "prev":
			script = flux.TemplateFill
		case "linear":
			script = flux.TemplateLinear
		}
	}

	tmpl, err := template.New("").Parse(script)
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("Template parse: %v", err)))
		return
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, vars)
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("Template exec: %v", err)))
		return
	}

	result, err := influxdb.QueryAPI().Query(context.Background(), buffer.String())
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("InfluxDB Query: %v", err)))
		return
	}

	rows := make([]map[string]interface{}, 0)

	for result.Next() {
		values := result.Record().Values()

		for key, val := range values {
			if t, ok := val.(time.Time); ok {
				values[key] = modutil.TimeFormat(t)
			}
		}

		rows = append(rows, values)
	}

	ctx.JSON(util.Success(gin.H{
		"flux": buffer.String(),
		"rows": rows,
	}))
}

func tryConvertTimeZone(before string) string {
	t, err := modutil.ParseTime(before)
	if err == nil {
		return t.UTC().Format("2006-01-02T15:04:05Z")
	}

	return before
}
