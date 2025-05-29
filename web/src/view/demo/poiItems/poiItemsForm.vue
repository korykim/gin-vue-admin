<template>
    <div>
      <div class="gva-form-box">
        <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
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
    createPoiItems,
    updatePoiItems,
    findPoiItems
  } from '@/api/demo/poiItems'
  
  defineOptions({
      name: 'PoiItemsForm'
  })
  
  // 自动获取字典
  import { useRoute, useRouter } from "vue-router"
  import { ElMessage } from 'element-plus'
  import { ref, reactive } from 'vue'
  // 图片选择组件
  import SelectImage from '@/components/selectImage/selectImage.vue'
  
  const route = useRoute()
  const router = useRouter()
  
  // 提交按钮loading
  const btnLoading = ref(false)
  
  const type = ref('')
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
  
  // 初始化方法
  const init = async () => {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
      if (route.query.id) {
        const res = await findPoiItems({ ID: route.query.id })
        if (res.code === 0) {
          // 处理照片数据
          if (res.data.photos && Array.isArray(res.data.photos)) {
            res.data.photos = res.data.photos.map(item => {
              if (typeof item === 'string') return item
              return item.url || ''
            })
          }
          formData.value = res.data
          type.value = 'update'
        }
      } else {
        type.value = 'create'
      }
  }
  
  init()
  // 保存按钮
  const save = async() => {
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
  