// 自动生成模板ProductImportRecords
package kotra

import (
	"time"
)

// 未准入境产品表 结构体  ProductImportRecords
type ProductImportRecords struct {
	Id                  *int       `json:"id" form:"id" gorm:"primarykey;comment:自增主键;column:id;size:10;"`                                            //自增主键
	HsCode              *string    `json:"hsCode" form:"hsCode" gorm:"comment:HS编码;column:hs_code;size:20;"`                                          //HS编码
	SingleNumber        *string    `json:"singleNumber" form:"singleNumber" gorm:"comment:检验检疫编号;column:single_number;size:100;"`                     //检验检疫编号
	ProductName         *string    `json:"productName" form:"productName" gorm:"comment:产品名称;column:product_name;"`                                   //产品名称
	PlaceOfOrigin       *string    `json:"placeOfOrigin" form:"placeOfOrigin" gorm:"comment:原产地;column:place_of_origin;size:50;"`                     //原产地
	FactoryInformation  *string    `json:"factoryInformation" form:"factoryInformation" gorm:"type:text;comment:生产企业信息;column:factory_information;"`  //生产企业信息
	ImporterInformation *string    `json:"importerInformation" form:"importerInformation" gorm:"comment:进口商信息;column:importer_information;size:200;"` //进口商信息
	RecordNumber        *string    `json:"recordNumber" form:"recordNumber" gorm:"comment:进口商备案号;column:record_number;size:100;"`                     //进口商备案号
	Weight              *string    `json:"weight" form:"weight" gorm:"comment:重量(kg);column:weight;size:200;"`                                        //重量(kg)
	IssueFacts          *string    `json:"issueFacts" form:"issueFacts" gorm:"type:text;comment:未准入境的事实;column:issue_facts;"`                         //未准入境的事实
	ImportPort          *string    `json:"importPort" form:"importPort" gorm:"comment:进口口岸;column:import_port;size:50;"`                              //进口口岸
	Years               *int       `json:"years" form:"years" gorm:"comment:年份;column:years;size:10;"`                                                //年份
	Month               *int       `json:"month" form:"month" gorm:"comment:月份;column:month;"`                                                        //月份
	CreateTime          *time.Time `json:"createTime" form:"createTime" gorm:"comment:创建时间;column:create_time;"`                                      //创建时间
	UpdateTime          *time.Time `json:"updateTime" form:"updateTime" gorm:"comment:更新时间;column:update_time;"`                                      //更新时间
}

// TableName 未准入境产品表 ProductImportRecords自定义表名 product_import_records
func (ProductImportRecords) TableName() string {
	return "product_import_records"
}
