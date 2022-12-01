<template>
  <div>
    <aside class="page-header">
      圆铜(铝)线规格
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input-number
          v-model="listLrwsQuery.field_eq_std_diameter"
          placeholder="标称直径（mm）"
          clearable
          style="width: 180px; margin-right: 10px"
          size="small"
        />
        <el-input-number
          v-model="listLrwsQuery.field_eq_std_diameter"
          placeholder="截面积（mm^2）"
          clearable
          style="width: 180px; margin-right: 10px"
          size="small"
        />
        <el-input-number
          v-model="listLrwsQuery.field_eq_std_diameter"
          placeholder="最大外径（mm）"
          clearable
          style="width: 180px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLrwsQuery.field_eq_lrws_creator"
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
          @click="getLrwsList"
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
      <lrws-create
        v-model="showCreate"
        @get-lrws-list="getLrwsList"
      />
      <lrws-view
        v-model="showDetail"
        :lrws="lrwsDetail"
      />
      <lrws-update
        v-model="showUpdate"
        :lrws="lrwsDetail"
        @get-lrws-list="getLrwsList"
      />
      <el-table
        v-loading="loading"
        :data="lrwsList"
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
          label="标称直径（mm）"
          prop="std_diameter"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.std_diameter }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="截面积（mm^2）"
          prop="section_area"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.section_area }}mm^2</span>
          </template>
        </el-table-column>
        <el-table-column
          label="最大外径（mm）"
          prop="max_diameter"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.max_diameter }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lrws_creator"
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
  <!--            @click="showLrwsDetail(row)"-->
  <!--          >-->
  <!--            查看详情-->
  <!--          </el-button>-->
            <el-button
              size="mini"
              type="success"
              @click="showLrwsUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLrws(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLrwsQuery.page"
          v-model:page-size="listLrwsQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLrwsList"
          @current-change="getLrwsList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLrwsLibraries, GetStdLrwsLibraries, ListLrwsQuery, Lrws, ExportStdlrwsLibraries } from '@/api/standard_libraries/lrws'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LrwsCreate from './lrwsCreate.vue'
import LrwsView from './lrwsView.vue'
import LrwsUpdate from './lrwsUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LrwsCreate,
    LrwsView,
    LrwsUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLrwsQuery = reactive({
      field_eq_std_diameter: null,
      field_eq_section_area: null,
      field_eq_max_diameter: null,
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
    const lrwsList = ref(Array<Lrws>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlrwsLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '圆铜（铝）线规格.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '圆铜（铝）线规格.xlsx'
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
    async function getLrwsList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLrwsQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLrwsQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLrwsQuery.field_gt_created_at = 0
        listLrwsQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLrwsQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLrwsQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLrwsQuery.field_gt_updated_at = 0
        listLrwsQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLrwsLibraries(listLrwsQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lrwsList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLrwsQuery.order_by = column.prop
      listLrwsQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLrwsList()
    }
    onMounted(() => {
      getLrwsList()
    })
    function deleteLrws(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLrwsLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLrwsList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lrws: Lrws = {
      id: 0,
      std_diameter: 0,
      section_area: 0,
      max_diameter: 0
    }
    const lrwsDetail = ref(lrws)
    function showLrwsDetail(lrws: Lrws) {
      showDetail.value = true
      lrwsDetail.value = lrws
    }
    function showLrwsUpdate(lrws: Lrws) {
      showUpdate.value = true
      lrwsDetail.value = lrws
    }
    return {
      status,
      handleDownload,
      loading,
      lrwsList,
      LrwsCreate,
      created_at_range,
      updated_at_range,
      listLrwsQuery,
      total,
      getLrwsList,
      sortChange,
      deleteLrws,
      showCreate,
      showDetail,
      showUpdate,
      lrwsDetail,
      showLrwsUpdate,
      showLrwsDetail,
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
