
// 自动生成模板EmployeeProfile
package demo
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 员工信息 结构体  EmployeeProfile
type EmployeeProfile struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`  //姓名
  Position  *string `json:"position" form:"position" gorm:"column:position;"`  //职位
  Phone  *string `json:"phone" form:"phone" gorm:"column:phone;"`  //电话
  Email  *string `json:"email" form:"email" gorm:"column:email;"`  //邮箱
  StartDate  *time.Time `json:"startDate" form:"startDate" gorm:"column:start_date;"`  //入职日期
}


// TableName 员工信息 EmployeeProfile自定义表名 demo_ToyShopEmployees
func (EmployeeProfile) TableName() string {
    return "demo_ToyShopEmployees"
}





