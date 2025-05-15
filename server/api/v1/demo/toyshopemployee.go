package demo

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/demo"
    demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type EmployeeProfileApi struct {}



// CreateEmployeeProfile 创建员工信息
// @Tags EmployeeProfile
// @Summary 创建员工信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.EmployeeProfile true "创建员工信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /TSE/createEmployeeProfile [post]
func (TSEApi *EmployeeProfileApi) CreateEmployeeProfile(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var TSE demo.EmployeeProfile
	err := c.ShouldBindJSON(&TSE)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = TSEService.CreateEmployeeProfile(ctx,&TSE)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteEmployeeProfile 删除员工信息
// @Tags EmployeeProfile
// @Summary 删除员工信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.EmployeeProfile true "删除员工信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /TSE/deleteEmployeeProfile [delete]
func (TSEApi *EmployeeProfileApi) DeleteEmployeeProfile(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := TSEService.DeleteEmployeeProfile(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEmployeeProfileByIds 批量删除员工信息
// @Tags EmployeeProfile
// @Summary 批量删除员工信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /TSE/deleteEmployeeProfileByIds [delete]
func (TSEApi *EmployeeProfileApi) DeleteEmployeeProfileByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := TSEService.DeleteEmployeeProfileByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEmployeeProfile 更新员工信息
// @Tags EmployeeProfile
// @Summary 更新员工信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.EmployeeProfile true "更新员工信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /TSE/updateEmployeeProfile [put]
func (TSEApi *EmployeeProfileApi) UpdateEmployeeProfile(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var TSE demo.EmployeeProfile
	err := c.ShouldBindJSON(&TSE)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = TSEService.UpdateEmployeeProfile(ctx,TSE)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEmployeeProfile 用id查询员工信息
// @Tags EmployeeProfile
// @Summary 用id查询员工信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询员工信息"
// @Success 200 {object} response.Response{data=demo.EmployeeProfile,msg=string} "查询成功"
// @Router /TSE/findEmployeeProfile [get]
func (TSEApi *EmployeeProfileApi) FindEmployeeProfile(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reTSE, err := TSEService.GetEmployeeProfile(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reTSE, c)
}
// GetEmployeeProfileList 分页获取员工信息列表
// @Tags EmployeeProfile
// @Summary 分页获取员工信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.EmployeeProfileSearch true "分页获取员工信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /TSE/getEmployeeProfileList [get]
func (TSEApi *EmployeeProfileApi) GetEmployeeProfileList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo demoReq.EmployeeProfileSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := TSEService.GetEmployeeProfileInfoList(ctx,pageInfo)
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

// GetEmployeeProfilePublic 不需要鉴权的员工信息接口
// @Tags EmployeeProfile
// @Summary 不需要鉴权的员工信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /TSE/getEmployeeProfilePublic [get]
func (TSEApi *EmployeeProfileApi) GetEmployeeProfilePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    TSEService.GetEmployeeProfilePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的员工信息接口信息",
    }, "获取成功", c)
}
