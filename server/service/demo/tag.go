package demo

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
	demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
)

type TagService struct{}

// CreateTag 创建标签记录
func (tagService *TagService) CreateTag(ctx context.Context, tag *demo.Tag) (err error) {
	err = global.GVA_DB.Create(tag).Error
	return err
}

// DeleteTag 删除标签记录
func (tagService *TagService) DeleteTag(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&demo.Tag{}, "id = ?", ID).Error
	return err
}

// DeleteTagsByIds 批量删除标签记录
func (tagService *TagService) DeleteTagsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]demo.Tag{}, "id in ?", IDs).Error
	return err
}

// UpdateTag 更新标签记录
func (tagService *TagService) UpdateTag(ctx context.Context, tag demo.Tag) (err error) {
	err = global.GVA_DB.Model(&demo.Tag{}).Where("id = ?", tag.ID).Updates(&tag).Error
	return err
}

// GetTag 根据ID获取标签记录
func (tagService *TagService) GetTag(ctx context.Context, ID string) (tag demo.Tag, err error) {
	err = global.GVA_DB.Where("id = ?", ID).Preload("Posters").First(&tag).Error
	return
}

// GetTagInfoList 分页获取标签记录列表
func (tagService *TagService) GetTagInfoList(ctx context.Context, info demoReq.TagSearch) (list []demo.Tag, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&demo.Tag{})
	var tags []demo.Tag
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

	err = db.Preload("Posters").Find(&tags).Error
	return tags, total, err
}

// GetTagPublic 不需要鉴权的标签接口
func (tagService *TagService) GetTagPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
