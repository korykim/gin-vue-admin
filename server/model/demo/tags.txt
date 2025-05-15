
// 自动生成模板Tags
package demo
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// tags表 结构体  Tags
type Tags struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:标签名称;column:name;size:50;" binding:"required"`  //标签名称
  Color  *string `json:"color" form:"color" gorm:"comment:标签颜色;column:color;size:20;"`  //标签颜色
}


// TableName tags表 Tags自定义表名 tags
func (Tags) TableName() string {
    return "tags"
}





