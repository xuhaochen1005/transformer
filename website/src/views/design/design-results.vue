<template>
  <div class="app-container">
    <aside class="page-header">
      产品设计结果
    </aside>
    <div class="filter-container">
      <el-input
        v-model="listDesignTaskResultsQuery['field_lk_p.project_name']"
        placeholder="项目名称"
        clearable
        style="width: 200px"
        class="filter-item"
      />
      <el-input
        v-model="listDesignTaskResultsQuery['field_lk_p.product_model']"
        placeholder="产品型号"
        clearable
        class="filter-item"
        style="width: 200px"
      />
      <el-input
        v-model="listDesignTaskResultsQuery['field_lk_p.serial_code']"
        placeholder="任务书编号"
        clearable
        class="filter-item"
        style="width: 200px"
      />
      <el-button
        v-waves
        type="primary"
        class="filter-item"
        icon="el-icon-search"
        @click="getDesignResultsList"
      >
        搜索
      </el-button>
      <el-button
        v-waves
        :loading="downloadLoading"
        class="filter-item"
        type="primary"
        icon="el-icon-download"
        @click="handleDownload"
      >
        导出
      </el-button>
    </div>
    <!--    <DesignResults-create-->
    <!--      v-model="showCreate"-->
    <!--      @get-DesignResults-list="getDesignResultsList"-->
    <!--    />-->
    <!--    <design-results-view-->
    <!--      v-model="showDetail"-->
    <!--      :design-results="DesignResultsDetail"-->
    <!--    />-->
    <!--    <design-results-update-->
    <!--      v-model="showUpdate"-->
    <!--      :DesignResults="DesignResultsDetail"-->
    <!--      @get-DesignResults-list="getDesignResultsList"-->
    <!--    />-->
    <el-table
      v-loading="loading"
      :data="DesignTaskResultsList"
      border
      fit
      highlight-current-row
      style="width: 100%"
      @sort-change="sortChange"
    >
      <el-table-column type="expand">
        <template #default="{row}">
          <design-results-table
            :current-design-results="row.final_design_results"
            :show-design-results-mode="'view'"
          />
        </template>
      </el-table-column>
      <el-table-column
        label="ID"
        prop="id"
        sortable="custom"
        align="center"
        width="100px"
      >
        <template #default="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="项目名称"
        prop="p.project_name"
        sortable="custom"
        align="center"
        min-width="280px"
      >
        <template #default="{row}">
          <span>{{ row.design_project.project_name }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="产品型号"
        prop="design_project.product_model"
        align="center"
        width="200px"
      >
        <template #default="{row}">
          <span>{{ row.design_project.product_model }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="任务书编号"
        prop="product_name"
        align="center"
        width="150px"
      >
        <template #default="{row}">
          <span>{{ row.design_project.serial_code }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="设计人员"
        align="center"
        width="130px"
      >
        <template #default="{row}">
          <span>{{ row.design_project.designer_user.user_name_zh }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="编制人员"
        align="center"
        width="130px"
      >
        <template #default="{row}">
          <span>{{ row.design_project.drafting_user.user_name_zh }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="审核人员"
        align="center"
        width="130px"
      >
        <template #default="{row}">
          <span>{{ row.design_project.check_user.user_name_zh }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="更新时间"
        prop="created_at"
        sortable="custom"
        align="center"
        width="200px"
      >
        <template #default="{row}">
          <span>{{ UnixTime2HumanWithStr(row.updated_at) }}</span>
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
    </el-table>
    <div style="margin-top: 20px">
      <el-pagination
        v-show="total > 0"
        v-model:currentPage="listDesignTaskResultsQuery.page"
        v-model:page-size="listDesignTaskResultsQuery.limit"
        :total="total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="getDesignResultsList"
        @current-change="getDesignResultsList"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Ref } from 'vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteDesignResults, GetDesignResults, getDesignTaskResultsList } from '@/api/design'
import {
  DesignResults,
  DesignResultStatus,
  DesignTask,
  DesignTaskQuery,
  DesignTaskResultsQuery
} from '@/model/design'
import dayjs from 'dayjs'
import { useRouter } from 'vue-router'
import { ExportDesign } from '@/api/design_results'
import DesignResultsTable from '@/views/design/components/design-results-table.vue'
import { decodeObjectBytes } from '@/utils/jsonutils'

export default defineComponent({
  components: {
    DesignResultsTable
  },
  setup() {
    const router = useRouter()
    const listDesignTaskResultsQuery : DesignTaskResultsQuery = reactive({
      'field_eq_design_results.result_status': DesignResultStatus.ResultStatusFinsihed,
      page: 1,
      limit: 10,
      order: '',
      order_by: ''
    })
    const loading = ref(false)
    const downloading = ref(false)
    const DesignTaskResultsList = ref(Array<DesignTask>())
    const total = ref(0)
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function handleExport(designResults:DesignResults, id:number) {
      downloading.value = true
      const response = await ExportDesign(id)
      if (window.navigator && window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(response.data)
      } else {
        const objectUrl = URL.createObjectURL(new Blob([response.data]))
        const link = document.createElement('a')
        link.style.display = 'none'
        link.href = objectUrl
        let productName = ''
        productName = designResults.product_name + '.xlsx'
        link.download = productName
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(objectUrl)
      }
      downloading.value = false
    }
    async function getDesignResultsList() {
      loading.value = true
      const response = await getDesignTaskResultsList(listDesignTaskResultsQuery)

      if (response.data.code === 200) {
        total.value = response.data.total
        decodeObjectBytes(response.data.spec)
        response.data.spec.forEach(item => {
          item.final_design_results.forEach((iitem) => {
            iitem.design_task = JSON.parse(JSON.stringify(item))
          })
        })
        DesignTaskResultsList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    function sortChange(column: any) {
      listDesignTaskResultsQuery.order_by = column.prop
      listDesignTaskResultsQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getDesignResultsList()
    }
    onMounted(() => {
      getDesignResultsList()
    })
    function deleteDesignResults(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteDesignResults(id)
            if (response.data.code === 200) {
              ElMessage.success('记录删除成功!')
              await getDesignResultsList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const showDetail = ref(false)
    const showUpdate = ref(false)

    return {
      handleExport,
      downloading,
      status,
      loading,
      DesignTaskResultsList,
      listDesignTaskResultsQuery,
      total,
      UnixTime2Human,
      getDesignResultsList,
      sortChange,
      deleteDesignResults,
      showCreate,
      showDetail,
      showUpdate,
      UnixTime2HumanWithStr
    }
  }
})
</script>
<style lang="scss" scoped>
.filter-container {
  .container-row:first-child {
    margin-bottom: 10px;
  }
  .filter-item:not(button) {
    margin-right: 10px;
  }
}
:deep {
  .project-table-expand {
    &.el-form {
      width: 1100px;
    }
    .el-form-item {
      width: 100%;
    }
    input,textarea,.el-input,.el-radio {
          cursor: text;
          pointer-events: none;
          border: 0 !important;
    }
    .el-select {
      pointer-events: none;
    }
    .el-input__suffix {
      display: none;
    }
    .el-select,.el-input,.el-input-number {
        width: 100% !important;
        input {
          text-align: left !important;
          padding-left: 30px !important;
        }
      }
      .el-input-number__decrease {
        display: none;
      }
      .el-input-number__increase  {
        display: none;
      }
      .el-form-item__label {
        width: 140px !important;
      }

  }

}
</style>
