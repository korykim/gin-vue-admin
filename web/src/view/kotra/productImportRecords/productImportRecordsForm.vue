
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
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
  createProductImportRecords,
  updateProductImportRecords,
  findProductImportRecords
} from '@/api/kotra/productImportRecords'

defineOptions({
    name: 'ProductImportRecordsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findProductImportRecords({ ID: route.query.id })
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
