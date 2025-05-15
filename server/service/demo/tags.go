package demo

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
	demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
)

type TagsService struct{}

// CreateTags 创建tags表记录
// Author [yourname](https://github.com/yourname)
func (tagsService *TagsService) CreateTags(ctx context.Context, tags *demo.Tag) (err error) {
	err = global.GVA_DB.Create(tags).Error
	return err
}

// DeleteTags 删除tags表记录
// Author [yourname](https://github.com/yourname)
func (tagsService *TagsService) DeleteTags(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&demo.Tag{}, "id = ?", ID).Error
	return err
}

// DeleteTagsByIds 批量删除tags表记录
// Author [yourname](https://github.com/yourname)
func (tagsService *TagsService) DeleteTagsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]demo.Tag{}, "id in ?", IDs).Error
	return err
}

// UpdateTags 更新tags表记录
// Author [yourname](https://github.com/yourname)
func (tagsService *TagsService) UpdateTags(ctx context.Context, tags demo.Tag) (err error) {
	err = global.GVA_DB.Model(&demo.Tag{}).Where("id = ?", tags.ID).Updates(&tags).Error
	return err
}

// GetTags 根据ID获取tags表记录
// Author [yourname](https://github.com/yourname)
func (tagsService *TagsService) GetTags(ctx context.Context, ID string) (tags demo.Tag, err error) {
	err = global.GVA_DB.Where("id = ?", ID).Preload("Posters").First(&tags).Error
	return
}

// GetTagsInfoList 分页获取tags表记录
// Author [yourname](https://github.com/yourname)
func (tagsService *TagsService) GetTagsInfoList(ctx context.Context, info demoReq.TagsSearch) (list []demo.Tag, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&demo.Tag{})
	var tagss []demo.Tag
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Preload("Posters").Find(&tagss).Error
	return tagss, total, err
}
func (tagsService *TagsService) GetTagsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
