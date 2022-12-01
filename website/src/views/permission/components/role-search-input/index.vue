<template>
  <el-select
    ref="selectRef"
    v-model="handleValue"
    :model-value="handleValue"
    filterable
    :clearable="clearable"
    remote
    :multiple="false"
    :remote-method="searchRoles"
    :loading="roleSearchLoading"
    :placeholder="placeholder"
    style="width: 200px"
    @visible-change="handleVisiable"
    @change="handleChange"
  >
    <el-option
      v-for="item in searchRoleList"
      :key="item"
      :label="item.cn_name"
      :value="item.name"
    />
  </el-select>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, reactive, Ref, ref, watch, watchEffect } from 'vue'
import { Role, RoleQuery } from '@/model/permission'
import { GetRoleList } from '@/api/permission'

export default defineComponent({
  props: {
    modelValue: String,
    placeholder: {
      type: String,
      default: () => '搜索角色'
    },
    clearable: {
      type: Boolean,
      default: true
    },
    username: {
      type: String,
      default: () => ''
    }
  },
  emits: ['update:modelValue', 'handleChange', 'handleVisible', 'ref'],
  setup(props, ctx) {
    const handleValue = computed({
      get: () => props.modelValue,
      set: (value) => {
        ctx.emit('update:modelValue', value)
      }
    })
    const roleSearchLoading = ref(false)
    const searchRoleList: Ref<Role[]> = ref([])
    const searchRoleQuery: RoleQuery = reactive({
      username: props.username,
      field_eq_name: '',
      field_lk_cn_name: '',
      page: 1,
      limit: 5,
      order: '',
      order_by: ''
    })
    // 初始化角色名称
    async function initCurrentRole() {
      searchRoleQuery.username = props.username
      searchRoleQuery.field_eq_name = props.modelValue
      searchRoleQuery.limit = 5
      const res = await GetRoleList(searchRoleQuery)
      if (res.data.code === 200) {
        searchRoleList.value = res.data.spec
      }
      searchRoleQuery.limit = 5
      searchRoleQuery.field_eq_name = ''
    }
    watch(() => props.modelValue, initCurrentRole)
    async function searchRoles(queryString: string) {
      roleSearchLoading.value = true
      searchRoleQuery.field_lk_cn_name = queryString
      const res = await GetRoleList(searchRoleQuery)
      if (res.data.code === 200) {
        searchRoleList.value = res.data.spec
      }
      roleSearchLoading.value = false
    }
    function handleChange() {
      ctx.emit('handleChange', true)
    }
    const selectRef = ref(null)
    function handleVisiable(value: boolean) {
      ctx.emit('handleVisible', value)
    }
    onMounted(async () => {
      initCurrentRole()
      ctx.emit('ref', selectRef.value)
    })
    return {
      handleValue,
      roleSearchLoading,
      searchRoleList,
      selectRef,
      searchRoles,
      handleChange,
      handleVisiable
    }
  }
})
</script>

<style scoped>

</style>
