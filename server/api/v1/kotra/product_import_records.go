package kotra

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/kotra"
    kotraReq "github.com/flipped-aurora/gin-vue-admin/server/model/kotra/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ProductImportRecordsApi struct {}



// CreateProductImportRecords 创建未准入境产品表
// @Tags ProductImportRecords
// @Summary 创建未准入境产品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body kotra.ProductImportRecords true "创建未准入境产品表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /productImportRecords/createProductImportRecords [post]
func (productImportRecordsApi *ProductImportRecordsApi) CreateProductImportRecords(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var productImportRecords kotra.ProductImportRecords
	err := c.ShouldBindJSON(&productImportRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productImportRecordsService.CreateProductImportRecords(ctx,&productImportRecords)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteProductImportRecords 删除未准入境产品表
// @Tags ProductImportRecords
// @Summary 删除未准入境产品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body kotra.ProductImportRecords true "删除未准入境产品表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /productImportRecords/deleteProductImportRecords [delete]
func (productImportRecordsApi *ProductImportRecordsApi) DeleteProductImportRecords(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := productImportRecordsService.DeleteProductImportRecords(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteProductImportRecordsByIds 批量删除未准入境产品表
// @Tags ProductImportRecords
// @Summary 批量删除未准入境产品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /productImportRecords/deleteProductImportRecordsByIds [delete]
func (productImportRecordsApi *ProductImportRecordsApi) DeleteProductImportRecordsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := productImportRecordsService.DeleteProductImportRecordsByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateProductImportRecords 更新未准入境产品表
// @Tags ProductImportRecords
// @Summary 更新未准入境产品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body kotra.ProductImportRecords true "更新未准入境产品表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /productImportRecords/updateProductImportRecords [put]
func (productImportRecordsApi *ProductImportRecordsApi) UpdateProductImportRecords(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var productImportRecords kotra.ProductImportRecords
	err := c.ShouldBindJSON(&productImportRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productImportRecordsService.UpdateProductImportRecords(ctx,productImportRecords)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindProductImportRecords 用id查询未准入境产品表
// @Tags ProductImportRecords
// @Summary 用id查询未准入境产品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询未准入境产品表"
// @Success 200 {object} response.Response{data=kotra.ProductImportRecords,msg=string} "查询成功"
// @Router /productImportRecords/findProductImportRecords [get]
func (productImportRecordsApi *ProductImportRecordsApi) FindProductImportRecords(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reproductImportRecords, err := productImportRecordsService.GetProductImportRecords(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reproductImportRecords, c)
}
// GetProductImportRecordsList 分页获取未准入境产品表列表
// @Tags ProductImportRecords
// @Summary 分页获取未准入境产品表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query kotraReq.ProductImportRecordsSearch true "分页获取未准入境产品表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /productImportRecords/getProductImportRecordsList [get]
func (productImportRecordsApi *ProductImportRecordsApi) GetProductImportRecordsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo kotraReq.ProductImportRecordsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := productImportRecordsService.GetProductImportRecordsInfoList(ctx,pageInfo)
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

// GetProductImportRecordsPublic 不需要鉴权的未准入境产品表接口
// @Tags ProductImportRecords
// @Summary 不需要鉴权的未准入境产品表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /productImportRecords/getProductImportRecordsPublic [get]
func (productImportRecordsApi *ProductImportRecordsApi) GetProductImportRecordsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    productImportRecordsService.GetProductImportRecordsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的未准入境产品表接口信息",
    }, "获取成功", c)
}
