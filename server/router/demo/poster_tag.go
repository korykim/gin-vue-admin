package demo

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PosterTagRouter struct{}

// InitPosterTagRouter 初始化 PosterTag 路由信息
func (s *PosterTagRouter) InitPosterTagRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	posterTagRouter := Router.Group("posterTag").Use(middleware.OperationRecord())
	posterTagRouterWithoutRecord := Router.Group("posterTag")
	var posterTagApi = v1.ApiGroupApp.DemoApiGroup.PosterTagApi
	{
		posterTagRouter.POST("addTagsToPoster", posterTagApi.AddTagsToPoster)             // 为海报添加标签
		posterTagRouter.DELETE("removeTagsFromPoster", posterTagApi.RemoveTagsFromPoster) // 从海报中移除标签
	}

	{
		// 不需要记录操作日志的路由
		posterTagRouterWithoutRecord.GET("getPosterTags", posterTagApi.GetPosterTags)                 // 获取海报的所有标签
		posterTagRouterWithoutRecord.GET("getTagPosters", posterTagApi.GetTagPosters)                 // 获取标签关联的所有海报
		posterTagRouterWithoutRecord.GET("getPosterTagRelations", posterTagApi.GetPosterTagRelations) // 分页获取海报标签关联
	}
}
