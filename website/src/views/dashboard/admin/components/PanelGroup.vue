<!--
 * @Description:
 * @Author: ZY
 * @Date: 2021-01-18 10:48:43
 * @LastEditors: SCY
 * @LastEditTime: 2021-04-02 10:00:26
-->
<template>
  <el-row
    :gutter="20"
    class="panel-group"
  >
    <el-col

      :span="4"
      class="card-panel-col"
    >
      <div
        class="card-panel"
      >
        <div class="card-panel-icon-wrapper icon-people">
          <img
            :src="designImg"
          >
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            设计任务数
          </div>
          <CountTo
            ref="myCount"
            :start-val="0"
            :end-val="myCount.design_project_count"
            :duration="26"
            class="card-panel-num"
          />
        </div>
      </div>
    </el-col>
    <el-col

      :span="4"
      class="card-panel-col"
    >
      <div
        class="card-panel"
      >
        <div class="card-panel-icon-wrapper icon-message">
          <img
            :src="librariesImg"
          >
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            标准库|表数
          </div>
          <CountTo
            :start-val="0"
            :end-val="myCount.libraries_count"
            :duration="3000"
            class="card-panel-num"
          />
        </div>
      </div>
    </el-col>
    <el-col
      :span="4"
      class="card-panel-col"
    >
      <div
        class="card-panel"
      >
        <div class="card-panel-icon-wrapper icon-money">
          <img
            :src="designResultImg"
          >
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            设计结果数
          </div>
          <CountTo
            :start-val="0"
            :end-val="myCount.design_results_count"
            :duration="3200"
            class="card-panel-num"
          />
        </div>
      </div>
    </el-col>
    <el-col

      :span="4"
      class="card-panel-col"
    >
      <div
        class="card-panel"
      >
        <div class="card-panel-icon-wrapper icon-shopping">
          <img
            :src="usersImg"
          >
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            系统使用人数
          </div>
          <CountTo
            :start-val="0"
            :end-val="myCount.user_count"
            :duration="3600"
            class="card-panel-num"
          />
        </div>
      </div>
    </el-col>
    <el-col

      :span="4"
      class="card-panel-col"
    >
      <div
        class="card-panel"
      >
        <div class="card-panel-icon-wrapper icon-people">
          <img
            :src="deliverCountImg"
          >
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            已交货产品数
          </div>
          <CountTo
            :start-val="0"
            :end-val="myCount.design_project_delivered_count"
            :duration="26"
            class="card-panel-num"
          />
        </div>
      </div>
    </el-col>
    <el-col

      :span="4"
      class="card-panel-col"
    >
      <div
        class="card-panel"
      >
        <div class="card-panel-icon-wrapper icon-message">
          <img
            :src="delayImg"
          >
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            设计延期数
          </div>
          <CountTo
            :start-val="0"
            :end-val="myCount.design_project_exceeded_count"
            :duration="3000"
            class="card-panel-num"
          />
        </div>
      </div>
    </el-col>
  </el-row>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, ref } from 'vue'
import { CountTo } from 'vue3-count-to'
import designImg from '@/assets/images/home/design.png'
import librariesImg from '@/assets/images/home/libraries.png'
import designResultImg from '@/assets/images/home/design-result.png'
import usersImg from '@/assets/images/home/users.png'
import delayImg from '@/assets/images/home/delay.png'
import deliverCountImg from '@/assets/images/home/deliver-count.png'
export default defineComponent({
  components: {
    CountTo
  },
  props: {
    lineChartData: Object
  },
  emits: [],
  setup(props, { emit }) {
    const myCount = computed(() => {
      return props.lineChartData
    })
    onMounted(() => {
      console.log((myCount.value as any).value)
    })

    return {
      myCount,
      designImg,
      librariesImg,
      designResultImg,
      usersImg,
      delayImg,
      deliverCountImg
    }
  }
})
</script>

<style lang="scss" scoped>
.panel-group {
  margin-top: 18px;

  .card-panel-col {
    margin-bottom: 32px;
  }

  .card-panel {
    height: 108px;
    cursor: pointer;
    font-size: 12px;
    position: relative;
    overflow: hidden;
    color: #666;
    background: #fff;
    box-shadow: 4px 4px 40px rgba(0, 0, 0, .05);
    border-color: rgba(0, 0, 0, .05);
    height: 124px;
    background: #FFFFFF;
    border-radius: 10px;
    img{
      width:60px;
      height: 60px;
      display: inline-block;
    }
    .icon-people {
      color: #40c9c6;
    }

    .icon-message {
      color: #36a3f7;
    }

    .icon-money {
      color: #f4516c;
    }

    .icon-shopping {
      color: #34bfa3
    }

    .card-panel-icon-wrapper {
      float: left;
      margin: 14px 0 0 14px;
      width: 40%;
      padding: 16px;
      transition: all 0.38s ease-out;
      border-radius: 6px;
    }

    .card-panel-icon {
      float: left;
      font-size: 48px;
    }

    .card-panel-description {
      float: right;
      font-weight: bold;
      width: 49%;
      margin: 26px 0;

      .card-panel-text {
        line-height: 18px;
        color: rgba(0, 0, 0, 0.45);
        font-size: 16px;
        margin-bottom: 12px;
      }

      .card-panel-num {
        font-size: 20px;
      }
    }
  }
}

@media (max-width:1512px) {
  .card-panel-icon-wrapper {
    margin: 14px 0 0 0 !important;
  }
}
@media (max-width:550px) {
  .card-panel-description {
    display: none;
  }

  .card-panel-icon-wrapper {
    float: none !important;
    width: 100%;
    height: 100%;
    margin: 0 !important;

    svg {
      display: block;
      margin: 14px auto !important;
      float: none !important;
    }
  }
}
</style>
