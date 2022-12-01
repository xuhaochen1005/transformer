1<template>
  <div>
    <aside class="page-header">
      引线损耗
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLlsQuery.field_eq_rated_capacity"
          filterable
          allow-create
          placeholder="额定容量(kVA)"
          clearable
          style="width: 220px; margin-right: 10px"
          size="small"
        >
          <el-option
            v-for="item in capacity"
            :key="item"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <UserSearchInput
          v-model="listLlsQuery.field_eq_lls_creator"
          style="width: 180px"
          size="small"
          class="filter-item"
          :placeholder="'创建者'"
          clearable
        />
        <el-date-picker
          v-model="created_at_range"
          unlink-panels
          style="margin-top: 10px;width:280px;margin-left:10px"
          type="datetimerange"
          size="small"
          start-placeholder="创建起始日期"
          end-placeholder="创建截止日期"
        />
        <el-date-picker
          v-model="updated_at_range"
          unlink-panels
          style="margin-top: 10px;width:280px;margin-left:10px"
          type="datetimerange"
          size="small"
          start-placeholder="更新起始日期"
          end-placeholder="更新截止日期"
        />
        <el-button
          v-waves
          style="margin-left:20px"
          type="primary"
          icon="el-icon-search"
          @click="getLlsList"
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
      <lls-create
        v-model="showCreate"
        @get-lls-list="getLlsList"
      />
      <lls-view
        v-model="showDetail"
        :lls="llsDetail"
      />
      <lls-update
        v-model="showUpdate"
        :lls="llsDetail"
        @get-lls-list="getLlsList"
      />
      <el-table
        v-loading="loading"
        :data="llsList"
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
          label="额定容量(kVA)"
          prop="rated_capacity"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.rated_capacity }}kVA</span>
          </template>
        </el-table-column>
        <el-table-column
          label="引线损耗（W）"
          prop="lead_loss"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.lead_loss }}W</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lls_creator"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.creator_user?.user_name_zh }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建时间"
          prop="created_at"
          sortable="custom"
          align="center"
          width="180px"
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
          width="180px"
        >
          <template #default="{row}">
            <span>{{ UnixTime2HumanWithStr(row.updated_at) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          align="center"
          fixed="right"
          width="250px"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="success"
              @click="showLlsUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLls(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLlsQuery.page"
          v-model:page-size="listLlsQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLlsList"
          @current-change="getLlsList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  DeleteStdLlsLibraries,
  ExportStdllsLibraries,
  GetStdLlsLibraries,
  ListLlsQuery,
  Lls
} from '@/api/standard_libraries/lls'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LlsCreate from './llsCreate.vue'
import LlsView from './llsView.vue'
import LlsUpdate from './llsUpdate.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import dayjs from 'dayjs'
import { GetStdLrrLibraries, Lrr } from '@/api/standard_libraries/lrr'
import { ExportStdlifhvLibraries } from '@/api/standard_libraries/lifhv'

export default defineComponent({
  components: {
    UserSearchInput,
    LlsCreate,
    LlsView,
    LlsUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLlsQuery = reactive({
      field_eq_rated_capacity: null,
      field_eq_lead_loss: null,
      field_gt_created_at: 0,
      field_lt_created_at: 0,
      field_gt_updated_at: 0,
      field_lt_updated_at: 0,
      page: 1,
      limit: 10,
      order: '',
      order_by: ''
    })
    const loading = ref(false)
    const llsList = ref(Array<Lls>())
    const total = ref(0)
    const capacity : Ref<{label: string, value: number}[]> = ref([])
    const getStdLlsLibraries = async () => {
      const res = await GetStdLlsLibraries({
        field_eq_rated_capacity: 0,
        field_eq_lead_loss: 0,
        limit: 9999,
        page: 1,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        capacity.value = res.data.spec.map((i: Lls) => {
          return {
            label: i.rated_capacity + 'kVA',
            value: i.rated_capacity
          }
        })
      }
    }
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }

    async function handleDownload() {
      loading.value = true
      const res = await ExportStdllsLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '引线损耗.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '引线损耗.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLlsList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLlsQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLlsQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLlsQuery.field_gt_created_at = 0
        listLlsQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLlsQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLlsQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLlsQuery.field_gt_updated_at = 0
        listLlsQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLlsLibraries(listLlsQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        llsList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLlsQuery.order_by = column.prop
      listLlsQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLlsList()
    }
    onMounted(() => {
      getLlsList()
    })
    function deleteLls(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLlsLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLlsList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lls: Lls = {
      id: 0,
      rated_capacity: 0,
      lead_loss: 0
    }
    const llsDetail = ref(lls)
    function showLlsDetail(lls: Lls) {
      showDetail.value = true
      llsDetail.value = lls
    }
    function showLlsUpdate(lls: Lls) {
      showUpdate.value = true
      llsDetail.value = lls
    }
    const paramLoading = ref(true)
    onMounted(async () => {
      await Promise.all([
        // 额定容量
        getStdLlsLibraries()
      ])
      paramLoading.value = false
    })
    return {
      status,
      handleDownload,
      loading,
      llsList,
      LlsCreate,
      created_at_range,
      updated_at_range,
      listLlsQuery,
      total,
      capacity,
      UnixTime2Human,
      getLlsList,
      sortChange,
      deleteLls,
      showCreate,
      showDetail,
      showUpdate,
      llsDetail,
      showLlsUpdate,
      showLlsDetail,
      UnixTime2HumanWithStr
    }
  }
})
</script>
<style lang="scss" scoped>
  .page-header {
    line-height: 24px;
    font-size: 24px;
    color: #409eff;
    background: none;
    padding: 0;
    margin:20px;
  }
</style>
