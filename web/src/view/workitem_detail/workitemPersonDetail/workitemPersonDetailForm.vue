<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="借用结束时间:" prop="borrowEndTime">
          <el-date-picker v-model="formData.borrowEndTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="借用开始时间:" prop="borrowStartTime">
          <el-date-picker v-model="formData.borrowStartTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true"  placeholder="请输入邮箱" />
       </el-form-item>
        <el-form-item label="预估人力成本:" prop="estimatedCost">
          <el-input v-model.number="formData.estimatedCost" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="预估人天:" prop="estimatedDays">
          <el-input v-model.number="formData.estimatedDays" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否结束借用[1:结束,0:不结束]:" prop="isEstimatedFinished">
          <el-input v-model.number="formData.isEstimatedFinished" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="岗位:" prop="position">
          <el-input v-model="formData.position" :clearable="true"  placeholder="请输入岗位" />
       </el-form-item>
        <el-form-item label="能力等级:" prop="positionLevel">
          <el-input v-model="formData.positionLevel" :clearable="true"  placeholder="请输入能力等级" />
       </el-form-item>
        <el-form-item label="项目角色ID:" prop="roleId">
          <el-input v-model.number="formData.roleId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="项目角色名称:" prop="roleName">
          <el-input v-model="formData.roleName" :clearable="true"  placeholder="请输入项目角色名称" />
       </el-form-item>
        <el-form-item label="结算人天:" prop="settleDays">
          <el-input v-model.number="formData.settleDays" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="正在结算的天数:" prop="settleLockDays">
          <el-input v-model.number="formData.settleLockDays" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="结算人力成本:" prop="settledCost">
          <el-input v-model.number="formData.settledCost" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="已经结算的天数:" prop="settledDays">
          <el-input v-model.number="formData.settledDays" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户名:" prop="userName">
          <el-input v-model="formData.userName" :clearable="true"  placeholder="请输入用户名" />
       </el-form-item>
        <el-form-item label="工作项id:" prop="workitemId">
          <el-input v-model.number="formData.workitemId" :clearable="true" placeholder="请输入" />
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
  createWorkitemPersonDetail,
  updateWorkitemPersonDetail,
  findWorkitemPersonDetail
} from '@/api/workitem_detail/workitemPersonDetail'

defineOptions({
    name: 'WorkitemPersonDetailForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            borrowEndTime: new Date(),
            borrowStartTime: new Date(),
            email: '',
            estimatedCost: 0,
            estimatedDays: 0,
            isEstimatedFinished: 0,
            position: '',
            positionLevel: '',
            roleId: 0,
            roleName: '',
            settleDays: 0,
            settleLockDays: 0,
            settledCost: 0,
            settledDays: 0,
            userId: 0,
            userName: '',
            workitemId: 0,
        })
// 验证规则
const rule = reactive({
               workitemId : [{
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
      const res = await findWorkitemPersonDetail({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reworkitemPersonDetail
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createWorkitemPersonDetail(formData.value)
               break
             case 'update':
               res = await updateWorkitemPersonDetail(formData.value)
               break
             default:
               res = await createWorkitemPersonDetail(formData.value)
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
