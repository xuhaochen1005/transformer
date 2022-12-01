<template>
  <el-dialog
    v-model="show"
    title="调压范围"
    show-close
    width="40%"
  >
    <el-form
      label-width="100px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="调压范围下界（%）"
        label-width="170px"
      >
        {{ lrr.regulate_range_min }}
      </el-form-item>
      <el-form-item
        label="调压范围上界（%）"
        label-width="170px"
      >
        {{ lrr.regulate_range_max }}
      </el-form-item>
      <el-form-item
        label="调压范围步长（%）"
        label-width="170px"
      >
        {{ lrr.regulate_range_step }}
      </el-form-item>
      <el-form-item
        label="创建时间"
        label-width="150px"
      >
        {{ lrr.created_at }}
      </el-form-item>
      <el-form-item
        label="更新时间"
        label-width="150px"
      >
        {{ lrr.updated_at }}
      </el-form-item>
      <el-form-item
        label="删除时间"
        label-width="150px"
      >
        {{ lrr.deleted_at }}
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
  name: 'LrrView',
  props: {
    modelValue: Boolean,
    lrr: {
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
