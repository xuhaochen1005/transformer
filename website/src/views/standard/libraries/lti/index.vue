<template>
  <div>
    <aside class="page-header">
      端绝缘距离
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
          v-model="lowVol"
          placeholder="额定低压（kV）"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        />
        <el-select
          v-model="listLtiQuery.field_eq_terminal_insulate"
          placeholder="端绝缘距离（mm）"
          clearable
          style="width: 170px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="15mm"
            value="15"
          />
          <el-option
            label="20mm"
            value="20"
          />
          <el-option
            label="25mm"
            value="25"
          />
          <el-option
            label="30mm"
            value="30"
          />
          <el-option
            label="35mm"
            value="35"
          />
          <el-option
            label="40mm"
            value="40"
          />
        </el-select>
        <UserSearchInput
          v-model="listLtiQuery.field_eq_lti_creator"
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
          @click="getLtiList"
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
      <Lti-create
        v-model="showCreate"
        @get-Lti-list="getLtiList"
      />
      <Lti-view
        v-model="showDetail"
        :lti="LtiDetail"
      />
      <Lti-update
        v-model="showUpdate"
        :lti="LtiDetail"
        @get-Lti-list="getLtiList"
      />
      <el-table
        v-loading="loading"
        :data="LtiList"
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
          label="额定高压下界（kV）"
          prop="rated_high_vol_min"
          sortable="custom"
          align="center"
          min-width="180px"
        >
          <template #default="{row}">
            <span>{{ row.rated_high_vol_min }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定高压上界（kV）"
          prop="rated_high_vol_max"
          sortable="custom"
          align="center"
          min-width="180px"
        >
          <template #default="{row}">
            <span>{{ row.rated_low_vol_max }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定低压下界（kV）"
          prop="rated_low_vol_min"
          sortable="custom"
          align="center"
          min-width="180px"
        >
          <template #default="{row}">
            <span>{{ row.rated_low_vol_min }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定低压上界（kV）"
          prop="rated_low_vol_max"
          sortable="custom"
          align="center"
          min-width="180px"
        >
          <template #default="{row}">
            <span>{{ row.rated_low_vol_max }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="端绝缘距离（mm）"
          prop="terminal_insulate"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.terminal_insulate }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lti_creator"
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
              @click="showLtiDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLtiUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLti(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLtiQuery.page"
          v-model:page-size="listLtiQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLtiList"
          @current-change="getLtiList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLtiLibraries, GetStdLtiLibraries, ListLtiQuery, Lti, ExportStdltiLibraries } from '@/api/standard_libraries/lti'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LtiCreate from './ltiCreate.vue'
import LtiView from './ltiView.vue'
import LtiUpdate from './ltiUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LtiCreate,
    LtiView,
    LtiUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const highVol: Ref<number | undefined> = ref(undefined)
    const lowVol: Ref<number | undefined> = ref(undefined)
    const listLtiQuery = reactive({
      // field_eq_rated_high_vol_min: null,
      // field_eq_rated_high_vol_max: null,
      // field_eq_rated_low_vol_min: null,
      // field_eq_rated_low_vol_max: null,
      field_lt_rated_high_vol_min: highVol,
      field_ge_rated_high_vol_max: highVol,
      field_lt_rated_low_vol_min: lowVol,
      field_ge_rated_low_vol_max: lowVol,
      field_eq_terminal_insulate: null,
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
    const LtiList = ref(Array<Lti>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdltiLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '端绝缘距离.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '端绝缘距离.xlsx'
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
    async function getLtiList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLtiQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLtiQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLtiQuery.field_gt_created_at = 0
        listLtiQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLtiQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLtiQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLtiQuery.field_gt_updated_at = 0
        listLtiQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLtiLibraries(listLtiQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        LtiList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLtiQuery.order_by = column.prop
      listLtiQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLtiList()
    }
    onMounted(() => {
      getLtiList()
    })
    function deleteLti(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLtiLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLtiList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lti: Lti = {
      id: 0,
      rated_high_vol_min: 0,
      rated_high_vol_max: 0,
      rated_low_vol_min: 0,
      rated_low_vol_max: 0,
      terminal_insulate: 0,
      created_at: '',
      updated_at: ''
    }
    const LtiDetail = ref(lti)
    function showLtiDetail(Lti: Lti) {
      showDetail.value = true
      LtiDetail.value = Lti
    }
    function showLtiUpdate(Lti: Lti) {
      showUpdate.value = true
      LtiDetail.value = Lti
    }
    return {
      status,
      handleDownload,
      loading,
      LtiList,
      LtiCreate,
      created_at_range,
      updated_at_range,
      listLtiQuery,
      total,
      highVol,
      lowVol,
      getLtiList,
      sortChange,
      deleteLti,
      showCreate,
      showDetail,
      showUpdate,
      LtiDetail,
      showLtiUpdate,
      showLtiDetail,
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
