<template>
  <div>
    <aside class="page-header">
      导线表
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLwireQuery.field_eq_wire_type"
          placeholder="导线型号"
          clearable
          style="width: 200px;; margin-right: 10px"
          size="small"
        />
        <el-select
          v-model="listLwireQuery.field_eq_wire_material"
          filterable
          allow-create
          placeholder="导线材质"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="铜线"
            value="铜线"
          />
          <el-option
            label="铜箔"
            value="铜箔"
          />
          <el-option
            label="铝线"
            value="铝线"
          />
          <el-option
            label="铝箔"
            value="铝箔"
          />
        </el-select>
        <el-select
          v-model="listLwireQuery.field_eq_wire_shape"
          placeholder="导线形状"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="圆"
            value="圆"
          />
          <el-option
            label="扁"
            value="扁"
          />
        </el-select>
        <el-input-number
          v-model="listLwireQuery.field_eq_wire_density"
          placeholder="导线密度(kg/m^3)"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        />
        <el-input-number
          v-model="listLwireQuery.field_eq_wire_price"
          placeholder="导线价格(元/m)"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLwireQuery.field_eq_lwire_creator"
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
          @click="getLwireList"
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
      <lwire-create
        v-model="showCreate"
        @get-lwire-list="getLwireList"
      />
      <lwire-view
        v-model="showDetail"
        :lwire="lwireDetail"
      />
      <lwire-update
        v-model="showUpdate"
        :lwire="lwireDetail"
        @get-lwire-list="getLwireList"
      />
      <el-table
        v-loading="loading"
        :data="lwireList"
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
          label="导线型号"
          prop="wire_type"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.wire_type }}</span>
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
          label="导线形状"
          prop="wire_shape"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.wire_shape }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="导线密度(kg/m^3)"
          prop="wire_density"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.wire_density }}kg/m^3</span>
          </template>
        </el-table-column>
        <el-table-column
          label="导线价格（元/m）"
          prop="wire_price"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.wire_price }}元/m</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lwire_creator"
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
              @click="showLwireUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLwire(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLwireQuery.page"
          v-model:page-size="listLwireQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLwireList"
          @current-change="getLwireList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlwireLibraries, GetStdlwireLibraries, ListLwireQuery, Lwire, ExportStdlwireLibraries } from '@/api/standard_libraries/lwire'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LwireCreate from './lwireCreate.vue'
import LwireView from './lwireView.vue'
import LwireUpdate from './lwireUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LwireCreate,
    LwireView,
    LwireUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLwireQuery = reactive({
      field_eq_wire_type: null,
      field_eq_wire_material: null,
      field_eq_wire_shape: null,
      field_eq_wire_density: null,
      field_eq_wire_price: undefined,
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
    const lwireList = ref(Array<Lwire>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlwireLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '导线表.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '导线表.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLwireList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLwireQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLwireQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLwireQuery.field_gt_created_at = 0
        listLwireQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLwireQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLwireQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLwireQuery.field_gt_updated_at = 0
        listLwireQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlwireLibraries(listLwireQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lwireList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLwireQuery.order_by = column.prop
      listLwireQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLwireList()
    }
    onMounted(() => {
      getLwireList()
    })
    function deleteLwire(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlwireLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLwireList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lwire: Lwire = {
      id: 0,
      wire_type: null,
      wire_material: null,
      wire_shape: null,
      wire_density: null
    }
    const lwireDetail = ref(lwire)
    function showLwireDetail(lwire: Lwire) {
      showDetail.value = true
      lwireDetail.value = lwire
    }
    function showLwireUpdate(lwire: Lwire) {
      showUpdate.value = true
      lwireDetail.value = lwire
    }
    return {
      handleDownload,
      status,
      loading,
      lwireList,
      LwireCreate,
      created_at_range,
      updated_at_range,
      listLwireQuery,
      total,
      getLwireList,
      sortChange,
      deleteLwire,
      showCreate,
      showDetail,
      showUpdate,
      lwireDetail,
      showLwireUpdate,
      showLwireDetail,
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
