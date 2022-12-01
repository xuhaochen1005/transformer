<template>
  <div>
    <aside class="page-header">
      主风道宽
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="highVol"
          placeholder="额定高压(kV)"
          clearable
          style="width: 170px;margin-right: 10px"
          size="small"
        />
        <el-select
          v-model="listLmadQuery.field_eq_main_air_duct_min"
          filterable
          allow-create
          placeholder="主风道初选最小值（mm）"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        >
          <el-option
            v-for="item in airDuct"
            :key="item"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <UserSearchInput
          v-model="listLmadQuery.field_eq_lmad_creator"
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
          @click="getLmadList"
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
      <Lmad-create
        v-model="showCreate"
        @get-Lmad-list="getLmadList"
      />
      <Lmad-view
        v-model="showDetail"
        :lmad="LmadDetail"
      />
      <Lmad-update
        v-model="showUpdate"
        :lmad="LmadDetail"
        @get-Lmad-list="getLmadList"
      />
      <el-table
        v-loading="loading"
        :data="LmadList"
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
          width="80px"
        >
          <template #default="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定高压下界（kV）"
          prop="rated_high_vol_min"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.rated_high_vol_min }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定高压上界（kV）"
          prop="rated_high_vol_max"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.rated_high_vol_max }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="主风道初选最小值（mm）"
          prop="main_air_duct_min"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.main_air_duct_min }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lmad_creator"
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
          min-width="110px"
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
          min-width="110px"
        >
          <template #default="{row}">
            <span>{{ UnixTime2HumanWithStr(row.updated_at) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          align="center"
          width="250px"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="success"
              @click="showLmadUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLmad(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLmadQuery.page"
          v-model:page-size="listLmadQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLmadList"
          @current-change="getLmadList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLmadLibraries, GetStdLmadLibraries, ListLmadQuery, Lmad, ExportStdlmadLibraries } from '@/api/standard_libraries/lmad'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LmadCreate from './lmadCreate.vue'
import LmadView from './lmadView.vue'
import LmadUpdate from './lmadUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LmadCreate,
    LmadView,
    LmadUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const highVol: Ref<number | undefined> = ref(undefined)
    const listLmadQuery = reactive({
      // field_eq_rated_high_vol_min: null,
      // field_eq_rated_high_vol_max: null,
      field_lt_rated_high_vol_min: highVol,
      field_ge_rated_high_vol_max: highVol,
      field_eq_main_air_duct_min: null,
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
    const LmadList = ref(Array<Lmad>())
    const total = ref(0)
    const airDuct : Ref<{label: number, value: number}[]> = ref([])
    const getStdLmadLibraries = async () => {
      const res = await GetStdLmadLibraries({
        field_eq_main_air_duct_min: 0,
        limit: 9999,
        page: 1,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        airDuct.value = res.data.spec.map((i: Lmad) => {
          return {
            label: i.main_air_duct_min + 'mm',
            value: i.main_air_duct_min
          }
        })
      }
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlmadLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '主风道.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '主风道.xlsx'
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
    async function getLmadList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLmadQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLmadQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLmadQuery.field_gt_created_at = 0
        listLmadQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLmadQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLmadQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLmadQuery.field_gt_updated_at = 0
        listLmadQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLmadLibraries(listLmadQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        LmadList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLmadQuery.order_by = column.prop
      listLmadQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLmadList()
    }
    onMounted(() => {
      getLmadList()
    })
    function deleteLmad(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLmadLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLmadList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lmad: Lmad = {
      id: 0,
      rated_high_vol_min: 0,
      rated_high_vol_max: 0,
      main_air_duct_min: 0,
      created_at: '',
      updated_at: '',
      deleted_at: ''
    }
    const LmadDetail = ref(lmad)
    function showLmadDetail(Lmad: Lmad) {
      showDetail.value = true
      LmadDetail.value = Lmad
    }
    function showLmadUpdate(Lmad: Lmad) {
      showUpdate.value = true
      LmadDetail.value = Lmad
    }
    const paramLoading = ref(true)
    onMounted(async () => {
      await Promise.all([
        getStdLmadLibraries()
      ])
      paramLoading.value = false
    })
    return {
      status,
      handleDownload,
      loading,
      LmadList,
      LmadCreate,
      airDuct,
      created_at_range,
      updated_at_range,
      listLmadQuery,
      total,
      highVol,
      getLmadList,
      sortChange,
      deleteLmad,
      showCreate,
      showDetail,
      showUpdate,
      LmadDetail,
      showLmadUpdate,
      showLmadDetail,
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
