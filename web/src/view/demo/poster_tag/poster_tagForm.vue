<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="120px">
        <el-form-item label="操作类型:" prop="actionType">
          <el-select v-model="formData.actionType" placeholder="请选择操作类型" @change="handleActionTypeChange">
            <el-option label="添加标签" value="add" />
            <el-option label="移除标签" value="remove" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="选择海报:" prop="posterId">
          <el-select v-model="formData.posterId" filterable placeholder="请选择海报" @change="handlePosterChange">
            <el-option v-for="poster in posterOptions" :key="poster.value" :label="poster.label" :value="poster.value" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="选择标签:" prop="tagIds">
          <el-select v-model="formData.tagIds" multiple filterable placeholder="请选择标签" style="width: 100%">
            <el-option v-for="tag in tagOptions" :key="tag.value" :label="tag.label" :value="tag.value" />
          </el-select>
        </el-form-item>
        
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  addTagsToPoster,
  removeTagsFromPoster,
  getPosterTags
} from '@/api/demo/poster_tag'

import { getTagsList } from '@/api/demo/tags'
import { getPostersList } from '@/api/demo/posters'

defineOptions({
  name: 'PosterTagForm'
})

// 自动获取字典
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive, onMounted } from 'vue'

const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

// 海报选项
const posterOptions = ref([])
// 标签选项
const tagOptions = ref([])

// 表单数据
const formData = ref({
  actionType: 'add', // 默认是添加操作
  posterId: '',
  tagIds: []
})

// 验证规则
const rule = reactive({
  actionType: [
    { required: true, message: '请选择操作类型', trigger: 'change' }
  ],
  posterId: [
    { required: true, message: '请选择海报', trigger: 'change' }
  ],
  tagIds: [
    { required: true, message: '请选择标签', trigger: 'change' },
    { type: 'array', min: 1, message: '请至少选择一个标签', trigger: 'change' }
  ]
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
  await getOptions()
  
  // 如果URL有参数，则设置表单初始值
  if (route.query.posterId) {
    formData.value.posterId = Number(route.query.posterId)
    await handlePosterChange()
  }
  
  if (route.query.actionType) {
    formData.value.actionType = route.query.actionType
  }
}

// 获取选项数据
const getOptions = async () => {
  try {
    // 获取所有海报
    const postersRes = await getPostersList({ page: 1, pageSize: 100 })
    if (postersRes.code === 0) {
      posterOptions.value = postersRes.data.list.map(item => ({
        value: item.ID,
        label: `${item.title} (ID: ${item.ID})`
      }))
    }
    
    // 获取所有标签
    const tagsRes = await getTagsList({ page: 1, pageSize: 100 })
    if (tagsRes.code === 0) {
      tagOptions.value = tagsRes.data.list.map(item => ({
        value: item.ID,
        label: `${item.name} (ID: ${item.ID})`
      }))
    }
  } catch (error) {
    console.error('获取选项数据失败:', error)
  }
}

// 当选择海报变化时
const handlePosterChange = async () => {
  // 如果是移除操作，需要获取当前海报的标签
  if (formData.value.actionType === 'remove' && formData.value.posterId) {
    try {
      const res = await getPosterTags({ posterId: formData.value.posterId })
      if (res.code === 0) {
        // 更新标签选项为当前海报的标签
        formData.value.tagIds = []
        const posterTags = res.data
        tagOptions.value = posterTags.map(tag => ({
          value: tag.ID,
          label: `${tag.name} (ID: ${tag.ID})`
        }))
      }
    } catch (error) {
      console.error('获取海报标签失败:', error)
    }
  }
}

// 当操作类型变化时
const handleActionTypeChange = async () => {
  // 清空已选标签
  formData.value.tagIds = []
  
  // 如果是移除操作且已选择海报，则获取该海报的标签
  if (formData.value.actionType === 'remove' && formData.value.posterId) {
    await handlePosterChange()
  } else if (formData.value.actionType === 'add') {
    // 如果是添加操作，重新获取所有标签
    const tagsRes = await getTagsList({ page: 1, pageSize: 100 })
    if (tagsRes.code === 0) {
      tagOptions.value = tagsRes.data.list.map(item => ({
        value: item.ID,
        label: `${item.name} (ID: ${item.ID})`
      }))
    }
  }
}

onMounted(() => {
  init()
})

// 保存按钮
const save = async () => {
  btnLoading.value = true
  elFormRef.value?.validate(async (valid) => {
    if (!valid) {
      btnLoading.value = false
      return
    }
    
    try {
      let res
      const params = {
        posterId: formData.value.posterId,
        tagIds: formData.value.tagIds
      }
      
      if (formData.value.actionType === 'add') {
        res = await addTagsToPoster(params)
      } else {
        res = await removeTagsFromPoster(params)
      }
      
      btnLoading.value = false
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: formData.value.actionType === 'add' ? '添加标签成功' : '移除标签成功'
        })
        // 操作成功后返回列表页
        back()
      }
    } catch (error) {
      console.error('操作失败:', error)
      btnLoading.value = false
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}
</script>

<style>
</style> 