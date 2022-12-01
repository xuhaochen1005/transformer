<template>
  <el-dialog
    v-model="show"
    class="custom-dialog"
    width="800px"
  >
    <template #title>
      <div>
        <span class="el-dialog__title">反馈详情</span>
      </div>
    </template>
    <el-form
      :model="handleValue"
      label-width="100px"
      label-suffix=":"
    >
      <el-form-item
        label="ID"
        prop="id"
        label-width=""
      >
        {{ handleValue.id }}
      </el-form-item>
      <el-form-item
        label="反馈标题"
        prop="title"
        label-width=""
      >
        {{ handleValue.title }}
      </el-form-item>
      <el-form-item
        label="反馈内容"
        prop="content"
        label-width=""
      >
        <div
          class="message-content"
          v-html="handleValue.content"
        />
      </el-form-item>
      <el-form-item
        label="&emsp;创建者"
        label-width=""
      >
        {{ handleValue.creator_user?.user_name_zh }}
      </el-form-item>
      <!--      <el-form-item-->
      <!--        label="&emsp;接收者"-->
      <!--        prop="send_to"-->
      <!--        label-width=""-->
      <!--      >-->
      <!--        {{ tempDetail.send_to_user?.name }}-->
      <!--      </el-form-item>-->
    </el-form>
    <template #footer>
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
import { Feedback } from '@/model/feedback'
export default defineComponent({
  props: {
    modelValue: {
      type: Boolean,
      default: false
    },
    handleValue: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue'],
  setup(props, { emit }) {
    const show = computed({
      get: () => props.modelValue,
      set: (v) => emit('update:modelValue', v)
    })
    const defaultData: Feedback = {
      created_at: '', creator: 0, updated_at: '', content: '', id: 0, title: ''
    } as Feedback
    return {
      show
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
