<template>
  <el-select
    v-model="handleValue"
    :model-value="handleValue===0?null:handleValue"
    :multiple="multiple"
    filterable
    clearable
    remote
    :remote-method="searchUser"
    :loading="loading"
    :placeholder="placeholder"
  >
    <el-option
      v-for="item in selectOptions"
      :key="item.value"
      :label="item.label"
      :value="item.value"
    >
      <span style="float: left">{{ item.label }}</span>
      <span style="float: right; color: #8492a6; font-size: 13px">{{ onlyName ? item.value : item.name }}</span>
    </el-option>
  </el-select>
</template>

<script lang="ts">
import {
  defineComponent,
  onMounted,
  computed,
  ref, watch, PropType
} from 'vue'
import { GetUsers } from '@/api/user'
import type { User, UserQuery } from '@/model/common'
import { RoleList } from '@/model/permission'

export default defineComponent({
  props: {
    placeholder: {
      type: String,
      default: '搜索用户'
    },
    multiple: {
      type: Boolean,
      default: false
    },
    // 只提供用户名，返回也是用户名
    onlyName: {
      type: Boolean,
      default: false
    },
    modelValue: {
      type: null,
      default: null
    },
    roles: {
      type: Array as PropType<RoleList[]>,
      default: () => []
    },
    roleName: {
      type: String,
      default: () => ''
    }
  },
  emits: ['update:modelValue'],
  setup(props, ctx) {
    const loading = ref(false)
    const userList = ref(Array<User>())
    const selectOptions = computed(() => userList.value.map((user) => {
      return { value: props.onlyName ? user.name : user.id, label: user.user_name_zh, name: user.name }
    }))
    const handleValue : any = computed({
      get: () => props.modelValue,
      set: (value) => {
        ctx.emit('update:modelValue', value)
      }
    })
    const userQuery : UserQuery = {
      'field_eq_user.id': undefined,
      'field_lk_user.user_name_zh': '',
      field_eq_name: '',
      role_name: props.roleName,
      roles: props.roles,
      'field_lk_user.name': '',
      field_eq_department_id: null,
      'field_eq_user.status': null,
      page: 1,
      limit: 10,
      order: '',
      order_by: ''
    }
    async function searchUser(query : string | number) {
      const timer = setTimeout(() => {
        loading.value = true
      }, 300)
      if (typeof query === 'number') {
        userQuery['field_lk_user.user_name_zh'] = undefined
        userQuery['field_eq_user.id'] = query
      } else {
        userQuery['field_lk_user.user_name_zh'] = query
        userQuery['field_eq_user.id'] = undefined
      }
      const response = await GetUsers(userQuery)
      if (response.data.code === 200) {
        userList.value = response.data.spec
      }
      clearTimeout(timer)
      loading.value = false
    }
    async function propsUpdate() {
      handleValue.value = props.modelValue
      userQuery.role_name = props.roleName
      await searchUser(props.modelValue)
    }
    watch(props, propsUpdate)
    onMounted(propsUpdate)
    return {
      loading,
      handleValue,
      selectOptions,
      searchUser
    }
  }
})
</script>

<style lang="scss" scoped>

</style>
