<template>
  <div>
    <aside class="page-header">
      硅钢片性能
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLspdQuery.field_eq_stack_type"
          placeholder="硅钢片型号"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        />
        <el-input-number
          v-model="listLspdQuery.field_eq_core_flux_density"
          placeholder="铁心磁密（T）"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLspdQuery.field_eq_lspd_creator"
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
          @click="getLspdList"
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
      <lspd-create
        v-model="showCreate"
        @get-lspd-list="getLspdList"
      />
      <lspd-view
        v-model="showDetail"
        :lspd="lspdDetail"
      />
      <lspd-update
        v-model="showUpdate"
        :lspd="lspdDetail"
        @get-lspd-list="getLspdList"
      />
      <el-table
        v-loading="loading"
        :data="lspdList"
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
          label="硅钢片型号"
          prop="stack_type"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_type }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="铁心磁密（T）"
          prop="core_flux_density"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.core_flux_density }}T</span>
          </template>
        </el-table-column>
        <el-table-column
          label="单位铁损（W/kg）"
          prop="unit_iron_loss"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.unit_iron_loss }}W/kg</span>
          </template>
        </el-table-column>
        <el-table-column
          label="单位质量磁化容量（VA/kg）"
          prop="unit_mass_magnet_capacity"
          sortable="custom"
          align="center"
          min-width="250px"
        >
          <template #default="{row}">
            <span>{{ row.unit_mass_magnet_capacity }}VA/kg</span>
          </template>
        </el-table-column>
        <el-table-column
          label="单位面积接缝伏安值（VA/cm2）"
          prop="unit_area_seam_va"
          sortable="custom"
          align="center"
          min-width="260px"
        >
          <template #default="{row}">
            <span>{{ row.unit_area_seam_va }}VA/cm2</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lspd_creator"
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
            <el-button
              size="mini"
              type="primary"
              @click="showLspdDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLspdUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLspd(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLspdQuery.page"
          v-model:page-size="listLspdQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLspdList"
          @current-change="getLspdList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLspdLibraries, GetStdLspdLibraries, ListLspdQuery, Lspd, ExportStdlspdLibraries } from '@/api/standard_libraries/lspd'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LspdCreate from './lspdCreate.vue'
import LspdView from './lspdView.vue'
import LspdUpdate from './lspdUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LspdCreate,
    LspdView,
    LspdUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLspdQuery = reactive({
      field_eq_stack_type: null,
      field_eq_core_flux_density: null,
      field_eq_unit_iron_loss: null,
      field_eq_unit_mass_magnet_capacity: null,
      field_eq_unit_area_seam_va: null,
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
    const lspdList = ref(Array<Lspd>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlspdLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '硅钢片性能数据.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '硅钢片性能数据.xlsx'
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
    async function getLspdList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLspdQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLspdQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLspdQuery.field_gt_created_at = 0
        listLspdQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLspdQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLspdQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLspdQuery.field_gt_updated_at = 0
        listLspdQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLspdLibraries(listLspdQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lspdList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLspdQuery.order_by = column.prop
      listLspdQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLspdList()
    }
    onMounted(() => {
      getLspdList()
    })
    function deleteLspd(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLspdLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLspdList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lspd: Lspd = {
      id: 0,
      stack_type: '',
      core_flux_density: 0,
      unit_iron_loss: 0,
      unit_mass_magnet_capacity: 0,
      unit_area_seam_va: 0
    }
    const lspdDetail = ref(lspd)
    function showLspdDetail(lspd: Lspd) {
      showDetail.value = true
      lspdDetail.value = lspd
    }
    function showLspdUpdate(lspd: Lspd) {
      showUpdate.value = true
      lspdDetail.value = lspd
    }
    return {
      status,
      handleDownload,
      loading,
      lspdList,
      LspdCreate,
      created_at_range,
      updated_at_range,
      listLspdQuery,
      total,
      getLspdList,
      sortChange,
      deleteLspd,
      showCreate,
      showDetail,
      showUpdate,
      lspdDetail,
      showLspdUpdate,
      showLspdDetail,
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
