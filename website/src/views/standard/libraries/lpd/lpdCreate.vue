<template>
  <el-dialog
    v-model="show"
    title="新增相间距离最小值"
    show-close
    width="30%"
  >
    <el-form
      ref="lpdCreateForm"
      :model="lpd"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="相间距离最小值（mm）"
        prop="phase_dist_min"
        label-width=""
      >
        <el-input-number
          v-model="lpd.phase_dist_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="电压下限（kV）"
        prop="vol_min"
        label-width=""
      >
        <el-input-number
          v-model="lpd.vol_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="电压上限（kV）"
        prop="vol_max"
        label-width=""
      >
        <el-input-number
          v-model="lpd.vol_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLpd"
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
import { CreateStdLpdLibraries } from '@/api/standard_libraries/lpd'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LpdCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lpd-list'],
  setup(props, context) {
    const lpdCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lpdCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lpd = reactive({
      phase_dist_min: 0,
      vol_min: 0,
      vol_max: 0
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lpd.phase_dist_min == null) {
        return callback(new Error('请输入相间距离最小值（mm）：'))
      }
      if (lpd.phase_dist_min < 0) {
        return callback(new Error('相间距离最小值不能小于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lpd.vol_min == null) {
        return callback(new Error('请输入电压下限（kV）：'))
      }
      if (lpd.vol_min < 0) {
        return callback(new Error('电压下限不能小于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lpd.vol_max == null) {
        return callback(new Error('请输入电压上限（kV）：'))
      }
      if (lpd.vol_max < 0) {
        return callback(new Error('电压上限不能小于0'))
      }
      callback()
    }

    const rules = {
      phase_dist_min: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      vol_min: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      vol_max: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
    }
    function createLpd() {
      (lpdCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLpdLibraries(lpd)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lpd-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLpd() {
    //   const response = await CreateStdLpdLibraries(lpd)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lpd-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lpdCreateForm,
      show,
      lpd,
      rules,
      createLpd
    }
  }
})
</script>
