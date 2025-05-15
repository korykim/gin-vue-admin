package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ToyStoreRouter struct {}

// InitToyStoreRouter 初始化 玩具店 路由信息
func (s *ToyStoreRouter) InitToyStoreRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	TSRouter := Router.Group("TS").Use(middleware.OperationRecord())
	TSRouterWithoutRecord := Router.Group("TS")
	TSRouterWithoutAuth := PublicRouter.Group("TS")
	{
		TSRouter.POST("createToyStore", TSApi.CreateToyStore)   // 新建玩具店
		TSRouter.DELETE("deleteToyStore", TSApi.DeleteToyStore) // 删除玩具店
		TSRouter.DELETE("deleteToyStoreByIds", TSApi.DeleteToyStoreByIds) // 批量删除玩具店
		TSRouter.PUT("updateToyStore", TSApi.UpdateToyStore)    // 更新玩具店
	}
	{
		TSRouterWithoutRecord.GET("findToyStore", TSApi.FindToyStore)        // 根据ID获取玩具店
		TSRouterWithoutRecord.GET("getToyStoreList", TSApi.GetToyStoreList)  // 获取玩具店列表
	}
	{
	    TSRouterWithoutAuth.GET("getToyStoreDataSource", TSApi.GetToyStoreDataSource)  // 获取玩具店数据源
	    TSRouterWithoutAuth.GET("getToyStorePublic", TSApi.GetToyStorePublic)  // 玩具店开放接口
	}
}
