<template>
  <el-dialog
    v-model="show"
    title="迭片系数"
    show-close
  >
    <el-form
      label-width="220px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="编号"
        label-width=""
      >
        {{ lsf.id }}
      </el-form-item>
      <el-form-item
        label="迭片片厚（mm）"
        label-width=""
      >
        {{ lsf.stack_thick }}
      </el-form-item>
      <el-form-item
        label="迭片系数"
        label-width=""
      >
        {{ lsf.stack_factor }}
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
  name: 'LsfView',
  props: {
    modelValue: Boolean,
    lsf: {
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
