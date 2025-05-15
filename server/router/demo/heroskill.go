package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HeroSkillRouter struct {}

// InitHeroSkillRouter 初始化 技能 路由信息
func (s *HeroSkillRouter) InitHeroSkillRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	heroskillRouter := Router.Group("heroskill").Use(middleware.OperationRecord())
	heroskillRouterWithoutRecord := Router.Group("heroskill")
	heroskillRouterWithoutAuth := PublicRouter.Group("heroskill")
	{
		heroskillRouter.POST("createHeroSkill", heroskillApi.CreateHeroSkill)   // 新建技能
		heroskillRouter.DELETE("deleteHeroSkill", heroskillApi.DeleteHeroSkill) // 删除技能
		heroskillRouter.DELETE("deleteHeroSkillByIds", heroskillApi.DeleteHeroSkillByIds) // 批量删除技能
		heroskillRouter.PUT("updateHeroSkill", heroskillApi.UpdateHeroSkill)    // 更新技能
	}
	{
		heroskillRouterWithoutRecord.GET("findHeroSkill", heroskillApi.FindHeroSkill)        // 根据ID获取技能
		heroskillRouterWithoutRecord.GET("getHeroSkillList", heroskillApi.GetHeroSkillList)  // 获取技能列表
	}
	{
	    heroskillRouterWithoutAuth.GET("getHeroSkillPublic", heroskillApi.GetHeroSkillPublic)  // 技能开放接口
	}
}
