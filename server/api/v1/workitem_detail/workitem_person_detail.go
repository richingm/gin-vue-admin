package workitem_detail

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/workitem_detail"
	workitem_detailReq "github.com/flipped-aurora/gin-vue-admin/server/model/workitem_detail/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WorkitemPersonDetailApi struct {
}

var workitemPersonDetailService = service.ServiceGroupApp.Workitem_detailServiceGroup.WorkitemPersonDetailService

// CreateWorkitemPersonDetail 创建workitemPersonDetail表
// @Tags WorkitemPersonDetail
// @Summary 创建workitemPersonDetail表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body workitem_detail.WorkitemPersonDetail true "创建workitemPersonDetail表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /workitemPersonDetail/createWorkitemPersonDetail [post]
func (workitemPersonDetailApi *WorkitemPersonDetailApi) CreateWorkitemPersonDetail(c *gin.Context) {
	var workitemPersonDetail workitem_detail.WorkitemPersonDetail
	err := c.ShouldBindJSON(&workitemPersonDetail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := workitemPersonDetailService.CreateWorkitemPersonDetail(&workitemPersonDetail); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWorkitemPersonDetail 删除workitemPersonDetail表
// @Tags WorkitemPersonDetail
// @Summary 删除workitemPersonDetail表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body workitem_detail.WorkitemPersonDetail true "删除workitemPersonDetail表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /workitemPersonDetail/deleteWorkitemPersonDetail [delete]
func (workitemPersonDetailApi *WorkitemPersonDetailApi) DeleteWorkitemPersonDetail(c *gin.Context) {
	ID := c.Query("ID")
	if err := workitemPersonDetailService.DeleteWorkitemPersonDetail(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWorkitemPersonDetailByIds 批量删除workitemPersonDetail表
// @Tags WorkitemPersonDetail
// @Summary 批量删除workitemPersonDetail表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /workitemPersonDetail/deleteWorkitemPersonDetailByIds [delete]
func (workitemPersonDetailApi *WorkitemPersonDetailApi) DeleteWorkitemPersonDetailByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := workitemPersonDetailService.DeleteWorkitemPersonDetailByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWorkitemPersonDetail 更新workitemPersonDetail表
// @Tags WorkitemPersonDetail
// @Summary 更新workitemPersonDetail表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body workitem_detail.WorkitemPersonDetail true "更新workitemPersonDetail表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /workitemPersonDetail/updateWorkitemPersonDetail [put]
func (workitemPersonDetailApi *WorkitemPersonDetailApi) UpdateWorkitemPersonDetail(c *gin.Context) {
	var workitemPersonDetail workitem_detail.WorkitemPersonDetail
	err := c.ShouldBindJSON(&workitemPersonDetail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := workitemPersonDetailService.UpdateWorkitemPersonDetail(workitemPersonDetail); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWorkitemPersonDetail 用id查询workitemPersonDetail表
// @Tags WorkitemPersonDetail
// @Summary 用id查询workitemPersonDetail表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query workitem_detail.WorkitemPersonDetail true "用id查询workitemPersonDetail表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /workitemPersonDetail/findWorkitemPersonDetail [get]
func (workitemPersonDetailApi *WorkitemPersonDetailApi) FindWorkitemPersonDetail(c *gin.Context) {
	ID := c.Query("ID")
	if reworkitemPersonDetail, err := workitemPersonDetailService.GetWorkitemPersonDetail(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reworkitemPersonDetail": reworkitemPersonDetail}, c)
	}
}

// GetWorkitemPersonDetailList 分页获取workitemPersonDetail表列表
// @Tags WorkitemPersonDetail
// @Summary 分页获取workitemPersonDetail表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query workitem_detailReq.WorkitemPersonDetailSearch true "分页获取workitemPersonDetail表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workitemPersonDetail/getWorkitemPersonDetailList [get]
func (workitemPersonDetailApi *WorkitemPersonDetailApi) GetWorkitemPersonDetailList(c *gin.Context) {
	var pageInfo workitem_detailReq.WorkitemPersonDetailSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := workitemPersonDetailService.GetWorkitemPersonDetailInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
