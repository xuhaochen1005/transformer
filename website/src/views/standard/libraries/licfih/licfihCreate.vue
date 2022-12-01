<template>
  <el-dialog
    v-model="show"
    title="新增内线圈至铁芯距离标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="licfihCreateForm"
      :rules="rules"
      :model="licfih"
      label-width="220px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定高压下界(kV)"
        prop="rated_high_vol_min"
        label-width=""
      >
        <el-input-number
          v-model="licfih.rated_high_vol_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定高压上界(kV)"
        prop="rated_high_vol_max"
        label-width=""
      >
        <el-input-number
          v-model="licfih.rated_high_vol_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="内线圈至铁芯距离最小值(mm)"
        prop="inner_coil_form_iron_heart_min"
        label-width=""
      >
        <el-input-number
          v-model="licfih.inner_coil_form_iron_heart_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLicfih"
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
import { CreateStdlicfihLibraries } from '@/api/standard_libraries/licfih'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LicfihCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-licfih-list'],
  setup(props, context) {
    const licfihCreateForm = ref(null)
    const licfih = reactive({
      id: 0,
      inner_coil_form_iron_heart_min: 0,
      rated_high_vol_min: 0,
      rated_high_vol_max: 0
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (licfihCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (licfih.rated_high_vol_min == null) {
        return (new Error('请输入额定高压下界（kV）：'))
      }
      if (licfih.rated_high_vol_min < 0) {
        return (new Error('额定高压下界必须大于0！'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (licfih.rated_high_vol_max == null) {
        return (new Error('请输入额定高压上界（kV）：'))
      }
      if (licfih.rated_high_vol_max < 0) {
        return (new Error('额定高压上界必须大于0！'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (licfih.inner_coil_form_iron_heart_min == null) {
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
    function createLicfih() {
      (licfihCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlicfihLibraries(licfih)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-licfih-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLicfih() {
    //   const response = await CreateStdlicfihLibraries(licfih)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-licfih-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      licfihCreateForm,
      show,
      licfih,
      rules,
      createLicfih
    }
  }
})
</script>
