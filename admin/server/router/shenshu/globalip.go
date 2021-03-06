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

type globalIpForm struct {
	Name   string   `json:"name" validate:"required,max=254"`
	Type   int      `json:"type" validate:"required,min=1,max=2"`
	IP     []string `json:"ip" validate:"required,dive,ip"`
	Remark string   `json:"remark" validate:"max=254"`
}

func AddGlobalIP(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     globalIpForm
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

	ip := service.GlobalIP{
		Name: form.Name,
		Type: form.Type,
		IP:   form.IP,

		Remark: form.Remark,
	}

	err = ip.Save()
	if err != nil {
		log.Logger.Error("GlobalIP", zap.String("add", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.IPAddFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

type queryGlobalIPForm struct {
	Name     string `form:"name" validate:"max=254"`
	Type     int    `form:"type" validate:"required,min=1,max=2"`
	Page     int    `form:"page" validate:"required,min=1,max=50"`
	PageSize int    `form:"size" validate:"required,min=1"`
}

func GetGlobalIPs(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryGlobalIPForm
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

	ipSrv := service.GlobalIP{
		Name:     form.Name,
		Type:     form.Type,
		Page:     form.Page,
		PageSize: form.PageSize,
	}
	ips, count, err := ipSrv.GetList()
	if err != nil {
		log.Logger.Error("IP", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.IPGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = ips
	data["total"] = count
	appG.Response(httpCode, errCode, "", data)
}

func UpdateGlobalIP(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     globalIpForm
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

	ipSrv := service.GlobalIP{
		ID:   formId.ID,
		Name: form.Name,
		Type: form.Type,
		IP:   form.IP,

		Remark: form.Remark,
	}

	err = ipSrv.Save()
	if err != nil {
		log.Logger.Error("IP", zap.String("put", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.IPPutFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func DeleteGlobalIP(c *gin.Context) {
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

	IPsrv := service.GlobalIP{
		ID: formId.ID,
	}
	err = IPsrv.Delete()
	if err != nil {
		log.Logger.Error("IP", zap.String("delete", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.IPDeleteFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func GetGlobalIP(c *gin.Context) {
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

	ipSrv := service.GlobalIP{
		ID: formId.ID,
	}
	idIP, err := ipSrv.Get()
	if err != nil {
		log.Logger.Error("IP", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.IPGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["item"] = idIP
	appG.Response(httpCode, errCode, "", data)
}
