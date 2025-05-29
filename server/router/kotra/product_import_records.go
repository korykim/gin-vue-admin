package kotra

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProductImportRecordsRouter struct {}

// InitProductImportRecordsRouter 初始化 未准入境产品表 路由信息
func (s *ProductImportRecordsRouter) InitProductImportRecordsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	productImportRecordsRouter := Router.Group("productImportRecords").Use(middleware.OperationRecord())
	productImportRecordsRouterWithoutRecord := Router.Group("productImportRecords")
	productImportRecordsRouterWithoutAuth := PublicRouter.Group("productImportRecords")
	{
		productImportRecordsRouter.POST("createProductImportRecords", productImportRecordsApi.CreateProductImportRecords)   // 新建未准入境产品表
		productImportRecordsRouter.DELETE("deleteProductImportRecords", productImportRecordsApi.DeleteProductImportRecords) // 删除未准入境产品表
		productImportRecordsRouter.DELETE("deleteProductImportRecordsByIds", productImportRecordsApi.DeleteProductImportRecordsByIds) // 批量删除未准入境产品表
		productImportRecordsRouter.PUT("updateProductImportRecords", productImportRecordsApi.UpdateProductImportRecords)    // 更新未准入境产品表
	}
	{
		productImportRecordsRouterWithoutRecord.GET("findProductImportRecords", productImportRecordsApi.FindProductImportRecords)        // 根据ID获取未准入境产品表
		productImportRecordsRouterWithoutRecord.GET("getProductImportRecordsList", productImportRecordsApi.GetProductImportRecordsList)  // 获取未准入境产品表列表
	}
	{
	    productImportRecordsRouterWithoutAuth.GET("getProductImportRecordsPublic", productImportRecordsApi.GetProductImportRecordsPublic)  // 未准入境产品表开放接口
	}
}
