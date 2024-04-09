<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="创建时间:" prop="createdAt">
          <el-date-picker v-model="formData.createdAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="删除时间:" prop="deletedAt">
          <el-date-picker v-model="formData.deletedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="ID:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="重要程度:" prop="importLevel">
           <el-select v-model="formData.importLevel" placeholder="请选择重要程度" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in import_levelOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="知识库名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入知识库名称" />
       </el-form-item>
        <el-form-item label="备注:" prop="notes">
          <el-input v-model="formData.notes" :clearable="true"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item label="父知识库:" prop="pid">
          <el-input v-model.number="formData.pid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新时间:" prop="updatedAt">
          <el-date-picker v-model="formData.updatedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createKnowledges,
  updateKnowledges,
  findKnowledges
} from '@/api/knowledge/knowledges'

defineOptions({
    name: 'KnowledgesForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const import_levelOptions = ref([])
const formData = ref({
            createdAt: new Date(),
            deletedAt: new Date(),
            id: 0,
            importLevel: '',
            name: '',
            notes: '',
            pid: 0,
            updatedAt: new Date(),
        })
// 验证规则
const rule = reactive({
               importLevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
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
      const res = await findKnowledges({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reknowledges
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    import_levelOptions.value = await getDictFunc('import_level')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createKnowledges(formData.value)
               break
             case 'update':
               res = await updateKnowledges(formData.value)
               break
             default:
               res = await createKnowledges(formData.value)
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
