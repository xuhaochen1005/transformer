<template>
  <el-dialog
    v-model="show"
    title="引线损耗"
    show-close
  >
    <el-form
      label-width="250px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="编号"
        label-width=""
      >
        {{ lls.id }}
      </el-form-item>
      <el-form-item
        label="额定容量(kVA)"
        label-width=""
      >
        {{ lls.rated_capacity }}
      </el-form-item>
      <el-form-item
        label="引线损耗(W)"
        label-width=""
      >
        {{ lls.lead_loss }}
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="show = false">
        关闭
      </el-button>
    </template>
  </el-dialog>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue'

export default defineComponent({
  name: 'LlsView',
  props: {
    modelValue: Boolean,
    lls: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue'],
  setup(props, context) {
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        context.emit('update:modelValue', value)
      }
    })
    return {
      show
    }
  }
})
</script>
