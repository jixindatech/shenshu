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

type queryCCEventForm struct {
	Start    int64 `form:"start" validate:"required"`
	End      int64 `form:"end" validate:"required"`
	Page     int   `form:"page" validate:"required,min=1"`
	PageSize int   `form:"size" validate:"required,min=1,max=50"`
}

func GetCCEvents(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryCCEventForm
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

	ccSrv := service.CCEvent{
		Start:    form.Start,
		End:      form.End,
		Page:     form.Page,
		PageSize: form.PageSize,
	}
	events, err := ccSrv.GetList()
	if err != nil {
		log.Logger.Error("CCEvent", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.CCEventGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = events["data"]
	data["total"] = events["count"]
	appG.Response(httpCode, errCode, "", data)
}
