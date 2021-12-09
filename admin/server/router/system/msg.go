package system

import (
	"admin/core/log"
	"admin/server/pkg/app"
	"admin/server/pkg/e"
	"admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type msgForm struct {
	Name    string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required"`
	Remark  string `json:"remark"`
}

func AddMsg(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     msgForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}

	msgSrv := service.Msg{
		Name:    form.Name,
		Content: form.Content,
		Remark:  form.Remark,
	}
	err = msgSrv.Save()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.MsgAddFailed
		log.Logger.Error("msg", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}

func DeleteMsg(c *gin.Context) {
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

	msgSrv := service.Msg{
		ID: formId.ID,
	}
	err = msgSrv.Delete()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.MsgDeleteFailed
		log.Logger.Error("msg", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}

type queryMsgsForm struct {
	Name     string `form:"name" validate:"omitempty,max=254"`
	Page     int    `form:"page" validate:"required,gte=1"`
	PageSize int    `form:"size" validate:"required,gte=10,lte=50"`
}

func GetMsgs(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryMsgsForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}

	msgSrv := service.Msg{
		Name:     form.Name,
		PageSize: form.PageSize,
		Page:     form.Page,
	}

	data := make(map[string]interface{})
	msg, total, err := msgSrv.GetList()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.MsgAddFailed
		log.Logger.Error("msg", zap.String("err", err.Error()))
	} else {
		data["list"] = msg
		data["total"] = total
	}

	appG.Response(httpCode, errCode, "", data)
}

func GetMsg(c *gin.Context) {
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

	msgSrv := service.Msg{
		ID: form.ID,
	}
	msg, err := msgSrv.Get()
	if err != nil {
		log.Logger.Error("msg", zap.String("err", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.MsgGetFailed
	}

	data := make(map[string]interface{})
	if msg != nil && msg.ID == 0 {
		httpCode = http.StatusInternalServerError
		errCode = e.MsgGetFailed
	} else {
		data["item"] = msg
	}

	appG.Response(httpCode, errCode, "", data)
}

func UpdateMsg(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     msgForm
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

	msgSrv := service.Msg{
		ID:      formId.ID,
		Name:    form.Name,
		Content: form.Content,
		Remark:  form.Remark,
	}
	err = msgSrv.Save()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.MsgUpdateFailed
		log.Logger.Error("msg", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}

type msgUsersForm struct {
	Users []uint `json:"users" validate:"required"`
}

func SendMsg(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     msgUsersForm
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

	msgSrv := service.Msg{
		ID: formId.ID,
	}

	err = msgSrv.SendMsgs(form.Users)
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.MsgAddFailed
		log.Logger.Error("msg", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}
