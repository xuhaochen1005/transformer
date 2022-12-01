<template>
  <div>
    <aside class="page-header">
      高压外模台账
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLhvomlQuery.field_eq_mod_type"
          filterable
          allow-create
          placeholder="模具类型"
          clearable
          style="width: 110px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="硬模"
            value="硬模"
          />
          <el-option
            label="软模"
            value="软模"
          />
          <el-option
            label="双凸台软模"
            value="双凸台软模"
          />
          <el-option
            label="软模(矩形)"
            value="软模(矩形)"
          />
          <el-option
            label="软模(长圆)"
            value="软模(长圆)"
          />
        </el-select>
        <el-input
          v-model="listLhvomlQuery.field_eq_mod_size"
          placeholder="模具尺寸(mm)"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLhvomlQuery.field_eq_mod_height"
          placeholder="模具高度(mm)"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLhvomlQuery.field_eq_mod_num"
          placeholder="模具编号"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLhvomlQuery.field_eq_remark"
          placeholder="备注"
          clearable
          style="width: 110px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLhvomlQuery.field_eq_lhvoml_creator"
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
          @click="getLhvomlList"
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
      <Lhvoml-create
        v-model="showCreate"
        @get-Lhvoml-list="getLhvomlList"
      />
      <Lhvoml-view
        v-model="showDetail"
        :lhvoml="LhvomlDetail"
      />
      <Lhvoml-update
        v-model="showUpdate"
        :lhvoml="LhvomlDetail"
        @get-Lhvoml-list="getLhvomlList"
      />
      <el-table
        v-loading="loading"
        :data="LhvomlList"
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
          min-width="80px"
        >
          <template #default="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具尺寸"
          prop="mod_size"
          sortable="custom"
          align="center"
          min-width="110px"
        >
          <template #default="{row}">
            <span>{{ row.mod_size }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具类型"
          prop="mod_type"
          sortable="custom"
          align="center"
          min-width="110px"
        >
          <template #default="{row}">
            <span>{{ row.mod_type }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具高度(mm)"
          prop="mod_height"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.mod_height }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具数量(个)"
          prop="mod_amount"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.mod_amount }}个</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具图号"
          prop="mod_pic"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.mod_pic }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具编号"
          prop="mod_num"
          sortable="custom"
          align="center"
          min-width="130x"
        >
          <template #default="{row}">
            <span>{{ row.mod_num }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="凸台高度(封板尺寸)"
          prop="boss_width"
          sortable="custom"
          align="center"
          min-width="180x"
        >
          <template #default="{row}">
            <span>{{ row.boss_width }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="是否开槽"
          prop="groove"
          sortable="custom"
          align="center"
          min-width="130x"
        >
          <template #default="{row}">
            <span>{{ row.groove }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="螺母尺寸(mm)"
          prop="nut_size"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.nut_size }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="A头尺寸(mm)"
          prop="a_size"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.a_size }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="抽头孔距(mm)"
          prop="tap_hole_distance"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.tap_hole_distance }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="封板图号"
          prop="closure_pic"
          sortable="custom"
          align="center"
          min-width="130px"
        >
          <template #default="{row}">
            <span>{{ row.closure_pic }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="备注"
          prop="remark"
          sortable="custom"
          align="center"
          min-width="130px"
        >
          <template #default="{row}">
            <span>{{ row.remark }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lhvoml_creator"
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
          min-width="250px"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="primary"
              @click="showLhvomlDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLhvomlUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLhvoml(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLhvomlQuery.page"
          v-model:page-size="listLhvomlQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLhvomlList"
          @current-change="getLhvomlList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLhvomlLibraries, GetStdLhvomlLibraries, ListLhvomlQuery, Lhvoml, ExportStdlhvomlLibraries } from '@/api/standard_libraries/lhvoml'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LhvomlCreate from './lhvomlCreate.vue'
import LhvomlView from './lhvomlView.vue'
import LhvomlUpdate from './lhvomlUpdate.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import dayjs from 'dayjs'

export default defineComponent({
  components: {
    UserSearchInput,
    LhvomlCreate,
    LhvomlView,
    LhvomlUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLhvomlQuery = reactive({
      id: 0,
      field_eq_mod_size: '',
      field_eq_mod_type: '',
      field_eq_mod_height: null,
      field_eq_mod_amount: null,
      field_eq_mod_pic: '',
      field_eq_mod_num: '',
      field_eq_mod_suit: '',
      field_eq_boss_width: '',
      field_eq_groove: '',
      field_eq_nut_size: null,
      field_eq_a_size: null,
      field_eq_tap_hole_distance: null,
      field_eq_closure_pic: '',
      field_eq_remark: '',
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
    const LhvomlList = ref(Array<Lhvoml>())
    const total = ref(0)
    const downloading = ref(false)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      downloading.value = true
      const res = await ExportStdlhvomlLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '高压外模台账.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '高压外模台账.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      downloading.value = false
    }
    async function getLhvomlList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLhvomlQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLhvomlQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLhvomlQuery.field_gt_created_at = 0
        listLhvomlQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLhvomlQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLhvomlQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLhvomlQuery.field_gt_updated_at = 0
        listLhvomlQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLhvomlLibraries(listLhvomlQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        LhvomlList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLhvomlQuery.order_by = column.prop
      listLhvomlQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLhvomlList()
    }
    onMounted(() => {
      getLhvomlList()
    })
    function deleteLhvoml(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLhvomlLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLhvomlList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lhvoml: Lhvoml = {
      id: 0,
      mod_size: '',
      mod_type: '',
      mod_height: null,
      mod_amount: null,
      mod_pic: '',
      mod_num: '',
      mod_suit: '',
      boss_width: '',
      groove: '',
      nut_size: 0,
      a_size: null,
      tap_hole_distance: null,
      closure_pic: '',
      remark: ''
    }
    const LhvomlDetail = ref(lhvoml)
    function showLhvomlDetail(Lhvoml: Lhvoml) {
      showDetail.value = true
      LhvomlDetail.value = Lhvoml
    }
    function showLhvomlUpdate(Lhvoml: Lhvoml) {
      showUpdate.value = true
      LhvomlDetail.value = Lhvoml
    }
    return {
      status,
      handleDownload,
      loading,
      LhvomlList,
      LhvomlCreate,
      created_at_range,
      updated_at_range,
      listLhvomlQuery,
      total,
      getLhvomlList,
      sortChange,
      deleteLhvoml,
      showCreate,
      showDetail,
      showUpdate,
      LhvomlDetail,
      showLhvomlUpdate,
      showLhvomlDetail,
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
