import service from '@/utils/request'
// @Tags ProductImportRecords
// @Summary 创建未准入境产品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductImportRecords true "创建未准入境产品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /productImportRecords/createProductImportRecords [post]
export const createProductImportRecords = (data) => {
  return service({
    url: '/productImportRecords/createProductImportRecords',
    method: 'post',
    data
  })
}

// @Tags ProductImportRecords
// @Summary 删除未准入境产品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductImportRecords true "删除未准入境产品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /productImportRecords/deleteProductImportRecords [delete]
export const deleteProductImportRecords = (params) => {
  return service({
    url: '/productImportRecords/deleteProductImportRecords',
    method: 'delete',
    params
  })
}

// @Tags ProductImportRecords
// @Summary 批量删除未准入境产品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除未准入境产品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /productImportRecords/deleteProductImportRecords [delete]
export const deleteProductImportRecordsByIds = (params) => {
  return service({
    url: '/productImportRecords/deleteProductImportRecordsByIds',
    method: 'delete',
    params
  })
}

// @Tags ProductImportRecords
// @Summary 更新未准入境产品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductImportRecords true "更新未准入境产品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /productImportRecords/updateProductImportRecords [put]
export const updateProductImportRecords = (data) => {
  return service({
    url: '/productImportRecords/updateProductImportRecords',
    method: 'put',
    data
  })
}

// @Tags ProductImportRecords
// @Summary 用id查询未准入境产品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ProductImportRecords true "用id查询未准入境产品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /productImportRecords/findProductImportRecords [get]
export const findProductImportRecords = (params) => {
  return service({
    url: '/productImportRecords/findProductImportRecords',
    method: 'get',
    params
  })
}

// @Tags ProductImportRecords
// @Summary 分页获取未准入境产品表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取未准入境产品表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /productImportRecords/getProductImportRecordsList [get]
export const getProductImportRecordsList = (params) => {
  return service({
    url: '/productImportRecords/getProductImportRecordsList',
    method: 'get',
    params
  })
}

// @Tags ProductImportRecords
// @Summary 不需要鉴权的未准入境产品表接口
// @Accept application/json
// @Produce application/json
// @Param data query kotraReq.ProductImportRecordsSearch true "分页获取未准入境产品表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /productImportRecords/getProductImportRecordsPublic [get]
export const getProductImportRecordsPublic = () => {
  return service({
    url: '/productImportRecords/getProductImportRecordsPublic',
    method: 'get',
  })
}
