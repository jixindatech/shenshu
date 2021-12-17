package shenshu

import (
	"admin/core/log"
	"admin/server/pkg/app"
	"admin/server/pkg/e"
	"admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/datatypes"
	"net/http"
)

type ruleSpecificForm struct {
	Name     string         `json:"name" validate:"required,max=254"`
	Rules    datatypes.JSON `form:"rules" valid:"Required;"`
	Action   int            `form:"action" valid:"Range(1,4)"`
	Priority int            `form:"priority" valid:"Min(1)"`
	Status   int            `form:"status" valid:"Range(0,1)"`
	Remark   string         `form:"remark"`
}

func AddRuleSpecific(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     ruleSpecificForm
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

	ruleSrv := service.RuleSpecific{
		Name:      form.Name,
		RuleGroup: formId.ID,
		Rules:     form.Rules,
		Priority:  form.Priority,
		Action:    form.Action,
		Status:    form.Status,
		Remark:    form.Remark,
	}

	err = ruleSrv.Save()
	if err != nil {
		log.Logger.Error("rulespecific", zap.String("add", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleSpecificAddFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

type queryRuleSpecificForm struct {
	Name     string `form:"name" validate:"max=254"`
	Page     int    `form:"page" validate:"min=1"`
	PageSize int    `form:"size" validate:"required,min=1,max=50"`
}

func GetRuleSpecifics(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     queryRuleSpecificForm
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

	ruleSrv := service.RuleSpecific{
		RuleGroup: formId.ID,
		Name:      form.Name,
		Page:      form.Page,
		PageSize:  form.PageSize,
	}
	rules, count, err := ruleSrv.GetList()
	if err != nil {
		log.Logger.Error("rulespecific", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleSpecificGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = rules
	data["total"] = count
	appG.Response(httpCode, errCode, "", data)
}

func GetRuleSpecific(c *gin.Context) {
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

	ruleSrv := service.RuleSpecific{
		ID: formId.ID,
	}
	rule, err := ruleSrv.Get()
	if err != nil {
		log.Logger.Error("rulespecific", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleSpecificGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["item"] = rule
	appG.Response(httpCode, errCode, "", data)
}

func UpdateRuleSpecific(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     ruleSpecificForm
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

	ruleSrv := service.RuleSpecific{
		ID:       formId.ID,
		Name:     form.Name,
		Rules:    form.Rules,
		Priority: form.Priority,
		Action:   form.Action,
		Status:   form.Status,
		Remark:   form.Remark,
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

func DeleteRuleSpecific(c *gin.Context) {
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

	ruleSrv := service.RuleSpecific{
		ID: formId.ID,
	}
	err = ruleSrv.Delete()
	if err != nil {
		log.Logger.Error("rulespecific", zap.String("delete", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleBatchDeleteFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}
