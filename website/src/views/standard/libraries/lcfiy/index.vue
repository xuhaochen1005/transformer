<template>
  <div>
    <aside class="page-header">
      线圈端部距铁轭距离
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="highVol"
          placeholder="额定高压(kV)"
          clearable
          style="width: 150px;margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="lowVol"
          placeholder="额定低压(kV)"
          clearable
          style="width: 150px;margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLcfiyQuery.field_eq_low_vol_coil_from_iron_yoke"
          placeholder="低压线圈端部距铁轭距离(mm)"
          clearable
          style="width: 220px;margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLcfiyQuery.field_eq_high_vol_coil_from_iron_yoke"
          placeholder="高压线圈端部距铁轭距离(mm)"
          clearable
          style="width: 220px;margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLcfiyQuery.field_eq_lcc_creator"
          style="width: 150px;margin-right: 10px"
          size="small"
          class="filter-item"
          :placeholder="'创建者'"
          clearable
        />
        <el-date-picker
          v-model="created_at_range"
          unlink-panels
          style="margin-top: 10px;width:280px"
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
          @click="getLcfiyList"
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
      <lcfiy-create
        v-model="showCreate"
        @get-lcfiy-list="getLcfiyList"
      />
      <lcfiy-view
        v-model="showDetail"
        :lcfiy="lcfiyDetail"
      />
      <lcfiy-update
        v-model="showUpdate"
        :lcfiy="lcfiyDetail"
        @get-lcfiy-list="getLcfiyList"
      />
      <el-table
        v-loading="loading"
        :data="lcfiyList"
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
          label="额定高压下界(kV)"
          prop="rated_high_vol_min"
          sortable="custom"
          align="center"
          width="170px"
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
          width="170px"
        >
          <template #default="{row}">
            <span>{{ row.rated_high_vol_max }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定低压下界(kV)"
          prop="rated_low_vol_min"
          sortable="custom"
          align="center"
          width="170px"
        >
          <template #default="{row}">
            <span>{{ row.rated_low_vol_min }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定低压上界(kV)"
          prop="rated_low_vol_max"
          sortable="custom"
          align="center"
          width="170px"
        >
          <template #default="{row}">
            <span>{{ row.rated_low_vol_max }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="低压线圈端部距铁轭距离(mm)"
          :show-overflow-tooltip="true"
          prop="low_vol_coil_from_iron_yoke"
          sortable="custom"
          align="center"
          width="250px"
        >
          <template #default="{row}">
            <span>{{ row.low_vol_coil_from_iron_yoke }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="高压线圈端部距铁轭距离(mm)"
          prop="high_vol_coil_from_iron_yoke"
          sortable="custom"
          align="center"
          width="250px"
        >
          <template #default="{row}">
            <span>{{ row.high_vol_coil_from_iron_yoke }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lcfiy_creator"
          sortable="custom"
          width="150px"
          align="center"
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
          width="200px"
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
          min-width="250"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="primary"
              @click="showLcfiyDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLcfiyUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLcfiy(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLcfiyQuery.page"
          v-model:page-size="listLcfiyQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager,  next, jumper"
          @size-change="getLcfiyList"
          @current-change="getLcfiyList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlcfiyLibraries, GetStdlcfiyLibraries, ListLcfiyQuery, Lcfiy, ExportStdlcfiyLibraries } from '@/api/standard_libraries/lcfiy'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LcfiyCreate from './lcfiyCreate.vue'
import LcfiyView from './lcfiyView.vue'
import LcfiyUpdate from './lcfiyUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LcfiyCreate,
    LcfiyView,
    LcfiyUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const highVol: Ref<number | undefined> = ref(undefined)
    const lowVol: Ref<number | undefined> = ref(undefined)
    const listLcfiyQuery = reactive({
      // field_eq_rated_high_vol_min: null,
      // field_eq_rated_high_vol_max: null,
      field_lt_rated_high_vol_min: highVol,
      field_ge_rated_high_vol_max: highVol,
      // field_eq_rated_low_vol_min: null,
      // field_eq_rated_low_vol_max: null,
      field_lt_rated_low_vol_min: lowVol,
      field_ge_rated_low_vol_max: lowVol,
      field_eq_low_vol_coil_from_iron_yoke: null,
      field_eq_high_vol_coil_from_iron_yoke: null,
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
    const lcfiyList = ref(Array<Lcfiy>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlcfiyLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '线圈端部距铁轭距离.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '线圈端部距铁轭距离.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLcfiyList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLcfiyQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLcfiyQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLcfiyQuery.field_gt_created_at = 0
        listLcfiyQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLcfiyQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLcfiyQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLcfiyQuery.field_gt_updated_at = 0
        listLcfiyQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlcfiyLibraries(listLcfiyQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lcfiyList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLcfiyQuery.order_by = column.prop
      listLcfiyQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLcfiyList()
    }
    onMounted(() => {
      getLcfiyList()
    })
    function deleteLcfiy(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlcfiyLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLcfiyList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lcfiy: Lcfiy = {
      id: 0,
      rated_high_vol_min: null,
      rated_high_vol_max: null,
      rated_low_vol_min: null,
      rated_low_vol_max: null,
      low_vol_coil_from_iron_yoke: null,
      high_vol_coil_from_iron_yoke: null
    }
    const lcfiyDetail = ref(lcfiy)
    function showLcfiyDetail(lcfiy: Lcfiy) {
      showDetail.value = true
      lcfiyDetail.value = lcfiy
    }
    function showLcfiyUpdate(lcfiy: Lcfiy) {
      showUpdate.value = true
      lcfiyDetail.value = lcfiy
    }
    return {
      handleDownload,
      status,
      loading,
      lcfiyList,
      LcfiyCreate,
      created_at_range,
      updated_at_range,
      listLcfiyQuery,
      total,
      highVol,
      lowVol,
      getLcfiyList,
      sortChange,
      deleteLcfiy,
      showCreate,
      showDetail,
      showUpdate,
      lcfiyDetail,
      showLcfiyUpdate,
      showLcfiyDetail,
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
