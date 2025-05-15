// 自动生成模板Hero
package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 英雄信息 结构体  Hero
type Hero struct {
	global.GVA_MODEL
	Name   *string `json:"name" form:"name" gorm:"column:name;" binding:"required"` //名字
	Age    *int    `json:"age" form:"age" gorm:"column:age;"`                       //年龄
	Power  *string `json:"power" form:"power" gorm:"column:power;"`                 //能力
	Origin *string `json:"origin" form:"origin" gorm:"column:origin;"`              //起源
	Image  *string `json:"image" form:"image" gorm:"column:image;"`                 //图片

	//HeroSkills []HeroSkill `json:"heroSkills" form:"heroSkills" gorm:"many2many:hero_skills;"`
}

// TableName 英雄信息 Hero自定义表名 demo_hero
func (Hero) TableName() string {
	return "demo_hero"
}
