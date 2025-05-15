
package demo

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
    demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
)

type EmployeeProfileService struct {}
// CreateEmployeeProfile 创建员工信息记录
// Author [yourname](https://github.com/yourname)
func (TSEService *EmployeeProfileService) CreateEmployeeProfile(ctx context.Context, TSE *demo.EmployeeProfile) (err error) {
	err = global.GVA_DB.Create(TSE).Error
	return err
}

// DeleteEmployeeProfile 删除员工信息记录
// Author [yourname](https://github.com/yourname)
func (TSEService *EmployeeProfileService)DeleteEmployeeProfile(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&demo.EmployeeProfile{},"id = ?",ID).Error
	return err
}

// DeleteEmployeeProfileByIds 批量删除员工信息记录
// Author [yourname](https://github.com/yourname)
func (TSEService *EmployeeProfileService)DeleteEmployeeProfileByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]demo.EmployeeProfile{},"id in ?",IDs).Error
	return err
}

// UpdateEmployeeProfile 更新员工信息记录
// Author [yourname](https://github.com/yourname)
func (TSEService *EmployeeProfileService)UpdateEmployeeProfile(ctx context.Context, TSE demo.EmployeeProfile) (err error) {
	err = global.GVA_DB.Model(&demo.EmployeeProfile{}).Where("id = ?",TSE.ID).Updates(&TSE).Error
	return err
}

// GetEmployeeProfile 根据ID获取员工信息记录
// Author [yourname](https://github.com/yourname)
func (TSEService *EmployeeProfileService)GetEmployeeProfile(ctx context.Context, ID string) (TSE demo.EmployeeProfile, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&TSE).Error
	return
}
// GetEmployeeProfileInfoList 分页获取员工信息记录
// Author [yourname](https://github.com/yourname)
func (TSEService *EmployeeProfileService)GetEmployeeProfileInfoList(ctx context.Context, info demoReq.EmployeeProfileSearch) (list []demo.EmployeeProfile, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&demo.EmployeeProfile{})
    var TSEs []demo.EmployeeProfile
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
           orderMap["ID"] = true
           orderMap["CreatedAt"] = true
         	orderMap["name"] = true
         	orderMap["start_date"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&TSEs).Error
	return  TSEs, total, err
}
func (TSEService *EmployeeProfileService)GetEmployeeProfilePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
