<template>
  <div>
    <aside class="page-header">
      铁心磁密
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="capacity"
          placeholder="额定容量(kVA)"
          clearable
          style="width: 170px;margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="magnetDensity"
          placeholder="铁心磁密(T)"
          clearable
          style="width: 170px;margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLmdQuery.field_eq_lmd_creator"
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
          @click="getLmdList"
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
      <Lmd-create
        v-model="showCreate"
        @get-Lmd-list="getLmdList"
      />
      <Lmd-view
        v-model="showDetail"
        :lmd="LmdDetail"
      />
      <Lmd-update
        v-model="showUpdate"
        :lmd="LmdDetail"
        @get-Llvoml-list="getLmdList"
      />
      <el-table
        v-loading="loading"
        :data="LmdList"
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
          label="额定容量下界（kVA）"
          prop="rated_capacity_min"
          sortable="custom"
          align="center"
          min-width="190px"
        >
          <template #default="{row}">
            <span>{{ row.rated_capacity_min }}kVA</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定容量上界（kVA）"
          prop="rated_capacity_max"
          sortable="custom"
          align="center"
          min-width="190px"
        >
          <template #default="{row}">
            <span>{{ row.rated_capacity_max }}kVA</span>
          </template>
        </el-table-column>
        <el-table-column
          label="铁心磁密下界（T）"
          prop="magnet_density_min"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.magnet_density_min }}T</span>
          </template>
        </el-table-column>
        <el-table-column
          label="铁心磁密上界（T）"
          prop="magnet_density_max"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.magnet_density_max }}T</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lmd_creator"
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
          min-width="250px"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="primary"
              @click="showLmdDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLmdUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLmd(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLmdQuery.page"
          v-model:page-size="listLmdQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLmdList"
          @current-change="getLmdList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLmdLibraries, GetStdLmdLibraries, ListLmdQuery, Lmd, ExportStdlmdLibraries } from '@/api/standard_libraries/lmd'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LmdCreate from './lmdCreate.vue'
import LmdView from './lmdView.vue'
import LmdUpdate from './lmdUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LmdCreate,
    LmdView,
    LmdUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const capacity: Ref<number | undefined> = ref(undefined)
    const magnetDensity: Ref<number | undefined> = ref(undefined)
    const listLmdQuery = reactive({
      // field_eq_rated_capacity_min: null,
      // field_eq_rated_capacity_max: null,
      // field_eq_magnet_density_min: null,
      // field_eq_magnet_density_max: null,
      field_lt_rated_capacity_min: capacity,
      field_ge_rated_capacity_max: capacity,
      field_le_magnet_density_min: magnetDensity,
      field_ge_magnet_density_max: magnetDensity,
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
    const LmdList = ref(Array<Lmd>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlmdLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '铁心磁密Bm初选值.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '铁心磁密Bm初选值.xlsx'
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
    async function getLmdList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLmdQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLmdQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLmdQuery.field_gt_created_at = 0
        listLmdQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLmdQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLmdQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLmdQuery.field_gt_updated_at = 0
        listLmdQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLmdLibraries(listLmdQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        LmdList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLmdQuery.order_by = column.prop
      listLmdQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLmdList()
    }
    onMounted(() => {
      getLmdList()
    })
    function deleteLmd(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLmdLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLmdList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lmd: Lmd = {
      id: 0,
      rated_capacity_min: 0,
      rated_capacity_max: 0,
      magnet_density_min: 0,
      magnet_density_max: 0,
      created_at: '',
      updated_at: ''
    }
    const LmdDetail = ref(lmd)
    function showLmdDetail(Lmd: Lmd) {
      showDetail.value = true
      LmdDetail.value = Lmd
    }
    function showLmdUpdate(Lmd: Lmd) {
      showUpdate.value = true
      LmdDetail.value = Lmd
    }
    return {
      status,
      handleDownload,
      loading,
      LmdList,
      LmdCreate,
      created_at_range,
      updated_at_range,
      listLmdQuery,
      total,
      capacity,
      magnetDensity,
      getLmdList,
      sortChange,
      deleteLmd,
      showCreate,
      showDetail,
      showUpdate,
      LmdDetail,
      showLmdUpdate,
      showLmdDetail,
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
