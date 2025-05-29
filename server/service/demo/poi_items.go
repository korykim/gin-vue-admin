package demo

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
	demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
	"gorm.io/gorm/clause"
)

type PoiItemsService struct{}

// CreatePoiItems 创建poiItems表记录
// Author [yourname](https://github.com/yourname)
func (poiItemsService *PoiItemsService) CreatePoiItems(ctx context.Context, poiItems *demo.PoiItems) (err error) {
	err = global.GVA_DB.Create(poiItems).Error
	return err
}

// CreatePoiItemsInBatches 批量创建poiItems表记录，处理冲突情况
// Author [yourname](https://github.com/yourname)
func (poiItemsService *PoiItemsService) CreatePoiItemsInBatches(ctx context.Context, poiItems []*demo.PoiItems) (err error) {
	// 使用OnConflict处理冲突情况，根据id字段判断冲突
	// DoNothing表示冲突时不做任何操作
	err = global.GVA_DB.Clauses(clause.OnConflict{
		// 假设id是主键或唯一索引字段
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).CreateInBatches(poiItems, 100).Error
	return err
}

// DeletePoiItems 删除poiItems表记录
// Author [yourname](https://github.com/yourname)
func (poiItemsService *PoiItemsService) DeletePoiItems(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&demo.PoiItems{}, "id = ?", id).Error
	return err
}

// DeletePoiItemsByIds 批量删除poiItems表记录
// Author [yourname](https://github.com/yourname)
func (poiItemsService *PoiItemsService) DeletePoiItemsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]demo.PoiItems{}, "id in ?", ids).Error
	return err
}

// UpdatePoiItems 更新poiItems表记录
// Author [yourname](https://github.com/yourname)
func (poiItemsService *PoiItemsService) UpdatePoiItems(ctx context.Context, poiItems demo.PoiItems) (err error) {
	err = global.GVA_DB.Model(&demo.PoiItems{}).Where("id = ?", poiItems.Id).Updates(&poiItems).Error
	return err
}

// GetPoiItems 根据id获取poiItems表记录
// Author [yourname](https://github.com/yourname)
func (poiItemsService *PoiItemsService) GetPoiItems(ctx context.Context, id string) (poiItems demo.PoiItems, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&poiItems).Error
	return
}

// GetPoiItemsInfoList 分页获取poiItems表记录
// Author [yourname](https://github.com/yourname)
func (poiItemsService *PoiItemsService) GetPoiItemsInfoList(ctx context.Context, info demoReq.PoiItemsSearch) (list []demo.PoiItems, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&demo.PoiItems{})
	var poiItemss []demo.PoiItems
	// 如果有条件搜索 下方会自动创建搜索语句

	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	if info.Address != nil && *info.Address != "" {
		db = db.Where("address LIKE ?", "%"+*info.Address+"%")
	}
	if info.Tel != nil && *info.Tel != "" {
		db = db.Where("tel LIKE ?", "%"+*info.Tel+"%")
	}
	if info.Pname != nil && *info.Pname != "" {
		db = db.Where("pname = ?", *info.Pname)
	}
	if info.Cityname != nil && *info.Cityname != "" {
		db = db.Where("cityname = ?", *info.Cityname)
	}
	if info.Adname != nil && *info.Adname != "" {
		db = db.Where("adname = ?", *info.Adname)
	}
	if info.PoiType != nil && *info.PoiType != "" {
		db = db.Where("poiType = ?", *info.PoiType)
	}

	if info.BigCategory != nil && *info.BigCategory != "" {
		db = db.Where("BigCategory = ?", *info.BigCategory)
	}
	if info.MidCategory != nil && *info.MidCategory != "" {
		db = db.Where("MidCategory = ?", *info.MidCategory)
	}
	if info.SubCategory != nil && *info.SubCategory != "" {
		db = db.Where("SubCategory = ?", *info.SubCategory)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["pname"] = true
	orderMap["cityname"] = true
	orderMap["adname"] = true
	orderMap["BigCategory"] = true
	orderMap["MidCategory"] = true
	orderMap["SubCategory"] = true
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

	err = db.Find(&poiItemss).Error
	return poiItemss, total, err
}
func (poiItemsService *PoiItemsService) GetPoiItemsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
