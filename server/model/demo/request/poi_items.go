package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PoiItemsSearch struct {
	Name        *string `json:"name" form:"name"`
	Address     *string `json:"address" form:"address"`
	Tel         *string `json:"tel" form:"tel"`
	Pname       *string `json:"pname" form:"pname"`
	Cityname    *string `json:"cityname" form:"cityname"`
	Adname      *string `json:"adname" form:"adname"`
	PoiType     *string `json:"poiType" form:"poiType"`
	BigCategory *string `json:"BigCategory" form:"BigCategory"`
	MidCategory *string `json:"MidCategory" form:"MidCategory"`
	SubCategory *string `json:"SubCategory" form:"SubCategory"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
