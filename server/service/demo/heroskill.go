
package demo

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
    demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
)

type HeroSkillService struct {}
// CreateHeroSkill 创建技能记录
// Author [yourname](https://github.com/yourname)
func (heroskillService *HeroSkillService) CreateHeroSkill(ctx context.Context, heroskill *demo.HeroSkill) (err error) {
	err = global.GVA_DB.Create(heroskill).Error
	return err
}

// DeleteHeroSkill 删除技能记录
// Author [yourname](https://github.com/yourname)
func (heroskillService *HeroSkillService)DeleteHeroSkill(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&demo.HeroSkill{},"id = ?",ID).Error
	return err
}

// DeleteHeroSkillByIds 批量删除技能记录
// Author [yourname](https://github.com/yourname)
func (heroskillService *HeroSkillService)DeleteHeroSkillByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]demo.HeroSkill{},"id in ?",IDs).Error
	return err
}

// UpdateHeroSkill 更新技能记录
// Author [yourname](https://github.com/yourname)
func (heroskillService *HeroSkillService)UpdateHeroSkill(ctx context.Context, heroskill demo.HeroSkill) (err error) {
	err = global.GVA_DB.Model(&demo.HeroSkill{}).Where("id = ?",heroskill.ID).Updates(&heroskill).Error
	return err
}

// GetHeroSkill 根据ID获取技能记录
// Author [yourname](https://github.com/yourname)
func (heroskillService *HeroSkillService)GetHeroSkill(ctx context.Context, ID string) (heroskill demo.HeroSkill, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&heroskill).Error
	return
}
// GetHeroSkillInfoList 分页获取技能记录
// Author [yourname](https://github.com/yourname)
func (heroskillService *HeroSkillService)GetHeroSkillInfoList(ctx context.Context, info demoReq.HeroSkillSearch) (list []demo.HeroSkill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&demo.HeroSkill{})
    var heroskills []demo.HeroSkill
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    
    if info.SkillName != nil && *info.SkillName != "" {
        db = db.Where("skill_name LIKE ?", "%"+ *info.SkillName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
           orderMap["ID"] = true
           orderMap["CreatedAt"] = true
         	orderMap["skill_name"] = true
         	orderMap["skill_level"] = true
         	orderMap["skill_cooldown"] = true
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

	err = db.Find(&heroskills).Error
	return  heroskills, total, err
}
func (heroskillService *HeroSkillService)GetHeroSkillPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
