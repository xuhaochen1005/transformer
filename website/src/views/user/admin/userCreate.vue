<template>
  <el-dialog
    v-model="show"
    title="添加新用户"
    width="25%"
    show-close
  >
    <el-form
      ref="userCreateForm"
      :model="user"
      :rules="rules"
      label-width="100px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="&emsp;账号名"
        prop="name"
        label-width=""
      >
        <el-input
          v-model="user.name"
          placeholder="账号名"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="用户名"
        prop="user_name_zh"
        label-width=""
      >
        <el-input
          v-model="user.user_name_zh"
          placeholder="用户名"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="用户密码"
        prop="password"
        label-width=""
      >
        <el-input
          v-model="user.password"
          placeholder="用户密码"
          type="password"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="确认密码"
        prop="repeat_password"
        label-width=""
      >
        <el-input
          v-model="user.repeat_password"
          placeholder="确认密码"
          type="password"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="手机号码"
        prop="telephone"
        label-width=""
      >
        <el-input
          v-model="user.telephone"
          placeholder="手机号码"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="电子邮箱"
        prop="email"
        label-width=""
      >
        <el-input
          v-model="user.email"
          placeholder="电子邮箱"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="所属部门"
        prop="department_id"
        label-width=""
      >
        <el-select
          v-model="user.department_id"
          style="width: 100%;"
          placeholder="请选择所属部门"
          clearable
        >
          <el-option
            v-for="item in departmentOptions"
            :key="item"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="账号状态"
        prop="status"
        label-width=""
      >
        <el-select
          v-model="user.status"
          style="width: 100%;"
          placeholder="请设置账号状态"
          clearable
        >
          <el-option
            label="正常"
            :value="1"
          />
          <el-option
            label="停用"
            :value="2"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="&emsp;用户说明"
        prop="info"
        label-width=""
      >
        <el-input
          v-model="user.info"
          placeholder="用户说明"
          type="textarea"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createUser"
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
import { defineComponent, computed, reactive, ref } from 'vue'
import { UnwrapNestedRefs } from '@vue/reactivity'
import { CreateUser } from '@/api/user'
import { ElMessage } from 'element-plus'
import type { User } from '@/model/common'

export default defineComponent({
  name: 'UserCreate',
  props: {
    modelValue: Boolean,
    departmentOptions: Array
  },
  emits: ['update:modelValue', 'get-user-list'],
  setup(props, context) {
    const userCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (userCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const user: UnwrapNestedRefs<User> = reactive({
      name: '',
      user_name_zh: '',
      password: '',
      repeat_password: '',
      telephone: '',
      email: '',
      status: 1,
      info: ''
    })
    function validateUsername(rule: any, value: any, callback: any) {
      if (!user.name) {
        return callback(new Error('请输入用户账号'))
      }
      if (user.name.length < 4 || user.name.length > 32) {
        return callback(new Error('用户账号长度只能为4到32'))
      }
      callback()
    }
    function validateUserNameZh(rule: any, value: any, callback: any) {
      if (!user.user_name_zh) {
        return callback(new Error('请输入用户名称'))
      }
      if (user.user_name_zh.length < 2 || user.user_name_zh.length > 32) {
        return callback(new Error('用户名称长度只能为2到32'))
      }
      callback()
    }
    function validatePassword(rule: any, value: any, callback: any) {
      if (!user.password) {
        return callback(new Error('请输入用户密码'))
      }
      if (user.password.length < 8 || user.password.length > 32) {
        return callback(new Error('密码长度只能为8到32'))
      }
      callback()
    }
    function validateRepeatPassword(rule: any, value: any, callback: any) {
      if (!user.repeat_password) {
        return callback(new Error('请再次输入用户密码'))
      }
      if (user.password !== user.repeat_password) {
        return callback(new Error('两次输入的密码不一致'))
      }
      callback()
    }
    function validateTelephone(rule: any, value: any, callback: any) {
      if (!user.telephone) {
        return callback(new Error('请设置用户手机号码'))
      }
      if (user.telephone.length !== 11) {
        return callback(new Error('手机号码必须是11位'))
      }
      callback()
    }
    function validateEmail(rule: any, value: any, callback: any) {
      if (!user.email) {
        return callback(new Error('请设置用户邮箱'))
      }
      if (user.email.length < 6 || user.email.length > 32) {
        return callback(new Error('邮箱格式不正确'))
      }
      callback()
    }
    function validateDepartmentId(rule: any, value: any, callback: any) {
      if (!user.department_id) {
        return callback(new Error('请设置用户所属部门'))
      }
      callback()
    }
    function validateStatus(rule: any, value: any, callback: any) {
      if (!user.status) {
        return callback(new Error('请设置用户账号状态'))
      }
      callback()
    }
    const rules = {
      name: [
        { required: true, validator: validateUsername, trigger: 'blur' }
      ],
      user_name_zh: [
        { required: true, validator: validateUserNameZh, trigger: 'blur' }
      ],
      password: [
        { required: true, validator: validatePassword, trigger: 'blur' }
      ],
      repeat_password: [
        { required: true, validator: validateRepeatPassword, trigger: 'blur' }
      ],
      telephone: [
        { required: true, validator: validateTelephone, trigger: 'blur' }
      ],
      email: [
        { required: true, validator: validateEmail, trigger: 'blur' }
      ],
      department_id: [
        { required: true, validator: validateDepartmentId, trigger: 'change' }
      ],
      status: [
        { required: true, validator: validateStatus, trigger: 'change' }
      ]
    }
    function createUser() {
      (userCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateUser(user)
          if (response.data.code === 200) {
            ElMessage.success('新用户创建成功')
            context.emit('get-user-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的用户信息')
        }
      })
    }
    return {
      userCreateForm,
      show,
      user,
      rules,
      createUser
    }
  }
})
</script>
