<template>
  <div>
    <aside class="page-header">
      层间绝缘
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLliQuery.field_eq_winding_type"
          placeholder="绕制类型"
          clearable
          style="width: 130px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="lowVol"
          placeholder="额定低压(kV)"
          clearable
          style="width: 170px;margin-right: 10px"
          size="small"
        />
<!--        <el-input-->
<!--          v-model="listLliQuery.field_eq_layer_vol_min"-->
<!--          placeholder="层间电压下界(V)"-->
<!--          clearable-->
<!--          style="width: 170px;margin-right: 10px"-->
<!--          size="small"-->
<!--        />-->
<!--        <el-input-->
<!--          v-model="listLliQuery.field_eq_layer_vol_max"-->
<!--          placeholder="层间电压上界(V)"-->
<!--          clearable-->
<!--          style="width: 170px;margin-right: 10px"-->
<!--          size="small"-->
<!--        />-->
        <el-input
          v-model="layerVol"
          placeholder="层间电压(V)"
          clearable
          style="width: 170px;margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLliQuery.field_eq_layer_insulate"
          placeholder="层间绝缘距离(mm)"
          clearable
          style="width: 170px;margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLliQuery.field_eq_lli_creator"
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
          @click="getLliList"
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
      <lli-create
        v-model="showCreate"
        @get-lli-list="getLliList"
      />
      <lli-view
        v-model="showDetail"
        :lli="lliDetail"
      />
      <lli-update
        v-model="showUpdate"
        :lli="lliDetail"
        @get-lcfiy-list="getLliList"
      />
      <el-table
        v-loading="loading"
        :data="lliList"
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
          label="绕制类型"
          prop="winding_type"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.winding_type }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定低压下界(kV)"
          prop="rated_low_vol_min"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.rated_low_vol_min }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定低压上界(kV)"
          prop="rated_low_vol_max"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.rated_low_vol_max }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="层间电压下界(V)"
          prop="layer_vol_min"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.layer_vol_min }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="层间电压上界(V)"
          :show-overflow-tooltip="true"
          prop="layer_vol_max"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.layer_vol_max }}V</span>
          </template>
        </el-table-column>
        <el-table-column
          label="层间绝缘距离(mm)"
          prop="layer_insulate"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.layer_insulate }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lli_creator"
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
          width="250px"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="primary"
              @click="showLliDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLliUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLli(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLliQuery.page"
          v-model:page-size="listLliQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLliList"
          @current-change="getLliList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlliLibraries, GetStdlliLibraries, ListLliQuery, Lli, ExportStdlliLibraries } from '@/api/standard_libraries/lli'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LliCreate from './lliCreate.vue'
import LliView from './lliView.vue'
import LliUpdate from './lliUpdate.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import dayjs from 'dayjs'

export default defineComponent({
  components: {
    UserSearchInput,
    LliCreate,
    LliView,
    LliUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const lowVol: Ref<number | undefined> = ref(undefined)
    const layerVol: Ref<number | undefined> = ref(undefined)
    const listLliQuery = reactive({
      field_eq_winding_type: '',
      // field_eq_rated_low_vol_min: null,
      // field_eq_rated_low_vol_max: null,
      // field_eq_layer_vol_min: null,
      // field_eq_layer_vol_max: null,
      field_lt_rated_low_vol_min: lowVol,
      field_ge_rated_low_vol_max: lowVol,
      field_lt_layer_vol_min: layerVol,
      field_ge_layer_vol_max: layerVol,
      field_eq_layer_insulate: null,
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
    const lliList = ref(Array<Lli>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlliLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '层间绝缘距离.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '层间绝缘距离.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLliList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLliQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLliQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLliQuery.field_gt_created_at = 0
        listLliQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLliQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLliQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLliQuery.field_gt_updated_at = 0
        listLliQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlliLibraries(listLliQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lliList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLliQuery.order_by = column.prop
      listLliQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLliList()
    }
    onMounted(() => {
      getLliList()
    })
    function deleteLli(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlliLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLliList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lli: Lli = {
      id: 0,
      winding_type: '',
      rated_low_vol_min: null,
      rated_low_vol_max: null,
      layer_vol_min: null,
      layer_vol_max: null,
      layer_insulate: null
    }
    const lliDetail = ref(lli)
    function showLliDetail(lli: Lli) {
      showDetail.value = true
      lliDetail.value = lli
    }
    function showLliUpdate(lli: Lli) {
      showUpdate.value = true
      lliDetail.value = lli
    }
    return {
      handleDownload,
      status,
      loading,
      lliList,
      LliCreate,
      created_at_range,
      updated_at_range,
      listLliQuery,
      total,
      layerVol,
      lowVol,
      getLliList,
      sortChange,
      deleteLli,
      showCreate,
      showDetail,
      showUpdate,
      lliDetail,
      showLliUpdate,
      showLliDetail,
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
