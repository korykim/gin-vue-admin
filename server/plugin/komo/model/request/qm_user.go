package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type QmUserSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type QmUser struct {
	Username  *string `json:"username" form:"username" gorm:"column:username;" binding:"required"` //用户名
	Password  *string `json:"password" form:"password" gorm:"column:password;" binding:"required"` //密码
	Nickname  *string `json:"nickname" form:"nickname" gorm:"column:nickname;"`                    //昵称
	Captcha   string  `json:"captcha"`                                                             // 验证码
	CaptchaId string  `json:"captchaId"`                                                           // 验证码ID
}

// QmLogin 用户登录结构体
type QmLogin struct {
	Username  *string `json:"username" form:"username" binding:"required"` // 用户名
	Password  *string `json:"password" form:"password" binding:"required"` // 密码
	Captcha   string  `json:"captcha"`                                     // 验证码
	CaptchaId string  `json:"captchaId"`                                   // 验证码ID
}

type AdminChangePasswordReq struct {
	UserID   uint   `json:"userID" form:"userID"`
	Password string `json:"password" form:"password"`
}
