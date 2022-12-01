<template>
  <el-dialog
    v-model="show"
    title="新增线圈接法标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lccCreateForm"
      :rules="rules"
      :model="lcc"
      label-width="150px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="低压线圈接法"
        prop="low_vol_coil_connect"
        label-width=""
      >
        <el-input
          v-model="lcc.low_vol_coil_connect"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="高压线圈接法"
        prop="high_vol_coil_connect"
        label-width=""
      >
        <el-input
          v-model="lcc.high_vol_coil_connect"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLcc"
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
import { CreateStdlccLibraries } from '@/api/standard_libraries/lcc'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LccCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lcc-list'],
  setup(props, context) {
    const lccCreateForm = ref(null)
    const lcc = reactive({
      id: 0,
      low_vol_coil_connect: '',
      high_vol_coil_connect: ''
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lccCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!lcc.low_vol_coil_connect) {
        return callback(new Error('请输入低压线圈连接方式'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (!lcc.high_vol_coil_connect) {
        return callback(new Error('请输入高压线圈连接方式'))
      }
      callback()
    }

    const rules = {
      low_vol_coil_connect: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      high_vol_coil_connect: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ]
    }
    function createLcc() {
      (lccCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlccLibraries(lcc)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lcc-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLcc() {
    //   const response = await CreateStdlccLibraries(lcc)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lcc-list')
    //     show.value = false
    //   }
    // }
    return {
      lccCreateForm,
      show,
      lcc,
      rules,
      createLcc
    }
  }
})
</script>
