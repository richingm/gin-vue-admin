package workitem_detail

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WorkitemPersonDetailRouter struct {
}

// InitWorkitemPersonDetailRouter 初始化 workitemPersonDetail表 路由信息
func (s *WorkitemPersonDetailRouter) InitWorkitemPersonDetailRouter(Router *gin.RouterGroup) {
	workitemPersonDetailRouter := Router.Group("workitemPersonDetail").Use(middleware.OperationRecord())
	workitemPersonDetailRouterWithoutRecord := Router.Group("workitemPersonDetail")
	var workitemPersonDetailApi = v1.ApiGroupApp.Workitem_detailApiGroup.WorkitemPersonDetailApi
	{
		workitemPersonDetailRouter.POST("createWorkitemPersonDetail", workitemPersonDetailApi.CreateWorkitemPersonDetail)             // 新建workitemPersonDetail表
		workitemPersonDetailRouter.DELETE("deleteWorkitemPersonDetail", workitemPersonDetailApi.DeleteWorkitemPersonDetail)           // 删除workitemPersonDetail表
		workitemPersonDetailRouter.DELETE("deleteWorkitemPersonDetailByIds", workitemPersonDetailApi.DeleteWorkitemPersonDetailByIds) // 批量删除workitemPersonDetail表
		workitemPersonDetailRouter.PUT("updateWorkitemPersonDetail", workitemPersonDetailApi.UpdateWorkitemPersonDetail)              // 更新workitemPersonDetail表
	}
	{
		workitemPersonDetailRouterWithoutRecord.GET("findWorkitemPersonDetail", workitemPersonDetailApi.FindWorkitemPersonDetail)       // 根据ID获取workitemPersonDetail表
		workitemPersonDetailRouterWithoutRecord.GET("getWorkitemPersonDetailList", workitemPersonDetailApi.GetWorkitemPersonDetailList) // 获取workitemPersonDetail表列表
	}
}
