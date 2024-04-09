<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="知识库:" prop="knowledgeId">
          <el-input v-model.number="formData.knowledgeId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="父id:" prop="pid">
          <el-input v-model.number="formData.pid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入标题" />
       </el-form-item>
        <el-form-item label="重要程度:" prop="importanceLevel">
           <el-select v-model="formData.importanceLevel" placeholder="请选择重要程度" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in import_levelOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="理解程度:" prop="understandLevel">
           <el-select v-model="formData.understandLevel" placeholder="请选择理解程度" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in understand_levelOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="最后查看时间:" prop="lastViewedAt">
          <el-date-picker v-model="formData.lastViewedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="内容:" prop="content">
          <RichEdit v-model="formData.content"/>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createArticles,
  updateArticles,
  findArticles
} from '@/api/article/articles'

defineOptions({
    name: 'ArticlesForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const import_levelOptions = ref([])
const understand_levelOptions = ref([])
const formData = ref({
            knowledgeId: 0,
            pid: 0,
            title: '',
            importanceLevel: '',
            understandLevel: '',
            lastViewedAt: new Date(),
            content: '',
        })
// 验证规则
const rule = reactive({
               knowledgeId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               importanceLevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               understandLevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               content : [{
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
      const res = await findArticles({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rearticles
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    import_levelOptions.value = await getDictFunc('import_level')
    understand_levelOptions.value = await getDictFunc('understand_level')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createArticles(formData.value)
               break
             case 'update':
               res = await updateArticles(formData.value)
               break
             default:
               res = await createArticles(formData.value)
               break
           }
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
