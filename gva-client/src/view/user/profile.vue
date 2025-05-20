<template>
  <div class="profile-container">
    <el-card class="profile-card">
      <template #header>
        <div class="card-header">
          <h3>个人资料</h3>
          <el-button type="primary" size="small" @click="refreshUserInfo">刷新</el-button>
        </div>
      </template>
      <div v-if="!loading" class="profile-content">
        <div class="avatar-container">
          <el-avatar :size="100" :src="userInfo.avatar" />
          <h2 class="username">{{ userInfo.nickname || userInfo.username }}</h2>
        </div>
        <el-divider />
        <div class="info-container">
          <el-descriptions :column="1" border>
            <el-descriptions-item label="用户名">{{ userInfo.username }}</el-descriptions-item>
            <el-descriptions-item label="昵称">{{ userInfo.nickname }}</el-descriptions-item>
            <el-descriptions-item label="性别">{{ userInfo.gender }}</el-descriptions-item>
            <el-descriptions-item label="UUID">{{ userInfo.uuid }}</el-descriptions-item>
            <el-descriptions-item label="创建时间">{{ formatDate(userInfo.CreatedAt) }}</el-descriptions-item>
            <el-descriptions-item label="更新时间">{{ formatDate(userInfo.UpdatedAt) }}</el-descriptions-item>
            <el-descriptions-item label="个人介绍">{{ userInfo.introduction }}</el-descriptions-item>
          </el-descriptions>
        </div>
      </div>
      <div v-else class="loading-container">
        <el-skeleton :rows="6" animated />
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
import { ElMessage } from 'element-plus'

const userStore = useUserStore()
const loading = ref(true)
const userInfo = ref({})

// 格式化日期
const formatDate = (dateString) => {
  if (!dateString) return '未知'
  const date = new Date(dateString)
  return date.toLocaleString()
}

// 获取用户信息
const getUserInfo = async () => {
  loading.value = true
  try {
    const data = await userStore.GetUserInfo()
    if (data) {
      userInfo.value = data
    }
  } catch (error) {
    ElMessage.error('获取用户信息失败')
  } finally {
    loading.value = false
  }
}

// 刷新用户信息
const refreshUserInfo = () => {
  getUserInfo()
}

onMounted(() => {
  getUserInfo()
})
</script>

<style lang="scss" scoped>
.profile-container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.profile-card {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.profile-content {
  padding: 20px 0;
}

.avatar-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 20px;
}

.username {
  margin-top: 15px;
  font-size: 20px;
  color: #303133;
}

.info-container {
  margin-top: 20px;
}

.loading-container {
  padding: 20px;
}
</style>