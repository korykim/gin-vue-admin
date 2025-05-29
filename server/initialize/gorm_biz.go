package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/kotra"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(demo.Student{}, demo.Poster{}, demo.Tag{}, demo.PosterTag{}, demo.EmployeeProfile{}, demo.ToyStore{}, demo.PoiItems{}, kotra.ProductImportRecords{})
	if err != nil {
		return err
	}
	return nil
}
