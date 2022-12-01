<template>
  <div style="margin: 10px">
    <aside class="page-header">
      用户管理
    </aside>
    <div style="margin-bottom: 10px">
      <el-input
        v-model="query.field_eq_user_name_zh"
        placeholder="用户名"
        clearable
        style="width: 200px; margin-right: 10px"
        @keyup.enter="getUserList"
      />
      <el-select
        v-model="query.field_eq_department_id"
        placeholder="部门"
        clearable
        style="width: 200px; margin-right: 10px"
      >
        <el-option
          v-for="item in departmentOptions"
          :key="item"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
      <el-select
        v-model="query['field_eq_user.status']"
        placeholder="状态"
        clearable
        style="width: 200px; margin-right: 10px"
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
      <el-button
        v-waves
        type="primary"
        icon="el-icon-search"
        @click="getUserList"
      >
        搜索
      </el-button>
      <el-button
        v-waves
        type="primary"
        icon="el-icon-edit"
        @click="showCreate = true"
      >
        添加
      </el-button>
    </div>
    <user-create
      v-model="showCreate"
      :department-options="departmentOptions"
      @get-user-list="getUserList"
    />
    <user-view
      v-model="showDetail"
      :user="userDetail"
    />
    <user-update
      v-model="showUpdate"
      :department-options="departmentOptions"
      :user="userDetail"
      @get-user-list="getUserList"
    />
    <el-table
      v-loading="loading"
      :data="userList"
      border
      fit
      highlight-current-row
      style="width: 100%"
      @sort-change="sortChange"
    >
      <el-table-column
        label="编号"
        prop="id"
        sortable="custom"
        align="center"
        width="75px"
      >
        <template #default="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="账号"
        prop="name"
        sortable="custom"
        align="center"
        width="150px"
      >
        <template #default="{row}">
          <span>{{ row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="名称"
        prop="user_name_zh"
        sortable="custom"
        align="center"
        width="150px"
      >
        <template #default="{row}">
          <span>{{ row.user_name_zh }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="部门"
        prop="department_id"
        sortable="custom"
        align="center"

        min-width="100px"
      >
        <template #default="{row}">
          <span>{{ row.department_name }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="邮箱"
        prop="email"
        sortable="custom"
        align="center"
        min-width="150px"
      >
        <template #default="{row}">
          <span>{{ row.email }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="手机"
        prop="telephone"
        sortable="custom"
        align="center"
        min-width="100px"
      >
        <template #default="{row}">
          <span>{{ row.telephone }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="状态"
        prop="status"
        align="center"
      >
        <template #default="{row}">
          <span>{{ status.get(row.status) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="创建时间"
        prop="created_at"
        align="center"
        min-width="150px"
      >
        <template #default="{row}">
          <span>{{ UnixTime2HumanWithStr(row.created_at) }}</span>
        </template>
      </el-table-column>
      <!--      <el-table-column-->
      <!--        label="更新时间"-->
      <!--        prop="updated_at"-->
      <!--        align="center"-->
      <!--        min-width="100px"-->
      <!--      >-->
      <!--        <template #default="{row}">-->
      <!--          <span>{{ UnixTime2HumanWithStr(row.updated_at) }}</span>-->
      <!--        </template>-->
      <!--      </el-table-column>-->
      <el-table-column
        label="操作"
        align="center"
        min-width="230px"
      >
        <template #default="{row}">
          <el-button
            size="mini"
            type="primary"
            @click="showUserDetail(row)"
          >
            查看详情
          </el-button>
          <el-button
            size="mini"
            type="success"
            @click="showUserUpdate(row)"
          >
            修改
          </el-button>
          <el-button
            size="mini"
            type="danger"
            @click="deleteUser(row.id)"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <div style="margin-top: 20px">
      <el-pagination
        v-show="total > 0"
        v-model:current-page="query.page"
        v-model:page-size="query.limit"
        :total="total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="getUserList"
        @current-change="getUserList"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, Ref, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { GetUsers, DeleteUser } from '@/api/user'
import { User } from '@/model/common'
import UserCreate from './userCreate.vue'
import UserView from './userView.vue'
import UserUpdate from './userUpdate.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import type { Department, DepartmentQuery } from '@/model/department'
import { GetDepartments } from '@/api/department'

export default defineComponent({
  components: {
    UserCreate,
    UserView,
    UserUpdate
  },
  setup() {
    const status = new Map([[1, '正常'], [2, '停用']])
    const query = reactive({
      field_eq_user_name_zh: '',
      field_eq_name: '',
      field_eq_department_id: null,
      'field_eq_user.status': null,
      page: 1,
      limit: 10,
      order: '',
      order_by: ''
    })
    const loading = ref(false)
    const userList = ref(Array<User>())
    const total = ref(0)
    async function getUserList() {
      loading.value = true
      console.log(query)
      const response = await GetUsers(query)
      console.log(response)
      if (response.data.code === 200) {
        total.value = response.data.total
        userList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    const departments : Ref<Department[]> = ref([])
    const departmentOptions = computed(() => {
      return departments.value.map((de) => { return { label: de.name, value: de.id } })
    })
    const departmentQuery : DepartmentQuery = {
      field_lk_name: '',
      page: 1,
      limit: 9999,
      order: '',
      order_by: ''
    }
    async function getDepartments() {
      const res = await GetDepartments(departmentQuery)
      if (res.data.code === 200) {
        departments.value = res.data.spec
      }
    }
    function sortChange(column: any) {
      query.order_by = column.prop
      query.order = column.order === 'descending' ? 'desc' : 'asc'
      getUserList()
    }
    onMounted(() => {
      getUserList()
    })
    function deleteUser(id: number) {
      ElMessageBox.confirm('是否删除此用户，请注意此操作不可恢复！', '确认删除用户', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteUser(id)
            if (response.data.code === 200) {
              ElMessage.success('用户删除成功!')
              await getUserList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const user: User = {
      user_name_zh: '',
      name: '',
      telephone: '',
      email: '',
      department_id: 1,
      status: 0
    }
    const userDetail = ref(user)
    function showUserDetail(user: User) {
      showDetail.value = true
      userDetail.value = user
    }
    function showUserUpdate(user: User) {
      showUpdate.value = true
      userDetail.value = user
    }
    onMounted(() => {
      getDepartments()
    })
    return {
      status,
      query,
      loading,
      userList,
      total,
      departments,
      departmentOptions,
      UnixTime2HumanWithStr,
      getUserList,
      sortChange,
      deleteUser,
      showCreate,
      showDetail,
      showUpdate,
      userDetail,
      showUserUpdate,
      showUserDetail
    }
  }
})
</script>
<style lang="scss" scoped>

  .filter-container {
    .container-row:first-child {
      margin-bottom: 10px;
    }
    .filter-item:not(button) {
      margin-right: 10px;
    }
  }
  .form-icon {
    font-size: 18px;
    &.warning {
      color: $yellow;
    }
  }
  .custom-table-popper {
    max-width: 500px;
  }
</style>
