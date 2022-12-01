<template>
  <div>
    <aside class="page-header">
      相间距离最小值
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input-number
          v-model="listLpdQuery.field_eq_phase_dist_min"
          placeholder="相间距离最小值（mm）"
          clearable
          style="width: 180px; margin-right: 10px"
          size="small"
        />
        <el-input-number
          v-model="vol"
          placeholder="电压（kV）"
          clearable
          style="width: 180px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLpdQuery.field_eq_lpd_creator"
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
          @click="getLpdList"
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
      <lpd-create
        v-model="showCreate"
        @get-lpd-list="getLpdList"
      />
      <lpd-view
        v-model="showDetail"
        :lpd="lpdDetail"
      />
      <lpd-update
        v-model="showUpdate"
        :lpd="lpdDetail"
        @get-lpd-list="getLpdList"
      />
      <el-table
        v-loading="loading"
        :data="lpdList"
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
          label="相间距离最小值（mm）"
          prop="phase_dist_min"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.phase_dist_min }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="电压下限（kV）"
          prop="vol_min"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.vol_min }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="电压上限（kV）"
          prop="vol_max"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.vol_max }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lpd_creator"
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
          width="250px"
        >
          <template #default="{row}">
            <!--          <el-button-->
            <!--            size="mini"-->
            <!--            type="primary"-->
            <!--            @click="showLpdDetail(row)"-->
            <!--          >-->
            <!--            查看详情-->
            <!--          </el-button>-->
            <el-button
              size="mini"
              type="success"
              @click="showLpdUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLpd(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLpdQuery.page"
          v-model:page-size="listLpdQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLpdList"
          @current-change="getLpdList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLpdLibraries, GetStdLpdLibraries, ListLpdQuery, Lpd, ExportStdlpdLibraries } from '@/api/standard_libraries/lpd'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LpdCreate from './lpdCreate.vue'
import LpdView from './lpdView.vue'
import LpdUpdate from './lpdUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LpdCreate,
    LpdView,
    LpdUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const vol: Ref<number | undefined> = ref(undefined)
    const listLpdQuery = reactive({
      field_eq_phase_dist_min: null,
      // field_eq_vol_min: null,
      // field_eq_vol_max: null,
      field_lt_vol_min: vol,
      field_ge_vol_max: vol,
      field_eq_lpd_creator: null,
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
    const lpdList = ref(Array<Lpd>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlpdLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '相间距离最小值.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '相间距离最小值.xlsx'
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
    async function getLpdList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLpdQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLpdQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLpdQuery.field_gt_created_at = 0
        listLpdQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLpdQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLpdQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLpdQuery.field_gt_updated_at = 0
        listLpdQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLpdLibraries(listLpdQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lpdList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLpdQuery.order_by = column.prop
      listLpdQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLpdList()
    }
    onMounted(() => {
      getLpdList()
    })
    function deleteLpd(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLpdLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLpdList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lpd: Lpd = {
      id: 0,
      phase_dist_min: 0,
      vol_min: 0,
      vol_max: 0
    }
    const lpdDetail = ref(lpd)
    function showLpdDetail(lpd: Lpd) {
      showDetail.value = true
      lpdDetail.value = lpd
    }
    function showLpdUpdate(lpd: Lpd) {
      showUpdate.value = true
      lpdDetail.value = lpd
    }
    return {
      status,
      handleDownload,
      loading,
      lpdList,
      LpdCreate,
      created_at_range,
      updated_at_range,
      listLpdQuery,
      total,
      vol,
      getLpdList,
      sortChange,
      deleteLpd,
      showCreate,
      showDetail,
      showUpdate,
      lpdDetail,
      showLpdUpdate,
      showLpdDetail,
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
