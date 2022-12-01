<template>
  <div>
    <aside class="page-header">
      迭片系数
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input-number
          v-model="listLsfQuery.field_eq_stack_thick"
          placeholder="迭片片厚（mm）"
          clearable
          style="width: 220px; margin-right: 10px"
          size="small"
        />
        <el-input-number
          v-model="listLsfQuery.field_eq_stack_factor"
          placeholder="迭片系数"
          clearable
          style="width: 220px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLsfQuery.field_eq_lsf_creator"
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
          @click="getLsfList"
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
      <lsf-create
        v-model="showCreate"
        @get-lsf-list="getLsfList"
      />
      <lsf-view
        v-model="showDetail"
        :lsf="lsfDetail"
      />
      <lsf-update
        v-model="showUpdate"
        :lsf="lsfDetail"
        @get-lsf-list="getLsfList"
      />
      <el-table
        v-loading="loading"
        :data="lsfList"
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
          label="迭片片厚（mm）"
          prop="stack_thick"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.stack_thick }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="迭片系数"
          prop="stack_factor"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.stack_factor }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lsf_creator"
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
          width="250px"
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
          width="250px"
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
  <!--            @click="showLsfDetail(row)"-->
  <!--          >-->
  <!--            查看详情-->
  <!--          </el-button>-->
            <el-button
              size="mini"
              type="success"
              @click="showLsfUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLsf(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLsfQuery.page"
          v-model:page-size="listLsfQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLsfList"
          @current-change="getLsfList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLsfLibraries, GetStdLsfLibraries, ListLsfQuery, Lsf, ExportStdlsfLibraries } from '@/api/standard_libraries/lsf'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LsfCreate from './lsfCreate.vue'
import LsfView from './lsfView.vue'
import LsfUpdate from './lsfUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LsfCreate,
    LsfView,
    LsfUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLsfQuery = reactive({
      field_eq_stack_thick: null,
      field_eq_stack_factor: null,
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
    const lsfList = ref(Array<Lsf>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlsfLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '铁芯迭片系数.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '铁芯迭片系数.xlsx'
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
    async function getLsfList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLsfQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLsfQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLsfQuery.field_gt_created_at = 0
        listLsfQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLsfQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLsfQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLsfQuery.field_gt_updated_at = 0
        listLsfQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLsfLibraries(listLsfQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lsfList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLsfQuery.order_by = column.prop
      listLsfQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLsfList()
    }
    onMounted(() => {
      getLsfList()
    })
    function deleteLsf(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLsfLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLsfList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lsf: Lsf = {
      id: 0,
      stack_thick: 0,
      stack_factor: 0
    }
    const lsfDetail = ref(lsf)
    function showLsfDetail(lsf: Lsf) {
      showDetail.value = true
      lsfDetail.value = lsf
    }
    function showLsfUpdate(lsf: Lsf) {
      showUpdate.value = true
      lsfDetail.value = lsf
    }
    return {
      status,
      handleDownload,
      loading,
      lsfList,
      LsfCreate,
      created_at_range,
      updated_at_range,
      listLsfQuery,
      total,
      getLsfList,
      sortChange,
      deleteLsf,
      showCreate,
      showDetail,
      showUpdate,
      lsfDetail,
      showLsfUpdate,
      showLsfDetail,
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
