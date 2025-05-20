<template>
  <div class="login-container">
    <div class="login-box">
      <div class="login-logo">
        <img src="/public/logo.png" alt="Logo">
        <h2>Gin-Vue-Admin</h2>
      </div>
      <el-form ref="loginFormRef" :model="loginForm" :rules="loginRules" class="login-form">
        <el-form-item prop="username">
          <el-input 
            v-model="loginForm.username" 
            placeholder="用户名"
            clearable
          >
            <template #prefix>
              <el-icon><User /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input 
            v-model="loginForm.password" 
            type="password" 
            placeholder="密码" 
            show-password
            clearable
          >
            <template #prefix>
              <el-icon><Lock /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="loading" class="login-button" @click="submitForm">登录</el-button>
        </el-form-item>
        <div class="login-options">
          <el-button link @click="goToRegister">没有账号？去注册</el-button>
        </div>
      </el-form>
    </div>
    <div class="login-background">
      <div class="background-shape"></div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import { ElMessage } from 'element-plus'
import { User, Lock } from '@element-plus/icons-vue'
import Cookies from 'js-cookie'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const loading = ref(false)
const loginFormRef = ref(null)

const loginForm = reactive({
  username: '',
  password: ''
})

const loginRules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
}

const submitForm = () => {
  if (!loginFormRef.value) return
  
  loginFormRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        // 确保用户名和密码不为空
        if (!loginForm.username || !loginForm.password) {
          ElMessage.error('用户名和密码不能为空')
          return
        }
        
        console.log('提交登录表单:', loginForm) // 调试输出
        
        const success = await userStore.LoginIn(loginForm)
        if (success) {
          ElMessage.success('登录成功')
          
          // 检查 token 是否已保存到 cookie
          const token = Cookies.get('x-token')
          console.log('登录后的 token:', token)
          
          // 如果有重定向参数，跳转到对应页面，否则跳转到控制台
          const redirectPath = route.query.redirect || '/dashboard'
          router.push(redirectPath)
        }
      } catch (error) {
        console.error('登录错误:', error)
        ElMessage.error('登录失败: ' + (error.message || '未知错误'))
      } finally {
        loading.value = false
      }
    } else {
      ElMessage.error('请填写完整的登录信息')
      return false
    }
  })
}

const goToRegister = () => {
  router.push('/register')
}

// 测试输入框是否正常工作
onMounted(() => {
  console.log('登录页面已加载')
  
  // 检查是否已有 token
  const token = Cookies.get('x-token') || localStorage.getItem('x-token')
  console.log('已存储的 token:', token)
  
  // 如果已有 token，尝试自动登录
  if (token) {
    userStore.GetUserInfo().then(data => {
      if (data) {
        router.push('/')
      }
    })
  }
})
</script>

<style lang="scss" scoped>
.login-container {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f5f7fa;
  position: relative;
  overflow: hidden;
}

.login-box {
  width: 400px;
  padding: 40px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  z-index: 2;
}

.login-logo {
  text-align: center;
  margin-bottom: 30px;
  
  img {
    width: 80px;
    height: 80px;
  }
  
  h2 {
    margin-top: 10px;
    font-size: 24px;
    color: #409EFF;
  }
}

.login-form {
  margin-top: 20px;
}

.login-button {
  width: 100%;
}

.login-options {
  display: flex;
  justify-content: center;
  margin-top: 15px;
}

.login-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 1;
}

.background-shape {
  position: absolute;
  top: -50px;
  right: -50px;
  width: 300px;
  height: 300px;
  border-radius: 50%;
  background: linear-gradient(45deg, #409EFF, #36D1DC);
  opacity: 0.2;
}

.background-shape::before {
  content: '';
  position: absolute;
  bottom: -80px;
  left: -120px;
  width: 400px;
  height: 400px;
  border-radius: 50%;
  background: linear-gradient(45deg, #36D1DC, #5B86E5);
  opacity: 0.3;
}
</style> 