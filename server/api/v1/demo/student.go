package demo

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/demo"
    demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type StudentApi struct {}



// CreateStudent 创建学生
// @Tags Student
// @Summary 创建学生
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.Student true "创建学生"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /student/createStudent [post]
func (studentApi *StudentApi) CreateStudent(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var student demo.Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = studentService.CreateStudent(ctx,&student)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteStudent 删除学生
// @Tags Student
// @Summary 删除学生
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.Student true "删除学生"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /student/deleteStudent [delete]
func (studentApi *StudentApi) DeleteStudent(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := studentService.DeleteStudent(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteStudentByIds 批量删除学生
// @Tags Student
// @Summary 批量删除学生
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /student/deleteStudentByIds [delete]
func (studentApi *StudentApi) DeleteStudentByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := studentService.DeleteStudentByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateStudent 更新学生
// @Tags Student
// @Summary 更新学生
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body demo.Student true "更新学生"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /student/updateStudent [put]
func (studentApi *StudentApi) UpdateStudent(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var student demo.Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = studentService.UpdateStudent(ctx,student)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindStudent 用id查询学生
// @Tags Student
// @Summary 用id查询学生
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询学生"
// @Success 200 {object} response.Response{data=demo.Student,msg=string} "查询成功"
// @Router /student/findStudent [get]
func (studentApi *StudentApi) FindStudent(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	restudent, err := studentService.GetStudent(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(restudent, c)
}
// GetStudentList 分页获取学生列表
// @Tags Student
// @Summary 分页获取学生列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.StudentSearch true "分页获取学生列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /student/getStudentList [get]
func (studentApi *StudentApi) GetStudentList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo demoReq.StudentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := studentService.GetStudentInfoList(ctx,pageInfo)
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

// GetStudentPublic 不需要鉴权的学生接口
// @Tags Student
// @Summary 不需要鉴权的学生接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /student/getStudentPublic [get]
func (studentApi *StudentApi) GetStudentPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    studentService.GetStudentPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的学生接口信息",
    }, "获取成功", c)
}
