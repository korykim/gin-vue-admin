package demo

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	StudentApi

	PostersApi
	TagsApi
	PosterTagApi
}

var (
	studentService = service.ServiceGroupApp.DemoServiceGroup.StudentService

	postersService   = service.ServiceGroupApp.DemoServiceGroup.PostersService
	tagsService      = service.ServiceGroupApp.DemoServiceGroup.TagsService
	posterTagService = service.ServiceGroupApp.DemoServiceGroup.PosterTagService
)
