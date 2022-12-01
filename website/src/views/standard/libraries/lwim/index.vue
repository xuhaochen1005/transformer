<template>
  <div>
    <aside class="page-header">
      绕组外绝缘介质
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLwimQuery.field_eq_wind_insulate_media"
          filterable
          allow-create
          placeholder="绕组外绝缘介质"
          clearable
          style="width: 220px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="浇注式"
            value="浇注式"
          />
        </el-select>
        <el-select
          v-model="listLwimQuery.field_eq_wind_insulate_media_sign"
          filterable
          allow-create
          placeholder="绕组外绝缘介质代号"
          clearable
          style="width: 220px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="C"
            value="C"
          />
        </el-select>
        <UserSearchInput
          v-model="listLwimQuery.field_eq_lwim_creator"
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
          @click="getLwimList"
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
      <lwim-create
        v-model="showCreate"
        @get-lwim-list="getLwimList"
      />
      <lwim-view
        v-model="showDetail"
        :lwim="lwimDetail"
      />
      <lwim-update
        v-model="showUpdate"
        :lwim="lwimDetail"
        @get-lwim-list="getLwimList"
      />
      <el-table
        v-loading="loading"
        :data="lwimList"
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
          label="绕组外绝缘介质"
          prop="wind_insulate_media"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.wind_insulate_media }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="绕组外绝缘介质代号"
          prop="wind_insulate_media_sign"
          sortable="custom"
          align="center"
          min-width="250px"
        >
          <template #default="{row}">
            <span>{{ row.wind_insulate_media_sign }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lwim_creator"
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
              @click="showLwimUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLwim(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLwimQuery.page"
          v-model:page-size="listLwimQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLwimList"
          @current-change="getLwimList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLwimLibraries, GetStdLwimLibraries, ListLwimQuery, Lwim, ExportStdlwimLibraries } from '@/api/standard_libraries/lwim'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LwimCreate from './lwimCreate.vue'
import LwimView from './lwimView.vue'
import LwimUpdate from './lwimUpdate.vue'
import dayjs from 'dayjs'

export default defineComponent({
  components: {
    UserSearchInput,
    LwimCreate,
    LwimView,
    LwimUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLwimQuery = reactive({
      field_eq_wind_insulate_media: null,
      field_eq_wind_insulate_media_sign: null,
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
    const lwimList = ref(Array<Lwim>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlwimLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '绕组外绝缘介质.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '绕组外绝缘介质.xlsx'
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
    async function getLwimList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLwimQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLwimQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLwimQuery.field_gt_created_at = 0
        listLwimQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLwimQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLwimQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLwimQuery.field_gt_updated_at = 0
        listLwimQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLwimLibraries(listLwimQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lwimList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLwimQuery.order_by = column.prop
      listLwimQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLwimList()
    }
    onMounted(() => {
      getLwimList()
    })
    function deleteLwim(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLwimLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLwimList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lwim: Lwim = {
      id: 0,
      wind_insulate_media: '',
      wind_insulate_media_sign: ''
    }
    const lwimDetail = ref(lwim)
    function showLwimDetail(lwim: Lwim) {
      showDetail.value = true
      lwimDetail.value = lwim
    }
    function showLwimUpdate(lwim: Lwim) {
      showUpdate.value = true
      lwimDetail.value = lwim
    }
    return {
      status,
      handleDownload,
      loading,
      lwimList,
      LwimCreate,
      created_at_range,
      updated_at_range,
      listLwimQuery,
      total,
      UnixTime2Human,
      getLwimList,
      sortChange,
      deleteLwim,
      showCreate,
      showDetail,
      showUpdate,
      lwimDetail,
      showLwimUpdate,
      showLwimDetail,
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
