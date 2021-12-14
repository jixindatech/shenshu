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

type ruleForm struct {
	Name     string         `json:"name" validate:"required,max=254"`
	Rules    datatypes.JSON `json:"rules" validate:"required"`
	Action   int            `json:"action" validate:"min=1,max=4"`
	Priority int            `json:"priority" validate:"min=1"`
	Status   int            `json:"status" validate:"min=1,max=2"`
	Remark   string         `json:"remark"`
}

func AddRule(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     ruleForm
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

	ruleSrv := service.Rule{
		RuleGroup: formId.ID,
		Name:      form.Name,
		Rules:     form.Rules,
		Action:    form.Action,
		Priority:  form.Priority,
		Status:    form.Status,
		Remark:    form.Remark,
	}

	err = ruleSrv.Save()
	if err != nil {
		log.Logger.Error("Rule", zap.String("add", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleAddFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func GetRule(c *gin.Context) {
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

	ruleSrv := service.Rule{
		ID: formId.ID,
	}
	rule, err := ruleSrv.Get()
	if err != nil {
		log.Logger.Error("Rule", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["item"] = rule
	appG.Response(httpCode, errCode, "", data)
}

type queryRuleForm struct {
	Name     string `form:"name" validate:"max=254"`
	Page     int    `form:"page" validate:"min=1"`
	PageSize int    `form:"size" validate:"required,min=1,max=50"`
}

func GetRules(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     queryRuleForm
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

	ruleSrv := service.Rule{
		RuleGroup: formId.ID,
		Name:      form.Name,
		Page:      form.Page,
		PageSize:  form.PageSize,
	}
	rules, count, err := ruleSrv.GetList()
	if err != nil {
		log.Logger.Error("Rule", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = rules
	data["total"] = count
	appG.Response(httpCode, errCode, "", data)
}

func UpdateRule(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     ruleForm
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

	ruleSrv := service.Rule{
		ID:       formId.ID,
		Name:     form.Name,
		Rules:    form.Rules,
		Action:   form.Action,
		Priority: form.Priority,
		Status:   form.Status,
		Remark:   form.Remark,
	}

	err = ruleSrv.Save()
	if err != nil {
		log.Logger.Error("Rule", zap.String("add", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleAddFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func DeleteRule(c *gin.Context) {
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

	ruleSrv := service.Rule{
		ID: formId.ID,
	}
	err = ruleSrv.Delete()
	if err != nil {
		log.Logger.Error("Rule", zap.String("delete", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleDeleteFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}
