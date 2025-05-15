package demo

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	StudentApi
	HeroApi
	HeroSkillApi
	PostersApi
	TagsApi
	PosterTagApi
}

var (
	studentService   = service.ServiceGroupApp.DemoServiceGroup.StudentService
	heroService      = service.ServiceGroupApp.DemoServiceGroup.HeroService
	heroskillService = service.ServiceGroupApp.DemoServiceGroup.HeroSkillService
	postersService   = service.ServiceGroupApp.DemoServiceGroup.PostersService
	tagsService      = service.ServiceGroupApp.DemoServiceGroup.TagsService
	posterTagService = service.ServiceGroupApp.DemoServiceGroup.PosterTagService
)
