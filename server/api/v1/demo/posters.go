package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
	demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PostersApi struct{}

// CreatePosters 创建posters表
// @Tags Posters
// @Summary 创建posters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.Poster true "创建posters表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /posters/createPosters [post]
func (postersApi *PostersApi) CreatePosters(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var posters demo.Poster
	err := c.ShouldBindJSON(&posters)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = postersService.CreatePosters(ctx, &posters)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	// 返回创建的海报数据，包括ID
	response.OkWithData(posters, c)
}

// DeletePosters 删除posters表
// @Tags Posters
// @Summary 删除posters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.Poster true "删除posters表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /posters/deletePosters [delete]
func (postersApi *PostersApi) DeletePosters(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := postersService.DeletePosters(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePostersByIds 批量删除posters表
// @Tags Posters
// @Summary 批量删除posters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /posters/deletePostersByIds [delete]
func (postersApi *PostersApi) DeletePostersByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := postersService.DeletePostersByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePosters 更新posters表
// @Tags Posters
// @Summary 更新posters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.Poster true "更新posters表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /posters/updatePosters [put]
func (postersApi *PostersApi) UpdatePosters(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var posters demo.Poster
	err := c.ShouldBindJSON(&posters)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = postersService.UpdatePosters(ctx, posters)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPosters 用id查询posters表
// @Tags Posters
// @Summary 用id查询posters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询posters表"
// @Success 200 {object} response.Response{data=demo.Poster,msg=string} "查询成功"
// @Router /posters/findPosters [get]
func (postersApi *PostersApi) FindPosters(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reposters, err := postersService.GetPosters(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reposters, c)
}

// GetPostersList 分页获取posters表列表
// @Tags Posters
// @Summary 分页获取posters表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.PostersSearch true "分页获取posters表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /posters/getPostersList [get]
func (postersApi *PostersApi) GetPostersList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo demoReq.PostersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := postersService.GetPostersInfoList(ctx, pageInfo)
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

// GetPostersPublic 不需要鉴权的posters表接口
// @Tags Posters
// @Summary 不需要鉴权的posters表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /posters/getPostersPublic [get]
func (postersApi *PostersApi) GetPostersPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	postersService.GetPostersPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的posters表接口信息",
	}, "获取成功", c)
}
