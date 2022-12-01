<template>
  <div>
    <aside class="page-header">
      箔绕低压模具台账
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLflvmlQuery.field_eq_mod_size"
          placeholder="模具尺寸"
          clearable
          style="width: 150px;; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLflvmlQuery.field_eq_platform_height"
          placeholder="平台高(mm)"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLflvmlQuery.field_eq_mod_height"
          placeholder="模具高度(mm)"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLflvmlQuery.field_eq_mod_pic"
          placeholder="模具图号"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLflvmlQuery.field_eq_mod_num"
          placeholder="模具编号"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLflvmlQuery.field_eq_mod_amount"
          placeholder="模具数量(个)"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLflvmlQuery.field_eq_mod_suit"
          placeholder="适用产品型号"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLflvmlQuery.field_eq_remark"
          placeholder="备注"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLflvmlQuery.field_eq_lflvml_creator"
          style="width: 200px"
          size="small"
          class="filter-item"
          :placeholder="'创建者'"
          clearable
        />
        <el-date-picker
          v-model="created_at_range"
          unlink-panels
          style="margin-top: 10px;width:400px;margin-left:10px"
          type="datetimerange"
          size="small"
          start-placeholder="创建起始日期"
          end-placeholder="创建截止日期"
        />
        <el-date-picker
          v-model="updated_at_range"
          unlink-panels
          style="margin-top: 10px;width:400px;margin-left:10px"
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
          @click="getLflvmlList"
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
      <lflvml-create
        v-model="showCreate"
        @get-lflvml-list="getLflvmlList"
      />
      <lflvml-view
        v-model="showDetail"
        :lflvml="lflvmlDetail"
      />
      <lflvml-update
        v-model="showUpdate"
        :lflvml="lflvmlDetail"
        @get-lflvml-list="getLflvmlList"
      />
      <el-table
        v-loading="loading"
        :data="lflvmlList"
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
          label="模具尺寸(mm)"
          prop="mode_size"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.mod_size }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="平台高(mm)"
          prop="platform_height"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.platform_height }}</span>
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
            <span>{{ row.mod_height }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具图号"
          prop="mod_pic"
          sortable="custom"
          align="center"
          min-width="120px"
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
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.mod_num }}</span>
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
            <span>{{ row.mod_amount }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="适用产品型号"
          prop="mod_suit"
          sortable="custom"
          align="center"
          min-width="220px"
        >
          <template #default="{row}">
            <span>{{ row.mod_suit }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="备注"
          prop="remark"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.remark }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lflvml_creator"
          sortable="custom"
          align="center"
          width="100px"
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
              @click="showLflvmlDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLflvmlUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLflvml(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLflvmlQuery.page"
          v-model:page-size="listLflvmlQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLflvmlList"
          @current-change="getLflvmlList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  DeleteStdlflvmlLibraries,
  GetStdlflvmlLibraries,
  ListLflvmlQuery,
  Lflvml,
  ExportStdlflvmlLibraries
} from '@/api/standard_libraries/lflvml'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LflvmlCreate from './lflvmlCreate.vue'
import LflvmlView from './lflvmlView.vue'
import LflvmlUpdate from './lflvmlUpdate.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import dayjs from 'dayjs'

export default defineComponent({
  components: {
    UserSearchInput,
    LflvmlCreate,
    LflvmlView,
    LflvmlUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLflvmlQuery = reactive({
      field_eq_mod_size: '',
      field_eq_platform_height: null,
      field_eq_mod_height: null,
      field_eq_mod_pic: '',
      field_eq_mod_num: '',
      field_eq_mod_amount: null,
      field_eq_mod_suit: '',
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
    const lflvmlList = ref(Array<Lflvml>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlflvmlLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '箔绕低压模具台账.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '箔绕低压模具台账.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLflvmlList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLflvmlQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLflvmlQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLflvmlQuery.field_gt_created_at = 0
        listLflvmlQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLflvmlQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLflvmlQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLflvmlQuery.field_gt_updated_at = 0
        listLflvmlQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlflvmlLibraries(listLflvmlQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lflvmlList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLflvmlQuery.order_by = column.prop
      listLflvmlQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLflvmlList()
    }
    onMounted(() => {
      getLflvmlList()
    })
    function deleteLflvml(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlflvmlLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLflvmlList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lflvml: Lflvml = {
      id: 0,
      mod_size: '',
      platform_height: null,
      mod_height: null,
      mod_pic: '',
      mod_num: '',
      mod_amount: null,
      mod_suit: '',
      remark: ''
    }
    const lflvmlDetail = ref(lflvml)
    function showLflvmlDetail(lflvml: Lflvml) {
      showDetail.value = true
      lflvmlDetail.value = lflvml
    }
    function showLflvmlUpdate(lflvml: Lflvml) {
      showUpdate.value = true
      lflvmlDetail.value = lflvml
    }
    return {
      handleDownload,
      status,
      loading,
      lflvmlList,
      LflvmlCreate,
      created_at_range,
      updated_at_range,
      listLflvmlQuery,
      total,
      getLflvmlList,
      sortChange,
      deleteLflvml,
      showCreate,
      showDetail,
      showUpdate,
      lflvmlDetail,
      showLflvmlUpdate,
      showLflvmlDetail,
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
