<template>
  <el-dialog
    v-model="show"
    title="用户详情"
    show-close
  >
    <el-form
      label-width="100px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="用户编号"
        label-width=""
      >
        {{ user.id }}
      </el-form-item>
      <el-form-item
        label="用户账号"
        label-width=""
      >
        {{ user.name }}
      </el-form-item>
      <el-form-item
        label="用户名称"
        label-width=""
      >
        {{ user.user_name_zh }}
      </el-form-item>
      <el-form-item
        label="手机号码"
        label-width=""
      >
        {{ user.telephone }}
      </el-form-item>
      <el-form-item
        label="电子邮箱"
        label-width=""
      >
        {{ user.email }}
      </el-form-item>
      <el-form-item
        label="所属部门"
        label-width=""
      >
        {{ user.department_name }}
      </el-form-item>
      <el-form-item
        label="账号状态"
        label-width=""
      >
        {{ status.get(user.status) }}
      </el-form-item>
      <el-form-item
        label="用户说明"
        label-width=""
      >
        {{ user.info }}
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
  name: 'UserView',
  props: {
    modelValue: Boolean,
    user: {
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
