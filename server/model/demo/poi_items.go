// 自动生成模板PoiItems
package demo

import (
	"gorm.io/datatypes"
)

// poiItems表 结构体  PoiItems
type PoiItems struct {
	Id           *string        `json:"id" form:"id" gorm:"primarykey;comment:POI唯一标识;column:id;size:20;"`                         //POI唯一标识
	Name         *string        `json:"name" form:"name" gorm:"comment:POI名称;column:name;size:100;"`                               //POI名称
	Type         *string        `json:"type" form:"type" gorm:"comment:类型描述;column:type;size:100;"`                                //类型描述
	Location     datatypes.JSON `json:"location" form:"location" gorm:"comment:位置坐标[经度,纬度];column:location;" swaggertype:"object"` //位置坐标[经度,纬度]
	Address      *string        `json:"address" form:"address" gorm:"comment:详细地址;column:address;size:255;"`                       //详细地址
	Tel          *string        `json:"tel" form:"tel" gorm:"comment:联系电话;column:tel;size:100;"`                                   //联系电话
	Distance     *int           `json:"distance" form:"distance" gorm:"comment:距离(米);column:distance;size:10;"`                    //距离(米)
	Shopinfo     *string        `json:"shopinfo" form:"shopinfo" gorm:"comment:商铺信息;column:shopinfo;size:50;"`                     //商铺信息
	Website      *string        `json:"website" form:"website" gorm:"comment:网站;column:website;size:255;"`                         //网站
	Pcode        *string        `json:"pcode" form:"pcode" gorm:"comment:省份编码;column:pcode;size:10;"`                              //省份编码
	Citycode     *string        `json:"citycode" form:"citycode" gorm:"comment:城市编码;column:citycode;size:10;"`                     //城市编码
	Adcode       *string        `json:"adcode" form:"adcode" gorm:"comment:区域编码;column:adcode;size:10;"`                           //区域编码
	Postcode     *string        `json:"postcode" form:"postcode" gorm:"comment:邮政编码;column:postcode;size:10;"`                     //邮政编码
	Pname        *string        `json:"pname" form:"pname" gorm:"comment:省份名称;column:pname;size:50;"`                              //省份名称
	Cityname     *string        `json:"cityname" form:"cityname" gorm:"comment:城市名称;column:cityname;size:50;"`                     //城市名称
	Adname       *string        `json:"adname" form:"adname" gorm:"comment:区域名称;column:adname;size:50;"`                           //区域名称
	Email        *string        `json:"email" form:"email" gorm:"comment:电子邮箱;column:email;size:100;"`                             //电子邮箱
	Photos       datatypes.JSON `json:"photos" form:"photos" gorm:"comment:照片信息;column:photos;" swaggertype:"array,object"`        //照片信息
	EntrLocation *string        `json:"entrLocation" form:"entrLocation" gorm:"comment:入口位置;column:entr_location;"`                //入口位置
	ExitLocation *string        `json:"exitLocation" form:"exitLocation" gorm:"comment:出口位置;column:exit_location;"`                //出口位置
	Groupbuy     *bool          `json:"groupbuy" form:"groupbuy" gorm:"comment:是否支持团购;column:groupbuy;"`                           //是否支持团购
	Discount     *bool          `json:"discount" form:"discount" gorm:"comment:是否有优惠;column:discount;"`                            //是否有优惠
	IndoorMap    *bool          `json:"indoorMap" form:"indoorMap" gorm:"comment:是否有室内地图;column:indoor_map;"`                      //是否有室内地图
	Rating       *float64       `json:"rating" form:"rating" gorm:"comment:评分;column:rating;size:2;"`                              //评分
	Cost         *float64       `json:"cost" form:"cost" gorm:"comment:消费金额;column:cost;size:10;"`                                 //消费金额
	PoiType      *string        `json:"poiType" form:"poiType" gorm:"comment:POI类型编码;column:poiType;size:10;"`                     //POI类型编码
	BigCategory  *string        `json:"BigCategory" form:"BigCategory" gorm:"comment:大类别;column:BigCategory;size:50;"`             //大类别
	MidCategory  *string        `json:"MidCategory" form:"MidCategory" gorm:"comment:中类别;column:MidCategory;size:50;"`             //中类别
	SubCategory  *string        `json:"SubCategory" form:"SubCategory" gorm:"comment:子类别;column:SubCategory;size:50;"`             //子类别
}

// TableName poiItems表 PoiItems自定义表名 poi_items
func (PoiItems) TableName() string {
	return "poi_items"
}
