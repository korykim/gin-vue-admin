package kotra

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ProductImportRecordsApi }

var productImportRecordsService = service.ServiceGroupApp.KotraServiceGroup.ProductImportRecordsService
