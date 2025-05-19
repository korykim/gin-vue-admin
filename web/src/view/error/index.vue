<template>
  <div class="error-page">
    <div class="error-container">
      <h1>404</h1>
      <h2>页面不存在</h2>
      <p>抱歉，您访问的页面不存在或已被移除</p>
      <el-button type="primary" @click="goHome">返回首页</el-button>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'

const router = useRouter()
const userStore = useUserStore()

const goHome = () => {
  // 检查用户是否有默认路由，如果有则跳转到默认路由，否则跳转到登录页
  if (userStore.userInfo?.authority?.defaultRouter) {
    router.push({ name: userStore.userInfo.authority.defaultRouter })
  } else {
    router.push({ path: '/login' })
  }
}
</script>

<style lang="scss" scoped>
.error-page {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f5f7fa;
  
  .error-container {
    text-align: center;
    padding: 40px;
    
    h1 {
      font-size: 80px;
      color: #409eff;
      margin-bottom: 20px;
    }
    
    h2 {
      font-size: 30px;
      color: #606266;
      margin-bottom: 20px;
    }
    
    p {
      font-size: 16px;
      color: #909399;
      margin-bottom: 30px;
    }
  }
}
</style> 