<template>
  <div>
    <div id="jsmind_container" style="width: 99%; height: 700px;"></div>
  </div>
</template>

<script>
import { onMounted } from 'vue';
import jsMind from 'jsmind';
import 'jsmind/style/jsmind.css'; // 引入样式文件

export default {
  name: 'JsMindVue',
  setup() {
    onMounted(() => {
      const options = {
        container: 'jsmind_container',
        theme: 'primary',
        editable: true,
        view: {
                    engine: 'canvas', // 思维导图各节点之间线条的绘制引擎
                    hmargin: 10, // 思维导图距容器外框的最小水平距离
                    vmargin: 10, // 思维导图距容器外框的最小垂直距离
                    line_width: 1, // 思维导图线条的粗细
                    line_color: '#11f' // 思维导图线条的颜色
                },
                layout: {
                    hspace: 60, // 节点之间的水平间距
                    vspace: 10, // 节点之间的垂直间距
                    pspace: 10 // 节点与连接线之间的水平间距（用于容纳节点收缩/展开控制器）
                },
                shortcut: {
                    enable: false // 是否启用快捷键 默认为true
                }
      };

      const mind = {
        "meta": {
          "name": "demo",
          "author": "hizzgdev@163.com",
          "version": "0.2"
        },
        "format": "node_tree",
        "data": {}
      };

      const jm = new jsMind(options);
      jm.show(mind);
    });
  }
};
</script>

<style scoped>
#jsmind_container {
  border: 1px solid #ccc;
  background: #f4f4f4;
}
</style>