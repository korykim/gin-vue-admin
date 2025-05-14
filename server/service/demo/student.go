
package demo

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
    demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
)

type StudentService struct {}
// CreateStudent 创建学生记录
// Author [yourname](https://github.com/yourname)
func (studentService *StudentService) CreateStudent(ctx context.Context, student *demo.Student) (err error) {
	err = global.GVA_DB.Create(student).Error
	return err
}

// DeleteStudent 删除学生记录
// Author [yourname](https://github.com/yourname)
func (studentService *StudentService)DeleteStudent(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&demo.Student{},"id = ?",ID).Error
	return err
}

// DeleteStudentByIds 批量删除学生记录
// Author [yourname](https://github.com/yourname)
func (studentService *StudentService)DeleteStudentByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]demo.Student{},"id in ?",IDs).Error
	return err
}

// UpdateStudent 更新学生记录
// Author [yourname](https://github.com/yourname)
func (studentService *StudentService)UpdateStudent(ctx context.Context, student demo.Student) (err error) {
	err = global.GVA_DB.Model(&demo.Student{}).Where("id = ?",student.ID).Updates(&student).Error
	return err
}

// GetStudent 根据ID获取学生记录
// Author [yourname](https://github.com/yourname)
func (studentService *StudentService)GetStudent(ctx context.Context, ID string) (student demo.Student, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&student).Error
	return
}
// GetStudentInfoList 分页获取学生记录
// Author [yourname](https://github.com/yourname)
func (studentService *StudentService)GetStudentInfoList(ctx context.Context, info demoReq.StudentSearch) (list []demo.Student, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&demo.Student{})
    var students []demo.Student
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&students).Error
	return  students, total, err
}
func (studentService *StudentService)GetStudentPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
