<template>
  <el-dialog
    v-model="show"
    class="custom-dialog"
    width="800px"
  >
    <template #title>
      <div>
        <span class="el-dialog__title">发送消息</span>
      </div>
    </template>
    <el-form
      ref="msgCreateForm"
      :model="tempMsgDetail"
      :rules="msgFormRules"
      label-width="100px"
      label-suffix=":"
    >
      <el-form-item
        label="消息标题"
        prop="title"
        label-width=""
      >
        <el-input
          v-model="tempMsgDetail.title"
          placeholder="消息标题"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="消息内容"
        prop="content"
        label-width=""
      >
        <el-input
          v-model="tempMsgDetail.content"
          placeholder="消息内容"
          type="textarea"
          :rows="10"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="&emsp;接收人"
        prop="send_to"
        label-width=""
      >
        <UserSearchInput
          v-model="tempMsgDetail.send_to"
          style="width: 90%;"
          multiple
          :users="tempMsgDetail.send_to_user"
        />
        <el-tooltip
          content="接收人为空则为所有用户"
          placement="top"
          style="margin-left: 10px;"
        >
          <i
            class="form-icon warning el-icon-warning-outline"
          />
        </el-tooltip>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="confirmMsgCreate()"
      >
        发送消息
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
import {
  computed,
  defineComponent, onMounted,
  ref
} from 'vue'

import { Message } from '@/model/message'
import { MessageNotify } from '@/api/message'
import { ElMessage } from 'element-plus'
import UserSearchInput from '@/components/user-search-input/index.vue'
export default defineComponent({
  components: {
    UserSearchInput
  },
  props: {
    modelValue: Boolean
  },
  setup(props, ctx) {
    const show = computed({
      get: () => props.modelValue,
      set: (v) => ctx.emit('update:modelValue', v)
    })
    const emptyMsg: Message = {
      id: 0,
      title: '',
      content: '',
      send_by: 0,
      send_by_user: undefined,
      send_to: 0,
      send_to_user: undefined,
      status: 0
    }
    enum messageStatus {
      UNREAD = 1,
      HASREAD = 2
    }
    const messageStatusOptions = [
      { label: '未读', value: messageStatus.UNREAD },
      { label: '已读', value: messageStatus.HASREAD }
    ]

    const tempMsgDetail = ref(emptyMsg)
    // 消息通用验证表单
    function validateMsgTitle(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('标题不可为空'))
      }
      if (value.length > 64) {
        callback(new Error('消息标题长度不能大于64'))
      }
      callback()
    }
    function validateMsgContent(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('内容不可为空'))
      }
      callback()
    }
    function validateMsgStatus(rule: any, value: any, callback: any) {
      if (!(value in messageStatus)) {
        return callback(new Error('状态只能为已读或未读'))
      }
      callback()
    }
    const msgFormRules = {
      title: [
        { required: true, validator: validateMsgTitle, trigger: 'blur' }
      ],
      content: [
        { required: true, validator: validateMsgContent, trigger: 'blur' }
      ],
      status: [
        { required: true, validator: validateMsgStatus, trigger: 'blur' }
      ]
    }

    // 创建
    const messageCreate: Message = {
      id: 0,
      title: '',
      content: '',
      send_by: 0,
      send_to: 0,
      status: 0
    }
    const msgCreateForm = ref(null)
    function confirmMsgCreate() {
      (msgCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const tm = tempMsgDetail.value
          messageCreate.title = tm.title
          messageCreate.content = tm.content
          messageCreate.send_to = tm.send_to
          messageCreate.status = tm.status
          const response = await MessageNotify(messageCreate)
          if (response.data.code === 200) {
            ElMessage.success('消息发送成功')
            show.value = false
          }
        }
      })
    }
    return {
      show,
      tempMsgDetail,
      msgFormRules,
      msgCreateForm,
      messageStatusOptions,
      confirmMsgCreate
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
