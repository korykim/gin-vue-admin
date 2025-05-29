package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		demoRouter := router.RouterGroupApp.Demo
		demoRouter.InitStudentRouter(privateGroup, publicGroup)
		demoRouter.InitPostersRouter(privateGroup, publicGroup)
		demoRouter.InitTagsRouter(privateGroup, publicGroup)
		demoRouter.InitPosterTagRouter(privateGroup, publicGroup)
		demoRouter.InitEmployeeProfileRouter(privateGroup, publicGroup)
		demoRouter.InitToyStoreRouter(privateGroup, publicGroup)
		demoRouter.InitPoiItemsRouter(privateGroup, publicGroup)
	}
	{
		kotraRouter := router.RouterGroupApp.Kotra
		kotraRouter.InitProductImportRecordsRouter(privateGroup, publicGroup)
	}
}
