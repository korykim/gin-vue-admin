package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PoiItemsRouter struct{}

// InitPoiItemsRouter 初始化 poiItems表 路由信息
func (s *PoiItemsRouter) InitPoiItemsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	poiItemsRouter := Router.Group("poiItems").Use(middleware.OperationRecord())
	poiItemsRouterWithoutRecord := Router.Group("poiItems")
	poiItemsRouterWithoutAuth := PublicRouter.Group("poiItems")
	{
		poiItemsRouter.POST("createPoiItems", poiItemsApi.CreatePoiItems)                   // 新建poiItems表
		poiItemsRouter.DELETE("deletePoiItems", poiItemsApi.DeletePoiItems)                 // 删除poiItems表
		poiItemsRouter.DELETE("deletePoiItemsByIds", poiItemsApi.DeletePoiItemsByIds)       // 批量删除poiItems表
		poiItemsRouter.PUT("updatePoiItems", poiItemsApi.UpdatePoiItems)                    // 更新poiItems表
		poiItemsRouter.POST("createPoiItemsInBatches", poiItemsApi.CreatePoiItemsInBatches) // 批量创建poiItems表
	}
	{
		poiItemsRouterWithoutRecord.GET("findPoiItems", poiItemsApi.FindPoiItems)       // 根据ID获取poiItems表
		poiItemsRouterWithoutRecord.GET("getPoiItemsList", poiItemsApi.GetPoiItemsList) // 获取poiItems表列表
	}
	{
		poiItemsRouterWithoutAuth.GET("getPoiItemsPublic", poiItemsApi.GetPoiItemsPublic) // poiItems表开放接口
	}
}
