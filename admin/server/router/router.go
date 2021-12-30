package router

import (
	"admin/core/log"
	"admin/server/pkg/app"
	"admin/server/router/nginx"
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

	nginxApis := r.Group("/nginx", auth())
	{
		nginxApis.POST("/ssl", nginx.AddSSL)
		nginxApis.GET("/ssl", nginx.GetSSLs)
		nginxApis.GET("/ssl/:id", nginx.GetSSL)
		nginxApis.PUT("/ssl/:id", nginx.UpdateSSL)
		nginxApis.DELETE("/ssl/:id", nginx.DeleteSSL)

		nginxApis.POST("/upstream", nginx.AddUpstream)
		nginxApis.GET("/upstream", nginx.GetUpstreams)
		nginxApis.GET("/upstream/:id", nginx.GetUpstream)
		nginxApis.PUT("/upstream/:id", nginx.UpdateUpstream)
		nginxApis.DELETE("/upstream/:id", nginx.DeleteUpstream)

		nginxApis.POST("/site", nginx.AddSite)
		nginxApis.GET("/site", nginx.GetSites)
		nginxApis.GET("/site/:id", nginx.GetSite)
		nginxApis.PUT("/site/:id", nginx.UpdateSite)
		nginxApis.DELETE("/site/:id", nginx.DeleteSite)
	}

	shenshuApis := r.Group("/shenshu", auth())
	{
		shenshuApis.POST("/globalip", shenshu.AddGlobalIP)
		shenshuApis.GET("/globalip", shenshu.GetGlobalIPs)
		shenshuApis.GET("/globalip/:id", shenshu.GetGlobalIP)
		shenshuApis.PUT("/globalip/:id", shenshu.UpdateGlobalIP)
		shenshuApis.DELETE("/globalip/:id", shenshu.DeleteGlobalIP)

		shenshuApis.POST("/site/:id/ip", shenshu.AddIP)
		shenshuApis.GET("/site/:id/ip", shenshu.GetIPs)
		shenshuApis.GET("/site/ip/:id", shenshu.GetIP)
		shenshuApis.PUT("/site/ip/:id", shenshu.UpdateIP)
		shenshuApis.DELETE("/site/ip/:id", shenshu.DeleteIP)

		shenshuApis.POST("/site/:id/cc", shenshu.AddCC)
		shenshuApis.GET("/site/:id/cc", shenshu.GetCCs)
		shenshuApis.GET("/site/cc/:id", shenshu.GetCC)
		shenshuApis.PUT("/site/cc/:id", shenshu.UpdateCC)
		shenshuApis.DELETE("/site/cc/:id", shenshu.DeleteCC)

		shenshuApis.POST("/batchgroup", shenshu.AddBatchGroup)
		shenshuApis.GET("/batchgroup", shenshu.GetBatchGroups)
		shenshuApis.GET("/batchgroup/:id", shenshu.GetBatchGroup)
		shenshuApis.PUT("/batchgroup/:id", shenshu.UpdateBatchGroup)
		shenshuApis.DELETE("/batchgroup/:id", shenshu.DeleteBatchGroup)

		shenshuApis.POST("/batchgroup/:id/rule", shenshu.AddRuleBatch)
		shenshuApis.GET("/batchgroup/:id/rule", shenshu.GetRuleBatchs)
		shenshuApis.GET("/batchgroup/rule/:id", shenshu.GetRuleBatch)
		shenshuApis.PUT("/batchgroup/rule/:id", shenshu.UpdateRuleBatch)
		shenshuApis.DELETE("/batchgroup/rule/:id", shenshu.DeleteRuleBatch)

		shenshuApis.POST("/specificgroup", shenshu.AddSpecificGroup)
		shenshuApis.GET("/specificgroup", shenshu.GetSpecificGroups)
		shenshuApis.GET("/specificgroup/:id", shenshu.GetSpecificGroup)
		shenshuApis.PUT("/specificgroup/:id", shenshu.UpdateSpecificGroup)
		shenshuApis.DELETE("/specificgroup/:id", shenshu.DeleteSpecificGroup)

		shenshuApis.POST("/specificgroup/:id/rule", shenshu.AddRuleSpecific)
		shenshuApis.GET("/specificgroup/:id/rule", shenshu.GetRuleSpecifics)
		shenshuApis.GET("/specificgroup/rule/:id", shenshu.GetRuleSpecific)
		shenshuApis.PUT("/specificgroup/rule/:id", shenshu.UpdateRuleSpecific)
		shenshuApis.DELETE("/specificgroup/rule/:id", shenshu.DeleteRuleSpecific)

		shenshuApis.PUT("/site/:id/rulegroup", shenshu.UpdateSiteRuleGroup)
		shenshuApis.GET("/site/:id/rulegroup", shenshu.GetSiteRuleGroup)

		shenshuApis.POST("/site/:id/enable", shenshu.EnableSiteConfig)

		shenshuApis.GET("/event/cc", shenshu.GetCCEvents)
	}
	return r, nil
}
