<template>
  <el-dialog
    v-model="show"
    title="新增导线绝缘厚度表标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lwiCreateForm"
      :rules="rules"
      :model="lwi"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="导线编号"
        prop="wire_id"
        label-width=""
      >
        <el-input-number
          v-model="lwi.wire_id"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="线厚区间下界(mm)"
        prop="width_min"
        label-width=""
      >
        <el-input-number
          v-model="lwi.width_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="线厚区间上界(mm)"
        prop="width_max"
        label-width=""
      >
        <el-input-number
          v-model="lwi.width_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="轴向绝缘厚(mm)"
        prop="axial_insulation"
        label-width=""
      >
        <el-input-number
          v-model="lwi.axial_insulation"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="辐向绝缘厚(mm)"
        prop="radial_insulation"
        label-width=""
      >
        <el-input-number
          v-model="lwi.radial_insulation"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLwi"
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
import { CreateStdlwiLibraries } from '@/api/standard_libraries/lwi'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LwiCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lwi-list'],
  setup(props, context) {
    const lwiCreateForm = ref(null)
    const lwi = reactive({
      id: 0,
      wire_id: 0,
      width_min: 0,
      width_max: 0,
      axial_insulation: 0,
      radial_insulation: 0
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lwiCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!lwi.wire_id) {
        return callback(new Error('请输入导线编号：'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lwi.width_min == null) {
        return callback(new Error('请输入线厚区间下界（mm）：'))
      }
      if (lwi.width_min < 0) {
        return callback(new Error('线厚区间下界不能小于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lwi.width_max == null) {
        return callback(new Error('请输入线厚区间上界（mm）：'))
      }
      if (lwi.width_max < 0) {
        return callback(new Error('线厚区间上界不能小于0'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lwi.axial_insulation == null) {
        return callback(new Error('请输入轴向绝缘厚（mm）：'))
      }
      if (lwi.axial_insulation < 0) {
        return callback(new Error('轴向绝缘厚不能小于0'))
      }
      callback()
    }

    function validateFactor_5(rule: any, value: any, callback: any) {
      if (lwi.radial_insulation == null) {
        return callback(new Error('请输入辐向绝缘厚（mm）：'))
      }
      if (lwi.radial_insulation < 0) {
        return callback(new Error('辐向绝缘厚不能小于0'))
      }
      callback()
    }

    const rules = {
      wire_id: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      width_min: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],

      width_max: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      axial_insulation: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ],
      radial_insulation: [
        { required: true, validator: validateFactor_5, trigger: 'blur' }
      ],
    }
    function createLwi() {
      (lwiCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlwiLibraries(lwi)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lwi-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLwi() {
    //   const response = await CreateStdlwiLibraries(lwi)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lwi-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lwiCreateForm,
      show,
      lwi,
      rules,
      createLwi
    }
  }
})
</script>
