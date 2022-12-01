<!--
 * @Description: 首页雷达图
 * @Author: ZY
 * @Date: 2021-01-18 11:23:22
 * @LastEditors: SCY
 * @LastEditTime: 2021-04-06 14:50:31
-->
<template>
  <div
    id="homeRadarCharts"
    :class="className"
    :style="{height: height, width: width}"
  />
</template>
<script lang="ts">
import { defineComponent, onActivated, onBeforeUnmount, onDeactivated, onMounted, nextTick } from 'vue'
import resize from '@/components/charts/mixins/resize'
import { init } from 'echarts'

const animationDuration = 3000
export default defineComponent({
  props: {
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
  setup() {
    const {
      mounted,
      chart,
      beforeDestroy,
      activated,
      deactivated
    } = resize()

    const initChart = () => {
      const radarChart = init(document.getElementById('homeRadarCharts') as HTMLDivElement, 'macarons')
      radarChart.setOption({
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'shadow'
          }
        },
        radar: {
          radius: '66%',
          center: ['50%', '42%'],
          splitNumber: 8,

          splitArea: {
            areaStyle: {
              color: '#C5CEE6',
              opacity: 1,
              shadowBlur: 45,
              shadowColor: 'rgba(0,0,0,.1)',
              shadowOffsetX: 0,
              shadowOffsetY: 15
            }
          },
          indicator: [
            { name: '额定容量', max: 10000 },
            { name: '额定频率', max: 20000 },
            { name: '额定高压', max: 20000 },
            { name: '相数', max: 20000 },
            { name: '额定低压', max: 20000 },
            { name: '感应高压', max: 20000 }
          ]
        },
        legend: {
          left: 'center',
          bottom: '10',
          data: ['温升指标', '综合成本指标', '损耗指标']
        },
        series: [{
          color: ['#435EBE', '#5DDAB4', '#57CAEB'],
          type: 'radar',
          symbolSize: 0,
          areaStyle: {
            shadowBlur: 13,
            shadowColor: 'rgba(0,0,0,.2)',
            shadowOffsetX: 0,
            shadowOffsetY: 10,
            opacity: 1
          },
          data: [
            {
              value: [5000, 7000, 12000, 11000, 15000, 14000],
              name: '温升指标'
            },
            {
              value: [4000, 9000, 15000, 15000, 13000, 11000],
              name: '综合成本指标'
            },
            {
              value: [5500, 11000, 12000, 15000, 12000, 12000],
              name: '损耗指标'
            }
          ],
          animationDuration: animationDuration
        }]
      } as any)
      chart.value = radarChart
    }

    onMounted(() => {
      mounted()
      nextTick(() => {
        initChart()
      })
    })

    onBeforeUnmount(() => {
      beforeDestroy()
      if (!chart.value) {
        return
      }
      chart.value.dispose()
      chart.value = null
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
