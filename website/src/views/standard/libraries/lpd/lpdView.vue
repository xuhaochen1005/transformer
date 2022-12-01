<template>
  <el-dialog
    v-model="show"
    title="相间距离最小值"
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
        {{ lpd.id }}
      </el-form-item>
      <el-form-item
        label="相间距离最小值（mm）"
        label-width=""
      >
        {{ lpd.phase_dist_min }}
      </el-form-item>
      <el-form-item
        label="电压下限（kV）"
        label-width=""
      >
        {{ lpd.vol_min }}
      </el-form-item>
      <el-form-item
        label="电压上限（kV）"
        label-width=""
      >
        {{ lpd.vol_max }}
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
  name: 'LpdView',
  props: {
    modelValue: Boolean,
    lpd: {
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
