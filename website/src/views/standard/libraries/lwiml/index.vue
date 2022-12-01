<template>
  <div>
    <aside class="page-header">
      绕线内模台账
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLwimlQuery.field_eq_mod_diameter"
          placeholder="模具外径（mm）"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLwimlQuery.field_eq_mod_num"
          placeholder="模具编号"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLwimlQuery.field_eq_mod_pic"
          placeholder="模具图号"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLwimlQuery.field_eq_remark"
          placeholder="备注"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLwimlQuery.field_eq_wiml_creator"
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
          @click="getLwimlList"
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
      <Lwiml-create
        v-model="showCreate"
        @get-Lwiml-list="getLwimlList"
      />
      <Lwiml-view
        v-model="showDetail"
        :lwiml="LwimlDetail"
      />
      <Lwiml-update
        v-model="showUpdate"
        :lwiml="LwimlDetail"
        @get-Lwiml-list="getLwimlList"
      />
      <el-table
        v-loading="loading"
        :data="LwimlList"
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
          label="模具直径(mm)"
          prop="mod_diameter"
          sortable="custom"
          align="center"
          min-width="180px"
        >
          <template #default="{row}">
            <span>{{ row.mod_diameter }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具规格（mm*mm）"
          prop="mod_size"
          sortable="custom"
          align="center"
          min-width="200px"
        >
          <template #default="{row}">
            <span>{{ row.mod_size }}mm*mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具数量"
          prop="mod_amount"
          sortable="custom"
          align="center"
          min-width="110px"
        >
          <template #default="{row}">
            <span>{{ row.mod_amount }}个</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具图号"
          prop="mod_pic"
          sortable="custom"
          align="center"
          min-width="130x"
        >
          <template #default="{row}">
            <span>{{ row.mod_pic }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具编号"
          prop="mod_num"
          sortable="custom"
          align="center"
          min-width="130x"
        >
          <template #default="{row}">
            <span>{{ row.mod_num }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="备注"
          prop="remark"
          sortable="custom"
          align="center"
          min-width="130px"
        >
          <template #default="{row}">
            <span>{{ row.remark }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lwiml_creator"
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
            <span>{{ UnixTime2HumanWithStr( row.created_at) }}</span>
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
              type="primary"
              @click="showLwimlDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLwimlUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLwiml(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLwimlQuery.page"
          v-model:page-size="listLwimlQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLwimlList"
          @current-change="getLwimlList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import { DeleteStdLwimlLibraries, GetStdLwimlLibraries, ListLwimlQuery, Lwiml, ExportStdlwimlLibraries } from '@/api/standard_libraries/lwiml'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LwimlCreate from './lwimlCreate.vue'
import LwimlView from './lwimlView.vue'
import LwimlUpdate from './lwimlUpdate.vue'
import dayjs from 'dayjs'

export default defineComponent({
  components: {
    UserSearchInput,
    LwimlCreate,
    LwimlView,
    LwimlUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLwimlQuery = reactive({
      field_eq_mod_diameter: null,
      field_eq_mod_size: null,
      field_eq_mod_amount: null,
      field_eq_mod_pic: null,
      field_eq_mod_num: null,
      field_eq_mod_suit: null,
      field_eq_remark: null,
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
    const LwimlList = ref(Array<Lwiml>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlwimlLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '绕线内模台账.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '绕线内模台账.xlsx'
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
    async function getLwimlList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLwimlQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLwimlQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLwimlQuery.field_gt_created_at = 0
        listLwimlQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLwimlQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLwimlQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLwimlQuery.field_gt_updated_at = 0
        listLwimlQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLwimlLibraries(listLwimlQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        LwimlList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLwimlQuery.order_by = column.prop
      listLwimlQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLwimlList()
    }
    onMounted(() => {
      getLwimlList()
    })
    function deleteLwiml(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLwimlLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLwimlList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lwiml: Lwiml = {
      id: 0,
      mod_diameter: 0,
      mod_size: '',
      mod_pic: '',
      mod_amount: 0,
      mod_num: '',
      mod_suit: '',
      remark: '',
      created_at: '',
      updated_at: ''
    }
    const LwimlDetail = ref(lwiml)
    function showLwimlDetail(Lwiml: Lwiml) {
      showDetail.value = true
      LwimlDetail.value = Lwiml
    }
    function showLwimlUpdate(Lwiml: Lwiml) {
      showUpdate.value = true
      LwimlDetail.value = Lwiml
    }
    return {
      status,
      handleDownload,
      loading,
      LwimlList,
      LwimlCreate,
      created_at_range,
      updated_at_range,
      listLwimlQuery,
      total,
      getLwimlList,
      sortChange,
      deleteLwiml,
      showCreate,
      showDetail,
      showUpdate,
      LwimlDetail,
      showLwimlUpdate,
      showLwimlDetail,
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
