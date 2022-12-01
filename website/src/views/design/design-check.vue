<template>
  <div class="app-container">
    <aside class="page-header">
      设计复核
    </aside>
    <div class="filter-container">
      <el-input
        v-model="listQuery.project_name"
        placeholder="项目名称"
        clearable
        style="width: 200px"
        class="filter-item"
      />
      <el-input
        v-model="listQuery.serial_code"
        placeholder="任务书编号"
        clearable
        class="filter-item"
        style="width: 200px"
      />
      <el-select
        v-model="listQuery['field_eq_approve_instance.status']"
        placeholder="审核状态"
        clearable
        class="filter-item"
        style="width: 200px"
      >
        <el-option
          v-for="item in statusOptions"
          :key="item"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
      <el-button
        v-waves
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="handleFilter"
      >
        搜索
      </el-button>
      <!--      <el-button-->
      <!--        v-waves-->
      <!--        :loading="downloadLoading"-->
      <!--        class="filter-item"-->
      <!--        type="primary"-->
      <!--        icon="el-icon-download"-->
      <!--        @click="handleDownload"-->
      <!--      >-->
      <!--        导出-->
      <!--      </el-button>-->
    </div>
    <div class="main-table">
      <el-table
        key="design_check"
        v-loading="listLoading"
        :data="dataList"
        border
        fit
        highlight-current-row
        style="width: 100%"
        row-key="dataParse.serial_code"
        :expand-row-keys="[listQuery.serial_code]"
        @sort-change="sortChange"
      >
        <el-table-column type="expand">
          <template #default="props">
            <el-form
              label-position="left"
              inline
              class="approve-table-expand"
            >
              <el-timeline>
                <span
                  v-for="nodeInstance in props.row.approve_node_instances"
                  :key="nodeInstance.id"
                >
                  <el-timeline-item
                    :timestamp="UnixTime2HumanWithStr(nodeInstance.created_at)"
                    :color="statusOptions.find(o => o.value === nodeInstance.status)?.color"
                    placement="top"
                  >
                    <el-card>
                      <h4>{{ nodeInstance.approve_node.name }}</h4>
                      <p>
                        {{ getApproveDesc(nodeInstance) }}
                      </p>
                    </el-card>
                  </el-timeline-item>
                  <el-timeline-item
                    v-if="[ApproveNodeStatus.ApproveNodeRejected,ApproveNodeStatus.ApproveNodeClosed].includes(nodeInstance.status)"
                    :timestamp="UnixTime2HumanWithStr(nodeInstance.updated_at)"
                    :color="statusOptions.find(o => o.value === nodeInstance.status)?.color"
                    placement="top"
                  >
                    <el-card>
                      <h4>{{ '设计结果不通过' }}</h4>
                      <p>
                        {{ getApproveDesc(nodeInstance) }}
                      </p>
                    </el-card>
                  </el-timeline-item>
                </span>
              </el-timeline>
            </el-form>
          </template>
        </el-table-column>
        <el-table-column
          label="ID"
          sortable="custom"
          align="center"
          min-width="80px"
        >
          <template #default="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="项目名称"
          align="center"
          width="280px"
        >
          <template #default="{row}">
            <span>{{ row.dataParse?.project_name }}</span>
          </template>
        </el-table-column>

        <el-table-column
          label="任务书编号"
          prop="id"
          align="center"
          width="100px"
        >
          <template #default="{row}">
            <span>{{ row.dataParse?.serial_code }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="设计人员"
          prop="id"
          align="center"
          width="100px"
        >
          <template #default="{row}">
            <span>{{ row.applicant_user.user_name_zh }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="计算进度"
          align="center"
          width="100px"
        >
          <template #default="{row}">
            <el-tooltip
              v-if="row.design_task?.task_status === DesignTaskStatus.DesignTaskCanceled"
              popper-class="msg-tooltip"
            >
              <template
                #content
              >
                <span v-html="`<div>设计任务计算失败，请点击<a href='/#/user/profile?task_code=${row.design_task.design_project.serial_code}'>任务地址</a>进行修改</div>`" />
              </template>
              <el-tag
                type="danger"
              >
                <span>计算失败</span>
              </el-tag>
            </el-tooltip>
            <el-tag
              v-else
              v-loading="refreshApproveList.includes(row.id)"
              :type="row.design_task.result_progress === 1 ? 'success' : 'warning'"
            >
              <span>{{ (row.design_task.result_progress * 100).toFixed(2) }} %</span>
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          label="当前复核流程"
          prop="id"
          align="center"
          width="170px"
        >
          <template #default="{row}">
            <span>{{ row.approve_node_instances?.find((i) => i.next === 0)?.approve_node.name }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="当前复核状态"
          prop="id"
          align="center"
          width="110px"
        >
          <template #default="{row}">
            <el-tag :type="statusOptions.find(o => o.value === row.approve_node_instances?.find((i) => i.next === 0)?.status)?.type">
              {{ statusOptions.find(o => o.value === row.approve_node_instances?.find((i) => i.next === 0)?.status)?.label }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          label="审批单状态"
          prop="status"
          sortable="custom"
          align="center"
          width="120px"
        >
          <template #default="{row}">
            <el-tag :type="statusOptions.find(o => o.value === row.status)?.type">
              {{ statusOptions.find(o => o.value === row.status)?.label }}
            </el-tag>
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
          label="操作"
          align="center"
          fixed="right"
          min-width="500px"
          class-name="fixed-width"
        >
          <template #default="{row}">
            <div style="display: flex;flex-wrap: nowrap;justify-content: space-evenly;padding: 0 10px;">
              <span>
                <el-tooltip
                  v-if="row.design_task.task_status !== DesignTaskStatus.DesignTaskFinished"
                  :content="row.design_task.task_status < DesignTaskStatus.DesignTaskFinished ? '设计任务未计算完成' : '已选择设计结果'"
                  placement="top"
                >
                  <span
                    v-role="[RoleList.Designer]"
                  >
                    <el-button
                      disabled
                      type="warning"
                      size="mini"
                      @click="handleShowResults(row)"
                    >
                      选择设计结果
                    </el-button>
                  </span>
                </el-tooltip>

                <span
                  v-else
                  v-role="[RoleList.Designer]"
                >
                  <el-button
                    type="warning"
                    size="mini"
                    @click="handleShowResults(row)"
                  >
                    选择设计结果
                  </el-button>
                </span>
              </span>
              <span>
                <el-tooltip
                  v-if="row.design_task.task_status < DesignTaskStatus.DesignTaskResultSelected"
                  :content="'设计结果计算中'"
                  placement="top"
                >
                  <span>
                    <el-button
                      style="pointer-events: none;"
                      disabled
                      type="primary"
                      size="mini"
                      @click="handleDesignResult(row)"
                    >
                      查看设计结果
                    </el-button>
                  </span>
                </el-tooltip>
                <el-button
                  v-else
                  type="primary"
                  size="mini"
                  @click="handleDesignResult(row)"
                >
                  查看设计结果
                </el-button>
              </span>
              <span>
                <el-tooltip
                  v-if="!couldApprove(row)"
                  placement="top"
                  :content="row.status === ApproveStatus.ApproveNodeInaction ? '暂时无法审批': '已审批'"
                >
                  <span>
                    <el-button
                      :disabled="!couldApprove(row)"
                      size="mini"
                      type="success"
                      @click="handleApprove(row,ApproveNodeStatus.ApproveNodeAccepted)"
                    >
                      通过
                    </el-button>

                  </span>
                </el-tooltip>
                <el-button
                  v-else
                  size="mini"
                  type="success"
                  @click="handleApprove(row,ApproveNodeStatus.ApproveNodeAccepted)"
                >
                  通过
                </el-button>
              </span>
              <span>
                <el-tooltip
                  v-if="!couldApprove(row)"
                  placement="top"
                  :content="row.status === ApproveStatus.ApproveNodeInaction ? '暂时无法审批': '已审批'"
                >
                  <span>
                    <el-button
                      disabled
                      size="mini"
                      type="danger"
                      @click="handleApprove(row,ApproveNodeStatus.ApproveNodeRejected)"
                    >
                      不通过
                    </el-button>
                  </span>

                </el-tooltip>
                <el-button
                  v-else
                  size="mini"
                  type="danger"
                  @click="handleApprove(row,ApproveNodeStatus.ApproveNodeRejected)"
                >
                  不通过
                </el-button>

              </span>
              <span>
                <el-tooltip
                  v-if="!couldApprove(row)"
                  placement="top"
                  :content="row.status === ApproveStatus.ApproveNodeInaction ? '暂时无法审批': '已审批'"
                >
                  <span>
                    <el-button
                      disabled
                      size="mini"
                      type="danger"
                      @click="handleApprove(row,ApproveNodeStatus.ApproveNodeClosed)"
                    >
                      关闭
                    </el-button>
                  </span>
                </el-tooltip>
                <el-button
                  v-else
                  size="mini"
                  type="danger"
                  @click="handleApprove(row,ApproveNodeStatus.ApproveNodeClosed)"
                >
                  关闭
                </el-button>
              </span>
            </div>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div
      class="pagination-footer"
      style="margin-top: 1%"
    >
      <el-pagination
        v-model:current-page="listQuery.page"
        v-model:page-size="listQuery.limit"
        :total="total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleFilter"
        @current-change="handleFilter"
      />
    </div>
    <div class="dialogs-main">
      <el-dialog
        v-model="showDesignResults"
        custom-class="dialog-results"
        title="任务设计计算结果"
        width="80%"
      >
        <design-results-table
          :current-design-results="currentDesignResults"
          :show-design-results-mode="showDesignResultsMode"
          @handleSelectionChange="handleSelectionChange"
        />
        <template #footer>
          <span class="dialog-footer">
            <el-button
              v-if="showDesignResultsMode === 'create'"
              type="primary"
              style="margin-right: 10px;"
              @click="handleSelectDesignResults"
            >确认选择结果</el-button>
            <el-button @click="showDesignResults = false">关闭</el-button>
          </span>
        </template>
      </el-dialog>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, h, onMounted, reactive, ref, Ref } from 'vue'
import {
  ApproveInstance,
  ApproveInstanceQuery,
  ApproveNodeDesignTaskHeadID,
  ApproveNodeInstance,
  ApproveNodeStatus,
  ApproveStatus
} from '@/model/approve'

import { GetApproveInstances, UpdateApproveNodeInstance } from '@/api/approve'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import { useStore } from '@/store'
import {
  DecodeDesignTask,
  DesignResults,
  DesignResultStatus,
  DesignTaskQuery,
  DesignTaskStatus,
  TempTaskResults
} from '@/model/design'
import { decodeObjectBytes, encodeBytesData, stringify } from '@/utils/jsonutils'
import { ElLoading, ElMessage, ElMessageBox } from 'element-plus'
import { useRoute, useRouter } from 'vue-router'
import { RoleList } from '@/model/permission'
import { GetDesignResults, getDesignTaskList, getSpecDesignTask, UpdateDesignResults } from '@/api/design'
import design from '@/router/asyncRoutes/design'
import { CreateDesignResults } from '@/api/design_results'
import DesignResultsTable from '@/views/design/components/design-results-table.vue'

export default defineComponent({
  components: {
    DesignResultsTable
  },
  setup(props, context) {
    const route = useRoute()
    const router = useRouter()
    const state = useStore().state
    const listQuery : ApproveInstanceQuery = reactive({
      'field_eq_approve_instance.status': undefined,
      project_name: '',
      serial_code: '',
      limit: 10,
      order: 'asc',
      order_by: 'created_at desc,approve_instance.status',
      page: 1
    })
    const total = ref(0)
    const listLoading = ref(false)
    const statusOptions = [
      { type: '', color: '#409eff', label: '进行中', value: ApproveNodeStatus.ApproveNodeInaction },
      { type: 'success', color: '#67c23a', label: '已通过', value: ApproveNodeStatus.ApproveNodeAccepted },
      { type: 'danger', color: '#f56c6c', label: '被拒绝', value: ApproveNodeStatus.ApproveNodeRejected },
      { type: 'danger', color: '#fff', label: '被关闭', value: ApproveNodeStatus.ApproveNodeClosed }
    ]
    const dataList : Ref<ApproveInstance[]> = ref([])
    const getList = async () => {
      listLoading.value = true
      const res = await GetApproveInstances(listQuery)
      if (res.data.code === 200) {
        const taskList: number[] = []
        const tempDataList = res.data.spec.map((item) => {
          try {
            item.dataParse = JSON.parse(Buffer.from(item.data, 'base64').toString())
            // 设计审批
            if (item.approve_id === 1 && item.dataParse && item.dataParse.id) {
              taskList.push(item.dataParse.id)
            }
          } catch (e) {
            console.log('JSON parse error', e)
          }
          return item
        })
        const res2 = await getDesignTaskList({
          'field_eq_design_task.id': taskList, limit: 9999, order: '', order_by: '', page: 1
        } as DesignTaskQuery)
        if (res2.data.code === 200) {
          tempDataList.forEach((item) => {
            if (item.dataParse && item.dataParse.id) {
              const task = (res2.data.spec as DecodeDesignTask[]).find((task) => task.id === item.dataParse!.id)
              if (task) {
                decodeObjectBytes(task)
                item.design_task = task
              }
            }
          })
        }
        dataList.value = tempDataList
        total.value = res.data.total
      }
      listLoading.value = false
    }

    const handleFilter = () => {
      getList()
    }
    const sortChange = (column : any) => {
      listQuery.order_by = column.prop
      listQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getList()
    }
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
        case '总工审核设计结果': {
          result += '总工 '
        }
      }
      result += nodeInstance.approval_user?.user_name_zh + ' 于 '
      if (nodeInstance.status === ApproveNodeStatus.ApproveNodeInaction) {
        result += UnixTime2HumanWithStr(nodeInstance.created_at!) + ' '
      } else {
        result += UnixTime2HumanWithStr(nodeInstance.updated_at!) + ' '
      }
      if (nodeInstance.approve_node?.name === '设计人员提交设计结果') {
        result += '提交'
      } else {
        result += statusOptions.find(o => o.value === nodeInstance.status)?.label
      }
      return result
    }
    // 是否可以操作
    const couldApprove = (nodeInstance: ApproveInstance): Boolean => {
      if (!nodeInstance.approve_node_instances) {
        return false
      }
      if (nodeInstance.approve_node_instances.length <= 0) {
        return false
      }
      if (nodeInstance.status !== ApproveStatus.ApproveNodeInaction) {
        return false
      }
      if (nodeInstance.design_task!.task_status < DesignTaskStatus.DesignTaskResultSelected) {
        return false
      }
      const last = nodeInstance.approve_node_instances!.find(i => i.next === 0)
      // 设计节点由系统自动审核，不需要人工审核
      if (last?.approve_node!.id === ApproveNodeDesignTaskHeadID) {
        return false
      }
      // 当前任务的设计主管可以执行审批
      if (nodeInstance.dataParse!.dermander_id === state.user.id) {
        return true
      }

      // 只有审批人是自己时才允许进行审批
      return nodeInstance.approve_node_instances!.find((i) => i.next === 0)?.approval_user?.name === state.user.name
    }
    // 设计结果
    const handleDesignResult = async (approveInstance: ApproveInstance) => {
      const loading = ElLoading.service({
        lock: true,
        background: 'rgba(0, 0, 0, 0.5)'
      })
      try {
        const res = await getSpecDesignTask(approveInstance.design_task!.id)
        if (res.data.code === 200) {
          console.log(res.data.spec)
          decodeObjectBytes(res.data.spec)
          const designTask = res.data.spec
          const tempDesignTask = JSON.parse(JSON.stringify(designTask))
          if (designTask) {
            if (designTask.task_status >= DesignTaskStatus.DesignTaskResultSelected) {
              currentDesignResults.value = designTask.final_design_results.map((item) => {
                item.design_task = tempDesignTask
                return item
              })
              showDesignResultsMode.value = 'view'
              showDesignResults.value = true
              return
            }
          }
        }
      } catch (e) {
        console.log(e)
      } finally {
        loading.close()
      }

      ElMessage.info('设计结果正在计算中，暂时不能查看设计结果')
    }

    const comment = ref('')
    const approveNodeInstance = async (approveInstance: ApproveInstance, status: ApproveNodeStatus) => {
      try {
        const currentNode = JSON.parse(stringify(approveInstance.approve_node_instances!.find((i) => i.next === 0)!)) as ApproveNodeInstance
        currentNode.status = status
        currentNode.comment = comment.value
        currentNode.data = Buffer.from(stringify({
          id: approveInstance.dataParse!.id,
          comment: comment.value
        })).toString('base64')
        listLoading.value = true
        const res = await UpdateApproveNodeInstance(currentNode)
        if (res.data.code === 200) {
          ElMessage.success('审批操作成功')
          getList()
        }
      } catch (e) {
        console.log(e)
      }
      listLoading.value = false
    }
    const handleApprove = async (approveInstance: ApproveInstance, status: ApproveNodeStatus) => {
      if (!approveInstance) {
        return
      }
      if (!approveInstance.approve_node_instances) {
        return
      }
      const last = approveInstance.approve_node_instances!.find((i) => i.next === 0)
      if (!last) {
        return
      }
      // 当前节点已经被拒绝，暂时不可操作
      if (last!.status === ApproveStatus.ApproveNodeRejected && status !== ApproveNodeStatus.ApproveNodeClosed) {
        ElMessage.info('当前审批节点已被拒绝，暂时不可操作')
        return
      }
      const loading = ElLoading.service({
        lock: true,
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.5)'
      })
      const res = await getSpecDesignTask(approveInstance.dataParse!.id)
      if (res.data.code !== 200) {
        loading.close()
        return
      }
      const designTask = res.data.spec
      if (!designTask) {
        loading.close()
        return
      }
      if (designTask.task_status <= DesignTaskStatus.DesignTaskFinished) {
        ElMessage.info('设计结果正在计算中，暂时不能进行审批操作')
        loading.close()
        return
      }
      if (status === ApproveNodeStatus.ApproveNodeRejected) {
        comment.value = ''
        await ElMessageBox({
          title: '审批',
          message: h('textarea', {
            value: '',
            onInput: (e: any) => {
              comment.value = e.target.value
            },
            rows: 5,
            style: {
              width: '100%'
            },
            placeholder: '请输入不通过原因'
          }),
          showCancelButton: true,
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          beforeClose: async (action, instance, done) => {
            if (action === 'confirm') {
              instance.confirmButtonLoading = true
              await approveNodeInstance(approveInstance, status)
            }
            instance.confirmButtonLoading = false
            done()
            loading.close()
          }
        })
        return
      }
      loading.close()
      await approveNodeInstance(approveInstance, status)
    }

    const showDesignResults = ref(false)
    const showDesignResultsMode: Ref<'create'|'view'> = ref('create')
    const currentDesignResults : Ref<DesignResults[]> = ref([])
    const currentDesignTaskID = ref(0)
    const handleShowResults = async (approveInstance: ApproveInstance) => {
      finalDesignResults.value = []
      const loading = ElLoading.service({
        lock: true,
        background: 'rgba(0, 0, 0, 0.5)'
      })
      const res = await getSpecDesignTask(approveInstance.design_task!.id)
      if (res.data.code !== 200) {
        return
      }
      const designTask = res.data.spec
      decodeObjectBytes(designTask)
      if (designTask) {
        if (designTask.task_status !== DesignTaskStatus.DesignTaskFinished) {
          ElMessage.info(designTask.task_status > DesignTaskStatus.DesignTaskFinished ? '已经选择过设计结果' : '设计结果计算中')
          loading.close()
          return
        }
        try {
          const taskResults = designTask.design_results as Object as TempTaskResults
          currentDesignResults.value = []
          taskResults['pops'].forEach((item: any[], index: number) => {
            item[0].project_name = designTask.design_project?.project_name
            item[0].product_name = designTask.design_project?.product_name
            item[0].project_code = designTask.design_project?.product_code
            item[0].total_cost = taskResults['objs'][index]['总成本']
            item[0].design_project = designTask.design_project
            item[0].design_task = JSON.parse(JSON.stringify(designTask))
            item[0].design_task_id = designTask.id
            item[0].design_project_id = designTask.design_project_id
            item[0].cost_result = taskResults['objs'][index]
            currentDesignResults.value.push(item[0])
          })
        } catch (e) {
          console.log(designTask!.design_results, e)
        }
        if (currentDesignResults.value.length > 0) {
          showDesignResultsMode.value = 'create'
          showDesignResults.value = true
          currentDesignTaskID.value = designTask.id
          loading.close()
          return
        }
      }
      loading.close()
      ElMessage.info('设计结果正在计算中，暂时不能选择设计结果')
    }
    // 查看设计详情
    const currentDesignResultDetail = ref<DesignResults | undefined>(undefined)
    const showDesignResultDetail = ref(false)
    const handleSelectionChange = (results: DesignResults[]) => {
      finalDesignResults.value = results
    }
    const finalDesignResults : Ref<DesignResults[]> = ref([])
    const handleSelectDesignResults = async () => {
      if (finalDesignResults.value.length === 0) {
        ElMessage.info('请选择设计结果')
        return
      }
      const loading = ElLoading.service({
        target: '.dialog-results',
        lock: true,
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.5)'
      })
      const data = JSON.parse(JSON.stringify(finalDesignResults.value.map((result: DesignResults) => {
        result.design_task = undefined
        result.design_project = undefined
        result.hv_result = encodeBytesData(result.hv_result) as Object
        result.lv_result = encodeBytesData(result.lv_result) as Object
        result.core_result = encodeBytesData(result.core_result) as Object
        result.cost_result = encodeBytesData(result.cost_result) as Object
        result.temp_result = encodeBytesData(result.temp_result) as Object
        result.continue_vars = encodeBytesData(result.continue_vars) as Object
        result.performance_result = encodeBytesData(result.performance_result) as Object
        result.impedance_result = encodeBytesData(result.impedance_result) as Object
        return result
      }))) as DesignResults[]

      const res = await CreateDesignResults(data)
      if (res.data.code === 200) {
        ElMessage.success('操作成功')
        showDesignResults.value = false
        await getList()
      }
      loading.close()
    }
    const refreshApproveList : Ref<number[]> = ref([])
    onMounted(() => {
      if (route.query.code) {
        listQuery.serial_code = route.query.code!.toString()
      }
      getList()
      // 刷新列表
      const timer = setInterval(async () => {
        let refreshList: Number[] = []

        dataList.value.forEach((item) => {
          if (item.design_task?.task_status === DesignTaskStatus.DesignTaskStarted && item.design_task.result_progress < 1) {
            refreshList.push(item.design_task.id)
            refreshApproveList.value.push(item.id)
          }
        })
        if (refreshList.length === 0) {
          return
        }
        const res = await getDesignTaskList({
          'field_eq_design_task.id': refreshList, limit: 9999, order: '', order_by: '', page: 1
        } as DesignTaskQuery)
        if (res.data.code === 200) {
          dataList.value.forEach((item) => {
            if (item.dataParse && item.dataParse.id) {
              const task = res.data.spec.find((task) => task.id === item.dataParse!.id)
              if (task) {
                item.design_task = task
                if (task.result_progress === 1 || task.task_status !== DesignTaskStatus.DesignTaskStarted) {
                  getList()
                }
              }
            }
          })
        }
        refreshList = []
        refreshApproveList.value = []
      }, 10000)
    })
    return {
      RoleList,
      state,
      UnixTime2HumanWithStr,
      handleFilter,
      listQuery,
      total,
      listLoading,
      statusOptions,
      dataList,
      getApproveDesc,
      couldApprove,
      sortChange,
      showDesignResults,
      showDesignResultsMode,
      handleShowResults,
      currentDesignResults,
      currentDesignTaskID,
      DesignTaskStatus,
      handleDesignResult,
      handleApprove,
      ApproveStatus,
      ApproveNodeStatus,
      currentDesignResultDetail,
      showDesignResultDetail,
      handleSelectionChange,
      finalDesignResults,
      handleSelectDesignResults,
      refreshApproveList
    }
  }
})
</script>

<style lang="scss" scoped>
.el-button+.el-button {
  margin-left: 0;
}
.filter-container {
  .container-row:first-child {
    margin-bottom: 10px;
  }
  .filter-item:not(button) {
    margin-right: 10px;
  }
}

</style>
<style lang="scss">
.msg-tooltip {
  a {
    border-bottom: 1px solid $--color-primary;
    color: $--color-primary;
  }
}
</style>
