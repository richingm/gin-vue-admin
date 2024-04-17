package article

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ArticlesRouter struct {
}

// InitArticlesRouter 初始化 articles表 路由信息
func (s *ArticlesRouter) InitArticlesRouter(Router *gin.RouterGroup) {
	articlesRouter := Router.Group("articles").Use(middleware.OperationRecord())
	articlesRouterWithoutRecord := Router.Group("articles")
	var articlesApi = v1.ApiGroupApp.ArticleApiGroup.ArticlesApi
	{
		articlesRouter.POST("createArticles", articlesApi.CreateArticles)             // 新建articles表
		articlesRouter.DELETE("deleteArticles", articlesApi.DeleteArticles)           // 删除articles表
		articlesRouter.DELETE("deleteArticlesByIds", articlesApi.DeleteArticlesByIds) // 批量删除articles表
		articlesRouter.PUT("updateArticles", articlesApi.UpdateArticles)              // 更新articles表
	}
	{
		articlesRouterWithoutRecord.GET("findArticles", articlesApi.FindArticles)                // 根据ID获取articles表
		articlesRouterWithoutRecord.GET("getArticlesList", articlesApi.GetArticlesList)          // 获取articles表列表
		articlesRouterWithoutRecord.GET("getArticlesMind", articlesApi.GetArticlesByKnowledgeId) // 获取article脑图
	}
}
