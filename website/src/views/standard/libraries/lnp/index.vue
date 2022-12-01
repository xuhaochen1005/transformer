<template>
  <div>
    <aside class="page-header">
      噪声预测值
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
          v-model="highVol"
          placeholder="额定高压(kV)"
          clearable
          style="width: 170px;margin-right: 10px"
          size="small"
        />
        <el-select
          v-model="listLnpQuery.cool_type"
          filterable
          allow-create
          placeholder="冷却方式"
          clearable
          style="width: 130px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="AN"
            value="AN"
          />
        </el-select>
        <UserSearchInput
          v-model="listLnpQuery.field_eq_lnp_creator"
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
          @click="getLnpList"
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
      <Lnp-create
        v-model="showCreate"
        @get-Lnp-list="getLnpList"
      />
      <Lnp-view
        v-model="showDetail"
        :lnp="LnpDetail"
      />
      <Lnp-update
        v-model="showUpdate"
        :lnp="LnpDetail"
        @get-Lnp-list="getLnpList"
      />
      <el-table
        v-loading="loading"
        :data="LnpList"
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
          label="额定高压下界（kV）"
          prop="rated_high_vol_min"
          sortable="custom"
          align="center"
          min-width="190px"
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
          min-width="190px"
        >
          <template #default="{row}">
            <span>{{ row.rated_high_vol_max }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="冷却方式"
          prop="cool_type"
          sortable="custom"
          align="center"
          min-width="110x"
        >
          <template #default="{row}">
            <span>{{ row.cool_type }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="噪声预测（dB）"
          prop="noise_predict"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.noise_predict }}dB</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lnp_creator"
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
              @click="showLnpDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLnpUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLnp(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLnpQuery.page"
          v-model:page-size="listLnpQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLnpList"
          @current-change="getLnpList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLnpLibraries, GetStdLnpLibraries, ListLnpQuery, Lnp, ExportStdlnpLibraries } from '@/api/standard_libraries/lnp'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LnpCreate from './lnpCreate.vue'
import LnpView from './lnpView.vue'
import LnpUpdate from './lnpUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LnpCreate,
    LnpView,
    LnpUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const highVol: Ref<number | undefined> = ref(undefined)
    const capacity: Ref<number | undefined> = ref(undefined)
    const listLnpQuery = reactive({
      // field_eq_rated_capacity_min: null,
      // field_eq_rated_capacity_max: null,
      // field_eq_rated_high_vol_min: null,
      // field_eq_rated_high_vol_max: null,
      field_lt_rated_capacity_min: capacity,
      field_ge_rated_capacity_max: capacity,
      field_lt_rated_high_vol_min: highVol,
      field_ge_rated_high_vol_max: highVol,
      field_eq_cool_type: null,
      field_eq_noise_predict: null,
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
    const LnpList = ref(Array<Lnp>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlnpLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '噪声预测.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '噪声预测.xlsx'
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
    async function getLnpList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLnpQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLnpQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLnpQuery.field_gt_created_at = 0
        listLnpQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLnpQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLnpQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLnpQuery.field_gt_updated_at = 0
        listLnpQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLnpLibraries(listLnpQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        LnpList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLnpQuery.order_by = column.prop
      listLnpQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLnpList()
    }
    onMounted(() => {
      getLnpList()
    })
    function deleteLnp(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLnpLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLnpList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lnp: Lnp = {
      id: 0,
      rated_capacity_min: 0,
      rated_capacity_max: 0,
      rated_high_vol_min: 0,
      rated_high_vol_max: 0,
      cool_type: '',
      noise_predict: 0,
      created_at: '',
      updated_at: '',
      deleted_at: ''
    }
    const LnpDetail = ref(lnp)
    function showLnpDetail(Lnp: Lnp) {
      showDetail.value = true
      LnpDetail.value = Lnp
    }
    function showLnpUpdate(Lnp: Lnp) {
      showUpdate.value = true
      LnpDetail.value = Lnp
    }
    return {
      status,
      handleDownload,
      loading,
      LnpList,
      LnpCreate,
      created_at_range,
      updated_at_range,
      listLnpQuery,
      total,
      capacity,
      highVol,
      getLnpList,
      sortChange,
      deleteLnp,
      showCreate,
      showDetail,
      showUpdate,
      LnpDetail,
      showLnpUpdate,
      showLnpDetail,
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
