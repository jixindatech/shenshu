package shenshu

import (
	"admin/core/log"
	"admin/server/pkg/app"
	"admin/server/pkg/e"
	"admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type ccForm struct {
	Name      string `json:"name" validate:"required,max=254"`
	Site      uint   `json:"site" validate:"required,min=1"`
	Mode      string `json:"mode" validate:"required"`
	Method    string `json:"method" validate:"required"`
	URI       string `json:"uri" validate:"required"`
	Match     string `json:"match" validate:"required"`
	Threshold int    `json:"threshold" validate:"required,min=1"`
	Duration  int    `json:"duration" validate:"required,min=1"`
	Action    string `json:"action" validate:"required"`
	Remark    string `json:"remark" validate:"max=254"`
}

func AddCC(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     ccForm
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

	ccSrv := service.CC{
		Name:      form.Name,
		Site:      form.Site,
		Mode:      form.Mode,
		Method:    form.Method,
		URI:       form.URI,
		Match:     form.Match,
		Threshold: form.Threshold,
		Duration:  form.Duration,
		Action:    form.Action,
		Remark:    form.Remark,
	}

	err = ccSrv.Save()
	if err != nil {
		log.Logger.Error("CC", zap.String("add", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.CCAddFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func GetCC(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		httpCode = http.StatusOK
		formId   app.IDForm
		errCode  = e.SUCCESS
	)

	err := app.BindUriAndValid(c, &formId)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	ccSrv := service.CC{
		ID: formId.ID,
	}
	cc, err := ccSrv.Get()
	if err != nil {
		log.Logger.Error("CC", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.CCGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["item"] = cc
	appG.Response(httpCode, errCode, "", data)
}

type queryCCForm struct {
	Name     string `form:"name" validate:"max=254"`
	Page     int    `form:"page" validate:"min=1"`
	PageSize int    `form:"size" validate:"required,min=1,max=50"`
}

func GetCCs(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryCCForm
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

	ccSrv := service.CC{
		Name:     form.Name,
		Page:     form.Page,
		PageSize: form.PageSize,
	}
	ccs, count, err := ccSrv.GetList()
	if err != nil {
		log.Logger.Error("CC", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.CCGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = ccs
	data["total"] = count
	appG.Response(httpCode, errCode, "", data)
}

func UpdateCC(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     ccForm
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

	ccSrv := service.CC{
		ID:        formId.ID,
		Name:      form.Name,
		Site:      form.Site,
		Mode:      form.Mode,
		Method:    form.Method,
		URI:       form.URI,
		Match:     form.Match,
		Threshold: form.Threshold,
		Duration:  form.Duration,
		Action:    form.Action,
		Remark:    form.Remark,
	}

	err = ccSrv.Save()
	if err != nil {
		log.Logger.Error("CC", zap.String("add", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.CCAddFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}
