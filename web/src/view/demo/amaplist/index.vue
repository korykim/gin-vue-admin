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
          <el-button type="danger" :disabled="!multipleSelection.length" @click="batchDeletePois">批量删除</el-button>
          <el-tag v-if="multipleSelection.length" type="info" class="ml-2">已选择 {{ multipleSelection.length }} 项</el-tag>
        </div>
        <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="id"
          @selection-change="handleSelectionChange"
          v-loading="loading"
          height="800"
           
        >
          <el-table-column type="selection" width="55" />
          
          <el-table-column align="left" label="名称" prop="name" width="120" />
          
          <el-table-column align="left" label="ID" prop="id" width="120">
            <template #default="scope">
              <a href="javascript:;" class="id-link" @click="showPoiDetail(scope.row.id)">{{ scope.row.id }}</a>
            </template>
          </el-table-column>
          
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

      <!-- POI详情弹窗 -->
      <el-dialog
        v-model="poiDetailVisible"
        title="POI详情信息"
        width="70%"
        :close-on-click-modal="false"
        destroy-on-close
      >
        <div v-loading="poiDetailLoading">
          <div v-if="poiDetail" class="poi-detail-container">
            <div class="poi-detail-header">
              <h2>{{ poiDetail.name }}</h2>
              <div class="poi-basic-info">
                <div><span class="info-label">ID:</span> {{ poiDetail.id }}</div>
                <div><span class="info-label">类型:</span> {{ poiDetail.type }}</div>
                <div><span class="info-label">地址:</span> {{ poiDetail.address }}</div>
                <div><span class="info-label">位置:</span> {{ poiDetail.location }}</div>
              </div>
            </div>

            <!-- 联系方式和营业信息 -->
            <div v-if="poiDetail.business" class="poi-section">
              <h3>营业信息</h3>
              <div class="poi-business-info">
                <div v-if="poiDetail.business.business_area"><span class="info-label">商圈:</span> {{ poiDetail.business.business_area }}</div>
                <div v-if="poiDetail.business.tel"><span class="info-label">电话:</span> {{ poiDetail.business.tel }}</div>
                <div v-if="poiDetail.business.rating"><span class="info-label">评分:</span> {{ poiDetail.business.rating }}</div>
                <div v-if="poiDetail.business.opentime_today"><span class="info-label">今日营业时间:</span> {{ poiDetail.business.opentime_today }}</div>
                <div v-if="poiDetail.business.opentime_week"><span class="info-label">每周营业时间:</span> {{ poiDetail.business.opentime_week }}</div>
                <div v-if="poiDetail.business.rectag"><span class="info-label">推荐标签:</span> {{ poiDetail.business.rectag }}</div>
                <div v-if="poiDetail.business.keytag"><span class="info-label">关键标签:</span> {{ poiDetail.business.keytag }}</div>
              </div>
            </div>

            <!-- 照片展示 -->
            <div v-if="poiDetail.photos && poiDetail.photos.length > 0" class="poi-section">
              <h3>照片</h3>
              <div class="poi-photos">
                <el-image 
                  v-for="(photo, index) in poiDetail.photos" 
                  :key="index" 
                  style="width: 120px; height: 120px; margin-right: 10px; margin-bottom: 10px;" 
                  :src="photo.url" 
                  :preview-src-list="poiDetail.photos.map(p => p.url)"
                  fit="cover"
                  :initial-index="index"
                  preview-teleported
                />
              </div>
            </div>

            <!-- 导航信息 -->
            <div v-if="poiDetail.navi" class="poi-section">
              <h3>导航信息</h3>
              <div class="poi-navi-info">
                <div v-if="poiDetail.navi.entr_location"><span class="info-label">入口位置:</span> {{ poiDetail.navi.entr_location }}</div>
                <div v-if="poiDetail.navi.gridcode"><span class="info-label">网格编码:</span> {{ poiDetail.navi.gridcode }}</div>
              </div>
            </div>

            <!-- 原始数据 -->
            <div class="poi-section">
              <h3>原始JSON数据</h3>
              <json-viewer
                :value="parseJsonIfPossible(poiDetailJson)"
                copyable
                boxed
                sort
                expanded
                theme="light"
                class="json-viewer-container"
              />
            </div>
          </div>
          <el-empty v-else description="暂无详情数据" />
        </div>
      </el-dialog>

    </div>

  </template>
  
  <script setup>
  import { ref, onMounted, onActivated } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import * as XLSX from 'xlsx'
  import JumpTop from '@/components/jumptop/jumptop.vue'
  import axios from 'axios'
  import { JsonViewer } from 'vue3-json-viewer'
  import 'vue3-json-viewer/dist/index.css'

  
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
  const multipleTable = ref(null)
  
  // 添加POI详情相关变量
  const poiDetailVisible = ref(false)
  const poiDetailLoading = ref(false)
  const poiDetail = ref(null)
  const poiDetailJson = ref('')
  
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
      
      // 定义导出字段顺序
      const exportFields = [
        'BigCategory', 'MidCategory', 'SubCategory', 'adcode', 'address', 
        'adname', 'citycode', 'cityname', 'cost', 'discount', 'distance', 
        'email', 'entrLocation', 'exitLocation', 'groupbuy', 'id', 'indoorMap', 
        'location', 'name', 'pcode', 'photos', 'pname', 'poiType', 'postcode', 
        'rating', 'shopinfo', 'tel', 'type', 'website'
      ]
      
      // 自定义表头映射
      const headerMap = {
        'BigCategory': '大类别',
        'MidCategory': '中类别',
        'SubCategory': '子类别',
        'adcode': '区域编码',
        'address': '详细地址',
        'adname': '区域名称',
        'citycode': '城市编码',
        'cityname': '城市名称',
        'cost': '消费金额',
        'discount': '是否有优惠',
        'distance': '距离(米)',
        'email': '电子邮箱',
        'entrLocation': '入口位置',
        'exitLocation': '出口位置',
        'groupbuy': '是否支持团购',
        'id': 'POI唯一标识',
        'indoorMap': '是否有室内地图',
        'location': '位置坐标[经度,纬度]',
        'name': 'POI名称',
        'pcode': '省份编码',
        'photos': '照片信息',
        'pname': '省份名称',
        'poiType': 'POI类型编码',
        'postcode': '邮政编码',
        'rating': '评分',
        'shopinfo': '商铺信息',
        'tel': '联系电话',
        'type': '类型描述',
        'website': '网站'
      }
      
      // 准备数据 - 按照指定字段顺序导出
      const exportData = allPoisData.value.map(poi => {
        const rowData = {}
        
        // 按照指定顺序填充字段
        exportFields.forEach(field => {
          if (field === 'photos') {
            // 将照片数组转换为JSON字符串，确保空值也是有效的JSON格式
            if (poi[field] && Array.isArray(poi[field]) && poi[field].length > 0) {
              rowData[field] = JSON.stringify(poi[field])
            } else {
              rowData[field] = '[]' // 空数组的JSON字符串
            }
          } else if (field === 'location') {
            // 将位置坐标数组转换为JSON字符串，确保空值也是有效的JSON格式
            if (poi[field] && Array.isArray(poi[field]) && poi[field].length > 0) {
              rowData[field] = JSON.stringify(poi[field])
            } else {
              rowData[field] = '[]' // 空数组的JSON字符串
            }
          } else if (field === 'entrLocation' || field === 'exitLocation') {
            // 处理可能为null的位置字段
            if (poi[field] === null || poi[field] === undefined) {
              rowData[field] = 'null' // JSON格式的null
            } else if (typeof poi[field] === 'object') {
              rowData[field] = JSON.stringify(poi[field])
            } else {
              rowData[field] = poi[field]
            }
          } else if (field === 'cost') {
            // 处理消费金额字段，为空时使用0.00填充
            if (poi[field] === null || poi[field] === undefined || poi[field] === '') {
              rowData[field] = '0.00'
            } else {
              // 确保显示两位小数
              rowData[field] = Number(poi[field]).toFixed(2)
            }
          } else if (field === 'rating') {
            // 处理评分字段，为空时使用0.0填充
            if (poi[field] === null || poi[field] === undefined || poi[field] === '') {
              rowData[field] = '0.0'
            } else {
              // 确保显示一位小数
              rowData[field] = Number(poi[field]).toFixed(1)
            }
          } else if (field === 'discount' || field === 'groupbuy' || field === 'indoorMap') {
            // 布尔值转为0/1
            rowData[field] = poi[field] ? 1 : 0
          } else {
            // 其他字段直接复制
            rowData[field] = poi[field] !== undefined ? poi[field] : ''
          }
        })
        
        return rowData
      })
      
      // 创建工作表
      const worksheet = XLSX.utils.json_to_sheet(exportData, { 
        header: exportFields 
      })
      
      // 添加自定义表头
      // 在A1单元格前添加一个空行用于放置自定义表头
      XLSX.utils.sheet_add_aoa(worksheet, [
        exportFields.map(field => headerMap[field])
      ], { origin: "A1" })
      
      // 添加工作表到工作簿
      XLSX.utils.book_append_sheet(workbook, worksheet, 'Sheet1')
      
      // 设置列宽
      const columnWidths = exportFields.map(field => {
        if (['address', 'photos', 'location'].includes(field)) {
          return { wch: 60 } // 较长字段，增加宽度
        } else if (['name', 'type', 'BigCategory', 'MidCategory', 'SubCategory'].includes(field)) {
          return { wch: 25 } // 中等长度字段
        } else {
          return { wch: 15 } // 默认宽度
        }
      })
      
      worksheet['!cols'] = columnWidths
      
      // 导出Excel文件
      XLSX.writeFile(workbook, 'POI数据.xlsx')
      
      ElMessage.success('导出成功')
    } catch (error) {
      console.error('导出Excel失败:', error)
      ElMessage.error('导出Excel失败: ' + error.message)
    }
  }
  
  // 尝试将字符串解析为JSON，如果失败则返回原始字符串
  const parseJsonIfPossible = (str) => {
    if (typeof str !== 'string') {
      return str
    }
    
    try {
      return JSON.parse(str)
    } catch (e) {
      console.error('JSON解析失败:', e)
      return str
    }
  }
  
  // 显示POI详情
  const showPoiDetail = async (id) => {
    poiDetailVisible.value = true
    poiDetailLoading.value = true
    poiDetail.value = null
    poiDetailJson.value = ''
    
    try {
      // 替换为您的实际高德地图API Key
      const apiKey = '53269938ccac7faafdd3737e06a0b4f8' // 实际使用时应该换成真实的key
      const response = await axios.get(`https://restapi.amap.com/v5/place/detail`, {
        params: {
          id: id,
          show_fields: 'business,children,indoor,photos,navi',
          key: apiKey
        }
      })
      
      console.log('高德地图API返回数据:', response.data)
      
      if (response.data.status === '1' && response.data.pois && response.data.pois.length > 0) {
        poiDetail.value = response.data.pois[0]
        poiDetailJson.value = JSON.stringify(response.data, null, 2)
      } else {
        ElMessage.warning('未找到该POI的详细信息')
      }
    } catch (error) {
      console.error('获取POI详情失败:', error)
      ElMessage.error('获取POI详情失败: ' + (error.message || '未知错误'))
    } finally {
      poiDetailLoading.value = false
    }
  }
  
  // 批量删除选中的POI数据
  const batchDeletePois = () => {
    if (multipleSelection.value.length === 0) {
      ElMessage.warning('请先选择要删除的数据')
      return
    }
    
    ElMessageBox.confirm(`确定要删除选中的 ${multipleSelection.value.length} 条POI数据吗?`, '批量删除', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      try {
        // 获取所有选中项的ID
        const selectedIds = multipleSelection.value.map(item => item.id)
        
        // 过滤掉选中的数据
        const newPoisData = allPoisData.value.filter(item => !selectedIds.includes(item.id))
        
        // 保存到localStorage
        localStorage.setItem('amapPois', JSON.stringify(newPoisData))
        
        // 提示成功
        ElMessage.success(`成功删除 ${multipleSelection.value.length} 条数据`)
        
        // 清空选择
        multipleSelection.value = []
        
        // 重新加载数据
        loadPoisData()
      } catch (error) {
        console.error('批量删除POI数据失败:', error)
        ElMessage.error('批量删除POI数据失败: ' + error.message)
      }
    })
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
    background-color: var(--el-bg-color);
    padding: 16px;
    box-sizing: border-box;
    overflow-y: auto; /* 允许垂直滚动 */
  }
  
  .gva-search-box {
    margin-bottom: 16px;
    background-color: var(--el-bg-color-overlay);
    padding: 16px;
    border-radius: 4px;
    box-shadow: var(--el-box-shadow-light);
  }
  
  // .content-main {
  //   flex: 1;
  //   display: flex;
  //   flex-direction: column;
  //   background-color: var(--el-bg-color-overlay);
  //   border-radius: 4px;
  //   box-shadow: var(--el-box-shadow-light);
  //   padding: 16px;
  //   height: calc(100vh - 150px);
  //   overflow: hidden; /* 防止内容溢出 */
  // }

  .content-main{
    max-height: 90vh;
    overflow-y: auto;
  }
  
  .gva-btn-list {
    margin-bottom: 16px;
    display: flex;
    gap: 8px;
    align-items: center;
  }
  
  .ml-2 {
    margin-left: 8px;
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
    width: 8px;
    height: 8px;
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
  
  /* POI详情样式 */
  .poi-detail-container {
    padding: 0 10px;
    max-height: 70vh;
    overflow-y: auto;
  }
  
  .poi-detail-header {
    margin-bottom: 20px;
    
    h2 {
      margin-top: 0;
      margin-bottom: 10px;
      color: var(--el-color-primary);
    }
  }
  
  .poi-section {
    margin-bottom: 20px;
    
    h3 {
      margin-top: 5px;
      margin-bottom: 10px;
      font-size: 16px;
      color: var(--el-color-primary-light-3);
      border-bottom: 1px solid var(--el-border-color-light);
      padding-bottom: 5px;
    }
  }
  
  .poi-basic-info,
  .poi-business-info,
  .poi-navi-info {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 10px;
    
    div {
      margin-bottom: 5px;
    }
  }
  
  .info-label {
    font-weight: bold;
    color: var(--el-text-color-secondary);
    margin-right: 5px;
  }
  
  .poi-photos {
    display: flex;
    flex-wrap: wrap;
  }
  
  .id-link {
    color: var(--el-color-primary);
    text-decoration: none;
    cursor: pointer;
    
    &:hover {
      text-decoration: underline;
    }
  }
  
  /* JSON查看器样式 */
  .json-viewer-container {
    max-height: 400px;
    overflow-y: auto;
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
  
  /* 添加全局滚动条样式 */
  /* html, body {
    overflow-x: hidden;
  } */
  
  .amaplist-container {
    overflow-y: auto;
    scrollbar-width: thin;
    scrollbar-color: #888 #f1f1f1;
  }
  </style> 