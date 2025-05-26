<template>
    <div class="amaplist-container">
      <JumpTop :right="50" :bottom="20" text="回到顶部" />

      <div class="gva-search-box">
        <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
          <el-form-item label="名称" prop="name">
            <el-input v-model="searchInfo.name" placeholder="搜索POI名称" />
          </el-form-item>
          <el-form-item label="类型" prop="type">
            <el-input v-model="searchInfo.type" placeholder="搜索POI类型" />
          </el-form-item>
          <el-form-item label="地区" prop="adname">
            <el-input v-model="searchInfo.adname" placeholder="搜索地区" />
          </el-form-item>
          <el-form-item label="地址" prop="address">
            <el-input v-model="searchInfo.address" placeholder="搜索地址关键词" />
          </el-form-item>
  
          <el-form-item>
            <el-button type="primary" @click="onSubmit">查询</el-button>
            <el-button @click="onReset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>

      <div class="content-main">
        <div class="gva-btn-list">
          <el-button type="primary" @click="exportToExcel">导出数据</el-button>
          <el-button type="danger" @click="clearAllPois">清空所有数据</el-button>
        </div>
        <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="id"
          @selection-change="handleSelectionChange"
          v-loading="loading"
          height="calc(100% - 100px)"
        >
          <el-table-column type="selection" width="55" />
          
          <el-table-column align="left" label="名称" prop="name" width="120" />
          
          <el-table-column align="left" label="ID" prop="id" width="120" />
          
          <el-table-column align="left" label="类型" prop="type" width="180" />
          
          <el-table-column align="left" label="电话" prop="tel" width="120" />
          
          <el-table-column align="left" label="地址" prop="address" width="180" />
          
          <el-table-column align="left" label="省份" prop="pname" width="100" />
          
          <el-table-column align="left" label="城市" prop="cityname" width="100" />
          
          <el-table-column align="left" label="区域" prop="adname" width="100" />
          
          <el-table-column label="照片" prop="photos" min-width="200">
            <template #default="scope">
              <div class="photo-list">
                <template v-if="scope.row.photos && scope.row.photos.length > 0">
                  <el-image 
                    v-for="(photo, index) in scope.row.photos" 
                    :key="index" 
                    style="width: 80px; height: 80px; margin-right: 5px;" 
                    :src="photo.url" 
                    :preview-src-list="getPhotoUrls(scope.row.photos)"
                    fit="cover"
                    :initial-index="index"
                    preview-teleported
                  />
                </template>
                <span v-else class="no-photo">暂无照片</span>
              </div>
            </template>
          </el-table-column>
          
          <el-table-column align="left" label="操作" fixed="right" width="120">
            <template #default="scope">
              <el-button type="primary" link @click="deleteRow(scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        
        <el-empty v-if="!loading && tableData.length === 0" description="暂无数据" />
        
        <div class="gva-pagination" v-if="total > 0">
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

    </div>

  </template>
  
  <script setup>
  import { ref, onMounted, onActivated } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import * as XLSX from 'xlsx'
  import JumpTop from '@/components/jumptop/jumptop.vue'

  
  defineOptions({
    name: 'amaplist'
  })
  
  // 表格控制部分
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({})
  const allPoisData = ref([])
  const loading = ref(false)
  
  // 多选数据
  const multipleSelection = ref([])
  
  // 处理多选
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
  }
  
  // 获取照片URL列表
  const getPhotoUrls = (photos) => {
    if (!photos || !photos.length) return []
    return photos.map(photo => photo.url)
  }
  
  // 加载POI数据
  const loadPoisData = () => {
    loading.value = true
    try {
      const savedPoisStr = localStorage.getItem('amapPois')
      if (savedPoisStr) {
        allPoisData.value = JSON.parse(savedPoisStr)
        console.log(`从localStorage加载了${allPoisData.value.length}个POI点`)
        filterTableData()
      } else {
        allPoisData.value = []
        tableData.value = []
        total.value = 0
      }
    } catch (error) {
      console.error('加载POI数据失败:', error)
      ElMessage.error('加载POI数据失败')
      allPoisData.value = []
      tableData.value = []
      total.value = 0
    } finally {
      loading.value = false
    }
  }
  
  // 过滤表格数据
  const filterTableData = () => {
    let filteredData = [...allPoisData.value]
    
    // 应用搜索过滤
    if (searchInfo.value.name) {
      filteredData = filteredData.filter(item => 
        item.name && item.name.toLowerCase().includes(searchInfo.value.name.toLowerCase())
      )
    }
    
    if (searchInfo.value.type) {
      filteredData = filteredData.filter(item => 
        item.type && item.type.toLowerCase().includes(searchInfo.value.type.toLowerCase())
      )
    }
    
    if (searchInfo.value.adname) {
      filteredData = filteredData.filter(item => 
        item.adname && item.adname.toLowerCase().includes(searchInfo.value.adname.toLowerCase())
      )
    }
    
    // 添加地址模糊搜索
    if (searchInfo.value.address) {
      filteredData = filteredData.filter(item => 
        item.address && item.address.toLowerCase().includes(searchInfo.value.address.toLowerCase())
      )
    }
    
    // 更新总数
    total.value = filteredData.length
    
    // 分页
    const start = (page.value - 1) * pageSize.value
    const end = start + pageSize.value
    tableData.value = filteredData.slice(start, end)
  }
  
  // 搜索
  const onSubmit = () => {
    page.value = 1
    filterTableData()
  }
  
  // 重置
  const onReset = () => {
    searchInfo.value = {}
    page.value = 1
    filterTableData()
  }
  
  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    filterTableData()
  }
  
  // 修改页面容量
  const handleCurrentChange = (val) => {
    page.value = val
    filterTableData()
  }
  
  // 删除单行数据
  const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除这条POI数据吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      try {
        const newPoisData = allPoisData.value.filter(item => item.id !== row.id)
        localStorage.setItem('amapPois', JSON.stringify(newPoisData))
        ElMessage.success('删除成功')
        loadPoisData() // 重新加载数据
      } catch (error) {
        console.error('删除POI数据失败:', error)
        ElMessage.error('删除POI数据失败')
      }
    })
  }
  
  // 清空所有POI数据
  const clearAllPois = () => {
    ElMessageBox.confirm('确定要清空所有POI数据吗?', '警告', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      try {
        localStorage.removeItem('amapPois')
        ElMessage.success('已清空所有POI数据')
        loadPoisData() // 重新加载数据
      } catch (error) {
        console.error('清空POI数据失败:', error)
        ElMessage.error('清空POI数据失败')
      }
    })
  }
  
  // 导出Excel
  const exportToExcel = () => {
    try {
      if (!allPoisData.value || allPoisData.value.length === 0) {
        ElMessage.warning('没有数据可导出')
        return
      }
      
      // 创建工作簿
      const workbook = XLSX.utils.book_new()
      
      // 准备数据 - 处理照片数据，将照片URL数组转换为字符串
      const exportData = allPoisData.value.map(item => {
        const newItem = { ...item }
        // 处理照片数据，将对象数组转换为URL字符串，用分号分隔
        if (newItem.photos && newItem.photos.length) {
          newItem.photoUrls = newItem.photos.map(photo => photo.url).join('; ')
        } else {
          newItem.photoUrls = ''
        }
        // 删除原始照片对象数组，避免导出复杂对象
        delete newItem.photos
        return newItem
      })
      
      // 创建工作表
      const worksheet = XLSX.utils.json_to_sheet(exportData)
      
      // 添加工作表到工作簿
      XLSX.utils.book_append_sheet(workbook, worksheet, 'POI数据')
      
      // 设置列宽
      const columnWidths = [
        { wch: 20 }, // 名称
        { wch: 20 }, // ID
        { wch: 15 }, // 类型
        { wch: 15 }, // 电话
        { wch: 30 }, // 地址
        { wch: 10 }, // 省份
        { wch: 10 }, // 城市
        { wch: 10 }, // 区域
        { wch: 50 }  // 照片URLs
      ]
      worksheet['!cols'] = columnWidths
      
      // 导出Excel文件
      XLSX.writeFile(workbook, 'POI数据.xlsx')
      
      ElMessage.success('导出成功')
    } catch (error) {
      console.error('导出Excel失败:', error)
      ElMessage.error('导出Excel失败: ' + error.message)
    }
  }
  
  // 组件挂载时加载数据
  onMounted(() => {
    loadPoisData()
  })
  
  // 当组件被激活时重新加载数据（用于标签页切换）
  onActivated(() => {
    loadPoisData()
  })
  </script>
  
  <style lang="scss" scoped>
  .amaplist-container {
    height: 100vh;
    width: 100%;
    display: flex;
    flex-direction: column;
    overflow-y: auto;
    background-color: var(--el-bg-color);
    padding: 16px;
    box-sizing: border-box;
  }
  
  .gva-search-box {
    margin-bottom: 16px;
    background-color: var(--el-bg-color-overlay);
    padding: 16px;
    border-radius: 4px;
    box-shadow: var(--el-box-shadow-light);
  }
  
  .content-main {
    flex: 1;
    display: flex;
    flex-direction: column;
    overflow: auto;
    background-color: var(--el-bg-color-overlay);
    border-radius: 4px;
    box-shadow: var(--el-box-shadow-light);
    padding: 16px;
    height: calc(100vh - 150px);
  }
  
  .gva-btn-list {
    margin-bottom: 16px;
    display: flex;
    gap: 8px;
  }
  
  .photo-list {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
  }
  
  .no-photo {
    color: #909399;
    font-size: 14px;
  }
  
  .gva-pagination {
    margin-top: 16px;
    padding: 8px 0;
    display: flex;
    justify-content: flex-end;
  }
  
  /* 自定义滚动条 */
  :deep(::-webkit-scrollbar) {
    width: 10px;
    height: 10px;
  }
  
  :deep(::-webkit-scrollbar-track) {
    background: #f1f1f1;
    border-radius: 4px;
  }
  
  :deep(::-webkit-scrollbar-thumb) {
    background: #888;
    border-radius: 4px;
  }
  
  :deep(::-webkit-scrollbar-thumb:hover) {
    background: #555;
  }
  
  /* 确保表格内容可以滚动 */
  .el-table {
    flex: 1;
    overflow: auto;
  }
  </style>
  
  <style>
  /* 全局样式，确保图片预览在最顶层 */
  .el-image-viewer__wrapper {
    z-index: 3000 !important;
  }
  
  /* 确保图片预览中的按钮可以点击 */
  .el-image-viewer__btn {
    z-index: 3100 !important;
    pointer-events: auto !important;
  }
  
  .el-image-viewer__actions {
    z-index: 3100 !important;
    pointer-events: auto !important;
  }
  
  .el-image-viewer__close {
    z-index: 3100 !important;
    pointer-events: auto !important;
  }
  
  .el-image-viewer__canvas {
    pointer-events: auto !important;
  }
  
  .el-image-viewer__mask {
    pointer-events: auto !important;
  }
  </style> 