package knowledge

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/knowledge"
	knowledgeReq "github.com/flipped-aurora/gin-vue-admin/server/model/knowledge/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type KnowledgesApi struct {
}

var knowledgesService = service.ServiceGroupApp.KnowledgeServiceGroup.KnowledgesService

// CreateKnowledges 创建knowledges表
// @Tags Knowledges
// @Summary 创建knowledges表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body knowledge.Knowledges true "创建knowledges表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /knowledges/createKnowledges [post]
func (knowledgesApi *KnowledgesApi) CreateKnowledges(c *gin.Context) {
	var knowledges knowledge.Knowledges
	err := c.ShouldBindJSON(&knowledges)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	knowledges.CreatedBy = utils.GetUserID(c)

	if err := knowledgesService.CreateKnowledges(&knowledges); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteKnowledges 删除knowledges表
// @Tags Knowledges
// @Summary 删除knowledges表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body knowledge.Knowledges true "删除knowledges表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /knowledges/deleteKnowledges [delete]
func (knowledgesApi *KnowledgesApi) DeleteKnowledges(c *gin.Context) {
	id := c.Query("id")
	userID := utils.GetUserID(c)
	if err := knowledgesService.DeleteKnowledges(id, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteKnowledgesByIds 批量删除knowledges表
// @Tags Knowledges
// @Summary 批量删除knowledges表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /knowledges/deleteKnowledgesByIds [delete]
func (knowledgesApi *KnowledgesApi) DeleteKnowledgesByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	userID := utils.GetUserID(c)
	if err := knowledgesService.DeleteKnowledgesByIds(ids, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateKnowledges 更新knowledges表
// @Tags Knowledges
// @Summary 更新knowledges表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body knowledge.Knowledges true "更新knowledges表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /knowledges/updateKnowledges [put]
func (knowledgesApi *KnowledgesApi) UpdateKnowledges(c *gin.Context) {
	var knowledges knowledge.Knowledges
	err := c.ShouldBindJSON(&knowledges)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	knowledges.UpdatedBy = utils.GetUserID(c)

	if err := knowledgesService.UpdateKnowledges(knowledges); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindKnowledges 用id查询knowledges表
// @Tags Knowledges
// @Summary 用id查询knowledges表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query knowledge.Knowledges true "用id查询knowledges表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /knowledges/findKnowledges [get]
func (knowledgesApi *KnowledgesApi) FindKnowledges(c *gin.Context) {
	id := c.Query("id")
	if reknowledges, err := knowledgesService.GetKnowledges(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reknowledges": reknowledges}, c)
	}
}

// GetKnowledgesList 分页获取knowledges表列表
// @Tags Knowledges
// @Summary 分页获取knowledges表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query knowledgeReq.KnowledgesSearch true "分页获取knowledges表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /knowledges/getKnowledgesList [get]
func (knowledgesApi *KnowledgesApi) GetKnowledgesList(c *gin.Context) {
	var pageInfo knowledgeReq.KnowledgesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := knowledgesService.GetKnowledgesInfoList(pageInfo); err != nil {
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
