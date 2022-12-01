<template>
  <el-dialog
    v-model="show"
    class="custom-dialog"
    width="800px"
  >
    <template #title>
      <div>
        <span class="el-dialog__title">创建反馈</span>
      </div>
    </template>
    <div>
      <el-form
        ref="createForm"
        :model="tempDetail"
        :rules="msgFormRules"
        label-suffix=":"
        label-width="100px"
      >
        <el-form-item
          label="反馈标题"
          label-width=""
          prop="title"
        >
          <el-input
            v-model="tempDetail.title"
            clearable
            placeholder="反馈标题"
          />
        </el-form-item>
        <el-form-item
          label="反馈内容"
          label-width=""
          prop="content"
        >
          <el-input
            v-model="tempDetail.content"
            :rows="10"
            clearable
            placeholder="反馈内容"
            type="textarea"
          />
        </el-form-item>
      </el-form>
    </div>
    <template #footer>
      <el-button
        type="primary"
        @click="confirmCreate()"
      >
        发送反馈
      </el-button>
      <el-button
        @click="show = false"
      >
        取消
      </el-button>
    </template>
  </el-dialog>
</template>

<script lang='ts'>
import { computed, defineComponent, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Feedback } from '@/model/feedback'
import { CreateFeedback } from '@/api/feedback'

export default defineComponent({
  props: {
    modelValue: {
      type: Boolean,
      default: false
    }
  },
  emits: ['update:modelValue', 'updateList'],
  setup(props, { emit }) {
    const show = computed({
      get() {
        return props.modelValue
      },
      set(val) {
        emit('update:modelValue', val)
      }
    })
    const defaultData: Feedback = {
      creator: 0, content: '', id: 0, title: ''
    } as Feedback
    const tempDetail = ref(JSON.parse(JSON.stringify(defaultData)))

    // 反馈通用验证表单
    function validateTitle(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('标题不可为空'))
      }
      if (value.length > 64) {
        callback(new Error('反馈标题长度不能大于64'))
      }
      callback()
    }

    function validateContent(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('内容不可为空'))
      }
      callback()
    }

    const msgFormRules = {
      title: [
        { required: true, validator: validateTitle, trigger: 'blur' }
      ],
      content: [
        { required: true, validator: validateContent, trigger: 'blur' }
      ]
    }

    const createForm = ref(null)

    function confirmCreate() {
      (createForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const tm = tempDetail.value
          const response = await CreateFeedback(tempDetail.value)
          if (response.data.code === 200) {
            ElMessage.success('反馈发送成功')
            show.value = false
            emit('updateList', true)
          }
        }
      })
    }

    return {
      show,
      tempDetail,
      msgFormRules,
      createForm,
      confirmCreate
    }
  }
})
</script>

<style lang="scss" scoped>
.form-icon {
  font-size: 18px;

  &.warning {
    color: $yellow;
  }
}
</style>
