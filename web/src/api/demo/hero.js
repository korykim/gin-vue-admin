import service from '@/utils/request'
// @Tags Hero
// @Summary 创建英雄信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Hero true "创建英雄信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hero/createHero [post]
export const createHero = (data) => {
  return service({
    url: '/hero/createHero',
    method: 'post',
    data
  })
}

// @Tags Hero
// @Summary 删除英雄信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Hero true "删除英雄信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hero/deleteHero [delete]
export const deleteHero = (params) => {
  return service({
    url: '/hero/deleteHero',
    method: 'delete',
    params
  })
}

// @Tags Hero
// @Summary 批量删除英雄信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除英雄信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hero/deleteHero [delete]
export const deleteHeroByIds = (params) => {
  return service({
    url: '/hero/deleteHeroByIds',
    method: 'delete',
    params
  })
}

// @Tags Hero
// @Summary 更新英雄信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Hero true "更新英雄信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hero/updateHero [put]
export const updateHero = (data) => {
  return service({
    url: '/hero/updateHero',
    method: 'put',
    data
  })
}

// @Tags Hero
// @Summary 用id查询英雄信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Hero true "用id查询英雄信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hero/findHero [get]
export const findHero = (params) => {
  return service({
    url: '/hero/findHero',
    method: 'get',
    params
  })
}

// @Tags Hero
// @Summary 分页获取英雄信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取英雄信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hero/getHeroList [get]
export const getHeroList = (params) => {
  return service({
    url: '/hero/getHeroList',
    method: 'get',
    params
  })
}

// @Tags Hero
// @Summary 不需要鉴权的英雄信息接口
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.HeroSearch true "分页获取英雄信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /hero/getHeroPublic [get]
export const getHeroPublic = () => {
  return service({
    url: '/hero/getHeroPublic',
    method: 'get',
  })
}
