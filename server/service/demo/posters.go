package demo

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
	demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
)

type PostersService struct{}

// CreatePosters 创建posters表记录
// Author [yourname](https://github.com/yourname)
func (postersService *PostersService) CreatePosters(ctx context.Context, posters *demo.Poster) (err error) {
	err = global.GVA_DB.Create(posters).Error
	return err
}

// DeletePosters 删除posters表记录
// Author [yourname](https://github.com/yourname)
func (postersService *PostersService) DeletePosters(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&demo.Poster{}, "id = ?", ID).Error
	return err
}

// DeletePostersByIds 批量删除posters表记录
// Author [yourname](https://github.com/yourname)
func (postersService *PostersService) DeletePostersByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]demo.Poster{}, "id in ?", IDs).Error
	return err
}

// UpdatePosters 更新posters表记录
// Author [yourname](https://github.com/yourname)
func (postersService *PostersService) UpdatePosters(ctx context.Context, posters demo.Poster) (err error) {
	err = global.GVA_DB.Model(&demo.Poster{}).Where("id = ?", posters.ID).Updates(&posters).Error
	return err
}

// GetPosters 根据ID获取posters表记录
// Author [yourname](https://github.com/yourname)
func (postersService *PostersService) GetPosters(ctx context.Context, ID string) (posters demo.Poster, err error) {
	err = global.GVA_DB.Where("id = ?", ID).Preload("Tags").First(&posters).Error
	return
}

// GetPostersInfoList 分页获取posters表记录
// Author [yourname](https://github.com/yourname)
func (postersService *PostersService) GetPostersInfoList(ctx context.Context, info demoReq.PostersSearch) (list []demo.Poster, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&demo.Poster{})
	var posterss []demo.Poster
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

	err = db.Preload("Tags").Find(&posterss).Error
	return posterss, total, err
}
func (postersService *PostersService) GetPostersPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
