
// 自动生成模板Student
package demo
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 学生 结构体  Student
type Student struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:姓名;column:name;" binding:"required"`  //名字
  Age  *int `json:"age" form:"age" gorm:"comment:年龄;column:age;"`  //年龄
  Sex  *string `json:"sex" form:"sex" gorm:"comment:性别;column:sex;"`  //性别
  Profile  datatypes.JSON `json:"profile" form:"profile" gorm:"comment:简介;column:profile;" swaggertype:"array,object"`  //简介
  Picture  string `json:"picture" form:"picture" gorm:"comment:头像;column:picture;"`  //头像
}


// TableName 学生 Student自定义表名 demo_student
func (Student) TableName() string {
    return "demo_student"
}





