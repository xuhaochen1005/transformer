<template>
  <el-dialog
    v-model="show"
    title="新增线圈端部距铁轭距离标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lpnCreateForm"
      :rules="rules"
      :model="lpn"
      label-width="220px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="相数"
        prop="phase_num"
        label-width=""
      >
        <el-input
          v-model="lpn.phase_num"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="相数代号"
        prop="phase_num_sign"
        label-width=""
      >
        <el-input
          v-model="lpn.phase_num_sign"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLpn"
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
import { CreateStdLpnLibraries } from '@/api/standard_libraries/lpn'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LpnCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lpn-list'],
  setup(props, context) {
    const lpnCreateForm = ref(null)
    const lpn = reactive({
      id: 0,
      phase: 0,
      phase_num: '',
      phase_num_sign: ''
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lpnCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!lpn.phase_num) {
        return callback(new Error('请输入相数'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (!lpn.phase_num_sign) {
        return callback(new Error('请输入相数代号'))
      }
      callback()
    }

    const rules = {
      phase_num: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      phase_num_sign: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ]
    }
    function createLpn() {
      (lpnCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLpnLibraries(lpn)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lpn-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLpn() {
    //   const response = await CreateStdLpnLibraries(lpn)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lpn-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lpnCreateForm,
      show,
      lpn,
      rules,
      createLpn
    }
  }
})
</script>
