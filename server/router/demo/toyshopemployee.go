package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EmployeeProfileRouter struct {}

// InitEmployeeProfileRouter 初始化 员工信息 路由信息
func (s *EmployeeProfileRouter) InitEmployeeProfileRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	TSERouter := Router.Group("TSE").Use(middleware.OperationRecord())
	TSERouterWithoutRecord := Router.Group("TSE")
	TSERouterWithoutAuth := PublicRouter.Group("TSE")
	{
		TSERouter.POST("createEmployeeProfile", TSEApi.CreateEmployeeProfile)   // 新建员工信息
		TSERouter.DELETE("deleteEmployeeProfile", TSEApi.DeleteEmployeeProfile) // 删除员工信息
		TSERouter.DELETE("deleteEmployeeProfileByIds", TSEApi.DeleteEmployeeProfileByIds) // 批量删除员工信息
		TSERouter.PUT("updateEmployeeProfile", TSEApi.UpdateEmployeeProfile)    // 更新员工信息
	}
	{
		TSERouterWithoutRecord.GET("findEmployeeProfile", TSEApi.FindEmployeeProfile)        // 根据ID获取员工信息
		TSERouterWithoutRecord.GET("getEmployeeProfileList", TSEApi.GetEmployeeProfileList)  // 获取员工信息列表
	}
	{
	    TSERouterWithoutAuth.GET("getEmployeeProfilePublic", TSEApi.GetEmployeeProfilePublic)  // 员工信息开放接口
	}
}
