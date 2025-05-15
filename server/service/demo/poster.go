package demo

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
	demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
)

type PosterService struct{}

// CreatePoster 创建poster记录
func (posterService *PosterService) CreatePoster(ctx context.Context, poster *demo.Poster) (err error) {
	err = global.GVA_DB.Create(poster).Error
	return err
}

// DeletePoster 删除poster记录
func (posterService *PosterService) DeletePoster(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&demo.Poster{}, "id = ?", ID).Error
	return err
}

// DeletePostersByIds 批量删除poster记录
func (posterService *PosterService) DeletePostersByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]demo.Poster{}, "id in ?", IDs).Error
	return err
}

// UpdatePoster 更新poster记录
func (posterService *PosterService) UpdatePoster(ctx context.Context, poster demo.Poster) (err error) {
	err = global.GVA_DB.Model(&demo.Poster{}).Where("id = ?", poster.ID).Updates(&poster).Error
	return err
}

// GetPoster 根据ID获取poster记录
func (posterService *PosterService) GetPoster(ctx context.Context, ID string) (poster demo.Poster, err error) {
	err = global.GVA_DB.Where("id = ?", ID).Preload("Tags").First(&poster).Error
	return
}

// GetPosterInfoList 分页获取poster记录列表
func (posterService *PosterService) GetPosterInfoList(ctx context.Context, info demoReq.PosterSearch) (list []demo.Poster, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&demo.Poster{})
	var posters []demo.Poster
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.Title != nil && *info.Title != "" {
		db = db.Where("title LIKE ?", "%"+*info.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["ID"] = true
	orderMap["CreatedAt"] = true
	orderMap["title"] = true
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

	err = db.Preload("Tags").Find(&posters).Error
	return posters, total, err
}

// GetPosterPublic 不需要鉴权的poster接口
func (posterService *PosterService) GetPosterPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
