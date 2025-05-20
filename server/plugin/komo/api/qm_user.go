package api

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/komo/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/komo/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/captcha"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var QmUser = new(qmUser)

type qmUser struct{}

// CreateQmUser 创建用户信息
// @Tags QmUser
// @Summary 创建用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.QmUser true "创建用户信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /qmUser/createQmUser [post]
func (a *qmUser) CreateQmUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.QmUser
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceQmUser.CreateQmUser(ctx, &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteQmUser 删除用户信息
// @Tags QmUser
// @Summary 删除用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.QmUser true "删除用户信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /qmUser/deleteQmUser [delete]
func (a *qmUser) DeleteQmUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceQmUser.DeleteQmUser(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteQmUserByIds 批量删除用户信息
// @Tags QmUser
// @Summary 批量删除用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /qmUser/deleteQmUserByIds [delete]
func (a *qmUser) DeleteQmUserByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceQmUser.DeleteQmUserByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateQmUser 更新用户信息
// @Tags QmUser
// @Summary 更新用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.QmUser true "更新用户信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /qmUser/updateQmUser [put]
func (a *qmUser) UpdateQmUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.QmUser
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceQmUser.UpdateQmUser(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindQmUser 用id查询用户信息
// @Tags QmUser
// @Summary 用id查询用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询用户信息"
// @Success 200 {object} response.Response{data=model.QmUser,msg=string} "查询成功"
// @Router /qmUser/findQmUser [get]
func (a *qmUser) FindQmUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reqmUser, err := serviceQmUser.GetQmUser(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reqmUser, c)
}

// GetQmUserList 分页获取用户信息列表
// @Tags QmUser
// @Summary 分页获取用户信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.QmUserSearch true "分页获取用户信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /qmUser/getQmUserList [get]
func (a *qmUser) GetQmUserList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.QmUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceQmUser.GetQmUserInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetQmUserPublic 不需要鉴权的用户信息接口
// @Tags QmUser
// @Summary 不需要鉴权的用户信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /qmUser/getQmUserPublic [get]
func (a *qmUser) GetQmUserPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceQmUser.GetQmUserPublic(ctx)
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的用户信息接口信息"}, "获取成功", c)
}

// AdminChangePassword AdminChangePassword
// @Tags QmUser
// @Summary AdminChangePassword
// @Accept application/json
// @Produce application/json
// @Param data body request.AdminChangePasswordReq true "AdminChangePassword"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /qmUser/adminChangePassword [PUT]
func (a *qmUser) AdminChangePassword(c *gin.Context) {

	var req request.AdminChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 创建业务用Context
	ctx := c.Request.Context()
	// 请添加自己的业务逻辑
	err = serviceQmUser.AdminChangePassword(ctx, req)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}

// Register 用户注册
// @Tags QmUser
// @Summary 用户注册
// @Accept application/json
// @Produce application/json
// @Param body body request.QmUser true "注册信息"
// @Success 200 {object} response.Response{data=object,msg=string} "注册成功"
// @Router /qmUser/register [POST]
func (a *qmUser) Register(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var registerReq request.QmUser
	err := c.ShouldBindJSON(&registerReq)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证必要的参数
	if registerReq.Username == nil || registerReq.Password == nil {
		response.FailWithMessage("用户名和密码不能为空", c)
		return
	}

	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	// 判断是否需要验证码
	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	// 验证码校验
	store := captcha.NewDefaultRedisStore()

	// 调试信息
	global.GVA_LOG.Info("注册尝试", zap.String("username", *registerReq.Username))

	if registerReq.CaptchaId != "" && registerReq.Captcha != "" {
		global.GVA_LOG.Info("验证码信息", zap.String("captchaId", registerReq.CaptchaId), zap.String("captcha", registerReq.Captcha))
	}

	// 验证码逻辑：如果需要验证码，则必须验证通过；如果不需要验证码，则直接通过
	if !oc || (registerReq.CaptchaId != "" && registerReq.Captcha != "" && store.Verify(registerReq.CaptchaId, registerReq.Captcha, true)) {
		// 创建用户模型
		var user model.QmUser
		user.Username = registerReq.Username
		user.Password = registerReq.Password

		// 昵称可选，如果提供则使用
		if registerReq.Nickname != nil {
			user.Nickname = registerReq.Nickname
		}

		// 调用service方法创建用户
		err = serviceQmUser.CreateQmUser(ctx, &user)
		if err != nil {
			global.GVA_LOG.Error("注册失败", zap.Error(err))
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("注册失败: "+err.Error(), c)
			return
		}

		global.GVA_LOG.Info("用户注册成功", zap.Uint("userId", user.ID), zap.String("username", *user.Username))
		response.OkWithMessage("注册成功", c)
		return
	}

	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", c)
}

// 用户登录API
// @Tags QmUser
// @Summary 用户登录
// @Accept application/json
// @Produce application/json
// @Param body body request.QmLogin true "用户登录请求参数"
// @Success 200 {object} response.Response{data=model.QmUser,msg=string} "登录成功"
// @Failure 400 {object} response.Response{data=object,msg=string} "登录失败"
// @Router /qmUser/login [POST]
func (a *qmUser) Login(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var loginReq request.QmLogin
	err := c.ShouldBindJSON(&loginReq)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(loginReq, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 确保用户名和密码不为空
	if loginReq.Username == nil || loginReq.Password == nil {
		response.FailWithMessage("用户名和密码不能为空", c)
		return
	}

	err = utils.Verify(loginReq, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	// 判断是否需要验证码
	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	// 验证码校验
	store := captcha.NewDefaultRedisStore()

	// 调试信息
	//global.GVA_LOG.Info("Login attempt", zap.String("username", *loginReq.Username))
	//global.GVA_LOG.Info("Captcha info", zap.String("captchaId", loginReq.CaptchaId), zap.String("captcha", loginReq.Captcha))

	// 修复验证码逻辑：如果需要验证码，则必须验证通过；如果不需要验证码，则直接通过
	if !oc || (loginReq.CaptchaId != "" && loginReq.Captcha != "" && store.Verify(loginReq.CaptchaId, loginReq.Captcha, true)) {
		user, err := serviceQmUser.Login(ctx, *loginReq.Username, *loginReq.Password)
		if err != nil {
			global.GVA_LOG.Error("登录失败! 用户名不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}

		if user.Enable != 1 {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户被禁止登录", c)
			return
		}

		// 记录用户ID用于调试
		//global.GVA_LOG.Info("用户登录成功", zap.Uint("userId", user.ID))

		token, claims, err := utils.LoginToken(&user)
		if err != nil {
			global.GVA_LOG.Error("生成token失败", zap.Error(err))
			response.FailWithMessage("登录失败: "+err.Error(), c)
			return
		}

		// 记录token信息用于调试
		// global.GVA_LOG.Info("Token生成成功",
		// 	zap.Uint("userId", claims.BaseClaims.ID),
		// 	zap.String("username", claims.Username))

		maxAge := int(claims.RegisteredClaims.ExpiresAt.Unix() - time.Now().Unix())
		utils.SetToken(c, token, maxAge)
		response.OkWithDetailed(gin.H{
			"user":      user,
			"token":     token,
			"expiresAt": claims.RegisteredClaims.ExpiresAt,
		}, "登录成功", c)
		return
	}

	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", c)
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}

// 获取用户信息
// @Tags QmUser
// @Summary 获取用户信息
// @Accept application/json
// @Produce application/json
// @Param data body map[string]interface{} true "用户信息参数"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取成功"
// @Router /qmUser/GetUserInfo [GET]
func (a *qmUser) GetUserInfo(c *gin.Context) {

	id := utils.GetUserID(c)
	if id == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	// 创建业务用Context
	ctx := c.Request.Context()
	// 调用server方法接收参数并获取返回参数
	user, err := serviceQmUser.GetUserInfo(ctx, id)
	if err != nil {
		// 检查是否是"记录未找到"错误
		if err.Error() == "record not found" {
			response.FailWithMessage("用户不存在，请重新登录", c)
			return
		}
		global.GVA_LOG.Error("获取用户信息失败", zap.Error(err))
		response.FailWithMessage("获取用户信息失败", c)
		return
	}
	response.OkWithData(user, c)
}

// 用户登出
// @Tags QmUser
// @Summary 用户登出
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "登出成功"
// @Router /qmUser/logout [GET]
func (a *qmUser) Logout(c *gin.Context) {

	utils.ClearToken(c)

	response.OkWithMessage("登出成功", c)
}
