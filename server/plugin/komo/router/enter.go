package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/komo/api"

var (
	Router    = new(router)
	apiQmUser = api.Api.QmUser
)

type router struct{ QmUser qmUser }
