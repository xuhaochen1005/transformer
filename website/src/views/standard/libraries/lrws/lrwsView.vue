<template>
  <el-dialog
    v-model="show"
    title="圆铜（铝）线规格"
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
        {{ lrws.id }}
      </el-form-item>
      <el-form-item
        label="标称直径（mm）"
        label-width=""
      >
        {{ lrws.std_diameter }}
      </el-form-item>
      <el-form-item
        label="截面积（mm^2）"
        label-width=""
      >
        {{ lrws.section_area }}
      </el-form-item>
      <el-form-item
        label="最大外径（mm）"
        label-width=""
      >
        {{ lrws.max_diameter }}
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
  name: 'LrwsView',
  props: {
    modelValue: Boolean,
    lrws: {
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
