package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/demo"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/kotra"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Demo    demo.RouterGroup
	Kotra   kotra.RouterGroup
}
