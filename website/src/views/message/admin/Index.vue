<template>
  <div class="app-container">
    <aside class="page-header">
      消息管理
    </aside>
    <div class="main-table">
      <el-tabs
        v-model="activeTab"
      >
        <div
          v-if="activeTab === 'messageList'"
          class="filter-container"
        >
          <div class="container-row">
            <el-input
              v-model="query.field_lk_title"
              placeholder="标题"
              style="width: 200px"
              clearable
              class="filter-item"
              @keyup.enter="handleFilter"
            />
            <el-input
              v-model="query.field_lk_content"
              placeholder="消息内容"
              style="width: 200px"
              clearable
              class="filter-item"
              @keyup.enter="handleFilter"
            />
            <el-select
              v-model="query.field_eq_status"
              placeholder="消息状态"
              clearable
              style="width: 120px"
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
              icon="el-icon-download"
            >
              导出
            </el-button>
          </div>
        </div>
        <div
          v-if="activeTab === 'messageTemplate'"
          class="filter-container"
        >
          <div class="container-row">
            <el-input
              v-model="msgTemplateQuery.field_lk_name"
              placeholder="模板名称"
              style="width: 200px"
              clearable
              class="filter-item"
              @keyup.enter="handleMsgTemplateFilter"
            />
            <el-input
              v-model="msgTemplateQuery.field_lk_title"
              placeholder="模板标题"
              style="width: 200px"
              clearable
              class="filter-item"
              @keyup.enter="handleMsgTemplateFilter"
            />
            <el-input
              v-model="msgTemplateQuery.field_lk_content"
              placeholder="内容"
              style="width: 200px"
              clearable
              class="filter-item"
              @keyup.enter="handleMsgTemplateFilter"
            />
            <el-button
              v-waves
              class="filter-item"
              type="primary"
              icon="el-icon-search"
              @click="handleMsgTemplateFilter"
            >
              搜索
            </el-button>
            <el-button
              class="filter-item"
              style="margin-left: 10px"
              type="primary"
              icon="el-icon-edit"
              @click="handleMsgTemplateCreate()"
            >
              添加
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
        </div>
        <el-tab-pane
          label="消息列表"
          name="messageList"
        >
          <el-table
            key="messageList"
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
              width="75"
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
            >
              <template #default="{row}">
                <span>{{ row.title }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="发送方"
              prop="send_by_user"
              sortable="custom"
              align="center"
              width="150"
            >
              <template #default="{row}">
                <span>{{ row.send_by_user.user_name_zh }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="接收方"
              prop="send_to_user"
              sortable="custom"
              align="center"
              width="150"
            >
              <template #default="{row}">
                <span>{{ row.send_to_user.user_name_zh }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="消息状态"
              prop="status"
              sortable="custom"
              align="center"
              width="150"
            >
              <template #default="{row}">
                <template
                  v-for="option in messageStatusOptions"
                  :key="option"
                >
                  <span v-if="option.value === row.status">
                    {{ option.label }}
                  </span>
                </template>
              </template>
            </el-table-column>
            <el-table-column
              label="创建时间"
              prop="created_at"
              sortable="custom"
              align="center"
              width="250"
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
              width="250"
            >
              <template #default="{row}">
                <span>{{ UnixTime2HumanWithStr(row.updated_at) }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="操作"
              width="300"
              align="center"
              class-name="fixed-width"
            >
              <template #default="{row}">
                <el-button
                  type="primary"
                  size="mini"
                  @click="handleMsgDetail(row)"
                >
                  查看详情
                </el-button>
                <el-button
                  size="mini"
                  type="success"
                  @click="handleMsgUpdate(row)"
                >
                  修改
                </el-button>
                <el-button
                  v-if="row.status !== 'deleted'"
                  size="mini"
                  type="danger"
                  @click="deleteMessage(row.id)"
                >
                  删除
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
        <el-tab-pane
          label="消息模板"
          name="messageTemplate"
        >
          <el-table
            key="msgTemplateList"
            v-loading="msgTemplateListLoading"
            :data="msgTemplateList"
            border
            fit
            highlight-current-row
            style="width: 100%"
            @sort-change="msgTemplateSortChange"
          >
            <el-table-column
              label="ID"
              prop="id"
              sortable="custom"
              align="center"
              width="75"
            >
              <template #default="{row}">
                <span>{{ row.id }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="模板名称"
              prop="name"
              sortable="custom"
              align="center"
              width="150"
            >
              <template #default="{row}">
                <span>{{ row.name }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="模板标题"
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
              label="模板内容"
              prop="content"
              sortable="custom"
              align="center"
            >
              <template #default="{row}">
                <el-tooltip
                  class="item"
                  popper-class="custom-table-popper"
                  :visible-arrow="false"
                  :content="row.content"
                >
                  <span class="text-ellipsis">{{ row.content }}</span>
                </el-tooltip>
              </template>
            </el-table-column>
            <el-table-column
              label="创建者"
              prop="create_by_user"
              sortable="custom"
              align="center"
              width="100"
            >
              <template #default="{row}">
                <span>{{ row.create_by_user ? row.create_by_user.name : '' }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="创建时间"
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
              label="更新时间"
              prop="id"
              sortable="custom"
              align="center"
              width="200"
            >
              <template #default="{row}">
                <span>{{ UnixTime2HumanWithStr(row.updated_at) }}</span>
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
                  type="primary"
                  size="mini"
                  @click="handleMsgTemplateDetail(row)"
                >
                  查看详情
                </el-button>
                <el-button
                  v-if="row.status !== 'published'"
                  size="mini"
                  type="success"
                  @click="handleMsgTemplateUpdate(row)"
                >
                  修改
                </el-button>
                <el-button
                  v-if="row.status !== 'deleted'"
                  size="mini"
                  type="danger"
                  @click="deleteMsgTemplate(row.id)"
                >
                  删除
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
      <div style="margin-top: 20px;">
        <el-pagination
          v-if="activeTab === 'messageList' && total > 0"
          v-model:currentPage="query.page"
          v-model:page-size="query.limit"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getMessageList()"
          @current-change="getMessageList()"
        />
        <el-pagination
          v-if="activeTab === 'messageTemplate' && msgTemplateTotal > 0"
          v-model:currentPage="msgTemplateQuery.page"
          v-model:page-size="msgTemplateQuery.limit"
          :total="msgTemplateTotal"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="getMsgTemplateList()"
          @current-change="getMsgTemplateList()"
        />
      </div>
    </div>
    <div
      class="pagination-footer"
      style="margin-top: 1%"
    />
    <!--消息详情-->
    <el-dialog
      v-model="showMsgDetail"
      width="800px"
      title="消息详情"
      show-close
    >
      <el-form
        :model="tempMsgDetail"
        label-width="100px"
        label-suffix=":"
      >
        <el-form-item
          label="ID"
          prop="id"
          label-width=""
        >
          {{ tempMsgDetail.id }}
        </el-form-item>
        <el-form-item
          label="消息标题"
          prop="title"
          label-width=""
        >
          {{ tempMsgDetail.title }}
        </el-form-item>
        <el-form-item
          label="消息内容"
          prop="content"
          label-width=""
        >
          <div
            class="message-content"
            v-html="tempMsgDetail.content"
          />
        </el-form-item>
        <el-form-item
          label="发送者"
          prop="send_by"
          label-width=""
        >
          {{ tempMsgDetail.send_by_user?.user_name_zh }}
        </el-form-item>
        <el-form-item
          label="接收者"
          prop="send_to"
          label-width=""
        >
          {{ tempMsgDetail.send_to_user?.user_name_zh }}
        </el-form-item>
        <el-form-item
          label="消息状态"
          prop="status"
          label-width=""
        >
          <template
            v-for="option in messageStatusOptions"
            :key="option"
          >
            <span v-if="option.value === tempMsgDetail.status">
              {{ option.label }}
            </span>
          </template>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showMsgDetail = false">
          关闭
        </el-button>
      </template>
    </el-dialog>
    <!--更新消息-->
    <el-dialog
      v-model="showMsgUpdate"
      width="800px"
      title="修改消息"
      show-close
    >
      <el-form
        ref="msgUpdateForm"
        :model="tempMsgDetail"
        :rules="msgFormRules"
        label-width="100px"
        label-suffix=":"
      >
        <el-form-item
          label="消息标题"
          prop="title"
          label-width=""
        >
          <el-input
            v-model="tempMsgDetail.title"
            clearable
          />
        </el-form-item>
        <el-form-item
          label="消息内容"
          prop="content"
          label-width=""
        >
          <el-input
            v-model="tempMsgDetail.content"
            type="textarea"
            :rows="5"
            clearable
          />
        </el-form-item>
        <el-form-item
          label="发送人"
          prop="send_by"
          label-width=""
        >
          <UserSearchInput
            v-model="tempMsgDetail.send_by"
            style="width: 100%;"
            :users="tempMsgDetail.send_by_user"
          />
        </el-form-item>
        <el-form-item
          label="接收人"
          prop="send_to"
          label-width=""
        >
          <UserSearchInput
            v-model="tempMsgDetail.send_to"
            style="width: 100%;"
            :users="tempMsgDetail.send_to_user"
          />
        </el-form-item>
        <el-form-item
          label="消息状态"
          prop="status"
          label-width=""
        >
          <el-select
            v-model="tempMsgDetail.status"
            placeholder="消息状态"
            clearable
            style="width: 100%;"
          >
            <el-option
              v-for="item in messageStatusOptions"
              :key="item"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button
          type="primary"
          @click="confirmMsgUpdate()"
        >
          确定
        </el-button>
        <el-button @click="showMsgUpdate = false">
          取消
        </el-button>
      </template>
    </el-dialog>
    <!--    &lt;!&ndash;发送消息&ndash;&gt;-->
    <!--    <el-dialog-->
    <!--      v-model="showMsgCreate"-->
    <!--      width="600px"-->
    <!--      title="发送消息"-->
    <!--      show-close-->
    <!--    >-->
    <!--      <el-form-->
    <!--        ref="msgCreateForm"-->
    <!--        :model="tempMsgDetail"-->
    <!--        :rules="msgFormRules"-->
    <!--        label-width="100px"-->
    <!--        label-suffix=":"-->
    <!--      >-->
    <!--        <el-form-item-->
    <!--          label="标题"-->
    <!--          prop="title"-->
    <!--          label-width=""-->
    <!--        >-->
    <!--          <el-input-->
    <!--            v-model="tempMsgDetail.title"-->
    <!--            clearable-->
    <!--          />-->
    <!--        </el-form-item>-->
    <!--        <el-form-item-->
    <!--          label="内容"-->
    <!--          prop="content"-->
    <!--          label-width=""-->
    <!--        >-->
    <!--          <el-input-->
    <!--            v-model="tempMsgDetail.content"-->
    <!--            type="textarea"-->
    <!--            :row="5"-->
    <!--            clearable-->
    <!--          />-->
    <!--        </el-form-item>-->
    <!--        <el-form-item-->
    <!--          label="接收人"-->
    <!--          prop="send_to"-->
    <!--          label-width=""-->
    <!--        >-->
    <!--          <UserSearchInput-->
    <!--            v-model="tempMsgDetail.send_to"-->
    <!--            multiple-->
    <!--            :users="tempMsgDetail.send_to_user"-->
    <!--          />-->
    <!--          <el-tooltip-->
    <!--            content="接收人为空则为所有用户"-->
    <!--            placement="top"-->
    <!--            style="margin-left: 10px;"-->
    <!--          >-->
    <!--            <i-->
    <!--              class="form-icon warning el-icon-warning-outline"-->
    <!--            />-->
    <!--          </el-tooltip>-->
    <!--        </el-form-item>-->
    <!--      </el-form>-->
    <!--      <template #footer>-->
    <!--        <el-button-->
    <!--          type="primary"-->
    <!--          @click="confirmMsgCreate()"-->
    <!--        >-->
    <!--          确定-->
    <!--        </el-button>-->
    <!--        <el-button @click="showMsgCreate = false">-->
    <!--          取消-->
    <!--        </el-button>-->
    <!--      </template>-->
    <!--    </el-dialog>-->

    <!--模板详情-->
    <el-dialog
      v-model="showMsgTemplateDetail"
      width="600px"
      title="模板详情"
      show-close
    >
      <el-form
        v-model="tempMsgTemplateDetail"
        label-width="100px"
        label-suffix=":"
      >
        <el-form-item
          label="模板编号"
          prop="id"
          label-width=""
        >
          {{ tempMsgTemplateDetail.id }}
        </el-form-item>
        <el-form-item
          label="模板名称"
          prop="name"
          label-width=""
        >
          {{ tempMsgTemplateDetail.name }}
        </el-form-item>
        <el-form-item
          label="模板标题"
          prop="title"
          label-width=""
        >
          {{ tempMsgTemplateDetail.title }}
        </el-form-item>
        <el-form-item
          label="模板内容"
          prop="content"
          label-width=""
        >
          <div v-html="tempMsgTemplateDetail.content" />
        </el-form-item>
        <el-form-item
          label="创建者"
          prop="create_by"
          label-width=""
        >
          {{ tempMsgTemplateDetail.create_by_user?.name }}
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showMsgTemplateDetail = false">
          关闭
        </el-button>
      </template>
    </el-dialog>
    <!--更新模板-->
    <el-dialog
      v-model="showMsgTemplateUpdate"
      width="600px"
      title="修改模板"
      show-close
    >
      <el-form
        ref="msgTemplateUpdateForm"
        :model="tempMsgTemplateDetail"
        :rules="msgTemplateFormRules"
        label-width="100px"
        label-suffix=":"
      >
        <el-form-item
          label="模板名称"
          prop="name"
          label-width=""
        >
          <el-input
            v-model="tempMsgTemplateDetail.name"
            clearable
          />
        </el-form-item>
        <el-form-item
          label="模板标题"
          prop="title"
          label-width=""
        >
          <el-input
            v-model="tempMsgTemplateDetail.title"
            clearable
          />
        </el-form-item>
        <el-form-item
          type="textarea"
          label="模板内容"
          prop="content"
          label-width=""
        >
          <el-input
            v-model="tempMsgTemplateDetail.content"
            type="textarea"
            rows="5"
            clearable
          />
        </el-form-item>
        <el-form-item
          label="创建者"
          prop="create_by"
          label-width=""
        >
          <UserSearchInput
            v-model="tempMsgTemplateDetail.create_by"
            :users="tempMsgTemplateDetail.create_by_user"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button
          type="primary"
          @click="confirmMsgTemplateUpdate()"
        >
          确定
        </el-button>
        <el-button @click="showMsgTemplateUpdate = false">
          取消
        </el-button>
      </template>
    </el-dialog>
    <!--创建模板-->
    <el-dialog
      v-model="showMsgTemplateCreate"
      width="600px"
      title="创建消息模板"
      show-close
    >
      <el-form
        ref="msgTemplateCreateForm"
        :model="tempMsgTemplateDetail"
        :rules="msgTemplateFormRules"
        label-width="100px"
        label-suffix=":"
      >
        <el-form-item
          label="模板名称"
          prop="name"
          label-width=""
        >
          <el-input
            v-model="tempMsgTemplateDetail.name"
            clearable
          />
        </el-form-item>
        <el-form-item
          label="模板标题"
          prop="title"
          label-width=""
        >
          <el-input
            v-model="tempMsgTemplateDetail.title"
            clearable
          />
        </el-form-item>
        <el-form-item
          label="模板内容"
          prop="content"
          label-width=""
        >
          <el-input
            v-model="tempMsgTemplateDetail.content"
            type="textarea"
            :rows="5"
            clearable
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button
          type="primary"
          @click="confirmMsgTemplateCreate()"
        >
          确定
        </el-button>
        <el-button @click="showMsgTemplateCreate = false">
          取消
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>
<script lang="ts">
import {
  defineComponent,
  onMounted,
  reactive,
  watch,
  ref
} from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  GetUserMessages,
  GetMessageTemplates,
  UpdateUserMessage,
  MessageNotify,
  CreateMessageTemplate, UpdateMessageTemplate, DeleteMessage, DeleteMessageTemplate
} from '@/api/message'
import type { MessageQuery, MessageTemplateQuery, Message, MessageTemplate } from '@/model/message'
import { MessageStatus } from '@/model/message'
import UserSearchInput from '@/components/user-search-input/index.vue'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'

export default defineComponent({
  components: {
    UserSearchInput
  },
  setup(props, context) {
    const query: MessageQuery = reactive({
      field_lk_title: '',
      field_lk_content: '',
      field_eq_status: null,
      send_by: null,
      send_to: null,
      status: null,
      self: 2,
      page: 1,
      limit: 10,
      order: 'desc',
      order_by: 'created_at'
    })
    // 列表
    const firstTab = 'messageList'
    const activeTab = ref(firstTab)
    const listLoading = ref(false)
    const messageList = ref(Array<Message>())
    const total = ref(0)
    async function getMessageList() {
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
    // 详情
    const currentMsgTemplate = ref(null)
    const showMsgDetail = ref(false)
    const emptyMsg: Message = {
      id: 0,
      title: '',
      content: '',
      send_by: 0,
      send_by_user: undefined,
      send_to: 0,
      send_to_user: undefined,
      status: 0
    }
    const tempMsgDetail = ref(emptyMsg)
    function handleMsgDetail(message: Message) {
      tempMsgDetail.value = message
      showMsgDetail.value = true
    }
    // 消息通用验证表单
    function validateMsgTitle(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('标题不可为空'))
      }
      callback()
    }
    function validateMsgContent(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('消息内容不可为空'))
      }
      callback()
    }
    function validateMsgStatus(rule: any, value: any, callback: any) {
      if (!(Object.values(MessageStatus).includes(value as MessageStatus))) {
        return callback(new Error('消息状态只能为已读或未读'))
      }
      callback()
    }
    const msgFormRules = {
      title: [
        { required: true, validator: validateMsgTitle, trigger: 'blur' }
      ],
      content: [
        { required: true, validator: validateMsgContent, trigger: 'blur' }
      ],
      status: [
        { required: true, validator: validateMsgStatus, trigger: 'blur' }
      ]
    }
    // 更新
    const messageUpdate = {
      id: 0,
      title: '',
      content: '',
      send_by: 0,
      send_to: 0,
      status: 0
    }
    const showMsgUpdate = ref(false)
    function handleMsgUpdate(message: Message) {
      tempMsgDetail.value = JSON.parse(JSON.stringify(message))
      showMsgUpdate.value = true
    }
    const msgUpdateForm = ref(null)
    watch(showMsgUpdate, function(newValue) {
      if (!newValue) {
        (msgUpdateForm.value as any).clearValidate()
      }
    })
    function confirmMsgUpdate() {
      (msgUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const tempMsg = tempMsgDetail.value
          messageUpdate.id = tempMsg.id
          messageUpdate.title = tempMsg.title
          messageUpdate.content = tempMsg.content
          messageUpdate.send_to = tempMsg.send_to
          messageUpdate.send_by = tempMsg.send_by
          messageUpdate.status = tempMsg.status
          const response = await UpdateUserMessage(messageUpdate)
          if (response.data.code === 200) {
            ElMessage.success('消息更新成功')
            await getMessageList()
            showMsgUpdate.value = false
          }
        }
      })
    }
    // 创建
    const messageCreate: Message = {
      id: 0,
      title: '',
      content: '',
      send_by: 0,
      send_to: 0,
      status: 0
    }
    const showMsgCreate = ref(false)
    const msgCreateForm = ref(null)
    function handleMsgCreate() {
      // 重置临时对象
      tempMsgDetail.value = emptyMsg
      tempMsgDetail.value.status = 1
      showMsgCreate.value = true
    }
    function confirmMsgCreate() {
      (msgCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const tm = tempMsgDetail.value
          messageCreate.title = tm.title
          messageCreate.content = tm.content
          messageCreate.send_to = tm.send_to
          messageCreate.status = tm.status
          const response = await MessageNotify(messageCreate)
          if (response.data.code === 200) {
            ElMessage.success('消息发送成功')
            showMsgCreate.value = false
            await getMessageList()
          }
        }
      })
    }
    watch(showMsgCreate, (newValue) => {
      if (!newValue) {
        (msgCreateForm.value as any).clearValidate()
        tempMsgDetail.value = emptyMsg
      }
    })
    // 删除
    function deleteMessage(id : number) {
      ElMessageBox.confirm('是否删除此条消息', '确认删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteMessage(id)
            if (response.data.code === 200) {
              ElMessage.success('消息删除成功!')
              await getMessageList()
            }
          }
        }
      })
    }

    // 消息模板
    // 列表
    const msgTemplateListLoading = ref(false)
    const msgTemplateQuery : MessageTemplateQuery = reactive({
      field_lk_name: '',
      field_lk_title: '',
      field_lk_content: '',
      page: 1,
      limit: 10,
      order: 'desc',
      order_by: 'id'
    })
    const msgTemplateList = ref(Array<MessageTemplate>())
    const msgTemplateTotal = ref(0)
    async function getMsgTemplateList() {
      const response = await GetMessageTemplates(msgTemplateQuery)
      if (response.data.code === 200) {
        msgTemplateTotal.value = response.data.total
        msgTemplateList.value = response.data.total ? response.data.spec : []
      }
    }
    function handleMsgTemplateFilter() {
      getMsgTemplateList()
    }
    // 模板详情
    const showMsgTemplateDetail = ref(false)
    const emptyMsgTemplate: MessageTemplate = {
      id: 0,
      name: '',
      title: '',
      content: '',
      create_by: 0,
      created_at: '',
      updated_at: ''
    }
    const tempMsgTemplateDetail = ref(emptyMsgTemplate)
    function handleMsgTemplateDetail(msgTemplate : MessageTemplate) {
      tempMsgTemplateDetail.value = msgTemplate
      showMsgTemplateDetail.value = true
    }
    function msgTemplateSortChange(column: any) {
      msgTemplateQuery.order_by = column.prop
      msgTemplateQuery.order = column.order === 'descending' ? 'desc' : 'asc'
      getMsgTemplateList()
    }
    // 模板通用验证
    const msgTempValidator = {
      name: (rule: any, value: any, callback: any) => {
        if (!value) {
          return callback(new Error('模板名称不可为空'))
        }
        callback()
      },
      title: (rule: any, value: any, callback: any) => {
        if (!value) {
          return callback(new Error('模板标题不可为空'))
        }
        callback()
      },
      content: (rule: any, value: any, callback: any) => {
        if (!value) {
          return callback(new Error('模板内容不可为空'))
        }
        callback()
      },
      create_by: (rule: any, value: any, callback: any) => {
        if (!value) {
          return callback(new Error('创建者不可为空'))
        }
        callback()
      }
    }
    const msgTemplateFormRules = {
      name: [{ required: true, validator: msgTempValidator.name, trigger: 'blur' }],
      title: [{ required: true, validator: msgTempValidator.title, trigger: 'blur' }],
      content: [{ required: true, validator: msgTempValidator.content, trigger: 'blur' }],
      create_by: [{ required: true, validator: msgTempValidator.create_by, trigger: 'blur' }]
    }
    // 更新模板
    const showMsgTemplateUpdate = ref(false)
    function handleMsgTemplateUpdate(msgTemplate : MessageTemplate) {
      tempMsgTemplateDetail.value = JSON.parse(JSON.stringify(msgTemplate))
      showMsgTemplateUpdate.value = true
    }
    const msgTemplateUpdateForm = ref(null)
    function confirmMsgTemplateUpdate() {
      (msgTemplateUpdateForm.value as any).validate(async (valid: boolean) => {
        if (!valid) {
          return
        }
        const tempMsgTemplateUpdate: MessageTemplate = {
          id: 0,
          name: '',
          title: '',
          content: '',
          create_by: 0
        }
        const tmt = tempMsgTemplateDetail.value
        tempMsgTemplateUpdate.id = tmt.id
        tempMsgTemplateUpdate.name = tmt.name
        tempMsgTemplateUpdate.title = tmt.title
        tempMsgTemplateUpdate.content = tmt.content
        tempMsgTemplateUpdate.create_by = tmt.create_by
        const response = await UpdateMessageTemplate(tempMsgTemplateUpdate)
        if (response.data.code === 200) {
          ElMessage.success('消息模板更新成功')
          await getMsgTemplateList()
          showMsgTemplateUpdate.value = false
        }
      })
    }
    // 创建模板
    const showMsgTemplateCreate = ref(false)
    function handleMsgTemplateCreate() {
      tempMsgTemplateDetail.value = emptyMsgTemplate
      showMsgTemplateCreate.value = true
    }
    const msgTemplateCreateForm = ref(null)
    function confirmMsgTemplateCreate() {
      console.log('msgTemplateCreateForm');
      (msgTemplateCreateForm.value as any).validate(async (valid: boolean) => {
        console.log('msgTemplateCreateForm valid', valid)
        if (!valid) {
          return
        }
        const tempMsgTemplateCreate: MessageTemplate = {
          id: 0,
          name: '',
          title: '',
          content: '',
          create_by: 0
        }
        const tmt = tempMsgTemplateDetail.value
        tempMsgTemplateCreate.name = tmt.name
        tempMsgTemplateCreate.title = tmt.title
        tempMsgTemplateCreate.content = tmt.content
        const response = await CreateMessageTemplate(tempMsgTemplateCreate)
        if (response.data.code === 200) {
          ElMessage.success('消息模板创建成功')
          await getMsgTemplateList()
          showMsgTemplateCreate.value = false
        }
      })
    }
    // 删除模板
    function deleteMsgTemplate(id : number) {
      ElMessageBox.confirm('是否删除此条消息模板', '确认删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteMessageTemplate(id)
            if (response.data.code === 200) {
              ElMessage.success('消息模板删除成功!')
              await getMsgTemplateList()
            }
          }
        }
      })
    }

    onMounted(() => {
      getMessageList()
      getMsgTemplateList()
    })
    return {
      // util method
      UnixTime2HumanWithStr,
      // 消息列表
      query,
      activeTab,
      listLoading,
      messageStatusOptions,
      messageList,
      total,
      tempMsgDetail,
      currentMsgTemplate,
      showMsgDetail,
      showMsgUpdate,
      msgUpdateForm,
      msgFormRules,
      showMsgCreate,
      msgCreateForm,
      getMessageList,
      messageSortChange,
      handleFilter,
      handleMsgDetail,
      handleMsgUpdate,
      confirmMsgUpdate,
      handleMsgCreate,
      confirmMsgCreate,
      deleteMessage,
      // 消息模板列表
      msgTemplateListLoading,
      msgTemplateQuery,
      msgTemplateTotal,
      msgTemplateList,
      showMsgTemplateDetail,
      msgTemplateFormRules,
      msgTemplateUpdateForm,
      msgTemplateCreateForm,
      tempMsgTemplateDetail,
      showMsgTemplateUpdate,
      showMsgTemplateCreate,
      msgTemplateSortChange,
      getMsgTemplateList,
      handleMsgTemplateFilter,
      handleMsgTemplateDetail,
      handleMsgTemplateUpdate,
      handleMsgTemplateCreate,
      confirmMsgTemplateUpdate,
      confirmMsgTemplateCreate,
      deleteMsgTemplate
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
:deep() {
  .message-content {
    font-size: 16px;
    a {
      border-bottom: 1px solid $--color-primary;
      color: $--color-primary;
      font-size: 18px;
    }
  }
}

</style>
