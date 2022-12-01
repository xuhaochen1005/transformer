<template>
  <el-dialog
    v-model="show"
    title="上传文档"
    show-close
    width="30%"
  >
    <el-form
      ref="docsCreateForm"
      :model="docs"
      :rules="rules"
      label-width="100px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="添加文档"
        prop="name"
        label-width="150px"
      >
        <el-upload
          ref="uploadRef"
          action=""
          :http-request="httpRequest"
          :on-success="handleSuccess"
          :limit="1"
          :file-list="fileList"
          :auto-upload="false"
        >
          <el-button
            type="primary"
            style="margin-top:10px;margin-left: 10px"
            class="filter-item"
          >
            <i class="el-icon-upload el-icon--right" />
            上传文件
          </el-button>
        </el-upload>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createDocs"
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
import { defineComponent, computed, reactive, ref, PropType } from 'vue'
import { UDocuments } from '@/api/document'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'DocsCreate',
  props: {
    modelValue: Boolean,
    category: {
      type: String as PropType<'设计文档'|'培训文档'|'系统文档'>,
      default: ''
    }
  },
  emits: ['update:modelValue', 'get-docs-list'],
  setup(props, context) {
    const fileList = ref([])
    const docsCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          // (docsCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const docs = reactive({
      document_type: '',
      name: '',
      uuid: '',
      location: '',
      created_by: '',
      deleted_by: '',
      file_size: null,
      category: props.category
    })
    function handleSuccess() {
      fileList.value = []
    }
    async function httpRequest(param:any) {
      // console.log(param)
      const data = new FormData()
      data.append('file', param.file)
      data.append('category', props.category)
      // console.log(data)
      const response = await UDocuments(docs, data)
      if (response.data.code === 200) {
        ElMessage.success('文件上传成功!')
        context.emit('get-docs-list')
        show.value = false
      } else {
        ElMessage.error('请提供正确的文档信息')
      }
    }
    const uploadRef = ref(null)
    const createDocs = () => {
      (docsCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          (uploadRef.value as any).submit()
        }
      }
      )
    }
    function validateCategory(rule: any, value: any, callback: any) {
      if (!docs.category) {
        return callback(new Error('请选择文件归属'))
      }
      callback()
    }
    const rules = {
      category: [
        { required: true, validator: validateCategory, trigger: 'blur' }
      ]

    }
    return {
      handleSuccess,
      fileList,
      httpRequest,
      docsCreateForm,
      show,
      docs,
      uploadRef,
      rules,
      createDocs
    }
  }
})
</script>
