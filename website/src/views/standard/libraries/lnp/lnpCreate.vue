<template>
  <el-dialog
    v-model="show"
    title="新增噪声预测"
    show-close
    width="30%"
  >
    <el-form
      ref="lnpCreateForm"
      :model="lnp"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定容量下界（kVA）"
        prop="rated_capacity_min"
        label-width=""
      >
        <el-input-number
          v-model="lnp.rated_capacity_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定容量上界（kVA）"
        prop="rated_capacity_max"
        label-width=""
      >
        <el-input-number
          v-model="lnp.rated_capacity_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定高压下界（kV）"
        prop="rated_high_vol_min"
        label-width=""
      >
        <el-input-number
          v-model="lnp.rated_high_vol_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定高压上界（kV）"
        prop="rated_high_vol_max"
        label-width=""
      >
        <el-input-number
          v-model="lnp.rated_high_vol_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="冷却方式"
        prop="cool_type"
        label-width=""
      >
        <el-input
          v-model="lnp.cool_type"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="噪声预测（dB）"
        prop="noise_predict"
        label-width=""
      >
        <el-input-number
          v-model="lnp.noise_predict"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLnp"
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
import { CreateStdLnpLibraries } from '@/api/standard_libraries/lnp'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LnpCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lnp-list'],
  setup(props, context) {
    const lnpCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lnpCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lnp = reactive({
      rated_capacity_min: 0,
      rated_capacity_max: 0,
      rated_high_vol_min: 0,
      rated_high_vol_max: 0,
      cool_type: '',
      noise_predict: 0
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
    function createLnp() {
      (lnpCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLnpLibraries(lnp)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lnp-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLnp() {
    //   const response = await CreateStdLnpLibraries(lnp)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lnp-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lnpCreateForm,
      show,
      lnp,
      rules,
      createLnp
    }
  }
})
</script>
