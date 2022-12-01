<template>
  <div>
    <aside class="page-header">
      冲击电压
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input-number
          v-model="highVol"
          placeholder="额定高压（kV）"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        />
        <el-input-number
          v-model="listLshvQuery.field_eq_shock_hold_vol"
          placeholder="冲击电压（kV）"
          clearable
          style="width: 220px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLshvQuery.field_eq_lshv_creator"
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
          @click="getLshvList"
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
      <lshv-create
        v-model="showCreate"
        @get-lshv-list="getLshvList"
      />
      <lshv-view
        v-model="showDetail"
        :lshv="lshvDetail"
      />
      <lshv-update
        v-model="showUpdate"
        :lshv="lshvDetail"
        @get-lshv-list="getLshvList"
      />
      <el-table
        v-loading="loading"
        :data="lshvList"
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
          label="额定高压下界（kV）"
          prop="std_diameter"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.rated_high_vol_min }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定高压上界（kV）"
          prop="section_area"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.rated_high_vol_max }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="冲击电压（kV）"
          prop="max_diameter"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.shock_hold_vol }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lshv_creator"
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
          width="200px"
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
            <!--          <el-button-->
            <!--            size="mini"-->
            <!--            type="primary"-->
            <!--            @click="showLshvDetail(row)"-->
            <!--          >-->
            <!--            查看详情-->
            <!--          </el-button>-->
            <el-button
              size="mini"
              type="success"
              @click="showLshvUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLshv(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLshvQuery.page"
          v-model:page-size="listLshvQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLshvList"
          @current-change="getLshvList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLshvLibraries, GetStdLshvLibraries, ListLshvQuery, Lshv, ExportStdlshvLibraries } from '@/api/standard_libraries/lshv'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LshvCreate from './lshvCreate.vue'
import LshvView from './lshvView.vue'
import LshvUpdate from './lshvUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LshvCreate,
    LshvView,
    LshvUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const highVol: Ref<number | undefined> = ref(undefined)
    const listLshvQuery = reactive({
      // field_eq_rated_high_vol_min: null,
      // field_eq_rated_high_vol_max: null,
      field_lt_rated_high_vol_min: highVol,
      field_ge_rated_high_vol_max: highVol,
      field_eq_shock_hold_vol: null,
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
    const lshvList = ref(Array<Lshv>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlshvLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '冲击电压.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '冲击电压.xlsx'
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
    async function getLshvList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLshvQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLshvQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLshvQuery.field_gt_created_at = 0
        listLshvQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLshvQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLshvQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLshvQuery.field_gt_updated_at = 0
        listLshvQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLshvLibraries(listLshvQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lshvList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLshvQuery.order_by = column.prop
      listLshvQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLshvList()
    }
    onMounted(() => {
      getLshvList()
    })
    function deleteLshv(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLshvLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLshvList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lshv: Lshv = {
      id: 0,
      rated_high_vol_min: 0,
      rated_high_vol_max: 0,
      shock_hold_vol: 0
    }
    const lshvDetail = ref(lshv)
    function showLshvDetail(lshv: Lshv) {
      showDetail.value = true
      lshvDetail.value = lshv
    }
    function showLshvUpdate(lshv: Lshv) {
      showUpdate.value = true
      lshvDetail.value = lshv
    }
    return {
      status,
      handleDownload,
      loading,
      lshvList,
      LshvCreate,
      created_at_range,
      updated_at_range,
      listLshvQuery,
      total,
      highVol,
      getLshvList,
      sortChange,
      deleteLshv,
      showCreate,
      showDetail,
      showUpdate,
      lshvDetail,
      showLshvUpdate,
      showLshvDetail,
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
