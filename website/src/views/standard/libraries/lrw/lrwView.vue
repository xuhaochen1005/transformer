<template>
  <el-dialog
    v-model="show"
    title="调压方式"
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
        {{ lrw.id }}
      </el-form-item>
      <el-form-item
        label="调压方式"
        label-width=""
      >
        {{ lrw.regulate_way }}
      </el-form-item>
      <el-form-item
        label="调压方式代号"
        label-width=""
      >
        {{ lrw.regulate_way_sign }}
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
  name: 'LrwView',
  props: {
    modelValue: Boolean,
    lrw: {
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
