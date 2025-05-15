package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
	demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TagsApi struct{}

// CreateTags 创建tags表
// @Tags Tags
// @Summary 创建tags表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.Tag true "创建tags表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /tags/createTags [post]
func (tagsApi *TagsApi) CreateTags(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var tags demo.Tag
	err := c.ShouldBindJSON(&tags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = tagsService.CreateTags(ctx, &tags)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteTags 删除tags表
// @Tags Tags
// @Summary 删除tags表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.Tag true "删除tags表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /tags/deleteTags [delete]
func (tagsApi *TagsApi) DeleteTags(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := tagsService.DeleteTags(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTagsByIds 批量删除tags表
// @Tags Tags
// @Summary 批量删除tags表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /tags/deleteTagsByIds [delete]
func (tagsApi *TagsApi) DeleteTagsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := tagsService.DeleteTagsByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTags 更新tags表
// @Tags Tags
// @Summary 更新tags表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.Tag true "更新tags表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /tags/updateTags [put]
func (tagsApi *TagsApi) UpdateTags(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var tags demo.Tag
	err := c.ShouldBindJSON(&tags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = tagsService.UpdateTags(ctx, tags)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTags 用id查询tags表
// @Tags Tags
// @Summary 用id查询tags表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询tags表"
// @Success 200 {object} response.Response{data=demo.Tag,msg=string} "查询成功"
// @Router /tags/findTags [get]
func (tagsApi *TagsApi) FindTags(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	retags, err := tagsService.GetTags(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(retags, c)
}

// GetTagsList 分页获取tags表列表
// @Tags Tags
// @Summary 分页获取tags表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.TagsSearch true "分页获取tags表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /tags/getTagsList [get]
func (tagsApi *TagsApi) GetTagsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo demoReq.TagsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := tagsService.GetTagsInfoList(ctx, pageInfo)
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

// GetTagsPublic 不需要鉴权的tags表接口
// @Tags Tags
// @Summary 不需要鉴权的tags表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tags/getTagsPublic [get]
func (tagsApi *TagsApi) GetTagsPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	tagsService.GetTagsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的tags表接口信息",
	}, "获取成功", c)
}
