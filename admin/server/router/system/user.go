package system

import (
	"admin/core/log"
	"admin/core/rbac"
	"admin/server/pkg/app"
	"admin/server/pkg/e"
	"admin/server/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

const (
	IDENTITY = "id"
)

type jwtUser struct {
	ID       uint
	Username string
	Role     string
}

type userForm struct {
	UserName    string `json:"username" validate:"required"`
	DisplayName string `json:"displayName"`
	LoginType   string `json:"loginType" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Phone       string `json:"phone" validate:"required,phone"`
	Status      int    `json:"status" validate:"required,gte=1,lte=2"`
	Role        string `json:"role" validate:"required"`
	Remark      string `json:"remark"`
}

func AddUser(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     userForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}

	userSrv := service.User{
		Username:    form.UserName,
		DisplayName: form.DisplayName,
		LoginType:   form.LoginType,
		Email:       form.Email,
		Phone:       form.Phone,
		Status:      form.Status,
		Role:        form.Role,
		Remark:      form.Remark,
	}
	err = userSrv.Save()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.UserAddFailed
		log.Logger.Error("user", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}

func DeleteUser(c *gin.Context) {
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

	userSrv := service.User{
		ID: formId.ID,
	}
	err = userSrv.Delete()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.UserDeleteFailed
		log.Logger.Error("user", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}

type queryUsersForm struct {
	UserName string `form:"username" validate:"omitempty,max=254"`
	Page     int    `form:"page" validate:"required,gte=1"`
	PageSize int    `form:"size" validate:"required,gte=10,lte=50"`
}

func GetUsers(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryUsersForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}

	userSrv := service.User{
		Username: form.UserName,
		PageSize: form.PageSize,
		Page:     form.Page,
	}

	data := make(map[string]interface{})
	user, total, err := userSrv.GetList()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.UserAddFailed
		log.Logger.Error("user", zap.String("err", err.Error()))
	} else {
		data["list"] = user
		data["total"] = total
	}

	appG.Response(httpCode, errCode, "", data)
}

func GetUserInfo(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	userJwt, exist := c.Get(IDENTITY)
	if !exist {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, errCode, "user jwt claim is empty", nil)
		return
	}

	userClaim := userJwt.(*jwtUser)
	id := userClaim.ID
	userSrv := service.User{
		ID: id,
	}

	user, err := userSrv.Get()
	if err != nil {
		log.Logger.Error("user", zap.String("err", err.Error()))
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	if user.ID == 0 {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, errCode, "user jwt claim is not correct", nil)
		return
	}

	data := make(map[string]interface{})
	data["introduction"] = user.Remark
	data["name"] = user.Username
	data["info"] = user

	var roles []string
	roles = append(roles, user.Role)
	data["roles"] = roles

	routes := rbac.GetRoleRoutes(user.Role)
	if routes != nil {
		data["routes"] = routes
	} else {
		data["routes"] = []struct{}{}
	}

	api := rbac.GetRoleApi(user.Role)
	if len(api) > 0 {
		data["api"] = api
	} else {
		data["api"] = []struct{}{}
	}

	appG.Response(httpCode, errCode, "", data)
	return
}

func GetUser(c *gin.Context) {
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

	userSrv := service.User{
		ID: form.ID,
	}
	user, err := userSrv.Get()
	if err != nil {
		log.Logger.Error("user", zap.String("err", err.Error()))
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

func UpdateUser(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     userForm
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

	userSrv := service.User{
		ID:          formId.ID,
		Username:    form.UserName,
		DisplayName: form.DisplayName,
		LoginType:   form.LoginType,
		Email:       form.Email,
		Phone:       form.Phone,
		Status:      form.Status,
		Role:        form.Role,
		Remark:      form.Remark,
	}

	err = userSrv.Save()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.UserUpdateFailed
		log.Logger.Error("user", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}

type userInfoForm struct {
	Username    string `json:"username" validate:"required"`
	DisplayName string `json:"displayName"`
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
	Email       string `json:"email" validate:"required,email"`
	Phone       string `json:"phone" validate:"required,phone"`
}

func UpdateUserInfo(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     userInfoForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, e.ERROR, err.Error(), nil)
		return
	}

	userJwt, exist := c.Get(IDENTITY)
	if !exist {
		httpCode = http.StatusBadRequest
		appG.Response(httpCode, errCode, "user jwt claim is empty", nil)
		return
	}

	userClaim := userJwt.(*jwtUser)
	id := userClaim.ID

	userSrv := service.User{
		ID:          id,
		Username:    form.Username,
		DisplayName: form.DisplayName,
		Email:       form.Email,
		Phone:       form.Phone,
	}
	if len(form.OldPassword) > 0 && len(form.NewPassword) > 0 {
		_, err := userSrv.GetLoginUser(true)
		if err != nil {
			httpCode = http.StatusInternalServerError
			errCode = e.UserUpdateFailed
			log.Logger.Error("user", zap.String("err", err.Error()))
		}

		userSrv.Password = form.NewPassword
	}

	err = userSrv.UpdateUserInfo()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.UserUpdateFailed
		log.Logger.Error("user", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}

func UpdateUserPassword(c *gin.Context) {
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

	userSrv := service.User{
		ID: formId.ID,
	}
	err = userSrv.UpdatePassword()
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = e.UserUpdateFailed
		log.Logger.Error("user", zap.String("err", err.Error()))
	}

	appG.Response(httpCode, errCode, "", nil)
}

type loginForm struct {
	UserName string `json:"username" validate:"required,gt=1"`
	Password string `json:"password" validate:"required,gt=1"`
}

func Login(c *gin.Context) (interface{}, error) {
	var form loginForm

	err := app.BindAndValid(c, &form)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	userSrv := service.User{
		Username: form.UserName,
		Password: form.Password,
	}

	user, err := userSrv.GetLoginUser(false)
	if err != nil {
		/* For security */
		log.Logger.Error("user", zap.String("err", err.Error()))
		return nil, fmt.Errorf("%s", "user login failed")
	}

	jwtuser := &jwtUser{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}

	return jwtuser, nil
}

func Logout(c *gin.Context) {

}
