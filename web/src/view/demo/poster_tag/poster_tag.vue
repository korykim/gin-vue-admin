<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="CreatedAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
      <el-form-item label="海报ID" prop="posterId">
        <el-input v-model="searchInfo.posterId" placeholder="搜索海报ID" />
      </el-form-item>
      
      <el-form-item label="标签ID" prop="tagId">
        <el-input v-model="searchInfo.tagId" placeholder="搜索标签ID" />
      </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openAddTagsDialog">添加标签</el-button>
            <el-button type="danger" icon="delete" @click="openRemoveTagsDialog">移除标签</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="ID" prop="ID" width="90" />
        
        <el-table-column sortable align="left" label="创建日期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column sortable align="left" label="海报ID" prop="poster_id" width="120" />

        <el-table-column sortable align="left" label="标签ID" prop="tag_id" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>

    <!-- 添加标签弹窗 -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="addTagsDialogVisible" :show-close="false" :before-close="closeAddTagsDialog">
       <template #header>
          <div class="flex justify-between items-center">
            <span class="text-lg">添加标签到海报</span>
            <div>
              <el-button :loading="btnLoading" type="primary" @click="confirmAddTags">确 定</el-button>
              <el-button @click="closeAddTagsDialog">取 消</el-button>
            </div>
          </div>
        </template>

        <el-form :model="addTagsForm" label-position="top" ref="addTagsFormRef" :rules="addTagsRule" label-width="120px">
          <el-form-item label="选择海报:" prop="posterId">
            <el-input v-model="addTagsForm.posterId" type="number" placeholder="请输入海报ID" />
          </el-form-item>
          <el-form-item label="选择标签:" prop="tagIds">
            <el-select v-model="addTagsForm.tagIds" multiple placeholder="请选择标签" style="width: 100%">
              <el-option v-for="tag in tagOptions" :key="tag.value" :label="tag.label" :value="tag.value" />
            </el-select>
          </el-form-item>
        </el-form>
    </el-drawer>

    <!-- 移除标签弹窗 -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="removeTagsDialogVisible" :show-close="false" :before-close="closeRemoveTagsDialog">
       <template #header>
          <div class="flex justify-between items-center">
            <span class="text-lg">从海报移除标签</span>
            <div>
              <el-button :loading="btnLoading" type="primary" @click="confirmRemoveTags">确 定</el-button>
              <el-button @click="closeRemoveTagsDialog">取 消</el-button>
            </div>
          </div>
        </template>

        <el-form :model="removeTagsForm" label-position="top" ref="removeTagsFormRef" :rules="removeTagsRule" label-width="120px">
          <el-form-item label="选择海报:" prop="posterId">
            <el-input v-model="removeTagsForm.posterId" type="number" placeholder="请输入海报ID" />
          </el-form-item>
          <el-form-item label="选择标签:" prop="tagIds">
            <el-select v-model="removeTagsForm.tagIds" multiple placeholder="请选择标签" style="width: 100%" @change="handlePosterChange">
              <el-option v-for="tag in tagOptions" :key="tag.value" :label="tag.label" :value="tag.value" />
            </el-select>
          </el-form-item>
        </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  addTagsToPoster,
  removeTagsFromPoster,
  getPosterTags,
  getPosterTagRelations
} from '@/api/demo/poster_tag'

import { getTagsList } from '@/api/demo/tags'
import { getPostersList } from '@/api/demo/posters'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, onMounted } from 'vue'
import { useAppStore } from "@/pinia"

defineOptions({
  name: 'PosterTag'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 标签选项
const tagOptions = ref([])
// 海报选项
const posterOptions = ref([])

// 添加标签表单
const addTagsForm = ref({
  posterId: '',
  tagIds: []
})

// 移除标签表单
const removeTagsForm = ref({
  posterId: '',
  tagIds: []
})

// 验证规则
const addTagsRule = reactive({
  posterId: [
    { required: true, message: '请输入海报ID', trigger: 'blur' },
    { type: 'number', message: '海报ID必须为数字', trigger: 'blur', transform: (value) => Number(value) }
  ],
  tagIds: [
    { required: true, message: '请选择标签', trigger: 'change' },
    { type: 'array', min: 1, message: '请至少选择一个标签', trigger: 'change' }
  ]
})

// 验证规则
const removeTagsRule = reactive({
  posterId: [
    { required: true, message: '请输入海报ID', trigger: 'blur' },
    { type: 'number', message: '海报ID必须为数字', trigger: 'blur', transform: (value) => Number(value) }
  ],
  tagIds: [
    { required: true, message: '请选择标签', trigger: 'change' },
    { type: 'array', min: 1, message: '请至少选择一个标签', trigger: 'change' }
  ]
})

const searchRule = reactive({
  CreatedAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elSearchFormRef = ref()
const addTagsFormRef = ref()
const removeTagsFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    CreatedAt: "CreatedAt",
    ID: "ID",
    poster_id: "PosterID",
    tag_id: "TagID"
  }

  let sort = sortMap[prop]
  if (!sort) {
    sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getPosterTagRelations({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 获取标签和海报选项
const getOptions = async () => {
  // 获取标签选项
  const tagsRes = await getTagsList({ page: 1, pageSize: 100 })
  if (tagsRes.code === 0) {
    tagOptions.value = tagsRes.data.list.map(item => ({
      value: item.ID,
      label: `${item.name} (ID: ${item.ID})`
    }))
  }

  // 获取海报选项
  const postersRes = await getPostersList({ page: 1, pageSize: 100 })
  if (postersRes.code === 0) {
    posterOptions.value = postersRes.data.list.map(item => ({
      value: item.ID,
      label: `${item.title} (ID: ${item.ID})`
    }))
  }
}

// 当选择海报时，获取该海报的所有标签
const handlePosterChange = async () => {
  if (removeTagsForm.value.posterId) {
    try {
      const res = await getPosterTags({ posterId: removeTagsForm.value.posterId })
      if (res.code === 0) {
        // 更新标签选项为当前海报的标签
        removeTagsForm.value.tagIds = []
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

onMounted(() => {
  getTableData()
  getOptions()
})

// ============== 表格控制部分结束 ===============

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除该关联关系吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const res = await removeTagsFromPoster({ posterId: row.poster_id, tagIds: [row.tag_id] })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        getTableData()
      }
    } catch (error) {
      console.error('删除失败:', error)
    }
  })
}

// 添加标签弹窗
const addTagsDialogVisible = ref(false)

// 打开添加标签弹窗
const openAddTagsDialog = () => {
  addTagsDialogVisible.value = true
  addTagsForm.value = {
    posterId: '',
    tagIds: []
  }
}

// 关闭添加标签弹窗
const closeAddTagsDialog = () => {
  addTagsDialogVisible.value = false
}

// 确认添加标签
const confirmAddTags = async () => {
  addTagsFormRef.value?.validate(async (valid) => {
    if (!valid) return
    
    btnLoading.value = true
    try {
      const res = await addTagsToPoster({
        posterId: Number(addTagsForm.value.posterId),
        tagIds: addTagsForm.value.tagIds
      })
      
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '添加标签成功'
        })
        closeAddTagsDialog()
        getTableData()
      }
    } catch (error) {
      console.error('添加标签失败:', error)
    } finally {
      btnLoading.value = false
    }
  })
}

// 移除标签弹窗
const removeTagsDialogVisible = ref(false)

// 打开移除标签弹窗
const openRemoveTagsDialog = () => {
  removeTagsDialogVisible.value = true
  removeTagsForm.value = {
    posterId: '',
    tagIds: []
  }
}

// 关闭移除标签弹窗
const closeRemoveTagsDialog = () => {
  removeTagsDialogVisible.value = false
}

// 确认移除标签
const confirmRemoveTags = async () => {
  removeTagsFormRef.value?.validate(async (valid) => {
    if (!valid) return
    
    btnLoading.value = true
    try {
      const res = await removeTagsFromPoster({
        posterId: Number(removeTagsForm.value.posterId),
        tagIds: removeTagsForm.value.tagIds
      })
      
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '移除标签成功'
        })
        closeRemoveTagsDialog()
        getTableData()
      }
    } catch (error) {
      console.error('移除标签失败:', error)
    } finally {
      btnLoading.value = false
    }
  })
}
</script>

<style>
</style> 