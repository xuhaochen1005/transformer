<template>
  <el-dialog
    v-model="show"
    title="修改内线圈至铁芯距离标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="licfihUpdateForm"
      :model="tmpLicfih"
      :rules="rules"
      label-width="250px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定高压下界(kV)"
        prop="rated_high_vol_min"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicfih.rated_high_vol_min"
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
          v-model="tmpLicfih.rated_high_vol_max"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="内线圈至铁芯距离最小值(mm)"
        prop="inner_coil_form_iron_heart_min"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicfih.inner_coil_form_iron_heart_min"
          style="width:100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLicfih"
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
import { Licfih, UpdateStdlicfihLibraries } from '@/api/standard_libraries/licfih'

export default defineComponent({
  name: 'LicfihUpdate',
  props: {
    modelValue: Boolean,
    licfih: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-licfih-list'],
  setup (props, context) {
    const licfihUpdateForm = ref(null)
    const licfih: Licfih = {
      id: 0,
      inner_coil_form_iron_heart_min: null,
      rated_high_vol_min: null,
      rated_high_vol_max: null
    }
    const tmpLicfih = ref(licfih)
    watch(props, () => {
      tmpLicfih.value.id = props.licfih.id!
      tmpLicfih.value.inner_coil_form_iron_heart_min = props.licfih.inner_coil_form_iron_heart_min
      tmpLicfih.value.rated_high_vol_min = props.licfih.rated_high_vol_min
      tmpLicfih.value.rated_high_vol_max = props.licfih.rated_high_vol_max
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (licfihUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (licfih.rated_high_vol_min === null) {
        return (new Error('请输入额定高压下界（kV）：'))
      }
      if (licfih.rated_high_vol_min < 0) {
        return (new Error('额定高压下界必须大于0！'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (licfih.rated_high_vol_max === null) {
        return (new Error('请输入额定高压上界（kV）：'))
      }
      if (licfih.rated_high_vol_max < 0) {
        return (new Error('额定高压上界必须大于0！'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (licfih.inner_coil_form_iron_heart_min === null) {
        return (new Error('请输入内线圈至铁芯最小距离（mm）：'))
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
      inner_coil_form_iron_heart_min: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ]
    }
    function updateLicfih() {
      (licfihUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlicfihLibraries(tmpLicfih.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-licfih-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLicfih() {
    //   const response = await UpdateStdlicfihLibraries(tmpLicfih.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-licfih-list')
    //     show.value = false
    //   }
    // }
    return {
      licfihUpdateForm,
      show,
      tmpLicfih,
      rules,
      updateLicfih
    }
  }
})
</script>
