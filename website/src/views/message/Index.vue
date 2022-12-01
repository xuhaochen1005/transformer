<template>
  <div class="app-container">
    <aside class="page-message-header">
      消息中心
    </aside>
    <div class="main-table">
      <div
        class="filter-container"
      >
        <div class="container-row">
          <UserSearchInput
            v-model="query.field_eq_send_by"
            placeholder="发送者"
            class="filter-item"
          />
          <el-input
            v-model="query.field_lk_title"
            placeholder="搜索消息"
            style="width: 200px"
            clearable
            class="filter-item"
            @keyup.enter="handleFilter"
          />
          <el-select
            v-model="query.field_eq_status"
            placeholder="消息状态"
            clearable
            style="width: 200px"
            class="filter-item"
          >
            <el-option
              v-for="item in messageStatusOptions"
              :key="item"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
          <el-button
            v-waves
            class="filter-item"
            type="primary"
            icon="el-icon-search"
            @click="handleFilter"
          >
            搜索
          </el-button>
          <!--            <el-button-->
          <!--              class="filter-item"-->
          <!--              style="margin-left: 10px"-->
          <!--              type="primary"-->
          <!--              icon="el-icon-edit"-->
          <!--              @click="handleMsgCreate()"-->
          <!--            >-->
          <!--              添加-->
          <!--            </el-button>-->
          <el-button
            v-waves
            class="filter-item"
            type="primary"
            icon="el-icon-edit"
            @click="showMsgCreate = true"
          >
            发送新消息
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
          :data="messageList"
          border
          fit
          highlight-current-row
          style="width: 100%"
          @sort-change="messageSortChange"
        >
          <el-table-column
            label="ID"
            prop="id"
            sortable="custom"
            align="center"
            width="90"
          >
            <template #default="{row}">
              <span>{{ row.id }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="标题"
            prop="title"
            sortable="custom"
            align="center"
            width="250"
          >
            <template #default="{row}">
              <span>{{ row.title }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="消息内容"
            prop="content"
            sortable="custom"
            align="center"
            min-width="300"
          >
            <template #default="{row}">
              <el-tooltip
                popper-class="msg-tooltip"
                :visible-arrow="false"
                effect="dark"
                :content="row.content"
                placement="top-start"
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
            label="状态"
            prop="status"
            sortable="custom"
            align="center"
            width="150"
          >
            <template #default="{row}">
              <span>{{ messageStatusOptions.find(o => o.value === row.status)?.label }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="发送者"
            prop="send_by"
            sortable="custom"
            align="center"
            width="150"
          >
            <template #default="{row}">
              <span>{{ row.send_by_user?.user_name_zh }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="发送时间"
            prop="created_at"
            sortable="custom"
            align="center"
            width="200"
          >
            <template #default="{row}">
              <span>{{ UnixTime2HumanWithStr(row.created_at) }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="操作"
            align="center"
            width="300"
            class-name="fixed-width"
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
                v-if="row.status === MessageStatus.HASREAD"
                size="mini"
                type="warning"
                @click="handleSetMsgUnRead(row)"
              >
                标为未读
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
            :total="total"
            :page-sizes="[10, 20, 50, 100]"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="getMessageList"
            @current-change="getMessageList"
          />
        </div>
      </div>
    </div>
    <div class="dialog-mains">
      <MessageCreateForm v-model="showMsgCreate" />
      <MessageDetailView
        v-model="showMsgDetail"
        :handle-value="currentRow"
        @handleSetMsgUnRead="handleSetMsgUnRead"
      />
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, onMounted, reactive, ref, Ref } from 'vue'
import { DeleteMessage, GetUserMessages, UpdateUserMessage } from '@/api/message'
import { ElMessage, ElMessageBox } from 'element-plus'
import type { Message, MessageQuery } from '@/model/message'
import { MessageStatus } from '@/model/message'
import UserSearchInput from '@/components/user-search-input/index.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import MessageCreateForm from './Create.vue'
import MessageDetailView from './View.vue'
import { onBeforeRouteUpdate, useRoute, useRouter } from 'vue-router'
import emitter from '@/utils/emitter'

export default defineComponent({
  components: {
    UserSearchInput,
    MessageCreateForm,
    MessageDetailView
  },
  setup(props, context) {
    const router = useRouter()
    const route = useRoute()
    const query: MessageQuery = reactive({
      field_eq_send_by: undefined,
      field_lk_title: '',
      field_lk_content: '',
      field_eq_status: null,
      send_by: null,
      send_to: null,
      status: null,
      self: 1,
      page: 1,
      limit: 10,
      order: 'asc',
      order_by: 'created_at desc,status'
    })
    // 列表
    const listLoading = ref(false)
    const messageList : Ref<Message[]> = ref([])
    const total = ref(0)
    async function getMessageList() {
      emitter.emit('messageUpdate', true)
      listLoading.value = true
      const response = await GetUserMessages(query)
      if (response.data.code === 200) {
        total.value = response.data.total
        messageList.value = response.data.total ? response.data.spec : []
      }
      listLoading.value = false
    }

    const messageStatusOptions = [
      { label: '未读', value: MessageStatus.UNREAD },
      { label: '已读', value: MessageStatus.HASREAD }
    ]
    function messageSortChange(column: any) {
      query.order_by = column.prop
      query.order = column.order === 'descending' ? 'desc' : 'asc'
      getMessageList()
    }
    // filter
    function handleFilter() {
      getMessageList()
    }
    const avatarUrl = 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'

    const showMsgCreate = ref(false)
    const emptyRow: Message = {
      content: '', id: 0, send_by: 0, send_to: 0, status: 0, title: ''
    }
    const currentRow: Ref<Message> = ref(JSON.parse(JSON.stringify(emptyRow)) as Message)
    const showMsgDetail = ref(false)
    async function handleDetail(row: Message) {
      await setMsgStatus(row, MessageStatus.HASREAD)
      currentRow.value = row
      showMsgDetail.value = true
      getMessageList()
    }
    async function setMsgStatus(row: Message, status : MessageStatus) {
      const updateMsg = JSON.parse(JSON.stringify(row)) as Message
      updateMsg.status = status
      const res = await UpdateUserMessage(updateMsg)
      if (res.data.code === 200) {
        row.status = status
        if (status === MessageStatus.UNREAD) {
          ElMessage.success('设置消息未读成功')
          getMessageList()
        }
      }
    }
    async function handleSetMsgUnRead(row : Message) {
      await setMsgStatus(row, MessageStatus.UNREAD)
    }
    async function handleDelete(row : Message) {
      ElMessageBox.confirm('是否删除此消息，请注意此操作不可恢复！', '确认删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const res = await DeleteMessage(row.id)
            if (res.data.code === 200) {
              ElMessage.success('消息删除成功!')
              getMessageList()
            }
          }
        }
      })
    }
    function handleOpenMsgDetail(id: number) {
      messageList.value.map((msg) => {
        if (msg.id === id) {
          handleDetail(msg)
          router.push('/message/index')
        }
      })
    }
    onBeforeRouteUpdate(({ hash }) => {
      if (hash) {
        const id = parseInt(hash.substring(1))
        handleOpenMsgDetail(id)
      }
    })
    onMounted(async () => {
      await getMessageList()
      if (route.hash) {
        const id = parseInt(route.hash.substring(1))
        handleOpenMsgDetail(id)
      }
    })
    return {
      MessageStatus,
      // util method
      UnixTime2HumanWithStr,
      // 消息列表
      avatarUrl,
      query,
      listLoading,
      messageStatusOptions,
      messageList,
      total,
      getMessageList,
      messageSortChange,
      handleFilter,
      showMsgCreate,
      handleDetail,
      showMsgDetail,
      currentRow,
      handleSetMsgUnRead,
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
