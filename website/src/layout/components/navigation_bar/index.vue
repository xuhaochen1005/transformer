<!--
 * @Description: 导航栏
 * @Author: ZY
 * @Date: 2020-12-17 15:52:19
 * @LastEditors: ZY
 * @LastEditTime: 2021-01-27 19:16:50
-->
<template>
  <div class="navbar">
    <Hamburger
      id="hamburger-container"
      :is-active="sidebar.opened"
      class="hamburger-container"
      @toggle-click="toggleSideBar"
    />
    <BreadCrumb
      id="breadcrumb-container"
      class="breadcrumb-container"
    />
    <div class="right-menu">
      <template v-if="device !== 'mobile'">
        <!-- <error-log class="errLog-container right-menu-item hover-effect" /> -->
        <MessagePopover />
        <Screenfull class="right-menu-item hover-effect" />
      </template>
      <el-dropdown
        class="avatar-container right-menu-item hover-effect"
        trigger="click"
      >
        <div class="avatar-wrapper">
          <img
            :src="avatar + '?imageView2/1/w/80/h/80'"
            class="user-avatar"
          >
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <router-link to="/user/profile/">
              <el-dropdown-item>
                设置
              </el-dropdown-item>
            </router-link>
            <router-link to="/">
              <el-dropdown-item>
                首页
              </el-dropdown-item>
            </router-link>
            <el-dropdown-item
              divided
              @click="logout"
            >
              <span style="display:block;">
                退出
              </span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script lang="ts">
import BreadCrumb from '@/components/bread-crumb/index.vue'
import Hamburger from '@/components/hamburger/index.vue'
import Screenfull from '@/components/screenfull/index.vue'
import MessagePopover from '@/components/message-popover/index.vue'

import { computed, defineComponent, reactive, toRefs } from 'vue'
import { useStore } from '@/store'
import { AppMutation } from '@/store/app/mutation'
import { UserMutation } from '@/store/user/mutation'
import { useRoute, useRouter } from 'vue-router'

export default defineComponent({
  components: {
    BreadCrumb,
    Hamburger,
    Screenfull,
    MessagePopover
  },
  setup() {
    const store = useStore()
    const route = useRoute()
    const router = useRouter()
    const sidebar = computed(() => {
      return store.state.app.sidebar
    })
    const device = computed(() => {
      return store.state.app.device.toString()
    })
    const avatar = computed(() => {
      return 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif'
    })
    const state = reactive({
      toggleSideBar: () => {
        store.commit(AppMutation.TOGGLE_SIDEBAR, false)
      },
      logout: () => {
        store.commit(UserMutation.LOGOUT)
        router.push(`/login?redirect=${route.fullPath}`).catch(err => {
          console.warn(err)
        })
      }
    })
    return {
      sidebar,
      device,
      avatar,
      ...toRefs(state)
    }
  }
})
</script>

<style lang="scss" scoped>
.navbar {
  height: 50px;
  overflow: hidden;
  position: relative;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);

  .hamburger-container {
    line-height: 46px;
    height: 100%;
    float: left;
    padding: 0 15px;
    cursor: pointer;
    transition: background 0.3s;
    -webkit-tap-highlight-color: transparent;

    &:hover {
      background: rgba(0, 0, 0, 0.025);
    }
  }

  .breadcrumb-container {
    float: left;
  }

  .errLog-container {
    display: inline-block;
    vertical-align: top;
  }

  .right-menu {
    float: right;
    height: 100%;
    line-height: 50px;

    &:focus {
      outline: none;
    }

    .right-menu-item {
      display: inline-block;
      padding: 0 8px;
      height: 100%;
      font-size: 18px;
      color: #5a5e66;
      vertical-align: text-bottom;

      &.hover-effect {
        cursor: pointer;
        transition: background 0.3s;

        &:hover {
          background: rgba(0, 0, 0, 0.025);
        }
      }
    }

    .avatar-container {
      .avatar-wrapper {
        margin-top: 5px;
        margin-right: 16px;
        margin-left: 16px;
        position: relative;

        .user-avatar {
          cursor: pointer;
          width: 40px;
          height: 40px;
          border-radius: 10px;
        }

        .el-icon-caret-bottom {
          cursor: pointer;
          position: absolute;
          right: -20px;
          top: 25px;
          font-size: 12px;
        }
      }
    }
  }
}
</style>
