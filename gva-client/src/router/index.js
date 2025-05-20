import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import Cookies from 'js-cookie'

const router = createRouter({
  history: createWebHistory(""),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import("@/view/home/index.vue"),
      meta: {
        title: '首页',
        requiresAuth: false
      }
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import("@/view/dashboard/index.vue"),
      meta: {
        title: '控制台',
        requiresAuth: true
      }
    },
    {
      path: '/login',
      name: 'login',
      component: () => import("@/view/login/index.vue"),
      meta: {
        title: '登录'
      }
    },
    {
      path: '/register',
      name: 'register',
      component: () => import("@/view/register/index.vue"),
      meta: {
        title: '注册'
      }
    },
    {
      path: '/profile',
      name: 'profile',
      component: () => import("@/view/user/profile.vue"),
      meta: {
        title: '个人资料',
        requiresAuth: true
      }
    }
  ]
})

// 全局前置守卫
router.beforeEach(async (to, from, next) => {
  // 设置页面标题
  document.title = to.meta.title ? `${to.meta.title} - Gin-Vue-Admin` : 'Gin-Vue-Admin'
  
  // 检查该路由是否需要登录权限
  if (to.matched.some(record => record.meta.requiresAuth)) {
    const userStore = useUserStore()
    const token = userStore.token || Cookies.get('x-token') || localStorage.getItem('x-token')
    
    // 如果没有token，重定向到登录页面
    if (!token) {
      console.log('没有找到token，重定向到登录页面')
      next({
        path: '/login',
        query: { redirect: to.fullPath } // 将目标路由的完整路径传递给登录页面
      })
      return
    }
    
    // 如果有token但没有用户信息，尝试获取用户信息
    if (token && (!userStore.userInfo.username || Object.keys(userStore.userInfo).length === 0)) {
      try {
        console.log('尝试获取用户信息')
        const userData = await userStore.GetUserInfo()
        if (!userData) {
          console.log('获取用户信息失败，可能token已过期')
          userStore.ClearStorage() // 清除无效token
          next({
            path: '/login',
            query: { redirect: to.fullPath }
          })
          return
        }
      } catch (error) {
        console.error('获取用户信息出错:', error)
        userStore.ClearStorage()
        next({
          path: '/login',
          query: { redirect: to.fullPath }
        })
        return
      }
    }
    
    // 有token且已获取用户信息，允许访问
    next()
  } else if (to.path === '/login' || to.path === '/register') {
    // 如果用户已登录且尝试访问登录或注册页面，重定向到控制台
    const userStore = useUserStore()
    const token = userStore.token || Cookies.get('x-token') || localStorage.getItem('x-token')
    
    if (token) {
      // 检查token是否有效
      try {
        const userData = await userStore.GetUserInfo()
        if (userData) {
          console.log('用户已登录，重定向到控制台')
          next({ path: '/dashboard' })
          return
        }
      } catch (error) {
        console.error('检查登录状态出错:', error)
        // token无效，清除并继续访问登录页
        userStore.ClearStorage()
      }
    }
    next()
  } else {
    next() // 确保一定要调用 next()
  }
})

export default router
