package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gofrs/uuid/v5"
	googluuid "github.com/google/uuid"
)

var _ system.Login = (*QmUser)(nil)

// QmUser 用户信息 结构体
type QmUser struct {
	global.GVA_MODEL
	Username     *string   `json:"username" form:"username" gorm:"column:username;" binding:"required"`                               //用户名
	Password     *string   `json:"-" form:"password" gorm:"column:password;" binding:"required"`                                      //密码
	Avatar       string    `json:"avatar" form:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;column:avatar;"` //头像
	UUID         uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;"`                                                              //UUID
	Nickname     *string   `json:"nickname" form:"nickname" gorm:"column:nickname;"`                                                  //昵称
	Gender       string    `json:"gender" form:"gender" gorm:"column:gender;type:enum('男','女');"`                                     //性别
	Introduction *string   `json:"introduction" form:"introduction" gorm:"column:introduction;"`
	AuthorityId  uint      `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:用户角色ID"` //权限ID
	Enable       int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
}

// TableName 用户信息 QmUser自定义表名 qm_client_user
func (QmUser) TableName() string {
	return "qm_client_user"
}

// 实现Login接口的方法

// GetUsername 获取用户名
func (u *QmUser) GetUsername() string {
	if u.Username == nil {
		return ""
	}
	return *u.Username
}

// GetNickname 获取昵称
func (u *QmUser) GetNickname() string {
	if u.Nickname == nil {
		return ""
	}
	return *u.Nickname
}

// GetUUID 获取UUID
func (u *QmUser) GetUUID() googluuid.UUID {
	// 转换UUID类型
	id, _ := googluuid.Parse(u.UUID.String())
	return id
}

// GetUserId 获取用户ID
func (u *QmUser) GetUserId() uint {
	return u.ID
}

// GetAuthorityId 获取权限ID
func (u *QmUser) GetAuthorityId() uint {
	return u.AuthorityId
}

// GetUserInfo 获取用户信息
func (u *QmUser) GetUserInfo() any {
	return *u
}

// GetEnable 获取用户是否被冻结
func (u *QmUser) GetEnable() int {
	return u.Enable
}
