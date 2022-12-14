<!--
 * @Description: app 布局入口
 * @Author: ZY
 * @Date: 2020-12-17 15:32:33
 * @LastEditors: SCY
 * @LastEditTime: 2021-04-06 14:47:00
-->
<template>
  <el-config-provider :locale="locale">
    <div
      :class="classObj"
      class="app-wrapper"
    >
      <div
        v-if="classObj.mobile && sidebar.opened"
        class="drawer-bg"
        @click="handleClickOutside"
      />
      <Sidebar class="sidebar-container" />
      <div
        class="hasTagsView main-container"
      >
        <div class="fixed-header">
          <Navbar />
          <TagsView />
        </div>
        <AppMain />
      </div>
    </div>
  </el-config-provider>
</template>

<script lang="ts">
import { DeviceType } from '@/store/state'
import { computed, defineComponent, onBeforeMount, onBeforeUnmount, onMounted, reactive, ref, toRefs } from 'vue'
import { useStore } from '@/store'
import { AppMutation } from '@/store/app/mutation'
import { AppMain, Navbar, TagsView, Sidebar } from './components'
import resize from './resize'
import { ElConfigProvider } from 'element-plus'
import zhCn from 'element-plus/lib/locale/lang/zh-cn'
export default defineComponent({
  name: 'Layout',
  components: {
    AppMain,
    Navbar,
    Sidebar,
    TagsView,
    ElConfigProvider
  },
  setup() {
    const store = useStore()
    const { sidebar, device, addEventListenerOnResize, resizeMounted, removeEventListenerResize, watchRouter } = resize()
    const state = reactive({
      handleClickOutside: () => {
        store.commit(AppMutation.CLOSE_SIDEBAR, false)
      }
    })

    const classObj = computed(() => {
      return {
        hideSidebar: !sidebar.value.opened,
        openSidebar: sidebar.value.opened,
        withoutAnimation: sidebar.value.withoutAnimation,
        mobile: device.value === DeviceType.Mobile
      }
    })

    watchRouter()
    onBeforeMount(() => {
      addEventListenerOnResize()
    })

    onMounted(() => {
      resizeMounted()
    })

    onBeforeUnmount(() => {
      removeEventListenerResize()
    })
    return {
      locale: zhCn,
      classObj,
      sidebar,
      ...toRefs(state)
    }
  }
})
</script>

<style lang="scss" scoped>
.app-wrapper {
  @include clearfix;
  position: relative;
  height: 100%;
  width: 100%;
}

.drawer-bg {
  background: #000;
  opacity: 0.3;
  width: 100%;
  top: 0;
  height: 100%;
  position: absolute;
  z-index: 999;
}

.main-container {
  min-height: 100%;
  transition: margin-left .28s;
  margin-left: $sideBarWidth;
  position: relative;
}

.sidebar-container {
  transition: width 0.28s;
  width: $sideBarWidth !important;
  height: 100%;
  position: fixed;
  font-size: 0px;
  top: 0;
  bottom: 0;
  left: 0;
  z-index: 1001;
  overflow: hidden;
  background-color: #ffffff !important;
}

.fixed-header {
  position: fixed;
  top: 0;
  right: 0;
  z-index: 9;
  width: calc(100% - #{$sideBarWidth});
  transition: width 0.28s;
}

.hideSidebar {
  .main-container {
    margin-left: 54px;
  }

  .sidebar-container {
    width: 54px !important;
  }

  .fixed-header {
    width: calc(100% - 54px)
  }
}

/* for mobile response 适配移动端 */
.mobile {
  .main-container {
    margin-left: 0px;
  }

  .sidebar-container {
    transition: transform .28s;
    width: $sideBarWidth !important;
  }

  &.openSidebar {
    position: fixed;
    top: 0;
  }

  &.hideSidebar {
    .sidebar-container {
      pointer-events: none;
      transition-duration: 0.3s;
      transform: translate3d(-$sideBarWidth, 0, 0);
    }
  }

  .fixed-header {
    width: 100%;
  }
}

.withoutAnimation {
  .main-container,
  .sidebar-container {
    transition: none;
  }
}

</style>
