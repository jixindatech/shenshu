package router

import (
	"admin/core/log"
	"admin/server/pkg/app"
	"admin/server/router/shenshu"
	"admin/server/router/system"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

func Setup(mode string) (g *gin.Engine, err error) {
	err = app.SetupValidate()
	if err != nil {
		return nil, err
	}

	r := gin.New()
	gin.SetMode(mode)
	r.Use(ginzap.Ginzap(log.Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(log.Logger, true))

	authMiddleware, err := system.GetJwtMiddleWare(system.Login, system.Logout)
	if err != nil {
		return nil, err
	}
	r.NoRoute(authMiddleware.MiddlewareFunc(), system.NoRoute)

	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth := authMiddleware.MiddlewareFunc
	apis := r.Group("/system", auth())
	{
		apis.GET("/user/refresh_token", authMiddleware.RefreshHandler)
		apis.POST("/user/logout", authMiddleware.LogoutHandler)
		apis.GET("/user/info", system.GetUserInfo)

		apis.POST("/user", system.AddUser)
		apis.GET("/user", system.GetUsers)
		apis.GET("/user/:id", system.GetUser)
		apis.PUT("/user/:id", system.UpdateUser)
		apis.PUT("/user", system.UpdateUserInfo)
		apis.PUT("/user/password/:id", system.UpdateUserPassword)
		apis.DELETE("/user/:id", system.DeleteUser)

		apis.GET("/email", system.GetEmail)
		apis.POST("/email", system.AddEmail)
		apis.PUT("/email/:id", system.UpdateEmail)

		apis.GET("/ldap", system.GetLdap)
		apis.POST("/ldap", system.AddLdap)
		apis.PUT("/ldap/:id", system.UpdateLdap)

		apis.GET("/txsms", system.GetTxsms)
		apis.POST("/txsms", system.AddTxsms)
		apis.PUT("/txsms/:id", system.UpdateTxsms)

		apis.POST("/msg", system.AddMsg)
		apis.GET("/msg", system.GetMsgs)
		apis.GET("/msg/:id", system.GetMsg)
		apis.PUT("/msg/:id", system.UpdateMsg)
		apis.DELETE("/msg/:id", system.DeleteMsg)
		apis.POST("/msg/:id/user", system.SendMsg)

	}

	shenshuApis := r.Group("/shenshu", auth())
	{
		shenshuApis.POST("/ssl", shenshu.AddSSL)
		shenshuApis.GET("/ssl", shenshu.GetSSLs)
		shenshuApis.GET("/ssl/:id", shenshu.GetSSL)
		shenshuApis.PUT("/ssl/:id", shenshu.UpdateSSL)
		shenshuApis.DELETE("/ssl/:id", shenshu.DeleteSSL)

		shenshuApis.POST("/upstream", shenshu.AddUpstream)
		shenshuApis.GET("/upstream", shenshu.GetUpstreams)
		shenshuApis.GET("/upstream/:id", shenshu.GetUpstream)
		shenshuApis.PUT("/upstream/:id", shenshu.UpdateUpstream)
		shenshuApis.DELETE("/upstream/:id", shenshu.DeleteUpstream)

		shenshuApis.POST("/site", shenshu.AddSite)
		shenshuApis.GET("/site", shenshu.GetSites)
		shenshuApis.GET("/site/:id", shenshu.GetSite)
		shenshuApis.PUT("/site/:id", shenshu.UpdateSite)
		shenshuApis.DELETE("/site/:id", shenshu.DeleteSite)
	}
	return r, nil
}
