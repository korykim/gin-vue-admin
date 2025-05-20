<template>
    <div class="relative isolate h-screen overflow-hidden box-border">
      <div class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80" aria-hidden="true">
        <div class="relative left-[calc(50%-11rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-gradient-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]" style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"></div>
      </div>
      
      <!-- 顶部导航栏 -->
      <div class="p-4 flex justify-between items-center">
        <div class="text-center w-16 h-16"><img class="w-full h-full" src="/public/logo.png"></div>
        <div class="flex items-center gap-4">
          <!-- 已登录状态 -->
          <template v-if="isLoggedIn">
            <div class="flex items-center gap-2">
              <el-avatar :size="32" :src="userInfo.avatar || '/public/default-avatar.png'" />
              <span>{{ userInfo.nickname || userInfo.username }}</span>
            </div>
            <el-dropdown>
              <el-button type="primary">
                菜单 <el-icon class="el-icon--right"><arrow-down /></el-icon>
              </el-button>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item @click="goToDashboard">控制台</el-dropdown-item>
                  <el-dropdown-item @click="goToProfile">个人资料</el-dropdown-item>
                  <el-dropdown-item divided @click="handleLogout">退出登录</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </template>
          
          <!-- 未登录状态 -->
          <template v-else>
            <el-button type="primary" @click="goToLogin">登录</el-button>
            <el-button @click="goToRegister">注册</el-button>
          </template>
        </div>
      </div>
      
      <div class="h-full flex flex-col items-center justify-center -mt-16">
        <div class="text-center w-32 h-32 mb-8"><img class="w-full h-full" src="/public/logo.png"> </div>
        
        <div class="text-center">
          <h1 class="text-4xl font-bold tracking-tight text-gray-900 sm:text-6xl">更快实现你的创意</h1>
          <p class="mt-6 text-lg leading-8 text-gray-600">使用Gin-Vue-Admin,让你用更少的时间，更稳定的架构，做更多的项目，写更优雅的代码。</p>
          
          <!-- 未登录用户的行动按钮 -->
          <div v-if="!isLoggedIn" class="mt-10 flex items-center justify-center gap-x-6">
            <el-button type="primary" size="large" @click="goToLogin">立即登录</el-button>
            <el-button size="large" @click="goToRegister">免费注册</el-button>
          </div>
          
          <!-- 已登录用户的欢迎信息 -->
          <div v-else class="mt-8 p-4 bg-white rounded-lg shadow-sm">
            <h2 class="text-2xl font-bold text-gray-800">欢迎回来，{{ userInfo.nickname || userInfo.username }}</h2>
            <p class="mt-2 text-gray-600">您已成功登录系统</p>
            <div class="mt-4">
              <el-button type="primary" @click="goToDashboard">进入控制台</el-button>
            </div>
          </div>
          
          <!-- 功能特点 -->
          <div class="mt-16">
            <h2 class="text-2xl font-bold mb-6">系统特点</h2>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div class="feature-card">
                <el-icon class="feature-icon"><document /></el-icon>
                <h3 class="feature-title">代码生成器</h3>
                <p class="feature-desc">自动生成标准化的CRUD代码，提高开发效率</p>
              </div>
              <div class="feature-card">
                <el-icon class="feature-icon"><lock /></el-icon>
                <h3 class="feature-title">权限管理</h3>
                <p class="feature-desc">灵活的RBAC权限控制系统</p>
              </div>
              <div class="feature-card">
                <el-icon class="feature-icon"><setting /></el-icon>
                <h3 class="feature-title">插件系统</h3>
                <p class="feature-desc">可扩展的插件架构，轻松添加新功能</p>
              </div>
            </div>
          </div>
        </div>
      </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { ElButton, ElAvatar, ElDropdown, ElDropdownMenu, ElDropdownItem, ElMessage } from "element-plus";
import { ArrowDown, Document, Lock, Setting } from '@element-plus/icons-vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '@/pinia/modules/user';
import Cookies from 'js-cookie';

const router = useRouter();
const userStore = useUserStore();
const userInfo = ref({});
const isLoggedIn = computed(() => {
  return userInfo.value && userInfo.value.username;
});

// 获取用户信息（如果已登录）
const getUserInfo = async () => {
  try {
    // 检查是否有token
    const token = userStore.token || Cookies.get('x-token') || localStorage.getItem('x-token');
    if (!token) {
      console.log('未登录状态');
      return;
    }
    
    console.log('首页获取用户信息');
    
    // 如果store中没有用户信息，尝试获取
    if (!userStore.userInfo.username || Object.keys(userStore.userInfo).length === 0) {
      const data = await userStore.GetUserInfo();
      if (data) {
        userInfo.value = data;
        console.log('获取到用户信息:', data);
      }
    } else {
      // 使用store中的用户信息
      userInfo.value = userStore.userInfo;
      console.log('使用store中的用户信息:', userInfo.value);
    }
  } catch (error) {
    console.error('获取用户信息失败', error);
  }
};

// 跳转到登录页面
const goToLogin = () => {
  router.push('/login');
};

// 跳转到注册页面
const goToRegister = () => {
  router.push('/register');
};

// 跳转到控制台
const goToDashboard = () => {
  router.push('/dashboard');
};

// 跳转到个人资料页面
const goToProfile = () => {
  router.push('/profile');
};

// 退出登录
const handleLogout = async () => {
  try {
    await userStore.Logout();
    ElMessage.success('退出登录成功');
    userInfo.value = {};
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


<style scoped  lang="scss">
.feature-card {
  padding: 20px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  text-align: center;
  transition: transform 0.3s;
  
  &:hover {
    transform: translateY(-5px);
  }
  
  .feature-icon {
    font-size: 32px;
    color: #409EFF;
    margin-bottom: 16px;
  }
  
  .feature-title {
    font-size: 18px;
    font-weight: bold;
    margin-bottom: 8px;
    color: #303133;
  }
  
  .feature-desc {
    color: #606266;
    font-size: 14px;
  }
}
</style>
