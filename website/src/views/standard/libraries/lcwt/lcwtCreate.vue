<template>
  <el-dialog
    v-model="show"
    title="新增线圈导线类型标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lcwtCreateForm"
      :rules="rules"
      :model="lcwt"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="线圈导线类型"
        prop="coil_wire_type"
        label-width=""
      >
        <el-input
          v-model="lcwt.coil_wire_type"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="线圈导线类型字母代号"
        prop="coil_wire_type_sign"
        label-width=""
      >
        <el-input
          v-model="lcwt.coil_wire_type_sign"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLcwt"
      >
        确定
      </el-button>
      <el-button @click="show = false">
        取消
      </el-button>
    </template>
  </el-dialog>
</template>
<script lang="ts">
import { defineComponent, computed, reactive, ref } from 'vue'
import { CreateStdlcwtLibraries } from '@/api/standard_libraries/lcwt'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LcwtCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lcwt-list'],
  setup(props, context) {
    const lcwtCreateForm = ref(null)
    const lcwt = reactive({
      id: 0,
      coil_wire_type: '',
      coil_wire_type_sign: ''
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lcwtCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!lcwt.coil_wire_type) {
        return (new Error('请输入线圈导线类型：'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (!lcwt.coil_wire_type_sign) {
        return (new Error('请输入线圈导线类型代号：'))
      }
      callback()
    }

    const rules = {
      coil_wire_type: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      coil_wire_type_sign: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
    }
    function createLcwt() {
      (lcwtCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlcwtLibraries(lcwt)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lcwt-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLcwt() {
    //   const response = await CreateStdlcwtLibraries(lcwt)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lcwt-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lcwtCreateForm,
      show,
      lcwt,
      rules,
      createLcwt
    }
  }
})
</script>
