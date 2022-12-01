<!--
 * @Description: admin 权限主页
 * @Author: ZY
 * @Date: 2021-01-15 18:44:25
 * @LastEditors: SCY
 * @LastEditTime: 2021-04-02 09:40:07
-->
<template>
  <div class="dashboard-editor-container">
    <PanelGroup
      :line-chart-data="statisticsData"
    />

    <el-row
      :gutter="32"
    >
      <el-col
        :xs="24"
        :sm="24"
        :lg="8"
      >
        <div class="chart-wrapper">
          <RadarChart />
        </div>
      </el-col>
      <el-col
        :xs="24"
        :sm="24"
        :lg="8"
      >
        <div class="chart-wrapper">
          <PieChart :chart-data="statisticsData" />
        </div>
      </el-col>
      <el-col
        :xs="24"
        :sm="24"
        :lg="8"
      >
        <div class="chart-wrapper">
          <BarChart :chart-data="statisticsData" />
        </div>
      </el-col>
    </el-row>

    <el-row :gutter="8">
      <el-col
        :xs="{span: 24}"
        :sm="{span: 24}"
        :md="{span: 24}"
        :lg="{span: 16}"
        :xl="{span: 16}"
        style="padding-right:8px;margin-bottom:30px;"
      >
        <DependsTable
          style="border-radius: 8px;"
          :design-project-list="designProjectList"
        />
      </el-col>
      <el-col
        :xs="{span: 24}"
        :sm="{span: 24}"
        :md="{span: 24}"
        :lg="{span: 8}"
        :xl="{span: 8}"
        style="padding-right:8px;margin-bottom:30px;"
      >
        <TodoList style="border-radius: 8px;" />
      </el-col>
    </el-row>

    <el-row>
      <el-col
        :xs="{span: 24}"
        :sm="{span: 24}"
        :md="{span: 24}"
        :lg="{span: 16}"
        :xl="{span: 16}"
        style="padding-right:8px;margin-bottom:30px;"
      >
        <UpdateTimeline style="border-radius: 16px;" />
      </el-col>
      <el-col
        :xs="{span: 24}"
        :sm="{span: 24}"
        :md="{span: 24}"
        :lg="{span: 8}"
        :xl="{span: 8}"
        style="padding-right:8px;margin-bottom:30px;"
      >
        <BoxCard
          style="border-radius: 8px;"
          :design-project-list="designProjectList"
        />
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, Ref, ref } from 'vue'
import PanelGroup from './components/PanelGroup.vue'
import RadarChart from './components/RadarChart.vue'
import PieChart from './components/PieChart.vue'
import BarChart from './components/BarChart.vue'
import DependsTable from './components/DependsTable.vue'
import TodoList from './components/todo-list/index.vue'
import BoxCard from './components/BoxCard.vue'
import UpdateTimeline from './components/UpdateTimeline.vue'
import { GetStatistics } from '@/api/statistics'
import type { StatisticsData } from '@/model/statistics'
import type { DesignProject } from '@/model/design'
import { GetDesignProjects } from '@/api/projects'
import { DesignProjectStatus } from '@/model/design'
import { ResourceList, ActionList } from '@/model/permission'

export default defineComponent({
  components: {

    PanelGroup,

    RadarChart,
    PieChart,
    BarChart,
    DependsTable,
    TodoList,
    BoxCard,
    UpdateTimeline
  },
  setup() {
    const emptyStatisticsData : StatisticsData = {
      design_project_check_unaccepted_count: 0,
      design_project_checked_count: 0,
      design_project_approve_unaccepted_count: 0,
      design_project_approved_count: 0,
      design_project_count: 0,
      design_project_created_count: 0,
      design_project_delivered_count: 0,
      design_project_exceeded_count: 0,
      design_project_finished_count: 0,
      design_project_imported_count: 0,
      design_project_review_unaccepted_count: 0,
      design_project_reviewed_count: 0,
      design_project_started_count: 0,
      design_project_weekly_status_count: [],
      design_project_overdue_count: 0,
      design_results_count: 0,
      libraries_count: 0,
      user_count: 0
    }
    const statisticsData : Ref<StatisticsData> = ref(emptyStatisticsData)
    const designProjectList : Ref<DesignProject[]> = ref([])
    onMounted(async () => {
      GetStatistics().then((statisticsRes) => {
        if (statisticsRes.data.code === 200) {
          statisticsData.value = statisticsRes.data.spec
        }
      })
      // 获取最近5个状态不为完成的，根据交货日期倒书序的任务书
      GetDesignProjects({
        field_ne_project_status: DesignProjectStatus.DesignProjectChecked,
        field_lk_project_name: '',
        field_lk_product_model: '',
        limit: 5,
        order: 'desc',
        order_by: 'deliver_at asc,project_status asc,created_at',
        page: 1
      }).then((designProjectRes) => {
        if (designProjectRes.data.code === 200) {
          designProjectList.value = designProjectRes.data.spec
        }
      })
    })

    return {
      statisticsData,
      designProjectList,
      ResourceList,
      ActionList
    }
  }
})
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 32px;
  background-color: #F2F7FF;
  position: relative;

  .github-corner{
    position: absolute;
    top: 0px;
    border: 0;
    right: 0;
  }

  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
    border-radius: 8px;
  }
}

@media (max-width:1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}
</style>
