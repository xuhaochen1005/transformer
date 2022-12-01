<template>
  <div>
    <aside class="page-header">
      线圈接法
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLccQuery.field_eq_low_vol_coil_connect"
          filterable
          allow-create
          clearable
          style="width: 200px; margin-right: 10px"
          placeholder="低压线圈接法"
          size="small"
        >
          <el-option
            label="y"
            value="y"
          />
          <el-option
            label="d"
            value="d"
          />
        </el-select>
        <el-select
          v-model="listLccQuery.field_eq_high_vol_coil_connect"
          filterable
          allow-create
          clearable
          style="width: 200px; margin-right: 10px"
          placeholder="高压线圈接法"
          size="small"
        >
          <el-option
            label="Y"
            value="Y"
          />
          <el-option
            label="D"
            value="D"
          />
        </el-select>
        <UserSearchInput
          v-model="listLccQuery.field_eq_lcc_creator"
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
          @click="getLccList"
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
      <lcc-create
        v-model="showCreate"
        @get-lcc-list="getLccList"
      />
      <lcc-view
        v-model="showDetail"
        :lcc="lccDetail"
      />
      <lcc-update
        v-model="showUpdate"
        :lcc="lccDetail"
        @get-lcc-list="getLccList"
      />
      <el-table
        v-loading="loading"
        :data="lccList"
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
        >
          <template #default="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="低压线圈接法"
          prop="low_vol_coil_connect"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.low_vol_coil_connect }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="高压线圈接法"
          prop="high_vol_coil_connect"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.high_vol_coil_connect }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lcc_creator"
          sortable="custom"
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
        >
          <template #default="{row}">
            <span>{{ UnixTime2HumanWithStr(row.updated_at) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          align="center"
          fixed="right"
          min-width="80px"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="success"
              @click="showLccUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLcc(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLccQuery.page"
          v-model:page-size="listLccQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLccList"
          @current-change="getLccList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlccLibraries, ExportStdlccLibraries, GetStdlccLibraries, Lcc } from '@/api/standard_libraries/lcc'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LccCreate from './lccCreate.vue'
import LccView from './lccView.vue'
import LccUpdate from './lccUpdate.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import dayjs from 'dayjs'

export default defineComponent({
  components: {
    UserSearchInput,
    LccCreate,
    LccView,
    LccUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLccQuery = reactive({
      field_eq_low_vol_coil_connect: '',
      field_eq_high_vol_coil_connect: '',
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
    const lccList = ref(Array<Lcc>())
    const total = ref(0)

    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }

    async function handleDownload() {
      const res = await ExportStdlccLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '线圈接法.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '线圈接法.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }

    async function getLccList() {
      loading.value = true
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLccQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLccQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLccQuery.field_gt_created_at = 0
        listLccQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLccQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLccQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLccQuery.field_gt_updated_at = 0
        listLccQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlccLibraries(listLccQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lccList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }

    function sortChange(column: any) {
      listLccQuery.order_by = column.prop
      listLccQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLccList()
    }

    onMounted(() => {
      getLccList()
    })

    function deleteLcc(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlccLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLccList()
            }
          }
        }
      })
    }

    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lcc: Lcc = {
      id: 0,
      low_vol_coil_connect: '',
      high_vol_coil_connect: ''
    }
    const lccDetail = ref(lcc)

    function showLccDetail(lcc: Lcc) {
      showDetail.value = true
      lccDetail.value = lcc
    }

    function showLccUpdate(lcc: Lcc) {
      showUpdate.value = true
      lccDetail.value = lcc
    }

    return {
      handleDownload,
      status,
      loading,
      lccList,
      LccCreate,
      created_at_range,
      updated_at_range,
      listLccQuery,
      total,
      getLccList,
      sortChange,
      deleteLcc,
      showCreate,
      showDetail,
      showUpdate,
      lccDetail,
      showLccUpdate,
      showLccDetail,
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
