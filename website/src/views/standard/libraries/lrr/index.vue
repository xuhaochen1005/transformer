<template>
  <div>
    <aside class="page-header">
      调压范围
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLrrQuery.field_eq_regulate_range_min"
          filterable
          allow-create
          placeholder="调压范围下界（%）"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="-2.5%"
            value="-2.5%"
          />
          <el-option
            label="-5%"
            value="-5%"
          />
          <el-option
            label="-10%"
            value="-10%"
          />
        </el-select>
        <el-select
          v-model="listLrrQuery.field_eq_regulate_range_max"
          filterable
          allow-create
          placeholder="调压范围上界（%）"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="2.5%"
            value="2.5%"
          />
          <el-option
            label="5%"
            value="5%"
          />
          <el-option
            label="10%"
            value="10%"
          />
        </el-select>
        <el-select
          v-model="listLrrQuery.field_eq_regulate_range_step"
          filterable
          allow-create
          placeholder="调压范围步长（%）"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="2.5%"
            value="2.5%"
          />
          <el-option
            label="5%"
            value="5%"
          />
        </el-select>
        <UserSearchInput
          v-model="listLrrQuery.field_eq_lrr_creator"
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
          @click="getLrrList"
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
      <Lrr-create
        v-model="showCreate"
        @get-Lrr-list="getLrrList"
      />
      <Lrr-view
        v-model="showDetail"
        :lrr="LrrDetail"
      />
      <Lrr-update
        v-model="showUpdate"
        :lrr="LrrDetail"
        @get-Llvoml-list="getLrrList"
      />
      <el-table
        v-loading="loading"
        :data="LrrList"
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
          label="调压范围下界（%）"
          prop="regulate_range_min"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.regulate_range_min }}%</span>
          </template>
        </el-table-column>
        <el-table-column
          label="调压范围上界（%）"
          prop="regulate_range_max"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.regulate_range_max }}%</span>
          </template>
        </el-table-column>
        <el-table-column
          label="调压范围步长（%）"
          prop="regulate_range_step"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.regulate_range_step }}%</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lrr_creator"
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
          width="250px"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="success"
              @click="showLrrUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLrr(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLrrQuery.page"
          v-model:page-size="listLrrQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLrrList"
          @current-change="getLrrList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLrrLibraries, GetStdLrrLibraries, ListLrrQuery, Lrr, ExportStdlrrLibraries } from '@/api/standard_libraries/lrr'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LrrCreate from './lrrCreate.vue'
import LrrView from './lrrView.vue'
import LrrUpdate from './lrrUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LrrCreate,
    LrrView,
    LrrUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLrrQuery = reactive({
      field_eq_regulate_range_min: null,
      field_eq_regulate_range_max: null,
      field_eq_regulate_range_step: null,
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
    const LrrList = ref(Array<Lrr>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlrrLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '调压范围.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '调压范围.xlsx'
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
    async function getLrrList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLrrQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLrrQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLrrQuery.field_gt_created_at = 0
        listLrrQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLrrQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLrrQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLrrQuery.field_gt_updated_at = 0
        listLrrQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLrrLibraries(listLrrQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        LrrList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLrrQuery.order_by = column.prop
      listLrrQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLrrList()
    }
    onMounted(() => {
      getLrrList()
    })
    function deleteLrr(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLrrLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLrrList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lrr: Lrr = {
      id: 0,
      regulate_range_min: 0,
      regulate_range_max: 0,
      regulate_range_step: 0,
      created_at: '',
      updated_at: '',
      deleted_at: ''
    }
    const LrrDetail = ref(lrr)
    function showLrrDetail(Lrr: Lrr) {
      showDetail.value = true
      LrrDetail.value = Lrr
    }
    function showLrrUpdate(Lrr: Lrr) {
      showUpdate.value = true
      LrrDetail.value = Lrr
    }
    return {
      status,
      handleDownload,
      loading,
      LrrList,
      LrrCreate,
      created_at_range,
      updated_at_range,
      listLrrQuery,
      total,
      getLrrList,
      sortChange,
      deleteLrr,
      showCreate,
      showDetail,
      showUpdate,
      LrrDetail,
      showLrrUpdate,
      showLrrDetail,
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
