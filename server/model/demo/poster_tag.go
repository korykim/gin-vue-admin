package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PosterTag 海报与标签多对多关系模型
type PosterTag struct {
	global.GVA_MODEL
	PosterID uint `json:"poster_id" gorm:"comment:海报ID;index"` // 海报ID
	TagID    uint `json:"tag_id" gorm:"comment:标签ID;index"`    // 标签ID
}
