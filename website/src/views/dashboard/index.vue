<template>
  <div class="dashboard-container">
    <component :is="currentRole" />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import AdminDashboard from './admin/index.vue'
import { checkRole } from '@/utils/permission'
import { RoleList } from '@/model/permission'
export default defineComponent({
  components: {
    AdminDashboard
  },
  setup() {
    const currentRole = ref('admin-dashboard')
    // if (checkRole([RoleList.Root, RoleList.ChiefEngineer, RoleList.Demander])) {
    //   currentRole.value = 'admin-dashboard'
    // }
    // if (checkRole([RoleList.Designer])) {
    //   currentRole.value = 'editor-dashboard'
    // }
    // if (checkRole([RoleList.Reviewer])) {
    //   currentRole.value = 'reviewer-dashboard'
    // }
    return {
      currentRole
    }
  }
})
</script>
