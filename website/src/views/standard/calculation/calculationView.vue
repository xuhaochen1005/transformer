<template>
  <el-dialog
    v-model="show"
    title="公式参数标准库"
    width="1000px"
    show-close
  >
    <el-form
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="编号"
        label-width=""
      >
        {{ calculation.id }}
      </el-form-item>
      <el-form-item
        label="公式中文名"
        label-width=""
      >
        {{ calculation.formula_name_zh }}
      </el-form-item>
      <div class="type">
        <el-form-item
          label="公式表达式"
          label-width=""
        >
          <math-jax
            style="text-align: left"
            :latex="calculation.formula_express"
            :block="true"
          />
  <!--        {{ calculation.formula_express }}-->
        </el-form-item>
<!--        <el-form-item-->
<!--          label="公式英文名"-->
<!--          label-width=""-->
<!--        >-->
<!--          {{ calculation.formula_name }}-->
<!--        </el-form-item>-->
<!--        <el-form-item-->
<!--          label="公式描述"-->
<!--          label-width=""-->
<!--        >-->
<!--          {{ calculation.formula_description }}-->
<!--        </el-form-item>-->
<!--        <el-form-item-->
<!--          label="公式参数"-->
<!--          label-width=""-->
<!--        >-->
<!--          {{ calculation.formula_param }}-->
<!--        </el-form-item>-->
        <el-form-item
          label="备注"
          label-width=""
        >
          <math-jax
            align="right"
            :latex="calculation.remark"
            :block="true"
          />
  <!--        {{ calculation.remark }}-->
        </el-form-item>
      </div>
    </el-form>
    <template #footer>
      <el-button @click="show = false">
        关闭
      </el-button>
    </template>
  </el-dialog>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue'

export default defineComponent({
  name: 'CalculationView',
  props: {
    modelValue: Boolean,
    calculation: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue'],
  setup(props, context) {
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        context.emit('update:modelValue', value)
      }
    })
    return {
      show
    }
  }
})
</script>

<style lang="scss">
.type mjx-container[jax="SVG"][display="true"] {
  text-align: left!important;
}
</style>
