
// 自动生成模板ToyStore
package demo
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 玩具店 结构体  ToyStore
type ToyStore struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:玩具店名称;column:name;" binding:"required"`  //店名
  Address  *string `json:"address" form:"address" gorm:"comment:玩具店地址;column:address;" binding:"required"`  //地址
  Phone  *string `json:"phone" form:"phone" gorm:"comment:玩具店联系电话;column:phone;"`  //电话
  OpenHours  *string `json:"openHours" form:"openHours" gorm:"comment:玩具店营业时间;column:open_hours;"`  //营业时间
  Owner  *string `json:"owner" form:"owner" gorm:"comment:玩具店店主姓名;column:owner;"`  //店主
  Description  *string `json:"description" form:"description" gorm:"comment:玩具店描述;column:description;"`  //描述
  Logo  string `json:"logo" form:"logo" gorm:"comment:玩具店LOGO图片;column:logo;"`  //LOGO
  Manager  datatypes.JSON `json:"manager" form:"manager" gorm:"comment:管理员;column:manager;" swaggertype:"array,object"`  //管理员
}


// TableName 玩具店 ToyStore自定义表名 demo_ToyStore
func (ToyStore) TableName() string {
    return "demo_ToyStore"
}





