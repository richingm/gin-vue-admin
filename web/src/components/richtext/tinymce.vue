<script setup >
import { ref, reactive } from 'vue';
import Editor from '@tinymce/tinymce-vue'
const key = '1linwngh9a1660apfjb30gkw6far7h4a93ex9wac8bw6pjl0'
let init = reactive({
    // 选择器
    selector: 'textarea',
    // 设置本地语言，在本地的路径
    language_url: '/tinymce/langs/zh-Hans.js',
    // 设置本地语言
    language: 'zh-Hans',
    // 设置工具栏
    toolbar: [
        'bold italic hr | fontsize forecolor backcolor | blocks blockquote removeformat | undo redo ',
        'bullist table insertdatetime | link charmap emoticons wordcount searchreplace code | codesample visualblocks image fullscreen preview',
    ],
    // 设置插件
    plugins: 'codesample lists advlist link autolink charmap emoticons fullscreen preview code searchreplace table visualblocks wordcount insertdatetime image',
})
let content = ref('')
</script>
<template>
     <Editor :api-key="key" v-model="content" :init="init" />
</template>