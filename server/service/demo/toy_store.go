
package demo

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
    demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
)

type ToyStoreService struct {}
// CreateToyStore 创建玩具店记录
// Author [yourname](https://github.com/yourname)
func (TSService *ToyStoreService) CreateToyStore(ctx context.Context, TS *demo.ToyStore) (err error) {
	err = global.GVA_DB.Create(TS).Error
	return err
}

// DeleteToyStore 删除玩具店记录
// Author [yourname](https://github.com/yourname)
func (TSService *ToyStoreService)DeleteToyStore(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&demo.ToyStore{},"id = ?",ID).Error
	return err
}

// DeleteToyStoreByIds 批量删除玩具店记录
// Author [yourname](https://github.com/yourname)
func (TSService *ToyStoreService)DeleteToyStoreByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]demo.ToyStore{},"id in ?",IDs).Error
	return err
}

// UpdateToyStore 更新玩具店记录
// Author [yourname](https://github.com/yourname)
func (TSService *ToyStoreService)UpdateToyStore(ctx context.Context, TS demo.ToyStore) (err error) {
	err = global.GVA_DB.Model(&demo.ToyStore{}).Where("id = ?",TS.ID).Updates(&TS).Error
	return err
}

// GetToyStore 根据ID获取玩具店记录
// Author [yourname](https://github.com/yourname)
func (TSService *ToyStoreService)GetToyStore(ctx context.Context, ID string) (TS demo.ToyStore, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&TS).Error
	return
}
// GetToyStoreInfoList 分页获取玩具店记录
// Author [yourname](https://github.com/yourname)
func (TSService *ToyStoreService)GetToyStoreInfoList(ctx context.Context, info demoReq.ToyStoreSearch) (list []demo.ToyStore, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&demo.ToyStore{})
    var TSs []demo.ToyStore
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
           orderMap["ID"] = true
           orderMap["CreatedAt"] = true
         	orderMap["name"] = true
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

	err = db.Find(&TSs).Error
	return  TSs, total, err
}
func (TSService *ToyStoreService)GetToyStoreDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	
	   manager := make([]map[string]any, 0)
	   
       
       global.GVA_DB.Table("demo_ToyShopEmployees").Where("deleted_at IS NULL").Select("name as label,id as value").Scan(&manager)
	   res["manager"] = manager
	return
}
func (TSService *ToyStoreService)GetToyStorePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
