package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PostersRouter struct {}

// InitPostersRouter 初始化 posters表 路由信息
func (s *PostersRouter) InitPostersRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	postersRouter := Router.Group("posters").Use(middleware.OperationRecord())
	postersRouterWithoutRecord := Router.Group("posters")
	postersRouterWithoutAuth := PublicRouter.Group("posters")
	{
		postersRouter.POST("createPosters", postersApi.CreatePosters)   // 新建posters表
		postersRouter.DELETE("deletePosters", postersApi.DeletePosters) // 删除posters表
		postersRouter.DELETE("deletePostersByIds", postersApi.DeletePostersByIds) // 批量删除posters表
		postersRouter.PUT("updatePosters", postersApi.UpdatePosters)    // 更新posters表
	}
	{
		postersRouterWithoutRecord.GET("findPosters", postersApi.FindPosters)        // 根据ID获取posters表
		postersRouterWithoutRecord.GET("getPostersList", postersApi.GetPostersList)  // 获取posters表列表
	}
	{
	    postersRouterWithoutAuth.GET("getPostersPublic", postersApi.GetPostersPublic)  // posters表开放接口
	}
}
