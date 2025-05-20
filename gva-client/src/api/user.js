import service from '@/utils/request'
// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data) => {
  return service({
    url: '/base/login',
    method: 'post',
    data: data
  })
}

// @Summary 获取验证码
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/captcha [post]
export const captcha = () => {
  return service({
    url: '/base/captcha',
    method: 'post'
  })
}

// @Summary 用户注册
// @Produce  application/json
// @Param data body {username:"string",password:"string",nickname:"string"}
// @Router /qmUser/register [post]
export const register = (data) => {
  return service({
    url: '/qmUser/register',
    method: 'post',
    data: data
  })
}

// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /qmUser/login [post]
export const qmLogin = (data) => {
  return service({
    url: '/qmUser/login',
    method: 'post',
    data: data
  })
}

// @Summary 获取用户信息
// @Produce  application/json
// @Router /qmUser/GetUserInfo [get]
export const getUserInfo = () => {
  return service({
    url: '/qmUser/GetUserInfo',
    method: 'get'
  })
}

// @Summary 用户登出
// @Produce  application/json
// @Router /qmUser/logout [get]
export const logout = () => {
  return service({
    url: '/qmUser/logout',  
    method: 'get'
  })
}


