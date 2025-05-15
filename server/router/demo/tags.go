package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TagsRouter struct {}

// InitTagsRouter 初始化 tags表 路由信息
func (s *TagsRouter) InitTagsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	tagsRouter := Router.Group("tags").Use(middleware.OperationRecord())
	tagsRouterWithoutRecord := Router.Group("tags")
	tagsRouterWithoutAuth := PublicRouter.Group("tags")
	{
		tagsRouter.POST("createTags", tagsApi.CreateTags)   // 新建tags表
		tagsRouter.DELETE("deleteTags", tagsApi.DeleteTags) // 删除tags表
		tagsRouter.DELETE("deleteTagsByIds", tagsApi.DeleteTagsByIds) // 批量删除tags表
		tagsRouter.PUT("updateTags", tagsApi.UpdateTags)    // 更新tags表
	}
	{
		tagsRouterWithoutRecord.GET("findTags", tagsApi.FindTags)        // 根据ID获取tags表
		tagsRouterWithoutRecord.GET("getTagsList", tagsApi.GetTagsList)  // 获取tags表列表
	}
	{
	    tagsRouterWithoutAuth.GET("getTagsPublic", tagsApi.GetTagsPublic)  // tags表开放接口
	}
}
