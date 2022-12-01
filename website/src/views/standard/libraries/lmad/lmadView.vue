<template>
  <el-dialog
    v-model="show"
    title="主风道初选"
    show-close
    width="40%"
  >
    <el-form
      label-width="100px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定高压下界（kV）"
        label-width="150px"
      >
        {{ lmad.rated_high_vol_min }}
      </el-form-item>
      <el-form-item
        label="额定高压上界（kV）"
        label-width="150px"
      >
        {{ lmad.rated_high_vol_max }}
      </el-form-item>
      <el-form-item
        label="主风道初选最小值（mm）"
        label-width="150px"
      >
        {{ lmad.main_air_duct_min }}
      </el-form-item>
      <el-form-item
        label="创建时间"
        label-width="150px"
      >
        {{ lmad.created_at }}
      </el-form-item>
      <el-form-item
        label="更新时间"
        label-width="150px"
      >
        {{ lmad.updated_at }}
      </el-form-item>
      <el-form-item
        label="删除时间"
        label-width="150px"
      >
        {{ lmad.deleted_at }}
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
  name: 'LmadView',
  props: {
    modelValue: Boolean,
    lmad: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue'],
  setup(props, context) {
    const status = new Map([[1, '正常'], [2, '停用']])
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        context.emit('update:modelValue', value)
      }
    })
    return {
      status,
      show
    }
  }
})
</script>
