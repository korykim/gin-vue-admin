package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Tag 标签模型
type Tag struct {
	global.GVA_MODEL
	Name    string   `json:"name" form:"name" gorm:"comment:标签名称;size:50"`         // 标签名称
	Color   string   `json:"color" form:"color" gorm:"comment:标签颜色;size:20"`       // 标签颜色
	Posters []Poster `json:"posters" form:"posters" gorm:"many2many:poster_tags;"` // 关联海报
}
