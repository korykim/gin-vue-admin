package demo

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/demo"
    demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type HeroSkillApi struct {}



// CreateHeroSkill 创建技能
// @Tags HeroSkill
// @Summary 创建技能
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.HeroSkill true "创建技能"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /heroskill/createHeroSkill [post]
func (heroskillApi *HeroSkillApi) CreateHeroSkill(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var heroskill demo.HeroSkill
	err := c.ShouldBindJSON(&heroskill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = heroskillService.CreateHeroSkill(ctx,&heroskill)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteHeroSkill 删除技能
// @Tags HeroSkill
// @Summary 删除技能
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.HeroSkill true "删除技能"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /heroskill/deleteHeroSkill [delete]
func (heroskillApi *HeroSkillApi) DeleteHeroSkill(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := heroskillService.DeleteHeroSkill(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteHeroSkillByIds 批量删除技能
// @Tags HeroSkill
// @Summary 批量删除技能
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /heroskill/deleteHeroSkillByIds [delete]
func (heroskillApi *HeroSkillApi) DeleteHeroSkillByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := heroskillService.DeleteHeroSkillByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateHeroSkill 更新技能
// @Tags HeroSkill
// @Summary 更新技能
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.HeroSkill true "更新技能"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /heroskill/updateHeroSkill [put]
func (heroskillApi *HeroSkillApi) UpdateHeroSkill(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var heroskill demo.HeroSkill
	err := c.ShouldBindJSON(&heroskill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = heroskillService.UpdateHeroSkill(ctx,heroskill)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindHeroSkill 用id查询技能
// @Tags HeroSkill
// @Summary 用id查询技能
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询技能"
// @Success 200 {object} response.Response{data=demo.HeroSkill,msg=string} "查询成功"
// @Router /heroskill/findHeroSkill [get]
func (heroskillApi *HeroSkillApi) FindHeroSkill(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reheroskill, err := heroskillService.GetHeroSkill(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reheroskill, c)
}
// GetHeroSkillList 分页获取技能列表
// @Tags HeroSkill
// @Summary 分页获取技能列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.HeroSkillSearch true "分页获取技能列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /heroskill/getHeroSkillList [get]
func (heroskillApi *HeroSkillApi) GetHeroSkillList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo demoReq.HeroSkillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := heroskillService.GetHeroSkillInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetHeroSkillPublic 不需要鉴权的技能接口
// @Tags HeroSkill
// @Summary 不需要鉴权的技能接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /heroskill/getHeroSkillPublic [get]
func (heroskillApi *HeroSkillApi) GetHeroSkillPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    heroskillService.GetHeroSkillPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的技能接口信息",
    }, "获取成功", c)
}
