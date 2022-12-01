<template>
  <el-dialog
    v-model="show"
    title="绕组温升限值"
    show-close
    width="25%"
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
        {{ ltr.id }}
      </el-form-item>
      <el-form-item
        label="绝缘系统温度（℃）"
        label-width=""
      >
        {{ ltr.temp }}
      </el-form-item>
      <el-form-item
        label="绝缘温度代号"
        label-width=""
      >
        {{ ltr.temp_sign }}
      </el-form-item>
      <el-form-item
        label="低压温升限值（K）"
        label-width=""
      >
        {{ ltr.low_vol_temp_rise }}
      </el-form-item>
      <el-form-item
        label="高压温升限值（K）"
        label-width=""
      >
        {{ ltr.high_vol_temp_rise }}
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
  name: 'LtrView',
  props: {
    modelValue: Boolean,
    ltr: {
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
