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
        <el-button type="primary" icon="search" @click="show_selected">查看</el-button>
      </el-form-item>
    </el-form>
    <div ref="jsMindContainer" id="jsMindContainer" style="height: 660px; border: 1px solid #ccc;"></div>
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

setOptions()


const jsMindContainer = ref(null); // Use Vue ref here
const jm = ref(null);
const mind = ref({});

const options = {
  container: 'jsMindContainer', // Make sure this matches the ID in your template
  editable: true,
  theme: 'primary',
  default_event_handle: {
      enable_mousedown_handle: true
    },
    event_handle: {
      click_element_handle: (el) => handleNodeClick(el)
    }
};

const show_selected = () => {
  console.log(jm)
    var selected_node = jm.get_selected_node();
    if (!!selected_node) {
        console.log(selected_node.topic)
    } else {
        console.log('nothing');
    }
}


const handleNodeClick = (el) => {
    console.log('xxxxxx')
  const nodeId = el.getAttribute('id');
  console.log(nodeId)
};


onMounted(() => {
  nextTick(() => {
    if (jsMindContainer.value) {
      jm.value = new jsMind(options);
    } else {
      console.error('jsMind container not found in DOM');
    }    
  });
});


const onSubmit = async () => {
  await fetchMindData(searchInfo.value.knowledgeId);
}

const selectedNodeId = ref(null); // 用于存储选中节点的 ID

const onShow = async () => {
  const node = jm.value.selected_node()
  console.log(node)
}

const fetchMindData = async (id) => {
  try {
    const res = await getArticlesMind({ knowledgeId: id });
    if (jm.value) {
      jm.value.show( res.data.rearticles);
    } else {
      jm.value = new jsMind(options);
      jm.value.show( res.data.rearticles);
    }
  } catch (error) {
    console.error('Failed to fetch articles mind:', error);
  }
};

</script>

<style>
/* Additional styles if needed */
</style>