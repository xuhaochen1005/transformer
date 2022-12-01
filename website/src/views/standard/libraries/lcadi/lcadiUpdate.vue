<template>
  <el-dialog
    v-model="show"
    title="修改线圈和风道绝缘"
    show-close
    width="30%"
  >
    <el-form
      ref="lcadiUpdateForm"
      :model="tmpLcadi"
      :rules="rules"
      label-width="250px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="绕制类型"
        prop="wind_type"
        label-width=""
      >
        <el-select
          v-model="tmpLcadi.wind_type"
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
          v-model="tmpLcadi.coil_inner_insulate"
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
          v-model="tmpLcadi.coil_outer_insulate"
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
          v-model="tmpLcadi.air_duct_thick"
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
          v-model="tmpLcadi.air_duct_in_out_insulate"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLcadi"
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
import { Lcadi, UpdateStdLcadiLibraries } from '@/api/standard_libraries/lcadi'

export default defineComponent({
  name: 'LcadiUpdate',
  props: {
    modelValue: Boolean,
    lcadi: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lcadi-list'],
  setup(props, context) {
    const lcadiUpdateForm = ref(null)
    const lcadi: Lcadi = {
      id: 0,
      wind_type: '',
      coil_inner_insulate: 0,
      coil_outer_insulate: 0,
      air_duct_thick: '',
      air_duct_in_out_insulate: null
    }
    const tmpLcadi = ref(lcadi)
    watch(props, () => {
      tmpLcadi.value.id = props.lcadi.id
      tmpLcadi.value.wind_type = props.lcadi.wind_type
      tmpLcadi.value.coil_inner_insulate = props.lcadi.coil_inner_insulate
      tmpLcadi.value.air_duct_thick = props.lcadi.air_duct_thick
      tmpLcadi.value.air_duct_in_out_insulate = props.lcadi.air_duct_in_out_insulate
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lcadiUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
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

    function updateLcadi() {
      (lcadiUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLcadiLibraries(tmpLcadi.value)
          if (response.data.code === 200) {
            ElMessage.success('更新成功')
            context.emit('get-lcadi-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLcadi() {
    //   const response = await UpdateStdLcadiLibraries(tmpLcadi.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lcadi-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lcadiUpdateForm,
      show,
      tmpLcadi,
      rules,
      updateLcadi
    }
  }
})
</script>
