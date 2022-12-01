<template>
  <el-dialog
    v-model="show"
    title="冷却方式标准库"
    show-close
  >
    <el-form
      label-width="100px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="编号"
        label-width=""
      >
        {{ lct.id }}
      </el-form-item>
      <el-form-item
        label="冷却方式"
        label-width=""
      >
        {{ lct.cool_type }}
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
  name: 'LctView',
  props: {
    modelValue: Boolean,
    lct: {
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
