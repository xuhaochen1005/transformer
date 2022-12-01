<template>
  <div>
    <aside class="page-header">
      涡流损耗系数
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLelfQuery.field_lk_rated_capacity"
          placeholder="额定容量(kVA)"
          clearable
          style="width: 200px;; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLelfQuery.field_lk_eddy_loss_factor"
          placeholder="涡流损耗系数(%)"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLelfQuery.field_eq_lelf_creator"
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
          @click="getLelfList"
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
      <lelf-create
        v-model="showCreate"
        @get-lelf-list="getLelfList"
      />
      <lelf-view
        v-model="showDetail"
        :lelf="lelfDetail"
      />
      <lelf-update
        v-model="showUpdate"
        :lelf="lelfDetail"
        @get-lelf-list="getLelfList"
      />
      <el-table
        v-loading="loading"
        :data="lelfList"
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
          label="额定容量(kVA)"
          prop="rated_capacity"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.rated_capacity }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="涡流损耗系数(%)"
          prop="eddy_loss_factor"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.eddy_loss_factor }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lelf_creator"
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
              @click="showLelfUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLelf(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLelfQuery.page"
          v-model:page-size="listLelfQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLelfList"
          @current-change="getLelfList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlelfLibraries, GetStdlelfLibraries, ListLelfQuery, Lelf, ExportStdlelfLibraries } from '@/api/standard_libraries/lelf'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LelfCreate from './lelfCreate.vue'
import LelfView from './lelfView.vue'
import LelfUpdate from './lelfUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LelfCreate,
    LelfView,
    LelfUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLelfQuery = reactive({
      field_lk_rated_capacity: null,
      field_lk_eddy_loss_factor: null,
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
    const lelfList = ref(Array<Lelf>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlelfLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '涡流损耗系数.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '涡流损耗系数.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLelfList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLelfQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLelfQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLelfQuery.field_gt_created_at = 0
        listLelfQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLelfQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLelfQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLelfQuery.field_gt_updated_at = 0
        listLelfQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlelfLibraries(listLelfQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lelfList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLelfQuery.order_by = column.prop
      listLelfQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLelfList()
    }
    onMounted(() => {
      getLelfList()
    })
    function deleteLelf(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlelfLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLelfList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lelf: Lelf = {
      id: 0,
      rated_capacity: null,
      eddy_loss_factor: null
    }
    const lelfDetail = ref(lelf)
    function showLelfDetail(lelf: Lelf) {
      showDetail.value = true
      lelfDetail.value = lelf
    }
    function showLelfUpdate(lelf: Lelf) {
      showUpdate.value = true
      lelfDetail.value = lelf
    }
    return {
      handleDownload,
      status,
      loading,
      lelfList,
      LelfCreate,
      created_at_range,
      updated_at_range,
      listLelfQuery,
      total,
      getLelfList,
      sortChange,
      deleteLelf,
      showCreate,
      showDetail,
      showUpdate,
      lelfDetail,
      showLelfUpdate,
      showLelfDetail,
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
