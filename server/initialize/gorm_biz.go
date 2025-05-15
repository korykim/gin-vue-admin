package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(demo.Student{}, demo.Hero{}, demo.HeroSkill{}, demo.Poster{}, demo.Tag{}, demo.Posters{}, demo.Tags{})
	if err != nil {
		return err
	}
	return nil
}
