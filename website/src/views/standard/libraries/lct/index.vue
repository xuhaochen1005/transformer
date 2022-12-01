<template>
  <div>
    <aside class="page-header">
      冷却方式
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLctQuery.field_eq_cool_type"
          filterable
          allow-create
          clearable
          style="width: 200px; margin-right: 10px"
          placeholder="冷却方式"
          size="small"
        >
          <el-option
            lable="AN"
            value="AN"
          />
          <el-option
            lable="AF"
            value="AF"
          />
        </el-select>
        <UserSearchInput
          v-model="listLctQuery.field_eq_lct_creator"
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
          @click="getLctList"
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
      <lct-create
        v-model="showCreate"
        @get-lct-list="getLctList"
      />
      <lct-view
        v-model="showDetail"
        :lct="lctDetail"
      />
      <lct-update
        v-model="showUpdate"
        :lct="lctDetail"
        @get-lct-list="getLctList"
      />
      <el-table
        v-loading="loading"
        :data="lctList"
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
          label="冷却方式"
          prop="cool_type"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.cool_type }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lct_creator"
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
              @click="showLctUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLct(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLctQuery.page"
          v-model:page-size="listLctQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLctList"
          @current-change="getLctList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlctLibraries, GetStdlctLibraries, ListLctQuery, Lct, ExportStdlctLibraries } from '@/api/standard_libraries/lct'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LctCreate from './lctCreate.vue'
import LctView from './lctView.vue'
import LctUpdate from './lctUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LctCreate,
    LctView,
    LctUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLctQuery = reactive({
      field_eq_cool_type: '',
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
    const lctList = ref(Array<Lct>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlctLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '冷却方式.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '冷却方式.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLctList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLctQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLctQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLctQuery.field_gt_created_at = 0
        listLctQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLctQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLctQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLctQuery.field_gt_updated_at = 0
        listLctQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlctLibraries(listLctQuery)
      if (response.data.code === 200) {
        // console.log(response.data)
        total.value = response.data.total
        lctList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLctQuery.order_by = column.prop
      listLctQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLctList()
    }
    onMounted(() => {
      getLctList()
    })
    function deleteLct(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlctLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLctList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lct: Lct = {
      id: 0,
      cool_type: ''
    }
    const lctDetail = ref(lct)
    function showLctDetail(lct: Lct) {
      showDetail.value = true
      lctDetail.value = lct
    }
    function showLctUpdate(lct: Lct) {
      showUpdate.value = true
      lctDetail.value = lct
    }
    return {
      handleDownload,
      status,
      loading,
      lctList,
      LctCreate,
      created_at_range,
      updated_at_range,
      listLctQuery,
      total,
      getLctList,
      sortChange,
      deleteLct,
      showCreate,
      showDetail,
      showUpdate,
      lctDetail,
      showLctUpdate,
      showLctDetail,
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
