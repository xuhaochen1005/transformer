<template>
  <div class="app-container">
    <aside class="page-header">
      {{ category }}
    </aside>
    <div class="main-table">
      <div
        class="filter-container"
      >
        <div class="container-row">
          <el-input
            v-model="listDocsQuery.field_lk_name"
            placeholder="文档名称"
            clearable
            class="filter-item"
            style="width: 150px"
          />
          <el-select
            v-model="listDocsQuery.field_eq_document_type"
            placeholder="文档类型"
            clearable
            allow-create
            style="width: 150px"
            class="filter-item"
          >
            <el-option
              label="xlsx"
              value=".xlsx"
            />
            <el-option
              label="doc"
              value=".doc"
            />
            <el-option
              label="pdf"
              value=".pdf"
            />
            <el-option
              label="图片"
              :value="['.jpg','.png','.gif','.jpeg']"
            />
          </el-select>
          <UserSearchInput
            v-model="listDocsQuery.field_eq_docs_creator"
            style="width: 180px"
            class="filter-item"
            :placeholder="'创建者'"
            clearable
          />
          <el-date-picker
            v-model="created_at_range"
            unlink-panels
            class="filter-item"
            style="width:280px"
            type="datetimerange"
            start-placeholder="创建起始日期"
            end-placeholder="创建截止日期"
          />
          <el-date-picker
            v-model="updated_at_range"
            unlink-panels
            class="filter-item"
            style="width:280px;margin-left:10px;margin-right:10px"
            type="datetimerange"
            start-placeholder="更新起始日期"
            end-placeholder="更新截止日期"
          />
          <el-button
            v-waves
            type="primary"
            class="filter-item"
            icon="el-icon-search"
            @click="getDocsList"
          >
            搜索
          </el-button>
          <el-button
            v-waves
            type="primary"
            class="filter-item"
            icon="el-icon-edit"
            @click="showCreate = true"
          >
            添加文档
          </el-button>
          <!--          <el-upload-->
          <!--            class="filter-item"-->
          <!--            action=""-->
          <!--            :http-request="httpRequest"-->
          <!--            :on-preview="handlePreview"-->
          <!--            :on-remove="handleRemove"-->
          <!--            :on-success="handleSuccess"-->
          <!--            :before-remove="beforeRemove"-->
          <!--            :limit="1"-->
          <!--            :file-list="fileList"-->
          <!--            :on-change="handleChange"-->
          <!--            :show-file-list="false"-->
          <!--            style="display: inline-block"-->
          <!--          >-->
          <!--            <el-tooltip-->
          <!--              class="item"-->
          <!--              effect="dark"-->
          <!--              content="只能上传图片、文档文件"-->
          <!--              placement="top"-->
          <!--            >-->
          <!--              <el-button-->
          <!--                type="primary"-->
          <!--                style="margin-top:10px;margin-left: 10px"-->
          <!--                class="filter-item"-->
          <!--              >-->
          <!--                <i class="el-icon-upload el-icon&#45;&#45;right" />-->
          <!--                上传文件-->
          <!--              </el-button>-->
          <!--            </el-tooltip>-->
          <!--          </el-upload>-->
        </div>
      </div>
      <div class="main-table">
        <el-table
          :data="docsList"
          border
          fit
          highlight-current-row
          style="width: 100%"
          @sort-change="sortChange"
        >
          <el-table-column
            label="ID"
            sortable="custom"
            align="center"
            width="100px"
          >
            <template #default="{row}">
              <span>{{ row.id }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="文档名称"
            sortable="custom"
            align="center"
            min-width="200px"
          >
            <template #default="{row}">
              <el-tooltip
                popper-class="msg-tooltip"
                :visible-arrow="false"
                effect="dark"
                :content="row.name"
                placement="top-start"
              >
                <span class="text-ellipsis">{{ row.name }}</span>
              </el-tooltip>
            </template>
          </el-table-column>
          <el-table-column
            label="文件大小(MB)"
            prop="file_size"
            sortable="custom"
            align="center"
            width="150px"
          >
            <template #default="{row}">
              <span>{{ row.file_size }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="文档类型"
            prop="document_type"
            sortable="custom"
            align="center"
            width="150px"
          >
            <template #default="{row}">
              <span>{{ row.document_type.replace(".",'') }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="创建者"
            prop="docs_creator"
            sortable="custom"
            width="150px"
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
            width="260px"
            class-name="fixed-width"
          >
            <template #default="{row}">
              <el-button
                size="mini"
                type="primary"
                :loading="downloading"
                @click="downloadDocs(row,row.id)"
              >
                下载
              </el-button>
              <el-button
                size="mini"
                type="danger"
                @click="deleteDocs(row.id)"
              >
                删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div
        class="pagination-footer"
        style="margin-top: 1%"
      >
        <el-pagination
          v-show="total > 0"
          v-model:currentPage="listDocsQuery.page"
          v-model:page-size="listDocsQuery.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getDocsList"
          @current-change="getDocsList"
        />
      </div>
      <docs-create
        v-model="showCreate"
        :category="category"
        @get-docs-list="getDocsList"
      />
      <!--      <docs-view-->
      <!--        v-model="showDetail"-->
      <!--        :docs="docsDetail"-->
      <!--      />-->
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, ref, onMounted, Prop, PropType } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { DeleteDocuments, GetDocuments, ListDocumentsQuery, Documents, DownloadDocuments, UploadDocuments } from '@/api/document'
import UserSearchInput from '@/components/user-search-input/index.vue'
import DocsCreate from './docsCreate.vue'
import dayjs from 'dayjs'

import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import { Response } from '@/model'
import { AxiosResponse } from 'axios'
export default defineComponent({
  components: {
    UserSearchInput,
    DocsCreate
  },
  props: {
    category: {
      type: String as PropType<'设计文档'|'培训文档'|'系统文档'>,
      default: '设计文档',
      required: true
    }
  },
  setup(props) {
    // eslint-disable-next-line @typescript-eslint/no-var-requires
    const created_at_range = ref(Array<Date>())
    const updated_at_range = ref(Array<Date>())
    const listDocsQuery = reactive({
      field_lk_name: '',
      field_eq_document_type: '',
      field_eq_category: props.category,
      field_gt_created_at: 0,
      field_lt_created_at: 0,
      field_gt_updated_at: 0,
      field_lt_updated_at: 0,
      page: 1,
      limit: 10,
      order: 'desc',
      order_by: 'created_at'
    })
    const loading = ref(false)
    const downloading = ref(false)
    const docsList = ref(Array<Documents>())
    const total = ref(0)
    const fileList = ref([])
    function handleSuccess() {
      fileList.value = []
      getDocsList()
    }
    // param是自带参数。 this.$refs.upload.submit() 会自动调用 httpRequest方法.在里面取得file
    async function httpRequest(param:any) {
      const data = new FormData()
      data.append('file', param.file)
      const response = await UploadDocuments(data)
      if (response.data.code === 200) {
        ElMessage.success('文件上传成功!')
      }
      getDocsList()
    }
    function UnixTime2Human(timestamp: string) {
      // UnixTime2Human(timestamp)

      return dayjs(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }
    async function getDocsList() {
      loading.value = true
      if (created_at_range.value && created_at_range.value.length === 2) {
        listDocsQuery.field_gt_created_at = created_at_range.value[0].getTime() / 1000
        listDocsQuery.field_lt_created_at = created_at_range.value[1].getTime() / 1000
      } else {
        listDocsQuery.field_gt_created_at = 0
        listDocsQuery.field_lt_created_at = 0
      }
      if (updated_at_range.value && updated_at_range.value.length === 2) {
        listDocsQuery.field_gt_updated_at = updated_at_range.value[0].getTime() / 1000
        listDocsQuery.field_lt_updated_at = updated_at_range.value[1].getTime() / 1000
      } else {
        listDocsQuery.field_gt_updated_at = 0
        listDocsQuery.field_lt_updated_at = 0
      }
      const response = await GetDocuments(listDocsQuery)
      // console.log(response)
      if (response.data.code === 200) {
        total.value = response.data.total
        docsList.value = response.data.total ? response.data.spec : []
      }
      loading.value = false
    }
    async function downloadDocs(docs: Documents, id:number) {
      downloading.value = true
      const res = await DownloadDocuments(id)
      if ((res.data as Blob).type === 'application/json') {
        const data : Response<any> = await new Promise((resolve, reject) => {
          const reader = new FileReader()
          reader.readAsText(res.data, 'utf-8')
          reader.onload = function() {
            try {
              resolve(JSON.parse(reader.result as string))
            } catch (error) {
              resolve({ code: 500, error: '文件解析失败', spec: undefined, total: 0 })
            }
          }
        })
        if (data.code !== 200) {
          downloading.value = false
          ElMessage.error(data.error)
        }
        return
      }
      try {
        if (window.navigator && window.navigator.msSaveBlob) {
          window.navigator.msSaveBlob(res.data, docs.name + docs.document_type)
        } else {
          const objectUrl = URL.createObjectURL(new Blob([res.data]))
          const link = document.createElement('a')
          link.style.display = 'none'
          link.href = objectUrl
          link.download = docs.name + docs.document_type
          document.body.appendChild(link)
          link.click()
          document.body.removeChild(link)
          URL.revokeObjectURL(objectUrl)
        }
      } catch (err) {
        console.log(err)
      }
      downloading.value = false
    }
    // console.log(res)
    function sortChange(column: any) {
      listDocsQuery.order_by = column.prop
      listDocsQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getDocsList()
    }
    onMounted(() => {
      getDocsList()
    })
    function deleteDocs(id: number) {
      ElMessageBox.confirm('是否删除此记录，请注意此操作不可恢复！', '确认删除记录', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteDocuments(id)
            if (response.data.code === 200) {
              ElMessage.success('文件删除成功!')
              await getDocsList()
            }
          }
        }
      })
    }
    const showCreate = ref(false)
    const docs: Documents = {
      id: 0,
      name: '',
      uuid: '',
      location: '',
      file_size: null,
      document_type: '',
      category: ''
    }
    const docsDetail = ref(docs)
    return {
      downloading,
      created_at_range,
      updated_at_range,
      handleSuccess,
      fileList,
      httpRequest,
      downloadDocs,
      status,
      loading,
      docsList,
      // DocsCreate,
      listDocsQuery,
      total,
      getDocsList,
      sortChange,
      deleteDocs,
      showCreate,
      docsDetail,
      UnixTime2Human,
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
  .form-icon {
    font-size: 18px;
    &.warning {
      color: $yellow;
    }
  }
  .custom-table-popper {
    max-width: 500px;
  }
  </style>
