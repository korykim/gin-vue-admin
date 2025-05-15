package demo

import (
	"time"

	"gorm.io/gorm"
)

// PosterTag 海报与标签多对多关系模型
type PosterTag struct {
	PosterID  uint           `json:"poster_id" gorm:"primaryKey;comment:海报ID"` // 海报ID
	TagID     uint           `json:"tag_id" gorm:"primaryKey;comment:标签ID"`    // 标签ID
	CreatedAt time.Time      `json:"created_at" gorm:"comment:创建时间"`           // 创建时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"-"`                               // 禁用软删除功能
}

// TableName 设置表名
func (PosterTag) TableName() string {
	return "poster_tags"
}
