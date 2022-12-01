<template>
  <div>
    <aside class="page-header">
      变压器用途
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLusQuery.field_eq_usage"
          placeholder="变压器用途"
          filterable
          allow-create
          clearable
          style="width: 250px; margin-right: 10px"
          size="small"
        >
          <el-option
            v-for="item in useFor"
            :key="item"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-select
          v-model="listLusQuery.field_eq_usage_sign"
          placeholder="变压器用途代号"
          filterable
          allow-create
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        >
          <el-option
            v-for="item in usagesSign"
            :key="item"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <UserSearchInput
          v-model="listLusQuery.field_eq_lus_creator"
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
          @click="getLusList"
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
      <lus-create
        v-model="showCreate"
        @get-lus-list="getLusList"
      />
      <lus-view
        v-model="showDetail"
        :lus="lusDetail"
      />
      <lus-update
        v-model="showUpdate"
        :lus="lusDetail"
        @get-lus-list="getLusList"
      />
      <el-table
        v-loading="loading"
        :data="lusList"
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
          width="75px"
        >
          <template #default="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="变压器用途"
          prop="usage"
          sortable="custom"
          align="center"
          min-width="250px"
        >
          <template #default="{row}">
            <span>{{ row.usage }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="变压器用途代号"
          prop="usage_sign"
          sortable="custom"
          align="center"
          min-width="250px"
        >
          <template #default="{row}">
            <span>{{ row.usage_sign }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lus_creator"
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
          width="200px"
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
          width="200px"
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
              @click="showLusUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLus(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLusQuery.page"
          v-model:page-size="listLusQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLusList"
          @current-change="getLusList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, onMounted, Ref} from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLusLibraries, GetStdLusLibraries, ListLusQuery, Lus, ExportStdlusLibraries } from '@/api/standard_libraries/lus'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LusCreate from './lusCreate.vue'
import LusView from './lusView.vue'
import LusUpdate from './lusUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import {GetStdLrrLibraries, Lrr} from "@/api/standard_libraries/lrr";

export default defineComponent({
  components: {
    UserSearchInput,
    LusCreate,
    LusView,
    LusUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLusQuery = reactive({
      field_eq_usage: null,
      field_eq_usage_sign: null,
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
    const lusList = ref(Array<Lus>())
    const total = ref(0)
    const useFor : Ref<{label: string, value: string}[]> = ref([])
    const usagesSign : Ref<{label: string, value: string}[]> = ref([])

    const getStdLusLibraries = async () => {
      const res = await GetStdLusLibraries({
        field_eq_usage: '',
        field_eq_usage_sign: '',
        limit: 9999,
        page: 1,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        useFor.value = res.data.spec.map((i: Lus) => {
          return {
            label: i.usage,
            value: i.usage
          }
        })
        usagesSign.value = res.data.spec.map((i: Lus) => {
          return {
            label: i.usage_sign,
            value: i.usage_sign
          }
        })
      }
    }

    async function handleDownload() {
      const res = await ExportStdlusLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '变压器用途.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '变压器用途.xlsx'
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
    async function getLusList() {
      loading.value = true
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLusQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLusQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLusQuery.field_gt_created_at = 0
        listLusQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLusQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLusQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLusQuery.field_gt_updated_at = 0
        listLusQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLusLibraries(listLusQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lusList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLusQuery.order_by = column.prop
      listLusQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLusList()
    }
    onMounted(() => {
      getLusList()
    })
    function deleteLus(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLusLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLusList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lus: Lus = {
      id: 0,
      usage: '',
      usage_sign: ''
    }
    const lusDetail = ref(lus)
    function showLusDetail(lus: Lus) {
      showDetail.value = true
      lusDetail.value = lus
    }
    function showLusUpdate(lus: Lus) {
      showUpdate.value = true
      lusDetail.value = lus
    }

    const paramLoading = ref(true)
    onMounted(async () => {
      await Promise.all([

        getStdLusLibraries()
      ])
      paramLoading.value = false
    })
    return {
      status,
      handleDownload,
      loading,
      lusList,
      LusCreate,
      useFor,
      usagesSign,
      created_at_range,
      updated_at_range,
      listLusQuery,
      total,
      getLusList,
      sortChange,
      deleteLus,
      showCreate,
      showDetail,
      showUpdate,
      lusDetail,
      showLusUpdate,
      showLusDetail,
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
