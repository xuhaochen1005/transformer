<template>
  <el-dialog
    v-model="show"
    title="修改相间距离最小值"
    show-close
    width="30%"
  >
    <el-form
      ref="lpdUpdateForm"
      :model="tmpLpd"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="相间距离最小值（mm）"
        prop="phase_dist_min"
      >
        <el-input-number
          v-model="tmpLpd.phase_dist_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="电压下限（kV）"
        prop="vol_min"
      >
        <el-input-number
          v-model="tmpLpd.vol_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="电压上限（kV）"
        prop="vol_max"
      >
        <el-input-number
          v-model="tmpLpd.vol_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLpd"
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
import { computed, defineComponent, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Lpd, UpdateStdLpdLibraries } from '@/api/standard_libraries/lpd'

export default defineComponent({
  name: 'LpdUpdate',
  props: {
    modelValue: Boolean,
    lpd: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lpd-list'],
  setup(props, context) {
    const lpdUpdateForm = ref(null)
    const lpd: Lpd = {
      id: 0,
      phase_dist_min: 0,
      vol_min: 0,
      vol_max: 0
    }
    const tmpLpd = ref(lpd)
    watch(props, () => {
      tmpLpd.value.id = props.lpd.id
      tmpLpd.value.phase_dist_min = props.lpd.phase_dist_min
      tmpLpd.value.vol_min = props.lpd.vol_min
      tmpLpd.value.vol_max = props.lpd.vol_max
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lpdUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
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
    function updateLpd() {
      (lpdUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLpdLibraries(tmpLpd.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lpd-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLpd() {
    //   const response = await UpdateStdLpdLibraries(tmpLpd.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lpd-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lpdUpdateForm,
      show,
      tmpLpd,
      rules,
      updateLpd
    }
  }
})
</script>
