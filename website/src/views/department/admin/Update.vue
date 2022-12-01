<template>
  <el-dialog
    v-model="show"
    class="custom-dialog"
    width="800px"
  >
    <template #title>
      <div>
        <span class="el-dialog__title">更新部门</span>
      </div>
    </template>
    <el-form
      ref="updateForm"
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
      <el-form-item
        label="部门状态"
        prop="title"
        label-width=""
      >
        <el-select
          v-model="tempDetail.status"
          style="width: 100%;"
        >
          <el-option
            v-for="item in statusOptions"
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
        @click="confirmUpdate()"
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
import { computed, defineComponent, ref, Ref, watch, watchEffect } from 'vue'

import { ElMessage } from 'element-plus'
import { Department, DepartmentStatus } from '@/model/department'
import { updateDepartment } from '@/api/department'

export default defineComponent({
  props: {
    modelValue: Boolean,
    handleValue: {
      type: Object,
      required: true
    }
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

    const tempDetail: Ref<Department> = ref(JSON.parse(JSON.stringify(props.handleValue)) as Department)
    watch(() => props.handleValue, function (value) {
      tempDetail.value = JSON.parse(JSON.stringify(value)) as Department
      console.log(tempDetail)
    })
    console.log(tempDetail.value)
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
      ],
      status: [
        { required: true, validator: validateStatus, trigger: 'blur' }
      ]
    }
    const updateForm = ref(null)
    function confirmUpdate() {
      (updateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await updateDepartment(tempDetail!.value)
          if (response.data.code === 200) {
            ElMessage.success('部门更新成功')
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
      updateForm,
      statusOptions,
      confirmUpdate
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
