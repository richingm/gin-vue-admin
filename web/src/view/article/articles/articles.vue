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
      
        <el-form-item label="知识库" prop="knowledgeId">
            
             <el-input v-model.number="searchInfo.knowledgeId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="标题" prop="title">
         <el-input v-model="searchInfo.title" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="重要程度" prop="importanceLevel">
            <el-select v-model="searchInfo.importanceLevel" clearable placeholder="请选择" @clear="()=>{searchInfo.importanceLevel=undefined}">
              <el-option v-for="(item,key) in import_levelOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="理解程度" prop="understandLevel">
            <el-select v-model="searchInfo.understandLevel" clearable placeholder="请选择" @clear="()=>{searchInfo.understandLevel=undefined}">
              <el-option v-for="(item,key) in understand_levelOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="addChild(0)">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="id" prop="id" width="120" />
        <el-table-column align="left" label="标题" prop="title" width="120" />
        <el-table-column align="left" label="父文章" prop="ptitle" width="120" />
        <el-table-column sortable align="left" label="知识库" prop="knowledgeName" width="120" />
        <el-table-column sortable align="left" label="重要程度" prop="importanceLevel" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.importanceLevel,import_levelOptions) }}
            </template>
        </el-table-column>
        <el-table-column sortable align="left" label="理解程度" prop="understandLevel" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.understandLevel,understand_levelOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="创建日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="更新日期" width="180">
             <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template>
        </el-table-column>
        <el-table-column sortable align="left" label="最近浏览" width="180">
            <template #default="scope">{{ formatDate(scope.row.lastViewedAt) }}</template>
         </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button
                          type="primary"
                          link
                          icon="plus"
                          @click="addChild(scope.row.id)"
                        >添加子文章</el-button>
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                            <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                            查看详情
                        </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateArticlesFunc(scope.row)">变更</el-button>
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
    <el-drawer size="1500" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
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
            <el-form-item label="知识库:"  prop="knowledgeId" >
              <el-input v-model.number="formData.knowledgeId" :clearable="true" placeholder="请输入知识库" />
            </el-form-item>
            <el-form-item label="父文章:"  prop="pid" >
              <el-input v-model.number="formData.pid" :clearable="true" placeholder="请输入父id" />
            </el-form-item>
            <el-form-item label="标题:"  prop="title" >
              <el-input v-model="formData.title" :clearable="true"  placeholder="请输入标题" />
            </el-form-item>
            <el-form-item label="重要程度:"  prop="importanceLevel" >
              <el-select v-model="formData.importanceLevel" placeholder="请选择重要程度" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in import_levelOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="理解程度:"  prop="understandLevel" >
              <el-select v-model="formData.understandLevel" placeholder="请选择理解程度" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in understand_levelOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="最后查看时间:"  prop="lastViewedAt" >
              <el-date-picker v-model="formData.lastViewedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="内容:"  prop="content" >
              <RichEdit v-model="formData.content"/>
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer size="1500" v-model="detailShow" :before-close="closeDetailShow" title="查看详情" destroy-on-close>
          <template #title>
             <div class="flex justify-between items-center">
               <span class="text-lg">查看详情</span>
             </div>
         </template>
        <el-descriptions :column="1" border>
                <el-descriptions-item label="知识库">
                        {{ formData.knowledgeId }}
                </el-descriptions-item>
                <el-descriptions-item label="父id">
                        {{ formData.pid }}
                </el-descriptions-item>
                <el-descriptions-item label="标题">
                        {{ formData.title }}
                </el-descriptions-item>
                <el-descriptions-item label="重要程度">
                        {{ filterDict(formData.importanceLevel,import_levelOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="理解程度">
                        {{ filterDict(formData.understandLevel,understand_levelOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="最后查看时间">
                      {{ formatDate(formData.lastViewedAt) }}
                </el-descriptions-item>
        </el-descriptions>
        <el-descriptions :column="100" border>
            <el-descriptions-item label="内容" :span="99">
                    <RichHtml v-model="formData.content"/>
            </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createArticles,
  deleteArticles,
  deleteArticlesByIds,
  updateArticles,
  findArticles,
  getArticlesList
} from '@/api/article/articles'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'
import RichHtml from '@/components/richtext/rich-view.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'Articles'
})

// 自动化生成的字典（可能为空）以及字段
const import_levelOptions = ref([])
const understand_levelOptions = ref([])
const formData = ref({
        knowledgeId: 0,
        pid: 0,
        ptitle: '',
        knowledgeName: '',
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
               },
              ],
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               importanceLevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               understandLevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               content : [{
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
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            knowledgeId: 'knowledge_id',
            pid: 'pid',
            importanceLevel: 'importance_level',
            understandLevel: 'understand_level',
            lastViewedAt: 'last_viewed_at',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}

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
  const table = await getArticlesList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    import_levelOptions.value = await getDictFunc('import_level')
    understand_levelOptions.value = await getDictFunc('understand_level')
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
            deleteArticlesFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteArticlesByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateArticlesFunc = async(row) => {
    const res = await findArticles({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rearticles
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteArticlesFunc = async (row) => {
    const res = await deleteArticles({ id: row.id })
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
  const res = await findArticles({ id: row.id })
  if (res.code === 0) {
    formData.value = res.data.rearticles
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          knowledgeId: 0,
          pid: 0,
          title: '',
          importanceLevel: '',
          understandLevel: '',
          lastViewedAt: new Date(),
          }
}

const addChild = (id) => {
  type.value = 'create'
  formData.value.pid = id
  dialogFormVisible.value = true
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
        knowledgeId: 0,
        pid: 0,
        title: '',
        importanceLevel: '',
        understandLevel: '',
        lastViewedAt: new Date(),
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
