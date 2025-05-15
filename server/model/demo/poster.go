package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Poster 海报模型
type Poster struct {
	global.GVA_MODEL
	Title string `json:"title" form:"title" gorm:"comment:海报标题;size:255"`   // 海报标题
	Image string `json:"image" form:"image" gorm:"comment:海报图片地址;size:255"` // 海报图片
	Tags  []Tag  `json:"tags" form:"tags" gorm:"many2many:poster_tags;"`    // 关联标签
}
