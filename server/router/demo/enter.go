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
	EmployeeProfileRouter
	ToyStoreRouter
	PosterRouter
	TagRouter // InitDemoRouter 初始化Demo路由组
	// 新增路由
	// 新增API
	// 新增路由初始化
	PoiItemsRouter
}

var (
	studentApi   = api.ApiGroupApp.DemoApiGroup.StudentApi
	postersApi   = api.ApiGroupApp.DemoApiGroup.PostersApi
	tagsApi      = api.ApiGroupApp.DemoApiGroup.TagsApi
	posterTagApi = api.ApiGroupApp.DemoApiGroup.PosterTagApi
	TSEApi       = api.ApiGroupApp.DemoApiGroup.EmployeeProfileApi
	TSApi        = api.ApiGroupApp.DemoApiGroup.ToyStoreApi
	posterApi    = api.ApiGroupApp.DemoApiGroup.PosterApi
	tagApi       = api.ApiGroupApp.DemoApiGroup.TagApi
	poiItemsApi  = api.ApiGroupApp.DemoApiGroup.PoiItemsApi
)
var RouterGroupApp = new(RouterGroup)

func (s *RouterGroup) InitDemoRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	s.StudentRouter.InitStudentRouter(Router, PublicRouter)
	s.PostersRouter.InitPostersRouter(Router, PublicRouter)
	s.TagsRouter.InitTagsRouter(Router, PublicRouter)
	s.PosterTagRouter.InitPosterTagRouter(Router, PublicRouter)
	s.EmployeeProfileRouter.InitEmployeeProfileRouter(Router, PublicRouter)
	s.ToyStoreRouter.InitToyStoreRouter(Router, PublicRouter)
	s.PosterRouter.InitPosterRouter(Router, PublicRouter)
	s.TagRouter.InitTagRouter(Router, PublicRouter)
}
