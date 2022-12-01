<template>
  <el-dialog
    v-model="show"
    title="线圈接法标准库"
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
        {{ lcc.id }}
      </el-form-item>
      <el-form-item
        label="低压线圈接法"
        label-width=""
      >
        {{ lcc.low_vol_coil_connect }}
      </el-form-item>
      <el-form-item
        label="高压线圈接法"
        label-width=""
      >
        {{ lcc.high_vol_coil_connect }}
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
  name: 'LccView',
  props: {
    modelValue: Boolean,
    lcc: {
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
