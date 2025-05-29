import service from '@/utils/request'
// @Tags PoiItems
// @Summary 创建poiItems表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PoiItems true "创建poiItems表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /poiItems/createPoiItems [post]
export const createPoiItems = (data) => {
  // 处理index字段
  if (data.poiIndex !== undefined) {
    data.index = data.poiIndex
    delete data.poiIndex
  }
  return service({
    url: '/poiItems/createPoiItems',
    method: 'post',
    data
  })
}

// @Tags PoiItems 
// @Summary 批量创建poiItems表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PoiItems true "批量创建poiItems表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /poiItems/createPoiItemsInBatches [post]
export const createPoiItemsInBatches = (data) => {
  // 数据清理工作
  const cleanedData = data.map(item => {
    // 创建一个新对象，避免修改原始数据
    const cleanedItem = { ...item };
    
    // 处理location字段，确保为数组格式
    if (!cleanedItem.location) {
      cleanedItem.location = [];
    } else if (typeof cleanedItem.location === 'object' && !Array.isArray(cleanedItem.location)) {
      // 如果是对象格式 {lng, lat} 转换为数组 [lng, lat]
      if (cleanedItem.location.lng !== undefined && cleanedItem.location.lat !== undefined) {
        cleanedItem.location = [cleanedItem.location.lng, cleanedItem.location.lat];
      } else {
        cleanedItem.location = [];
      }
    }
    
    // 处理photos字段，确保为数组格式
    if (!cleanedItem.photos) {
      cleanedItem.photos = [];
    }
    
    // 处理cost字段，确保为数字类型，为空时设置为0.00
    if (cleanedItem.cost === undefined || cleanedItem.cost === null) {
      cleanedItem.cost = 0.00;
    } else if (typeof cleanedItem.cost === 'string') {
      // 将字符串转换为数字
      cleanedItem.cost = parseFloat(cleanedItem.cost) || 0.00;
    }
    
    // 处理入口位置和出口位置字段，后端结构体中是string类型
    // 确保字段名称与后端一致
    if (cleanedItem.entr_location) {
      // 如果是数组，转为JSON字符串
      if (Array.isArray(cleanedItem.entr_location)) {
        cleanedItem.entrLocation = JSON.stringify(cleanedItem.entr_location);
      } else {
        cleanedItem.entrLocation = String(cleanedItem.entr_location);
      }
      // 删除原字段，使用Go结构体中定义的字段名
      delete cleanedItem.entr_location;
    }
    
    if (cleanedItem.exit_location) {
      // 如果是数组，转为JSON字符串
      if (Array.isArray(cleanedItem.exit_location)) {
        cleanedItem.exitLocation = JSON.stringify(cleanedItem.exit_location);
      } else {
        cleanedItem.exitLocation = String(cleanedItem.exit_location);
      }
      // 删除原字段，使用Go结构体中定义的字段名
      delete cleanedItem.exit_location;
    }
    
    // 处理rating字段，确保为数值类型
    if (cleanedItem.rating !== undefined && typeof cleanedItem.rating === 'string') {
      cleanedItem.rating = parseFloat(cleanedItem.rating) || 0;
    }
    
    // 处理BigCategory、MidCategory、SubCategory字段
    if (cleanedItem.type && typeof cleanedItem.type === 'string') {
      const typeArr = cleanedItem.type.split(';');
      cleanedItem.BigCategory = typeArr[0] || '';
      cleanedItem.MidCategory = typeArr[1] || '';
      cleanedItem.SubCategory = typeArr[2] || '';
    }
    
    // 移除不需要的字段
    delete cleanedItem._idx;
    delete cleanedItem.index;
    
    return cleanedItem;
  });
  
  return service({
    url: '/poiItems/createPoiItemsInBatches',
    method: 'post',
    data: cleanedData
  });
}


// @Tags PoiItems
// @Summary 删除poiItems表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PoiItems true "删除poiItems表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /poiItems/deletePoiItems [delete]
export const deletePoiItems = (params) => {
  return service({
    url: '/poiItems/deletePoiItems',
    method: 'delete',
    params
  })
}

// @Tags PoiItems
// @Summary 批量删除poiItems表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除poiItems表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /poiItems/deletePoiItems [delete]
export const deletePoiItemsByIds = (params) => {
  return service({
    url: '/poiItems/deletePoiItemsByIds',
    method: 'delete',
    params
  })
}

// @Tags PoiItems
// @Summary 更新poiItems表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PoiItems true "更新poiItems表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /poiItems/updatePoiItems [put]
export const updatePoiItems = (data) => {
  // 处理index字段
  if (data.poiIndex !== undefined) {
    data.index = data.poiIndex
    delete data.poiIndex
  }
  return service({
    url: '/poiItems/updatePoiItems',
    method: 'put',
    data
  })
}

// @Tags PoiItems
// @Summary 用id查询poiItems表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PoiItems true "用id查询poiItems表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /poiItems/findPoiItems [get]
export const findPoiItems = (params) => {
  return service({
    url: '/poiItems/findPoiItems',
    method: 'get',
    params
  }).then(res => {
    // 处理返回数据
    if (res.code === 0 && res.data && res.data.index !== undefined) {
      res.data.poiIndex = res.data.index
    }
    return res
  })
}

// @Tags PoiItems
// @Summary 分页获取poiItems表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取poiItems表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /poiItems/getPoiItemsList [get]
export const getPoiItemsList = (params) => {
  return service({
    url: '/poiItems/getPoiItemsList',
    method: 'get',
    params
  }).then(res => {
    // 处理返回数据
    if (res.code === 0 && res.data && res.data.list) {
      res.data.list.forEach(item => {
        if (item.index !== undefined) {
          item.poiIndex = item.index
        }
      })
    }
    return res
  })
}

// @Tags PoiItems
// @Summary 不需要鉴权的poiItems表接口
// @Accept application/json
// @Produce application/json
// @Param data query demoReq.PoiItemsSearch true "分页获取poiItems表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /poiItems/getPoiItemsPublic [get]
export const getPoiItemsPublic = () => {
  return service({
    url: '/poiItems/getPoiItemsPublic',
    method: 'get',
  })
}
