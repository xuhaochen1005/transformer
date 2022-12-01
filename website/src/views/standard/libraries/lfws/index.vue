<template>
  <div>
    <aside class="page-header">
      扁导线规格
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLfwsQuery.field_lk_std_length"
          placeholder="标称长(mm)"
          clearable
          style="width: 150px;; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLfwsQuery.field_lk_std_width"
          placeholder="标称宽(mm)"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLfwsQuery.field_lk_section_area"
          placeholder="截面积(mm^2)"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLfwsQuery.field_lk_round_corner"
          placeholder="圆角半径(mm)"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLfwsQuery.field_lk_remark"
          placeholder="备注"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLfwsQuery.field_eq_lfws_creator"
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
          @click="getLfwsList"
        >
          搜索
        </el-button>
        <el-button
          v-waves
          type="primary"
          icon="el-icon-edit"
          style="margin-top: 10px"
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
      <lfws-create
        v-model="showCreate"
        @get-lfws-list="getLfwsList"
      />
      <lfws-view
        v-model="showDetail"
        :lfws="lfwsDetail"
      />
      <lfws-update
        v-model="showUpdate"
        :lfws="lfwsDetail"
        @get-lfws-list="getLfwsList"
      />
      <el-table
        v-loading="loading"
        :data="lfwsList"
        border
        fit
        highlight-current-row
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
          label="标称长(mm)"
          prop="std_length"
          sortable="custom"
          align="center"
          min-width="140px"
        >
          <template #default="{row}">
            <span>{{ row.std_length }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="标称宽(mm)"
          prop="std_width"
          sortable="custom"
          align="center"
          min-width="140px"
        >
          <template #default="{row}">
            <span>{{ row.std_width }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="截面积(mm^2)"
          prop="section_area"
          sortable="custom"
          align="center"
          min-width="160px"
        >
          <template #default="{row}">
            <span>{{ row.section_area }}mm^2</span>
          </template>
        </el-table-column>
        <el-table-column
          label="圆角半径(mm)"
          prop="round_corner"
          sortable="custom"
          align="center"
          min-width="160px"
        >
          <template #default="{row}">
            <span>{{ row.round_corner }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="备注"
          prop="remark"
          sortable="custom"
          align="center"
          min-width="140px"
        >
          <template #default="{row}">
            <span>{{ row.remark }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lfws_creator"
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
              @click="showLfwsDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLfwsUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLfws(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLfwsQuery.page"
          v-model:page-size="listLfwsQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLfwsList"
          @current-change="getLfwsList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlfwsLibraries, GetStdlfwsLibraries, ListLfwsQuery, Lfws, ExportStdlfwsLibraries } from '@/api/standard_libraries/lfws'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LfwsCreate from './lfwsCreate.vue'
import LfwsView from './lfwsView.vue'
import LfwsUpdate from './lfwsUpdate.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import dayjs from 'dayjs'

export default defineComponent({
  components: {
    UserSearchInput,
    LfwsCreate,
    LfwsView,
    LfwsUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLfwsQuery = reactive({
      field_lk_std_length: null,
      field_lk_std_width: null,
      field_lk_section_area: null,
      field_lk_round_corner: null,
      field_lk_remark: '',
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
    const lfwsList = ref(Array<Lfws>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlfwsLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '扁铜（铝）线规格.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '扁铜（铝）线规格.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLfwsList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLfwsQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLfwsQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLfwsQuery.field_gt_created_at = 0
        listLfwsQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLfwsQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLfwsQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLfwsQuery.field_gt_updated_at = 0
        listLfwsQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlfwsLibraries(listLfwsQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lfwsList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLfwsQuery.order_by = column.prop
      listLfwsQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLfwsList()
    }
    onMounted(() => {
      getLfwsList()
    })
    function deleteLfws(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlfwsLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLfwsList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lfws: Lfws = {
      id: 0,
      std_length: null,
      std_width: null,
      section_area: null,
      round_corner: null,
      remark: ''
    }
    const lfwsDetail = ref(lfws)
    function showLfwsDetail(lfws: Lfws) {
      showDetail.value = true
      lfwsDetail.value = lfws
    }
    function showLfwsUpdate(lfws: Lfws) {
      showUpdate.value = true
      lfwsDetail.value = lfws
    }
    return {
      handleDownload,
      status,
      loading,
      lfwsList,
      LfwsCreate,
      created_at_range,
      updated_at_range,
      listLfwsQuery,
      total,
      getLfwsList,
      sortChange,
      deleteLfws,
      showCreate,
      showDetail,
      showUpdate,
      lfwsDetail,
      showLfwsUpdate,
      showLfwsDetail,
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
