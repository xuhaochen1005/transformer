<template>
  <div>
    <aside class="page-header">
      硅钢片
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLstackQuery.field_eq_type"
          placeholder="硅钢片型号"
          clearable
          style="width: 200px;; margin-right: 10px"
          size="small"
        />
        <el-input-number
          v-model="listLstackQuery.field_eq_density"
          placeholder="密度(kg/cm^3)"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        />
        <el-input-number
          v-model="listLstackQuery.field_eq_price"
          placeholder="单价(元/kg)"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLstackQuery.field_eq_lstack_creator"
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
          @click="getLstackList"
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
      <lstack-create
        v-model="showCreate"
        @get-lstack-list="getLstackList"
      />
      <lstack-view
        v-model="showDetail"
        :lstack="lstackDetail"
      />
      <lstack-update
        v-model="showUpdate"
        :lstack="lstackDetail"
        @get-lstack-list="getLstackList"
      />
      <el-table
        v-loading="loading"
        :data="lstackList"
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
          label="硅钢片型号"
          prop="type"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.type }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="密度(kg/cm^3)"
          prop="density"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.density }}kg/cm^3</span>
          </template>
        </el-table-column>
        <el-table-column
          label="单价(元/kg)"
          prop="price"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.price }}元/kg</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lstack_creator"
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
              @click="showLstackUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLstack(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLstackQuery.page"
          v-model:page-size="listLstackQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLstackList"
          @current-change="getLstackList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlstackLibraries, GetStdlstackLibraries, ListLstackQuery, Lstack, ExportStdlstackLibraries } from '@/api/standard_libraries/lstack'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LstackCreate from './lstackCreate.vue'
import LstackView from './lstackView.vue'
import LstackUpdate from './lstackUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LstackCreate,
    LstackView,
    LstackUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLstackQuery = reactive({
      field_eq_type: null,
      field_eq_density: null,
      field_eq_price: null,
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
    const lstackList = ref(Array<Lstack>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlstackLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '硅钢片.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '硅钢片.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLstackList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLstackQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLstackQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLstackQuery.field_gt_created_at = 0
        listLstackQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLstackQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLstackQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLstackQuery.field_gt_updated_at = 0
        listLstackQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlstackLibraries(listLstackQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lstackList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLstackQuery.order_by = column.prop
      listLstackQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLstackList()
    }
    onMounted(() => {
      getLstackList()
    })
    function deleteLstack(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlstackLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLstackList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lstack: Lstack = {
      id: 0,
      type: null,
      density: null,
      price: null
    }
    const lstackDetail = ref(lstack)
    function showLstackDetail(lstack: Lstack) {
      showDetail.value = true
      lstackDetail.value = lstack
    }
    function showLstackUpdate(lstack: Lstack) {
      showUpdate.value = true
      lstackDetail.value = lstack
    }
    return {
      handleDownload,
      status,
      loading,
      lstackList,
      LstackCreate,
      created_at_range,
      updated_at_range,
      listLstackQuery,
      total,
      getLstackList,
      sortChange,
      deleteLstack,
      showCreate,
      showDetail,
      showUpdate,
      lstackDetail,
      showLstackUpdate,
      showLstackDetail,
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
