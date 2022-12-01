<template>
  <div style="margin: 10px">
    <aside class="page-header">
      计算公式
    </aside>
    <div style="margin-bottom: 10px">
      <el-input
        v-model="listCalculationQuery.field_eq_formula_name_zh"
        placeholder="公式名"
        clearable
        style="width: 200px; margin-right: 10px"
      />
      <el-date-picker
        v-model="created_at_range"
        unlink-panels
        style="margin-left:10px"
        type="datetimerange"
        range-separator="至"
        start-placeholder="创建起始日期"
        end-placeholder="创建截止日期"
      />
      <el-date-picker
        v-model="updated_at_range"
        unlink-panels
        style="margin-left:10px"
        type="datetimerange"
        range-separator="至"
        start-placeholder="更新起始日期"
        end-placeholder="更新截止日期"
      />
      <el-button
        v-waves
        style="margin-left:20px"
        type="primary"
        icon="el-icon-search"
        @click="getCalculationList"
      >
        搜索
      </el-button>
<!--      <el-button-->
<!--        v-waves-->
<!--        type="primary"-->
<!--        icon="el-icon-edit"-->
<!--        @click="showCreate = true"-->
<!--      >-->
<!--        添加-->
<!--      </el-button>-->
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
    <Calculation-create
      v-model="showCreate"
      @get-Calculation-list="getCalculationList"
    />
    <Calculation-view
      v-model="showDetail"
      :calculation="CalculationDetail"
    />
    <Calculation-update
      v-model="showUpdate"
      :calculation="CalculationDetail"
      @get-Calculation-list="getCalculationList"
    />
    <el-table
      v-loading="loading"
      :data="CalculationList"
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
        label="公式名"
        prop="formula_name_zh"
        sortable="custom"
        align="center"
        min-width="200px"
      >
        <template #default="{row}">
          <span>{{ row.formula_name_zh }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="公式表达式"
        prop="formula_express"
        sortable="custom"
        align="center"
        min-width="450px"
      >
        <template #default="{row}">
          <math-jax
            :latex="row.formula_express"
            :block="true"
          />
        </template>
      </el-table-column>
      <el-table-column
        label="公式注释"
        prop="formula_description"
        sortable="custom"
        align="center"
        min-width="500px"
      >
        <template #default="{row}">
          <math-jax
            :latex="row.remark"
            :block="true"
          />
        </template>
      </el-table-column>
      <el-table-column
        label="创建时间"
        prop="created_at"
        sortable="custom"
        align="center"
        width="180px"
      >
        <template #default="{row}">
          <span>{{ UnixTime2HumanWithStr( row.created_at) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="更新时间"
        prop="updated_at"
        sortable="custom"
        align="center"
        width="180px"
      >
        <template #default="{row}">
          <span>{{ UnixTime2HumanWithStr(row.updated_at) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="操作"
        align="center"
        width="340px"
      >
        <template #default="{row}">
          <el-button
            size="mini"
            type="primary"
            @click="showCalculationDetail(row)"
          >
            查看详情
          </el-button>
          <el-button
            size="mini"
            type="success"
            @click="showCalculationUpdate(row)"
          >
            修改
          </el-button>
          <el-button
            size="mini"
            type="danger"
            @click="deleteCalculation(row.id)"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <div style="margin-top: 20px">
      <el-pagination
        v-show="total > 0"
        v-model:currentPage="listCalculationQuery.page"
        v-model:page-size="listCalculationQuery.limit"
        :total="total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="getCalculationList"
        @current-change="getCalculationList"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import { DeleteCalculation, GetSpecCalculation, ListCalculationQuery, Calculation } from '@/api/standard_libraries/calculation'
import CalculationCreate from './calculationCreate.vue'
import CalculationView from './calculationView.vue'
import CalculationUpdate from './calculationUpdate.vue'
import dayjs from 'dayjs'

export default defineComponent({
  components: {
    CalculationCreate,
    CalculationView,
    CalculationUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listCalculationQuery = reactive({
      field_eq_formula_name_zh: '',
      field_eq_formula_name: '',
      field_eq_formula_express: '',
      field_eq_formula_description: '',
      field_eq_formula_param: '',
      field_eq_remark: '',
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
    const CalculationList = ref(Array<Calculation>())
    const total = ref(0)
    // async function handleDownload() {
    //   loading.value = true
    //   const res = await ExportStdlwimlLibraries()
    //   if (window.navigator && window.navigator.msSaveBlob) {
    //     window.navigator.msSaveBlob(res.data, '绕线内模台账.xlsx')
    //   } else {
    //     const objectUrl = URL.createObjectURL(res.data)
    //     const link = document.createElement('a')
    //     link.style.display = 'none'
    //     link.href = objectUrl
    //     link.download = '绕线内模台账.xlsx'
    //     document.body.appendChild(link)
    //     link.click()
    //     document.body.removeChild(link)
    //     URL.revokeObjectURL(objectUrl)
    //   }
    //   loading.value = false
    // }
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function getCalculationList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listCalculationQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listCalculationQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listCalculationQuery.field_gt_created_at = 0
        listCalculationQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listCalculationQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listCalculationQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listCalculationQuery.field_gt_updated_at = 0
        listCalculationQuery.field_lt_updated_at = 0
      }
      const response = await GetSpecCalculation(listCalculationQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        CalculationList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listCalculationQuery.order_by = column.prop
      listCalculationQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getCalculationList()
    }
    onMounted(() => {
      getCalculationList()
    })
    function deleteCalculation(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteCalculation(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getCalculationList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const calculation: Calculation = {
      id: 0,
      formula_name_zh: '',
      formula_express: '',
      formula_name: '',
      formula_description: '',
      formula_param: '',
      remark: '',
      created_at: '',
      updated_at: ''

    }
    const CalculationDetail = ref(calculation)
    function showCalculationDetail(calculation: Calculation) {
      showDetail.value = true
      CalculationDetail.value = calculation
    }
    function showCalculationUpdate(calculation: Calculation) {
      showUpdate.value = true
      CalculationDetail.value = calculation
    }
    return {
      status,
      // handleDownload,
      loading,
      CalculationList,
      CalculationCreate,
      created_at_range,
      updated_at_range,
      listCalculationQuery,
      total,
      getCalculationList,
      sortChange,
      deleteCalculation,
      showCreate,
      showDetail,
      showUpdate,
      CalculationDetail,
      showCalculationUpdate,
      showCalculationDetail,
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
}
</style>
