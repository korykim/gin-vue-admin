package demo

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	StudentRouter

	PostersRouter
	TagsRouter
	PosterTagRouter
}

var (
	studentApi = api.ApiGroupApp.DemoApiGroup.StudentApi

	postersApi   = api.ApiGroupApp.DemoApiGroup.PostersApi
	tagsApi      = api.ApiGroupApp.DemoApiGroup.TagsApi
	posterTagApi = api.ApiGroupApp.DemoApiGroup.PosterTagApi
)

var RouterGroupApp = new(RouterGroup)

// InitDemoRouter 初始化Demo路由组
func (s *RouterGroup) InitDemoRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {

	s.StudentRouter.InitStudentRouter(Router, PublicRouter)
	s.PostersRouter.InitPostersRouter(Router, PublicRouter)
	s.TagsRouter.InitTagsRouter(Router, PublicRouter)
	s.PosterTagRouter.InitPosterTagRouter(Router, PublicRouter)
}
