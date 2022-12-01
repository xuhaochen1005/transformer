<template>
  <el-dialog
    v-model="show"
    title="修改噪声预测"
    show-close
    width="30%"
  >
    <el-form
      ref="lnpUpdateForm"
      :model="tmpLnp"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定容量下界（kVA）"
        prop="rated_capacity_min"
      >
        <el-input-number
          v-model="tmpLnp.rated_capacity_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定容量下界（kVA）"
        prop="rated_capacity_max"
      >
        <el-input-number
          v-model="tmpLnp.rated_capacity_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定高压下界（kV）"
        prop="rated_high_vol_min"
      >
        <el-input-number
          v-model="tmpLnp.rated_high_vol_min"
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
          v-model="tmpLnp.rated_high_vol_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="冷却方式"
        prop="cool_type"
      >
        <el-input
          v-model="tmpLnp.cool_type"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="噪声预测（dB）"
        prop="noise_predict"
      >
        <el-input-number
          v-model="tmpLnp.noise_predict"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLnp"
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
import { Lnp, UpdateStdLnpLibraries } from '@/api/standard_libraries/lnp'

export default defineComponent({
  name: 'LnpUpdate',
  props: {
    modelValue: Boolean,
    lnp: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lnp-list'],
  setup(props, context) {
    const lnpUpdateForm = ref(null)
    const lnp: Lnp = {
      id: 0,
      rated_capacity_min: 0,
      rated_capacity_max: 0,
      rated_high_vol_min: 0,
      rated_high_vol_max: 0,
      cool_type: '',
      noise_predict: 0
    }
    const tmpLnp = ref(lnp)
    watch(props, () => {
      tmpLnp.value.id = props.lnp.id
      tmpLnp.value.rated_capacity_min = props.lnp.rated_capacity_min
      tmpLnp.value.rated_capacity_max = props.lnp.rated_capacity_max
      tmpLnp.value.rated_high_vol_min = props.lnp.rated_high_vol_min
      tmpLnp.value.rated_high_vol_max = props.lnp.rated_high_vol_max
      tmpLnp.value.cool_type = props.lnp.cool_type
      tmpLnp.value.noise_predict = props.lnp.noise_predict
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lnpUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lnp.rated_capacity_min == null) {
        return callback(new Error('请输入额定容量下界（kVA）'))
      }
      if (lnp.rated_capacity_min < 0) {
        return callback(new Error('额定容量下界不能小于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lnp.rated_capacity_max == null) {
        return callback(new Error('请输入额定容量上界（kVA）'))
      }
      if (lnp.rated_capacity_max < 0) {
        return callback(new Error('额定容量上界不能小于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lnp.rated_high_vol_min == null) {
        return callback(new Error('请输入额定高压下界（kV）'))
      }
      if (lnp.rated_high_vol_min < 0) {
        return callback(new Error('额定高压下界不能小于0'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lnp.rated_high_vol_max == null) {
        return callback(new Error('请输入额定高压上界（kV）'))
      }
      if (lnp.rated_high_vol_max < 0) {
        return callback(new Error('额定高压上界不能小于0'))
      }
      callback()
    }

    function validateFactor_5(rule: any, value: any, callback: any) {
      if (!lnp.cool_type) {
        return callback(new Error('请输入冷却方式'))
      }
      callback()
    }

    function validateFactor_6(rule: any, value: any, callback: any) {
      if (lnp.noise_predict == null) {
        return callback(new Error('请输入噪声预测值（dB）'))
      }
      if (lnp.noise_predict < 0) {
        return callback(new Error('噪声预测值不能小于0'))
      }
      callback()
    }

    const rules = {
      rated_capacity_min: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      rated_capacity_max: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      rated_high_vol_min: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      rated_high_vol_max: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ],
      cool_type: [
        { required: true, validator: validateFactor_5, trigger: 'blur' }
      ],
      noise_predict: [
        { required: true, validator: validateFactor_6, trigger: 'blur' }
      ],

    }
    function updateLnp() {
      (lnpUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLnpLibraries(tmpLnp.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lnp-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLnp() {
    //   const response = await UpdateStdLnpLibraries(tmpLnp.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lnp-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lnpUpdateForm,
      show,
      tmpLnp,
      rules,
      updateLnp
    }
  }
})
</script>
