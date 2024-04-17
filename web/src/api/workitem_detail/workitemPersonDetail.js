import service from '@/utils/request'

// @Tags WorkitemPersonDetail
// @Summary 创建workitemPersonDetail表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkitemPersonDetail true "创建workitemPersonDetail表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /workitemPersonDetail/createWorkitemPersonDetail [post]
export const createWorkitemPersonDetail = (data) => {
  return service({
    url: '/workitemPersonDetail/createWorkitemPersonDetail',
    method: 'post',
    data
  })
}

// @Tags WorkitemPersonDetail
// @Summary 删除workitemPersonDetail表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkitemPersonDetail true "删除workitemPersonDetail表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /workitemPersonDetail/deleteWorkitemPersonDetail [delete]
export const deleteWorkitemPersonDetail = (params) => {
  return service({
    url: '/workitemPersonDetail/deleteWorkitemPersonDetail',
    method: 'delete',
    params
  })
}

// @Tags WorkitemPersonDetail
// @Summary 批量删除workitemPersonDetail表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除workitemPersonDetail表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /workitemPersonDetail/deleteWorkitemPersonDetail [delete]
export const deleteWorkitemPersonDetailByIds = (params) => {
  return service({
    url: '/workitemPersonDetail/deleteWorkitemPersonDetailByIds',
    method: 'delete',
    params
  })
}

// @Tags WorkitemPersonDetail
// @Summary 更新workitemPersonDetail表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkitemPersonDetail true "更新workitemPersonDetail表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /workitemPersonDetail/updateWorkitemPersonDetail [put]
export const updateWorkitemPersonDetail = (data) => {
  return service({
    url: '/workitemPersonDetail/updateWorkitemPersonDetail',
    method: 'put',
    data
  })
}

// @Tags WorkitemPersonDetail
// @Summary 用id查询workitemPersonDetail表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WorkitemPersonDetail true "用id查询workitemPersonDetail表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /workitemPersonDetail/findWorkitemPersonDetail [get]
export const findWorkitemPersonDetail = (params) => {
  return service({
    url: '/workitemPersonDetail/findWorkitemPersonDetail',
    method: 'get',
    params
  })
}

// @Tags WorkitemPersonDetail
// @Summary 分页获取workitemPersonDetail表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取workitemPersonDetail表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workitemPersonDetail/getWorkitemPersonDetailList [get]
export const getWorkitemPersonDetailList = (params) => {
  return service({
    url: '/workitemPersonDetail/getWorkitemPersonDetailList',
    method: 'get',
    params
  })
}
