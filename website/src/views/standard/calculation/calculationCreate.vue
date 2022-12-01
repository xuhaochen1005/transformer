<template>
  <el-dialog
    v-model="show"
    title="新增公式参数标准库"
    show-close
  >
    <el-form
      ref="calculationCreateForm"
      :rules="rules"
      :model="calculation"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="公式名"
        prop="formula_name_zh"
        label-width=""
      >
        <el-input
          v-model="calculation.formula_name_zh"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="公式表达式"
        prop="formula_express"
        label-width=""
      >
        <el-input
          v-model="calculation.formula_express"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
<!--      <el-form-item-->
<!--        label="公式英文"-->
<!--        prop="formula_name"-->
<!--        label-width=""-->
<!--      >-->
<!--        <el-input-->
<!--          v-model="calculation.formula_name"-->
<!--          size="medium"-->
<!--          style="width: 100%"-->
<!--          clearable-->
<!--        />-->
<!--      </el-form-item>-->
<!--      <el-form-item-->
<!--        label="公式描述"-->
<!--        prop="formula_description"-->
<!--        label-width=""-->
<!--      >-->
<!--        <el-input-->
<!--          v-model="calculation.formula_description"-->
<!--          size="medium"-->
<!--          style="width: 100%"-->
<!--          clearable-->
<!--        />-->
<!--      </el-form-item>-->
<!--      <el-form-item-->
<!--        label="公式参数"-->
<!--        prop="formula_param"-->
<!--        label-width=""-->
<!--      >-->
<!--        <el-input-->
<!--          v-model="calculation.formula_param"-->
<!--          size="medium"-->
<!--          style="width: 100%"-->
<!--          clearable-->
<!--        />-->
<!--      </el-form-item>-->
      <el-form-item
        label="备注"
        prop="remark"
        label-width=""
      >
        <el-input
          v-model="calculation.remark"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createCalculation"
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
import { CreateCalculation } from '@/api/standard_libraries/calculation'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'CalculationCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-calculation-list'],
  setup(props, context) {
    const calculationCreateForm = ref(null)
    const calculation = reactive({
      id: 0,
      formula_name_zh: '',
      formula_express: '',
      formula_name: '',
      formula_description: '',
      formula_param: '',
      remark: ''
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (calculationCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor(rule: any, value: any, callback: any) {
      if (!calculation.formula_name_zh) {
        return callback(new Error('请输入公式名：'))
      }
      if (calculation.formula_name_zh == null) {
        return callback(new Error('公式名不能为空！'))
      }
      callback()
    }
    const rules = {
      formula_name_zh: [
        { required: true, validator: validateFactor, trigger: 'blur' }
      ]

    }
    function createCalculation() {
      (calculationCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateCalculation(calculation)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-calculation-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLfp() {
    //   const response = await CreateStdlfpLibraries(lfp)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lfp-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      calculationCreateForm,
      show,
      calculation,
      rules,
      createCalculation
    }
  }
})
</script>
