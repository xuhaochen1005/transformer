<template>
  <el-dialog
    v-model="show"
    title="修改工频耐压标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lifhvUpdateForm"
      :model="tmpLifhv"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定高压下界(kV)"
        prop="rated_high_vol_min"
        label-width=""
      >
        <el-input-number
          v-model="tmpLifhv.rated_high_vol_min"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定高压上界(kV)"
        prop="rated_high_vol_max"
        label-width=""
      >
        <el-input-number
          v-model="tmpLifhv.rated_high_vol_max"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="工频耐压(kV)"
        prop="industry_freq_hold_vol"
        label-width=""
      >
        <el-input-number
          v-model="tmpLifhv.industry_freq_hold_vol"
          style="width:100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLifhv"
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
import { Lifhv, UpdateStdlifhvLibraries } from '@/api/standard_libraries/lifhv'

export default defineComponent({
  name: 'LifhvUpdate',
  props: {
    modelValue: Boolean,
    lifhv: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lifhv-list'],
  setup (props, context) {
    const lifhvUpdateForm = ref(null)
    const lifhv: Lifhv = {
      id: 0,
      industry_freq_hold_vol: null,
      rated_high_vol_min: null,
      rated_high_vol_max: null
    }
    const tmpLifhv = ref(lifhv)
    watch(props, () => {
      tmpLifhv.value.id = props.lifhv.id!
      tmpLifhv.value.industry_freq_hold_vol = props.lifhv.industry_freq_hold_vol
      tmpLifhv.value.rated_high_vol_min = props.lifhv.rated_high_vol_min
      tmpLifhv.value.rated_high_vol_max = props.lifhv.rated_high_vol_max
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lifhvUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lifhv.rated_high_vol_min == null) {
        return callback(new Error('请输入额定高压下界(kV)'))
      }
      if (lifhv.rated_high_vol_min <= 0) {
        return callback(new Error('额定高压下界必须大于0！'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lifhv.rated_high_vol_max == null) {
        return callback(new Error('请输入额定高压上界(kV)'))
      }
      if (lifhv.rated_high_vol_max <= 0) {
        return callback(new Error('额定高压上界必须大于0！'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lifhv.industry_freq_hold_vol == null) {
        return callback(new Error('请输入工频耐压(kV)：'))
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
      industry_freq_hold_vol: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ]
    }
    function updateLifhv() {
      (lifhvUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlifhvLibraries(tmpLifhv.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lifhv-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLifhv() {
    //   const response = await UpdateStdlifhvLibraries(tmpLifhv.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lifhv-list')
    //     show.value = false
    //   }
    // }
    return {
      lifhvUpdateForm,
      show,
      tmpLifhv,
      rules,
      updateLifhv
    }
  }
})
</script>
