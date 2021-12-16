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

func EnableSiteRuleGroup(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		// formId   app.IDForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	appG.Response(httpCode, errCode, "", nil)
}

type siteRuleGroupForm struct {
	IDs []uint `json:"ids" validate:"required,dive,min=1"`
}

func UpdateSiteRuleGroup(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     siteRuleGroupForm
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

	siteSrv := &service.Site{
		ID:  formId.ID,
		Ids: form.IDs,
	}
	err = siteSrv.UpdatRuleGroup()
	if err != nil {
		log.Logger.Error("SiteRuleGroup", zap.String("update", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.SitePutRuleGroupFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}
	appG.Response(httpCode, errCode, "", nil)
}

func GetSiteRuleGroup(c *gin.Context) {
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

	siteSrv := &service.Site{
		ID: formId.ID,
	}

	ids, err := siteSrv.GetRuleGroup()
	if err != nil {
		log.Logger.Error("SiteRuleGroup", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.SiteGetRuleGroupFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	if ids == nil {
		ids = []uint{}
	}

	data := make(map[string]interface{})
	data["ids"] = ids

	appG.Response(httpCode, errCode, "", data)
}

func EnableSiteConfig(c *gin.Context) {
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

	siteSrv := &service.Site{
		ID: formId.ID,
	}

	err = siteSrv.Enable()
	if err != nil {
		log.Logger.Error("SiteRuleGroup", zap.String("enable", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.SiteEnableSiteConfig
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}
