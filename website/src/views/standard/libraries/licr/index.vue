<template>
  <div>
    <aside class="page-header">
      铁心表(圆形)
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLicrQuery.field_eq_core_diameter"
          placeholder="铁心直径(mm)"
          clearable
          style="width: 150px;; margin-right: 10px"
          size="small"
        />
        <el-select
          v-model="listLicrQuery.field_eq_pull_plate"
          placeholder="有无拉板"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        >
          <el-option
            label="有"
            value="1"
          />
          <el-option
            label="无"
            value="2"
          />
        </el-select>
        <el-input
          v-model="listLicrQuery.field_eq_section_area"
          placeholder="截面积(mm^2)"
          clearable
          style="width: 160px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLicrQuery.field_eq_iron_corner_weight"
          placeholder="角重(kg)"
          clearable
          style="width: 160px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLicrQuery.field_eq_licr_creator"
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
          @click="getLicrList"
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
      <licr-create
        v-model="showCreate"
        @get-licr-list="getLicrList"
      />
      <licr-view
        v-model="showDetail"
        :licr="licrDetail"
      />
      <licr-update
        v-model="showUpdate"
        :licr="licrDetail"
        @get-licr-list="getLicrList"
      />
      <el-table
        v-loading="loading"
        :data="licrList"
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
          label="铁心直径(mm)"
          prop="core_diameter"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.core_diameter }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="有无拉板"
          prop="pull_plate"
          sortable="custom"
          align="center"
          width="120px"
        >
          <template #default="{row}">
            <span>{{ pullplate.get(row.pull_plate) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="截面积(mm^2)"
          prop="section_area"
          sortable="custom"
          align="center"
          width="160px"
        >
          <template #default="{row}">
            <span>{{ row.section_area }}mm^2</span>
          </template>
        </el-table-column>
        <el-table-column
          label="角重(kg)"
          prop="iron_corner_weight"
          sortable="custom"
          align="center"
          width="130px"
        >
          <template #default="{row}">
            <span>{{ row.iron_corner_weight }}kg</span>
          </template>
        </el-table-column>
        <el-table-column
          label="片宽(1级)(mm)"
          prop="stack_width_1"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_width_1 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="片宽(2级)(mm)"
          prop="stack_width_2"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_width_2 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="片宽(3级)(mm)"
          prop="stack_width_3"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_width_3 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="片宽(4级)(mm)"
          prop="stack_width_4"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_width_4 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="片宽(5级)(mm)"
          prop="stack_width_5"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_width_5 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="片宽(6级)(mm)"
          prop="stack_width_6"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_width_6 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="片宽(7级)(mm)"
          prop="stack_width_7"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_width_7 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="片宽(8级)(mm)"
          prop="stack_width_8"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_width_8 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="迭厚(1级)(mm)"
          prop="stack_thick_1"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_thick_1 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="迭厚(2级)(mm)"
          prop="stack_thick_2"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_thick_2 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="迭厚(3级)(mm)"
          prop="stack_thick_3"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_thick_3 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="迭厚(4级)(mm)"
          prop="stack_thick_4"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_thick_4 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="迭厚(5级)(mm)"
          prop="stack_thick_5"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_thick_5 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="迭厚(6级)(mm)"
          prop="stack_thick_6"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_thick_6 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="迭厚(7级)(mm)"
          prop="stack_thick_7"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_thick_7 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="迭厚(8级)(mm)"
          prop="stack_thick_8"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_thick_8 }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="总迭厚(mm)"
          prop="stack_thick_sum"
          sortable="custom"
          align="center"
          width="200px"
        >
          <template #default="{row}">
            <span>{{ row.stack_thick_sum }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="主级实迭厚迭厚(mm)"
          prop="main_level_real_stack_thick"
          sortable="custom"
          align="center"
          width="200px"
        >
          <template #default="{row}">
            <span>{{ row.main_level_real_stack_thick }}mm</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建者"
          prop="licr_creator"
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
              @click="showLicrDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLicrUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLicr(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLicrQuery.page"
          v-model:page-size="listLicrQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLicrList"
          @current-change="getLicrList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlicrLibraries, GetStdlicrLibraries, ListLicrQuery, Licr, LicrPullPlate, ExportStdlicrLibraries } from '@/api/standard_libraries/licr'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LicrCreate from './licrCreate.vue'
import LicrView from './licrView.vue'
import LicrUpdate from './licrUpdate.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import dayjs from 'dayjs'

export default defineComponent({
  components: {
    UserSearchInput,
    LicrCreate,
    LicrView,
    LicrUpdate
  },
  setup() {
    const pullplate = new Map([[2, '无'], [1, '有']])
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLicrQuery = reactive({
      field_eq_core_diameter: null,
      field_eq_section_area: null,
      field_eq_iron_corner_weight: null,
      field_eq_stack_width_1: null,
      field_eq_stack_width_2: null,
      field_eq_stack_width_3: null,
      field_eq_stack_width_4: null,
      field_eq_stack_width_5: null,
      field_eq_stack_width_6: null,
      field_eq_stack_width_7: null,
      field_eq_stack_width_8: null,
      field_eq_stack_thick_1: null,
      field_eq_stack_thick_2: null,
      field_eq_stack_thick_3: null,
      field_eq_stack_thick_4: null,
      field_eq_stack_thick_5: null,
      field_eq_stack_thick_6: null,
      field_eq_stack_thick_7: null,
      field_eq_stack_thick_8: null,
      field_eq_stack_thick_sum: null,
      field_eq_main_level_real_stack_thick: null,
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
    const licrList = ref(Array<Licr>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlicrLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '铁心表（圆形）.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '铁心表（圆形）.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLicrList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLicrQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLicrQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLicrQuery.field_gt_created_at = 0
        listLicrQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLicrQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLicrQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLicrQuery.field_gt_updated_at = 0
        listLicrQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlicrLibraries(listLicrQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        licrList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLicrQuery.order_by = column.prop
      listLicrQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLicrList()
    }
    onMounted(() => {
      getLicrList()
    })
    function deleteLicr(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlicrLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLicrList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const licr: Licr = {
      id: 0,
      core_diameter: null,
      pull_plate: LicrPullPlate.Yes,
      section_area: null,
      iron_corner_weight: null,
      stack_width_1: null,
      stack_width_2: null,
      stack_width_3: null,
      stack_width_4: null,
      stack_width_5: null,
      stack_width_6: null,
      stack_width_7: null,
      stack_width_8: null,
      stack_thick_1: null,
      stack_thick_2: null,
      stack_thick_3: null,
      stack_thick_4: null,
      stack_thick_5: null,
      stack_thick_6: null,
      stack_thick_7: null,
      stack_thick_8: null,
      stack_thick_sum: null,
      main_level_real_stack_thick: null
    }
    const licrDetail = ref(licr)
    function showLicrDetail(licr: Licr) {
      showDetail.value = true
      licrDetail.value = licr
    }
    function showLicrUpdate(lico: Licr) {
      showUpdate.value = true
      licrDetail.value = lico
    }
    return {
      pullplate,
      status,
      handleDownload,
      loading,
      licrList,
      LicrCreate,
      created_at_range,
      updated_at_range,
      listLicrQuery,
      total,
      getLicrList,
      sortChange,
      deleteLicr,
      showCreate,
      showDetail,
      showUpdate,
      licrDetail,
      showLicrUpdate,
      showLicrDetail,
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
