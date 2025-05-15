
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="店名:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入店名" />
</el-form-item>
        <el-form-item label="地址:" prop="address">
    <el-input v-model="formData.address" :clearable="false" placeholder="请输入地址" />
</el-form-item>
        <el-form-item label="电话:" prop="phone">
    <el-input v-model="formData.phone" :clearable="false" placeholder="请输入电话" />
</el-form-item>
        <el-form-item label="营业时间:" prop="openHours">
    <el-input v-model="formData.openHours" :clearable="false" placeholder="请输入营业时间" />
</el-form-item>
        <el-form-item label="店主:" prop="owner">
    <el-input v-model="formData.owner" :clearable="false" placeholder="请输入店主" />
</el-form-item>
        <el-form-item label="描述:" prop="description">
    <el-input v-model="formData.description" :clearable="false" placeholder="请输入描述" />
</el-form-item>
        <el-form-item label="LOGO:" prop="logo">
    <SelectImage
     v-model="formData.logo"
     file-type="image"
    />
</el-form-item>
        <el-form-item label="管理员:" prop="manager">
    <el-select multiple v-model="formData.manager" placeholder="请选择管理员" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in dataSource.manager" :key="key" :label="item.label" :value="item.value" />
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
    getToyStoreDataSource,
  createToyStore,
  updateToyStore,
  findToyStore
} from '@/api/demo/toyStore'

defineOptions({
    name: 'ToyStoreForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 数组控制组件
import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            name: '',
            address: '',
            phone: '',
            openHours: '',
            owner: '',
            description: '',
            logo: "",
            manager: [],
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               address : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getToyStoreDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findToyStore({ ID: route.query.id })
      if (res.code === 0) {
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
               res = await createToyStore(formData.value)
               break
             case 'update':
               res = await updateToyStore(formData.value)
               break
             default:
               res = await createToyStore(formData.value)
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
