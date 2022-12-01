<!--
 * @Description: 依赖包表格
 * @Author: ZY
 * @Date: 2021-01-18 11:09:25
 * @LastEditors: SCY
 * @LastEditTime: 2021-04-02 09:37:05
-->

<template>
  <el-card
    class="card"
    shadow="never"
    style="border-radius: 8px;"
  >
    <div style="margin-bottom: 15px;">
      <span>产品设计任务概要</span>
    </div>
    <el-empty
      v-if="projectList.length === 0"
      :image-size="200"
    />
    <div v-else>
      <div class="bottom-btn">
        <router-link
          v-for="item in projectList"
          :key="item"
          :to="`/design/design_projects/?code=${item.serial_code}`"
        >
          <el-button type="success">
            {{ item.product_model }}
          </el-button>
        </router-link>
      </div>
      <table
        class="table"
        style="alignment: left"
      >
        <tr style="text-align: left">
          <td align="left">
            项目/产品名称
          </td>
          <td style="text-align: left">
            产品型号
          </td>
          <td style="text-align: left">
            设计/校核
          </td>
          <td style="text-align: left">
            产品交货日期
          </td>
          <td

            style="text-align: left"
          >
            产品当期状态
          </td>
        </tr>
        <tr
          v-for="item in projectList"
          :key="item"
        >
          <td style="text-align: left">
            {{ item.project_name }}/{{ item.product_name }}
          </td>
          <td style="text-align: left">
            {{ item.product_model }}
          </td>
          <td style="text-align: left">
            {{ item.designer_user.user_name_zh }}/{{ item.check_user.user_name_zh }}
          </td>
          <td style="text-align: left">
            {{ UnixTime2HumanWithUnix(item.deliver_at) }}
          </td>
          <td
            v-if="statusText[item.project_status]==='设计任务未开始'"
            style="text-align: left"
          >
            <el-tag type="danger">
              {{ statusText[item.project_status] }}
            </el-tag>
          </td>
          <td
            v-else-if="statusText[item.project_status]==='设计任务被驳回'"
            style="text-align: left"
          >
            <el-tag type="warning">
              {{ statusText[item.project_status] }}
            </el-tag>
          </td>
          <td
            v-else
            style="text-align: left"
          >
            <el-tag>
              {{ statusText[item.project_status] }}
            </el-tag>
          </td>
        </tr>
      </table>
    </div>
  </el-card>
</template>

<script lang="ts">
import { computed, defineComponent, reactive, toRefs } from 'vue'
import type { DesignProject } from '@/model/design'
import { UnixTime2HumanWithUnix } from '@/utils/timeutils'

export default defineComponent({
  props: {
    designProjectList: Array
  },
  setup(props) {
    const projectList = computed(() => {
      console.log(props.designProjectList)
      return props.designProjectList as DesignProject[]
    })
    const statusText = [
      '',
      '设计任务未开始',
      '设计任务未开始',
      '设计任务已开始',
      '设计任务计算完成',
      '设计任务校核不通过',
      '设计任务待校核',
      '设计任务审核不通过',
      '设计任务审核通过',
      '设计任务审核不通过',
      '设计任务已完成'
    ]
    return {
      UnixTime2HumanWithUnix,
      projectList,
      statusText
    }
  }
})
</script>

<style lang="scss" scoped>
.table {
  width: 100%;
  color: #666;
  border-collapse: collapse;
  background-color: #fff;

  td {
    position: relative;
    min-height: 20px;
    padding: 9px 15px;
    font-size: 14px;
    line-height: 20px;
    border: 1px solid #e6e6e6;

    &:nth-child(odd) {
      width: 20%;
      text-align: right;
      background-color: #f7f7f7;
    }
  }
}

.bottom-btn {
  button {
    margin: 0 10px 15px 0;
  }
}
</style>
