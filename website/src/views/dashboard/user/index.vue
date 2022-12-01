<!--
 * @Description: admin 权限主页
 * @Author: ZY
 * @Date: 2021-01-15 18:44:25
 * @LastEditors: SCY
 * @LastEditTime: 2021-04-02 09:40:07
-->
<template>
  <div
    class="dashboard-user-container"
  >
    <el-row
      :gutter="8"
      style="height: 667px;"
    >
      <el-col
        :xs="{span: 24}"
        :sm="{span: 24}"
        :md="{span: 24}"
        :lg="{span: 16}"
        :xl="{span: 16}"
        style="padding-right:8px;margin-bottom:30px;"
      >
        <UpdateTimeline style="border-radius: 16px;" />
      </el-col>
      <el-col
        :xs="{span: 24}"
        :sm="{span: 24}"
        :md="{span: 24}"
        :lg="{span: 8}"
        :xl="{span: 8}"
        style="padding-right:8px;margin-bottom:30px;"
      >
        <TodoList style="border-radius: 8px;" />
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue'
import TodoList from '../admin/components/todo-list/index.vue'
import UpdateTimeline from '../admin/components/UpdateTimeline.vue'
import { ResourceList, ActionList } from '@/model/permission'

export default defineComponent({
  components: {
    TodoList,
    UpdateTimeline
  },
  setup() {
    onMounted(async () => {
    })

    return {
      ResourceList,
      ActionList
    }
  }
})
</script>

<style lang="scss" scoped>
.dashboard-user-container {
  padding: 32px;
  background-color: #F2F7FF;
  position: relative;
  height: calc(100vh-116px);

}

</style>
