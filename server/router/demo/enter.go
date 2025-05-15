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
	EmployeeProfileRouter // InitDemoRouter 初始化Demo路由组
	ToyStoreRouter
	// 新增路由
	PosterRouter
	TagRouter
}

var (
	studentApi   = api.ApiGroupApp.DemoApiGroup.StudentApi
	postersApi   = api.ApiGroupApp.DemoApiGroup.PostersApi
	tagsApi      = api.ApiGroupApp.DemoApiGroup.TagsApi
	posterTagApi = api.ApiGroupApp.DemoApiGroup.PosterTagApi
	TSEApi       = api.ApiGroupApp.DemoApiGroup.EmployeeProfileApi
	TSApi        = api.ApiGroupApp.DemoApiGroup.ToyStoreApi
	// 新增API
	posterApi = api.ApiGroupApp.DemoApiGroup.PosterApi
	tagApi    = api.ApiGroupApp.DemoApiGroup.TagApi
)
var RouterGroupApp = new(RouterGroup)

func (s *RouterGroup) InitDemoRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	s.StudentRouter.InitStudentRouter(Router, PublicRouter)
	s.PostersRouter.InitPostersRouter(Router, PublicRouter)
	s.TagsRouter.InitTagsRouter(Router, PublicRouter)
	s.PosterTagRouter.InitPosterTagRouter(Router, PublicRouter)
	s.EmployeeProfileRouter.InitEmployeeProfileRouter(Router, PublicRouter)
	s.ToyStoreRouter.InitToyStoreRouter(Router, PublicRouter)
	// 新增路由初始化
	s.PosterRouter.InitPosterRouter(Router, PublicRouter)
	s.TagRouter.InitTagRouter(Router, PublicRouter)
}
