<template>
  <el-dialog
    v-model="show"
    title="修改线圈端部距铁轭距离标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lpnUpdateForm"
      :model="tmpLpn"
      :rules="rules"
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
          v-model="tmpLpn.phase_num"
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
          v-model="tmpLpn.phase_num_sign"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLpn"
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
import { computed, defineComponent, reactive, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Lpn, UpdateStdLpnLibraries } from '@/api/standard_libraries/lpn'

export default defineComponent({
  name: 'LpnUpdate',
  props: {
    modelValue: Boolean,
    lpn: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lpn-list'],
  setup (props, context) {
    const lpnUpdateForm = ref(null)
    const lpn: Lpn = {
      id: 0,
      phase: 0,
      phase_num: null,
      phase_num_sign: null
    }
    const tmpLpn = ref(lpn)
    watch(props, () => {
      tmpLpn.value.id = props.lpn.id!
      tmpLpn.value.phase_num = props.lpn.phase_num
      tmpLpn.value.phase_num_sign = props.lpn.phase_num_sign
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lpnUpdateForm.value as any).clearValidate()
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
    function updateLpn() {
      (lpnUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLpnLibraries(tmpLpn.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lpn-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLpn() {
    //   const response = await UpdateStdLpnLibraries(tmpLpn.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lpn-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lpnUpdateForm,
      show,
      tmpLpn,
      rules,
      updateLpn
    }
  }
})
</script>
