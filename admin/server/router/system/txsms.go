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

type txsmsForm struct {
	Host      string `json:"host" validate:"required"`
	SecretId  string `json:"secretId" validate:"required"`
	SecretKey string `json:"secretKey" validate:"required"`
}

func AddTxsms(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     txsmsForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}
	txsmsSrv := service.Txsms{
		Host:      form.Host,
		SecretId:  form.SecretId,
		SecretKey: form.SecretKey,
	}

	err = txsmsSrv.Save()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.TxsmsAddFailed
		log.Logger.Error("txsms", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}

func GetTxsms(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	txsmsSrv := service.Txsms{}
	txsms, err := txsmsSrv.Get()
	if err != nil {
		log.Logger.Error("txsms", zap.String("err", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.TxsmsGetFailed
	}

	data := make(map[string]interface{})
	data["item"] = txsms

	appG.Response(httpCode, errCode, "", data)

}

func UpdateTxsms(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formID   app.IDForm
		form     txsmsForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindUriAndValid(c, &formID)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}

	if formID.ID != 1 {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, "", nil)
		return
	}

	err = app.BindAndValid(c, &form)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}
	txsmsSrv := service.Txsms{
		ID:        formID.ID,
		Host:      form.Host,
		SecretId:  form.SecretId,
		SecretKey: form.SecretKey,
	}

	err = txsmsSrv.Save()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.TxsmsUpdateFailed
		log.Logger.Error("txsms", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}
