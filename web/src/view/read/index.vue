<template>
  <div>
    <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="知识库" style="width:20%" prop="knowledgeId">
        <el-cascader
        v-model.number="searchInfo.knowledgeId"
        style="width:100%"
        :disabled="false"
        :options="knowledgeOption"
        :props="{ checkStrictly: true, label: 'name', value: 'id', emitPath: false }"
        :show-all-levels="false"
        filterable
        />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
        </el-form-item>
    </el-form>
    <div ref="jsMindContainer" style="height: 660px; border: 1px solid #ccc;"></div>
  </div>
</template>

<script setup>
import {
  getKnowledgesOptions
} from '@/api/knowledge/knowledges'

import { ref, reactive } from 'vue'

const searchInfo = ref({
        knowledgeId: 0})

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

const knowledgeOption = ref([])
const parentArticleOption = ref([
  {
    id: 0,
    name: '根文章'
  }
])

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    knowledgeOption.value = []
    setKnowledgeOptions(knowledgeOption.value, false)

    parentArticleOption.value = [
        {
          id: 0,
          name: '根文章'
        }
      ]
    setParentArticleOption(parentArticleOption.value, 0, '', false)
}

const setParentArticleOption = (optionsData, id, name, disabled) => {
        if (id > 0 ) {
            const option = {
                    name: name,
                    id: id,
                    disabled: disabled || id === formData.value.pid
                  }
            optionsData.push(option)
        }
}

const setKnowledgeOptions = async (optionsData, disabled) => {
  const knowledgeData = await getKnowledgesOptions({})
  if (knowledgeData.data.knowledges) {
    knowledgeData.data.knowledges.forEach(item => {
      const option = {
        name: item.name,
        id: item.id,
        disabled: disabled || item.id === formData.value.knowledgeId
      }
      optionsData.push(option)
    });
  }
}

const jsMindContainer = ref(null);
    const jm = ref(null);

setOptions()

const onSubmit = () => {
    console.log('xxx')
    const options = {
            container: jsMindContainer.value,
                editable: true,
                theme: 'primary'
              }


          const mind = {
            "meta": {
              "name": "demo",
              "author": "hizzgdev@163.com",
              "version": "0.2"
            },
            "format": "node_tree",
            "data": {
                                "id": "root",
                                "topic": "jsMind",
                                "children": [
                                  {
                                    "id": "easy",
                                    "topic": "Easy",
                                    "direction": "left",
                                    "children": [
                                      {"id": "easy1", "topic": "Easy to show"},
                                      {"id": "easy2", "topic": "Easy to edit"},
                                      // ... other existing children ...
                                      {
                                        "id": "easy4",
                                        "topic": "Easy to embed",
                                        "children": [
                                          {
                                            "id": "level1",
                                            "topic": "Level 1",
                                            "children": [
                                              {
                                                "id": "level2",
                                                "topic": "Level 2",
                                                "children": [
                                                  {
                                                    "id": "level3",
                                                    "topic": "Level 3",
                                                    "children": [
                                                      {
                                                        "id": "level4",
                                                        "topic": "Level 4",
                                                        "children": [
                                                          {
                                                            "id": "level5",
                                                            "topic": "Level 5",
                                                            "children": [
                                                              {
                                                                "id": "level6",
                                                                "topic": "Level 4",
                                                                "children": [
                                                                  {
                                                                    "id": "level7",
                                                                    "topic": "Level 5",
                                                                    "children": [
                                                                          {
                                                                            "id": "level8",
                                                                            "topic": "Level 4",
                                                                            "children": [
                                                                              {
                                                                                "id": "level9",
                                                                                "topic": "Level 5"
                                                                                // This is the fifth level of nesting
                                                                              }
                                                                            ]
                                                                          }
                                                                        ]
                                                                  }
                                                                ]
                                                              }
                                                            ]
                                                          }
                                                        ]
                                                      }
                                                    ]
                                                  }
                                                ]
                                              }
                                            ]
                                          }
                                        ]
                                      }
                                      // ... continue with other nodes as required ...
                                    ]
                                  },
                                  // ... other existing nodes ...
                                  {
                                    "id": "other",
                                    "topic": "test node",
                                    "direction": "right",
                                    "children": [
                                      {
                                        "id": "easy11",
                                        "topic": "Easy",
                                        "direction": "left",
                                        "children": [
                                          // Add nested levels here as well if required
                                        ]
                                      }
                                    ]
                                  }
                                ]
                              }
          }

      if (jm.value) {
        // If jm already exists, use the existing instance to show the new mind map
        jm.value.show(mind);
      } else {
        // If jm doesn't exist, create a new jsMind instance and show the mind map
        jm.value = new jsMind(options);
        jm.value.show(mind);
      }
}
</script>
<script>
import jsMind from 'jsmind'
import 'jsmind/style/jsmind.css'

export default {
  name: 'JsMindDemo',
  mounted() {
    this.initJsMind()
  },
  methods: {
    initJsMind() {

    }
  }
}
</script>

<style>
/* Additional styles if needed */
</style>