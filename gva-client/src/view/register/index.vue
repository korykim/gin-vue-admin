<template>
  <div class="register-container">
    <div class="register-box">
      <div class="register-logo">
        <img src="/public/logo.png" alt="Logo">
        <h2>Gin-Vue-Admin</h2>
      </div>
      <el-form ref="registerFormRef" :model="registerForm" :rules="registerRules" class="register-form">
        <el-form-item prop="username">
          <el-input 
            v-model="registerForm.username" 
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
            v-model="registerForm.password" 
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
        <el-form-item prop="confirmPassword">
          <el-input 
            v-model="registerForm.confirmPassword" 
            type="password" 
            placeholder="确认密码" 
            show-password
            clearable
          >
            <template #prefix>
              <el-icon><Lock /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item prop="nickname">
          <el-input 
            v-model="registerForm.nickname" 
            placeholder="昵称"
            clearable
          >
            <template #prefix>
              <el-icon><UserFilled /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="loading" class="register-button" @click="submitForm">注册</el-button>
        </el-form-item>
        <div class="register-options">
          <el-button link @click="goToLogin">已有账号？去登录</el-button>
        </div>
      </el-form>
    </div>
    <div class="register-background">
      <div class="background-shape"></div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import { ElMessage } from 'element-plus'
import { User, Lock, UserFilled } from '@element-plus/icons-vue'

const router = useRouter()
const userStore = useUserStore()
const loading = ref(false)
const registerFormRef = ref(null)

const registerForm = reactive({
  username: '',
  password: '',
  confirmPassword: '',
  nickname: ''
})

const validatePass2 = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请再次输入密码'))
  } else if (value !== registerForm.password) {
    callback(new Error('两次输入密码不一致'))
  } else {
    callback()
  }
}

const registerRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于 6 个字符', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请再次输入密码', trigger: 'blur' },
    { validator: validatePass2, trigger: 'blur' }
  ],
  nickname: [
    { required: true, message: '请输入昵称', trigger: 'blur' }
  ]
}

const submitForm = () => {
  if (!registerFormRef.value) return
  
  registerFormRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        // 验证表单数据
        if (!registerForm.username || !registerForm.password || !registerForm.nickname) {
          ElMessage.error('请填写完整的注册信息')
          return
        }
        
        // 验证两次密码是否一致
        if (registerForm.password !== registerForm.confirmPassword) {
          ElMessage.error('两次输入的密码不一致')
          return
        }
        
        console.log('提交注册表单:', registerForm) // 调试输出
        
        const registerData = {
          username: registerForm.username,
          password: registerForm.password,
          nickname: registerForm.nickname
        }
        
        const success = await userStore.Register(registerData)
        if (success) {
          ElMessage.success('注册成功，请登录')
          router.push('/login')
        }
      } catch (error) {
        console.error('注册错误:', error)
        ElMessage.error('注册失败: ' + (error.message || '未知错误'))
      } finally {
        loading.value = false
      }
    } else {
      ElMessage.error('请填写完整的注册信息')
      return false
    }
  })
}

const goToLogin = () => {
  router.push('/login')
}
</script>

<style lang="scss" scoped>
.register-container {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f5f7fa;
  position: relative;
  overflow: hidden;
}

.register-box {
  width: 400px;
  padding: 40px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  z-index: 2;
}

.register-logo {
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

.register-form {
  margin-top: 20px;
}

.register-button {
  width: 100%;
}

.register-options {
  display: flex;
  justify-content: center;
  margin-top: 15px;
}

.register-background {
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