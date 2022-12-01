<template>
  <div>
    <aside class="page-header">
      箔绕扁形模具台账
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLfomlQuery.field_eq_mod_size"
          placeholder="模具尺寸(mm)"
          clearable
          style="width: 130px;; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLfomlQuery.field_eq_mod_height"
          placeholder="模具高度(mm)"
          clearable
          style="width: 130px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLfomlQuery.field_eq_mod_pic"
          placeholder="模具图号"
          clearable
          style="width: 130px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLfomlQuery.field_eq_mod_num"
          placeholder="模具编号"
          clearable
          style="width: 130px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLfomlQuery.field_eq_mod_suit"
          placeholder="适用产品型号"
          clearable
          style="width: 130px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLfomlQuery.field_eq_remark"
          placeholder="备注"
          clearable
          style="width: 100px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLfomlQuery.field_eq_lfoml_creator"
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
          style="margin-left:20px;margin-top: 10px"
          type="primary"
          icon="el-icon-search"
          @click="getLfomlList"
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
      <lfoml-create
        v-model="showCreate"
        @get-lfoml-list="getLfomlList"
      />
      <lfoml-view
        v-model="showDetail"
        :lfoml="lfomlDetail"
      />
      <lfoml-update
        v-model="showUpdate"
        :lfoml="lfomlDetail"
        @get-lfoml-list="getLfomlList"
      />
      <el-table
        v-loading="loading"
        :data="lfomlList"
        border
        fit
        highlight-current-row
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
          label="模具尺寸(mm)"
          prop="mode_size"
          sortable="custom"
          align="center"
          min-width="140px"
        >
          <template #default="{row}">
            <span>{{ row.mod_size }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具高度(mm)"
          prop="mod_height"
          sortable="custom"
          align="center"
          min-width="160px"
        >
          <template #default="{row}">
            <span>{{ row.mod_height }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具图号"
          prop="mod_pic"
          sortable="custom"
          align="center"
          min-width="160px"
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
          min-width="140px"
        >
          <template #default="{row}">
            <span>{{ row.mod_num }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="模具数量(个)"
          prop="mod_amount"
          sortable="custom"
          align="center"
          min-width="140px"
        >
          <template #default="{row}">
            <span>{{ row.mod_amount }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="适用产品型号"
          prop="mod_suit"
          sortable="custom"
          align="center"
          min-width="220px"
        >
          <template #default="{row}">
            <span>{{ row.mod_suit }}</span>
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
          prop="lfoml_creator"
          sortable="custom"
          align="center"
          width="120px"
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
          width="180px"
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
          width="180px"
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
              @click="showLfomlDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLfomlUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLfoml(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLfomlQuery.page"
          v-model:page-size="listLfomlQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLfomlList"
          @current-change="getLfomlList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlfomlLibraries, GetStdlfomlLibraries, ListLfomlQuery, Lfoml, ExportStdlfomlLibraries } from '@/api/standard_libraries/foml'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LfomlCreate from './lfomlCreate.vue'
import LfomlView from './lfomlView.vue'
import LfomlUpdate from './lfomlUpdate.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import dayjs from 'dayjs'

export default defineComponent({
  components: {
    UserSearchInput,
    LfomlCreate,
    LfomlView,
    LfomlUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLfomlQuery = reactive({
      field_eq_mod_size: '',
      field_eq_mod_height: null,
      field_eq_mod_pic: '',
      field_eq_mod_num: '',
      field_eq_mod_amount: null,
      field_eq_mod_suit: '',
      field_eq_remark: '',
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
    const lfomlList = ref(Array<Lfoml>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlfomlLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '箔绕扁形模具台账.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '箔绕扁形模具台账.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLfomlList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLfomlQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLfomlQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLfomlQuery.field_gt_created_at = 0
        listLfomlQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLfomlQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLfomlQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLfomlQuery.field_gt_updated_at = 0
        listLfomlQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlfomlLibraries(listLfomlQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lfomlList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLfomlQuery.order_by = column.prop
      listLfomlQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLfomlList()
    }
    onMounted(() => {
      getLfomlList()
    })
    function deleteLfoml(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlfomlLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLfomlList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lfoml: Lfoml = {
      id: 0,
      mod_size: '',
      mod_height: null,
      mod_pic: '',
      mod_num: '',
      mod_amount: null,
      mod_suit: '',
      remark: ''
    }
    const lfomlDetail = ref(lfoml)
    function showLfomlDetail(lfoml: Lfoml) {
      showDetail.value = true
      lfomlDetail.value = lfoml
    }
    function showLfomlUpdate(lfoml: Lfoml) {
      showUpdate.value = true
      lfomlDetail.value = lfoml
    }
    return {
      handleDownload,
      status,
      loading,
      lfomlList,
      LfomlCreate,
      created_at_range,
      updated_at_range,
      listLfomlQuery,
      total,
      UnixTime2Human,
      getLfomlList,
      sortChange,
      deleteLfoml,
      showCreate,
      showDetail,
      showUpdate,
      lfomlDetail,
      showLfomlUpdate,
      showLfomlDetail,
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
