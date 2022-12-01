<template>
  <div>
    <el-dialog
      v-model="show"
      width="800px"
      title="权限指派"
    >
      <template #footer>
        <el-button
          @click="show = false"
        >
          关闭
        </el-button>
      </template>
      <div style="margin: 10px 0;">
        <el-button
          type="primary"
          @click="handleCreate()"
        >
          指派
        </el-button>
      </div>
      <el-table
        border
        :loading="loading"
        :data="list"
        style="width: 100%"
      >
        <el-table-column
          label="用户名"
          width="180"
        >
          <template #default="scope">
            <span>{{ scope.row.user_name_zh }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
        >
          <template #default="scope">
            <el-button
              type="success"
              @click="confirmDeleteUserRole(scope.row)"
            >
              回收
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
    <el-dialog
      v-model="showCreate"
      width="700px"
      title="权限指派"
    >
      <template #footer>
        <el-button
          :disabled="tempDetail.username === ''"
          @click="confirmUserRoleCreate"
        >
          确认
        </el-button>
        <el-button
          @click="showCreate = false"
        >
          取消
        </el-button>
      </template>
      <el-form>
        <el-form-item
          label="用户名"
          prop="username"
        >
          <UserSearchInput
            v-model="tempDetail.username"
            :role-name="handleValue.name"
            clearable
            :only-name="true"
            placeholder="搜索用户"
          />
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, reactive, Ref, ref, watch } from 'vue'
import type { Role, UserRole, UserRoleQuery, UserRoles } from '@/model/permission'
import { CreateUserRole, DeleteUserRole, GetUserRole } from '@/api/permission'
import { ElMessage } from 'element-plus'
import UserSearchInput from '@/components/user-search-input/index.vue'
export default defineComponent({
  components: {
    UserSearchInput
  },
  props: {
    modelValue: Boolean,
    handleValue: Object
  },
  setup(props, ctx) {
    const show = computed({
      get: () => props.modelValue,
      set: (v) => {
        if (!v) {
          tempDetail.value.username = ''
        }
        ctx.emit('update:modelValue', v)
      }
    })
    const query : UserRoleQuery = reactive({
      field_lk_role_name: '',
      field_lk_username: '',
      field_eq_role_name: '',
      order: '',
      order_by: '',
      page: 0,
      limit: 0
    })
    const loading = ref(false)
    const list : Ref<UserRoles[]> = ref([])
    async function getUserRole() {
      loading.value = true
      query.field_eq_role_name = (props.handleValue as Role).name
      query.field_eq_temporary = 1
      const res = await GetUserRole(query)
      if (res.data.code === 200) {
        list.value = res.data.spec
      }
      loading.value = false
    }
    async function confirmDeleteUserRole(row: UserRoles) {
      if (row.user_role[0]) {
        const res = await DeleteUserRole(row.user_role[0].id!)
        if (res.data.code === 200) {
          ElMessage.success('权限回收成功')
          getUserRole()
        }
      }
    }
    const tempDetail: Ref<UserRole> = ref({
      role_name: '', username: '', user_name_zh: ''
    })
    const showCreate = ref(false)
    function handleCreate() {
      tempDetail.value.role_name = ''
      tempDetail.value.username = ''
      showCreate.value = true
    }
    async function confirmUserRoleCreate() {
      tempDetail.value.temporary = 1
      tempDetail.value.role_name = (props.handleValue as Role).name
      const res = await CreateUserRole(tempDetail.value)
      if (res.data.code === 200) {
        ElMessage.success('权限指派成功')
        getUserRole()
      }
      showCreate.value = false
    }

    watch(() => props.handleValue, function() {
      getUserRole()
    })
    onMounted(() => {
      if (props.handleValue) {
        getUserRole()
      }
    })
    return {
      show,
      loading,
      list,
      confirmDeleteUserRole,
      showCreate,
      tempDetail,
      handleCreate,
      confirmUserRoleCreate

    }
  }
})
</script>

<style scoped>

</style>
