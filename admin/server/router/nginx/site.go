package nginx

import (
	"admin/core/log"
	"admin/server/pkg/app"
	"admin/server/pkg/e"
	"admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type siteForm struct {
	Name       string `json:"name" validate:"required,max=254"`
	Host       string `json:"host" validate:"required"`
	Path       string `json:"path" validate:"required"`
	UptreamRef uint   `json:"upstreamRef" validate:"required"`
	Remark     string `json:"remark" validate:"max=254"`
}

func AddSite(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     siteForm
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

	Site := service.Site{
		Name:        form.Name,
		Host:        form.Host,
		Path:        form.Path,
		UpstreamRef: form.UptreamRef,

		Remark: form.Remark,
	}

	err = Site.Save()
	if err != nil {
		log.Logger.Error("Site", zap.String("add", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.SiteAddFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func GetSite(c *gin.Context) {
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

	Site := service.Site{
		ID: formId.ID,
	}
	idSite, err := Site.Get()
	if err != nil {
		log.Logger.Error("Site", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.SiteGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["item"] = idSite
	appG.Response(httpCode, errCode, "", data)
}

type querySiteForm struct {
	Name     string `form:"name" validate:"max=254"`
	Page     int    `form:"page" validate:"required,min=1,max=50"`
	PageSize int    `form:"size" validate:"required,min=1"`
}

func GetSites(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     querySiteForm
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

	Site := service.Site{
		Name:     form.Name,
		Page:     form.Page,
		PageSize: form.PageSize,
	}
	Sites, count, err := Site.GetList()
	if err != nil {
		log.Logger.Error("Site", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.SiteGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = Sites
	data["total"] = count
	appG.Response(httpCode, errCode, "", data)
}

func UpdateSite(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     siteForm
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

	Site := service.Site{
		ID:          formId.ID,
		Host:        form.Host,
		Path:        form.Path,
		UpstreamRef: form.UptreamRef,

		Remark: form.Remark,
	}
	err = Site.Save()
	if err != nil {
		log.Logger.Error("Site", zap.String("put", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.SitePutFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func DeleteSite(c *gin.Context) {
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

	Site := service.Site{
		ID: formId.ID,
	}
	err = Site.Delete()
	if err != nil {
		log.Logger.Error("Site", zap.String("delete", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.SiteDeleteFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}
