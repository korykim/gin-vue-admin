package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
	demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PoiItemsApi struct{}

// CreatePoiItems 创建poiItems表
// @Tags PoiItems
// @Summary 创建poiItems表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.PoiItems true "创建poiItems表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /poiItems/createPoiItems [post]
func (poiItemsApi *PoiItemsApi) CreatePoiItems(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var poiItems demo.PoiItems
	err := c.ShouldBindJSON(&poiItems)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = poiItemsService.CreatePoiItems(ctx, &poiItems)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// CreatePoiItemsInBatches 批量创建poiItems表
// @Tags PoiItems
// @Summary 批量创建poiItems表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.PoiItems true "批量创建poiItems表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /poiItems/createPoiItemsInBatches [post]
func (poiItemsApi *PoiItemsApi) CreatePoiItemsInBatches(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var poiItems []*demo.PoiItems
	err := c.ShouldBindJSON(&poiItems)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = poiItemsService.CreatePoiItemsInBatches(ctx, poiItems)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePoiItems 删除poiItems表
// @Tags PoiItems
// @Summary 删除poiItems表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.PoiItems true "删除poiItems表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /poiItems/deletePoiItems [delete]
func (poiItemsApi *PoiItemsApi) DeletePoiItems(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := poiItemsService.DeletePoiItems(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePoiItemsByIds 批量删除poiItems表
// @Tags PoiItems
// @Summary 批量删除poiItems表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /poiItems/deletePoiItemsByIds [delete]
func (poiItemsApi *PoiItemsApi) DeletePoiItemsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := poiItemsService.DeletePoiItemsByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePoiItems 更新poiItems表
// @Tags PoiItems
// @Summary 更新poiItems表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.PoiItems true "更新poiItems表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /poiItems/updatePoiItems [put]
func (poiItemsApi *PoiItemsApi) UpdatePoiItems(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var poiItems demo.PoiItems
	err := c.ShouldBindJSON(&poiItems)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = poiItemsService.UpdatePoiItems(ctx, poiItems)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPoiItems 用id查询poiItems表
// @Tags PoiItems
// @Summary 用id查询poiItems表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query string true "用id查询poiItems表"
// @Success 200 {object} response.Response{data=demo.PoiItems,msg=string} "查询成功"
// @Router /poiItems/findPoiItems [get]
func (poiItemsApi *PoiItemsApi) FindPoiItems(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	repoiItems, err := poiItemsService.GetPoiItems(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(repoiItems, c)
}

// GetPoiItemsList 分页获取poiItems表列表
// @Tags PoiItems
// @Summary 分页获取poiItems表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.PoiItemsSearch true "分页获取poiItems表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /poiItems/getPoiItemsList [get]
func (poiItemsApi *PoiItemsApi) GetPoiItemsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo demoReq.PoiItemsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := poiItemsService.GetPoiItemsInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetPoiItemsPublic 不需要鉴权的poiItems表接口
// @Tags PoiItems
// @Summary 不需要鉴权的poiItems表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /poiItems/getPoiItemsPublic [get]
func (poiItemsApi *PoiItemsApi) GetPoiItemsPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	poiItemsService.GetPoiItemsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的poiItems表接口信息",
	}, "获取成功", c)
}
