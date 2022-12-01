<template>
  <el-dialog
    v-model="show"
    title="修改冲击电压"
    show-close
    width="30%"
  >
    <el-form
      ref="lshvUpdateForm"
      :model="tmpLshv"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定高压下界（kV）"
        prop="rated_high_vol_min"
      >
        <el-input-number
          v-model="tmpLshv.rated_high_vol_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定高压上界（kV）"
        prop="rated_high_vol_max"
      >
        <el-input-number
          v-model="tmpLshv.rated_high_vol_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="冲击电压（kV）"
        prop="shock_hold_vol"
      >
        <el-input-number
          v-model="tmpLshv.shock_hold_vol"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLshv"
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
import { Lshv, UpdateStdLshvLibraries } from '@/api/standard_libraries/lshv'

export default defineComponent({
  name: 'LshvUpdate',
  props: {
    modelValue: Boolean,
    lshv: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lshv-list'],
  setup(props, context) {
    const lshvUpdateForm = ref(null)
    const lshv: Lshv = {
      id: 0,
      rated_high_vol_min: 0,
      rated_high_vol_max: 0,
      shock_hold_vol: 0
    }
    const tmpLshv = ref(lshv)
    watch(props, () => {
      tmpLshv.value.id = props.lshv.id
      tmpLshv.value.rated_high_vol_min = props.lshv.rated_high_vol_min
      tmpLshv.value.rated_high_vol_max = props.lshv.rated_high_vol_max
      tmpLshv.value.shock_hold_vol = props.lshv.shock_hold_vol
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lshvUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lshv.rated_high_vol_min == null) {
        return callback(new Error('请输入额定高压下界（kV）：'))
      }
      if (lshv.rated_high_vol_min < 0) {
        return callback(new Error('额定高压下界不能小于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lshv.rated_high_vol_max == null) {
        return callback(new Error('请输入额定高压上界（kV）：'))
      }
      if (lshv.rated_high_vol_max < 0) {
        return callback(new Error('额定高压上界不能小于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lshv.shock_hold_vol == null) {
        return callback(new Error('请输入冲击电压（kV）：'))
      }
      if (lshv.shock_hold_vol < 0) {
        return callback(new Error('冲击电压不能小于0'))
      }
      callback()
    }

    const rules = {
      rated_high_vol_min: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      rated_high_vol_max: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      shock_hold_vol: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ]

    }
    function updateLshv() {
      (lshvUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLshvLibraries(tmpLshv.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lshv-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLshv() {
    //   const response = await UpdateStdLshvLibraries(tmpLshv.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lshv-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lshvUpdateForm,
      show,
      tmpLshv,
      rules,
      updateLshv
    }
  }
})
</script>
