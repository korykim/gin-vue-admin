package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PosterTagApi struct{}

// AddTagsToPoster 为海报添加标签
// @Tags PosterTag
// @Summary 为海报添加标签
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body struct{PosterID uint `json:"posterId"`;TagIDs []uint `json:"tagIds"`} true "海报ID和标签ID列表"
// @Success 200 {object} response.Response{msg=string} "添加成功"
// @Router /posterTag/addTagsToPoster [post]
func (posterTagApi *PosterTagApi) AddTagsToPoster(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var reqData struct {
		PosterID uint   `json:"posterId"`
		TagIDs   []uint `json:"tagIds"`
	}

	err := c.ShouldBindJSON(&reqData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = posterTagService.AddTagsToPoster(ctx, reqData.PosterID, reqData.TagIDs)
	if err != nil {
		global.GVA_LOG.Error("添加标签失败!", zap.Error(err))
		response.FailWithMessage("添加标签失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("添加标签成功", c)
}

// RemoveTagsFromPoster 从海报中移除标签
// @Tags PosterTag
// @Summary 从海报中移除标签
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body struct{PosterID uint `json:"posterId"`;TagIDs []uint `json:"tagIds"`} true "海报ID和标签ID列表"
// @Success 200 {object} response.Response{msg=string} "移除成功"
// @Router /posterTag/removeTagsFromPoster [delete]
func (posterTagApi *PosterTagApi) RemoveTagsFromPoster(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var reqData struct {
		PosterID uint   `json:"posterId"`
		TagIDs   []uint `json:"tagIds"`
	}

	err := c.ShouldBindJSON(&reqData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = posterTagService.RemoveTagsFromPoster(ctx, reqData.PosterID, reqData.TagIDs)
	if err != nil {
		global.GVA_LOG.Error("移除标签失败!", zap.Error(err))
		response.FailWithMessage("移除标签失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("移除标签成功", c)
}

// GetPosterTags 获取海报的所有标签
// @Tags PosterTag
// @Summary 获取海报的所有标签
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param posterId query uint true "海报ID"
// @Success 200 {object} response.Response{data=[]demo.Tag,msg=string} "获取成功"
// @Router /posterTag/getPosterTags [get]
func (posterTagApi *PosterTagApi) GetPosterTags(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var reqData struct {
		PosterID uint `form:"posterId" binding:"required"`
	}

	if err := c.ShouldBindQuery(&reqData); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	tags, err := posterTagService.GetPosterTags(ctx, reqData.PosterID)
	if err != nil {
		global.GVA_LOG.Error("获取标签失败!", zap.Error(err))
		response.FailWithMessage("获取标签失败:"+err.Error(), c)
		return
	}
	response.OkWithData(tags, c)
}

// GetTagPosters 获取标签关联的所有海报
// @Tags PosterTag
// @Summary 获取标签关联的所有海报
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param tagId query uint true "标签ID"
// @Success 200 {object} response.Response{data=[]demo.Poster,msg=string} "获取成功"
// @Router /posterTag/getTagPosters [get]
func (posterTagApi *PosterTagApi) GetTagPosters(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var reqData struct {
		TagID uint `form:"tagId" binding:"required"`
	}

	if err := c.ShouldBindQuery(&reqData); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	posters, err := posterTagService.GetTagPosters(ctx, reqData.TagID)
	if err != nil {
		global.GVA_LOG.Error("获取海报失败!", zap.Error(err))
		response.FailWithMessage("获取海报失败:"+err.Error(), c)
		return
	}
	response.OkWithData(posters, c)
}

// GetPosterTagRelations 分页获取海报标签关联列表
// @Tags PosterTag
// @Summary 分页获取海报标签关联列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.PosterTagSearch true "分页获取海报标签关联列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /posterTag/getPosterTagRelations [get]
func (posterTagApi *PosterTagApi) GetPosterTagRelations(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo demoReq.PosterTagSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := posterTagService.GetPosterTagRelations(ctx, pageInfo)
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

// GetTagsUsageStats 获取标签使用统计
// @Tags PosterTag
// @Summary 获取标签使用统计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]map[string]interface{},msg=string} "获取成功"
// @Router /posterTag/getTagsUsageStats [get]
func (posterTagApi *PosterTagApi) GetTagsUsageStats(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	stats, err := posterTagService.GetTagsUsageStats(ctx)
	if err != nil {
		global.GVA_LOG.Error("获取标签使用统计失败!", zap.Error(err))
		response.FailWithMessage("获取标签使用统计失败:"+err.Error(), c)
		return
	}
	response.OkWithData(stats, c)
}
