<template>
  <div>
    <aside class="page-header">
      温升限值
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLtrQuery.field_eq_temp"
          filterable
          allow-create
          placeholder="绝缘系统温度（℃）"
          clearable
          style="width: 180px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="105℃"
            value="105"
          />
          <el-option
            label="120℃"
            value="120"
          />
          <el-option
            label="130℃"
            value="130"
          />
          <el-option
            label="155℃"
            value="155"
          />
          <el-option
            label="180℃"
            value="180"
          />
          <el-option
            label="200℃"
            value="200"
          />
          <el-option
            label="220℃"
            value="220"
          />
        </el-select>
        <el-select
          v-model="listLtrQuery.field_eq_temp_sign"
          placeholder="绝缘温度代号"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="A"
            value="A"
          />
          <el-option
            label="E"
            value="E"
          />
          <el-option
            label="B"
            value="B"
          />
          <el-option
            label="F"
            value="F"
          />
          <el-option
            label="H"
            value="H"
          />
          <el-option
            label="/"
            value="/"
          />
        </el-select>
        <el-select
          v-model="listLtrQuery.field_eq_low_vol_temp_rise"
          placeholder="低压温升限值（K）"
          clearable
          style="width: 180px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="60K"
            value="60"
          />
          <el-option
            label="75k"
            value="75"
          />
          <el-option
            label="80k"
            value="80"
          />
          <el-option
            label="100K"
            value="100"
          />
          <el-option
            label="125k"
            value="125"
          />
          <el-option
            label="135k"
            value="135"
          />
          <el-option
            label="150k"
            value="150"
          />
        </el-select>
        <el-select
          v-model="listLtrQuery.field_eq_low_vol_temp_rise"
          placeholder="高压温升限值（K）"
          clearable
          style="width: 180px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="60K"
            value="60"
          />
          <el-option
            label="75k"
            value="75"
          />
          <el-option
            label="80k"
            value="80"
          />
          <el-option
            label="100K"
            value="100"
          />
          <el-option
            label="125k"
            value="125"
          />
          <el-option
            label="135k"
            value="135"
          />
          <el-option
            label="150k"
            value="150"
          />
        </el-select>
        <UserSearchInput
          v-model="listLtrQuery.field_eq_ltr_creator"
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
          @click="getLtrList"
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
      <ltr-create
        v-model="showCreate"
        @get-ltr-list="getLtrList"
      />
      <ltr-view
        v-model="showDetail"
        :ltr="ltrDetail"
      />
      <ltr-update
        v-model="showUpdate"
        :ltr="ltrDetail"
        @get-ltr-list="getLtrList"
      />
      <el-table
        v-loading="loading"
        :data="ltrList"
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
          label="绝缘系统温度（℃）"
          prop="temp"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.temp }}℃</span>
          </template>
        </el-table-column>
        <el-table-column
          label="绝缘温度代号"
          prop="temp_sign"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.temp_sign }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="低压温升限值（K）"
          prop="low_vol_temp_rise"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.low_vol_temp_rise }}k</span>
          </template>
        </el-table-column>
        <el-table-column
          label="高压温升限值（K）"
          prop="high_vol_temp_rise"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.high_vol_temp_rise }}k</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="ltr_creator"
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
              @click="showLtrDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLtrUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLtr(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLtrQuery.page"
          v-model:page-size="listLtrQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLtrList"
          @current-change="getLtrList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLtrLibraries, GetStdLtrLibraries, Ltr, ExportStdltrLibraries } from '@/api/standard_libraries/ltr'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LtrCreate from './ltrCreate.vue'
import LtrView from './ltrView.vue'
import LtrUpdate from './ltrUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LtrCreate,
    LtrView,
    LtrUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLtrQuery = reactive({
      field_eq_temp: null,
      field_eq_temp_sign: null,
      field_eq_low_vol_temp_rise: null,
      field_eq_high_vol_temp_rise: null,
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
    const ltrList = ref(Array<Ltr>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdltrLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '绕组温升.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '绕组温升.xlsx'
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

    async function getLtrList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLtrQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLtrQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLtrQuery.field_gt_created_at = 0
        listLtrQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLtrQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLtrQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLtrQuery.field_gt_updated_at = 0
        listLtrQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLtrLibraries(listLtrQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        ltrList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }

    function sortChange(column: any) {
      listLtrQuery.order_by = column.prop
      listLtrQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLtrList()
    }

    onMounted(() => {
      getLtrList()
    })

    function deleteLtr(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLtrLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLtrList()
            }
          }
        }
      })
    }

    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const ltr: Ltr = {
      id: 0,
      temp: 0,
      temp_sign: '',
      low_vol_temp_rise: 0,
      high_vol_temp_rise: 0
    }
    const ltrDetail = ref(ltr)

    function showLtrDetail(ltr: Ltr) {
      showDetail.value = true
      ltrDetail.value = ltr
    }

    function showLtrUpdate(ltr: Ltr) {
      showUpdate.value = true
      ltrDetail.value = ltr
    }

    return {
      status,
      handleDownload,
      loading,
      ltrList,
      LtrCreate,
      created_at_range,
      updated_at_range,
      listLtrQuery,
      total,
      getLtrList,
      sortChange,
      deleteLtr,
      showCreate,
      showDetail,
      showUpdate,
      ltrDetail,
      showLtrUpdate,
      showLtrDetail,
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
