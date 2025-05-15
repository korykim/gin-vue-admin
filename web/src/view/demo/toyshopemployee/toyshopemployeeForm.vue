
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="姓名:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入姓名" />
</el-form-item>
        <el-form-item label="职位:" prop="position">
    <el-input v-model="formData.position" :clearable="false" placeholder="请输入职位" />
</el-form-item>
        <el-form-item label="电话:" prop="phone">
    <el-input v-model="formData.phone" :clearable="false" placeholder="请输入电话" />
</el-form-item>
        <el-form-item label="邮箱:" prop="email">
    <el-input v-model="formData.email" :clearable="false" placeholder="请输入邮箱" />
</el-form-item>
        <el-form-item label="入职日期:" prop="startDate">
    <el-date-picker v-model="formData.startDate" type="date" style="width:100%" placeholder="选择日期" :clearable="false" />
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
  createEmployeeProfile,
  updateEmployeeProfile,
  findEmployeeProfile
} from '@/api/demo/toyshopemployee'

defineOptions({
    name: 'EmployeeProfileForm'
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
            name: '',
            position: '',
            phone: '',
            email: '',
            startDate: new Date(),
        })
// 验证规则
const rule = reactive({
               name : [{
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
      const res = await findEmployeeProfile({ ID: route.query.id })
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
               res = await createEmployeeProfile(formData.value)
               break
             case 'update':
               res = await updateEmployeeProfile(formData.value)
               break
             default:
               res = await createEmployeeProfile(formData.value)
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
