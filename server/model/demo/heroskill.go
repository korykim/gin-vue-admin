// 自动生成模板HeroSkill
package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 技能 结构体  HeroSkill
type HeroSkill struct {
	global.GVA_MODEL
	SkillName        *string  `json:"skillName" form:"skillName" gorm:"column:skill_name;" binding:"required"`              //技能名
	SkillDescription *string  `json:"skillDescription" form:"skillDescription" gorm:"column:skill_description;"type:text;"` //描述
	SkillLevel       *int     `json:"skillLevel" form:"skillLevel" gorm:"column:skill_level;" binding:"required"`           //等级
	SkillEffect      *string  `json:"skillEffect" form:"skillEffect" gorm:"column:skill_effect;"`                           //效果
	SkillCooldown    *float64 `json:"skillCooldown" form:"skillCooldown" gorm:"column:skill_cooldown;" binding:"required"`  //冷却
	//Heros            []Hero   `json:"heros" form:"heros" gorm:"many2many:hero_skills;"`
}

// TableName 技能 HeroSkill自定义表名 demo_HeroSkill
func (HeroSkill) TableName() string {
	return "demo_HeroSkill"
}
