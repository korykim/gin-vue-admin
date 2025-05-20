package service

import (
	"context"
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/komo/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/komo/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/komo/plugin"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid/v5"
	"go.uber.org/zap"
)

var QmUser = new(qmUser)

type qmUser struct{}

// CreateQmUser 创建用户信息记录
// Author [yourname](https://github.com/yourname)
func (s *qmUser) CreateQmUser(ctx context.Context, qmUser *model.QmUser) (err error) {

	// 查询username是否已经注册
	var count int64
	err = global.GVA_DB.Model(&model.QmUser{}).Where("username = ?", qmUser.Username).Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("用户名已被注册")
	}

	pwd := *qmUser.Password
	qmUser.Password = new(string)
	*qmUser.Password = utils.BcryptHash(pwd)
	qmUser.UUID, _ = uuid.NewV4()

	qmUser.Introduction = new(string)
	*qmUser.Introduction = "这位用户很懒，什么都没有留下"

	// 设置默认性别，避免数据库错误
	if qmUser.Gender == "" {
		qmUser.Gender = "男"
	}

	// 设置用户权限ID
	qmUser.AuthorityId = plugin.Config.AuthorityId

	err = global.GVA_DB.Create(qmUser).Error
	return err
}

// DeleteQmUser 删除用户信息记录
// Author [yourname](https://github.com/yourname)
func (s *qmUser) DeleteQmUser(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.QmUser{}, "id = ?", ID).Error
	return err
}

// DeleteQmUserByIds 批量删除用户信息记录
// Author [yourname](https://github.com/yourname)
func (s *qmUser) DeleteQmUserByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.QmUser{}, "id in ?", IDs).Error
	return err
}

// UpdateQmUser 更新用户信息记录
// Author [yourname](https://github.com/yourname)
func (s *qmUser) UpdateQmUser(ctx context.Context, qmUser model.QmUser) (err error) {
	// 创建一个新的结构体，只包含允许更新的字段
	updateData := map[string]interface{}{
		"avatar":       qmUser.Avatar,
		"nickname":     qmUser.Nickname,
		"gender":       qmUser.Gender,
		"introduction": qmUser.Introduction,
		"updated_at":   qmUser.UpdatedAt,
	}

	err = global.GVA_DB.Model(&model.QmUser{}).Where("id = ?", qmUser.ID).Updates(updateData).Error
	return err
}

// GetQmUser 根据ID获取用户信息记录
// Author [yourname](https://github.com/yourname)
func (s *qmUser) GetQmUser(ctx context.Context, ID string) (qmUser model.QmUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&qmUser).Error
	return
}

// GetQmUserInfoList 分页获取用户信息记录
// Author [yourname](https://github.com/yourname)
func (s *qmUser) GetQmUserInfoList(ctx context.Context, info request.QmUserSearch) (list []model.QmUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.QmUser{})
	var qmUsers []model.QmUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&qmUsers).Error
	return qmUsers, total, err
}

func (s *qmUser) GetQmUserPublic(ctx context.Context) {
	// 实现公共用户功能
}

// AdminChangePassword AdminChangePassword
// Author [yourname](https://github.com/yourname)
func (s *qmUser) AdminChangePassword(ctx context.Context, req request.AdminChangePasswordReq) (err error) {
	newPassword := utils.BcryptHash(req.Password)
	db := global.GVA_DB.Model(&model.QmUser{}).Where("id = ?", req.UserID).Update("password", newPassword)
	return db.Error
}

// Register 注册用户
func (s *qmUser) Register(ctx context.Context, body map[string]interface{}) (err error) {
	// 获取用户名和密码
	usernameStr := body["username"].(string)
	passwordStr := body["password"].(string)

	// 创建用户模型
	username := new(string)
	*username = usernameStr

	password := new(string)
	*password = passwordStr

	// 设置权限ID
	authorityId := plugin.Config.AuthorityId

	user := model.QmUser{
		Username:    username,
		Password:    password,
		AuthorityId: authorityId,
	}

	// 保存用户
	err = global.GVA_DB.Create(&user).Error
	return err
}

// Login 用户登录
func (s *qmUser) Login(ctx context.Context, username, password string) (user model.QmUser, err error) {
	// 根据用户名查找用户
	err = global.GVA_DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	// 检查用户名是否为空
	if user.Username == nil || *user.Username == "" {
		return user, errors.New("用户名或密码错误")
	}

	// 验证密码
	if !utils.BcryptCheck(password, *user.Password) {
		return user, errors.New("用户名或密码错误")
	}

	return
}

// GetUserInfo 获取用户信息
func (s *qmUser) GetUserInfo(ctx context.Context, userId uint) (user model.QmUser, err error) {
	// 记录正在查询的用户ID
	global.GVA_LOG.Info("正在查询用户信息",
		zap.Uint("userId", userId),
		zap.String("表名", model.QmUser{}.TableName()))

	// 根据ID查询用户
	err = global.GVA_DB.First(&user, "id = ?", userId).Error

	// 记录查询结果
	if err != nil {
		global.GVA_LOG.Error("查询用户信息失败",
			zap.Error(err),
			zap.Uint("userId", userId))
	} else {
		global.GVA_LOG.Info("查询用户信息成功",
			zap.Uint("userId", userId),
			zap.String("username", *user.Username))
	}

	return
}
