<template>
  <div class="amap-page">
    <div class="amap-container">
      <MapContainer />
    </div>
  </div>
</template>

<script setup>
import MapContainer from '@/components/amappoi/MapContainer.vue'
import { defineOptions } from 'vue'

defineOptions({
    name: 'amap'
})
</script>

<style scoped>
.amap-page {
  @apply bg-white dark:bg-slate-800 rounded-md p-4 h-full w-full;
}

.amap-container {
  @apply h-[calc(100vh-120px)] w-full relative;
}
</style>

<style>
/* 提高图片预览模态框的z-index */
.el-image-viewer__wrapper {
  z-index: 3000 !important; /* 确保高于地图的z-index */
}

/* 确保遮罩层也有足够高的z-index */
.el-image-viewer__mask {
  z-index: 2999 !important;
}
</style>
