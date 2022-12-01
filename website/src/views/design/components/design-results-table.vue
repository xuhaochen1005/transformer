<template>
  <div>
    <el-table
      ref="singleTableRef"
      :highlight-current-row="showDesignResultsMode === 'create'"
      border
      :default-sort="{prop: `performance_result['总成本']`, order: 'ascending'}"
      :data="currentDesignResults"
      fit
      @current-change="handleSelectionChange"
    >
      <el-table-column
        label="ID"
        align="center"
        width="50px"
      >
        <template #default="{$index}">
          <span>{{ $index + 1 }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="总成本"
        :sortable="showDesignResultsMode === 'create'"
        align="center"
        prop="total_cost"
        min-width="90px"
      >
        <template #default="{row}">
          <span>{{ row.cost_result['总成本']?.toFixed(2) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="总损耗"
        prop="performance_result['总损耗']"
        align="center"
        :sortable="showDesignResultsMode === 'create'"
        min-width="90px"
      >
        <template #default="{row}">
          <span>{{ row.performance_result['总损耗']?.toFixed(2) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="线圈损耗"
        prop="performance_result['线圈损耗']"
        align="center"
        :sortable="showDesignResultsMode === 'create'"
        min-width="103px"
      >
        <template #default="{row}">
          <span>{{ row.performance_result['线圈损耗']?.toFixed(2) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="引线损耗"
        prop="performance_result['引线损耗']"
        align="center"
        :sortable="showDesignResultsMode === 'create'"
        min-width="103px"
      >
        <template #default="{row}">
          <span>{{ row.performance_result['引线损耗']?.toFixed(2) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="负载损耗"
        prop="performance_result['负载损耗']"
        align="center"
        :sortable="showDesignResultsMode === 'create'"
        min-width="103px"
      >
        <template #default="{row}">
          <span>{{ row.performance_result['负载损耗']?.toFixed(2) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="空载损耗"
        prop="performance_result['空载损耗']"
        align="center"
        :sortable="showDesignResultsMode === 'create'"
        min-width="103px"
      >
        <template #default="{row}">
          <span>{{ row.performance_result['空载损耗']?.toFixed(2) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="线圈损耗"
        prop="performance_result['线圈损耗']"
        align="center"
        :sortable="showDesignResultsMode === 'create'"
        min-width="103px"
      >
        <template #default="{row}">
          <span>{{ row.performance_result['线圈损耗']?.toFixed(2) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="铁心温升"
        prop="performance_result['铁心温升']"
        align="center"
        :sortable="showDesignResultsMode === 'create'"
        min-width="103px"
      >
        <template #default="{row}">
          <span>{{ row.performance_result['铁心温升']?.toFixed(2) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="低压温升"
        prop="performance_result['低压温升']"
        align="center"
        :sortable="showDesignResultsMode === 'create'"
        min-width="103px"
      >
        <template #default="{row}">
          <span>{{ row.performance_result['低压温升']?.toFixed(2) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="高压温升"
        prop="performance_result['高压温升']"
        align="center"
        :sortable="showDesignResultsMode === 'create'"
        min-width="103px"
      >
        <template #default="{row}">
          <span>{{ row.performance_result['高压温升']?.toFixed(2) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="空载电流"
        prop="performance_result['空载电流']"
        align="center"
        :sortable="showDesignResultsMode === 'create'"
        min-width="103px"
      >
        <template #default="{row}">
          <span>{{ row.performance_result['空载电流']?.toFixed(2) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="阻抗电压"
        prop="performance_result['阻抗电压']"
        align="center"
        :sortable="showDesignResultsMode === 'create'"
        min-width="103px"
      >
        <template #default="{row}">
          <span>{{ row.performance_result['阻抗电压']?.toFixed(2) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="噪音预测"
        prop="performance_result['噪音预测']"
        align="center"
        :sortable="showDesignResultsMode === 'create'"
        min-width="103px"
      >
        <template #default="{row}">
          <span>{{ row.performance_result['噪音预测']?.toFixed(2) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="操作"
        fixed="right"
        align="center"
      >
        <template #default="{row}">
          <el-button
            type="text"
            size="small"
            @click="currentDesignResultDetail = row;showDesignResultDetail = true"
          >
            查看详情
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog
      v-model="showDesignResultDetail"
      custom-class="dialog-result-detail"
      width="80%"
    >
      <design-results-detail
        :handle-value="currentDesignResultDetail"
      />
    </el-dialog>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType, ref, Ref, watch } from 'vue'
import { DesignResults } from '@/model/design'
import DesignResultsDetail from '@/views/design/design-results-detail.vue'
import { ElTable } from 'element-plus'

export default defineComponent({
  components: { DesignResultsDetail },
  props: {
    currentDesignResults: {
      type: Array as PropType<DesignResults[]>,
      default: () => []
    },
    showDesignResultsMode: {
      type: String as PropType<'create' | 'view'>,
      default: 'create'
    }
  },
  emits: ['handleSelectionChange'],
  setup(props, { emit }) {
    const multipleTableRef = ref<InstanceType<typeof ElTable>>()
    const handleSelectionChange = (result: DesignResults | undefined) => {
      if (result) {
        emit('handleSelectionChange', [result])
      }
    }
    const showDesignResultDetail = ref(false)
    const currentDesignResultDetail : Ref<DesignResults | null> = ref(null)
    return {
      handleSelectionChange,
      showDesignResultDetail,
      currentDesignResultDetail
    }
  }
})
</script>

<style scoped lang="scss">
</style>
