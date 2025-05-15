
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名字:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入名字" />
</el-form-item>
        <el-form-item label="年龄:" prop="age">
    <el-input v-model.number="formData.age" :clearable="false" placeholder="请输入年龄" />
</el-form-item>
        <el-form-item label="能力:" prop="power">
    <el-input v-model="formData.power" :clearable="false" placeholder="请输入能力" />
</el-form-item>
        <el-form-item label="起源:" prop="origin">
    <el-input v-model="formData.origin" :clearable="false" placeholder="请输入起源" />
</el-form-item>
        <el-form-item label="图片:" prop="image">
    <el-input v-model="formData.image" :clearable="false" placeholder="请输入图片" />
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
  createHero,
  updateHero,
  findHero
} from '@/api/demo/hero'

defineOptions({
    name: 'HeroForm'
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
            age: 0,
            power: '',
            origin: '',
            image: '',
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
      const res = await findHero({ ID: route.query.id })
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
               res = await createHero(formData.value)
               break
             case 'update':
               res = await updateHero(formData.value)
               break
             default:
               res = await createHero(formData.value)
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
