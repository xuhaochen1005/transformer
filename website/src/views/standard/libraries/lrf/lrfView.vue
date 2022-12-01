<template>
  <el-dialog
    v-model="show"
    title="额定频率"
    show-close
    width="40%"
  >
    <el-form
      label-width="100px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定频率（Hz）"
        label-width="150px"
      >
        {{ lrf.rated_freq }}
      </el-form-item>
      <el-form-item
        label="创建时间"
        label-width="150px"
      >
        {{ lrf.created_at }}
      </el-form-item>
      <el-form-item
        label="更新时间"
        label-width="150px"
      >
        {{ lrf.updated_at }}
      </el-form-item>
      <el-form-item
        label="删除时间"
        label-width="150px"
      >
        {{ lrf.deleted_at }}
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
  name: 'LrfView',
  props: {
    modelValue: Boolean,
    lrf: {
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
