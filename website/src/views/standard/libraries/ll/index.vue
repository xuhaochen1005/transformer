<template>
  <div>
    <aside class="page-header">
      损耗相关（空载损耗，负载损耗，空载电流，短路阻抗）
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLlQuery.field_eq_loss_level"
          filterable
          allow-create
          placeholder="损耗水平代号"
          clearable
          style="width: 130px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="10"
            value="10"
          />
          <el-option
            label="11"
            value="11"
          />
          <el-option
            label="12"
            value="12"
          />
          <el-option
            label="13"
            value="13"
          />
        </el-select>
        <el-select
          v-model="listLlQuery.field_eq_regulate_way"
          filterable
          allow-create
          placeholder="调压方式"
          clearable
          style="width: 110px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="无励磁调压"
            value="无励磁调压"
          />
          <el-option
            label="有载调压"
            value="有载调压"
          />
        </el-select>
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
        <el-input
          v-model="lowVol"
          placeholder="额定低压(kV)"
          clearable
          style="width: 170px;margin-right: 10px"
          size="small"
        />
        <el-select
          v-model="listLlQuery.field_eq_regulate_range_min"
          placeholder="调压下界(%)"
          clearable
          style="width: 130px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="0%"
            value="0%"
          />
          <el-option
            label="-10%"
            value="-10%"
          />
          <el-option
            label="-5%"
            value="-5%"
          />
          <el-option
            label="-2.5%"
            value="-2.5%"
          />
        </el-select>
        <el-select
          v-model="listLlQuery.field_eq_regulate_range_max"
          placeholder="调压上界(%)"
          clearable
          style="width: 130px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="0%"
            value="0%"
          />
          <el-option
            label="2.5%"
            value="2.5%"
          />
          <el-option
            label="5%"
            value="5%"
          />
          <el-option
            label="10%"
            value="10%"
          />
        </el-select>
        <el-select
          v-model="listLlQuery.field_eq_regulate_range_step"
          placeholder="调压步长(%)"
          clearable
          style="width: 130px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="0%"
            value="0%"
          />
          <el-option
            label="2.5%"
            value="2.5%"
          />
          <el-option
            label="5%"
            value="5%"
          />
        </el-select>
        <el-select
          v-model="listLlQuery.field_eq_temperature"
          placeholder="绝缘温度(K)"
          clearable
          style="width: 130px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="100K"
            value="100K"
          />
          <el-option
            label="120K"
            value="120K"
          />
          <el-option
            label="130K"
            value="130K"
          />
          <el-option
            label="145K"
            value="145K"
          />
          <el-option
            label="150K"
            value="150K"
          />
          <el-option
            label="180K"
            value="180K"
          />
        </el-select>
        <UserSearchInput
          v-model="listLlQuery.field_eq_ll_creator"
          style="width: 180px"
          size="small"
          class="filter-item"
          :placeholder="'创建者'"
          clearable
        />
        <div class="block">
          <el-date-picker
            v-model="created_at_range"
            unlink-panels
            style="margin-top: 10px;width:280px"
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
            @click="getLlList"
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
      </div>
      <Ll-create
        v-model="showCreate"
        @get-Ll-list="getLlList"
      />
      <Ll-view
        v-model="showDetail"
        :ll="LlDetail"
      />
      <Ll-update
        v-model="showUpdate"
        :ll="LlDetail"
        @get-Ll-list="getLlList"
      />
      <el-table
        v-loading="loading"
        :data="LlList"
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
          label="损耗水平代号"
          prop="regulate_way"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.loss_level }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="调压方式"
          prop="regulate_way"
          sortable="custom"
          align="center"
          min-width="110px"
        >
          <template #default="{row}">
            <span>{{ row.regulate_way }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定容量下界"
          prop="rated_capacity_min"
          sortable="custom"
          align="center"
          min-width="130px"
        >
          <template #default="{row}">
            <span>{{ row.rated_capacity_min }}kVA</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定容量上界"
          prop="rated_capacity_max"
          sortable="custom"
          align="center"
          min-width="130px"
        >
          <template #default="{row}">
            <span>{{ row.rated_capacity_max }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定高压下界"
          prop="rated_high_vol_min"
          sortable="custom"
          align="center"
          min-width="130px"
        >
          <template #default="{row}">
            <span>{{ row.rated_high_vol_min }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定高压上界"
          prop="rated_high_vol_max"
          sortable="custom"
          align="center"
          min-width="130x"
        >
          <template #default="{row}">
            <span>{{ row.rated_high_vol_max }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定低压下界"
          prop="rated_low_vol_min"
          sortable="custom"
          align="center"
          min-width="130px"
        >
          <template #default="{row}">
            <span>{{ row.rated_low_vol_min }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定低压上界"
          prop="rated_low_vol_max"
          sortable="custom"
          align="center"
          min-width="130px"
        >
          <template #default="{row}">
            <span>{{ row.rated_low_vol_max }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="调压范围下界"
          prop="regulate_range_min"
          sortable="custom"
          align="center"
          min-width="130px"
        >
          <template #default="{row}">
            <span>{{ row.regulate_range_min }}%</span>
          </template>
        </el-table-column>
        <el-table-column
          label="调压范围上界"
          prop="regulate_range_max"
          sortable="custom"
          align="center"
          min-width="130px"
        >
          <template #default="{row}">
            <span>{{ row.regulate_range_max }}%</span>
          </template>
        </el-table-column>
        <el-table-column
          label="调压范围步长"
          prop="regulate_range_step"
          sortable="custom"
          align="center"
          min-width="130px"
        >
          <template #default="{row}">
            <span>{{ row.regulate_range_step }}%</span>
          </template>
        </el-table-column>
        <el-table-column
          label="绝缘系统温度"
          prop="temperature"
          sortable="custom"
          align="center"
          min-width="130px"
        >
          <template #default="{row}">
            <span>{{ row.temperature }}K</span>
          </template>
        </el-table-column>
        <el-table-column
          label="负载损耗"
          prop="load_loss"
          sortable="custom"
          align="center"
          min-width="110px"
        >
          <template #default="{row}">
            <span>{{ row.load_loss }}W</span>
          </template>
        </el-table-column>
        <el-table-column
          label="空载损耗"
          prop="no_load_loss"
          sortable="custom"
          align="center"
          min-width="110px"
        >
          <template #default="{row}">
            <span>{{ row.no_load_loss }}W</span>
          </template>
        </el-table-column>
        <el-table-column
          label="空载电流"
          prop="no_load_current"
          sortable="custom"
          align="center"
          min-width="110px"
        >
          <template #default="{row}">
            <span>{{ row.no_load_current }}%</span>
          </template>
        </el-table-column>
        <el-table-column
          label="短路阻抗"
          prop="short_circuit_imped"
          sortable="custom"
          align="center"
          min-width="110px"
        >
          <template #default="{row}">
            <span>{{ row.short_circuit_imped }}%</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="ll_creator"
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
              @click="showLlDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLlUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLl(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLlQuery.page"
          v-model:page-size="listLlQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLlList"
          @current-change="getLlList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLlLibraries, GetStdLlLibraries, ListLlQuery, Ll, ExportStdllLibraries } from '@/api/standard_libraries/ll'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LlCreate from './llCreate.vue'
import LlView from './llView.vue'
import LlUpdate from './llUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LlCreate,
    LlView,
    LlUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const capacity: Ref<number | undefined> = ref(undefined)
    const highVol: Ref<number | undefined> = ref(undefined)
    const lowVol: Ref<number | undefined> = ref(undefined)
    const listLlQuery = reactive({
      field_eq_loss_level: null,
      field_eq_regulate_way: null,
      field_lt_rated_capacity_min: capacity,
      field_ge_rated_capacity_max: capacity,
      field_lt_rated_high_vol_min: highVol,
      field_ge_rated_high_vol_max: highVol,
      field_lt_rated_low_vol_min: lowVol,
      field_ge_rated_low_vol_max: lowVol,
      field_eq_regulate_range_min: null,
      field_eq_regulate_range_max: null,
      field_eq_regulate_range_step: null,
      field_eq_temperature: null,
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
    const LlList = ref(Array<Ll>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdllLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '损耗.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '损耗.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLlList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLlQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLlQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLlQuery.field_gt_created_at = 0
        listLlQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLlQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLlQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLlQuery.field_gt_updated_at = 0
        listLlQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLlLibraries(listLlQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        LlList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLlQuery.order_by = column.prop
      listLlQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLlList()
    }
    onMounted(() => {
      getLlList()
    })
    function deleteLl(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLlLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLlList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const ll: Ll = {
      id: 0,
      loss_level: 0,
      regulate_way: '',
      rated_capacity_min: 0,
      rated_capacity_max: 0,
      rated_high_vol_min: 0,
      rated_high_vol_max: 0,
      rated_low_vol_min: 0,
      rated_low_vol_max: 0,
      regulate_range_min: 0,
      regulate_range_max: 0,
      regulate_range_step: 0,
      temperature: 0,
      load_loss: 0,
      no_load_loss: 0,
      no_load_current: 0,
      short_circuit_imped: 0,
      created_at: '',
      updated_at: ''
    }
    const LlDetail = ref(ll)
    function showLlDetail(Ll: Ll) {
      showDetail.value = true
      LlDetail.value = Ll
    }
    function showLlUpdate(Ll: Ll) {
      showUpdate.value = true
      LlDetail.value = Ll
    }
    return {
      status,
      handleDownload,
      loading,
      LlList,
      LlCreate,
      created_at_range,
      updated_at_range,
      listLlQuery,
      total,
      capacity,
      highVol,
      lowVol,
      getLlList,
      sortChange,
      deleteLl,
      showCreate,
      showDetail,
      showUpdate,
      LlDetail,
      showLlUpdate,
      showLlDetail,
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
