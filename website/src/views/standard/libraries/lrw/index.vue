<template>
  <div>
    <aside class="page-header">
      调压方式
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLrwQuery.field_eq_regulate_way"
          filterable
          allow-create
          placeholder="调压方式"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="无励磁调压"
            value="无励磁调压"
          />
          <el-option
            label="有载调压"
            value="有载调压"
          />
        </el-select>
        <el-select
          v-model="listLrwQuery.field_eq_regulate_way_sign"
          placeholder="调压方式代号"
          filterable
          allow-create
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        >
          <el-option
            label=""
            value=""
          />
          <el-option
            label="Z"
            value="Z"
          />
        </el-select>
        <UserSearchInput
          v-model="listLrwQuery.field_eq_lrw_creator"
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
          @click="getLrwList"
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
      <lrw-create
        v-model="showCreate"
        @get-lrw-list="getLrwList"
      />
      <lrw-view
        v-model="showDetail"
        :lrw="lrwDetail"
      />
      <lrw-update
        v-model="showUpdate"
        :lrw="lrwDetail"
        @get-lrw-list="getLrwList"
      />
      <el-table
        v-loading="loading"
        :data="lrwList"
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
          label="调压方式"
          prop="regulate_way"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.regulate_way }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="调压方式代号"
          prop="regulate_way_sign"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.regulate_way_sign }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lrw_creator"
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
              type="success"
              @click="showLrwUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLrw(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLrwQuery.page"
          v-model:page-size="listLrwQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLrwList"
          @current-change="getLrwList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLrwLibraries, GetStdLrwLibraries, ListLrwQuery, Lrw, ExportStdlrwLibraries } from '@/api/standard_libraries/lrw'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LrwCreate from './lrwCreate.vue'
import LrwView from './lrwView.vue'
import LrwUpdate from './lrwUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LrwCreate,
    LrwView,
    LrwUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLrwQuery = reactive({
      field_eq_regulate_way: null,
      field_eq_regulate_way_sign: null,
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
    const lrwList = ref(Array<Lrw>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlrwLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '调压方式.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '调压方式.xlsx'
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
    async function getLrwList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLrwQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLrwQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLrwQuery.field_gt_created_at = 0
        listLrwQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLrwQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLrwQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLrwQuery.field_gt_updated_at = 0
        listLrwQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLrwLibraries(listLrwQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lrwList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLrwQuery.order_by = column.prop
      listLrwQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLrwList()
    }
    onMounted(() => {
      getLrwList()
    })
    function deleteLrw(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLrwLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLrwList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lrw: Lrw = {
      id: 0,
      regulate_way: '',
      regulate_way_sign: ''
    }
    const lrwDetail = ref(lrw)
    function showLrwDetail(lrw: Lrw) {
      showDetail.value = true
      lrwDetail.value = lrw
    }
    function showLrwUpdate(lrw: Lrw) {
      showUpdate.value = true
      lrwDetail.value = lrw
    }
    return {
      status,
      handleDownload,
      loading,
      lrwList,
      LrwCreate,
      created_at_range,
      updated_at_range,
      listLrwQuery,
      total,
      getLrwList,
      sortChange,
      deleteLrw,
      showCreate,
      showDetail,
      showUpdate,
      lrwDetail,
      showLrwUpdate,
      showLrwDetail,
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
