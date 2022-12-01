<template>
  <div>
    <aside class="page-header">
      线圈形状
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLcsQuery.field_eq_coil_shape"
          filterable
          allow-create
          clearable
          style="width: 200px; margin-right: 10px"
          placeholder="线圈形状"
          size="small"
        >
          <el-option
            label="圆形"
            value="圆形"
          />
          <el-option
            label="长圆形"
            value="长圆形"
          />
        </el-select>
        <UserSearchInput
          v-model="listLcsQuery.field_eq_lcs_creator"
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
          @click="getLcsList"
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
      <lcs-create
        v-model="showCreate"
        @get-lcs-list="getLcsList"
      />
      <lcs-view
        v-model="showDetail"
        :lcs="lcsDetail"
      />
      <lcs-update
        v-model="showUpdate"
        :lcs="lcsDetail"
        @get-lcs-list="getLcsList"
      />
      <el-table
        v-loading="loading"
        :data="lcsList"
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
          label="线圈形状"
          prop="coil_shape"
          sortable="custom"
          align="center"
        >
          <template #default="{row}">
            <span>{{ row.coil_shape }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lcs_creator"
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
              @click="showLcsUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLcs(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLcsQuery.page"
          v-model:page-size="listLcsQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLcsList"
          @current-change="getLcsList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlcsLibraries, ExportStdlcsLibraries, GetStdlcsLibraries, Lcs } from '@/api/standard_libraries/lcs'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LcsCreate from './lcsCreate.vue'
import LcsView from './lcsView.vue'
import LcsUpdate from './lcsUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LcsCreate,
    LcsView,
    LcsUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLcsQuery = reactive({
      field_eq_coil_shape: '',
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
    const lcsList = ref(Array<Lcs>())
    const total = ref(0)

    function change(cs: string) {
      listLcsQuery.field_eq_coil_shape = cs
    }

    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }

    async function handleDownload() {
      const res = await ExportStdlcsLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '线圈形状.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '线圈形状.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }

    async function getLcsList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLcsQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLcsQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLcsQuery.field_gt_created_at = 0
        listLcsQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLcsQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLcsQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLcsQuery.field_gt_updated_at = 0
        listLcsQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlcsLibraries(listLcsQuery)
      if (response.data.code === 200) {
        // console.log(response.data)
        total.value = response.data.total
        lcsList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }

    function sortChange(column: any) {
      listLcsQuery.order_by = column.prop
      listLcsQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLcsList()
    }

    onMounted(() => {
      getLcsList()
    })

    function deleteLcs(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlcsLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLcsList()
            }
          }
        }
      })
    }

    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lcs: Lcs = {
      id: 0,
      coil_shape: ''
    }
    const lcsDetail = ref(lcs)

    function showLcsDetail(lcs: Lcs) {
      showDetail.value = true
      lcsDetail.value = lcs
    }

    function showLcsUpdate(lcc: Lcs) {
      showUpdate.value = true
      lcsDetail.value = lcc
    }

    return {
      handleDownload,
      change,
      status,
      loading,
      lcsList,
      LcsCreate,
      created_at_range,
      updated_at_range,
      listLcsQuery,
      total,
      getLcsList,
      sortChange,
      deleteLcs,
      showCreate,
      showDetail,
      showUpdate,
      lcsDetail,
      showLcsUpdate,
      showLcsDetail,
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
