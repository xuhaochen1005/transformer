<template>
  <div>
    <aside class="page-header">
      低压外模台账
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-select
          v-model="listLlvomlQuery.field_eq_mod_type"
          filterable
          allow-create
          placeholder="模具类型"
          clearable
          style="width: 110px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="硬模"
            value="硬模"
          />
          <el-option
            label="软模"
            value="软模"
          />
          <el-option
            label="箔式浇注外模"
            value="箔式浇注外模"
          />
          <el-option
            label="水冷变浇注外模"
            value="水冷变浇注外模"
          />
        </el-select>
        <el-input
          v-model="listLlvomlQuery.field_eq_mod_diameter"
          placeholder="模具外径（mm）"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLlvomlQuery.field_eq_mod_height"
          placeholder="模具高度（mm）"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLlvomlQuery.field_eq_mod_num"
          placeholder="模具编号"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLlvomlQuery.field_eq_remark"
          placeholder="备注"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLlvomlQuery.field_eq_llvoml_creator"
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
          @click="getLlvomlList"
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
      <Llvoml-create
        v-model="showCreate"
        @get-Llvoml-list="getLlvomlList"
      />
      <Llvoml-view
        v-model="showDetail"
        :llvoml="LlvomlDetail"
      />
      <Llvoml-update
        v-model="showUpdate"
        :llvoml="LlvomlDetail"
        @get-Llvoml-list="getLlvomlList"
      />
      <el-table
        v-loading="loading"
        :data="LlvomlList"
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
          label="模具外径(mm)"
          prop="mod_diameter"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.mod_diameter }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具类型"
          prop="mod_type"
          sortable="custom"
          align="center"
          min-width="110px"
        >
          <template #default="{row}">
            <span>{{ row.mod_type }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具高度(mm)"
          prop="mod_height"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.mod_height }}mm</span>
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
          label="模具编号"
          prop="mod_num"
          sortable="custom"
          align="center"
          min-width="150x"
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
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.remark }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="llvoml_creator"
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
              type="primary"
              @click="showLlvomlDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLlvomlUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLlvoml(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLlvomlQuery.page"
          v-model:page-size="listLlvomlQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLlvomlList"
          @current-change="getLlvomlList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdLlvomlLibraries, GetStdLlvomlLibraries, ListLlvomlQuery, Llvoml, ExportStdllvomlLibraries } from '@/api/standard_libraries/llvoml'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LlvomlCreate from './llvomlCreate.vue'
import LlvomlView from './llvomlView.vue'
import LlvomlUpdate from './llvomlUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LlvomlCreate,
    LlvomlView,
    LlvomlUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLlvomlQuery = reactive({
      field_eq_mod_diameter: null,
      field_eq_mod_type: null,
      field_eq_mod_height: null,
      field_eq_mod_num: null,
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
    const LlvomlList = ref(Array<Llvoml>())
    const total = ref(0)
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdllvomlLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '低压外模台账.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '低压外模台账.xlsx'
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
    async function getLlvomlList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLlvomlQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLlvomlQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLlvomlQuery.field_gt_created_at = 0
        listLlvomlQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLlvomlQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLlvomlQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLlvomlQuery.field_gt_updated_at = 0
        listLlvomlQuery.field_lt_updated_at = 0
      }
      const response = await GetStdLlvomlLibraries(listLlvomlQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        LlvomlList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLlvomlQuery.order_by = column.prop
      listLlvomlQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLlvomlList()
    }
    onMounted(() => {
      getLlvomlList()
    })
    function deleteLlvoml(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdLlvomlLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLlvomlList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const llvoml: Llvoml = {
      id: 0,
      mod_diameter: 0,
      mod_type: '',
      mod_height: 0,
      mod_amount: 0,
      mod_num: '',
      mod_suit: '',
      remark: '',
      created_at: '',
      updated_at: '',
      deleted_at: ''
    }
    const LlvomlDetail = ref(llvoml)
    function showLlvomlDetail(Llvoml: Llvoml) {
      showDetail.value = true
      LlvomlDetail.value = Llvoml
    }
    function showLlvomlUpdate(Llvoml: Llvoml) {
      showUpdate.value = true
      LlvomlDetail.value = Llvoml
    }
    return {
      status,
      handleDownload,
      loading,
      LlvomlList,
      LlvomlCreate,
      created_at_range,
      updated_at_range,
      listLlvomlQuery,
      total,
      getLlvomlList,
      sortChange,
      deleteLlvoml,
      showCreate,
      showDetail,
      showUpdate,
      LlvomlDetail,
      showLlvomlUpdate,
      showLlvomlDetail,
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
