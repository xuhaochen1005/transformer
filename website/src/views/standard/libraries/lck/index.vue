<template>
  <div>
    <aside class="page-header">
      铁心直径经验系数K
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLckQuery.field_eq_winding_material"
          placeholder="绕组材质"
          clearable
          style="width: 100px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="capacity"
          placeholder="额定容量(kVA)"
          clearable
          style="width: 170px;margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="kVAlue"
          placeholder="k值"
          clearable
          style="width: 100px;margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLckQuery.field_eq_lck_creator"
          style="width: 145px"
          size="small"
          class="filter-item"
          :placeholder="'创建者'"
          clearable
        />
        <el-date-picker
          v-model="created_at_range"
          unlink-panels
          style="margin-top: 10px;width:250px;margin-left:10px"
          type="datetimerange"
          size="small"
          start-placeholder="创建起始日期"
          end-placeholder="创建截止日期"
        />
        <el-date-picker
          v-model="updated_at_range"
          unlink-panels
          style="margin-top: 10px;width:250px;margin-left:10px"
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
          @click="getLckList"
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
      <lck-create
        v-model="showCreate"
        @get-lck-list="getLckList"
      />
      <lck-view
        v-model="showDetail"
        :lck="lckDetail"
      />
      <lck-update
        v-model="showUpdate"
        :lck="lckDetail"
        @get-lck-list="getLckList"
      />
      <el-table
        v-loading="loading"
        :data="lckList"
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
          label="绕组材质"
          prop="winding_material"
          sortable="custom"
          align="center"
          min-width="150px"
        >
          <template #default="{row}">
            <span>{{ row.winding_material }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定容量下限(kVA)"
          prop="rated_capacity_min"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.rated_capacity_min }}kVA</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定容量上限(kVA)"
          prop="rated_capacity_max"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.rated_capacity_max }}kVA</span>
          </template>
        </el-table-column>
        <el-table-column
          label="k值下限"
          prop="k_min"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.k_min }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="k值上限"
          :show-overflow-tooltip="true"
          prop="k_max"
          sortable="custom"
          align="center"
          min-width="170px"
        >
          <template #default="{row}">
            <span>{{ row.k_max }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="lck_creator"
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
          width="250"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="primary"
              @click="showLckDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLckUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLck(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLckQuery.page"
          v-model:page-size="listLckQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLckList"
          @current-change="getLckList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlckLibraries, GetStdlckLibraries, ListLckQuery, Lck, ExportStdlckLibraries } from '@/api/standard_libraries/lck'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LckCreate from './lckCreate.vue'
import LckView from './lckView.vue'
import LckUpdate from './lckUpdate.vue'
import dayjs from 'dayjs'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput,
    LckCreate,
    LckView,
    LckUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const capacity: Ref<number | undefined> = ref(undefined)
    const kVAlue: Ref<number | undefined> = ref(undefined)
    const listLckQuery = reactive({
      field_eq_winding_material: '',
      // field_eq_rated_capacity_min: null,
      // field_eq_rated_capacity_max: null,
      field_lt_rated_capacity_min: capacity,
      field_ge_rated_capacity_max: capacity,
      // field_eq_k_min: null,
      // field_eq_k_max: null,
      field_lt_k_min: kVAlue,
      field_ge_k_max: kVAlue,
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
    const lckList = ref(Array<Lck>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlckLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '干式变压器铁心直径经验系数K.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '干式变压器铁心直径经验系数K.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLckList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLckQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLckQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLckQuery.field_gt_created_at = 0
        listLckQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLckQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLckQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLckQuery.field_gt_updated_at = 0
        listLckQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlckLibraries(listLckQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        lckList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLckQuery.order_by = column.prop
      listLckQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLckList()
    }
    onMounted(() => {
      getLckList()
    })
    function deleteLck(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlckLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLckList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lck: Lck = {
      id: 0,
      winding_material: '',
      rated_capacity_min: null,
      rated_capacity_max: null,
      k_min: null,
      k_max: null
    }
    const lckDetail = ref(lck)
    function showLckDetail(lck: Lck) {
      showDetail.value = true
      lckDetail.value = lck
    }
    function showLckUpdate(lck: Lck) {
      showUpdate.value = true
      lckDetail.value = lck
    }
    return {
      handleDownload,
      status,
      loading,
      lckList,
      LckCreate,
      created_at_range,
      updated_at_range,
      listLckQuery,
      total,
      capacity,
      kVAlue,
      getLckList,
      sortChange,
      deleteLck,
      showCreate,
      showDetail,
      showUpdate,
      lckDetail,
      showLckUpdate,
      showLckDetail,
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
