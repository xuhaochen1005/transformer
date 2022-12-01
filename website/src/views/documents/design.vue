<template>
  <div>
    <docs-table category="设计文档" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import docsTable from '@/views/documents/components/docsTable.vue'
export default defineComponent({
  components: { docsTable },

  setup() {
    return {}
  }
})
</script>

<style scoped lang="scss">
</style>
