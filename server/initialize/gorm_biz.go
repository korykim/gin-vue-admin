package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(demo.Student{})
	if err != nil {
		return err
	}
	return nil
}
