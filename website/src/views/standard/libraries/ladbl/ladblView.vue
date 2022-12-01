<template>
  <el-dialog
    v-model="show"
    title="风道条台账标准库"
    show-close
    width="25%"
  >
    <el-form
      label-width="160px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="编号"
        label-width=""
      >
        {{ ladbl.id }}
      </el-form-item>
      <el-form-item
        label="风道条厚度(mm)"
        label-width=""
      >
        {{ ladbl.air_duct_bar_thickness }}
      </el-form-item>
      <el-form-item
        label="风道条宽度(mm)"
        label-width=""
      >
        {{ ladbl.air_duct_bar_width }}
      </el-form-item>
      <el-form-item
        label="风道条长度(mm)"
        label-width=""
      >
        {{ ladbl.air_duct_bar_length }}
      </el-form-item>
      <el-form-item
        label="风道条数量(个)"
        label-width=""
      >
        {{ ladbl.air_duct_bar_num }}
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
  name: 'LadblView',
  props: {
    modelValue: Boolean,
    ladbl: {
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
