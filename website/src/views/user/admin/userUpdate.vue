<template>
  <el-dialog
    v-model="show"
    title="修改用户"
    width="25%"
    show-close
  >
    <el-form
      ref="userUpdateForm"
      :model="tmpUser"
      :rules="rules"
      label-width="100px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="用户账号"
        prop="name"
        label-width=""
      >
        <el-input
          v-model="tmpUser.name"
          placeholder="用户账号"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="用户名"
        prop="user_name_zh"
        label-width=""
      >
        <el-input
          v-model="tmpUser.user_name_zh"
          placeholder="用户名"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="新密码"
        prop="password"
        label-width=""
      >
        <el-input
          v-model="tmpUser.password"
          placeholder="新密码"
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
          v-model="tmpUser.repeat_password"
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
          v-model="tmpUser.telephone"
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
          v-model="tmpUser.email"
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
          v-model="tmpUser.department_id"
          placeholder="请选择所属部门"
          style="width: 100%;"
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
          v-model="tmpUser.status"
          placeholder="请设置账号状态"
          style="width: 100%;"
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
          v-model="tmpUser.info"
          placeholder="用户说明"
          type="textarea"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateUser"
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
import { computed, defineComponent, reactive, ref, watch } from 'vue'
import { UpdateUser } from '@/api/user'
import { ElMessage } from 'element-plus'
import type { User } from '@/model/common'
import department from '@/router/asyncRoutes/department'

export default defineComponent({
  name: 'UserUpdate',
  props: {
    modelValue: Boolean,
    user: {
      type: Object,
      required: true
    },
    departmentOptions: Array
  },
  emits: ['update:modelValue', 'get-user-list'],
  setup (props, context) {
    const userUpdateForm = ref(null)
    const user: User = {
      id: 0,
      name: '',
      user_name_zh: '',
      password: '',
      repeat_password: '',
      telephone: '',
      email: '',
      department_id: 1,
      status: 0,
      info: ''
    }
    const tmpUser = ref(user)
    watch(props, () => {
      tmpUser.value.id = props.user.id
      tmpUser.value.name = props.user.name
      tmpUser.value.user_name_zh = props.user.user_name_zh
      tmpUser.value.password = ''
      tmpUser.value.repeat_password = ''
      tmpUser.value.telephone = props.user.telephone
      tmpUser.value.email = props.user.email
      tmpUser.value.department_id = props.user.department_id
      tmpUser.value.status = props.user.status
      tmpUser.value.info = props.user.info
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (userUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateUsername(rule: any, value: any, callback: any) {
      if (!tmpUser.value.name) {
        return callback(new Error('请输入用户账号'))
      }
      if (tmpUser.value.name.length < 4 || tmpUser.value.name.length > 32) {
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
      if (tmpUser.value.password && (tmpUser.value.password.length < 8 || tmpUser.value.password.length > 32)) {
        return callback(new Error('密码长度只能为8到32'))
      }
      callback()
    }
    function validateRepeatPassword(rule: any, value: any, callback: any) {
      if (tmpUser.value.password && !tmpUser.value.repeat_password) {
        return callback(new Error('请再次输入用户密码'))
      }
      if (tmpUser.value.password !== tmpUser.value.repeat_password) {
        return callback(new Error('两次输入的密码不一致'))
      }
      callback()
    }
    function validateTelephone(rule: any, value: any, callback: any) {
      if (!tmpUser.value.telephone) {
        return callback(new Error('请设置用户手机号码'))
      }
      if (tmpUser.value.telephone.length !== 11) {
        return callback(new Error('手机号码必须是11位'))
      }
      callback()
    }
    function validateEmail(rule: any, value: any, callback: any) {
      if (!tmpUser.value.email) {
        return callback(new Error('请设置用户邮箱'))
      }
      if (tmpUser.value.email.length < 6 || tmpUser.value.email.length > 32) {
        return callback(new Error('邮箱格式不正确'))
      }
      callback()
    }
    function validateDepartmentId(rule: any, value: any, callback: any) {
      if (!tmpUser.value.department_id) {
        return callback(new Error('请设置用户所属部门'))
      }
      callback()
    }
    function validateStatus(rule: any, value: any, callback: any) {
      if (!tmpUser.value.status) {
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
    function updateUser() {
      (userUpdateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await UpdateUser(tmpUser.value)
          if (response.data.code === 200) {
            ElMessage.success('用户信息更新成功')
            context.emit('get-user-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的用户信息')
        }
      })
    }
    return {
      userUpdateForm,
      show,
      tmpUser,
      rules,
      updateUser
    }
  }
})
</script>
