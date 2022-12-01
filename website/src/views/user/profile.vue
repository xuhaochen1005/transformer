<!--
 * @Description:
 * @Autor: scy😊
 * @Date: 2021-01-23 11:02:49
 * @LastEditors: scy😊
 * @LastEditTime: 2021-01-23 11:02:50
-->
<template>
  <div class="app-container">
    <aside class="page-header">
      用户个人中心
    </aside>
    <el-collapse
      v-model="activeNames"
      style="margin-top: 1%"
    >
      <el-collapse-item
        title="账户信息"
        name="1"
      >
        <el-descriptions>
          <el-descriptions-item label="账号:">
            {{ UserBasic.name }}
          </el-descriptions-item>
          <el-descriptions-item label="用户名:">
            {{ UserBasic.user_name_zh }}
          </el-descriptions-item>
          <el-descriptions-item label="手机号:">
            {{ UserBasic.telephone }}
          </el-descriptions-item>
          <el-descriptions-item label="邮箱:">
            {{ UserBasic.email }}
          </el-descriptions-item>
          <el-descriptions-item label="用户状态:">
            <el-tag
              v-if="UserBasic.status==1"
              size="small"
            >
              正常
            </el-tag>
            <el-tag v-else>
              状态异常
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="创建时间:">
            {{ UnixTime2HumanWithStr(UserBasic.created_at) }}
          </el-descriptions-item>
          <el-descriptions-item label="更新时间:">
            {{ UnixTime2HumanWithStr(UserBasic.updated_at) }}
          </el-descriptions-item>
          <el-descriptions-item label="所属部门:">
            {{ UserBasic.department_name }}
          </el-descriptions-item>
        </el-descriptions>
      </el-collapse-item>
      <el-collapse-item
        title="我的审批"
        name="2"
      >
        <div class="filter-container">
          <el-input
            v-model="approveListQuery.project_name"
            placeholder="项目名称"
            clearable
            style="width: 200px"
            class="filter-item"
          />
          <el-input
            v-model="approveListQuery.serial_code"
            placeholder="任务书编号"
            clearable
            class="filter-item"
            style="width: 200px"
          />
          <el-select
            v-model="approveListQuery['field_eq_approve_instance.status']"
            placeholder="审核状态"
            clearable
            class="filter-item"
            style="width: 200px"
          >
            <el-option
              v-for="item in approveStatusOptions"
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
            @click="getApproveList"
          >
            搜索
          </el-button>
        </div>
        <div class="main-table">
          <el-table
            key="design_check"
            v-loading="approveListLoading"
            :data="approveDataList"
            border
            fit
            highlight-current-row
            style="width: 100%"
            @sort-change="approveSortChange"
          >
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
              min-width="280px"
            >
              <template #default="{row}">
                <span>{{ row.dataParse?.project_name }}</span>
              </template>
            </el-table-column>

            <el-table-column
              label="任务书编号"
              prop="id"
              align="center"
              min-width="120px"
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
              label="当前复核流程"
              align="center"
              min-width="180px"
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
                <span>{{ approveStatusOptions.find(o => o.value === row.approve_node_instances?.find((i) => i.next === 0)?.status)?.label }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="审批单状态"
              prop="status"
              align="center"
              width="120px"
            >
              <template #default="{row}">
                <span>{{ approveStatusOptions.find(o => o.value === row.status)?.label }}</span>
              </template>
            </el-table-column>

            <el-table-column
              label="创建时间"
              prop="created_at"
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
              width="150px"
              class-name="fixed-width"
            >
              <template #default="{row}">
                <el-button
                  type="primary"
                  size="mini"
                  @click="handleGotoApprove(row)"
                >
                  去审批
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <div
          class="pagination-footer"
          style="margin-top: 1%"
        >
          <el-pagination
            v-model:current-page="approveListQuery.page"
            v-model:page-size="approveListQuery.limit"
            :total="approveTotal"
            :page-sizes="[10, 20, 50, 100]"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="getApproveList"
            @current-change="getApproveList"
          />
        </div>
      </el-collapse-item>
      <el-collapse-item
        title="设计任务统计"
        name="3"
      >
        <div class="filter-container">
          <el-input
            v-model="listQuery['field_lk_p.project_name']"
            placeholder="项目名称"
            clearable
            style="width: 200px"
            class="filter-item"
          />
          <el-input
            v-model="listQuery['field_lk_p.product_model']"
            placeholder="产品型号"
            clearable
            style="width: 200px"
            class="filter-item"
          />
          <el-input
            v-model="listQuery['field_lk_p.serial_code']"
            placeholder="任务书编号"
            clearable
            style="width: 200px"
            class="filter-item"
          />
          <el-select
            v-model="listQuery.field_eq_task_status"
            style="width: 200px"
            placeholder="任务状态"
            clearable
            class="filter-item"
          >
            <el-option
              v-for="item in taskStatusSelectOptions"
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
        </div>
        <div class="main-table">
          <el-table
            key="tableKey"
            v-loading="listLoading"
            :data="dataList"
            border
            fit
            highlight-current-row
            style="width: 100%"
            @sort-change="taskSortChange"
          >
            <el-table-column
              label="ID"
              prop="id"
              sortable="custom"
              align="center"
              width="100px"
            >
              <template #default="{row}">
                <span>{{ row.id }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="项目名称"
              align="center"
              min-width="280px"
            >
              <template #default="{row}">
                <span>{{ row.design_project.project_name }}</span>
              </template>
            </el-table-column>

            <el-table-column
              label="产品型号"
              align="center"
              width="200px"
            >
              <template #default="{row}">
                <span>{{ row.design_project.product_model }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="任务书编号"
              align="center"
              width="130px"
            >
              <template #default="{row}">
                <span>{{ row.design_project.serial_code }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="任务状态"
              prop="task_status"
              sortable="custom"
              align="center"
              width="130px"
            >
              <template #default="{row}">
                <span>{{ taskStatusOptions.find(o => o.value === row.task_status)?.label }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="备注"
              prop="comment"
              sortable="custom"
              align="center"
              width="244"
            >
              <template #default="{row}">
                <span>{{ row.comment ? row.comment : '无' }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="创建时间"
              prop="created_at"
              sortable="custom"
              align="center"
              width="180"
            >
              <template #default="{row}">
                <span>{{ UnixTime2HumanWithStr(row.created_at) }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="产品交货时间"
              align="center"
              width="170px"
            >
              <template #default="{row}">
                <span>{{ UnixTime2HumanWithUnix(row.design_project.deliver_at) }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="任务设计时间"
              align="center"
              width="170px"
            >
              <template #default="{row}">
                <span>{{ UnixTime2HumanWithUnix(row.design_project.design_finish_at) }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="操作"
              fixed="right"
              align="center"
              min-width="280px"
              class-name="fixed-width"
            >
              <template #default="{row}">
                <el-button
                  type="success"
                  size="mini"
                  @click="handleShowDetail(row)"
                >
                  查看详情
                </el-button>
                <el-button
                  type="primary"
                  size="mini"
                  @click="handleExport(row)"
                >
                  导出设计参数
                </el-button>
                <el-button
                  type="warning"
                  size="mini"
                  @click="handleModify(row)"
                >
                  修改
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <div
          class="pagination-footer"
          style="margin-top: 1%"
        >
          <el-pagination
            v-show="total > 0"
            v-model:current-page="listQuery.page"
            v-model:page-size="listQuery.limit"
            :total="total"
            :page-sizes="[10, 20, 50, 100]"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="handleFilter"
            @current-change="handleFilter"
          />
        </div>
      </el-collapse-item>

      <!--          <el-collapse-item
        title="可控 Controllability"
        name="4"
      >
        <div>用户决策：根据场景可给予用户操作建议或安全提示，但不能代替用户进行决策；</div>
        <div>结果可控：用户可以自由的进行操作，包括撤销、回退和终止当前操作等。</div>
      </el-collapse-item>-->
    </el-collapse>
  </div>
</template>

<script lang="ts">
import { useStore } from '@/store'
import { defineComponent, onMounted, reactive, ref, Ref, toRefs } from 'vue'
import type { DesignTaskQuery } from '@/model/design'
import { DesignProject, DesignTask, DesignTaskStatus, SeamType } from '@/model/design'
import { GetUserBasic } from '@/api/user'
import { UnixTime2HumanWithStr, UnixTime2HumanWithUnix } from '@/utils/timeutils'
import { ExportDesignTask, getDesignTaskList } from '@/api/design'
import { GetApproveInstances } from '@/api/approve'
import { ApproveInstance, ApproveInstanceQuery, ApproveNodeInstance, ApproveNodeStatus } from '@/model/approve'
import { useRoute, useRouter } from 'vue-router'
import { ExportDesignProject } from '@/api/projects'
import { Response } from '@/model'
import { ElMessage } from 'element-plus'
import { get } from 'js-cookie'
export default defineComponent({
  components: {

  },
  setup() {
    const route = useRoute()
    const router = useRouter()
    const dataMap = reactive({

      activeNames: '1',
      UserBasic: {},
      total: 0,
      query: {
      },
      loading: false
    })

    const approveListQuery : ApproveInstanceQuery = reactive({
      'field_eq_approve_instance.status': undefined,
      project_name: '',
      serial_code: '',
      limit: 10,
      order: 'asc',
      order_by: 'approve_instance.status',
      page: 1,
      self: 1
    })
    const approveTotal = ref(0)
    const approveListLoading = ref(false)
    const approveStatusOptions = [
      { color: '#409eff', label: '进行中', value: ApproveNodeStatus.ApproveNodeInaction },
      { color: '#67c23a', label: '已通过', value: ApproveNodeStatus.ApproveNodeAccepted },
      { color: '#f56c6c', label: '被拒绝', value: ApproveNodeStatus.ApproveNodeRejected },
      { color: '#fff', label: '被关闭', value: ApproveNodeStatus.ApproveNodeClosed }
    ]
    const approveDataList : Ref<ApproveInstance[]> = ref([])
    const getApproveList = async () => {
      approveListLoading.value = true
      const res = await GetApproveInstances(approveListQuery)
      if (res.data.code === 200) {
        approveDataList.value = res.data.spec.map((item) => {
          try {
            item.dataParse = JSON.parse(Buffer.from(item.data, 'base64').toString())
          } catch (e) {
            console.log('JSON parse error', e)
          }
          return item
        })
        approveTotal.value = res.data.total
      }
      approveListLoading.value = false
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
        result += approveStatusOptions.find(o => o.value === nodeInstance.status)?.label
      }
      return result
    }

    const handleGotoApprove = (nodeInstance: ApproveNodeInstance) => {
      router.push(`/design/design_check?code=${nodeInstance.dataParse.serial_code}`)
    }

    const listQuery = reactive({
      'field_lk_p.serial_code': '',
      field_eq_task_status: undefined,
      page: 1,
      limit: 10,
      order: 'desc',
      order_by: 'created_at'
    })
    const showDetail = ref(false)
    const taskStatusOptions = [
      { label: '撤销', value: DesignTaskStatus.DesignTaskCanceled },
      { label: '创建中', value: DesignTaskStatus.DesignTaskCreated },
      { label: '计算中', value: DesignTaskStatus.DesignTaskStarted },
      { label: '计算完成', value: DesignTaskStatus.DesignTaskFinished },
      { label: '已选择结果', value: DesignTaskStatus.DesignTaskResultSelected },
      { label: '复核不通过', value: DesignTaskStatus.DesignTaskReviewUnaccepted },
      { label: '复核通过', value: DesignTaskStatus.DesignTaskReviewed },
      { label: '审核不通过', value: DesignTaskStatus.DesignTaskApproveUnaccepted },
      { label: '审核通过', value: DesignTaskStatus.DesignTaskApproved },
      { label: '审核不通过', value: DesignTaskStatus.DesignTaskCheckUnaccepted },
      { label: '审核通过', value: DesignTaskStatus.DesignTaskChecked }
    ]
    const taskStatusSelectOptions = [
      { label: '撤销', value: DesignTaskStatus.DesignTaskCanceled },
      { label: '创建中', value: DesignTaskStatus.DesignTaskCreated },
      { label: '计算中', value: DesignTaskStatus.DesignTaskStarted },
      { label: '已完成', value: DesignTaskStatus.DesignTaskFinished },
      { label: '复核不通过', value: DesignTaskStatus.DesignTaskReviewUnaccepted },
      { label: '复核通过', value: DesignTaskStatus.DesignTaskReviewed },
      { label: '审核不通过', value: [DesignTaskStatus.DesignTaskApproveUnaccepted, DesignTaskStatus.DesignTaskCheckUnaccepted] },
      { label: '审核通过', value: [DesignTaskStatus.DesignTaskApproved, DesignTaskStatus.DesignTaskChecked] }
    ]
    const SeamTypeOptions = [
      { label: '全斜', value: SeamType.FullOblique }
    ]
    const listLoading = ref(false)
    const total = ref(0)
    const dataList : Ref<DesignTask[]> = ref([])
    const defaultTableRow : DesignTask = {
      approve: 0,
      approve_by_creator: false,
      comment: '',
      created_at: '',
      creator: 0,
      result_progress: 0,
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
      },
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
    const specRangeOptions: Ref<number[]> = ref([])
    const tempTableRow : Ref<DesignTask> = ref(JSON.parse(JSON.stringify(defaultTableRow)) as DesignTask)
    const store = useStore()
    const detailAction = ref('create')

    const getBasicUser = async () => {
      const data = await GetUserBasic()
      dataMap.UserBasic = data?.data.spec ?? []
      dataMap.total = data?.data.total ?? 0

      // Just to simulate the time of the request
      setTimeout(() => {
        dataMap.loading = true
      }, 0.5 * 1000)
    }
    const getList = async () => {
      listLoading.value = true
      const response = await getDesignTaskList(listQuery)
      if (response.status === 200) {
        total.value = response.data.total
        dataList.value = response.data.total ? response.data.spec : []
      }
      listLoading.value = false
    }
    const handleShowDetail = async (row: DesignTask) => {
      router.push(`/design/design_create?id=${row.id}`)
    }
    const downloading = ref(false)
    async function handleExport(designTask : DesignTask) {
      downloading.value = true
      const res = await ExportDesignTask(designTask.id)
      if ((res.data as Blob).type === 'application/json') {
        const data : Response<any> = await new Promise((resolve, reject) => {
          const reader = new FileReader()
          reader.readAsText(res.data, 'utf-8')
          reader.onload = function() {
            try {
              resolve(JSON.parse(reader.result as string))
            } catch (error) {
              resolve({ code: 500, error: '文件解析失败', spec: undefined, total: 0 })
            }
          }
        })
        if (data.code !== 200) {
          downloading.value = false
          ElMessage.error(data.error)
        }
        return
      }
      // console.log(res)
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data)
      } else {
        const objectUrl = URL.createObjectURL(new Blob([res.data]))
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '设计参数: ' + designTask.design_project?.project_name + UnixTime2HumanWithUnix(Date.now() / 1000) + '.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      downloading.value = false
    }

    const handleModify = async (row: DesignTask) => {
      if (row.task_status >= DesignTaskStatus.DesignTaskStarted) {
        ElMessage.error('已经开始的任务不能修改')
        return
      }
      router.push(`/design/design_create?id=${row.id}&modify=1`)
    }
    const handleFilter = () => {
      getList()
    }

    const taskSortChange = (column: any) => {
      listQuery.order_by = column.prop
      listQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getList()
    }

    const approveSortChange = (column: any) => {
      approveListQuery.order_by = column.prop
      approveListQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getApproveList()
    }

    onMounted(() => {
      if (route.query.task_code) {
        listQuery['field_lk_p.serial_code'] = String(route.query.task_code)
        dataMap.activeNames = '3'
      }
      getBasicUser()
      getApproveList()
      getList()
    })
    return {
      UnixTime2HumanWithStr,
      UnixTime2HumanWithUnix,
      ...toRefs(dataMap),
      approveTotal,
      approveListQuery,
      approveDataList,
      approveListLoading,
      approveStatusOptions,
      ApproveNodeStatus,
      getApproveDesc,
      getApproveList,
      handleGotoApprove,
      listQuery,
      taskSortChange,
      approveSortChange,
      defaultTableRow,
      taskStatusOptions,
      taskStatusSelectOptions,
      SeamTypeOptions,
      listLoading,
      total,
      dataList,
      showDetail,
      specRangeOptions,
      handleFilter,
      handleShowDetail,
      handleModify,
      handleExport
    }
  }
})

// const defaultProfile: IProfile = {
//   name: 'Loading...',
//   email: 'Loading...',
//   avatar: 'Loading...',
//   roles: 'Loading...'
// }

// @Component({
//   name: 'Profile',
//   components: {
//     Account,
//     Activity,
//     Timeline,
//     UserCard
//   }
// })
// export default class extends Vue {
//   private user = defaultProfile
//   private activeTab = 'activity'

//   get name() {
//     return UserModule.name
//   }

//   get email() {
//     return UserModule.email
//   }

//   get avatar() {
//     return UserModule.avatar
//   }

//   get roles() {
//     return UserModule.roles
//   }

//   created() {
//     this.getUser()
//   }

//   private getUser() {
//     this.user = {
//       name: this.name,
//       email: this.email,
//       avatar: this.avatar,
//       roles: this.roles.join(' | ')
//     }
//   }
// }
</script>

<style lang="scss" scoped>
.card-header {
  display: flex;

  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}
.page-header {
  line-height: 24px;
  font-size: 24px;
  color: #409eff;
  background: none;
  padding: 0;
}

.filter-container {
  .container-row:first-child {
    margin-bottom: 10px;
  }
  .filter-item:not(button) {
    margin-right: 10px;
  }
}
.form-icon {
  font-size: 18px;
  &.warning {
    color: $yellow;
  }
}
.custom-table-popper {
  max-width: 500px;
}
</style>
