import service from '@/utils/request'
// @Tags EmployeeProfile
// @Summary 创建员工信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EmployeeProfile true "创建员工信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /TSE/createEmployeeProfile [post]
export const createEmployeeProfile = (data) => {
  return service({
    url: '/TSE/createEmployeeProfile',
    method: 'post',
    data
  })
}

// @Tags EmployeeProfile
// @Summary 删除员工信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EmployeeProfile true "删除员工信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TSE/deleteEmployeeProfile [delete]
export const deleteEmployeeProfile = (params) => {
  return service({
    url: '/TSE/deleteEmployeeProfile',
    method: 'delete',
    params
  })
}

// @Tags EmployeeProfile
// @Summary 批量删除员工信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除员工信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TSE/deleteEmployeeProfile [delete]
export const deleteEmployeeProfileByIds = (params) => {
  return service({
    url: '/TSE/deleteEmployeeProfileByIds',
    method: 'delete',
    params
  })
}

// @Tags EmployeeProfile
// @Summary 更新员工信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EmployeeProfile true "更新员工信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TSE/updateEmployeeProfile [put]
export const updateEmployeeProfile = (data) => {
  return service({
    url: '/TSE/updateEmployeeProfile',
    method: 'put',
    data
  })
}

// @Tags EmployeeProfile
// @Summary 用id查询员工信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.EmployeeProfile true "用id查询员工信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TSE/findEmployeeProfile [get]
export const findEmployeeProfile = (params) => {
  return service({
    url: '/TSE/findEmployeeProfile',
    method: 'get',
    params
  })
}

// @Tags EmployeeProfile
// @Summary 分页获取员工信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取员工信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TSE/getEmployeeProfileList [get]
export const getEmployeeProfileList = (params) => {
  return service({
    url: '/TSE/getEmployeeProfileList',
    method: 'get',
    params
  })
}

// @Tags EmployeeProfile
// @Summary 不需要鉴权的员工信息接口
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.EmployeeProfileSearch true "分页获取员工信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /TSE/getEmployeeProfilePublic [get]
export const getEmployeeProfilePublic = () => {
  return service({
    url: '/TSE/getEmployeeProfilePublic',
    method: 'get',
  })
}
