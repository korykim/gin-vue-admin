package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StudentRouter struct {}

// InitStudentRouter 初始化 学生 路由信息
func (s *StudentRouter) InitStudentRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	studentRouter := Router.Group("student").Use(middleware.OperationRecord())
	studentRouterWithoutRecord := Router.Group("student")
	studentRouterWithoutAuth := PublicRouter.Group("student")
	{
		studentRouter.POST("createStudent", studentApi.CreateStudent)   // 新建学生
		studentRouter.DELETE("deleteStudent", studentApi.DeleteStudent) // 删除学生
		studentRouter.DELETE("deleteStudentByIds", studentApi.DeleteStudentByIds) // 批量删除学生
		studentRouter.PUT("updateStudent", studentApi.UpdateStudent)    // 更新学生
	}
	{
		studentRouterWithoutRecord.GET("findStudent", studentApi.FindStudent)        // 根据ID获取学生
		studentRouterWithoutRecord.GET("getStudentList", studentApi.GetStudentList)  // 获取学生列表
	}
	{
	    studentRouterWithoutAuth.GET("getStudentPublic", studentApi.GetStudentPublic)  // 学生开放接口
	}
}
