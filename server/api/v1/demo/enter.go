package demo

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	StudentApi
	PostersApi
	TagsApi
	PosterTagApi
	EmployeeProfileApi
	ToyStoreApi
	PosterApi
	TagApi
	PoiItemsApi
}

var (
	studentService   = service.ServiceGroupApp.DemoServiceGroup.StudentService
	postersService   = service.ServiceGroupApp.DemoServiceGroup.PostersService
	tagsService      = service.ServiceGroupApp.DemoServiceGroup.TagsService
	posterTagService = service.ServiceGroupApp.DemoServiceGroup.PosterTagService
	TSEService       = service.ServiceGroupApp.DemoServiceGroup.EmployeeProfileService
	TSService        = service.ServiceGroupApp.DemoServiceGroup.ToyStoreService
	posterService    = service.ServiceGroupApp.DemoServiceGroup.PosterService
	tagService       = service.ServiceGroupApp.DemoServiceGroup.TagService
	poiItemsService  = service.ServiceGroupApp.DemoServiceGroup.PoiItemsService
)
