<template>
  <el-dialog
    v-model="show"
    title="新增线圈和风道绝缘"
    show-close
    width="30%"
  >
    <el-form
      ref="lcadiCreateForm"
      :rules="rules"
      :model="lcadi"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="绕制类型"
        prop="wind_type"
        label-width=""
      >
        <el-select
          v-model="lcadi.wind_type"
          filterable
          allow-create
          size="medium"
          style="width: 100%"
          clearable
        >
          <el-option
            label="箔绕"
            value="箔绕"
          />
          <el-option
            label="线绕"
            value="线绕"
          />
        </el-select>

      </el-form-item>
      <el-form-item
        label="线圈内层绝缘（mm）"
        prop="coil_inner_insulate"
        label-width=""
      >
        <el-input-number
          v-model="lcadi.coil_inner_insulate"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="线圈外层绝缘（mm）"
        prop="coil_outer_insulate"
        label-width=""
      >
        <el-input-number
          v-model="lcadi.coil_outer_insulate"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="风道厚度可选值（mm）"
        prop="air_duct_thick"
        label-width=""
      >
        <el-input
          v-model="lcadi.air_duct_thick"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="风道内外层绝缘（mm）"
        prop="air_duct_in_out_insulate"
        label-width=""
      >
        <el-input-number
          v-model="lcadi.air_duct_in_out_insulate"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLcadi"
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
import { CreateStdLcadiLibraries } from '@/api/standard_libraries/lcadi'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LcadiCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lcadi-list'],
  setup(props, context) {
    const lcadiCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lcadiCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lcadi = reactive({
      wind_type: '',
      coil_inner_insulate: 0,
      coil_outer_insulate: 0,
      air_duct_thick: '',
      air_duct_in_out_insulate: 0
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!lcadi.wind_type) {
        return callback(new Error('请输入绕制类型'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lcadi.coil_inner_insulate == null) {
        return callback(new Error('请输入线圈内绝缘距离（mm）'))
      }
      if (lcadi.coil_inner_insulate < 0) {
        return callback(new Error('线圈内绝缘距离不能小于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lcadi.coil_outer_insulate == null) {
        return callback(new Error('请输入线圈外绝缘距离（mm）'))
      }
      if (lcadi.coil_outer_insulate < 0) {
        return callback(new Error('线圈外绝缘距离不能小于0'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lcadi.air_duct_thick == null) {
        return callback(new Error('请输入风道条厚度可选值（mm）'))
      }
      callback()
    }

    function validateFactor_5(rule: any, value: any, callback: any) {
      if (lcadi.air_duct_in_out_insulate == null) {
        return callback(new Error('请输入风道内外层绝缘（mm）'))
      }
      if (lcadi.air_duct_in_out_insulate < 0) {
        return callback(new Error('风道内外层绝缘不能小于0'))
      }
      callback()
    }

    const rules = {
      wind_type: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      coil_inner_insulate: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      coil_outer_insulate: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      air_duct_thick: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ],
      air_duct_in_out_insulate: [
        { required: true, validator: validateFactor_5, trigger: 'blur' }
      ],
    }
    function createLcadi() {
      (lcadiCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLcadiLibraries(lcadi)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lcadi-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLcadi() {
    //   const response = await CreateStdLcadiLibraries(lcadi)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lcadi-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lcadiCreateForm,
      show,
      lcadi,
      rules,
      createLcadi
    }
  }
})
</script>
