package demo

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/demo"
    demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type HeroApi struct {}



// CreateHero 创建英雄信息
// @Tags Hero
// @Summary 创建英雄信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.Hero true "创建英雄信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /hero/createHero [post]
func (heroApi *HeroApi) CreateHero(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var hero demo.Hero
	err := c.ShouldBindJSON(&hero)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = heroService.CreateHero(ctx,&hero)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteHero 删除英雄信息
// @Tags Hero
// @Summary 删除英雄信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.Hero true "删除英雄信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /hero/deleteHero [delete]
func (heroApi *HeroApi) DeleteHero(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := heroService.DeleteHero(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteHeroByIds 批量删除英雄信息
// @Tags Hero
// @Summary 批量删除英雄信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /hero/deleteHeroByIds [delete]
func (heroApi *HeroApi) DeleteHeroByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := heroService.DeleteHeroByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateHero 更新英雄信息
// @Tags Hero
// @Summary 更新英雄信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.Hero true "更新英雄信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /hero/updateHero [put]
func (heroApi *HeroApi) UpdateHero(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var hero demo.Hero
	err := c.ShouldBindJSON(&hero)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = heroService.UpdateHero(ctx,hero)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindHero 用id查询英雄信息
// @Tags Hero
// @Summary 用id查询英雄信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询英雄信息"
// @Success 200 {object} response.Response{data=demo.Hero,msg=string} "查询成功"
// @Router /hero/findHero [get]
func (heroApi *HeroApi) FindHero(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	rehero, err := heroService.GetHero(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rehero, c)
}
// GetHeroList 分页获取英雄信息列表
// @Tags Hero
// @Summary 分页获取英雄信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.HeroSearch true "分页获取英雄信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /hero/getHeroList [get]
func (heroApi *HeroApi) GetHeroList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo demoReq.HeroSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := heroService.GetHeroInfoList(ctx,pageInfo)
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

// GetHeroPublic 不需要鉴权的英雄信息接口
// @Tags Hero
// @Summary 不需要鉴权的英雄信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /hero/getHeroPublic [get]
func (heroApi *HeroApi) GetHeroPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    heroService.GetHeroPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的英雄信息接口信息",
    }, "获取成功", c)
}
