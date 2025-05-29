package kotra

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/kotra"
	kotraReq "github.com/flipped-aurora/gin-vue-admin/server/model/kotra/request"
)

type ProductImportRecordsService struct{}

// CreateProductImportRecords 创建未准入境产品表记录
// Author [yourname](https://github.com/yourname)
func (productImportRecordsService *ProductImportRecordsService) CreateProductImportRecords(ctx context.Context, productImportRecords *kotra.ProductImportRecords) (err error) {
	err = global.GVA_DB.Create(productImportRecords).Error
	return err
}

// DeleteProductImportRecords 删除未准入境产品表记录
// Author [yourname](https://github.com/yourname)
func (productImportRecordsService *ProductImportRecordsService) DeleteProductImportRecords(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&kotra.ProductImportRecords{}, "id = ?", id).Error
	return err
}

// DeleteProductImportRecordsByIds 批量删除未准入境产品表记录
// Author [yourname](https://github.com/yourname)
func (productImportRecordsService *ProductImportRecordsService) DeleteProductImportRecordsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]kotra.ProductImportRecords{}, "id in ?", ids).Error
	return err
}

// UpdateProductImportRecords 更新未准入境产品表记录
// Author [yourname](https://github.com/yourname)
func (productImportRecordsService *ProductImportRecordsService) UpdateProductImportRecords(ctx context.Context, productImportRecords kotra.ProductImportRecords) (err error) {
	err = global.GVA_DB.Model(&kotra.ProductImportRecords{}).Where("id = ?", productImportRecords.Id).Updates(&productImportRecords).Error
	return err
}

// GetProductImportRecords 根据id获取未准入境产品表记录
// Author [yourname](https://github.com/yourname)
func (productImportRecordsService *ProductImportRecordsService) GetProductImportRecords(ctx context.Context, id string) (productImportRecords kotra.ProductImportRecords, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&productImportRecords).Error
	return
}

// GetProductImportRecordsInfoList 分页获取未准入境产品表记录
// Author [yourname](https://github.com/yourname)
func (productImportRecordsService *ProductImportRecordsService) GetProductImportRecordsInfoList(ctx context.Context, info kotraReq.ProductImportRecordsSearch) (list []kotra.ProductImportRecords, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&kotra.ProductImportRecords{})
	var productImportRecordss []kotra.ProductImportRecords
	// 如果有条件搜索 下方会自动创建搜索语句

	if info.HsCode != nil && *info.HsCode != "" {
		db = db.Where("hs_code = ?", *info.HsCode)
	}
	if info.ProductName != nil && *info.ProductName != "" {
		db = db.Where("product_name LIKE ?", "%"+*info.ProductName+"%")
	}
	if info.FactoryInformation != nil && *info.FactoryInformation != "" {
		db = db.Where("factory_information LIKE ?", "%"+*info.FactoryInformation+"%")
	}
	if info.ImporterInformation != nil && *info.ImporterInformation != "" {
		db = db.Where("importer_information LIKE ?", "%"+*info.ImporterInformation+"%")
	}
	if info.RecordNumber != nil && *info.RecordNumber != "" {
		db = db.Where("record_number = ?", *info.RecordNumber)
	}
	if info.ImportPort != nil && *info.ImportPort != "" {
		db = db.Where("import_port = ?", *info.ImportPort)
	}
	if info.Years != nil && *info.Years > 0 {
		db = db.Where("years = ?", *info.Years)
	}
	if info.Month != nil && *info.Month > 0 {
		db = db.Where("month = ?", *info.Month)
	}
	if info.PlaceOfOrigin != nil && *info.PlaceOfOrigin != "" {
		db = db.Where("place_of_origin = ?", *info.PlaceOfOrigin)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["years"] = true
	orderMap["month"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&productImportRecordss).Error
	return productImportRecordss, total, err
}
func (productImportRecordsService *ProductImportRecordsService) GetProductImportRecordsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
