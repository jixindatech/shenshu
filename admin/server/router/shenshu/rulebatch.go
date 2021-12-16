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

type ruleBatchForm struct {
	Name    string `json:"name" validate:"required,max=254"`
	Pattern string `json:"pattern" validate:"required,max=254"`
	Action  int    `json:"action" validate:"required,min=1,max=4"`
	Status  int    `json:"status" validate:"required,min=1,max=2"`
	Remark  string `json:"remark"`
}

func AddRuleBatch(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     ruleBatchForm
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

	ruleSrv := service.RuleBatch{
		RuleGroup: formId.ID,
		Name:      form.Name,
		Pattern:   form.Pattern,
		Action:    form.Action,
		Status:    form.Status,
		Remark:    form.Remark,
	}

	err = ruleSrv.Save()
	if err != nil {
		log.Logger.Error("rulebatch", zap.String("add", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleBatchAddFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func GetRuleBatch(c *gin.Context) {
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

	ruleSrv := service.RuleBatch{
		ID: formId.ID,
	}
	rule, err := ruleSrv.Get()
	if err != nil {
		log.Logger.Error("rulebatch", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleBatchGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["item"] = rule
	appG.Response(httpCode, errCode, "", data)
}

type queryRuleBatchForm struct {
	Name     string `form:"name" validate:"max=254"`
	Page     int    `form:"page" validate:"min=1"`
	PageSize int    `form:"size" validate:"required,min=1,max=50"`
}

func GetRuleBatchs(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     queryRuleBatchForm
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

	ruleSrv := service.RuleBatch{
		RuleGroup: formId.ID,
		Name:      form.Name,
		Page:      form.Page,
		PageSize:  form.PageSize,
	}
	rules, count, err := ruleSrv.GetList()
	if err != nil {
		log.Logger.Error("rulebatch", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleBatchGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = rules
	data["total"] = count
	appG.Response(httpCode, errCode, "", data)
}

func UpdateRuleBatch(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     ruleBatchForm
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

	ruleSrv := service.RuleBatch{
		ID:      formId.ID,
		Name:    form.Name,
		Pattern: form.Pattern,
		Action:  form.Action,
		Status:  form.Status,
		Remark:  form.Remark,
	}

	err = ruleSrv.Save()
	if err != nil {
		log.Logger.Error("rulebatch", zap.String("add", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleBatchAddFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func DeleteRuleBatch(c *gin.Context) {
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

	ruleSrv := service.RuleBatch{
		ID: formId.ID,
	}
	err = ruleSrv.Delete()
	if err != nil {
		log.Logger.Error("rulebatch", zap.String("delete", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleBatchDeleteFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}
