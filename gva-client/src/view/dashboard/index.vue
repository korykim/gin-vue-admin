<template>
  <div class="dashboard-container">
    <!-- 顶部导航栏 -->
    <div class="top-navbar">
      <div class="logo">
        <img src="/public/logo.png" alt="Logo">
        <h2>Gin-Vue-Admin</h2>
      </div>
      <div class="user-info">
        <div v-if="userInfo.username" class="user-avatar">
          <el-avatar :size="40" :src="userInfo.avatar || '/public/default-avatar.png'" />
          <span>{{ userInfo.nickname || userInfo.username }}</span>
        </div>
        <el-dropdown>
          <el-button type="primary" size="small">
            操作 <el-icon class="el-icon--right"><arrow-down /></el-icon>
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="goToProfile">个人资料</el-dropdown-item>
              <el-dropdown-item @click="goToHome">返回首页</el-dropdown-item>
              <el-dropdown-item divided @click="handleLogout">退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>

    <!-- 主要内容区域 -->
    <div class="main-content">
      <div class="welcome-section">
        <h1>欢迎来到控制台，{{ userInfo.nickname || userInfo.username }}！</h1>
        <p>这里是您的个人控制中心，您可以管理您的账户和使用系统功能。</p>
      </div>
      
      <!-- 功能卡片区域 -->
      <div class="feature-cards">
        <el-row :gutter="20">
          <el-col :span="8">
            <el-card shadow="hover" @click="goToProfile">
              <template #header>
                <div class="card-header">
                  <el-icon><user /></el-icon>
                  <span>个人资料</span>
                </div>
              </template>
              <div class="card-content">
                查看和编辑您的个人信息
              </div>
            </el-card>
          </el-col>
          
          <el-col :span="8">
            <el-card shadow="hover">
              <template #header>
                <div class="card-header">
                  <el-icon><document /></el-icon>
                  <span>我的项目</span>
                </div>
              </template>
              <div class="card-content">
                管理您的项目和任务
              </div>
            </el-card>
          </el-col>
          
          <el-col :span="8">
            <el-card shadow="hover">
              <template #header>
                <div class="card-header">
                  <el-icon><setting /></el-icon>
                  <span>系统设置</span>
                </div>
              </template>
              <div class="card-content">
                配置系统参数和选项
              </div>
            </el-card>
          </el-col>
        </el-row>
      </div>
      
      <!-- 统计信息 -->
      <div class="statistics-section">
        <h2>系统概览</h2>
        <el-row :gutter="20">
          <el-col :span="6">
            <div class="stat-card">
              <div class="stat-value">12</div>
              <div class="stat-label">项目总数</div>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="stat-card">
              <div class="stat-value">45</div>
              <div class="stat-label">任务总数</div>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="stat-card">
              <div class="stat-value">8</div>
              <div class="stat-label">待处理任务</div>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="stat-card">
              <div class="stat-value">32</div>
              <div class="stat-label">已完成任务</div>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { ElButton, ElAvatar, ElDropdown, ElDropdownMenu, ElDropdownItem, ElMessage, ElRow, ElCol, ElCard } from 'element-plus';
import { ArrowDown, User, Document, Setting } from '@element-plus/icons-vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '@/pinia/modules/user';
import Cookies from 'js-cookie';

const router = useRouter();
const userStore = useUserStore();
const userInfo = ref({});

// 获取用户信息
const getUserInfo = async () => {
  try {
    console.log('控制台获取用户信息');
    
    // 如果store中没有用户信息，尝试获取
    if (!userStore.userInfo.username || Object.keys(userStore.userInfo).length === 0) {
      const data = await userStore.GetUserInfo();
      if (data) {
        userInfo.value = data;
        console.log('获取到用户信息:', data);
      } else {
        console.log('获取用户信息失败');
        checkLoginStatus();
      }
    } else {
      // 使用store中的用户信息
      userInfo.value = userStore.userInfo;
      console.log('使用store中的用户信息:', userInfo.value);
    }
  } catch (error) {
    console.error('获取用户信息失败', error);
    checkLoginStatus();
  }
};

// 检查登录状态
const checkLoginStatus = () => {
  const token = userStore.token || Cookies.get('x-token') || localStorage.getItem('x-token');
  if (!token) {
    console.log('未找到token，重定向到登录页');
    ElMessage.warning('请先登录');
    router.push('/login');
  }
};

// 跳转到个人资料页面
const goToProfile = () => {
  router.push('/profile');
};

// 跳转到首页
const goToHome = () => {
  router.push('/');
};

// 退出登录
const handleLogout = async () => {
  try {
    await userStore.Logout();
    ElMessage.success('退出登录成功');
    router.push('/login');
  } catch (error) {
    console.error('退出登录失败', error);
    ElMessage.error('退出登录失败');
  }
};

onMounted(() => {
  getUserInfo();
});
</script>

<style scoped lang="scss">
.dashboard-container {
  min-height: 100vh;
  background-color: #f5f7fa;
}

.top-navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 20px;
  background-color: white;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  
  .logo {
    display: flex;
    align-items: center;
    
    img {
      width: 40px;
      height: 40px;
      margin-right: 10px;
    }
    
    h2 {
      font-size: 18px;
      color: #409EFF;
      margin: 0;
    }
  }
  
  .user-info {
    display: flex;
    align-items: center;
    gap: 15px;
    
    .user-avatar {
      display: flex;
      align-items: center;
      gap: 8px;
      
      span {
        font-size: 14px;
      }
    }
  }
}

.main-content {
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.welcome-section {
  margin-bottom: 30px;
  padding: 30px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  text-align: center;
  
  h1 {
    font-size: 24px;
    color: #303133;
    margin-bottom: 10px;
  }
  
  p {
    color: #606266;
    font-size: 16px;
  }
}

.feature-cards {
  margin-bottom: 30px;
  
  .card-header {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 16px;
    font-weight: bold;
  }
  
  .card-content {
    padding: 10px 0;
    color: #606266;
  }
  
  .el-card {
    cursor: pointer;
    transition: transform 0.3s;
    
    &:hover {
      transform: translateY(-5px);
    }
  }
}

.statistics-section {
  padding: 20px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  
  h2 {
    font-size: 18px;
    margin-bottom: 20px;
    color: #303133;
  }
  
  .stat-card {
    padding: 20px;
    text-align: center;
    background-color: #f5f7fa;
    border-radius: 4px;
    
    .stat-value {
      font-size: 24px;
      font-weight: bold;
      color: #409EFF;
    }
    
    .stat-label {
      margin-top: 8px;
      color: #606266;
    }
  }
}
</style> 