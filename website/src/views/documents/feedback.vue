<template>
  <div class="app-container">
    <aside class="page-message-header">
      反馈中心
    </aside>
    <div class="main-table">
      <div
        class="filter-container"
      >
        <div class="container-row">
          <el-input
            v-model="query.field_lk_title"
            class="filter-item"
            clearable
            placeholder="搜索标题"
            style="width: 200px"
            @keyup.enter="handleFilter"
          />
          <el-input
            v-model="query.field_lk_content"
            class="filter-item"
            clearable
            placeholder="搜索内容"
            style="width: 200px"
            @keyup.enter="handleFilter"
          />
          <el-button
            v-waves
            class="filter-item"
            icon="el-icon-search"
            type="primary"
            @click="handleFilter"
          >
            搜索
          </el-button>
          <el-button
            v-waves
            class="filter-item"
            icon="el-icon-edit"
            type="primary"
            @click="showCreate = true"
          >
            发送新反馈
          </el-button>
          <!--          <el-button-->
          <!--            v-waves-->
          <!--            class="filter-item"-->
          <!--            type="primary"-->
          <!--            icon="el-icon-download"-->
          <!--          >-->
          <!--            导出-->
          <!--          </el-button>-->
        </div>
      </div>
      <div class="main-table">
        <el-table
          key="message-table"
          v-loading="listLoading"
          :data="feedbackList"
          border
          fit
          highlight-current-row
          style="width: 100%"
          @sort-change="sortChange"
        >
          <el-table-column
            align="center"
            label="ID"
            prop="id"
            sortable="custom"
            width="90"
          >
            <template #default="{row}">
              <span>{{ row.id }}</span>
            </template>
          </el-table-column>
          <el-table-column
            align="center"
            label="标题"
            prop="title"
            sortable="custom"
            width="250"
          >
            <template #default="{row}">
              <span>{{ row.title }}</span>
            </template>
          </el-table-column>
          <el-table-column
            align="center"
            label="反馈内容"
            min-width="300"
            prop="content"
            sortable="custom"
          >
            <template #default="{row}">
              <el-tooltip
                :content="row.content"
                :visible-arrow="false"
                effect="dark"
                placement="top-start"
                popper-class="msg-tooltip"
              >
                <template
                  #content
                >
                  <span v-html="row.content" />
                </template>
                <span
                  class="text-ellipsis"
                  v-html="row.content"
                />
              </el-tooltip>
            </template>
          </el-table-column>
          <el-table-column
            align="center"
            label="创建者"
            sortable="custom"
            width="150"
          >
            <template #default="{row}">
              <span>{{ row.creator_user?.user_name_zh }}</span>
            </template>
          </el-table-column>
          <el-table-column
            align="center"
            label="创建时间"
            prop="created_at"
            sortable="custom"
            width="200"
          >
            <template #default="{row}">
              <span>{{ UnixTime2HumanWithStr(row.created_at) }}</span>
            </template>
          </el-table-column>
          <el-table-column
            align="center"
            class-name="fixed-width"
            label="操作"
            width="300"
          >
            <template #default="{row}">
              <el-button
                size="mini"
                type="primary"
                @click="handleDetail(row)"
              >
                查看
              </el-button>
              <el-button
                size="mini"
                type="danger"
                @click="handleDelete(row)"
              >
                删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
        <div
          class="pagination-footer"
          style="margin-top: 1%"
        >
          <el-pagination

            v-model:current-page="query.page"
            v-model:page-size="query.limit"
            :page-sizes="[10, 20, 50, 100]"
            :total="total"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="getFeedbackList"
            @current-change="getFeedbackList"
          />
        </div>
      </div>
    </div>
    <div class="dialog-mains">
      <feedback-create
        v-if="showCreate"
        v-model:model-value="showCreate"
        @update-list="getFeedbackList()"
      />
      <feedback-view
        v-model:model-value="showMsgDetail"
        :handle-value="currentRow"
      />
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, onMounted, reactive, ref, Ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import type { Message } from '@/model/message'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import { useRoute, useRouter } from 'vue-router'
import { DeleteFeedback, FeedbackListQuery, GetFeedbackList } from '@/api/feedback'
import { Feedback } from '@/model/feedback'
import feedbackCreate from './feedback/Create.vue'
import feedbackView from './feedback/View.vue'

export default defineComponent({
  components: {
    feedbackCreate,
    feedbackView
  },
  setup(props, context) {
    const router = useRouter()
    const route = useRoute()
    const query: FeedbackListQuery = reactive({
      page: 1,
      limit: 10,
      order: 'desc',
      order_by: 'created_at'
    })
    // 列表
    const listLoading = ref(false)
    const feedbackList: Ref<Feedback[]> = ref([])
    const total = ref(0)

    async function getFeedbackList() {
      listLoading.value = true
      const response = await GetFeedbackList(query)
      if (response.data.code === 200) {
        total.value = response.data.total
        feedbackList.value = response.data.total ? response.data.spec : []
      }
      listLoading.value = false
    }

    function sortChange(column: any) {
      query.order_by = column.prop
      query.order = column.order === 'descending' ? 'desc' : 'asc'
      getFeedbackList()
    }

    // filter
    function handleFilter() {
      getFeedbackList()
    }

    const showCreate = ref(false)
    const emptyRow: Feedback = {
      content: '', created_at: '', creator: 0, id: 0, title: '', updated_at: ''
    } as Feedback
    const currentRow: Ref<Feedback> = ref(JSON.parse(JSON.stringify(emptyRow)) as Feedback)
    const showMsgDetail = ref(false)

    async function handleDetail(row: Feedback) {
      currentRow.value = row
      showMsgDetail.value = true
      getFeedbackList()
    }

    async function handleDelete(row: Message) {
      ElMessageBox.confirm('是否删除此反馈，请注意此操作不可恢复！', '确认删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const res = await DeleteFeedback(row.id)
            if (res.data.code === 200) {
              ElMessage.success('反馈删除成功!')
              getFeedbackList()
            }
          }
        }
      })
    }

    onMounted(async () => {
      await getFeedbackList()
    })
    return {
      // util method
      UnixTime2HumanWithStr,
      // 反馈列表
      query,
      listLoading,
      feedbackList,
      total,
      getFeedbackList,
      sortChange,
      handleFilter,
      showCreate,
      handleDetail,
      showMsgDetail,
      currentRow,
      handleDelete
    }
  }
})
</script>

<style lang="scss" scoped>
.page-message-header {
  line-height: 24px;
  font-size: 24px;
  color: #409eff;
  background: none;
  padding: 0;
}

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

.message-card {
  box-shadow: 0px 0px 10px 2px rgba(64, 158, 255, 0.25);
  padding: 10px 10px;
  min-height: 180px;
  margin-bottom: 15px;

  .message-header {
    height: 65px;
  }

  .message-avatar {
    float: left;
    height: 50px;
    width: 50px;
    margin-right: 10px;
    border-radius: 50%;
    overflow: hidden;

    img {
      width: 100%;
      height: 100%;
    }
  }

  .message-name {
    font-size: 18px;
    font-weight: bold;
    color: #409eff;
  }

  .message-desc {
    margin-top: 10px;
    font-size: 14px;
    color: #000;
  }

  .message-publish {
    margin-left: 10px;

    .publish-time {
      color: #409eff;
      margin-left: 10px;
    }
  }

  .message-title {
    margin-top: 10px;
    color: #f56c6c;
    font-size: 18px;
    padding: 0 10px;
  }

  .message-content {
    color: #000;
    font-size: 16px;
    padding: 0 10px;
    margin-top: 20px;
  }

}

</style>
<style lang="scss">
.msg-tooltip {
  a {
    border-bottom: 1px solid $--color-primary;
    color: $--color-primary;
  }

  max-width: 500px !important;
}
</style>
