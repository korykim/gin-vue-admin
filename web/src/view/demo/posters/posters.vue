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
      
            <el-form-item label="海报标题" prop="title">
  <el-input v-model="searchInfo.title" placeholder="搜索条件" />
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
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate  template-id="demo_Posters" />
            <ExportExcel  template-id="demo_Posters" filterDeleted/>
            <ImportExcel  template-id="demo_Posters" @on-success="getTableData" />
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
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column sortable align="left" label="海报标题" prop="title" width="120" />

            <el-table-column align="left" label="海报图片地址" prop="image" width="120" />

            <el-table-column align="left" label="标签" width="220">
              <template #default="scope">
                <el-tag
                  v-for="tag in scope.row.tags"
                  :key="tag.ID"
                  :type="tag.color"
                  class="tag-item"
                  style="margin-right: 5px; margin-bottom: 5px;"
                >
                  {{ tag.name }}
                </el-tag>
                <span v-if="!scope.row.tags || scope.row.tags.length === 0" class="no-tags-text">暂无标签</span>
              </template>
            </el-table-column>

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updatePostersFunc(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="collection" class="table-button" @click="manageTags(scope.row)">标签管理</el-button>
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
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="海报标题:" prop="title">
    <el-input v-model="formData.title" :clearable="true" placeholder="请输入海报标题" />
</el-form-item>
            <el-form-item label="海报图片地址:" prop="image">
    <el-input v-model="formData.image" :clearable="true" placeholder="请输入海报图片地址" />
</el-form-item>
            <el-form-item label="标签管理:">
              <div class="dialog-tags-container">
                <div class="current-tags" v-if="dialogPosterTags.length > 0">
                  <el-tag
                    v-for="tag in dialogPosterTags"
                    :key="tag.ID"
                    :type="tag.color"
                    class="tag-item"
                    closable
                    @close="removeDialogTag(tag)"
                  >
                    {{ tag.name }}
                  </el-tag>
                </div>
                <div class="no-tags" v-else>
                  <span class="no-tags-text">尚未添加标签</span>
                </div>
                
                <div class="add-tags">
                  <el-select
                    v-model="dialogSelectedTagIds"
                    multiple
                    filterable
                    placeholder="请选择要添加的标签"
                    style="width: 100%;"
                    @change="handleDialogTagsChange"
                  >
                    <el-option
                      v-for="tag in dialogAvailableTags"
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="海报标题">
    {{ detailFrom.title }}
</el-descriptions-item>
                    <el-descriptions-item label="海报图片地址">
    {{ detailFrom.image }}
</el-descriptions-item>
                    <el-descriptions-item label="标签">
    <div class="detail-tags">
      <el-tag
        v-for="tag in detailFrom.tags"
        :key="tag.ID"
        :type="tag.color"
        class="tag-item"
        style="margin-right: 5px; margin-bottom: 5px;"
      >
        {{ tag.name }}
      </el-tag>
      <span v-if="!detailFrom.tags || detailFrom.tags.length === 0" class="no-tags-text">暂无标签</span>
    </div>
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

    <!-- 标签管理弹窗 -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="tagsDialogVisible" :show-close="true" :before-close="closeTagsDialog" title="标签管理">
      <div class="tags-container">
        <div class="current-tags" v-if="currentPosterTags.length > 0">
          <h3>已添加标签</h3>
          <el-tag
            v-for="tag in currentPosterTags"
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
          <el-empty description="尚未添加标签"></el-empty>
        </div>
        
        <div class="add-tags">
          <h3>添加新标签</h3>
          <el-form :inline="true">
            <el-form-item>
              <el-select
                v-model="selectedTagIds"
                multiple
                filterable
                placeholder="请选择要添加的标签"
                style="width: 100%;"
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
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="addTags">添加标签</el-button>
            </el-form-item>
          </el-form>
        </div>

        <el-divider content-position="center">标签管理</el-divider>

        <div class="tags-crud">
          <div class="tags-crud-header">
            <h3>标签列表</h3>
            <el-button type="primary" size="small" icon="plus" @click="openTagDialog('create')">新增标签</el-button>
          </div>
          <el-table :data="allTags" style="width: 100%">
            <el-table-column prop="name" label="标签名称" width="120" />
            <el-table-column prop="color" label="颜色" width="180">
              <template #default="scope">
                <el-tag :type="scope.row.color">{{ scope.row.color }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="操作" width="150">
              <template #default="scope">
                <el-button type="primary" link icon="edit" @click="openTagDialog('update', scope.row)">编辑</el-button>
                <el-button type="primary" link icon="delete" @click="deleteTagConfirm(scope.row)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
    </el-drawer>

    <!-- 标签编辑弹窗 -->
    <el-dialog
      v-model="tagFormVisible"
      :title="tagFormType === 'create' ? '新增标签' : '编辑标签'"
      width="30%"
      :before-close="closeTagDialog"
    >
      <el-form :model="tagFormData" ref="tagFormRef" :rules="tagFormRules" label-width="80px">
        <el-form-item label="标签名称" prop="name">
          <el-input v-model="tagFormData.name" placeholder="请输入标签名称" />
        </el-form-item>
        <el-form-item label="标签颜色" prop="color">
          <el-select v-model="tagFormData.color" placeholder="请选择标签颜色">
            <el-option label="默认" value="" />
            <el-option label="成功" value="success" />
            <el-option label="警告" value="warning" />
            <el-option label="危险" value="danger" />
            <el-option label="信息" value="info" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="closeTagDialog">取消</el-button>
          <el-button type="primary" :loading="tagFormLoading" @click="submitTagForm">确定</el-button>
        </span>
      </template>
    </el-dialog>

  </div>
</template>

<script setup>
import {
  createPosters,
  deletePosters,
  deletePostersByIds,
  updatePosters,
  findPosters,
  getPostersList
} from '@/api/demo/posters'

import { 
  addTagsToPoster, 
  removeTagsFromPoster, 
  getPosterTags 
} from '@/api/demo/poster_tag'

import { 
  getTagsList,
  createTags,
  updateTags,
  deleteTags
} from '@/api/demo/tags'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'Posters'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
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

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    CreatedAt:"CreatedAt",
    ID:"ID",
            title: 'title',
  }

  let sort = sortMap[prop]
  if(!sort){
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
  const table = await getPostersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
    
    // 获取每个海报的标签
    await Promise.all(tableData.value.map(async (poster) => {
      try {
        const res = await getPosterTags({ posterId: poster.ID })
        if (res.code === 0) {
          poster.tags = res.data || []
        } else {
          poster.tags = []
        }
      } catch (error) {
        console.error('获取海报标签失败:', error)
        poster.tags = []
      }
    }))
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deletePostersFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deletePostersByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updatePostersFunc = async(row) => {
    const res = await findPosters({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
        await loadDialogTags(row.ID)
    }
}


// 删除行
const deletePostersFunc = async (row) => {
    const res = await deletePosters({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = async () => {
    type.value = 'create'
    dialogFormVisible.value = true
    await loadDialogTags()
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    // 重置表单和标签状态
    formData.value = {
        title: '',
        image: '',
    }
    dialogPosterTags.value = []
    dialogAvailableTags.value = []
    dialogSelectedTagIds.value = []
    dialogTagsToAdd.value = []
    dialogTagsToRemove.value = []
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createPosters(formData.value)
                  console.log("创建海报API返回结果:", JSON.stringify(res))
                  if (res.code === 0) {
                    // 获取新创建的海报ID
                    let newPosterId = null
                    
                    // 根据API返回格式获取ID
                    if (res.data && typeof res.data === 'object') {
                      // 如果返回的是对象
                      if (res.data.ID) {
                        newPosterId = res.data.ID
                      } else if (res.data.id) {
                        newPosterId = res.data.id
                      } else if (res.data.poster && res.data.poster.ID) {
                        newPosterId = res.data.poster.ID
                      }
                    } else if (typeof res.data === 'number') {
                      // 如果直接返回ID
                      newPosterId = res.data
                    }
                    
                    console.log("解析的新海报ID:", newPosterId)
                    
                    // 保存标签关系
                    if (newPosterId && dialogPosterTags.value.length > 0) {
                      await saveDialogTagsRelations(newPosterId)
                    } else if (dialogPosterTags.value.length > 0) {
                      console.error("无法获取新创建的海报ID，无法添加标签")
                      ElMessage.warning("无法获取新海报ID，标签添加失败")
                    }
                  }
                  break
                case 'update':
                  res = await updatePosters(formData.value)
                  if (res.code === 0) {
                    // 保存标签关系
                    console.log("更新海报成功，ID:", formData.value.ID)
                    await saveDialogTagsRelations(formData.value.ID)
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
                closeDialog()
                getTableData()
              }
      })
}

const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findPosters({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    
    // 获取海报的标签
    try {
      const tagsRes = await getPosterTags({ posterId: row.ID })
      if (tagsRes.code === 0) {
        detailFrom.value.tags = tagsRes.data || []
      } else {
        detailFrom.value.tags = []
      }
    } catch (error) {
      console.error('获取海报标签失败:', error)
      detailFrom.value.tags = []
    }
    
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}

// 标签管理相关
const tagsDialogVisible = ref(false)
const currentPosterId = ref(null)
const currentPosterTags = ref([])
const availableTags = ref([])
const selectedTagIds = ref([])
const allTags = ref([]) // 所有标签列表

// 标签表单相关
const tagFormVisible = ref(false)
const tagFormType = ref('create')
const tagFormData = ref({
  name: '',
  color: ''
})
const tagFormLoading = ref(false)
const tagFormRef = ref()
const tagFormRules = reactive({
  name: [
    { required: true, message: '请输入标签名称', trigger: 'blur' },
    { min: 1, max: 50, message: '长度在1到50个字符', trigger: 'blur' }
  ]
})

// 打开标签管理对话框
const manageTags = async (row) => {
  currentPosterId.value = row.ID
  tagsDialogVisible.value = true
  await loadPosterTags()
  await loadAllTags()
  await loadAvailableTags()
}

// 关闭标签管理对话框
const closeTagsDialog = () => {
  tagsDialogVisible.value = false
  currentPosterId.value = null
  currentPosterTags.value = []
  availableTags.value = []
  selectedTagIds.value = []
  allTags.value = []
}

// 加载所有标签
const loadAllTags = async () => {
  try {
    const res = await getTagsList({ page: 1, pageSize: 100 })
    if (res.code === 0) {
      allTags.value = res.data.list || []
    }
  } catch (error) {
    console.error('获取所有标签失败:', error)
  }
}

// 加载海报标签
const loadPosterTags = async () => {
  try {
    const res = await getPosterTags({ posterId: currentPosterId.value })
    if (res.code === 0) {
      currentPosterTags.value = res.data || []
    }
  } catch (error) {
    console.error('获取海报标签失败:', error)
  }
}

// 加载可用标签
const loadAvailableTags = async () => {
  try {
    const res = await getTagsList({ page: 1, pageSize: 100 })
    if (res.code === 0) {
      // 过滤掉已经添加的标签
      const currentTagIds = currentPosterTags.value.map(tag => tag.ID)
      availableTags.value = res.data.list.filter(tag => !currentTagIds.includes(tag.ID))
    }
  } catch (error) {
    console.error('获取标签列表失败:', error)
  }
}

// 添加标签
const addTags = async () => {
  if (selectedTagIds.value.length === 0) {
    ElMessage.warning('请选择要添加的标签')
    return
  }

  try {
    const res = await addTagsToPoster({
      posterId: currentPosterId.value,
      tagIds: selectedTagIds.value
    })

    if (res.code === 0) {
      ElMessage.success('添加标签成功')
      selectedTagIds.value = []
      await loadPosterTags()
      await loadAvailableTags()
    }
  } catch (error) {
    console.error('添加标签失败:', error)
  }
}

// 移除标签
const removeTag = async (tag) => {
  try {
    const res = await removeTagsFromPoster({
      posterId: currentPosterId.value,
      tagIds: [tag.ID]
    })

    if (res.code === 0) {
      ElMessage.success('移除标签成功')
      await loadPosterTags()
      await loadAvailableTags()
    }
  } catch (error) {
    console.error('移除标签失败:', error)
  }
}

// 打开标签表单对话框
const openTagDialog = (type, tag = null) => {
  tagFormType.value = type
  if (type === 'update' && tag) {
    // 编辑模式，填充表单数据
    tagFormData.value = { ...tag }
  } else {
    // 创建模式，重置表单
    tagFormData.value = {
      name: '',
      color: ''
    }
  }
  tagFormVisible.value = true
}

// 关闭标签表单对话框
const closeTagDialog = () => {
  tagFormVisible.value = false
  tagFormData.value = {
    name: '',
    color: ''
  }
}

// 提交标签表单
const submitTagForm = async () => {
  tagFormRef.value?.validate(async (valid) => {
    if (!valid) return
    
    tagFormLoading.value = true
    try {
      let res
      if (tagFormType.value === 'create') {
        // 创建新标签
        res = await createTags(tagFormData.value)
      } else {
        // 更新现有标签
        res = await updateTags(tagFormData.value)
      }
      
      if (res.code === 0) {
        ElMessage.success(tagFormType.value === 'create' ? '创建标签成功' : '更新标签成功')
        closeTagDialog()
        
        // 刷新标签列表
        await loadAllTags()
        await loadAvailableTags()
        
        // 如果是编辑了已添加的标签，也需要刷新当前海报的标签
        if (tagFormType.value === 'update') {
          await loadPosterTags()
        }
      }
    } catch (error) {
      console.error(tagFormType.value === 'create' ? '创建标签失败:' : '更新标签失败:', error)
    } finally {
      tagFormLoading.value = false
    }
  })
}

// 删除标签确认
const deleteTagConfirm = (tag) => {
  ElMessageBox.confirm('确定要删除此标签吗？删除后将无法恢复，且会从所有关联的海报中移除。', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const res = await deleteTags({ ID: tag.ID })
      if (res.code === 0) {
        ElMessage.success('删除标签成功')
        
        // 刷新标签列表
        await loadAllTags()
        await loadAvailableTags()
        
        // 如果删除的是已添加的标签，也需要刷新当前海报的标签
        if (currentPosterTags.value.some(t => t.ID === tag.ID)) {
          await loadPosterTags()
        }
      }
    } catch (error) {
      console.error('删除标签失败:', error)
    }
  }).catch(() => {
    // 取消删除，不做任何操作
  })
}

// 弹窗标签相关
const dialogPosterTags = ref([])
const dialogAvailableTags = ref([])
const dialogSelectedTagIds = ref([])
const dialogTagsToAdd = ref([])
const dialogTagsToRemove = ref([])

// 加载弹窗标签数据
const loadDialogTags = async (posterId = null) => {
  // 重置标签数据
  dialogPosterTags.value = []
  dialogSelectedTagIds.value = []
  
  // 加载所有标签
  try {
    const res = await getTagsList({ page: 1, pageSize: 100 })
    if (res.code === 0) {
      // 如果是编辑模式，需要获取当前海报的标签
      if (posterId) {
        const tagsRes = await getPosterTags({ posterId })
        if (tagsRes.code === 0) {
          dialogPosterTags.value = tagsRes.data || []
          
          // 过滤掉已经添加的标签
          const currentTagIds = dialogPosterTags.value.map(tag => tag.ID)
          dialogAvailableTags.value = res.data.list.filter(tag => !currentTagIds.includes(tag.ID))
        }
      } else {
        // 新增模式，所有标签都可用
        dialogAvailableTags.value = res.data.list
      }
    }
  } catch (error) {
    console.error('获取标签列表失败:', error)
  }
  
  // 重置待添加和待移除列表
  dialogTagsToAdd.value = []
  dialogTagsToRemove.value = []
}

// 处理弹窗标签选择变化
const handleDialogTagsChange = (value) => {
  // 计算新添加的标签
  const newTags = value.filter(id => !dialogPosterTags.value.some(tag => tag.ID === id))
  
  if (newTags.length > 0) {
    // 将新标签添加到dialogPosterTags中
    newTags.forEach(tagId => {
      const tag = dialogAvailableTags.value.find(t => t.ID === tagId)
      if (tag) {
        dialogPosterTags.value.push(tag)
        // 添加到待添加列表
        dialogTagsToAdd.value.push(tagId)
      }
    })
    
    // 清空选择
    dialogSelectedTagIds.value = []
  }
}

// 从弹窗中移除标签
const removeDialogTag = (tag) => {
  // 从当前标签列表中移除
  dialogPosterTags.value = dialogPosterTags.value.filter(t => t.ID !== tag.ID)
  
  // 如果标签原本就存在于海报中（编辑模式），添加到待移除列表
  if (type.value === 'update') {
    dialogTagsToRemove.value.push(tag.ID)
  }
  
  // 将标签添加回可选列表
  if (!dialogAvailableTags.value.some(t => t.ID === tag.ID)) {
    dialogAvailableTags.value.push(tag)
  }
}

// 保存弹窗标签关系
const saveDialogTagsRelations = async (posterId) => {
  if (!posterId) {
    console.error('保存标签关系失败: 无效的海报ID', posterId)
    ElMessage.error('保存标签关系失败: 无效的海报ID')
    return
  }
  
  // 新建海报时，添加所有标签
  if (type.value === 'create' && dialogPosterTags.value.length > 0) {
    const tagIds = dialogPosterTags.value.map(tag => tag.ID)
    try {
      console.log(`准备添加标签，海报ID:${posterId}，标签IDs:${JSON.stringify(tagIds)}`)
      const res = await addTagsToPoster({
        posterId: Number(posterId), // 确保是数字类型
        tagIds: tagIds
      })
      console.log('添加标签结果:', JSON.stringify(res))
      if (res.code !== 0) {
        ElMessage.error(`添加标签失败: ${res.msg}`)
      }
    } catch (error) {
      console.error('添加标签失败:', error)
      ElMessage.error(`添加标签出错: ${error.message || '未知错误'}`)
    }
  } 
  // 编辑海报时，处理添加和移除的标签
  else if (type.value === 'update') {
    // 添加新标签
    if (dialogTagsToAdd.value.length > 0) {
      try {
        console.log(`准备添加标签(编辑模式)，海报ID:${posterId}，标签IDs:${JSON.stringify(dialogTagsToAdd.value)}`)
        const res = await addTagsToPoster({
          posterId: Number(posterId), // 确保是数字类型
          tagIds: dialogTagsToAdd.value
        })
        console.log('添加标签结果(编辑模式):', JSON.stringify(res))
        if (res.code !== 0) {
          ElMessage.error(`添加标签失败: ${res.msg}`)
        }
      } catch (error) {
        console.error('添加标签失败:', error)
        ElMessage.error(`添加标签出错: ${error.message || '未知错误'}`)
      }
    }
    
    // 移除标签
    if (dialogTagsToRemove.value.length > 0) {
      try {
        const res = await removeTagsFromPoster({
          posterId: Number(posterId), // 确保是数字类型
          tagIds: dialogTagsToRemove.value
        })
        if (res.code !== 0) {
          ElMessage.error(`移除标签失败: ${res.msg}`)
        }
      } catch (error) {
        console.error('移除标签失败:', error)
        ElMessage.error(`移除标签出错: ${error.message || '未知错误'}`)
      }
    }
  }
}

</script>

<style>
.tags-container {
  padding: 0 16px;
}

.current-tags {
  margin-bottom: 24px;
}

.tag-item {
  margin-right: 8px;
  margin-bottom: 8px;
}

.add-tags {
  margin-top: 24px;
}

.no-tags-text {
  color: #909399;
  font-size: 12px;
  font-style: italic;
}

.dialog-tags-container {
  margin-top: 10px;
}

.tags-crud {
  margin-top: 20px;
}

.tags-crud-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}
</style>
