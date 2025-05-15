package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// PosterTagSearch 海报标签关联查询结构体
type PosterTagSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	PosterID       *uint      `json:"posterId" form:"posterId"` // 海报ID
	TagID          *uint      `json:"tagId" form:"tagId"`       // 标签ID
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
