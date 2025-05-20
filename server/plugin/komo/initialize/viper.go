package initialize

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/komo/plugin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Viper() {
	err := global.GVA_VP.UnmarshalKey("komo", &plugin.Config)
	if err != nil {
		err = errors.Wrap(err, "初始化配置文件失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}

	// 如果没有配置验证码，则使用默认配置
	if plugin.Config.Captcha.OpenCaptcha == 0 && plugin.Config.Captcha.OpenCaptchaTimeOut == 0 {
		plugin.Config.Captcha.OpenCaptcha = 3         // 默认错误三次后显示验证码
		plugin.Config.Captcha.OpenCaptchaTimeOut = 60 // 默认超时时间60秒
	}
}
