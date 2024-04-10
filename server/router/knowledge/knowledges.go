package knowledge

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type KnowledgesRouter struct {
}

// InitKnowledgesRouter 初始化 knowledges表 路由信息
func (s *KnowledgesRouter) InitKnowledgesRouter(Router *gin.RouterGroup) {
	knowledgesRouter := Router.Group("knowledges").Use(middleware.OperationRecord())
	knowledgesRouterWithoutRecord := Router.Group("knowledges")
	var knowledgesApi = v1.ApiGroupApp.KnowledgeApiGroup.KnowledgesApi
	{
		knowledgesRouter.POST("createKnowledges", knowledgesApi.CreateKnowledges)             // 新建knowledges表
		knowledgesRouter.DELETE("deleteKnowledges", knowledgesApi.DeleteKnowledges)           // 删除knowledges表
		knowledgesRouter.DELETE("deleteKnowledgesByIds", knowledgesApi.DeleteKnowledgesByIds) // 批量删除knowledges表
		knowledgesRouter.PUT("updateKnowledges", knowledgesApi.UpdateKnowledges)              // 更新knowledges表
	}
	{
		knowledgesRouterWithoutRecord.GET("findKnowledges", knowledgesApi.FindKnowledges)             // 根据ID获取knowledges表
		knowledgesRouterWithoutRecord.GET("getKnowledgesList", knowledgesApi.GetKnowledgesList)       // 获取knowledges表列表
		knowledgesRouterWithoutRecord.GET("getKnowledgesOptions", knowledgesApi.GetKnowledgesOptions) // 获取knowledgeoptions
	}
}
