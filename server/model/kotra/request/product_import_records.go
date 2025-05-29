package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ProductImportRecordsSearch struct {
	HsCode              *string `json:"hsCode" form:"hsCode"`
	ProductName         *string `json:"productName" form:"productName"`
	FactoryInformation  *string `json:"factoryInformation" form:"factoryInformation"`
	ImporterInformation *string `json:"importerInformation" form:"importerInformation"`
	RecordNumber        *string `json:"recordNumber" form:"recordNumber"`
	ImportPort          *string `json:"importPort" form:"importPort"`
	Years               *int    `json:"years" form:"years"`
	Month               *int    `json:"month" form:"month"`
	PlaceOfOrigin       *string `json:"placeOfOrigin" form:"placeOfOrigin"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
