<template>
  <div>
    <aside class="page-header">
      额定频率
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLrfQuery.field_eq_rated_freq"
          filterable
          allow-create
          placeholder="额定频率（Hz）"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        >
          <el-option
            v-for="item in ratedFreq"
            :key="item"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <UserSearchInput
          v-model="listLrfQuery.field_eq_lrf_creator"
          style="width: 180px"
          size="small"
          class="filter-item"
          :placeholder="'创建者'"
          clearable
        />
        <el-date-picker
          v-model="created_at_range"
          unlink-panels
          style="margin-top: 10px;width:280px;margin-left:10px"
          type="datetimerange"
          size="small"
          start-placeholder="创建起始日期"
          end-placeholder="创建截止日期"
        />
        <el-date-picker
          v-model="updated_at_range"
          unlink-panels
          style="margin-top: 10px;width:280px;margin-left:10px"
          type="datetimerange"
          size="small"
          start-placeholder="更新起始日期"
          end-placeholder="更新截止日期"
        />
        <el-button
          v-waves
          style="margin-left:20px"
          type="primary"
          icon="el-icon-search"
          @click="getLrfList"
        >
          搜索
        </el-button>
        <el-button
          v-waves
          type="primary"
          icon="el-icon-edit"
          @click="showCreate = true"
        >
          添加
        </el-button>
        <el-button
          v-waves
          :loading="loading"
          class="filter-item"
          type="primary"
          icon="el-icon-download"
          @click="handleDownload"
        >
          导出
        </el-button>
      </div>
      <Lrf-create
        v-model="showCreate"
        @get-Lrf-list="getLrfList"
      />
      <Lrf-view
        v-model="showDetail"
        :lrf="LrfDetail"
      />
      <Lrf-update
        v-model="showUpdate"
        :lrf="LrfDetail"
        @get-Lrf-list="getLrfList"
      />
      <el-table
        v-loading="loading"
        :data="LrfList"
        border
        fit
        highlight-current-row
        style="width: 100%"
        @sort-change="sortChange"
      >
        <el-table-column
          label="编号"
          prop="id"
          sortable="custom"
          align="center"
          min-width="80px"
        >
          <template #default="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定频率（Hz）"
          prop="mod_diameter"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.rated_freq }}Hz</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lrf_creator"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.creator_user?.user_name_zh }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建时间"
          prop="created_at"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ UnixTime2HumanWithStr(row.created_at) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="更新时间"
          prop="updated_at"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ UnixTime2HumanWithStr(row.updated_at) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          align="center"
          fixed="right"
          width="250px"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="success"
              @click="showLrfUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLrf(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLrfQuery.page"
          v-model:page-size="listLrfQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLrfList"
          @current-change="getLrfList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, onMounted, Ref} from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLrfLibraries, GetStdLrfLibraries, ListLrfQuery, Lrf, ExportStdlrfLibraries } from '@/api/standard_libraries/lrf'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LrfCreate from './lrfCreate.vue'
import LrfView from './lrfView.vue'
import LrfUpdate from './lrfUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import {GetStdLmadLibraries, Lmad} from "@/api/standard_libraries/lmad";

export default defineComponent({
  components: {
    UserSearchInput,
    LrfCreate,
    LrfView,
    LrfUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLrfQuery = reactive({
      field_eq_rated_freq: null,
      field_gt_created_at: 0,
      field_lt_created_at: 0,
      field_gt_updated_at: 0,
      field_lt_updated_at: 0,
      page: 1,
      limit: 10,
      order: '',
      order_by: ''
    })
    const loading = ref(false)
    const LrfList = ref(Array<Lrf>())
    const total = ref(0)
    const ratedFreq : Ref<{label: number, value: number}[]> = ref([])
    const getStdLrfLibraries = async () => {
      const res = await GetStdLrfLibraries({
        field_eq_rated_freq: 0,
        limit: 9999,
        page: 1,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        ratedFreq.value = res.data.spec.map((i: Lrf) => {
          return {
            label: i.rated_freq + 'Hz',
            value: i.rated_freq
          }
        })
      }
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlrfLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '额定频率.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '额定频率.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function getLrfList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLrfQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLrfQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLrfQuery.field_gt_created_at = 0
        listLrfQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLrfQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLrfQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLrfQuery.field_gt_updated_at = 0
        listLrfQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLrfLibraries(listLrfQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        LrfList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLrfQuery.order_by = column.prop
      listLrfQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLrfList()
    }
    onMounted(() => {
      getLrfList()
    })
    function deleteLrf(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLrfLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLrfList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lrf: Lrf = {
      id: 0,
      rated_freq: 0,
      created_at: '',
      updated_at: '',
      deleted_at: ''
    }
    const LrfDetail = ref(lrf)
    function showLrfDetail(Lrf: Lrf) {
      showDetail.value = true
      LrfDetail.value = Lrf
    }
    function showLrfUpdate(Lrf: Lrf) {
      showUpdate.value = true
      LrfDetail.value = Lrf
    }
    const paramLoading = ref(true)
    onMounted(async () => {
      await Promise.all([

        getStdLrfLibraries()
      ])
      paramLoading.value = false
    })
    return {
      status,
      handleDownload,
      loading,
      LrfList,
      LrfCreate,
      ratedFreq,
      created_at_range,
      updated_at_range,
      listLrfQuery,
      total,
      getLrfList,
      sortChange,
      deleteLrf,
      showCreate,
      showDetail,
      showUpdate,
      LrfDetail,
      showLrfUpdate,
      showLrfDetail,
      UnixTime2Human,
      UnixTime2HumanWithStr
    }
  }
})
</script>
<style lang="scss" scoped>
  .page-header {
    line-height: 24px;
    font-size: 24px;
    color: #409eff;
    background: none;
    padding: 0;
    margin:20px;
  }
</style>
