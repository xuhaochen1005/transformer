<template>
  <div>
    <aside class="page-header">
      电流密度
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLcdQuery.field_eq_winding_material"
          filterable
          allow-create
          clearable
          style="width: 200px; margin-right: 10px"
          placeholder="导线材料"
          size="small"
        >
          <el-option
            label="铜"
            value="铜"
          />
          <el-option
            label="铝"
            value="铝"
          />
        </el-select>
        <el-input
          v-model="currentDensity"
          placeholder="电流密度(A/mm^2)"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLcdQuery.field_eq_lcd_creator"
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
          @click="getLcdList"
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
      <lcd-create
        v-model="showCreate"
        @get-lcd-list="getLcdList"
      />
      <lcd-view
        v-model="showDetail"
        :lcd="lcdDetail"
      />
      <lcd-update
        v-model="showUpdate"
        :lcd="lcdDetail"
        @get-lcd-list="getLcdList"
      />
      <el-table
        v-loading="loading"
        :data="lcdList"
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
        >
          <template #default="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="导线材料"
          prop="winding_material"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.winding_material }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="电流密度下界(A/mm^2)"
          prop="current_density_min"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.current_density_min }}A/mm^2</span>
          </template>
        </el-table-column>
        <el-table-column
          label="电流密度上界(A/mm^2)"
          prop="current_density_max"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.current_density_max }}A/mm^2</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lcd_creator"
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
              @click="showLcdUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLcd(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLcdQuery.page"
          v-model:page-size="listLcdQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLcdList"
          @current-change="getLcdList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlcdLibraries, GetStdlcdLibraries, ListLcdQuery, Lcd, ExportStdlcdLibraries } from '@/api/standard_libraries/lcd'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LcdCreate from './lcdCreate.vue'
import LcdView from './lcdView.vue'
import LcdUpdate from './lcdUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LcdCreate,
    LcdView,
    LcdUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const select = ref(Array<Lcd>())
    const currentDensity: Ref<number | undefined> = ref(undefined)
    const listLcdQuery = reactive({
      field_eq_winding_material: '',
      // field_lk_current_density_min: null,
      // field_lk_current_density_max: null,
      field_lt_current_density_min: currentDensity,
      field_ge_current_density_max: currentDensity,
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
    const lcdList = ref(Array<Lcd>())
    const total = ref(0)
    async function Focus() {
      loading.value = true
      const response = await GetStdlcdLibraries(listLcdQuery)
      select.value = response.data.total ? response.data.spec : []
      loading.value = false
    }
    function cdChange(cd:string) {
      listLcdQuery.field_eq_winding_material = cd
    }
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlcdLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '电流密度.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '电流密度.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLcdList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLcdQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLcdQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLcdQuery.field_gt_created_at = 0
        listLcdQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLcdQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLcdQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLcdQuery.field_gt_updated_at = 0
        listLcdQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlcdLibraries(listLcdQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lcdList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLcdQuery.order_by = column.prop
      listLcdQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLcdList()
    }
    onMounted(() => {
      getLcdList()
    })
    function deleteLcd(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlcdLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLcdList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lcd: Lcd = {
      id: 0,
      winding_material: '',
      current_density_min: null,
      current_density_max: null
    }
    const lcdDetail = ref(lcd)
    function showLcdDetail(lcd: Lcd) {
      showDetail.value = true
      lcdDetail.value = lcd
    }
    function showLcdUpdate(lcd: Lcd) {
      showUpdate.value = true
      lcdDetail.value = lcd
    }
    return {
      handleDownload,
      select,
      Focus,
      cdChange,
      status,
      loading,
      lcdList,
      LcdCreate,
      created_at_range,
      updated_at_range,
      listLcdQuery,
      total,
      currentDensity,
      getLcdList,
      sortChange,
      deleteLcd,
      showCreate,
      showDetail,
      showUpdate,
      lcdDetail,
      showLcdUpdate,
      showLcdDetail,
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
