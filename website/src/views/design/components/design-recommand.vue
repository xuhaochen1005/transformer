<template>
  <el-row
    v-show="show"
    style="margin-top: 1%"
  >
    <div class="filter-container">
      <el-input
        v-model="recommendDesignTaskQuery.rated_capacity_range_max"
        placeholder="额定容量上限"
        clearable
        style="width: 200px"
        class="filter-item"
      />
      <el-input
        v-model="recommendDesignTaskQuery.rated_capacity_range_min"
        placeholder="额定容量下限"
        clearable
        class="filter-item"
        style="width: 200px"
      />
      <el-input
        v-model="recommendDesignTaskQuery.rated_low_range_max"
        placeholder="额定低压上限"
        clearable
        style="width: 200px"
        class="filter-item"
      />
      <el-input
        v-model="recommendDesignTaskQuery.rated_low_range_min"
        placeholder="额定低压下限"
        clearable
        class="filter-item"
        style="width: 200px"
      />

      <el-button
        v-waves
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="getRecommendTaskList"
      >
        搜索
      </el-button>
      <el-button
        class="filter-item"
        style="margin-left: 10px"
        type="primary"
        icon="el-icon-edit"
        @click="show = false"
      >
        关闭推荐设计
      </el-button>
    </div>
    <el-table
      key="result-like"
      v-loading="recommendLoading"
      :data="recommendTaskList"
      border
      fit
      highlight-current-row
      style="width: 100%"
    >
      <el-table-column
        label="相似度"
        sortable="custom"
        align="center"
        min-width="141px"
      >
        <template #default="{row}">
          <span>{{ getDesignTaskSimilarity(row) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="项目名称"
        prop="name"
        sortable="custom"
        width="200px"
        align="center"
      >
        <template #default="{row}">
          <span>{{ row.design_project.project_name }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="型号"
        prop="name"
        sortable="custom"
        width="200px"
        align="center"
      >
        <template #default="{row}">
          <span>{{ row.design_project.product_model }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="额定容量(kvA)"
        sortable="custom"
        align="center"
        width="130px"
      >
        <template #default="{row}">
          <span>{{ row.input.rated_capacity }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="额定高压(kV)"
        sortable="custom"
        align="center"
        width="130px"
      >
        <template #default="{row}">
          <span>{{ row.input.rated_high_vol }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="额定低压(kV)"
        sortable="custom"
        align="center"
        width="130px"
      >
        <template #default="{row}">
          <span>{{ row.input.rated_low_vol }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="阻抗电压(kV)"
        sortable="custom"
        align="center"
        width="130px"
      >
        <template #default="{row}">
          <span>{{ row.input.limit_impedance_vol }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="联结组别"
        sortable="custom"
        align="center"
        width="150px"
      >
        <template #default="{row}">
          <span>{{ row.input.connect_type }}</span>
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
          <span>{{ UnixTime2HumanWithStr(row.created_at) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="备注"
        prop="comment"
        sortable="custom"
        align="center"
        width="200px"
      >
        <template #default="{row}">
          <span>{{ row.comment }}</span>
        </template>
      </el-table-column>

      <el-table-column
        label="操作"
        fixed="right"
        align="center"
        min-width="273px"
        class-name="fixed-width"
      >
        <template #default="{row}">
          <el-button
            type="primary"
            size="mini"
            @click="handleDetail(row)"
          >
            查看详情
          </el-button>
          <el-button
            size="mini"
            type="warning"
            @click="handleUpdate(row)"
          >
            使用设计参数
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      v-if="recommendTotal > 0"
      style="margin-top: 1%;"
      :total="recommendTotal"
      :page-sizes="[10, 20, 50, 100]"
      layout="total, sizes, prev, pager, next, jumper"
      @size-change="getRecommendTaskList"
      @current-change="getRecommendTaskList"
    />
    <el-dialog
      v-model="showDetail"
      custom-class="detail-dialog"
      title="设计详情"
      width="80%"
    >
      <design-results-table
        show-design-results-mode="view"
        :current-design-results="recommendRow.final_design_results"
      />
      <template #footer>
        <el-button
          type="primary"
          @click="handleUpdate(recommendRow)"
        >
          使用设计参数
        </el-button>
        <el-button @click="showDetail = false">
          关闭
        </el-button>
      </template>
    </el-dialog>
  </el-row>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, reactive, Ref, ref, watch, watchEffect } from 'vue'
import { DecodeDesignTask, DesignProject, DesignTask, PullType, SeamType } from '@/model/design'
import { getRecommendDesignTaskList } from '@/api/design'
import { decodeObjectBytes, stringify } from '@/utils/jsonutils'
import { ElMessage } from 'element-plus'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import DesignResultsTable from '@/views/design/components/design-results-table.vue'
const defaultTableRow : DesignTask = {
  approve: 0,
  approve_by_creator: false,
  approve_node_instance: undefined,
  result_progress: 0,
  comment: '',
  created_at: '',
  creator: 0,
  deleted_at: '',
  design_project: {
    check_by: 0,
    check_user: undefined,
    checked_at: 0,
    comment: '',
    created_at: '',
    creator: 0,
    creator_user: undefined,
    deleted_at: '',
    deliver_at: 0,
    design_at: 0,
    design_comment: '',
    design_finish_at: 0,
    designer: 0,
    designer_signed: 0,
    designer_user: undefined,
    drafting_at: 0,
    drafting_by: 0,
    drafting_user: undefined,
    drawing_at: 0,
    id: 0,
    need_master_approve: 0,
    order_at: 0,
    product_altitude: 0,
    product_code: '',
    product_connection_group: '',
    product_cooling_method: '',
    product_frequency: 0,
    product_frequency_special: '',
    product_induct_high_vol: 0,
    product_industry_freq_hold_vol: '',
    product_insulation_high_ac: 0,
    product_insulation_high_li: 0,
    product_insulation_level: '',
    product_insulation_level_special: '',
    product_insulation_low_ac: 0,
    product_insulation_low_li: 0,
    product_ip_level: '',
    product_load_loss: 0,
    product_model: '',
    product_name: '',
    product_noload_loss: 0,
    product_phrase: '',
    product_rated_capacity: 0,
    product_rated_v_high: 0,
    product_rated_v_low: 0,
    product_short_circuit_resistance: 0,
    product_spec_shock_vol: '',
    product_tap_range: '',
    product_tap_range_special: '',
    product_temp_limit: 0,
    product_total_loss: 0,
    product_usage: '',
    product_wind_type: '',
    product_wire_material: '',
    product_wire_shape: '',
    project_name: '',
    project_status: 0,
    review_at: 0,
    reviewer: 0,
    reviewer_signed: 0,
    reviewer_user: undefined,
    serial_code: '',
    tech_requirments: '',
    updated_at: ''
  } as DesignProject,
  design_project_id: 0,
  design_results: '',
  designer: 0,
  final_design_results: [],
  id: 0,
  input: '',
  name: '',
  reviewer: 0,
  task_status: 0,
  updated_at: ''
}
export default defineComponent({
  components: {
    DesignResultsTable
  },

  props: {
    modelValue: Boolean,
    handleValue: {
      required: true,
      type: Object as PropType<DecodeDesignTask>,
      default: () => { return JSON.parse(JSON.stringify(defaultTableRow)) }
    }
  },
  emits: ['update:modelValue', 'update:handleValue'],
  setup(props, context) {
    // 相似设计
    const show = computed({
      get: () => props.modelValue,
      set: (value) => {
        context.emit('update:modelValue', value)
      }
    })
    const tempTableRow = computed({
      get: () => props.handleValue as DecodeDesignTask,
      set: (value) => {
        context.emit('update:handleValue', value)
      }
    })
    const recommendTotal = ref(0)
    const recommendTaskList : Ref<DecodeDesignTask[]> = ref([])
    const recommendLoading = ref(false)
    const downloadLoading = ref(false)
    const recommendDesignTaskQuery = reactive({
      rated_capacity_range_max: undefined,
      rated_capacity_range_min: undefined,
      rated_low_range_max: undefined,
      rated_low_range_min: undefined,
      'field_eq_JSON_EXTRACT(input,"$.phase_number")': ['1', '3'],
      'field_ge_JSON_EXTRACT(input,"$.rated_capacity")': 0,
      'field_le_JSON_EXTRACT(input,"$.rated_capacity")': 0,
      'field_ge_JSON_EXTRACT(input,"$.rated_high_vol")': 9,
      'field_le_JSON_EXTRACT(input,"$.rated_high_vol")': 11,
      'field_ge_JSON_EXTRACT(input,"$.rated_low_vol")': 0,
      'field_le_JSON_EXTRACT(input,"$.rated_low_vol")': 0,
      'field_re_JSON_EXTRACT(input,"$.connect_type")': '',
      'field_eq_JSON_EXTRACT(input,"$.limit_impedance_vol")': [4, 6, 8],
      field_re_name: '',
      page: 1,
      limit: 10,
      order: '',
      order_by: ''
    })
    watchEffect(() => {
      // 默认 误差50
      recommendDesignTaskQuery['field_le_JSON_EXTRACT(input,"$.rated_capacity")'] = recommendDesignTaskQuery.rated_capacity_range_max ?? (tempTableRow.value.input.rated_capacity * 100 + 50 * 100) / 100
      recommendDesignTaskQuery['field_ge_JSON_EXTRACT(input,"$.rated_capacity")'] = recommendDesignTaskQuery.rated_capacity_range_min ?? (tempTableRow.value.input.rated_capacity * 100 - 50 * 100) / 100
      // 额定高压 默认误差 10%
      recommendDesignTaskQuery['field_le_JSON_EXTRACT(input,"$.rated_high_vol")'] = (tempTableRow.value.input.rated_high_vol * 100 + tempTableRow.value.input.rated_high_vol * 100 * 0.1) / 100
      recommendDesignTaskQuery['field_ge_JSON_EXTRACT(input,"$.rated_high_vol")'] = (tempTableRow.value.input.rated_high_vol * 100 - tempTableRow.value.input.rated_high_vol * 100 * 0.1) / 100
      // 额定低压 误差 1
      recommendDesignTaskQuery['field_ge_JSON_EXTRACT(input,"$.rated_low_vol")'] = recommendDesignTaskQuery.rated_low_range_min ?? (tempTableRow.value.input.rated_low_vol * 100 - 1 * 100) / 100
      recommendDesignTaskQuery['field_le_JSON_EXTRACT(input,"$.rated_low_vol")'] = recommendDesignTaskQuery.rated_low_range_max ?? (tempTableRow.value.input.rated_low_vol * 100 + 1 * 100) / 100

      recommendDesignTaskQuery['field_re_JSON_EXTRACT(input,"$.connect_type")'] = tempTableRow.value!.input.connect_type.substr(0, 2)
      // TODO: regex (ZL|ZB)?([DS])(C)(Z?)(B?)(10|11|12|13|15|16)-(.*?)\/(.*?)\/(.*) mysql 查找取6
      recommendDesignTaskQuery.field_re_name = '(ZL|ZB)?([DS])(C)(Z?)(B?)(10|11|12|13|15|16)-(.*?)/(.*?)/(.*)'
    })
    async function getRecommendTaskList() {
      recommendLoading.value = true
      const response = await getRecommendDesignTaskList(recommendDesignTaskQuery)
      if (response.data.code === 200) {
        decodeObjectBytes(response.data.spec)
        recommendTaskList.value = response.data.spec as DecodeDesignTask[]
        recommendTotal.value = response.data.total
      }
      recommendLoading.value = false
    }
    watch(show, (v) => {
      if (!v) { return }
      getRecommendTaskList()
    })
    const getDesignTaskSimilarity = (row: DecodeDesignTask) => {
      let value = 60
      if (row.input.rated_capacity === tempTableRow.value.input.rated_capacity) {
        value += 5
      }
      if (row.input.phase_number === tempTableRow.value.input.phase_number) {
        value += 5
      }
      if (row.input.rated_high_vol === tempTableRow.value.input.rated_high_vol) {
        value += 5
      }
      if (row.input.rated_low_vol === tempTableRow.value.input.rated_low_vol) {
        value += 5
      }
      if (row.input.connect_type === tempTableRow.value.input.connect_type) {
        value += 5
      }
      if (row.input.limit_impedance_vol === tempTableRow.value.input.limit_impedance_vol) {
        value += 5
      }
      // 损耗水平
      const nameLike = row.name.match(/(ZL|ZB)?([DS])(C)(Z?)(B?)(10|11|12|13|15|16)/)
      if (nameLike) {
        if (nameLike[6]! === tempTableRow.value.name) {
          value += 5
        }
      }
      return `${value}%`
    }
    const showDetail = ref(false)

    const handleUpdate = (row: DecodeDesignTask) => {
      showDetail.value = false
      tempTableRow.value.input = JSON.parse(JSON.stringify(row.input))
      ElMessage.success('成功使用相识设计的参数')
    }
    const activeCollapseNames = ref(['1'])
    activeCollapseNames.value = ['1']
    function handleCollapseChange(value: any) {
      console.log('collapse change', activeCollapseNames, value)
    }

    const recommendRow : Ref<DesignTask> = ref(JSON.parse(JSON.stringify(defaultTableRow)) as DesignTask)
    const handleDetail = (row: DesignTask) => {
      const detail = JSON.parse(JSON.stringify(row)) as DesignTask
      detail.name = detail.design_project!.product_model
      detail.final_design_results.forEach(item => {
        const temp = JSON.parse(JSON.stringify(detail)) as DesignTask
        temp.final_design_results = []
        item.design_task = temp
      })
      recommendRow.value = detail
      showDetail.value = true
    }
    const pullTypeOptions = [
      { label: '有拉板', value: PullType.WithPull },
      { label: '无拉板', value: PullType.WithoutPull }
    ]
    const seamTypeOptions = [
      { label: '全斜', value: SeamType.FullOblique },
      { label: '半斜', value: SeamType.HalfOblique }
    ]
    return {
      UnixTime2HumanWithStr,
      show,
      recommendTotal,
      recommendTaskList,
      getRecommendTaskList,
      recommendLoading,
      recommendDesignTaskQuery,
      getDesignTaskSimilarity,
      handleUpdate,
      showDetail,
      activeCollapseNames,
      handleCollapseChange,
      recommendRow,
      handleDetail,
      pullTypeOptions,
      seamTypeOptions,
      downloadLoading
    }
  }
})
</script>

<style scoped lang="scss">
.filter-container {
.container-row:first-child {
  margin-bottom: 10px;
}
.filter-item:not(button) {
  margin-right: 10px;
}
}
:deep() {
  .detail-dialog {
    .product-model {
      label {
        border: 1px  ;
        line-height: 38px;
        width: 157px;
        text-align: center;
      }
      .el-input,.el-input-number {
        width: 100% !important;
        input {
          text-align: left !important;
          width: 100% !important;
        }
      }
    }
    label {
      width: 90px;
      text-align: center;
      font-size: 14px;
      margin: 0;
    }
    .is-in-pagination {
      .el-input {
        width: 50px !important;
        padding: 0 3px !important;
      }
    }
    .el-input:not(.el-select .el-input),.el-input-number{
      width: 130px;
    }
    /*.el-input-number.is-without-controls .el-input__inner {
      width: 550px;
      margin-outside: 1px;
      !*line-height: 30px;
      height: 28px;*!
    }*/
    .el-input-number {
      .el-input {
        width: 100% !important;
      }
    }
    .el-input{
      width: 100%;
      input {
        padding: 0 27px;
        text-align: center;
        color: #409eff;
      }
    }

    .data-unit {
      width: 35px;
      text-align: center;
      display: inline-block;
      font-size: 12px;
      font-weight: 700;
    }
    .value-to {
      width: 15px;
      line-height: 35px;
      text-align: center;
      font-size: 24px;
      font-weight: 700;
      margin-left: 22px;
    }
    .collapse-title{
      justify-content: center; /* 水平居中 */
      text-align: right;
      alignment: center;
      flex: 0 2 50%;
    }
    .multiple-select {
      .el-input {
        width: 100% !important;
      }
    }
    //disabled
    .is-disabled {
      .el-input__inner,.el-textarea__inner,.el-radio__label {
        background: #fff;
        color: #000 !important;
      }
      .is-checked {
        .el-radio__inner {
          background: $__color-primary !important;
          color: #fff !important;
          border-color: $__color-primary !important;
          &::after {
            background: #fff !important;
          }
        }
      }
    }
  }

}
</style>
