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

type queryRuleEventForm struct {
	SiteID   uint   `form:"site"`
	SiteName string `form:"name"`
	Start    int64  `form:"start" validate:"required"`
	End      int64  `form:"end" validate:"required"`
	Page     int    `form:"page" validate:"required,min=1"`
	PageSize int    `form:"size" validate:"required,min=1,max=50"`
}

func GetBatchRuleEvents(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryRuleEventForm
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

	ruleSrv := service.BatchRuleEvent{
		SiteID:   form.SiteID,
		SiteName: form.SiteName,
		Start:    form.Start,
		End:      form.End,
		Page:     form.Page,
		PageSize: form.PageSize,
	}
	events, err := ruleSrv.GetList()
	if err != nil {
		log.Logger.Error("RuleEvent", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.BatchRuleEventGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = events["data"]
	data["total"] = events["count"]
	appG.Response(httpCode, errCode, "", data)
}

func GetSpecificRuleEvents(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryRuleEventForm
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

	ruleSrv := service.SpecificRuleEvent{
		SiteID:   form.SiteID,
		SiteName: form.SiteName,
		Start:    form.Start,
		End:      form.End,
		Page:     form.Page,
		PageSize: form.PageSize,
	}
	events, err := ruleSrv.GetList()
	if err != nil {
		log.Logger.Error("RuleEvent", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.SpecificRuleEventGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = events["data"]
	data["total"] = events["count"]
	appG.Response(httpCode, errCode, "", data)
}

type querySiteEventInfoForm struct {
	SiteId uint  `form:"site" validate:"min=0"`
	Start  int64 `form:"start" validate:"required"`
	End    int64 `form:"end" validate:"required"`
}

func GetSiteRuleEventInfo(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     querySiteEventInfoForm
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

	ruleSrv := service.Site{
		ID: form.SiteId,
	}
	events, err := ruleSrv.GetRuleEventInfo(form.Start, form.End)
	if err != nil {
		log.Logger.Error("SiteEvent", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.SiteRuleEventGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["item"] = events
	appG.Response(httpCode, errCode, "", data)
}
