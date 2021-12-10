package nginx

import (
	"admin/core/log"
	"admin/server/pkg/app"
	"admin/server/pkg/e"
	"admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type sslForm struct {
	Name   string `json:"name" validate:"required,max=254"`
	Pub    string `json:"pub" validate:"required,max=5120"`
	Pri    string `json:"pri"  validate:"required,max=5120"`
	Remark string `json:"remark" validate:"max=254"`
}

func AddSSL(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     sslForm
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

	ssl := service.SSL{
		Name:   form.Name,
		Pub:    strings.TrimSpace(form.Pub),
		Pri:    strings.TrimSpace(form.Pri),
		Remark: form.Remark,
	}
	err = ssl.Save()
	if err != nil {
		log.Logger.Error("ssl", zap.String("add", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.SSLAddFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func GetSSL(c *gin.Context) {
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

	ssl := service.SSL{
		ID: formId.ID,
	}
	idSSL, err := ssl.Get()
	if err != nil {
		log.Logger.Error("user", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.UserGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["item"] = idSSL
	appG.Response(httpCode, errCode, "", data)
}

type querySSLForm struct {
	Name     string `form:"name" validate:"max=254"`
	Page     int    `form:"page" validate:"required,min=1,max=50"`
	PageSize int    `form:"size" validate:"required,min=1"`
}

func GetSSLs(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     querySSLForm
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

	ssl := service.SSL{
		Name:     form.Name,
		Page:     form.Page,
		PageSize: form.PageSize,
	}
	ssls, count, err := ssl.GetList()
	if err != nil {
		log.Logger.Error("ssl", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.SSLGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = ssls
	data["total"] = count
	appG.Response(httpCode, errCode, "", data)
}

func UpdateSSL(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     sslForm
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

	ssl := service.SSL{
		ID:     formId.ID,
		Name:   form.Name,
		Pub:    form.Pub,
		Pri:    form.Pri,
		Remark: form.Remark,
	}
	err = ssl.Save()
	if err != nil {
		log.Logger.Error("ssl", zap.String("put", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.SSLPutFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func DeleteSSL(c *gin.Context) {
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

	ssl := service.SSL{
		ID: formId.ID,
	}
	err = ssl.Delete()
	if err != nil {
		log.Logger.Error("ssl", zap.String("delete", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.SSLDeleteFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}
