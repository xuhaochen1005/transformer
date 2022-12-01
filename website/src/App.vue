<template>
  <router-view />
</template>

<script lang="ts">
import { defineComponent } from 'vue'
export default defineComponent({
})
</script>

<style lang="scss">
.el-submenu__title:focus, .el-submenu__title:hover {
  outline: 0 !important;
  color: #fff !important;
  background: #435EBE !important;
  border-radius: 8px !important;
}
.el-menu-item:hover {
  outline: 0 !important;
  color: #fff !important;
  background: #435EBE !important;
  border-radius: 8px !important;
}
.el-menu.el-menu--collapse.v-enter-to {
  .el-submenu__icon-arrow.el-icon-arrow-right {
    display: none;
  }
}
</style>
