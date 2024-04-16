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
import jsMind from 'jsmind'
import 'jsmind/style/jsmind.css'

import {
  getKnowledgesOptions
} from '@/api/knowledge/knowledges'

import {
  getArticlesMind
} from '@/api/article/articles'

import { ref, reactive, onMounted, nextTick } from 'vue'

const searchInfo = ref({knowledgeId: 0})

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
});

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

const jsMindContainer = ref('jsmind_container');
const jm = ref(null);

onMounted(() => {
  nextTick(() => {
    const options = {
      container: 'jsMindContainer', // ID of your container element
      editable: true,
      theme: 'primary'
    };

    const mind = ref({
          "meta": {
            "name": "demo",
            "author": "hizzgdev@163.com",
            "version": "0.2"
          },
          "format": "node_tree",
          "data": {}
    });

    if (jsMindContainer.value) {
      jm.value = new jsMind(options);
      // If you need to load initial data, call `fetchMindData` here
    } else {
      console.error('jsMind container not found in DOM');
    }
  });
});

setOptions()


const onSubmit = () => {
    fetchMindData(searchInfo.value.knowledgeId)
}

const fetchMindData = async (id) => {
  try {
    console.log(options)
    const res = await getArticlesMind({ knowledgeId: id });
    mind.data = res.data;
    if (jm.value) {
      jm.value.show(mind);
    } else {
      jm.value = new jsMind(options);
      jm.value.show(mind);
    }
  } catch (error) {
    console.error('Failed to fetch articles mind:', error);
  }
};
</script>

<style>
/* Additional styles if needed */
</style>