<template>
  <div>
    <div ref="jsMindContainer" style="height: 500px; border: 1px solid #ccc;"></div>
  </div>
</template>

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
      const options = {
        container: this.$refs.jsMindContainer,
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
            {"id": "easy", "topic": "Easy", "direction": "left", "children": [
              {"id": "easy1", "topic": "Easy to show"},
              {"id": "easy2", "topic": "Easy to edit"},
              {"id": "easy3", "topic": "Easy to store"},
              {"id": "easy4", "topic": "Easy to embed"}
            ]},
            {"id": "open", "topic": "Open Source", "direction": "right"},
            {"id": "powerful", "topic": "Powerful", "direction": "right"},
            {"id": "other", "topic": "test node", "direction": "right"}
          ]
        }
      }

      const jm = new jsMind(options)
      jm.show(mind)
    }
  }
}
</script>

<style>
/* Additional styles if needed */
</style>