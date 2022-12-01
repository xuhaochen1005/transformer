<template>
  <el-dialog
    v-model="show"
    title="修改文档信息"
    show-close
  >
    <el-form
      ref="docsUpdateForm"
      :model="tmpDocs"
      :rules="rules"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="文件名"
        label-width="200px"
        prop="name"
      >
        <el-input
          v-model="tmpDocs.name"
          type="name"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="文档类型(标准库、任务书等)"
        label-width="200px"
        prop="document_type"
      >
        <el-select
          v-model="tmpDocs.document_type"
          placeholder="文档类型"
          clearable
          style="width: 150px; margin-right: 10px"
        >
          <el-option
            label="标准库"
            value="标准库"
          />
          <el-option
            label="任务书"
            value="任务书"
          />
        </el-select>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateDocs"
      >
        确定
      </el-button>
      <el-button @click="show = false">
        取消
      </el-button>
    </template>
  </el-dialog>
</template>

<script lang="ts">
import { computed, defineComponent, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Documents, UpdateDocuments } from '@/api/document'

export default defineComponent({
  name: 'DocsUpdate',
  props: {
    modelValue: Boolean,
    docs: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-docs-list'],
  setup (props, context) {
    const docsUpdateForm = ref(null)
    const docs: Documents = {
      id: 0,
      name: '',
      uuid: '',
      location: '',
      file_size: null,
      document_type: '',
      category: ''
    }
    const tmpDocs = ref(docs)
    watch(props, () => {
      tmpDocs.value.id = props.docs.id
      tmpDocs.value.name = props.docs.name
      tmpDocs.value.document_type = props.docs.document_type
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (docsUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    async function updateDocs() {
      const response = await UpdateDocuments(tmpDocs.value)
      if (response.data.code === 200) {
        ElMessage.success('信息更新成功')
        context.emit('get-docs-list')
        show.value = false
      }
    }
    return {
      docsUpdateForm,
      show,
      tmpDocs,
      updateDocs
    }
  }
})
</script>
