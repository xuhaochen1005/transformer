<template>
  <div>
    <aside class="page-header">
      铁心表(长圆形)
    </aside>
    <div style="margin: 10px">
      <div style="margin-bottom: 10px">
        <el-input
          v-model="listLicoQuery.field_eq_core_diameter"
          placeholder="铁心直径(mm)"
          clearable
          style="width: 150px;; margin-right: 10px"
          size="small"
        />
        <el-select
          v-model="listLicoQuery.field_eq_pull_plate"
          placeholder="有无拉板"
          clearable
          style="width: 130px; margin-right: 10px"
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
          v-model="listLicoQuery.field_lk_section_area"
          placeholder="截面积(mm^2)"
          clearable
          style="width: 150px; margin-right: 10px"
          size="small"
        />
        <el-input
          v-model="listLicoQuery.field_eq_iron_corner_weight"
          placeholder="角重kg)"
          clearable
          style="width: 130px; margin-right: 10px"
          size="small"
        />
        <UserSearchInput
          v-model="listLicoQuery.field_eq_lico_creator"
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
          @click="getLicoList"
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
      <lico-create
        v-model="showCreate"
        @get-lico-list="getLicoList"
      />
      <lico-view
        v-model="showDetail"
        :lico="licoDetail"
      />
      <lico-update
        v-model="showUpdate"
        :lico="licoDetail"
        @get-lico-list="getLicoList"
      />
      <el-table
        v-loading="loading"
        :data="licoList"
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
          width="120px"
        >
          <template #default="{row}">
            <span>{{ row.iron_corner_weight }}kg</span>
          </template>
        </el-table-column>
        <el-table-column
          label="片宽(0级)(mm)"
          prop="stack_width_0"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_width_0 }}mm</span>
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
          label="迭厚(0级)(mm)"
          prop="stack_thick_0"
          sortable="custom"
          align="center"
          width="150px"
        >
          <template #default="{row}">
            <span>{{ row.stack_thick_0 }}mm</span>
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
          width="150px"
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
          prop="lico_creator"
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
          width="340px"
        >
          <template #default="{row}">
            <el-button
              size="mini"
              type="primary"
              @click="showLicoDetail(row)"
            >
              查看详情
            </el-button>
            <el-button
              size="mini"
              type="success"
              @click="showLicoUpdate(row)"
            >
              修改
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteLico(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listLicoQuery.page"
          v-model:page-size="listLicoQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getLicoList"
          @current-change="getLicoList"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteStdlicoLibraries, GetStdlicoLibraries, ListLicoQuery, Lico, LicoPullPlate, ExportStdlicoLibraries } from '@/api/standard_libraries/lico'
import UserSearchInput from '@/components/user-search-input/index.vue'
import LicoCreate from './licoCreate.vue'
import LicoView from './licoView.vue'
import LicoUpdate from './licoUpdate.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import dayjs from 'dayjs'

export default defineComponent({
  components: {
    UserSearchInput,
    LicoCreate,
    LicoView,
    LicoUpdate
  },
  setup() {
    const pullplate = new Map([[2, '无'], [1, '有']])
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listLicoQuery = reactive({
      field_eq_core_diameter: null,
      field_lk_section_area: null,
      field_eq_iron_corner_weight: null,
      field_eq_stack_width_0: null,
      field_eq_stack_width_1: null,
      field_eq_stack_width_2: null,
      field_eq_stack_width_3: null,
      field_eq_stack_width_4: null,
      field_eq_stack_width_5: null,
      field_eq_stack_width_6: null,
      field_eq_stack_width_7: null,
      field_eq_stack_width_8: null,
      field_eq_stack_thick_0: null,
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
    const licoList = ref(Array<Lico>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleDownload() {
      loading.value = true
      const res = await ExportStdlicoLibraries()
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(res.data, '铁心表（长圆形）.xlsx')
      } else {
        const objectUrl = URL.createObjectURL(res.data)
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        link.download = '铁心表（长圆形）.xlsx'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      loading.value = false
    }
    async function getLicoList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listLicoQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listLicoQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listLicoQuery.field_gt_created_at = 0
        listLicoQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listLicoQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listLicoQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listLicoQuery.field_gt_updated_at = 0
        listLicoQuery.field_lt_updated_at = 0
      }
      const response = await GetStdlicoLibraries(listLicoQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        licoList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listLicoQuery.order_by = column.prop
      listLicoQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getLicoList()
    }
    onMounted(() => {
      getLicoList()
    })
    function deleteLico(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteStdlicoLibraries(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getLicoList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)
    const lico: Lico = {
      id: 0,
      core_diameter: null,
      pull_plate: LicoPullPlate.Yes,
      section_area: null,
      iron_corner_weight: null,
      stack_width_0: null,
      stack_width_1: null,
      stack_width_2: null,
      stack_width_3: null,
      stack_width_4: null,
      stack_width_5: null,
      stack_width_6: null,
      stack_width_7: null,
      stack_width_8: null,
      stack_thick_0: null,
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
    const licoDetail = ref(lico)
    function showLicoDetail(lico: Lico) {
      showDetail.value = true
      licoDetail.value = lico
    }
    function showLicoUpdate(lico: Lico) {
      showUpdate.value = true
      licoDetail.value = lico
    }
    return {
      pullplate,
      status,
      handleDownload,
      loading,
      licoList,
      LicoCreate,
      created_at_range,
      updated_at_range,
      listLicoQuery,
      total,
      getLicoList,
      sortChange,
      deleteLico,
      showCreate,
      showDetail,
      showUpdate,
      licoDetail,
      showLicoUpdate,
      showLicoDetail,
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
