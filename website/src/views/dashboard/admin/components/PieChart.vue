<!--
 * @Description: 首页饼图组件
 * @Author: ZY
 * @Date: 2021-01-18 11:08:53
 * @LastEditors: ZY
 * @LastEditTime: 2021-01-26 10:19:45
-->
<template>
  <div
    id="homePieCharts"
    :class="className"
    :style="{height: height, width: width}"
  />
</template>

<script lang="ts">
import {
  defineComponent,
  onActivated,
  onBeforeUnmount,
  onDeactivated,
  onMounted,
  nextTick,
  computed,
  ComputedRef, watch
} from 'vue'
import resize from '@/components/charts/mixins/resize'
import { init, EChartsOption } from 'echarts'
import { StatisticsData } from '@/model/statistics'
export default defineComponent({
  props: {
    chartData: Object,
    className: {
      type: String,
      default: 'chart'
    },
    width: {
      type: String,
      default: '100%',
      required: true
    },
    height: {
      type: String,
      default: '300px',
      required: true
    }
  },
  setup(props) {
    const {
      mounted,
      chart,
      beforeDestroy,
      activated,
      deactivated
    } = resize()
    const chartData : ComputedRef<StatisticsData> = computed(() => {
      return props.chartData as StatisticsData
    })

    const initChart = () => {
      const pieChart = init(document.getElementById('homePieCharts') as HTMLDivElement, 'macarons')
      pieChart.setOption({
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b} : {c} ({d}%)'
        },
        legend: {
          left: 'center',
          bottom: '10',
          data: ['设计未开始', '设计进行中', '设计已逾期', '设计已完成']
        },
        series: [
          {
            color: ['#435EBE', '#5DDAB4', '#9694FF', '#FF7976', '#57CAEB'],
            name: '设计任务统计状态总览',
            type: 'pie',
            roseType: 'radius',
            radius: [15, 95],
            center: ['50%', '38%'],
            data: [
              { value: chartData.value.design_project_imported_count, name: '设计未开始' },
              {
                value: chartData.value.design_project_started_count +
              chartData.value.design_project_created_count +
              chartData.value.design_project_finished_count +
              chartData.value.design_project_reviewed_count +
              chartData.value.design_project_approved_count,
                name: '设计进行中'
              },
              {
                value: chartData.value.design_project_overdue_count,
                name: '设计已逾期'
              },
              { value: chartData.value.design_project_checked_count, name: '设计已完成' }
            ],
            animationEasing: 'cubicInOut',
            animationDuration: 2600
          }
        ]
      } as EChartsOption)
      chart.value = pieChart
    }
    watch(() => props.chartData, () => { initChart() })
    onMounted(() => {
      mounted()
      nextTick(() => {
        initChart()
      })
    })

    onBeforeUnmount(() => {
      // if (!chart.value) {
      //   return
      // }
      // chart.value.dispose()
      // chart.value = null
      beforeDestroy()
    })

    onActivated(() => {
      activated()
    })

    onDeactivated(() => {
      deactivated()
    })

    return {

    }
  }
})
</script>
