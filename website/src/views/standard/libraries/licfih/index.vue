<template>
  <div>
    <aside class="page-header">
      内线圈至铁心的距离
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="highVol"
          placeholder="额定高压(kV)"
          clearable
          style="width: 200px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLicfihQuery.field_eq_inner_coil_form_iron_heart_min"
          placeholder="内线圈至铁芯距离最小值(mm)"
          clearable
          style="width: 240px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLicfihQuery.field_eq_licfih_creator"
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
          @click="getLicfihList"
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
      <licfih-create
        v-model="showCreate"
        @get-licfih-list="getLicfihList"
      />
      <licfih-view
        v-model="showDetail"
        :licfih="licfihDetail"
      />
      <licfih-update
        v-model="showUpdate"
        :licfih="licfihDetail"
        @get-licfih-list="getLicfihList"
      />
      <el-table
        v-loading="loading"
        :data="licfihList"
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
          label="额定高压下界(kV)"
          prop="rated_high_vol_min"
          sortable="custom"
          align="center"
          min-width="180px"
        >
          <template #default="{row}">
            <span>{{ row.rated_high_vol_min }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="额定高压上界(kV)"
          prop="rated_high_vol_max"
          sortable="custom"
          align="center"
          min-width="180px"
        >
          <template #default="{row}">
            <span>{{ row.rated_high_vol_max }}kV</span>
          </template>
        </el-table-column>
        <el-table-column
          label="内线圈至铁芯距离最小值(mm)"
          prop="inner_coil_form_iron_heart_min"
          sortable="custom"
          align="center"
          min-width="250px"
        >
          <template #default="{row}">
            <span>{{ row.inner_coil_form_iron_heart_min }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="licfih_creator"
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
              @click="showLicfihUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLicfih(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLicfihQuery.page"
          v-model:page-size="listLicfihQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLicfihList"
          @current-change="getLicfihList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlicfihLibraries, GetStdlicfihLibraries, ListLicfihQuery, Licfih, ExportStdlicfihLibraries } from '@/api/standard_libraries/licfih'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LicfihCreate from './licfihCreate.vue'
import LicfihView from './licfihView.vue'
import LicfihUpdate from './licfihUpdate.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import dayjs from 'dayjs'

export default defineComponent({
  components: {
    UserSearchInput,
    LicfihCreate,
    LicfihView,
    LicfihUpdate
  },
  setup() {
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const highVol: Ref<number | undefined> = ref(undefined)
    const listLicfihQuery = reactive({
      field_eq_inner_coil_form_iron_heart_min: null,
      // field_eq_rated_high_vol_min: null,
      // field_eq_rated_high_vol_max: null,
      field_lt_rated_high_vol_min: highVol,
      field_ge_rated_high_vol_max: highVol,
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
    const licfihList = ref(Array<Licfih>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlicfihLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '内线圈至铁芯距离.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '内线圈至铁芯距离.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLicfihList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLicfihQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLicfihQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLicfihQuery.field_gt_created_at = 0
        listLicfihQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLicfihQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLicfihQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLicfihQuery.field_gt_updated_at = 0
        listLicfihQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlicfihLibraries(listLicfihQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        licfihList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLicfihQuery.order_by = column.prop
      listLicfihQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLicfihList()
    }
    onMounted(() => {
      getLicfihList()
    })
    function deleteLicfih(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlicfihLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLicfihList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const licfih: Licfih = {
      id: 0,
      inner_coil_form_iron_heart_min: null,
      rated_high_vol_min: null,
      rated_high_vol_max: null
    }
    const licfihDetail = ref(licfih)
    function showLicfihDetail(licfih: Licfih) {
      showDetail.value = true
      licfihDetail.value = licfih
    }
    function showLicfihUpdate(licfih: Licfih) {
      showUpdate.value = true
      licfihDetail.value = licfih
    }
    return {
      status,
      handleDownload,
      loading,
      licfihList,
      LicfihCreate,
      created_at_range,
      updated_at_range,
      listLicfihQuery,
      total,
      highVol,
      getLicfihList,
      sortChange,
      deleteLicfih,
      showCreate,
      showDetail,
      showUpdate,
      licfihDetail,
      showLicfihUpdate,
      showLicfihDetail,
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
