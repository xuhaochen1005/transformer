<template>
  <el-dialog
    v-model="show"
    class="custom-dialog"
    width="800px"
  >
    <template #title>
      <div>
        <span class="el-dialog__title">查看消息</span>
      </div>
    </template>
    <el-form
      :model="tempDetail"
      label-width="100px"
      label-suffix=":"
    >
      <el-form-item
        label="ID"
        prop="id"
        label-width=""
      >
        {{ tempDetail.id }}
      </el-form-item>
      <el-form-item
        label="消息标题"
        prop="title"
        label-width=""
      >
        {{ tempDetail.title }}
      </el-form-item>
      <el-form-item
        label="消息内容"
        prop="content"
        label-width=""
      >
        <div
          class="message-content"
          v-html="tempDetail.content"
        />
      </el-form-item>
      <el-form-item
        label="&emsp;发送者"
        prop="send_by"
        label-width=""
      >
        {{ tempDetail.send_by_user?.user_name_zh }}
      </el-form-item>
      <!--      <el-form-item-->
      <!--        label="&emsp;接收者"-->
      <!--        prop="send_to"-->
      <!--        label-width=""-->
      <!--      >-->
      <!--        {{ tempDetail.send_to_user?.name }}-->
      <!--      </el-form-item>-->
      <el-form-item
        label="消息状态"
        prop="status"
        label-width=""
      >
        <template
          v-for="option in messageStatusOptions"
          :key="option"
        >
          <span v-if="option.value === tempDetail.status">
            {{ option.label }}
          </span>
        </template>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        v-if="tempDetail.status === MessageStatus.HASREAD"
        type="primary"
        @click="$emit('handleSetMsgUnRead',handleValue);tempDetail.status = MessageStatus.UNREAD"
      >
        标为未读
      </el-button>
      <el-button
        @click="show = false"
      >
        关闭
      </el-button>
    </template>
  </el-dialog>
</template>

<script lang='ts'>
import {
  computed,
  defineComponent, onMounted,
  ref, watch, Ref
} from 'vue'

import type { Message } from '@/model/message'
import { MessageStatus } from '@/model/message'
export default defineComponent({
  props: {
    modelValue: Boolean,
    handleValue: {
      type: Object,
      required: true
    }
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
      send_to: 0,
      status: 0
    }
    const tempDetail : Ref<Message> = ref(JSON.parse(JSON.stringify(emptyMsg)) as Message)
    watch(() => props.handleValue, function (value) {
      console.log('watch', value)
      tempDetail.value = JSON.parse(JSON.stringify(value)) as Message
    })
    enum messageStatus {
      UNREAD = 1,
      HASREAD = 2
    }
    const messageStatusOptions = [
      { label: '未读', value: messageStatus.UNREAD },
      { label: '已读', value: messageStatus.HASREAD }
    ]
    return {
      show,
      tempDetail,
      messageStatusOptions,
      MessageStatus
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
:deep() {
  .message-content {
    font-size: 16px;
    a {
      border-bottom: 1px solid $--color-primary;
      color: $--color-primary;
      font-size: 18px;
    }
  }
}
</style>
