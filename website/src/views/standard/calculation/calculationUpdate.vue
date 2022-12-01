<template>
  <el-dialog
    v-model="show"
    title="修改公式参数标准库"
    show-close
    :close-on-click-modal="false"
  >
    <el-form
      ref="calculationUpdateForm"
      :model="tmpCalculation"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="公式中文名"
        prop="formula_name_zh"
        label-width=""
      >
        <el-input
          v-model="tmpCalculation.formula_name_zh"
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
        <div
          ref="formulaExpress"
          class="formula"
        />
      </el-form-item>
<!--      <el-form-item-->
<!--        label="公式名称"-->
<!--        prop="formula_name"-->
<!--        label-width=""-->
<!--      >-->
<!--        <el-input-->
<!--          v-model="tmpCalculation.formula_name"-->
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
<!--          v-model="tmpCalculation.formula_description"-->
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
<!--          v-model="tmpCalculation.formula_param"-->
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
        <div
          ref="formulaRemake"
          class="formula"
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateCalculation"
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
import { computed, defineComponent, ref, nextTick, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Calculation, UpdateCalculation } from '@/api/standard_libraries/calculation'
import { MathfieldElement } from 'mathlive'

export default defineComponent({
  name: 'CalculationUpdate',
  props: {
    modelValue: Boolean,
    calculation: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-calculation-list'],
  setup (props, context) {
    const calculationUpdateForm = ref(null)
    const formulaExpress = ref(null)
    const formulaExpressEl = new MathfieldElement({
      virtualKeyboardMode: 'onfocus'
    })
    const formulaRemake = ref(null)
    const formulaRemakeEl = new MathfieldElement({
      virtualKeyboardMode: 'onfocus'
    })
    const calculation: Calculation = {
      id: 0,
      formula_name: '',
      formula_name_zh: '',
      formula_express: '',
      formula_description: '',
      formula_param: '',
      remark: ''
    }
    const tmpCalculation = ref(calculation)
    watch(props, () => {
      tmpCalculation.value.id = props.calculation.id!
      tmpCalculation.value.formula_name = props.calculation.formula_name
      tmpCalculation.value.formula_name_zh = props.calculation.formula_name_zh
      tmpCalculation.value.formula_express = props.calculation.formula_express
      tmpCalculation.value.formula_description = props.calculation.formula_description
      tmpCalculation.value.formula_param = props.calculation.formula_param
      tmpCalculation.value.remark = props.calculation.remark
      if (props.modelValue) {
        nextTick(() => {
          formulaExpressEl.setValue(tmpCalculation.value.formula_express);
          (formulaExpress.value as any).appendChild(formulaExpressEl)
          formulaRemakeEl.setValue(tmpCalculation.value.remark);
          (formulaRemake.value as any).appendChild(formulaRemakeEl)
        })
      }
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (calculationUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    async function updateCalculation() {
      tmpCalculation.value.formula_express = formulaExpressEl.value
      tmpCalculation.value.remark = formulaRemakeEl.value
      const response = await UpdateCalculation(tmpCalculation.value)
      if (response.data.code === 200) {
        ElMessage.success('信息更新成功')
        context.emit('get-calculation-list')
        show.value = false
      }
    }
    return {
      calculationUpdateForm,
      formulaExpress,
      formulaRemake,
      show,
      tmpCalculation,
      updateCalculation
    }
  }
})
</script>

<style lang="scss">
.formula math-field {
  border-radius: 5px;
  padding: 0 8px;
  border: 1px solid rgba(0, 0, 0, .1);
}
.ML__keyboard {
  z-index: 9999;
}
.ML__keyboard .ML__keyboard--plate div .rows>ul>li aside {
  display: none;
}
</style>
