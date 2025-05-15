package demo

import (
	"context"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
	demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
	"gorm.io/gorm"
)

type PosterTagService struct{}

// AddTagsToPoster 为海报添加标签
// Author [yourname](https://github.com/yourname)
func (posterTagService *PosterTagService) AddTagsToPoster(ctx context.Context, posterID uint, tagIDs []uint) (err error) {
	// 查询海报是否存在
	var poster demo.Poster
	if err = global.GVA_DB.Where("id = ?", posterID).First(&poster).Error; err != nil {
		return errors.New(fmt.Sprintf("海报不存在: %v", err))
	}

	// 查询所有标签是否存在
	var count int64
	if err = global.GVA_DB.Model(&demo.Tag{}).Where("id IN ?", tagIDs).Count(&count).Error; err != nil {
		return err
	}
	if int(count) != len(tagIDs) {
		return errors.New("部分标签不存在")
	}

	// 查询要添加的标签
	var tags []demo.Tag
	if err = global.GVA_DB.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
		return err
	}

	// 开启事务
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 为海报添加标签，使用Unscoped()避免软删除影响
		if err = tx.Model(&poster).Association("Tags").Unscoped().Append(tags); err != nil {
			return err
		}
		return nil
	})
}

// RemoveTagsFromPoster 从海报中移除标签
// Author [yourname](https://github.com/yourname)
func (posterTagService *PosterTagService) RemoveTagsFromPoster(ctx context.Context, posterID uint, tagIDs []uint) (err error) {
	// 查询海报是否存在
	var poster demo.Poster
	if err = global.GVA_DB.Where("id = ?", posterID).First(&poster).Error; err != nil {
		return errors.New(fmt.Sprintf("海报不存在: %v", err))
	}

	// 查询要移除的标签
	var tags []demo.Tag
	if err = global.GVA_DB.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
		return err
	}

	// 开启事务
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 从海报中移除标签，使用Unscoped()避免软删除影响
		if err = tx.Model(&poster).Association("Tags").Unscoped().Delete(tags); err != nil {
			return err
		}
		return nil
	})
}

// GetPosterTags 获取海报的所有标签
// Author [yourname](https://github.com/yourname)
func (posterTagService *PosterTagService) GetPosterTags(ctx context.Context, posterID uint) (tags []demo.Tag, err error) {
	// 查询海报是否存在
	var poster demo.Poster
	if err = global.GVA_DB.Where("id = ?", posterID).First(&poster).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("海报不存在: %v", err))
	}

	// 查询海报的所有标签，使用Unscoped()避免软删除影响
	if err = global.GVA_DB.Model(&poster).Association("Tags").Unscoped().Find(&tags); err != nil {
		return nil, err
	}
	return tags, nil
}

// GetTagPosters 获取标签关联的所有海报
// Author [yourname](https://github.com/yourname)
func (posterTagService *PosterTagService) GetTagPosters(ctx context.Context, tagID uint) (posters []demo.Poster, err error) {
	// 查询标签是否存在
	var tag demo.Tag
	if err = global.GVA_DB.Where("id = ?", tagID).First(&tag).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("标签不存在: %v", err))
	}

	// 查询标签关联的所有海报，使用Unscoped()避免软删除影响
	if err = global.GVA_DB.Model(&tag).Association("Posters").Unscoped().Find(&posters); err != nil {
		return nil, err
	}
	return posters, nil
}

// GetPosterTagRelations 分页获取海报标签关联
// Author [yourname](https://github.com/yourname)
func (posterTagService *PosterTagService) GetPosterTagRelations(ctx context.Context, info demoReq.PosterTagSearch) (list []demo.PosterTag, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&demo.PosterTag{}).Unscoped()

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.PosterID != nil {
		db = db.Where("poster_id = ?", info.PosterID)
	}

	if info.TagID != nil {
		db = db.Where("tag_id = ?", info.TagID)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 设置排序
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["ID"] = true
	orderMap["CreatedAt"] = true
	orderMap["PosterID"] = true
	orderMap["TagID"] = true

	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	// 分页查询
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	var posterTags []demo.PosterTag
	err = db.Find(&posterTags).Error
	return posterTags, total, err
}

// GetTagsUsageStats 获取标签使用统计
// Author [yourname](https://github.com/yourname)
func (posterTagService *PosterTagService) GetTagsUsageStats(ctx context.Context) ([]map[string]interface{}, error) {
	var tagStats []struct {
		TagID uint `json:"tag_id"`
		Count int  `json:"count"`
	}

	// 通过关联表count每个标签的使用次数，使用Unscoped()禁用软删除条件
	if err := global.GVA_DB.Table("poster_tags").Unscoped().
		Select("tag_id, count(poster_id) as count").
		Group("tag_id").
		Find(&tagStats).Error; err != nil {
		return nil, err
	}

	// 获取所有标签信息
	var tags []demo.Tag
	if err := global.GVA_DB.Find(&tags).Error; err != nil {
		return nil, err
	}

	// 合并标签信息和使用统计
	result := make([]map[string]interface{}, 0, len(tags))
	tagStatsMap := make(map[uint]int)

	// 将标签使用次数放入map便于查找
	for _, stat := range tagStats {
		tagStatsMap[stat.TagID] = stat.Count
	}

	// 为每个标签创建结果对象
	for _, tag := range tags {
		count := tagStatsMap[tag.ID] // 如果没有找到，默认为0
		tagInfo := map[string]interface{}{
			"ID":         tag.ID,
			"name":       tag.Name,
			"color":      tag.Color,
			"usageCount": count,
		}
		result = append(result, tagInfo)
	}

	return result, nil
}
