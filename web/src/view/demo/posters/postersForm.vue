<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="海报标题:" prop="title">
    <el-input v-model="formData.title" :clearable="true" placeholder="请输入海报标题" />
</el-form-item>
        <el-form-item label="海报图片地址:" prop="image">
    <el-input v-model="formData.image" :clearable="true" placeholder="请输入海报图片地址" />
</el-form-item>
        <el-form-item label="标签管理:">
          <div class="tags-container">
            <div class="current-tags" v-if="posterTags.length > 0">
              <el-tag
                v-for="tag in posterTags"
                :key="tag.ID"
                :type="tag.color"
                class="tag-item"
                closable
                @close="removeTag(tag)"
              >
                {{ tag.name }}
              </el-tag>
            </div>
            <div class="no-tags" v-else>
              <p class="empty-tip">尚未添加标签</p>
            </div>
            
            <div class="add-tags">
              <el-select
                v-model="selectedTagIds"
                multiple
                filterable
                placeholder="请选择要添加的标签"
                style="width: 100%;"
                @change="handleTagsChange"
              >
                <el-option
                  v-for="tag in availableTags"
                  :key="tag.ID"
                  :label="tag.name"
                  :value="tag.ID"
                >
                  <el-tag :type="tag.color">{{ tag.name }}</el-tag>
                </el-option>
              </el-select>
            </div>
          </div>
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
  createPosters,
  updatePosters,
  findPosters
} from '@/api/demo/posters'

import { 
  addTagsToPoster, 
  removeTagsFromPoster, 
  getPosterTags 
} from '@/api/demo/poster_tag'

import { 
  getTagsList 
} from '@/api/demo/tags'

defineOptions({
    name: 'PostersForm'
})

// 引入所需的库
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive, onMounted } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

// 标签相关
const posterTags = ref([])
const availableTags = ref([])
const selectedTagIds = ref([])
const tagsToAdd = ref([])
const tagsToRemove = ref([])

const type = ref('')
const formData = ref({
            title: '',
            image: '',
        })
// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPosters({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()

// 加载海报标签
const loadPosterTags = async () => {
  if (route.query.id) {
    try {
      const res = await getPosterTags({ posterId: route.query.id })
      if (res.code === 0) {
        posterTags.value = res.data || []
      }
    } catch (error) {
      console.error('获取海报标签失败:', error)
    }
  } else {
    // 新建海报时，没有标签
    posterTags.value = []
  }
}

// 加载可用标签
const loadAvailableTags = async () => {
  try {
    const res = await getTagsList({ page: 1, pageSize: 100 })
    if (res.code === 0) {
      // 过滤掉已经添加的标签
      const currentTagIds = posterTags.value.map(tag => tag.ID)
      availableTags.value = res.data.list.filter(tag => !currentTagIds.includes(tag.ID))
    }
  } catch (error) {
    console.error('获取标签列表失败:', error)
  }
}

// 处理标签选择变化
const handleTagsChange = (value) => {
  // 计算新添加的标签
  const newTags = value.filter(id => !posterTags.value.some(tag => tag.ID === id))
  
  if (newTags.length > 0) {
    // 将新标签添加到posterTags中
    newTags.forEach(tagId => {
      const tag = availableTags.value.find(t => t.ID === tagId)
      if (tag) {
        posterTags.value.push(tag)
        // 添加到待添加列表
        tagsToAdd.value.push(tagId)
      }
    })
    
    // 清空选择
    selectedTagIds.value = []
  }
}

// 移除标签
const removeTag = (tag) => {
  // 从当前标签列表中移除
  posterTags.value = posterTags.value.filter(t => t.ID !== tag.ID)
  
  // 如果标签原本就存在于海报中（编辑模式），添加到待移除列表
  if (type.value === 'update') {
    tagsToRemove.value.push(tag.ID)
  }
  
  // 将标签添加回可选列表
  if (!availableTags.value.some(t => t.ID === tag.ID)) {
    availableTags.value.push(tag)
  }
}

// 保存标签关系
const saveTagsRelations = async (posterId) => {
  // 新建海报时，添加所有标签
  if (type.value === 'create' && posterTags.value.length > 0) {
    const tagIds = posterTags.value.map(tag => tag.ID)
    try {
      await addTagsToPoster({
        posterId: posterId,
        tagIds: tagIds
      })
    } catch (error) {
      console.error('添加标签失败:', error)
    }
  } 
  // 编辑海报时，处理添加和移除的标签
  else if (type.value === 'update') {
    // 添加新标签
    if (tagsToAdd.value.length > 0) {
      try {
        await addTagsToPoster({
          posterId: posterId,
          tagIds: tagsToAdd.value
        })
      } catch (error) {
        console.error('添加标签失败:', error)
      }
    }
    
    // 移除标签
    if (tagsToRemove.value.length > 0) {
      try {
        await removeTagsFromPoster({
          posterId: posterId,
          tagIds: tagsToRemove.value
        })
      } catch (error) {
        console.error('移除标签失败:', error)
      }
    }
  }
}

// 初始化标签数据
onMounted(async () => {
  await loadPosterTags()
  await loadAvailableTags()
  tagsToAdd.value = []
  tagsToRemove.value = []
})

// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createPosters(formData.value)
               if (res.code === 0 && posterTags.value.length > 0) {
                 // 保存标签关系
                 await saveTagsRelations(res.data.ID)
               }
               break
             case 'update':
               res = await updatePosters(formData.value)
               if (res.code === 0) {
                 // 保存标签关系
                 await saveTagsRelations(formData.value.ID)
               }
               break
             default:
               res = await createPosters(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
             // 返回列表页
             back()
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
.tags-container {
  margin-top: 10px;
}

.current-tags {
  margin-bottom: 12px;
}

.tag-item {
  margin-right: 8px;
  margin-bottom: 8px;
}

.empty-tip {
  color: #909399;
  font-size: 14px;
  margin-bottom: 12px;
}

.add-tags {
  margin-top: 12px;
}
</style>
