package demo

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/demo"
    demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ToyStoreApi struct {}



// CreateToyStore 创建玩具店
// @Tags ToyStore
// @Summary 创建玩具店
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.ToyStore true "创建玩具店"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /TS/createToyStore [post]
func (TSApi *ToyStoreApi) CreateToyStore(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var TS demo.ToyStore
	err := c.ShouldBindJSON(&TS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = TSService.CreateToyStore(ctx,&TS)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteToyStore 删除玩具店
// @Tags ToyStore
// @Summary 删除玩具店
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.ToyStore true "删除玩具店"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /TS/deleteToyStore [delete]
func (TSApi *ToyStoreApi) DeleteToyStore(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := TSService.DeleteToyStore(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteToyStoreByIds 批量删除玩具店
// @Tags ToyStore
// @Summary 批量删除玩具店
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /TS/deleteToyStoreByIds [delete]
func (TSApi *ToyStoreApi) DeleteToyStoreByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := TSService.DeleteToyStoreByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateToyStore 更新玩具店
// @Tags ToyStore
// @Summary 更新玩具店
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.ToyStore true "更新玩具店"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /TS/updateToyStore [put]
func (TSApi *ToyStoreApi) UpdateToyStore(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var TS demo.ToyStore
	err := c.ShouldBindJSON(&TS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = TSService.UpdateToyStore(ctx,TS)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindToyStore 用id查询玩具店
// @Tags ToyStore
// @Summary 用id查询玩具店
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询玩具店"
// @Success 200 {object} response.Response{data=demo.ToyStore,msg=string} "查询成功"
// @Router /TS/findToyStore [get]
func (TSApi *ToyStoreApi) FindToyStore(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reTS, err := TSService.GetToyStore(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reTS, c)
}
// GetToyStoreList 分页获取玩具店列表
// @Tags ToyStore
// @Summary 分页获取玩具店列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.ToyStoreSearch true "分页获取玩具店列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /TS/getToyStoreList [get]
func (TSApi *ToyStoreApi) GetToyStoreList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo demoReq.ToyStoreSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := TSService.GetToyStoreInfoList(ctx,pageInfo)
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
// GetToyStoreDataSource 获取ToyStore的数据源
// @Tags ToyStore
// @Summary 获取ToyStore的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /TS/getToyStoreDataSource [get]
func (TSApi *ToyStoreApi) GetToyStoreDataSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口为获取数据源定义的数据
    dataSource, err := TSService.GetToyStoreDataSource(ctx)
    if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
   		response.FailWithMessage("查询失败:" + err.Error(), c)
   		return
    }
   response.OkWithData(dataSource, c)
}

// GetToyStorePublic 不需要鉴权的玩具店接口
// @Tags ToyStore
// @Summary 不需要鉴权的玩具店接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /TS/getToyStorePublic [get]
func (TSApi *ToyStoreApi) GetToyStorePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    TSService.GetToyStorePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的玩具店接口信息",
    }, "获取成功", c)
}
