package apis

import (
	"go-admin/common/apis"

	"github.com/gin-gonic/gin"
)

type Article struct {
	apis.Api
}

// GetArticleList 获取文章列表
func (e Article) GetArticleList(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK("hello world ！", "success")
}
