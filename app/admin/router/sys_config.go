package router

import (
	"go-admin/app/admin/apis"
	middleware2 "go-admin/common/middleware"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysConfigRouter)
}

// 需认证的路由代码
func registerSysConfigRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysConfig{}
	r := v1.Group("/config").Use(authMiddleware.MiddlewareFunc()).Use(middleware2.AuthCheckRole())
	{

		r.GET("", api.GetSysConfigList)
		r.GET("/:id", api.GetSysConfig)
		r.POST("", api.InsertSysConfig)
		r.PUT("/:id", api.UpdateSysConfig)
		r.DELETE("", api.DeleteSysConfig)
	}

	r1 := v1.Group("/configKey").Use(authMiddleware.MiddlewareFunc())
	{
		r1.GET("/:configKey", api.GetSysConfigByKEYForService)
	}

	r2 := v1.Group("/app-config")
	{
		r2.GET("", api.GetSysConfigBySysApp)
	}

}
