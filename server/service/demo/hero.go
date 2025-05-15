
package demo

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
    demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
)

type HeroService struct {}
// CreateHero 创建英雄信息记录
// Author [yourname](https://github.com/yourname)
func (heroService *HeroService) CreateHero(ctx context.Context, hero *demo.Hero) (err error) {
	err = global.GVA_DB.Create(hero).Error
	return err
}

// DeleteHero 删除英雄信息记录
// Author [yourname](https://github.com/yourname)
func (heroService *HeroService)DeleteHero(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&demo.Hero{},"id = ?",ID).Error
	return err
}

// DeleteHeroByIds 批量删除英雄信息记录
// Author [yourname](https://github.com/yourname)
func (heroService *HeroService)DeleteHeroByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]demo.Hero{},"id in ?",IDs).Error
	return err
}

// UpdateHero 更新英雄信息记录
// Author [yourname](https://github.com/yourname)
func (heroService *HeroService)UpdateHero(ctx context.Context, hero demo.Hero) (err error) {
	err = global.GVA_DB.Model(&demo.Hero{}).Where("id = ?",hero.ID).Updates(&hero).Error
	return err
}

// GetHero 根据ID获取英雄信息记录
// Author [yourname](https://github.com/yourname)
func (heroService *HeroService)GetHero(ctx context.Context, ID string) (hero demo.Hero, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&hero).Error
	return
}
// GetHeroInfoList 分页获取英雄信息记录
// Author [yourname](https://github.com/yourname)
func (heroService *HeroService)GetHeroInfoList(ctx context.Context, info demoReq.HeroSearch) (list []demo.Hero, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&demo.Hero{})
    var heros []demo.Hero
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
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
         	orderMap["age"] = true
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

	err = db.Find(&heros).Error
	return  heros, total, err
}
func (heroService *HeroService)GetHeroPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
