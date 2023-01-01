package router

import (
	"go-admin/app/admin/apis"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerArticleRouter)
}

// 需认证的路由代码
func registerArticleRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Article{}
	r := v1.Group("")
	{
		r.GET("/articleList", api.GetArticleList)
	}
}
