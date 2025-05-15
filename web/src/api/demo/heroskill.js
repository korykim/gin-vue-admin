import service from '@/utils/request'
// @Tags HeroSkill
// @Summary 创建技能
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.HeroSkill true "创建技能"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /heroskill/createHeroSkill [post]
export const createHeroSkill = (data) => {
  return service({
    url: '/heroskill/createHeroSkill',
    method: 'post',
    data
  })
}

// @Tags HeroSkill
// @Summary 删除技能
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.HeroSkill true "删除技能"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /heroskill/deleteHeroSkill [delete]
export const deleteHeroSkill = (params) => {
  return service({
    url: '/heroskill/deleteHeroSkill',
    method: 'delete',
    params
  })
}

// @Tags HeroSkill
// @Summary 批量删除技能
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除技能"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /heroskill/deleteHeroSkill [delete]
export const deleteHeroSkillByIds = (params) => {
  return service({
    url: '/heroskill/deleteHeroSkillByIds',
    method: 'delete',
    params
  })
}

// @Tags HeroSkill
// @Summary 更新技能
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.HeroSkill true "更新技能"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /heroskill/updateHeroSkill [put]
export const updateHeroSkill = (data) => {
  return service({
    url: '/heroskill/updateHeroSkill',
    method: 'put',
    data
  })
}

// @Tags HeroSkill
// @Summary 用id查询技能
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.HeroSkill true "用id查询技能"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /heroskill/findHeroSkill [get]
export const findHeroSkill = (params) => {
  return service({
    url: '/heroskill/findHeroSkill',
    method: 'get',
    params
  })
}

// @Tags HeroSkill
// @Summary 分页获取技能列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取技能列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /heroskill/getHeroSkillList [get]
export const getHeroSkillList = (params) => {
  return service({
    url: '/heroskill/getHeroSkillList',
    method: 'get',
    params
  })
}

// @Tags HeroSkill
// @Summary 不需要鉴权的技能接口
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.HeroSkillSearch true "分页获取技能列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /heroskill/getHeroSkillPublic [get]
export const getHeroSkillPublic = () => {
  return service({
    url: '/heroskill/getHeroSkillPublic',
    method: 'get',
  })
}
