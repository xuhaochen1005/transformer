<template>
  <el-dialog
    v-model="show"
    title="线圈导线类型标准库"
    show-close
  >
    <el-form
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="编号"
        label-width=""
      >
        {{ lcwt.id }}
      </el-form-item>
      <el-form-item
        label="线圈导线类型"
        label-width=""
      >
        {{ lcwt.coil_wire_type }}
      </el-form-item>
      <el-form-item
        label="线圈导线类型字母代号"
        label-width=""
      >
        {{ lcwt.coil_wire_type_sign }}
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
  name: 'LcwtView',
  props: {
    modelValue: Boolean,
    lcwt: {
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
