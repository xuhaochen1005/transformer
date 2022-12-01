<template>
  <el-popover
    popper-class="message-popper"
    placement="bottom-end"
    :width="350"
    trigger="hover"
  >
    <template #reference>
      <span class="btn-message">
        <span class="message-container">
          <el-badge
            type="danger"
            :value="total"
            :max="99"
            :hidden="total === 0"
            class="msg-badge"
          >
            <svg-icon
              name="bells"
              :size="30"
            />
          </el-badge>
        </span>
      </span>
    </template>
    <div class="message-popover-container">
      <div class="header">
        消息中心
      </div>
      <div class="content-box">
        <div
          v-if="total > 0"
          class="msg-total-desc"
        >
          <div>一共有 <span class="msg-value">{{ total }}</span> 条未读消息</div>
          <el-button
            type="text"
            @click="allRead()"
          >
            全部设为已读
          </el-button>
        </div>
        <router-link
          v-for="message in unReadMessageList"
          :key="message"
          :to="`/message/index/#${message.id}`"
        >
          <div
            class="message-box"
          >
            <div class="message-header">
              <el-avatar
                size="small"
                :src="avatarUrl"
              />
              <span class="message-username">{{ message.send_by_user.user_name_zh }}</span>
            </div>
            <div class="message-content text-ellipsis">
              {{ message.title }}
            </div>
          </div>
        </router-link>
        <el-empty
          v-if="unReadMessageList.length <= 0"
          :image-size="100"
          description="暂时没有新消息"
        />
      </div>
    </div>
  </el-popover>
</template>

<script lang="ts">
import { defineComponent, onMounted, Ref, ref } from 'vue'
import { GetUserMessages, UpdateAllMessageRead } from '@/api/message'
import type { MessageQuery, Message } from '@/model/message'
import { MessageStatus } from '@/model/message'
import emitter from '@/utils/emitter'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

export default defineComponent({
  props: {},
  emits: [],
  setup(props, context) {
    const router = useRouter()
    const route = useRoute()
    const avatarUrl = 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'
    const messageQuery: MessageQuery = {
      field_lk_content: '',
      field_lk_title: '',
      self: 1,
      field_eq_status: MessageStatus.UNREAD,
      page: 1,
      limit: 3,
      order: 'desc',
      order_by: 'created_at'
    }
    const unReadMessageList: Ref<Message[]> = ref([])
    const total = ref(0)
    async function getMessageList() {
      if (route.name === 'Login') {
        return
      }
      const res = await GetUserMessages(messageQuery)
      if (res.data.code === 200) {
        unReadMessageList.value = res.data.spec
        total.value = res.data.total
      }
    }
    const allRead = async () => {
      const res = await UpdateAllMessageRead()
      if (res.data.code === 200) {
        getMessageList()
        ElMessage.success('全部设为已读成功')
      }
    }
    // 轮询获取最新消息
    setInterval(() => {
      getMessageList()
    }, process.env.NODE_ENV === 'development' ? 60000 : 5000)

    onMounted(() => {
      getMessageList()
      emitter.on('messageUpdate', e => {
        getMessageList()
      })
    })
    return {
      avatarUrl,
      unReadMessageList,
      total,
      allRead
    }
  }
})
</script>

<style lang="scss" scoped>
  .btn-message {
    width: 46px;
    height: 46px;
    position: relative;
    display: inline-block;
    vertical-align: top;
    .message-container {
      border-radius: 5px;
      margin-top: 5px;
      width: 40px;
      height: 40px;
      display: inline-block;
      line-height: 40px;
      text-align: center;
      position: relative;
      svg {
        margin-top: 5px;
      }
    }
  }
  .message-popover-container {
    height: 100%;
    border-radius: 10px;
    overflow: hidden;
    .header {
      width: 100%;
      height: 33px;
      background: $--color-primary;
      text-align: center;
      line-height: 33px;
      font-size: 18px;
      color: #fff;
    }
    .content-box {
      padding: 10px;
    }
    .message-box {
      cursor: pointer;
      padding: 5px;
      width: 100%;
      height: 73px;
      overflow: hidden;
      background: #FFFFFF;
      margin-bottom: 10px;
      box-shadow: 0px 0px 6px rgba(64, 158, 255, 0.3);
      border-radius: 5px;
      .message-header {
        height: 37px;
        line-height: 36px;
      }
      .message-username {
        vertical-align: top;
        margin-left: 5px;
        font-size: 16px;
        font-weight: bold;
        color: $--color-primary;
        line-height: 27px;
      }
      .message-content {
        height: 37px;
        width: 100%;
        overflow: hidden;
        font-size: 14px;
        color: #000;
        padding: 0 5px;
      }
    }
  }
  :deep() {
    .msg-badge {
      sup {
        top: 10px;
        right: 12px;
      }
    }
  }

</style>
<style lang="scss">
.el-popover.el-popper.message-popper {
  padding: 0;
  max-height: 350px;
  overflow: hidden;
  min-height: 80px;
}
.msg-value {
  color: $__color-danger;
  font-size: 16px;
  font-weight: 700;
}
.msg-total-desc {
  height: 55px;
  text-align: center;
  font-size: 14px;
}
.content-box {
  overflow: auto;
}

</style>
