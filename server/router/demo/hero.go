package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HeroRouter struct {}

// InitHeroRouter 初始化 英雄信息 路由信息
func (s *HeroRouter) InitHeroRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	heroRouter := Router.Group("hero").Use(middleware.OperationRecord())
	heroRouterWithoutRecord := Router.Group("hero")
	heroRouterWithoutAuth := PublicRouter.Group("hero")
	{
		heroRouter.POST("createHero", heroApi.CreateHero)   // 新建英雄信息
		heroRouter.DELETE("deleteHero", heroApi.DeleteHero) // 删除英雄信息
		heroRouter.DELETE("deleteHeroByIds", heroApi.DeleteHeroByIds) // 批量删除英雄信息
		heroRouter.PUT("updateHero", heroApi.UpdateHero)    // 更新英雄信息
	}
	{
		heroRouterWithoutRecord.GET("findHero", heroApi.FindHero)        // 根据ID获取英雄信息
		heroRouterWithoutRecord.GET("getHeroList", heroApi.GetHeroList)  // 获取英雄信息列表
	}
	{
	    heroRouterWithoutAuth.GET("getHeroPublic", heroApi.GetHeroPublic)  // 英雄信息开放接口
	}
}
