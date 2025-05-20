import { defineStore } from 'pinia'
import { ref } from "vue"
import { qmLogin, getUserInfo, register, logout } from '@/api/user'
import { ElMessage } from 'element-plus'
import Cookies from 'js-cookie'

export const useUserStore = defineStore('user', () => {
  // 从 cookie 或 localStorage 中恢复 token
  const storedToken = Cookies.get('x-token') || localStorage.getItem('x-token')
  
  const userInfo = ref({})
  const token = ref(storedToken || '')

  // 设置用户信息
  const setUserInfo = (val) => {
    userInfo.value = val
    // 可选：将用户信息存储到 localStorage
    localStorage.setItem('userInfo', JSON.stringify(val))
  }

  // 设置token
  const setToken = (val) => {
    token.value = val
    // 将 token 同时存储到 cookie 和 localStorage
    Cookies.set('x-token', val, { expires: 7 }) // 7天过期
    localStorage.setItem('x-token', val)
  }

  // 登录
  const LoginIn = async (loginInfo) => {
    try {
      // 验证登录信息
      if (!loginInfo.username || !loginInfo.password) {
        ElMessage.error('用户名和密码不能为空')
        return false
      }

      console.log('正在发送登录请求:', loginInfo)
      const res = await qmLogin(loginInfo)
      console.log('登录响应:', res)
      
      if (res.code === 0) {
        setUserInfo(res.data.user)
        setToken(res.data.token)
        return true
      } else {
        ElMessage.error(res.message || '登录失败')
        return false
      }
    } catch (err) {
      console.error('登录错误:', err)
      ElMessage.error(err.message || '登录失败')
      return false
    }
  }

  // 注册
  const Register = async (registerInfo) => {
    try {
      const res = await register(registerInfo)
      if (res.code === 0) {
        ElMessage.success('注册成功')
        return true
      } else {
        ElMessage.error(res.message || '注册失败')
        return false
      }
    } catch (err) {
      console.error('注册错误:', err)
      ElMessage.error(err.message || '注册失败')
      return false
    }
  }

  // 获取用户信息
  const GetUserInfo = async () => {
    try {
      // 如果没有 token，直接返回
      if (!token.value) {
        return null
      }
      
      const res = await getUserInfo()
      if (res.code === 0) {
        setUserInfo(res.data)
        return res.data
      } else {
        ElMessage.error(res.message || '获取用户信息失败')
        return null
      }
    } catch (err) {
      console.error('获取用户信息错误:', err)
      ElMessage.error(err.message || '获取用户信息失败')
      return null
    }
  }

  // 登出
  const Logout = async () => {
    try {
      // 调用后端登出接口
      const res = await logout()
      console.log('登出响应:', res)
      
      // 无论后端响应如何，都清除本地存储
      ClearStorage()
      
      if (res.code === 0) {
        return true
      } else {
        ElMessage.error(res.message || '登出失败')
        return false
      }
    } catch (err) {
      console.error('登出错误:', err)
      // 即使后端请求失败，也清除本地存储
      ClearStorage()
      throw err
    }
  }

  // 清除存储
  const ClearStorage = () => {
    userInfo.value = {}
    token.value = ''
    // 清除 cookie 和 localStorage 中的数据
    Cookies.remove('x-token')
    localStorage.removeItem('x-token')
    localStorage.removeItem('userInfo')
  }

  return {
    userInfo,
    token,
    setUserInfo,
    setToken,
    LoginIn,
    Register,
    GetUserInfo,
    Logout,
    ClearStorage
  }
})
