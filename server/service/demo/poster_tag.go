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
		// 为海报添加标签
		if err = tx.Model(&poster).Association("Tags").Append(tags); err != nil {
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
		// 从海报中移除标签
		if err = tx.Model(&poster).Association("Tags").Delete(tags); err != nil {
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

	// 查询海报的所有标签
	if err = global.GVA_DB.Model(&poster).Association("Tags").Find(&tags); err != nil {
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

	// 查询标签关联的所有海报
	if err = global.GVA_DB.Model(&tag).Association("Posters").Find(&posters); err != nil {
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
	db := global.GVA_DB.Model(&demo.PosterTag{})

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
