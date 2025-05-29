package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/demo"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/kotra"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	DemoServiceGroup    demo.ServiceGroup
	KotraServiceGroup   kotra.ServiceGroup
}
