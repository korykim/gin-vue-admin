import service from '@/utils/request'
// @Tags Tags
// @Summary 创建tags表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Tags true "创建tags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tags/createTags [post]
export const createTags = (data) => {
  return service({
    url: '/tags/createTags',
    method: 'post',
    data
  })
}

// @Tags Tags
// @Summary 删除tags表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Tags true "删除tags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tags/deleteTags [delete]
export const deleteTags = (params) => {
  return service({
    url: '/tags/deleteTags',
    method: 'delete',
    params
  })
}

// @Tags Tags
// @Summary 批量删除tags表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除tags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tags/deleteTags [delete]
export const deleteTagsByIds = (params) => {
  return service({
    url: '/tags/deleteTagsByIds',
    method: 'delete',
    params
  })
}

// @Tags Tags
// @Summary 更新tags表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Tags true "更新tags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tags/updateTags [put]
export const updateTags = (data) => {
  return service({
    url: '/tags/updateTags',
    method: 'put',
    data
  })
}

// @Tags Tags
// @Summary 用id查询tags表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Tags true "用id查询tags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tags/findTags [get]
export const findTags = (params) => {
  return service({
    url: '/tags/findTags',
    method: 'get',
    params
  })
}

// @Tags Tags
// @Summary 分页获取tags表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取tags表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tags/getTagsList [get]
export const getTagsList = (params) => {
  return service({
    url: '/tags/getTagsList',
    method: 'get',
    params
  })
}

// @Tags Tags
// @Summary 不需要鉴权的tags表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tags/getTagsPublic [get]
export const getTagsPublic = () => {
  return service({
    url: '/tags/getTagsPublic',
    method: 'get'
  })
}
