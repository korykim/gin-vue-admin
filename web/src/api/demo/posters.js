import service from '@/utils/request'
// @Tags Posters
// @Summary 创建posters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Posters true "创建posters表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /posters/createPosters [post]
export const createPosters = (data) => {
  return service({
    url: '/posters/createPosters',
    method: 'post',
    data
  })
}

// @Tags Posters
// @Summary 删除posters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Posters true "删除posters表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /posters/deletePosters [delete]
export const deletePosters = (params) => {
  return service({
    url: '/posters/deletePosters',
    method: 'delete',
    params
  })
}

// @Tags Posters
// @Summary 批量删除posters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除posters表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /posters/deletePosters [delete]
export const deletePostersByIds = (params) => {
  return service({
    url: '/posters/deletePostersByIds',
    method: 'delete',
    params
  })
}

// @Tags Posters
// @Summary 更新posters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Posters true "更新posters表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /posters/updatePosters [put]
export const updatePosters = (data) => {
  return service({
    url: '/posters/updatePosters',
    method: 'put',
    data
  })
}

// @Tags Posters
// @Summary 用id查询posters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Posters true "用id查询posters表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /posters/findPosters [get]
export const findPosters = (params) => {
  return service({
    url: '/posters/findPosters',
    method: 'get',
    params
  })
}

// @Tags Posters
// @Summary 分页获取posters表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取posters表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /posters/getPostersList [get]
export const getPostersList = (params) => {
  return service({
    url: '/posters/getPostersList',
    method: 'get',
    params
  })
}

// @Tags Posters
// @Summary 不需要鉴权的posters表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /posters/getPostersPublic [get]
export const getPostersPublic = () => {
  return service({
    url: '/posters/getPostersPublic',
    method: 'get'
  })
}
