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

type ldapForm struct {
	Type     string `json:"type" validate:"required"`
	Host     string `json:"host" validate:"required"`
	Port     int    `json:"port" validate:"required,min=1,max=65535"`
	DN       string `json:"dn" validate:"required"`
	BaseDN   string `json:"basedn" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

func AddLdap(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     ldapForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}
	ldapSrv := service.Ldap{
		Type:     form.Type,
		Host:     form.Host,
		Port:     form.Port,
		DN:       form.DN,
		BaseDN:   form.BaseDN,
		Password: form.Password,
	}

	err = ldapSrv.Save()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.LdapAddFailed
		log.Logger.Error("ldap", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}

func GetLdap(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	ldapSrv := service.Ldap{}
	ldap, err := ldapSrv.Get()
	if err != nil {
		log.Logger.Error("ldap", zap.String("err", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.LdapGetFailed
	}

	data := make(map[string]interface{})
	data["item"] = ldap

	appG.Response(httpCode, errCode, "", data)

}

func UpdateLdap(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formID   app.IDForm
		form     ldapForm
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
	ldapSrv := service.Ldap{
		ID:       formID.ID,
		Type:     form.Type,
		Host:     form.Host,
		Port:     form.Port,
		DN:       form.DN,
		BaseDN:   form.BaseDN,
		Password: form.Password,
	}

	err = ldapSrv.Save()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.LdapUpdateFailed
		log.Logger.Error("ldap", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}
