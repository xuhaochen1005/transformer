<template>
  <div class="app-container">
    <aside class="page-header">
      系统监控和维护
    </aside>
    <div class="main-table">
      <el-tabs
        v-model="activeName"
        @tab-click="handleClick"
      >
        <el-tab-pane
          label="系统监控"
          name="monitor"
        >
          <el-row :gutter="20">
            <el-col
              :span="12"
              style="height: 350px;"
            >
              <div
                class="grid-content bg-purple"
                style="height: 100%;"
              >
                <el-card
                  class="box-card"
                >
                  <template #header>
                    <div class="card-header">
                      <span>系统基本信息</span>
                    </div>
                  </template>
                  <el-descriptions
                    class="table-system-info"
                    :column="3"
                    border
                  >
                    <el-descriptions-item
                      label="主机名"
                      label-align="left"
                      align="center"
                      label-class-name="my-label"
                      class-name="my-content"
                    >
                      {{ systemHost.hostname }}
                    </el-descriptions-item>
                    <el-descriptions-item
                      label="运行时长"
                      label-align="left"
                      align="center"
                    >
                      {{ (systemHost.uptime/60/60).toFixed(2) }}小时
                    </el-descriptions-item>
                    <el-descriptions-item
                      label="开机时间"
                      label-align="left"
                      align="center"
                    >
                      {{ UnixTime2Human(systemHost.bootTime) }}
                    </el-descriptions-item>
                    <el-descriptions-item
                      label="操作系统"
                      label-align="left"
                      align="center"
                    >
                      <el-tag size="small">
                        {{ systemHost.os }}
                      </el-tag>
                    </el-descriptions-item>
                    <el-descriptions-item
                      label="主机平台"
                      label-align="left"
                      align="center"
                    >
                      {{ systemHost.platform }}
                    </el-descriptions-item>
                    <el-descriptions-item
                      label="内核版本"
                      label-align="left"
                      align="center"
                    >
                      <el-tag size="small">
                        {{ systemHost.kernelVersion }}
                      </el-tag>
                    </el-descriptions-item>
                    <el-descriptions-item
                      label="系统架构"
                      label-align="left"
                      align="center"
                    >
                      {{ systemHost.kernelArch }}
                    </el-descriptions-item>
                    <el-descriptions-item
                      label="主机ID"
                      label-align="left"
                      align="center"
                    >
                      {{ systemHost.hostid }}
                    </el-descriptions-item>
                  </el-descriptions>
                </el-card>
              </div>
            </el-col>
            <el-col
              :span="12"
              style="height: 350px;"
            >
              <div class="grid-content bg-purple">
                <el-card class="box-card">
                  <template #header>
                    <div class="card-header">
                      <span>系统内存-CPU信息</span>
                    </div>
                  </template>
                  <el-descriptions
                    :column="3"
                    border
                    class="table-system-cpu"
                  >
                    <el-descriptions-item
                      v-if="systemMemory.swap_mem"
                      label="交换内存总量"
                      label-align="left"
                      align="center"
                      class-name="my-content"
                    >
                      {{ (systemMemory.swap_mem.total/1024/1024 /1024).toFixed(2) }}GB
                    </el-descriptions-item>
                    <el-descriptions-item
                      v-if="systemMemory.swap_mem"
                      label="交换内存用量"
                      label-align="left"
                      align="center"
                    >
                      {{ (systemMemory.swap_mem.used /1024/1024 /1024).toFixed(2) }}GB
                    </el-descriptions-item>

                    <el-descriptions-item
                      v-if="systemMemory.swap_mem"
                      label="用量占比"
                      label-align="left"
                      align="center"
                    >
                      <el-tag>
                        {{ (systemMemory.swap_mem.usedPercent).toFixed(2) }}%
                      </el-tag>
                    </el-descriptions-item>
                    <el-descriptions-item
                      v-if=" systemMemory.virtaul_mem"
                      label="虚拟内存总量"
                      label-align="left"
                      align="center"
                      label-class-name="my-label"
                      class-name="my-content"
                      width="150px"
                    >
                      {{ (systemMemory.virtaul_mem.total/1024/1024 /1024).toFixed(2) }}GB
                    </el-descriptions-item>
                    <el-descriptions-item
                      v-if=" systemMemory.virtaul_mem"
                      label="虚拟内存用量"
                      label-align="left"

                      align="center"
                    >
                      {{ (systemMemory.virtaul_mem.used /1024/1024 /1024).toFixed(2) }}GB
                    </el-descriptions-item>

                    <el-descriptions-item

                      v-if=" systemMemory.virtaul_mem"
                      label="用量占比"
                      label-align="left"
                      align="center"
                    >
                      <el-tag>
                        {{ (systemMemory.virtaul_mem.usedPercent).toFixed(2) }} %
                      </el-tag>
                    </el-descriptions-item>
                    <el-descriptions-item
                      v-if="systemCpu.physical_cnt"
                      label="物理核数"
                      label-align="left"
                      align="center"
                      class-name="my-content"
                    >
                      {{ systemCpu.physical_cnt }}
                    </el-descriptions-item>
                    <el-descriptions-item
                      v-if="systemCpu.logical_cnt"
                      label="逻辑核数"
                      label-align="left"
                      align="center"
                    >
                      {{ systemCpu.logical_cnt }}
                    </el-descriptions-item>

                    <el-descriptions-item
                      v-if="systemCpu.total_percent"
                      label="cpu 总负载"
                      label-align="left"
                      align="center"
                    >
                      <el-tag>
                        {{ systemCpu.total_percent[0].toFixed(2) }}%
                      </el-tag>
                    </el-descriptions-item>
                    <el-descriptions-item
                      v-if=" systemCpu.load_avg"
                      label="最近一分钟负载"
                      label-align="left"
                      align="center"
                      label-class-name="my-label"
                      class-name="my-content"
                      width="150px"
                    >
                      <el-tag>
                        {{ (systemCpu.load_avg.load1*100).toFixed(2) }}%
                      </el-tag>
                    </el-descriptions-item>
                    <el-descriptions-item
                      v-if=" systemMemory.virtaul_mem"
                      label="最近五分钟负载"
                      label-align="left"

                      align="center"
                    >
                      <el-tag>
                        {{ (systemCpu.load_avg.load5*100).toFixed(2) }}%
                      </el-tag>
                    </el-descriptions-item>

                    <el-descriptions-item

                      v-if=" systemMemory.virtaul_mem"
                      label="最近十五分钟负载"
                      label-align="left"
                      align="center"
                    >
                      <el-tag>
                        {{ (systemCpu.load_avg.load15*100).toFixed(2) }}%
                      </el-tag>
                    </el-descriptions-item>
                  </el-descriptions>
                </el-card>
              </div>
            </el-col>
          </el-row>
        </el-tab-pane>
        <!--        <el-tab-pane
          label="系统维护"
          name="second"
        >
          系统维护
        </el-tab-pane>-->
      </el-tabs>
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
        :current-page="currentPage4"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>
<script lang="ts">
import {
  defineComponent,
  onMounted,
  reactive,
  watch,
  ref,
  toRefs
} from 'vue'
import { useRoute } from 'vue-router'
// import { UnixTime2Human } from '@/utils/timeutils'
// import { useStore } from '@/store'
import { GetSysBasic, GetSysCPU, GetSysMemory, GetSysHost, GetSysNet, GetSysDisk, GetSysProcess } from '@/api/maintenance'
import dayjs from 'dayjs'

export default defineComponent({
  setup() {
    const userNameRef = ref(null)
    const passwordRef = ref(null)
    const loginFormRef = ref(null)
    const route = useRoute()
    const state = reactive({
      loginForm: {
        username: 'admin',
        password: ''
      },
      loginRules: {
        username: [{ validator: userNameRef, trigger: 'blur' }],
        password: [{ validator: passwordRef, trigger: 'blur' }]
      },
      passwordType: 'password',
      loading: false,
      showDialog: false,
      capsTooltip: false,
      active: 'monitor',
      redirect: '',
      otherQuery: {},
      total: 0,
      userList: [],
      systemHost: {},
      systemCpu: {
        load_avg: {}
      },
      systemBasic: {},
      systemMemory: { },
      systemNet: {},
      systemDisk: {},
      systemProcess: {},
      listLoading: false,
      activeName: 'monitor',
      listQuery: {
        field_eq_name: '',
        field_eq_department_id: 1,
        'field_eq_user.status': 1,
        page: 1,
        limit: 10,
        importance: undefined,
        title: undefined,
        type: undefined,
        order_by: '+id',
        order: '',
        sort: ''
      }
    })

    const methods = reactive({
      UnixTime2Human: (timestamp: number) => {
        // UnixTime2Human(timestamp)
        return dayjs.unix(timestamp).format('YYYY-MM-DD HH:mm:ss')
      },
      sortByID: (order: string) => {
        if (order === 'ascending') {
          // listQuery.sort = '+id'
        } else {
          // listQuery.sort = '-id'
        }
        // dataMap.handleFilter()
      },
      getSortClass: (key: string) => {
        // const sort = dataMap.listQuery.sort
        // return sort === `+${key}` ? 'ascending' : 'descending'
        console.log(key)
      },

      async getSystemHost(page?: any, total?: any, limit?: any) {
        if (page) {
          state.listQuery.page = page
        }
        if (limit) {
          state.listQuery.limit = limit
        }
        state.listLoading = true
        const data = await GetSysHost(state.listQuery)
        state.systemHost = data?.data.spec ?? []
        state.total = data?.data.total ?? 0

        // Just to simulate the time of the request
        setTimeout(() => {
          state.listLoading = false
        }, 0.5 * 1000)
      },
      async getSystemCPU(page?: any, total?: any, limit?: any) {
        if (page) {
          state.listQuery.page = page
        }
        if (limit) {
          state.listQuery.limit = limit
        }
        state.listLoading = true
        const data = await GetSysCPU(state.listQuery)
        state.systemCpu = data?.data.spec ?? []
        state.total = data?.data.total ?? 0

        // Just to simulate the time of the request
        setTimeout(() => {
          state.listLoading = false
        }, 0.5 * 1000)
      },
      async getSystemBasic(page?: any, total?: any, limit?: any) {
        if (page) {
          state.listQuery.page = page
        }
        if (limit) {
          state.listQuery.limit = limit
        }
        state.listLoading = true
        const data = await GetSysBasic(state.listQuery)
        state.systemBasic = data?.data.spec ?? []
        state.total = data?.data.total ?? 0

        // Just to simulate the time of the request
        setTimeout(() => {
          state.listLoading = false
        }, 0.5 * 1000)
      },
      async getSystemMemory(page?: any, total?: any, limit?: any) {
        if (page) {
          state.listQuery.page = page
        }
        if (limit) {
          state.listQuery.limit = limit
        }
        state.listLoading = true
        const data = await GetSysMemory(state.listQuery)
        state.systemMemory = data?.data.spec ?? []
        state.total = data?.data.total ?? 0

        // Just to simulate the time of the request
        setTimeout(() => {
          state.listLoading = false
        }, 0.5 * 1000)
      },
      async getSystemProcess(page?: any, total?: any, limit?: any) {
        if (page) {
          state.listQuery.page = page
        }
        if (limit) {
          state.listQuery.limit = limit
        }
        state.listLoading = true
        const data = await GetSysProcess(state.listQuery)
        state.systemProcess = data?.data.spec ?? []
        state.total = data?.data.total ?? 0

        // Just to simulate the time of the request
        setTimeout(() => {
          state.listLoading = false
        }, 0.5 * 1000)
      },
      async getSystemDisk(page?: any, total?: any, limit?: any) {
        if (page) {
          state.listQuery.page = page
        }
        if (limit) {
          state.listQuery.limit = limit
        }
        state.listLoading = true
        const data = await GetSysDisk(state.listQuery)
        state.systemDisk = data?.data.spec ?? []
        state.total = data?.data.total ?? 0

        // Just to simulate the time of the request
        setTimeout(() => {
          state.listLoading = false
        }, 0.5 * 1000)
      },
      async getSystemNet(page?: any, total?: any, limit?: any) {
        if (page) {
          state.listQuery.page = page
        }
        if (limit) {
          state.listQuery.limit = limit
        }
        state.listLoading = true
        const data = await GetSysNet(state.listQuery)
        state.systemNet = data?.data.spec ?? []
        state.total = data?.data.total ?? 0

        // Just to simulate the time of the request
        setTimeout(() => {
          state.listLoading = false
        }, 0.5 * 1000)
      }
    })

    /*    function getOtherQuery(query: LocationQuery) {
      return Object.keys(query).reduce((acc, cur) => {
        if (cur !== 'redirect') {
          acc[cur] = query[cur]
        }
        return acc
      }, {} as LocationQuery)
    } */

    watch(() => route.query, query => {
      if (query) {
        state.redirect = query.redirect?.toString() ?? ''
        // state.otherQuery = getOtherQuery(query)
      }
    })

    onMounted(() => {
      /* if (state.loginForm.username === '') {
        (userNameRef.value as any).focus()
      } else if (state.loginForm.password === '') {
        (passwordRef.value as any).focus()
      } */
      //   methods.getUserList(null, null, 20)
      methods.getSystemHost(null, null, 20)
      methods.getSystemMemory(null, null, 20)
      methods.getSystemCPU(null, null, 20)
    })

    return {
      userNameRef,
      passwordRef,
      loginFormRef,
      ...toRefs(state),
      ...toRefs(methods)
    }
  }
})
</script>

<style lang="scss" scoped>

:deep() {
  .grid-content {
    height: 100%;
  }
  .box-card {
    height: 100%;
    .el-card__body {
      height: calc(100% - 56px);
    }

    .table-system-info,.table-system-cpu {
      height: 100%;
      .el-descriptions__body {
        height: 100%;
      }
      table {
        height: 100%;
        td {
          padding: 0 !important;
        }
      }
    }

    .table-system-cpu td {
      height: 25% !important;
    }
  }
}
</style>
