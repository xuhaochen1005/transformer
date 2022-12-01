<template>
  <el-dialog
    v-model="show"
    class="custom-dialog"
    width="800px"
  >
    <template #title>
      <div>
        <span class="el-dialog__title">创建部门</span>
      </div>
    </template>
    <el-form
      ref="createForm"
      :model="tempDetail"
      :rules="formRules"
      label-width="100px"
      label-suffix=":"
    >
      <el-form-item
        label="部门名称"
        prop="title"
        label-width=""
      >
        <el-input
          v-model="tempDetail.name"
          prop="name"
          placeholder="部门名称"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="confirmCreate()"
      >
        确认
      </el-button>
      <el-button
        @click="show = false"
      >
        取消
      </el-button>
    </template>
  </el-dialog>
</template>

<script lang='ts'>
import { computed, defineComponent, ref, Ref } from 'vue'

import { ElMessage } from 'element-plus'
import { Department, DepartmentStatus } from '@/model/department'
import { CreateDepartment } from '@/api/department'

export default defineComponent({
  props: {
    modelValue: Boolean
  },
  setup(props, ctx) {
    const show = computed({
      get: () => props.modelValue,
      set: (v) => ctx.emit('update:modelValue', v)
    })
    const emptyForm: Department = {
      name: '',
      status: DepartmentStatus.Normal
    }

    const statusOptions = [
      { label: '正常', value: DepartmentStatus.Normal },
      { label: '停用', value: DepartmentStatus.Disabled }
    ]

    const tempDetail: Ref<Department> = ref(JSON.parse(JSON.stringify(emptyForm)) as Department)
    // 消息通用验证表单
    function validateName(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('部门名称不可为空'))
      }
      callback()
    }
    function validateStatus(rule: any, value: any, callback: any) {
      if (!(value in DepartmentStatus)) {
        return callback(new Error('状态只能为已读或未读'))
      }
      callback()
    }
    const formRules = {
      name: [
        { required: true, validator: validateName, trigger: 'blur' }
      ]
    }
    const createForm = ref(null)
    function confirmCreate() {
      (createForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateDepartment(tempDetail.value)
          if (response.data.code === 200) {
            ElMessage.success('部门创建成功')
            ctx.emit('refresh', true)
            show.value = false
          }
        }
      })
    }
    return {
      show,
      tempDetail,
      formRules,
      createForm,
      statusOptions,
      confirmCreate
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
