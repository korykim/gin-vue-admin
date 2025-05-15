import service from '@/utils/request'
// @Tags ToyStore
// @Summary 创建玩具店
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ToyStore true "创建玩具店"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /TS/createToyStore [post]
export const createToyStore = (data) => {
  return service({
    url: '/TS/createToyStore',
    method: 'post',
    data
  })
}

// @Tags ToyStore
// @Summary 删除玩具店
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ToyStore true "删除玩具店"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TS/deleteToyStore [delete]
export const deleteToyStore = (params) => {
  return service({
    url: '/TS/deleteToyStore',
    method: 'delete',
    params
  })
}

// @Tags ToyStore
// @Summary 批量删除玩具店
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除玩具店"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TS/deleteToyStore [delete]
export const deleteToyStoreByIds = (params) => {
  return service({
    url: '/TS/deleteToyStoreByIds',
    method: 'delete',
    params
  })
}

// @Tags ToyStore
// @Summary 更新玩具店
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ToyStore true "更新玩具店"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TS/updateToyStore [put]
export const updateToyStore = (data) => {
  return service({
    url: '/TS/updateToyStore',
    method: 'put',
    data
  })
}

// @Tags ToyStore
// @Summary 用id查询玩具店
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ToyStore true "用id查询玩具店"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TS/findToyStore [get]
export const findToyStore = (params) => {
  return service({
    url: '/TS/findToyStore',
    method: 'get',
    params
  })
}

// @Tags ToyStore
// @Summary 分页获取玩具店列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取玩具店列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TS/getToyStoreList [get]
export const getToyStoreList = (params) => {
  return service({
    url: '/TS/getToyStoreList',
    method: 'get',
    params
  })
}
// @Tags ToyStore
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TS/findToyStoreDataSource [get]
export const getToyStoreDataSource = () => {
  return service({
    url: '/TS/getToyStoreDataSource',
    method: 'get',
  })
}

// @Tags ToyStore
// @Summary 不需要鉴权的玩具店接口
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.ToyStoreSearch true "分页获取玩具店列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /TS/getToyStorePublic [get]
export const getToyStorePublic = () => {
  return service({
    url: '/TS/getToyStorePublic',
    method: 'get',
  })
}
