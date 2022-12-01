<template>
  <div class="app-container">
    <aside class="page-header">
      部门管理
    </aside>
    <div class="filter-container">
      <el-input
        v-model="listQuery.field_lk_name"
        placeholder="部门"
        clearable
        style="width: 200px"
        class="filter-item"
      />
      <el-select
        v-model="listQuery.field_eq_status"
        placeholder="状态"
        clearable
        class="filter-item"
        style="width: 200px"
      >
        <el-option
          v-for="item in statusOptions"
          :key="item"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
      <el-button
        v-waves
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="handleFilter"
      >
        搜索
      </el-button>
      <el-button
        class="filter-item"
        style="margin-left: 10px"
        type="primary"
        icon="el-icon-edit"
        @click="showCreate = true"
      >
        添加
      </el-button>
      <el-button
        v-waves
        :loading="loading"
        class="filter-item"
        type="primary"
        icon="el-icon-download"
        @click="handleDownload"
      >
        导出
      </el-button>
    </div>
    <div class="main-table">
      <el-table
        key="department-table"
        v-loading="listLoading"
        :data="dataList"
        border
        fit
        highlight-current-row
        style="width: 100%"
        @sort-change="sortChange"
      >
        <el-table-column
          label="ID"
          prop="id"
          sortable="custom"
          align="center"
          width="150"
        >
          <template #default="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="部门名称"
          prop="name"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.name }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="部门状态"
          prop="status"
          sortable="custom"
          align="center"
          width="200"
        >
          <template #default="{row}">
            <span>{{ statusOptions.find(o => o.value === row.status)?.label }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建时间"
          prop="created_at"
          sortable="custom"
          align="center"
          width="250"
        >
          <template #default="{row}">
            <span>{{ UnixTime2HumanWithStr(row.created_at) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="更新时间"
          prop="updated_at"
          sortable="custom"
          align="center"
          width="250"
        >
          <template #default="{row}">
            <span>{{ UnixTime2HumanWithStr(row.updated_at) }}</span>
          </template>
        </el-table-column>

        <el-table-column
          label="操作"
          align="center"
          width="300"
          class-name="fixed-width"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="success"
              @click="handleUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="handleDelete(row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div
      class="pagination-footer"
      style="margin-top: 1%"
    >
      <el-pagination
        v-show="total > 0"
        v-model:current-page="listQuery.page"
        v-model:page-size="listQuery.limit"
        :total="total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="getList"
        @current-change="getList"
      />
    </div>
    <div class="dialogs-main">
      <DepartmentCreate
        v-model="showCreate"
        @refresh="getList()"
      />
      <DepartmentUpdate
        v-model="showUpdate"
        :handle-value="currentRow"
        @refresh="getList()"
      />
    </div>
  </div>
</template>
<script lang="ts">
import {
  defineComponent, isRef,
  onMounted,
  reactive,
  ref,
  Ref, toRaw
} from 'vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import type { Department, DepartmentQuery } from '@/model/department'
import { DepartmentStatus } from '@/model/department'
import { deleteDepartment, GetDepartments, ExportDepartment } from '@/api/department'
import DepartmentCreate from './Create.vue'
import DepartmentUpdate from './Update.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteUserRole } from '@/api/permission'
import { UnwrapNestedRefs } from '@vue/reactivity'
export default defineComponent({
  components: {
    DepartmentCreate,
    DepartmentUpdate
  },
  setup() {
    const listQuery: UnwrapNestedRefs<DepartmentQuery> = reactive({
      field_lk_name: '',
      page: 1,
      limit: 10,
      order: '',
      order_by: ''
    })
    const loading = ref(false)
    const listLoading = ref(false)
    const total = ref(0)
    const dataList : Ref<Department[]> = ref([])
    async function handleDownload() {
      loading.value = true
      const res = await ExportDepartment()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '部门.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '部门.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getList() {
      listLoading.value = true
      const res = await GetDepartments(listQuery)
      if (res.data.code === 200) {
        total.value = res.data.total
        dataList.value = res.data.spec
      }
      listLoading.value = false
    }
    function sortChange(column: any) {
      listQuery.order_by = column.prop
      listQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getList()
    }
    function handleFilter() {
      getList()
    }
    const statusOptions = [
      { label: '正常', value: DepartmentStatus.Normal },
      { label: '停用', value: DepartmentStatus.Disabled }
    ]
    const emptyRow: Department = {
      name: '',
      status: DepartmentStatus.Normal
    }
    const showCreate = ref(false)
    const showUpdate = ref(false)
    const currentRow : Ref<Department> = ref(JSON.parse(JSON.stringify(emptyRow)) as Department)
    function handleUpdate(row: Department) {
      currentRow.value = row
      showUpdate.value = true
    }
    function handleDelete(row: Department) {
      ElMessageBox.confirm('是否删除此部门', '确认删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await deleteDepartment(row.id!)
            if (response.data.code === 200) {
              ElMessage.success('用户角色关联删除成功!')
              getList()
            }
          }
        }
      })
    }

    onMounted(() => {
      getList()
    })

    return {
      handleDownload,
      UnixTime2HumanWithStr,
      listQuery,
      listLoading,
      total,
      dataList,
      getList,
      handleFilter,
      sortChange,
      statusOptions,
      showCreate,
      showUpdate,
      handleUpdate,
      currentRow,
      handleDelete

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
</style>
