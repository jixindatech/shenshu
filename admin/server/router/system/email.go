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

type emailForm struct {
	Host     string `json:"host" validate:"required"`
	Port     int    `json:"port" validate:"required,min=1,max=65535"`
	Sender   string `json:"sender" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func AddEmail(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     emailForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}
	emailSrv := service.Email{
		Host:     form.Host,
		Port:     form.Port,
		Sender:   form.Sender,
		Password: form.Password,
	}

	err = emailSrv.Save()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.EmailAddFailed
		log.Logger.Error("email", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}

func GetEmail(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	emailSrv := service.Email{}
	email, err := emailSrv.Get()
	if err != nil {
		log.Logger.Error("email", zap.String("err", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.EmailGetFailed
	}

	data := make(map[string]interface{})
	data["item"] = email

	appG.Response(httpCode, errCode, "", data)

}

func UpdateEmail(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formID   app.IDForm
		form     emailForm
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
	emailSrv := service.Email{
		ID:       formID.ID,
		Host:     form.Host,
		Port:     form.Port,
		Sender:   form.Sender,
		Password: form.Password,
	}

	err = emailSrv.Save()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.EmailUpdateFailed
		log.Logger.Error("email", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}
