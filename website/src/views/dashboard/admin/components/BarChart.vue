<!--
 * @Description: 首页柱形图小组件
 * @Author: ZY
 * @Date: 2021-01-15 18:48:40
 * @LastEditors: ZY
 * @LastEditTime: 2021-01-26 09:57:33
-->

<template>
  <div
    id="homebarcharts"
    :class="className"
    :style="{height: height, width: width}"
  />
</template>

<script lang="ts">
import { defineComponent, onActivated, onBeforeUnmount, onDeactivated, onMounted, nextTick, watch } from 'vue'
import resize from '@/components/charts/mixins/resize'
import { init, EChartsOption } from 'echarts'
import { GetStatistics } from '@/api/statistics'
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

    const animationDuration = 6000
    const initChart = () => {
      const barChart = init(document.getElementById('homebarcharts') as HTMLDivElement, 'macarons')
      const count = []
      for (const weekDate of (props.chartData as StatisticsData).design_project_weekly_status_count) {
        count.push(weekDate.design_project_count)
      }
      barChart.setOption({
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'shadow'
          }
        },
        grid: {
          top: 10,
          left: '2%',
          right: '2%',
          bottom: '3%',
          containLabel: true
        },
        xAxis: [{
          type: 'category',
          data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
          axisTick: {
            alignWithLabel: true
          }
        }],
        yAxis: [{
          type: 'value',
          minInterval: 1,
          axisTick: {
            show: false
          }
        }],
        series: [
          {
            name: '新增任务书数',
            type: 'bar',
            stack: 'vistors',
            // barWidth: '60%',
            data: count,
            animationDuration
          }
        ]
      } as EChartsOption)
      chart.value = barChart
    }

    watch(() => props.chartData, () => {
      initChart()
    })

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
