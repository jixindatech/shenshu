package nginx

import (
	"admin/core/log"
	"admin/server/pkg/app"
	"admin/server/pkg/e"
	"admin/server/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type upstreamForm struct {
	Name           string         `json:"name" validate:"required,max=254"`
	Lb             string         `json:"lb" validate:"required,lb"`
	Key            string         `json:"key"`
	Backend        []service.Node `json:"backend" validate:"required"`
	Retry          int            `json:"retry" validate:"required,min=1"`
	TimeoutConnect int            `json:"timeoutConnect" validate:"required,min=1"`
	TimeoutSend    int            `json:"timeoutSend" validate:"required,min=1"`
	TimeoutReceive int            `json:"timeoutReceive" validate:"required,min=1"`

	Remark string `json:"remark" validate:"max=254"`
}

func AddUpstream(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     upstreamForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	backends, err := json.Marshal(form.Backend)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	if form.Lb == "chash" && len(form.Key) == 0 {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, "chash need key", nil)
		return
	}

	if len(form.Key) > 0 {
		if form.Key != "remote-addr" {
			httpCode = e.InvalidParams
			errCode = e.ERROR
			appG.Response(httpCode, errCode, "LB key is incorrect", nil)
			return
		}
	}

	upstream := service.Upstream{
		Name:           form.Name,
		Lb:             form.Lb,
		Key:            form.Key,
		Backend:        []byte(backends),
		Retry:          form.Retry,
		TimeoutConnect: form.TimeoutConnect,
		TimeoutSend:    form.TimeoutSend,
		TimeoutReceive: form.TimeoutReceive,

		Remark: form.Remark,
	}

	err = upstream.Save()
	if err != nil {
		log.Logger.Error("upstream", zap.String("add", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.UpstreamAddFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func GetUpstream(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindUriAndValid(c, &formId)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	upstream := service.Upstream{
		ID: formId.ID,
	}
	idUpstream, err := upstream.Get()
	if err != nil {
		log.Logger.Error("upstream", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.UpstreamGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["item"] = idUpstream
	appG.Response(httpCode, errCode, "", data)
}

type queryUpstreamForm struct {
	Name     string `form:"name" validate:"max=254"`
	Page     int    `form:"page" validate:"min=0"`
	PageSize int    `form:"size" validate:"min=0,max=50"`
}

func GetUpstreams(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryUpstreamForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	upstream := service.Upstream{
		Name:     form.Name,
		Page:     form.Page,
		PageSize: form.PageSize,
	}
	upstreams, count, err := upstream.GetList()
	if err != nil {
		log.Logger.Error("upstream", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.UpstreamGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = upstreams
	data["total"] = count
	appG.Response(httpCode, errCode, "", data)
}

func UpdateUpstream(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     upstreamForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindUriAndValid(c, &formId)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	err = app.BindAndValid(c, &form)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	backends, err := json.Marshal(form.Backend)
	if err != nil {
		httpCode = e.InvalidParams
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	upstream := service.Upstream{
		ID:             formId.ID,
		Name:           form.Name,
		Lb:             form.Lb,
		Backend:        backends,
		Retry:          form.Retry,
		TimeoutConnect: form.TimeoutConnect,
		TimeoutSend:    form.TimeoutSend,
		TimeoutReceive: form.TimeoutReceive,

		Remark: form.Remark,
	}
	err = upstream.Save()
	if err != nil {
		log.Logger.Error("upstream", zap.String("put", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.UpstreamPutFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func DeleteUpstream(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindUriAndValid(c, &formId)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	upstream := service.Upstream{
		ID: formId.ID,
	}
	err = upstream.Delete()
	if err != nil {
		log.Logger.Error("upstream", zap.String("delete", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.UpstreamDeleteFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}
