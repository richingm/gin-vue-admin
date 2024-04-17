<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item label="工作项id" prop="workitemId">
            
             <el-input v-model.number="searchInfo.workitemId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
         <el-table-column align="left" label="借用结束时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.borrowEndTime) }}</template>
         </el-table-column>
         <el-table-column align="left" label="借用开始时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.borrowStartTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="邮箱" prop="email" width="120" />
        <el-table-column align="left" label="预估人力成本" prop="estimatedCost" width="120" />
        <el-table-column align="left" label="预估人天" prop="estimatedDays" width="120" />
        <el-table-column align="left" label="是否结束借用[1:结束,0:不结束]" prop="isEstimatedFinished" width="120" />
        <el-table-column align="left" label="岗位" prop="position" width="120" />
        <el-table-column align="left" label="能力等级" prop="positionLevel" width="120" />
        <el-table-column align="left" label="项目角色ID" prop="roleId" width="120" />
        <el-table-column align="left" label="项目角色名称" prop="roleName" width="120" />
        <el-table-column align="left" label="结算人天" prop="settleDays" width="120" />
        <el-table-column align="left" label="正在结算的天数" prop="settleLockDays" width="120" />
        <el-table-column align="left" label="结算人力成本" prop="settledCost" width="120" />
        <el-table-column align="left" label="已经结算的天数" prop="settledDays" width="120" />
        <el-table-column align="left" label="用户id" prop="userId" width="120" />
        <el-table-column align="left" label="用户名" prop="userName" width="120" />
        <el-table-column align="left" label="工作项id" prop="workitemId" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateWorkitemPersonDetailFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-drawer size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #title>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="借用结束时间:"  prop="borrowEndTime" >
              <el-date-picker v-model="formData.borrowEndTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="借用开始时间:"  prop="borrowStartTime" >
              <el-date-picker v-model="formData.borrowStartTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="邮箱:"  prop="email" >
              <el-input v-model="formData.email" :clearable="true"  placeholder="请输入邮箱" />
            </el-form-item>
            <el-form-item label="预估人力成本:"  prop="estimatedCost" >
              <el-input v-model.number="formData.estimatedCost" :clearable="true" placeholder="请输入预估人力成本" />
            </el-form-item>
            <el-form-item label="预估人天:"  prop="estimatedDays" >
              <el-input v-model.number="formData.estimatedDays" :clearable="true" placeholder="请输入预估人天" />
            </el-form-item>
            <el-form-item label="是否结束借用[1:结束,0:不结束]:"  prop="isEstimatedFinished" >
              <el-input v-model.number="formData.isEstimatedFinished" :clearable="true" placeholder="请输入是否结束借用[1:结束,0:不结束]" />
            </el-form-item>
            <el-form-item label="岗位:"  prop="position" >
              <el-input v-model="formData.position" :clearable="true"  placeholder="请输入岗位" />
            </el-form-item>
            <el-form-item label="能力等级:"  prop="positionLevel" >
              <el-input v-model="formData.positionLevel" :clearable="true"  placeholder="请输入能力等级" />
            </el-form-item>
            <el-form-item label="项目角色ID:"  prop="roleId" >
              <el-input v-model.number="formData.roleId" :clearable="true" placeholder="请输入项目角色ID" />
            </el-form-item>
            <el-form-item label="项目角色名称:"  prop="roleName" >
              <el-input v-model="formData.roleName" :clearable="true"  placeholder="请输入项目角色名称" />
            </el-form-item>
            <el-form-item label="结算人天:"  prop="settleDays" >
              <el-input v-model.number="formData.settleDays" :clearable="true" placeholder="请输入结算人天" />
            </el-form-item>
            <el-form-item label="正在结算的天数:"  prop="settleLockDays" >
              <el-input v-model.number="formData.settleLockDays" :clearable="true" placeholder="请输入正在结算的天数" />
            </el-form-item>
            <el-form-item label="结算人力成本:"  prop="settledCost" >
              <el-input v-model.number="formData.settledCost" :clearable="true" placeholder="请输入结算人力成本" />
            </el-form-item>
            <el-form-item label="已经结算的天数:"  prop="settledDays" >
              <el-input v-model.number="formData.settledDays" :clearable="true" placeholder="请输入已经结算的天数" />
            </el-form-item>
            <el-form-item label="用户id:"  prop="userId" >
              <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入用户id" />
            </el-form-item>
            <el-form-item label="用户名:"  prop="userName" >
              <el-input v-model="formData.userName" :clearable="true"  placeholder="请输入用户名" />
            </el-form-item>
            <el-form-item label="工作项id:"  prop="workitemId" >
              <el-input v-model.number="formData.workitemId" :clearable="true" placeholder="请输入工作项id" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer size="800" v-model="detailShow" :before-close="closeDetailShow" title="查看详情" destroy-on-close>
          <template #title>
             <div class="flex justify-between items-center">
               <span class="text-lg">查看详情</span>
             </div>
         </template>
        <el-descriptions :column="1" border>
                <el-descriptions-item label="借用结束时间">
                      {{ formatDate(formData.borrowEndTime) }}
                </el-descriptions-item>
                <el-descriptions-item label="借用开始时间">
                      {{ formatDate(formData.borrowStartTime) }}
                </el-descriptions-item>
                <el-descriptions-item label="邮箱">
                        {{ formData.email }}
                </el-descriptions-item>
                <el-descriptions-item label="预估人力成本">
                        {{ formData.estimatedCost }}
                </el-descriptions-item>
                <el-descriptions-item label="预估人天">
                        {{ formData.estimatedDays }}
                </el-descriptions-item>
                <el-descriptions-item label="是否结束借用[1:结束,0:不结束]">
                        {{ formData.isEstimatedFinished }}
                </el-descriptions-item>
                <el-descriptions-item label="岗位">
                        {{ formData.position }}
                </el-descriptions-item>
                <el-descriptions-item label="能力等级">
                        {{ formData.positionLevel }}
                </el-descriptions-item>
                <el-descriptions-item label="项目角色ID">
                        {{ formData.roleId }}
                </el-descriptions-item>
                <el-descriptions-item label="项目角色名称">
                        {{ formData.roleName }}
                </el-descriptions-item>
                <el-descriptions-item label="结算人天">
                        {{ formData.settleDays }}
                </el-descriptions-item>
                <el-descriptions-item label="正在结算的天数">
                        {{ formData.settleLockDays }}
                </el-descriptions-item>
                <el-descriptions-item label="结算人力成本">
                        {{ formData.settledCost }}
                </el-descriptions-item>
                <el-descriptions-item label="已经结算的天数">
                        {{ formData.settledDays }}
                </el-descriptions-item>
                <el-descriptions-item label="用户id">
                        {{ formData.userId }}
                </el-descriptions-item>
                <el-descriptions-item label="用户名">
                        {{ formData.userName }}
                </el-descriptions-item>
                <el-descriptions-item label="工作项id">
                        {{ formData.workitemId }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createWorkitemPersonDetail,
  deleteWorkitemPersonDetail,
  deleteWorkitemPersonDetailByIds,
  updateWorkitemPersonDetail,
  findWorkitemPersonDetail,
  getWorkitemPersonDetailList
} from '@/api/workitem_detail/workitemPersonDetail'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'WorkitemPersonDetail'
})

// 自动化生成的字典（可能为空）以及字段
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
               },
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

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
    pageSize.value = 10
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
  const table = await getWorkitemPersonDetailList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteWorkitemPersonDetailFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteWorkitemPersonDetailByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateWorkitemPersonDetailFunc = async(row) => {
    const res = await findWorkitemPersonDetail({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reworkitemPersonDetail
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWorkitemPersonDetailFunc = async (row) => {
    const res = await deleteWorkitemPersonDetail({ ID: row.ID })
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


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findWorkitemPersonDetail({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reworkitemPersonDetail
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
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
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
