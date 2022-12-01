<template>
  <div>
    <aside class="page-header">
      风道条台账
    </aside>
    <div style="margin:10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLadblQuery.field_lk_air_duct_bar_thickness"
          placeholder="风道条厚度(mm)"
          clearable
          style="width: 145px;margin-left:10px"
          size="small"
        />
        <el-input
          v-model="listLadblQuery.field_lk_air_duct_bar_width"
          placeholder="风道条宽度(mm)"
          clearable
          style="width: 145px;margin-left:10px"
          size="small"
        />
        <el-input
          v-model="listLadblQuery.field_lk_air_duct_bar_length"
          placeholder="风道条长度(mm)"
          clearable
          style="width: 145px;margin-left:10px"
          size="small"
        />
        <el-input
          v-model="listLadblQuery.field_lk_air_duct_bar_num"
          placeholder="风道条数量(个)"
          clearable
          style="width: 145px;margin-left:10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLadblQuery.field_eq_ladbl_creator"
          style="width: 145px;margin-left:10px"
          size="small"
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
          @click="getLadblList"
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
          :loading="downloading"
          class="filter-item"
          type="primary"
          icon="el-icon-download"
          @click="handleDownload"
        >
          导出
        </el-button>
      </div>
      <ladbl-create
        v-model="showCreate"
        @get-ladbl-list="getLadblList"
      />
      <ladbl-view
        v-model="showDetail"
        :ladbl="ladblDetail"
      />
      <ladbl-update
        v-model="showUpdate"
        :ladbl="ladblDetail"
        @get-ladbl-list="getLadblList"
      />
      <el-table
        v-loading="loading"
        :data="ladblList"
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
          width="80px"
        >
          <template #default="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="风道条厚度(mm)"
          prop="air_duct_bar_thickness"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.air_duct_bar_thickness }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="风道条宽度(mm)"
          prop="air_duct_bar_width"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.air_duct_bar_width }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="风道条长度(mm)"
          prop="air_duct_bar_length"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.air_duct_bar_length }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="风道条数量(个)"
          fit
          prop="air_duct_bar_num"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.air_duct_bar_num }}个</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="ladbl_creator"
          sortable="custom"
          align="center"
          width="180px"
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
          width="150px"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="success"
              @click="showLadblUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLadbl(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLadblQuery.page"
          v-model:page-size="listLadblQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLadblList"
          @current-change="getLadblList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { reactive, ref, onMounted, Ref, defineComponent } from 'vue'
import { ElMessage, ElMessageBox, ElConfigProvider } from 'element-plus'
import { DeleteStdladblLibraries, GetStdladblLibraries, ListLadblQuery, Ladbl, ExportStdladblLibraries } from '@/api/standard_libraries/ladbl'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LadblCreate from './ladblCreate.vue'
import LadblView from './ladblView.vue'
import LadblUpdate from './ladblUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
export default defineComponent({
  components: {
    UserSearchInput,
    LadblCreate,
    LadblView,
    LadblUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLadblQuery = reactive({
      field_lk_air_duct_bar_thickness: null,
      field_lk_air_duct_bar_width: null,
      field_lk_air_duct_bar_length: null,
      field_lk_air_duct_bar_num: null,
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
    const downloading = ref(false)
    const ladblList : Ref<Ladbl[]> = ref([])
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      downloading.value = true
      const res = await ExportStdladblLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '风道条台账.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '风道条台账.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      downloading.value = false
    }
    async function getLadblList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLadblQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLadblQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLadblQuery.field_gt_created_at = 0
        listLadblQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLadblQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLadblQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLadblQuery.field_gt_updated_at = 0
        listLadblQuery.field_lt_updated_at = 0
      }
      const response = await GetStdladblLibraries(listLadblQuery)
      if (response.data.code === 200) {
        total.value = response.data.total
        ladblList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLadblQuery.order_by = column.prop
      listLadblQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLadblList()
    }
    onMounted(() => {
      getLadblList()
    })
    function deleteLadbl(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdladblLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLadblList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const ladbl: Ladbl = {
      id: 0,
      air_duct_bar_thickness: 0,
      air_duct_bar_width: 0,
      air_duct_bar_length: 0,
      air_duct_bar_num: 0,
      ladbl_creator: 0
    }
    const ladblDetail = ref(ladbl)
    function showLadblDetail(ladbl: Ladbl) {
      showDetail.value = true
      ladblDetail.value = ladbl
    }
    function showLadblUpdate(ladbl: Ladbl) {
      showUpdate.value = true
      ladblDetail.value = ladbl
    }
    return {
      status,
      handleDownload,
      downloading,
      loading,
      ladblList,
      LadblCreate,
      created_at_range,
      updated_at_range,
      listLadblQuery,
      total,
      getLadblList,
      sortChange,
      deleteLadbl,
      showCreate,
      showDetail,
      showUpdate,
      ladblDetail,
      showLadblUpdate,
      showLadblDetail,
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
