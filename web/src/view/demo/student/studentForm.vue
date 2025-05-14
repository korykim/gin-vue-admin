
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名字:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入名字" />
</el-form-item>
        <el-form-item label="年龄:" prop="age">
    <el-input v-model.number="formData.age" :clearable="true" placeholder="请输入年龄" />
</el-form-item>
        <el-form-item label="性别:" prop="sex">
    <el-select v-model="formData.sex" placeholder="请选择性别" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="简介:" prop="profile">
    <SelectFile v-model="formData.profile" />
</el-form-item>
        <el-form-item label="头像:" prop="picture">
    <SelectImage
     v-model="formData.picture"
     file-type="image"
    />
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
  createStudent,
  updateStudent,
  findStudent
} from '@/api/demo/student'

defineOptions({
    name: 'StudentForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const genderOptions = ref([])
const formData = ref({
            name: '',
            age: undefined,
            sex: '',
            profile: [],
            picture: "",
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
      const res = await findStudent({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    genderOptions.value = await getDictFunc('gender')
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
               res = await createStudent(formData.value)
               break
             case 'update':
               res = await updateStudent(formData.value)
               break
             default:
               res = await createStudent(formData.value)
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
