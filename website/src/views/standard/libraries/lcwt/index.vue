<template>
  <div>
    <aside class="page-header">
      线圈导线类型
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="select.coil_wire_type"
          filterable
          allow-create
          clearable
          style="width: 200px; margin-right: 10px"
          placeholder="线圈导线类型"
          @focus="Focus"
          @change="cwtChange(select.coil_wire_type)"
          size="small"
        >
          <el-option
            v-for="item in select"
            :key="item.id"
            :lable="item.coil_wire_type"
            :value="item.coil_wire_type"
          />
        </el-select>
        <el-select
          v-model="select.coil_wire_type_sign"
          filterable
          allow-create
          clearable
          style="width: 200px; margin-right: 10px"
          placeholder="线圈导线类型字母代号"
          @focus="Focus"
          @change="cwtsChange(select.coil_wire_type_sign)"
          size="small"
        >
          <el-option
            v-for="item in select"
            :key="item.id"
            :lable="item.coil_wire_type_sign"
            :value="item.coil_wire_type_sign"
          />
        </el-select>
        <UserSearchInput
          v-model="listLcwtQuery.field_eq_lcwt_creator"
          style="width: 200px"
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
          @click="getLcwtList"
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
      <lcwt-create
        v-model="showCreate"
        @get-lcwt-list="getLcwtList"
      />
      <lcwt-view
        v-model="showDetail"
        :lcwt="lcwtDetail"
      />
      <lcwt-update
        v-model="showUpdate"
        :lcwt="lcwtDetail"
        @get-lcwt-list="getLcwtList"
      />
      <el-table
        v-loading="loading"
        :data="lcwtList"
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
          label="线圈导线类型"
          prop="coil_wire_type"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.coil_wire_type }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="线圈导线类型字母代号"
          prop="coil_wire_type_sign"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.coil_wire_type_sign }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lcwt_creator"
          sortable="custom"
          align="center"
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
        >
          <template #default="{row}">
            <span>{{ UnixTime2HumanWithStr(row.updated_at) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          align="center"
          fixed="right"
          min-width="80px"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="success"
              @click="showLcwtUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLcwt(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLcwtQuery.page"
          v-model:page-size="listLcwtQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLcwtList"
          @current-change="getLcwtList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlcwtLibraries, GetStdlcwtLibraries, ListLcwtQuery, Lcwt, ExportStdlcwtLibraries } from '@/api/standard_libraries/lcwt'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LcwtCreate from './lcwtCreate.vue'
import LcwtView from './lcwtView.vue'
import LcwtUpdate from './lcwtUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LcwtCreate,
    LcwtView,
    LcwtUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const select = ref(Array<Lcwt>())
    const listLcwtQuery = reactive({
      field_eq_coil_wire_type: '',
      field_eq_coil_wire_type_sign: '',
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
    const lcwtList = ref(Array<Lcwt>())
    const total = ref(0)
    async function Focus() {
      loading.value = true
      const response = await GetStdlcwtLibraries(listLcwtQuery)
      select.value = response.data.total ? response.data.spec : []
      loading.value = false
    }
    function cwtChange(cwt:string) {
      listLcwtQuery.field_eq_coil_wire_type = cwt
    }
    function cwtsChange(cwts:string) {
      listLcwtQuery.field_eq_coil_wire_type_sign = cwts
    }
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlcwtLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '线圈导线类型.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '线圈导线类型.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLcwtList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLcwtQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLcwtQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLcwtQuery.field_gt_created_at = 0
        listLcwtQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLcwtQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLcwtQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLcwtQuery.field_gt_updated_at = 0
        listLcwtQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlcwtLibraries(listLcwtQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lcwtList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLcwtQuery.order_by = column.prop
      listLcwtQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLcwtList()
    }
    onMounted(() => {
      getLcwtList()
    })
    function deleteLcwt(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlcwtLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLcwtList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lcwt: Lcwt = {
      id: 0,
      coil_wire_type: '',
      coil_wire_type_sign: ''
    }
    const lcwtDetail = ref(lcwt)
    function showLcwtDetail(lcwt: Lcwt) {
      showDetail.value = true
      lcwtDetail.value = lcwt
    }
    function showLcwtUpdate(lcwt: Lcwt) {
      showUpdate.value = true
      lcwtDetail.value = lcwt
    }
    return {
      handleDownload,
      select,
      Focus,
      cwtChange,
      cwtsChange,
      status,
      loading,
      lcwtList,
      LcwtCreate,
      created_at_range,
      updated_at_range,
      listLcwtQuery,
      total,
      getLcwtList,
      sortChange,
      deleteLcwt,
      showCreate,
      showDetail,
      showUpdate,
      lcwtDetail,
      showLcwtUpdate,
      showLcwtDetail,
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
