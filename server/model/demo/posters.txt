
// 自动生成模板Posters
package demo
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// posters表 结构体  Posters
type Posters struct {
    global.GVA_MODEL
  Title  *string `json:"title" form:"title" gorm:"comment:海报标题;column:title;size:255;" binding:"required"`  //海报标题
  Image  *string `json:"image" form:"image" gorm:"comment:海报图片地址;column:image;size:255;"`  //海报图片地址
}


// TableName posters表 Posters自定义表名 posters
func (Posters) TableName() string {
    return "posters"
}





