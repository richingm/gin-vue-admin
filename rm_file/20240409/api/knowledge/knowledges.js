import service from '@/utils/request'

// @Tags Knowledges
// @Summary 创建knowledges表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Knowledges true "创建knowledges表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /knowledges/createKnowledges [post]
export const createKnowledges = (data) => {
  return service({
    url: '/knowledges/createKnowledges',
    method: 'post',
    data
  })
}

// @Tags Knowledges
// @Summary 删除knowledges表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Knowledges true "删除knowledges表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /knowledges/deleteKnowledges [delete]
export const deleteKnowledges = (params) => {
  return service({
    url: '/knowledges/deleteKnowledges',
    method: 'delete',
    params
  })
}

// @Tags Knowledges
// @Summary 批量删除knowledges表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除knowledges表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /knowledges/deleteKnowledges [delete]
export const deleteKnowledgesByIds = (params) => {
  return service({
    url: '/knowledges/deleteKnowledgesByIds',
    method: 'delete',
    params
  })
}

// @Tags Knowledges
// @Summary 更新knowledges表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Knowledges true "更新knowledges表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /knowledges/updateKnowledges [put]
export const updateKnowledges = (data) => {
  return service({
    url: '/knowledges/updateKnowledges',
    method: 'put',
    data
  })
}

// @Tags Knowledges
// @Summary 用id查询knowledges表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Knowledges true "用id查询knowledges表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /knowledges/findKnowledges [get]
export const findKnowledges = (params) => {
  return service({
    url: '/knowledges/findKnowledges',
    method: 'get',
    params
  })
}

// @Tags Knowledges
// @Summary 分页获取knowledges表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取knowledges表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /knowledges/getKnowledgesList [get]
export const getKnowledgesList = (params) => {
  return service({
    url: '/knowledges/getKnowledgesList',
    method: 'get',
    params
  })
}
