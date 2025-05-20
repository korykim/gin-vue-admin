package config

type Config struct {
	AuthorityId uint    `mapstructure:"authority-id" json:"authority-id" yaml:"authority-id"`
	Captcha     Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}

type Captcha struct {
	OpenCaptcha        int `mapstructure:"open-captcha" json:"open-captcha" yaml:"open-captcha"`                         // 防爆破验证码开启此数，0代表每次登录都需要验证码，其他数字代表错误密码此数，如3代表错误三次后出现验证码
	OpenCaptchaTimeOut int `mapstructure:"open-captcha-timeout" json:"open-captcha-timeout" yaml:"open-captcha-timeout"` // 防爆破验证码超时时间，单位：s(秒)
}
