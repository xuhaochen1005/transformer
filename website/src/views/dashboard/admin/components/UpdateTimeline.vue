<template>
  <el-card
    v-loading="listLoading"
    style="height: 100%;position: relative;"
    class="card"
    shadow="never"
  >
    <div style="margin-bottom: 15px;">
      <span>产品设计审批概要</span>
    </div>
    <el-empty
      v-if="approveInstanceList.length === 0 && listLoading === false"
      :image-size="200"
    />
    <div v-else>
      <el-timeline>
        <el-timeline-item
          v-for="approveInstance in approveInstanceList"
          :key="approveInstance"
          :timestamp="UnixTime2HumanWithStr(approveInstance.created_at)"
          :color="statusOptions.find(o => o.value === approveInstance.status)?.color"
          placement="top"
        >
          <el-card>
            <h4>{{ approveInstance.dataParse.project_name + ' [' + approveInstance.dataParse.product_model + ']' }}</h4>
            <p v-for="nodeInstance in approveInstance.approve_node_instances">
              {{ getApproveDesc(nodeInstance) }}
            </p>
          </el-card>
        </el-timeline-item>
      </el-timeline>
      <div
        class="pagination-footer"
        style="margin-top: 1%"
      >
        <el-pagination
          v-model:current-page="query.page"
          v-model:page-size="query.limit"
          style="position: absolute; bottom: 10px;"
          :total="total"
          layout="total,  prev, pager, next, jumper"
          @size-change="getApproveList"
          @current-change="getApproveList"
        />
      </div>
    </div>
  </el-card>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, Ref, ref } from 'vue'
import type { ApproveInstance, ApproveInstanceQuery } from '@/model/approve'
import { GetApproveInstances } from '@/api/approve'
import { ApproveNodeInstance, ApproveNodeStatus } from '@/model/approve'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  setup() {
    const query : ApproveInstanceQuery = reactive({
      limit: 3,
      order: 'desc',
      order_by: 'created_at',
      page: 1
    })
    const approveInstanceList : Ref<ApproveInstance[]> = ref([])
    const statusOptions = [
      { color: '#409eff', label: '进行中', value: ApproveNodeStatus.ApproveNodeInaction },
      { color: '#67c23a', label: '已通过', value: ApproveNodeStatus.ApproveNodeAccepted },
      { color: '#f56c6c', label: '被拒绝', value: ApproveNodeStatus.ApproveNodeRejected },
      { color: '#fff', label: '被关闭', value: ApproveNodeStatus.ApproveNodeClosed }
    ]
    // 拼接审核流程描述
    const getApproveDesc = (nodeInstance : ApproveNodeInstance) => {
      let result = ''
      switch (nodeInstance.approve_node!.name!) {
        case '设计人员提交设计结果': {
          result += '设计人员 '
          break
        }
        case '复核人员复核设计结果': {
          result += '复核人员 '
          break
        }
        case '设计主管审核设计结果': {
          result += '设计主管 '
          break
        }
        case '总工审核设计结果': {
          result += '总工 '
          break
        }
      }
      result += nodeInstance.approval_user?.user_name_zh + ' 于 '
      if (nodeInstance.status === ApproveNodeStatus.ApproveNodeInaction) {
        result += UnixTime2HumanWithStr(nodeInstance.created_at!) + ' '
      } else {
        result += UnixTime2HumanWithStr(nodeInstance.updated_at!) + ' '
      }
      if (nodeInstance.approve_node?.name === '设计人员提交设计结果') {
        result += '提交设计'
      } else {
        result += statusOptions.find(o => o.value === nodeInstance.status)?.label
      }
      return result
    }
    const total = ref(0)
    const listLoading = ref(false)
    const getApproveList = async () => {
      listLoading.value = true
      const res = await GetApproveInstances(query)
      if (res.data.code === 200) {
        approveInstanceList.value = res.data.spec.map((item) => {
          try {
            item.dataParse = JSON.parse(Buffer.from(item.data, 'base64').toString())
          } catch (e) {
            console.log('JSON parse error', e)
          }
          return item
        })
        total.value = res.data.total
      }
      listLoading.value = false
    }
    onMounted(() => {
      getApproveList()
    })
    return {
      UnixTime2HumanWithStr,
      query,
      approveInstanceList,
      total,
      listLoading,
      getApproveList,
      statusOptions,
      getApproveDesc
    }
  }
})
</script>

<style scoped>

</style>
