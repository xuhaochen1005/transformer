<!--
 * @Description:
 * @Author: ZY
 * @Date: 2020-12-24 10:35:47
 * @LastEditors: SCY
 * @LastEditTime: 2021-04-06 14:16:03
-->
<template>
  <div
    :class="{'has-logo': showLogo}"
    class="sideWrap"
  >
    <SidebarLogo
      v-if="showLogo"
      :collapse="isCollapse"
    />
    <el-scrollbar wrap-class="scrollbar-wrapper">
      <el-menu
        :collapse="!isCollapse"
        :unique-opened="false"
        :default-active="activeMenu"
        :background-color="variables.menuBg"
        :text-color="variables.menuText"
        :active-text-color="menuActiveTextColor"
        mode="vertical"
      >
        <SidebarItem
          v-for="route in routes"
          :key="route.path"
          :item="route"
          :base-path="route.path"
          :is-collapse="isCollapse"
        />
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue'
import SidebarItem from './SidebarItem.vue'
import SidebarLogo from './SidebarLogo.vue'
import variables from '@/styles/_variables.scss'
import { useStore } from '@/store'
import { useRoute } from 'vue-router'
import { permissionRoutes } from '@/router'

export default defineComponent({
  components: {
    SidebarItem,
    SidebarLogo
  },
  setup() {
    const store = useStore()
    const route = useRoute()
    const sidebar = computed(() => {
      return store.state.app.sidebar
    })
    const routes = computed(() => {
      return permissionRoutes
    })
    const showLogo = true

    const menuActiveTextColor = '#f1ff18'

    const activeMenu = computed(() => {
      const { meta, path } = route
      if (meta && meta.activeMenu) {
        return meta.activeMenu
      }
      return path
    })

    const isCollapse = computed(() => {
      return sidebar.value.opened
    })

    return {
      sidebar,
      routes,
      showLogo,
      menuActiveTextColor,
      variables,
      activeMenu,
      isCollapse
    }
  }
})
</script>

<style lang="scss" scoped>
  #app .sidebar-container.sideWrap {
    background-color: $__color-primary !important;
  }

  :deep() {
    #app .sidebar-logo-container .sidebar-logo-link .sidebar-title {
      color: #fff !important;
    }
    .el-submenu__title {
      border-radius: 0 !important;
    }
    .el-submenu > .el-submenu__title {
      font-weight: 700;
    }
    .el-submenu>.el-submenu__title .el-submenu__icon-arrow {
      font-size: 20px;
      color: #fff;
      font-weight: 700;
    }
  }
</style>
<style lang="scss">
.sidebar-container {
  // reset element-ui css
  .horizontal-collapse-transition {
    transition: 0s width ease-in-out, 0s padding-left ease-in-out,
      0s padding-right ease-in-out;
  }

  .scrollbar-wrapper {
    overflow-x: hidden !important;
  }

  .el-scrollbar__view {
    height: 100%;
  }

  .el-scrollbar__bar {
    &.is-vertical {
      right: 0px;
    }

    &.is-horizontal {
      display: none;
    }
  }
}
</style>

<style lang="scss" scoped>
.el-scrollbar {
  height: 100%;
}

.has-logo {
  .el-scrollbar {
    height: calc(100vh - 100px);
  }
}

.el-menu {
  border: none;
  height: 100%;
  width: 100% !important;
}
</style>
