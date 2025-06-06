<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
            <el-form-item label="HS编码" prop="hsCode">
  <el-input v-model="searchInfo.hsCode" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="产品名称" prop="productName">
  <el-input v-model="searchInfo.productName" placeholder="搜索条件" />
</el-form-item>
            

<el-form-item label="原产地" prop="placeOfOrigin">
  <el-input v-model="searchInfo.placeOfOrigin" placeholder="搜索条件" />
</el-form-item>


            <el-form-item label="生产企业信息" prop="factoryInformation">
  <el-input v-model="searchInfo.factoryInformation" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="进口商信息" prop="importerInformation">
  <el-input v-model="searchInfo.importerInformation" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="进口商备案号" prop="recordNumber">
  <el-input v-model="searchInfo.recordNumber" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="进口口岸" prop="importPort">
  <el-input v-model="searchInfo.importPort" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="年份" prop="years">
  <el-input v-model.number="searchInfo.years" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="月份" prop="month">
  <el-input v-model.number="searchInfo.month" placeholder="搜索条件" />
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
            <ExportTemplate  template-id="kotra_ProductImportRecords" />
            <ExportExcel  template-id="kotra_ProductImportRecords" filterDeleted/>
            <ImportExcel  template-id="kotra_ProductImportRecords" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        height="800"
        >
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="HS编码" prop="hsCode" width="120" />

            <!-- <el-table-column align="left" label="检验检疫编号" prop="singleNumber" width="120" /> -->

            <el-table-column align="left" label="产品名称" prop="productName" width="220" />

            <el-table-column align="left" label="原产地" prop="placeOfOrigin" width="120" />

            <el-table-column align="left" label="生产企业信息" prop="factoryInformation" width="220" />

            <el-table-column align="left" label="进口商信息" prop="importerInformation" width="240" show-overflow-tooltip>
              <template #default="scope">
                <el-button 
                  type="primary" 
                  link 
                  size="small" 
                  @click="openQixin(scope.row.importerInformation)"
                >
                  {{ scope.row.importerInformation }}
                </el-button>
                <el-button
                  type="info"
                  link
                  size="small"
                  @click="copyText(scope.row.importerInformation)"
                >
                  <el-icon><DocumentCopy /></el-icon>
                </el-button>
              </template>
            </el-table-column>

            <el-table-column align="left" label="重量(kg)" prop="weight" width="80" />

            <el-table-column align="left" label="进口口岸" prop="importPort" width="120" />

            <el-table-column sortable align="left" label="年份" prop="years" width="60" />

            <el-table-column sortable align="left" label="月份" prop="month" width="60" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateProductImportRecordsFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
            <el-form-item label="HS编码:" prop="hsCode">
    <el-input v-model="formData.hsCode" :clearable="true" placeholder="请输入HS编码" />
</el-form-item>
            <el-form-item label="检验检疫编号:" prop="singleNumber">
    <el-input v-model="formData.singleNumber" :clearable="true" placeholder="请输入检验检疫编号" />
</el-form-item>
            <el-form-item label="产品名称:" prop="productName">
    <el-input v-model="formData.productName" :clearable="true" placeholder="请输入产品名称" />
</el-form-item>
            <el-form-item label="原产地:" prop="placeOfOrigin">
    <el-input v-model="formData.placeOfOrigin" :clearable="true" placeholder="请输入原产地" />
</el-form-item>
            <el-form-item label="生产企业信息:" prop="factoryInformation">
    <el-input v-model="formData.factoryInformation" :clearable="true" placeholder="请输入生产企业信息" />
</el-form-item>
            <el-form-item label="进口商信息:" prop="importerInformation">
    <el-input v-model="formData.importerInformation" :clearable="true" placeholder="请输入进口商信息" />
</el-form-item>
            <el-form-item label="进口商备案号:" prop="recordNumber">
    <el-input v-model="formData.recordNumber" :clearable="true" placeholder="请输入进口商备案号" />
</el-form-item>
            <el-form-item label="重量(kg):" prop="weight">
    <el-input v-model="formData.weight" :clearable="true" placeholder="请输入重量(kg)" />
</el-form-item>
            <el-form-item label="未准入境的事实:" prop="issueFacts">
    <el-input v-model="formData.issueFacts" :clearable="true" placeholder="请输入未准入境的事实" />
</el-form-item>
            <el-form-item label="进口口岸:" prop="importPort">
    <el-input v-model="formData.importPort" :clearable="true" placeholder="请输入进口口岸" />
</el-form-item>
            <el-form-item label="年份:" prop="years">
    <el-input v-model.number="formData.years" :clearable="true" placeholder="请输入年份" />
</el-form-item>
            <el-form-item label="月份:" prop="month">
    <el-input v-model.number="formData.month" :clearable="true" placeholder="请输入月份" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="HS编码">
    {{ detailFrom.hsCode }}
</el-descriptions-item>
                    <el-descriptions-item label="检验检疫编号">
    {{ detailFrom.singleNumber }}
</el-descriptions-item>
                    <el-descriptions-item label="产品名称">
    {{ detailFrom.productName }}
</el-descriptions-item>
                    <el-descriptions-item label="原产地">
    {{ detailFrom.placeOfOrigin }}
</el-descriptions-item>
                    <el-descriptions-item label="生产企业信息">
    {{ detailFrom.factoryInformation }}
</el-descriptions-item>
                    <el-descriptions-item label="进口商信息">
    {{ detailFrom.importerInformation }}
</el-descriptions-item>
                    <el-descriptions-item label="进口商备案号">
    {{ detailFrom.recordNumber }}
</el-descriptions-item>
                    <el-descriptions-item label="重量(kg)">
    {{ detailFrom.weight }}
</el-descriptions-item>
                    <el-descriptions-item label="未准入境的事实">
    {{ detailFrom.issueFacts }}
</el-descriptions-item>
                    <el-descriptions-item label="进口口岸">
    {{ detailFrom.importPort }}
</el-descriptions-item>
                    <el-descriptions-item label="年份">
    {{ detailFrom.years }}
</el-descriptions-item>
                    <el-descriptions-item label="月份">
    {{ detailFrom.month }}
</el-descriptions-item>
                    <el-descriptions-item label="创建时间">
    {{ detailFrom.createTime }}
</el-descriptions-item>
                    <el-descriptions-item label="更新时间">
    {{ detailFrom.updateTime }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createProductImportRecords,
  deleteProductImportRecords,
  deleteProductImportRecordsByIds,
  updateProductImportRecords,
  findProductImportRecords,
  getProductImportRecordsList
} from '@/api/kotra/productImportRecords'

// 全量引入格式化工具 请按需保留
//import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"
import { DocumentCopy } from '@element-plus/icons-vue'

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'ProductImportRecords'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            hsCode: '',
            singleNumber: '',
            productName: '',
            placeOfOrigin: '',
            factoryInformation: '',
            importerInformation: '',
            recordNumber: '',
            weight: '',
            issueFacts: '',
            importPort: '',
            years: undefined,
            month: undefined,
        })



// 验证规则
const rule = reactive({
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
            years: 'years',
            month: 'month',
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
  const table = await getProductImportRecordsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
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
            deleteProductImportRecordsFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteProductImportRecordsByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateProductImportRecordsFunc = async(row) => {
    const res = await findProductImportRecords({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteProductImportRecordsFunc = async (row) => {
    const res = await deleteProductImportRecords({ id: row.id })
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
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        hsCode: '',
        singleNumber: '',
        productName: '',
        placeOfOrigin: '',
        factoryInformation: '',
        importerInformation: '',
        recordNumber: '',
        weight: '',
        issueFacts: '',
        importPort: '',
        years: undefined,
        month: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createProductImportRecords(formData.value)
                  break
                case 'update':
                  res = await updateProductImportRecords(formData.value)
                  break
                default:
                  res = await createProductImportRecords(formData.value)
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
  const res = await findProductImportRecords({ id: row.id })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}

const openQixin = (companyName) => {
  const url = `https://www.qixin.com/search?key=${encodeURIComponent(companyName)}`;
  window.open(url, '_blank');
}

// 复制文本到剪贴板
const copyText = (text) => {
  if (navigator.clipboard && window.isSecureContext) {
    // 使用现代的Clipboard API
    navigator.clipboard.writeText(text).then(() => {
      ElMessage({
        type: 'success',
        message: '复制成功'
      })
    }).catch(() => {
      ElMessage({
        type: 'error',
        message: '复制失败'
      })
    })
  } else {
    // 兼容模式
    const textArea = document.createElement('textarea')
    textArea.value = text
    textArea.style.position = 'fixed'
    textArea.style.left = '-999999px'
    textArea.style.top = '-999999px'
    document.body.appendChild(textArea)
    textArea.focus()
    textArea.select()
    
    try {
      document.execCommand('copy')
      ElMessage({
        type: 'success',
        message: '复制成功'
      })
    } catch (err) {
      ElMessage({
        type: 'error',
        message: '复制失败'
      })
    }
    
    document.body.removeChild(textArea)
  }
}

</script>

<style>

</style>
