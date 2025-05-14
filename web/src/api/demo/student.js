import service from '@/utils/request'
// @Tags Student
// @Summary 创建学生
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Student true "创建学生"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /student/createStudent [post]
export const createStudent = (data) => {
  return service({
    url: '/student/createStudent',
    method: 'post',
    data
  })
}

// @Tags Student
// @Summary 删除学生
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Student true "删除学生"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /student/deleteStudent [delete]
export const deleteStudent = (params) => {
  return service({
    url: '/student/deleteStudent',
    method: 'delete',
    params
  })
}

// @Tags Student
// @Summary 批量删除学生
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除学生"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /student/deleteStudent [delete]
export const deleteStudentByIds = (params) => {
  return service({
    url: '/student/deleteStudentByIds',
    method: 'delete',
    params
  })
}

// @Tags Student
// @Summary 更新学生
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Student true "更新学生"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /student/updateStudent [put]
export const updateStudent = (data) => {
  return service({
    url: '/student/updateStudent',
    method: 'put',
    data
  })
}

// @Tags Student
// @Summary 用id查询学生
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Student true "用id查询学生"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /student/findStudent [get]
export const findStudent = (params) => {
  return service({
    url: '/student/findStudent',
    method: 'get',
    params
  })
}

// @Tags Student
// @Summary 分页获取学生列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取学生列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/getStudentList [get]
export const getStudentList = (params) => {
  return service({
    url: '/student/getStudentList',
    method: 'get',
    params
  })
}

// @Tags Student
// @Summary 不需要鉴权的学生接口
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.StudentSearch true "分页获取学生列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /student/getStudentPublic [get]
export const getStudentPublic = () => {
  return service({
    url: '/student/getStudentPublic',
    method: 'get',
  })
}
