<template>
  <div>
    <aside class="page-header">
      相数
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLpnQuery.field_eq_phase_num"
          filterable
          allow-create
          placeholder="相数"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="单相"
            value="单相"
          />
          <el-option
            label="三相"
            value="三相"
          />
        </el-select>
        <el-select
          v-model="listLpnQuery.field_eq_phase_num_sign"
          placeholder="相数代号"
          filterable
          allow-create
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="D"
            value="D"
          />
          <el-option
            label="S"
            value="S"
          />
        </el-select>
        <UserSearchInput
          v-model="listLpnQuery.field_eq_lpn_creator"
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
          @click="getLpnList"
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
      <lpn-create
        v-model="showCreate"
        @get-lpn-list="getLpnList"
      />
      <lpn-view
        v-model="showDetail"
        :lpn="lpnDetail"
      />
      <lpn-update
        v-model="showUpdate"
        :lpn="lpnDetail"
        @get-lpn-list="getLpnList"
      />
      <el-table
        v-loading="loading"
        :data="lpnList"
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
          label="相数"
          prop="phase_num"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.phase_num }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="相数代号"
          prop="phase_num_sign"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.phase_num_sign }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lpn_creator"
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
          width="300px"
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
          width="300px"
        >
          <template #default="{row}">
            <span>{{ UnixTime2HumanWithStr(row.updated_at) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          align="center"
          fixed="right"
          width="340px"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="success"
              @click="showLpnUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLpn(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLpnQuery.page"
          v-model:page-size="listLpnQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLpnList"
          @current-change="getLpnList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLpnLibraries, GetStdLpnLibraries, ListLpnQuery, Lpn, ExportStdlpnLibraries } from '@/api/standard_libraries/lpn'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LpnCreate from './lpnCreate.vue'
import LpnView from './lpnView.vue'
import LpnUpdate from './lpnUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LpnCreate,
    LpnView,
    LpnUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLpnQuery = reactive({
      field_eq_phase_num: null,
      field_eq_phase_num_sign: null,
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
    const lpnList = ref(Array<Lpn>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlpnLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '相数.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '相数.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function getLpnList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLpnQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLpnQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLpnQuery.field_gt_created_at = 0
        listLpnQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLpnQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLpnQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLpnQuery.field_gt_updated_at = 0
        listLpnQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLpnLibraries(listLpnQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lpnList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLpnQuery.order_by = column.prop
      listLpnQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLpnList()
    }
    onMounted(() => {
      getLpnList()
    })
    function deleteLpn(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLpnLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLpnList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lpn: Lpn = {
      id: 0,
      phase: 0,
      phase_num: '',
      phase_num_sign: ''
    }
    const lpnDetail = ref(lpn)
    function showLpnDetail(lpn: Lpn) {
      showDetail.value = true
      lpnDetail.value = lpn
    }
    function showLpnUpdate(lpn: Lpn) {
      showUpdate.value = true
      lpnDetail.value = lpn
    }
    return {
      status,
      handleDownload,
      loading,
      lpnList,
      LpnCreate,
      created_at_range,
      updated_at_range,
      listLpnQuery,
      total,
      getLpnList,
      sortChange,
      deleteLpn,
      showCreate,
      showDetail,
      showUpdate,
      lpnDetail,
      showLpnUpdate,
      showLpnDetail,
      UnixTime2Human,
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
