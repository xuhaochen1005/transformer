<template>
  <div>
    <aside class="page-header">
      线圈和风道绝缘
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLcadiQuery.field_eq_wind_type"
          filterable
          allow-create
          placeholder="绕制类型"
          clearable
          style="width: 220px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="箔绕"
            value="箔绕"
          />
          <el-option
            label="线绕"
            value="线绕"
          />
        </el-select>
        <UserSearchInput
          v-model="listLcadiQuery.field_eq_lcadi_creator"
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
          @click="getLcadiList"
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
      <lcadi-create
        v-model="showCreate"
        @get-lcadi-list="getLcadiList"
      />
      <lcadi-view
        v-model="showDetail"
        :lcadi="lcadiDetail"
      />
      <lcadi-update
        v-model="showUpdate"
        :lcadi="lcadiDetail"
        @get-lcadi-list="getLcadiList"
      />
      <el-table
        v-loading="loading"
        :data="lcadiList"
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
          label="绕制类型"
          prop="wind_type"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.wind_type }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="线圈内层绝缘（mm）"
          prop="coil_inner_insulate"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.coil_inner_insulate }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="线圈外层绝缘（mm）"
          prop="coil_outer_insulate"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.coil_outer_insulate }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="风道厚度可选值（mm）"
          prop="air_duct_thick"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.air_duct_thick }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="风道内外层绝缘（mm）"
          prop="air_duct_in_out_insulate"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.air_duct_in_out_insulate }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lcadi_creator"
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
              type="primary"
              @click="showLcadiDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLcadiUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLcadi(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
        <div style="margin-top: 20px">
          <el-pagination
            v-show="total > 0"
            v-model:currentPage="listLcadiQuery.page"
            v-model:page-size="listLcadiQuery.limit"
            :total="total"
            :page-sizes="[10, 20, 50, 100]"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="getLcadiList"
            @current-change="getLcadiList"
          />
        </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLcadiLibraries, GetStdLcadiLibraries, ListLcadiQuery, Lcadi, ExportStdlcadiLibraries } from '@/api/standard_libraries/lcadi'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LcadiCreate from './lcadiCreate.vue'
import LcadiView from './lcadiView.vue'
import LcadiUpdate from './lcadiUpdate.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import dayjs from 'dayjs'
export default defineComponent({
  components: {
    UserSearchInput,
    LcadiCreate,
    LcadiView,
    LcadiUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLcadiQuery = reactive({
      field_eq_wind_type: null,
      field_lk_coil_inner_insulate: null,
      field_lk_coil_outer_insulate: null,
      field_lk_air_duct_thick: null,
      field_lk_air_duct_in_out_insulate: null,
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
    const lcadiList = ref(Array<Lcadi>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlcadiLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '线圈与风道绝缘.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '线圈与风道绝缘.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLcadiList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLcadiQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLcadiQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLcadiQuery.field_gt_created_at = 0
        listLcadiQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLcadiQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLcadiQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLcadiQuery.field_gt_updated_at = 0
        listLcadiQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLcadiLibraries(listLcadiQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lcadiList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLcadiQuery.order_by = column.prop
      listLcadiQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLcadiList()
    }
    onMounted(() => {
      getLcadiList()
    })
    function deleteLcadi(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLcadiLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLcadiList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lcadi: Lcadi = {
      id: 0,
      wind_type: '',
      coil_inner_insulate: 0,
      coil_outer_insulate: 0,
      air_duct_thick: '',
      air_duct_in_out_insulate: 0
    }
    const lcadiDetail = ref(lcadi)
    function showLcadiDetail(lcadi: Lcadi) {
      showDetail.value = true
      lcadiDetail.value = lcadi
    }
    function showLcadiUpdate(lcadi: Lcadi) {
      showUpdate.value = true
      lcadiDetail.value = lcadi
    }
    return {
      handleDownload,
      status,
      loading,
      lcadiList,
      LcadiCreate,
      created_at_range,
      updated_at_range,
      listLcadiQuery,
      total,
      UnixTime2Human,
      getLcadiList,
      sortChange,
      deleteLcadi,
      showCreate,
      showDetail,
      showUpdate,
      lcadiDetail,
      showLcadiUpdate,
      showLcadiDetail,
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
