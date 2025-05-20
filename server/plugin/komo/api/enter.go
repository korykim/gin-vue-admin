package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/komo/service"

var (
	Api           = new(api)
	serviceQmUser = service.Service.QmUser
)

type api struct{ QmUser qmUser }
