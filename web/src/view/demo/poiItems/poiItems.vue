<template>
  <div class="amaplist-container">
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
            <el-form-item label="POI名称" prop="name">
  <el-input v-model="searchInfo.name" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="详细地址" prop="address">
  <el-input v-model="searchInfo.address" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="联系电话" prop="tel">
  <el-input v-model="searchInfo.tel" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="省份名称" prop="pname">
  <el-input v-model="searchInfo.pname" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="城市名称" prop="cityname">
  <el-input v-model="searchInfo.cityname" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="区域名称" prop="adname">
  <el-input v-model="searchInfo.adname" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="POI类型编码" prop="poiType">
  <el-input v-model="searchInfo.poiType" placeholder="搜索条件" />
</el-form-item>

<el-form-item label="大类别" prop="BigCategory">
  <el-input v-model="searchInfo.BigCategory" placeholder="搜索条件" />
</el-form-item>

<el-form-item label="中类别" prop="MidCategory">
  <el-input v-model="searchInfo.MidCategory" placeholder="搜索条件" />
</el-form-item>

<el-form-item label="子类别" prop="SubCategory">
  <el-input v-model="searchInfo.SubCategory" placeholder="搜索条件" />
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
            <ExportTemplate  template-id="demo_PoiItems" />
            <ExportExcel  template-id="demo_PoiItems" filterDeleted/>
            <ImportExcel  template-id="demo_PoiItems" @on-success="getTableData" />
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
        
            <el-table-column align="left" label="POI唯一标识" prop="id" width="120" />

            <el-table-column align="left" label="POI名称" prop="name" width="120" />

            <el-table-column align="left" label="类型描述" prop="type" width="120" />

            <el-table-column align="left" label="详细地址" prop="address" width="120" />

            <el-table-column align="left" label="联系电话" prop="tel" width="120" />

            <el-table-column sortable align="left" label="省份名称" prop="pname" width="120" />

            <el-table-column sortable align="left" label="城市名称" prop="cityname" width="120" />

            <el-table-column sortable align="left" label="区域名称" prop="adname" width="120" />

            <el-table-column label="照片信息" prop="photos" width="200">
   <template #default="scope">
      <div class="multiple-img-box">
         <template v-if="scope.row.photos && scope.row.photos.length > 0">
             <el-image 
                preview-teleported 
                v-for="(item,index) in scope.row.photos" 
                :key="index" 
                style="width: 80px; height: 80px" 
                :src="getUrl(item.url || item)" 
                fit="cover"
                :preview-src-list="getPreviewList(scope.row.photos)"
                :initial-index="index"
             />
         </template>
         <span v-else class="no-photo">暂无照片</span>
     </div>
   </template>
</el-table-column>
            <el-table-column align="left" label="POI类型编码" prop="poiType" width="120" />

            <el-table-column sortable align="left" label="大类别" prop="BigCategory" width="120" />

            <el-table-column sortable align="left" label="中类别" prop="MidCategory" width="120" />

            <el-table-column sortable align="left" label="子类别" prop="SubCategory" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updatePoiItemsFunc(scope.row)">编辑</el-button>
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
            <el-form-item label="POI唯一标识:" prop="id">
    <el-input v-model="formData.id" :clearable="true" placeholder="请输入POI唯一标识" />
</el-form-item>
            <el-form-item label="POI名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入POI名称" />
</el-form-item>
            <el-form-item label="类型描述:" prop="type">
    <el-input v-model="formData.type" :clearable="true" placeholder="请输入类型描述" />
</el-form-item>
            <el-form-item label="位置坐标[经度,纬度]:" prop="location">
    // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.location 后端会按照json的类型进行存取
    {{ formData.location }}
</el-form-item>
            <el-form-item label="详细地址:" prop="address">
    <el-input v-model="formData.address" :clearable="true" placeholder="请输入详细地址" />
</el-form-item>
            <el-form-item label="联系电话:" prop="tel">
    <el-input v-model="formData.tel" :clearable="true" placeholder="请输入联系电话" />
</el-form-item>
            <el-form-item label="距离(米):" prop="distance">
    <el-input v-model.number="formData.distance" :clearable="true" placeholder="请输入距离(米)" />
</el-form-item>
            <el-form-item label="商铺信息:" prop="shopinfo">
    <el-input v-model="formData.shopinfo" :clearable="true" placeholder="请输入商铺信息" />
</el-form-item>
            <el-form-item label="网站:" prop="website">
    <el-input v-model="formData.website" :clearable="true" placeholder="请输入网站" />
</el-form-item>
            <el-form-item label="省份编码:" prop="pcode">
    <el-input v-model="formData.pcode" :clearable="true" placeholder="请输入省份编码" />
</el-form-item>
            <el-form-item label="城市编码:" prop="citycode">
    <el-input v-model="formData.citycode" :clearable="true" placeholder="请输入城市编码" />
</el-form-item>
            <el-form-item label="区域编码:" prop="adcode">
    <el-input v-model="formData.adcode" :clearable="true" placeholder="请输入区域编码" />
</el-form-item>
            <el-form-item label="邮政编码:" prop="postcode">
    <el-input v-model="formData.postcode" :clearable="true" placeholder="请输入邮政编码" />
</el-form-item>
            <el-form-item label="省份名称:" prop="pname">
    <el-input v-model="formData.pname" :clearable="true" placeholder="请输入省份名称" />
</el-form-item>
            <el-form-item label="城市名称:" prop="cityname">
    <el-input v-model="formData.cityname" :clearable="true" placeholder="请输入城市名称" />
</el-form-item>
            <el-form-item label="区域名称:" prop="adname">
    <el-input v-model="formData.adname" :clearable="true" placeholder="请输入区域名称" />
</el-form-item>
            <el-form-item label="电子邮箱:" prop="email">
    <el-input v-model="formData.email" :clearable="true" placeholder="请输入电子邮箱" />
</el-form-item>
            <el-form-item label="照片信息:" prop="photos">
    <SelectImage
     multiple
     v-model="formData.photos"
     file-type="image"
     />
</el-form-item>
            <el-form-item label="入口位置:" prop="entrLocation">
    <el-input v-model="formData.entrLocation" :clearable="true" placeholder="请输入入口位置" />
</el-form-item>
            <el-form-item label="出口位置:" prop="exitLocation">
    <el-input v-model="formData.exitLocation" :clearable="true" placeholder="请输入出口位置" />
</el-form-item>
            <el-form-item label="是否支持团购:" prop="groupbuy">
    <el-switch v-model="formData.groupbuy" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="是否有优惠:" prop="discount">
    <el-switch v-model="formData.discount" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="是否有室内地图:" prop="indoorMap">
    <el-switch v-model="formData.indoorMap" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="评分:" prop="rating">
    <el-input-number v-model="formData.rating" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="消费金额:" prop="cost">
    <el-input-number v-model="formData.cost" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <!-- <el-form-item label="索引:" prop="Idx">
    <el-input v-model.number="formData.Idx" :clearable="true" placeholder="请输入索引" />
</el-form-item>
            <el-form-item label="序号:" prop="index">
    <el-input v-model.number="formData.poiIndex" :clearable="true" placeholder="请输入序号" />
</el-form-item> -->
            <el-form-item label="POI类型编码:" prop="poiType">
    <el-input v-model="formData.poiType" :clearable="true" placeholder="请输入POI类型编码" />
</el-form-item>
            <el-form-item label="大类别:" prop="BigCategory">
    <el-input v-model="formData.BigCategory" :clearable="true" placeholder="请输入大类别" />
</el-form-item>
            <el-form-item label="中类别:" prop="MidCategory">
    <el-input v-model="formData.MidCategory" :clearable="true" placeholder="请输入中类别" />
</el-form-item>
            <el-form-item label="子类别:" prop="SubCategory">
    <el-input v-model="formData.SubCategory" :clearable="true" placeholder="请输入子类别" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="POI唯一标识">
    {{ detailFrom.id }}
</el-descriptions-item>
                    <el-descriptions-item label="POI名称">
    {{ detailFrom.name }}
</el-descriptions-item>
                    <el-descriptions-item label="类型描述">
    {{ detailFrom.type }}
</el-descriptions-item>
                    <el-descriptions-item label="位置坐标[经度,纬度]">
    {{ detailFrom.location }}
</el-descriptions-item>
                    <el-descriptions-item label="详细地址">
    {{ detailFrom.address }}
</el-descriptions-item>
                    <el-descriptions-item label="联系电话">
    {{ detailFrom.tel }}
</el-descriptions-item>
                    <el-descriptions-item label="距离(米)">
    {{ detailFrom.distance }}
</el-descriptions-item>
                    <el-descriptions-item label="商铺信息">
    {{ detailFrom.shopinfo }}
</el-descriptions-item>
                    <el-descriptions-item label="网站">
    {{ detailFrom.website }}
</el-descriptions-item>
                    <el-descriptions-item label="省份编码">
    {{ detailFrom.pcode }}
</el-descriptions-item>
                    <el-descriptions-item label="城市编码">
    {{ detailFrom.citycode }}
</el-descriptions-item>
                    <el-descriptions-item label="区域编码">
    {{ detailFrom.adcode }}
</el-descriptions-item>
                    <el-descriptions-item label="邮政编码">
    {{ detailFrom.postcode }}
</el-descriptions-item>
                    <el-descriptions-item label="省份名称">
    {{ detailFrom.pname }}
</el-descriptions-item>
                    <el-descriptions-item label="城市名称">
    {{ detailFrom.cityname }}
</el-descriptions-item>
                    <el-descriptions-item label="区域名称">
    {{ detailFrom.adname }}
</el-descriptions-item>
                    <el-descriptions-item label="电子邮箱">
    {{ detailFrom.email }}
</el-descriptions-item>
                    <el-descriptions-item label="照片信息">
    <template v-if="detailFrom.photos && detailFrom.photos.length > 0">
        <el-image 
            style="width: 50px; height: 50px; margin-right: 10px" 
            :preview-src-list="getPreviewList(detailFrom.photos)" 
            :initial-index="index" 
            v-for="(item,index) in detailFrom.photos" 
            :key="index" 
            :src="getUrl(item.url || item)" 
            fit="cover" 
        />
    </template>
    <span v-else class="no-photo">暂无照片</span>
</el-descriptions-item>
                    <el-descriptions-item label="入口位置">
    {{ detailFrom.entrLocation }}
</el-descriptions-item>
                    <el-descriptions-item label="出口位置">
    {{ detailFrom.exitLocation }}
</el-descriptions-item>
                    <el-descriptions-item label="是否支持团购">
    {{ detailFrom.groupbuy }}
</el-descriptions-item>
                    <el-descriptions-item label="是否有优惠">
    {{ detailFrom.discount }}
</el-descriptions-item>
                    <el-descriptions-item label="是否有室内地图">
    {{ detailFrom.indoorMap }}
</el-descriptions-item>
                    <el-descriptions-item label="评分">
    {{ detailFrom.rating }}
</el-descriptions-item>
                    <el-descriptions-item label="消费金额">
    {{ detailFrom.cost }}
</el-descriptions-item>
                    <!-- <el-descriptions-item label="索引">
    {{ detailFrom.Idx }}
</el-descriptions-item>
                    <el-descriptions-item label="序号">
    {{ detailFrom.poiIndex }}
</el-descriptions-item> -->
                    <el-descriptions-item label="POI类型编码">
    {{ detailFrom.poiType }}
</el-descriptions-item>
                    <el-descriptions-item label="大类别">
    {{ detailFrom.BigCategory }}
</el-descriptions-item>
                    <el-descriptions-item label="中类别">
    {{ detailFrom.MidCategory }}
</el-descriptions-item>
                    <el-descriptions-item label="子类别">
    {{ detailFrom.SubCategory }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createPoiItems,
  deletePoiItems,
  deletePoiItemsByIds,
  updatePoiItems,
  findPoiItems,
  getPoiItemsList
} from '@/api/demo/poiItems'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

// 全量引入格式化工具 请按需保留
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
    name: 'PoiItems'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            id: '',
            name: '',
            type: '',
            location: {},
            address: '',
            tel: '',
            distance: undefined,
            shopinfo: '',
            website: '',
            pcode: '',
            citycode: '',
            adcode: '',
            postcode: '',
            pname: '',
            cityname: '',
            adname: '',
            email: '',
            photos: [],
            entrLocation: '',
            exitLocation: '',
            groupbuy: false,
            discount: false,
            indoorMap: false,
            rating: 0,
            cost: 0,
            //Idx: undefined,
            //poiIndex: undefined,
            poiType: '',
            BigCategory: '',
            MidCategory: '',
            SubCategory: '',
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
            pname: 'pname',
            cityname: 'cityname',
            adname: 'adname',
            BigCategory: 'BigCategory',
            MidCategory: 'MidCategory',
            SubCategory: 'SubCategory',
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
    if (searchInfo.value.groupbuy === ""){
        searchInfo.value.groupbuy=null
    }
    if (searchInfo.value.discount === ""){
        searchInfo.value.discount=null
    }
    if (searchInfo.value.indoorMap === ""){
        searchInfo.value.indoorMap=null
    }
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
  const table = await getPoiItemsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deletePoiItemsFunc(row)
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
      const res = await deletePoiItemsByIds({ ids })
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
const updatePoiItemsFunc = async(row) => {
    const res = await findPoiItems({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deletePoiItemsFunc = async (row) => {
    const res = await deletePoiItems({ id: row.id })
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
        id: '',
        name: '',
        type: '',
        location: {},
        address: '',
        tel: '',
        distance: undefined,
        shopinfo: '',
        website: '',
        pcode: '',
        citycode: '',
        adcode: '',
        postcode: '',
        pname: '',
        cityname: '',
        adname: '',
        email: '',
        photos: [],
        entrLocation: '',
        exitLocation: '',
        groupbuy: false,
        discount: false,
        indoorMap: false,
        rating: 0,
        cost: 0,
        Idx: undefined,
        poiIndex: undefined,
        poiType: '',
        BigCategory: '',
        MidCategory: '',
        SubCategory: '',
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
                  res = await createPoiItems(formData.value)
                  break
                case 'update':
                  res = await updatePoiItems(formData.value)
                  break
                default:
                  res = await createPoiItems(formData.value)
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
  const res = await findPoiItems({ id: row.id })
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

// 获取预览图片列表
const getPreviewList = (photos) => {
  if (!photos || !Array.isArray(photos)) return []
  return photos.map(item => {
    if (typeof item === 'string') return getUrl(item)
    return getUrl(item.url || '')
  })
}

</script>

<style lang="scss" scoped>

  
.no-photo {
  color: #909399;
  font-size: 14px;
  font-style: italic;
}

.gva-table-box {
  max-height: 80vh;
  overflow-y: auto;
}
</style>
