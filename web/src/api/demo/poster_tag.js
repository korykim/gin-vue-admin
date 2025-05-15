import service from '@/utils/request'

// @Tags PosterTag
// @Summary 为海报添加标签
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body struct{PosterID uint `json:"posterId"`;TagIDs []uint `json:"tagIds"`} true "海报ID和标签ID列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /posterTag/addTagsToPoster [post]
export const addTagsToPoster = (data) => {
  return service({
    url: '/posterTag/addTagsToPoster',
    method: 'post',
    data
  })
}

// @Tags PosterTag
// @Summary 从海报中移除标签
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body struct{PosterID uint `json:"posterId"`;TagIDs []uint `json:"tagIds"`} true "海报ID和标签ID列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"移除成功"}"
// @Router /posterTag/removeTagsFromPoster [delete]
export const removeTagsFromPoster = (data) => {
  return service({
    url: '/posterTag/removeTagsFromPoster',
    method: 'delete',
    data
  })
}

// @Tags PosterTag
// @Summary 获取海报的所有标签
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param posterId query uint true "海报ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /posterTag/getPosterTags [get]
export const getPosterTags = (params) => {
  return service({
    url: '/posterTag/getPosterTags',
    method: 'get',
    params
  })
}

// @Tags PosterTag
// @Summary 获取标签关联的所有海报
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param tagId query uint true "标签ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /posterTag/getTagPosters [get]
export const getTagPosters = (params) => {
  return service({
    url: '/posterTag/getTagPosters',
    method: 'get',
    params
  })
}

// @Tags PosterTag
// @Summary 分页获取海报标签关联列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取海报标签关联列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /posterTag/getPosterTagRelations [get]
export const getPosterTagRelations = (params) => {
  return service({
    url: '/posterTag/getPosterTagRelations',
    method: 'get',
    params
  })
}

// @Tags PosterTag
// @Summary 获取标签使用统计
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /posterTag/getTagsUsageStats [get]
export const getTagsUsageStats = () => {
  return service({
    url: '/posterTag/getTagsUsageStats',
    method: 'get'
  })
} 