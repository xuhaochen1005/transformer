<template>
  <div>
    <aside class="page-header">
      导线绝缘厚度表
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLwiQuery.field_eq_wire_id"
          placeholder="导线编号"
          clearable
          style="width: 150px;; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="wireWidth"
          placeholder="线厚(mm)"
          clearable
          style="width: 160px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLwiQuery.field_eq_axial_insulation"
          placeholder="轴向绝缘厚(mm)"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLwiQuery.field_eq_radial_insulation"
          placeholder="辐向绝缘厚(mm)"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLwiQuery.field_eq_lwi_creator"
          style="width: 150px"
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
          @click="getLwiList"
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
      <lwi-create
        v-model="showCreate"
        @get-lwi-list="getLwiList"
      />
      <lwi-view
        v-model="showDetail"
        :lwi="lwiDetail"
      />
      <lwi-update
        v-model="showUpdate"
        :lwi="lwiDetail"
        @get-lwi-list="getLwiList"
      />
      <el-table
        v-loading="loading"
        :data="lwiList"
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
          label="导线编号"
          prop="wire_id"
          sortable="custom"
          align="center"
          min-width="120px"
        >
          <template #default="{row}">
            <span>{{ row.wire_id }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="线厚区间下界(mm)"
          prop="width_min"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.width_min }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="线厚区间上界(mm)"
          prop="width_max"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.width_max }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="轴向绝缘厚(mm)"
          prop="axial_insulation"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.axial_insulation }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="辐向绝缘厚(mm)"
          prop="radial_insulation"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.radial_insulation }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lwi_creator"
          sortable="custom"
          align="center"
          min-width="120px"
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
          min-width="200px"
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
          min-width="200px"
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
              @click="showLwiDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLwiUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLwi(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLwiQuery.page"
          v-model:page-size="listLwiQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLwiList"
          @current-change="getLwiList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlwiLibraries, GetStdlwiLibraries, ListLwiQuery, Lwi, ExportStdlwiLibraries } from '@/api/standard_libraries/lwi'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LwiCreate from './lwiCreate.vue'
import LwiView from './lwiView.vue'
import LwiUpdate from './lwiUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LwiCreate,
    LwiView,
    LwiUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const wireWidth: Ref<number | undefined> = ref(undefined)
    const listLwiQuery = reactive({
      field_eq_wire_id: null,
      field_lt_width_min: wireWidth,
      field_ge_width_max: wireWidth,
      // field_eq_width_min: null,
      // field_eq_width_max: null,
      field_eq_axial_insulation: null,
      field_eq_radial_insulation: null,
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
    const lwiList = ref(Array<Lwi>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlwiLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '导线绝缘厚度表.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '导线绝缘厚度表.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLwiList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLwiQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLwiQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLwiQuery.field_gt_created_at = 0
        listLwiQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLwiQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLwiQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLwiQuery.field_gt_updated_at = 0
        listLwiQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlwiLibraries(listLwiQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lwiList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLwiQuery.order_by = column.prop
      listLwiQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLwiList()
    }
    onMounted(() => {
      getLwiList()
    })
    function deleteLwi(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlwiLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLwiList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lwi: Lwi = {
      id: 0,
      wire_id: null,
      width_min: null,
      width_max: null,
      axial_insulation: null,
      radial_insulation: null
    }
    const lwiDetail = ref(lwi)
    function showLwiDetail(lwi: Lwi) {
      showDetail.value = true
      lwiDetail.value = lwi
    }
    function showLwiUpdate(lwi: Lwi) {
      showUpdate.value = true
      lwiDetail.value = lwi
    }
    return {
      handleDownload,
      status,
      loading,
      lwiList,
      LwiCreate,
      created_at_range,
      updated_at_range,
      listLwiQuery,
      total,
      wireWidth,
      getLwiList,
      sortChange,
      deleteLwi,
      showCreate,
      showDetail,
      showUpdate,
      lwiDetail,
      showLwiUpdate,
      showLwiDetail,
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
