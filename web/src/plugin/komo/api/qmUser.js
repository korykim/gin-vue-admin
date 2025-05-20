import service from '@/utils/request'
// @Tags QmUser
// @Summary 创建用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.QmUser true "创建用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /qmUser/createQmUser [post]
export const createQmUser = (data) => {
  return service({
    url: '/qmUser/createQmUser',
    method: 'post',
    data
  })
}

// @Tags QmUser
// @Summary 删除用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.QmUser true "删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /qmUser/deleteQmUser [delete]
export const deleteQmUser = (params) => {
  return service({
    url: '/qmUser/deleteQmUser',
    method: 'delete',
    params
  })
}

// @Tags QmUser
// @Summary 批量删除用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /qmUser/deleteQmUser [delete]
export const deleteQmUserByIds = (params) => {
  return service({
    url: '/qmUser/deleteQmUserByIds',
    method: 'delete',
    params
  })
}

// @Tags QmUser
// @Summary 更新用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.QmUser true "更新用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /qmUser/updateQmUser [put]
export const updateQmUser = (data) => {
  return service({
    url: '/qmUser/updateQmUser',
    method: 'put',
    data
  })
}

// @Tags QmUser
// @Summary 用id查询用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.QmUser true "用id查询用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /qmUser/findQmUser [get]
export const findQmUser = (params) => {
  return service({
    url: '/qmUser/findQmUser',
    method: 'get',
    params
  })
}

// @Tags QmUser
// @Summary 分页获取用户信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /qmUser/getQmUserList [get]
export const getQmUserList = (params) => {
  return service({
    url: '/qmUser/getQmUserList',
    method: 'get',
    params
  })
}
// @Tags QmUser
// @Summary 不需要鉴权的用户信息接口
// @Accept application/json
// @Produce application/json
// @Param data query request.QmUserSearch true "分页获取用户信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /qmUser/getQmUserPublic [get]
export const getQmUserPublic = () => {
  return service({
    url: '/qmUser/getQmUserPublic',
    method: 'get',
  })
}

// AdminChangePassword AdminChangePassword
// @Tags QmUser
// @Summary AdminChangePassword
// @Accept application/json
// @Produce application/json
// @Param data body request.AdminChangePasswordReq true "AdminChangePassword"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /qmUser/adminChangePassword [PUT]
export const adminChangePassword = (data) => {
  return service({
    url: '/qmUser/adminChangePassword',
    method: 'PUT',
    data
  })
}
// 用户注册
export const register = (params) => {
  return service({
    url: "/qmUser/register",
    method: "POST",
    data: params
  })
}
// 用户登录前端逻辑
import request from "@/utils/request";

export const login = (username, password) => {
  return request({
    url: "/qmUser/login",
    method: "POST",
    data: {
      username,
      password
    }
  });
}
// 获取用户信息
export const GetUserInfo = (params) => {
  return service({
    url: "/qmUser/GetUserInfo",
    method: "GET",
    params: params
  })
}
