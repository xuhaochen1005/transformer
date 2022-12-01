<template>
  <div>
    <aside class="page-header">
      工频耐压
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="highVol"
          placeholder="额定高压(kV)"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLifhvQuery.field_eq_industry_freq_hold_vol"
          placeholder="工频耐压(kV)"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLifhvQuery.field_eq_lifhv_creator"
          style="width: 200px"
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
          @click="getLifhvList"
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
      <lifhv-create
        v-model="showCreate"
        @get-lifhv-list="getLifhvList"
      />
      <lifhv-view
        v-model="showDetail"
        :lifhv="lifhvDetail"
      />
      <lifhv-update
        v-model="showUpdate"
        :lifhv="lifhvDetail"
        @get-lifhv-list="getLifhvList"
      />
      <el-table
        v-loading="loading"
        :data="lifhvList"
        border
        fit
        highlight-current-row
        @sort-change="sortChange"
      >
        <el-table-column
          label="编号"
          prop="id"
          sortable="custom"
          align="center"
          min-width="75px"
        >
          <template #default="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定高压下界(kV)"
          prop="rated_high_vol_min"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.rated_high_vol_min }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定高压上界(kV)"
          prop="rated_high_vol_max"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.rated_high_vol_max }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="工频耐压"
          prop="industry_freq_hold_vol"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.industry_freq_hold_vol }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lifhv_creator"
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
          min-width="200px"
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
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ UnixTime2HumanWithStr(row.updated_at) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          align="center"
          fixed="right"
          min-width="250px"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="success"
              @click="showLifhvUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLifhv(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLifhvQuery.page"
          v-model:page-size="listLifhvQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLifhvList"
          @current-change="getLifhvList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlifhvLibraries, GetStdlifhvLibraries, ListLifhvQuery, Lifhv, ExportStdlifhvLibraries } from '@/api/standard_libraries/lifhv'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LifhvCreate from './lifhvCreate.vue'
import LifhvView from './lifhvView.vue'
import LifhvUpdate from './lifhvUpdate.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import dayjs from 'dayjs'

export default defineComponent({
  components: {
    UserSearchInput,
    LifhvCreate,
    LifhvView,
    LifhvUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const highVol: Ref<number | undefined> = ref(undefined)
    const listLifhvQuery = reactive({
      field_eq_industry_freq_hold_vol: null,
      field_lt_rated_high_vol_min: highVol,
      field_ge_rated_high_vol_max: highVol,
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
    const lifhvList = ref(Array<Lifhv>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlifhvLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '工频耐压.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '工频耐压.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLifhvList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLifhvQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLifhvQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLifhvQuery.field_gt_created_at = 0
        listLifhvQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLifhvQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLifhvQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLifhvQuery.field_gt_updated_at = 0
        listLifhvQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlifhvLibraries(listLifhvQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lifhvList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLifhvQuery.order_by = column.prop
      listLifhvQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLifhvList()
    }
    onMounted(() => {
      getLifhvList()
    })
    function deleteLifhv(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlifhvLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLifhvList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lifhv: Lifhv = {
      id: 0,
      industry_freq_hold_vol: null,
      rated_high_vol_min: null,
      rated_high_vol_max: null
    }
    const lifhvDetail = ref(lifhv)
    function showLifhvDetail(lifhv: Lifhv) {
      showDetail.value = true
      lifhvDetail.value = lifhv
    }
    function showLifhvUpdate(lifhv: Lifhv) {
      showUpdate.value = true
      lifhvDetail.value = lifhv
    }
    return {
      status,
      handleDownload,
      loading,
      lifhvList,
      LifhvCreate,
      created_at_range,
      updated_at_range,
      listLifhvQuery,
      total,
      highVol,
      getLifhvList,
      sortChange,
      deleteLifhv,
      showCreate,
      showDetail,
      showUpdate,
      lifhvDetail,
      showLifhvUpdate,
      showLifhvDetail,
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
