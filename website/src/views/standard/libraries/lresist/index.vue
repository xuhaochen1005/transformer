<template>
  <div>
    <aside class="page-header">
      电阻率
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLresistQuery.field_eq_wire_material"
          placeholder="导线材质"
          clearable
          style="width: 200px;; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLresistQuery.field_eq_temp"
          placeholder="温度（℃）"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLresistQuery.field_eq_lresist_creator"
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
          @click="getLresistList"
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
      <lresist-create
        v-model="showCreate"
        @get-lresist-list="getLresistList"
      />
      <lresist-view
        v-model="showDetail"
        :lresist="lresistDetail"
      />
      <lresist-update
        v-model="showUpdate"
        :lresist="lresistDetail"
        @get-lresist-list="getLresistList"
      />
      <el-table
        v-loading="loading"
        :data="lresistList"
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
          label="导线材质"
          prop="wire_material"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.wire_material }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="温度（℃）"
          prop="temp"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.temp }}℃</span>
          </template>
        </el-table-column>
        <el-table-column
          label="电阻率(Ω·m)"
          prop="resistivity"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.resistivity }}Ω·m</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lresist_creator"
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
              @click="showLresistUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLresist(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLresistQuery.page"
          v-model:page-size="listLresistQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLresistList"
          @current-change="getLresistList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlresistLibraries, GetStdlresistLibraries, ListLresistQuery, Lresist, ExportStdlresistLibraries } from '@/api/standard_libraries/lresist'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LresistCreate from './lresistCreate.vue'
import LresistView from './lresistView.vue'
import LresistUpdate from './lresistUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LresistCreate,
    LresistView,
    LresistUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLresistQuery = reactive({
      field_eq_wire_material: null,
      field_eq_temp: null,
      field_eq_resistivity: null,
      // field_eq_lresist_creator: '',
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
    const lresistList = ref(Array<Lresist>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlresistLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '电阻率.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '电阻率.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLresistList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLresistQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLresistQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLresistQuery.field_gt_created_at = 0
        listLresistQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLresistQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLresistQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLresistQuery.field_gt_updated_at = 0
        listLresistQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlresistLibraries(listLresistQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lresistList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLresistQuery.order_by = column.prop
      listLresistQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLresistList()
    }
    onMounted(() => {
      getLresistList()
    })
    function deleteLresist(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlresistLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLresistList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lresist: Lresist = {
      id: 0,
      wire_material: null,
      temp: null,
      resistivity: null
    }
    const lresistDetail = ref(lresist)
    function showLresistDetail(lresist: Lresist) {
      showDetail.value = true
      lresistDetail.value = lresist
    }
    function showLresistUpdate(lresist: Lresist) {
      showUpdate.value = true
      lresistDetail.value = lresist
    }
    return {
      handleDownload,
      status,
      loading,
      lresistList,
      LresistCreate,
      created_at_range,
      updated_at_range,
      listLresistQuery,
      total,
      getLresistList,
      sortChange,
      deleteLresist,
      showCreate,
      showDetail,
      showUpdate,
      lresistDetail,
      showLresistUpdate,
      showLresistDetail,
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
