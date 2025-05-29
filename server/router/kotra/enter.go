package kotra

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ProductImportRecordsRouter }

var productImportRecordsApi = api.ApiGroupApp.KotraApiGroup.ProductImportRecordsApi
