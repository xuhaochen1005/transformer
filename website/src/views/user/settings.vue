<template>
  <div class="app-container">
    <aside class="page-header">
      用户设置
    </aside>
    <el-tabs
      type="border-card"
      style="min-height: 750px;"
    >
      <el-tab-pane label="重置密码">
        <el-row>
          <el-col :span="10">
            <el-form
              ref="resetForm"
              :model="tempReset"
              :rules="resetRules"
              label-width="100px"
              label-suffix=":"
            >
              <el-form-item
                label="&emsp;原密码"
                prop="title"
                label-width=""
              >
                <el-input
                  v-model="tempReset.password"
                  type="password"
                  placeholder="请输入原密码"
                />
              </el-form-item>
              <el-form-item
                label="&emsp;新密码"
                prop="content"
                label-width=""
              >
                <el-input
                  v-model="tempReset.new_password"
                  type="password"
                  placeholder="请输入新密码"
                />
              </el-form-item>
              <el-form-item
                label="确认密码"
                prop="send_to"
                label-width=""
              >
                <el-input
                  v-model="tempReset.repeat_password"
                  type="password"
                  placeholder="请再输入一遍新密码"
                />
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>
        <el-row style="width: 700px;flex-direction: row-reverse">
          <div>
            <el-button
              type="primary"
              @click="confirmReset()"
            >
              确认重置
            </el-button>
          </div>
        </el-row>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script lang='ts'>
import {
  computed,
  defineComponent, onMounted, reactive,
  ref
} from 'vue'

import type { User } from '@/model/common'
export default defineComponent({
  props: {
    modelValue: Boolean
  },
  setup(props, ctx) {
    const resetForm = ref(null)
    const tempReset = reactive({
      password: '',
      new_password: '',
      repeat_password: ''
    })
    const resetRules = {

    }
    async function confirmReset() {

    }
    return {
      resetForm,
      tempReset,
      resetRules,
      confirmReset
    }
  }
})
</script>

<style lang="scss" scoped>
.form-icon {
  font-size: 18px;
  &.warning {
    color: $yellow;
  }
}
</style>
