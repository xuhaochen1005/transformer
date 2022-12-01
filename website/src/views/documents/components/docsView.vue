<template>
  <el-dialog
    v-model="show"
    title="文档信息"
    show-close
    width="40%"
  >
    <el-form
      label-width="200px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="文件名"
      >
        {{ docs.name }}
      </el-form-item>
      <el-form-item
        label="文件存储路径"
      >
        {{ docs.location }}
      </el-form-item>
      <el-form-item
        label="创建者"
      >
        {{ docs.coil_outer_insulate }}
      </el-form-item>
      <el-form-item
        label="文件大小(kb为单位)"
      >
        {{ docs.file_size }}
      </el-form-item>
      <el-form-item
        label="文档类型(标准库、任务书等)"
      >
        {{ docs.file_size }}
      </el-form-item>
      <el-form-item
        label="创建时间"
      >
        {{ UnixTime2HumanWithStr(docs.created_at) }}
      </el-form-item>
      <el-form-item
        label="更新时间"
      >
        {{ UnixTime2HumanWithStr(docs.updated_at) }}
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
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  name: 'DocsView',
  props: {
    modelValue: Boolean,
    docs: {
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
      show,
      UnixTime2HumanWithStr
    }
  }
})
</script>
