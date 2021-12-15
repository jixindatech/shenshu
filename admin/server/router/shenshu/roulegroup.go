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

type ruleGroupForm struct {
	Name     string         `json:"name" validate:"required"`
	Type     int            `json:"type" validate:"required,min=1,max=2"`
	Priority int            `json:"priority" validate:"required,min=1"`
	Status   int            `json:"status" validate:"required,min=1,max=4"`
	Level    int            `json:"level"`
	Decoder  datatypes.JSON `json:"decoder" validate:"required"`

	Remark string `json:"remark"`
}

func AddRuleGroup(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     ruleGroupForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}

	rgSrv := service.RuleGroup{
		Name:     form.Name,
		Type:     form.Type,
		Priority: form.Priority,
		Status:   form.Status,
		Level:    form.Level,
		Decoder:  form.Decoder,
		Remark:   form.Remark,
	}
	err = rgSrv.Save()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.RuleGroupAddFailed
		log.Logger.Error("rulegroup", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}

func DeleteRuleGroup(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindUriAndValid(c, &formId)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}

	rgSrv := service.RuleGroup{
		ID: formId.ID,
	}
	err = rgSrv.Delete()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.RuleGroupDeleteFailed
		log.Logger.Error("rulegroup", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}

type queryRuleGroupForm struct {
	Name     string `form:"name" validate:"omitempty,max=254"`
	Type     int    `form:"type" validate:"omitempty,min=1,max=2"`
	Page     int    `form:"page" validate:"min=0"`
	PageSize int    `form:"size" validate:"required,gte=10,lte=50"`
}

func GetRuleGroups(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryRuleGroupForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}

	rgSrv := service.RuleGroup{
		Name:     form.Name,
		Type:     form.Type,
		PageSize: form.PageSize,
		Page:     form.Page,
	}

	data := make(map[string]interface{})
	user, total, err := rgSrv.GetList()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.RuleGroupAddFailed
		log.Logger.Error("rulegroup", zap.String("err", err.Error()))
	} else {
		data["list"] = user
		data["total"] = total
	}

	appG.Response(httpCode, errCode, "", data)
}

func GetRuleGroup(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     app.IDForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindUriAndValid(c, &form)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}

	rgSrv := service.RuleGroup{
		ID: form.ID,
	}
	user, err := rgSrv.Get()
	if err != nil {
		log.Logger.Error("rulegroup", zap.String("err", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.UserGetFailed
	}

	data := make(map[string]interface{})
	if user != nil && user.ID == 0 {
		httpCode = http.StatusInternalServerError
		errCode = e.UserGetFailed
	} else {
		data["item"] = user
	}

	appG.Response(httpCode, errCode, "", data)
}

func UpdateRuleGroup(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     ruleGroupForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindUriAndValid(c, &formId)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}

	err = app.BindAndValid(c, &form)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}

	rgSrv := service.RuleGroup{
		ID:       formId.ID,
		Name:     form.Name,
		Priority: form.Priority,
		Status:   form.Status,
		Level:    form.Level,
		Decoder:  form.Decoder,
		Remark:   form.Remark,
	}

	err = rgSrv.Save()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.RuleGroupPutFailed
		log.Logger.Error("rulegroup", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}
