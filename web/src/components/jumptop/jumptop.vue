<template>
  <transition name="el-fade-in">
    <div 
      v-show="visible" 
      class="jump-top-container"
      :style="{ right: backTopRight, bottom: backTopBottom }"
      @click="handleClick"
    >
      <div class="jump-top-content">
        <el-icon :size="18"><CaretTop /></el-icon>
        <span v-if="showText" class="jump-top-text">{{ text }}</span>
      </div>
    </div>
  </transition>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { ElIcon } from 'element-plus'
import { CaretTop } from '@element-plus/icons-vue'

defineOptions({
  name: 'JumpTop'
})

// 定义组件属性
const props = defineProps({
  // 显示的文本
  text: {
    type: String,
    default: '顶部'
  },
  // 是否显示文本
  showText: {
    type: Boolean,
    default: true
  },
  // 按钮距离右侧的距离
  right: {
    type: Number,
    default: 40
  },
  // 按钮距离底部的距离
  bottom: {
    type: Number,
    default: 40
  },
  // 滚动高度达到此参数值才出现
  visibilityHeight: {
    type: Number,
    default: 200
  },
  // 自定义点击事件
  customClick: {
    type: Function,
    default: null
  },
  // 目标容器的选择器，默认为主内容区域
  target: {
    type: String,
    default: '#gva-base-load-dom'
  }
})

// 定义事件
const emit = defineEmits(['click'])

// 计算属性
const backTopRight = computed(() => `${props.right}px`)
const backTopBottom = computed(() => `${props.bottom}px`)

// 控制按钮显示
const visible = ref(false)

// 获取目标滚动容器
const getScrollContainer = () => {
  // 首先尝试获取指定的目标容器
  if (props.target) {
    const targetContainer = document.querySelector(props.target)
    if (targetContainer && isScrollable(targetContainer)) {
      return targetContainer
    }
  }
  
  // 然后尝试常见的布局容器
  const commonSelectors = [
    '.gva-container2', 
    '.gva-container', 
    '.overflow-auto',
    '.poi-list-container', 
    '.gva-table-box'
  ]
  
  for (const selector of commonSelectors) {
    const container = document.querySelector(selector)
    if (container && isScrollable(container)) {
      return container
    }
  }
  
  // 如果没有找到可滚动的容器，返回文档主体
  return document.documentElement || document.body
}

// 检查元素是否可滚动
const isScrollable = (element) => {
  if (!element) return false
  
  const style = window.getComputedStyle(element)
  const overflowY = style.getPropertyValue('overflow-y')
  
  return (
    overflowY === 'scroll' || 
    overflowY === 'auto' || 
    (element.scrollHeight > element.clientHeight)
  )
}

// 滚动监听函数
const handleScroll = () => {
  const container = getScrollContainer()
  if (container) {
    // 使用scrollTop来判断滚动位置
    visible.value = container.scrollTop > props.visibilityHeight
  }
}

// 点击事件处理
const handleClick = (e) => {
  const container = getScrollContainer()
  
  if (container) {
    // 平滑滚动到顶部
    container.scrollTo({
      top: 0,
      behavior: 'smooth'
    })
  }

  if (props.customClick) {
    props.customClick(e)
  }
  
  emit('click', e)
}

// 生命周期钩子
onMounted(() => {
  const container = getScrollContainer()
  
  if (container) {
    container.addEventListener('scroll', handleScroll)
    // 初始检查滚动位置
    handleScroll()
  }
})

onUnmounted(() => {
  const container = getScrollContainer()
  
  if (container) {
    container.removeEventListener('scroll', handleScroll)
  }
})
</script>

<style lang="scss" scoped>
.jump-top-container {
  position: fixed;
  cursor: pointer;
  z-index: 99;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 48px;
  height: 48px;
  border-radius: 4px;
  background-color: var(--el-color-primary);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  transition: all 0.3s;
  color: #fff;
  
  &:hover {
    background-color: var(--el-color-primary-dark-2);
  }
}

.jump-top-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 12px;
}

.jump-top-text {
  margin-top: 2px;
  line-height: 1;
}

/* 自定义过渡效果 */
.el-fade-in-enter-active,
.el-fade-in-leave-active {
  transition: opacity 0.3s, transform 0.3s;
}

.el-fade-in-enter-from,
.el-fade-in-leave-to {
  opacity: 0;
  transform: translateY(10px);
}
</style>
