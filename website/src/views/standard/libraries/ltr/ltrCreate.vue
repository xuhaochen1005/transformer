<template>
  <el-dialog
    v-model="show"
    title="新增绕组温升限值"
    show-close
    width="30%"
  >
    <el-form
      ref="ltrCreateForm"
      :model="ltr"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="绝缘系统温度（℃）"
        prop="temp"
        label-width=""
      >
        <el-input-number
          v-model="ltr.temp"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="绝缘温度代号"
        prop="temp_sign"
        label-width=""
      >
        <el-input
          v-model="ltr.temp_sign"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="低压温升限值（K）"
        prop="low_vol_temp_rise"
        label-width=""
      >
        <el-input-number
          v-model="ltr.low_vol_temp_rise"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="高压温升限值（K）"
        prop="high_vol_temp_rise"
        label-width=""
      >
        <el-input-number
          v-model="ltr.high_vol_temp_rise"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLtr"
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
import { CreateStdLtrLibraries } from '@/api/standard_libraries/ltr'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LtrCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-ltr-list'],
  setup(props, context) {
    const ltrCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (ltrCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const ltr = reactive({
      temp: 0,
      temp_sign: '',
      low_vol_temp_rise: 0,
      high_vol_temp_rise: 0
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (ltr.temp == null) {
        return callback(new Error('请输入绝缘系统温度（℃）：'))
      }
      if (ltr.temp < 0) {
        return callback(new Error('绝缘系统温度不能小于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (!ltr.temp_sign) {
        return callback(new Error('请输入绝缘系统温度字母代号：'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (ltr.low_vol_temp_rise == null) {
        return callback(new Error('请输入低压温升（K）：'))
      }
      if (ltr.low_vol_temp_rise < 0) {
        return callback(new Error('低压温升不能小于0'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (ltr.high_vol_temp_rise == null) {
        return callback(new Error('请输入高压温升（K）：'))
      }
      if (ltr.high_vol_temp_rise < 0) {
        return callback(new Error('高压温升不能小于0'))
      }
      callback()
    }

    const rules = {
      temp: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      temp_sign: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      low_vol_temp_rise: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      high_vol_temp_rise: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ],

    }
    function createLtr() {
      (ltrCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLtrLibraries(ltr)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-ltr-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLtr() {
    //   const response = await CreateStdLtrLibraries(ltr)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-ltr-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      ltrCreateForm,
      show,
      ltr,
      rules,
      createLtr
    }
  }
})
</script>
