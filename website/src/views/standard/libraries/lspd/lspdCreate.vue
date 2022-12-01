<template>
  <el-dialog
    v-model="show"
    title="新增硅钢片性能"
    show-close
    width="30%"
  >
    <el-form
      ref="lspdCreateForm"
      :model="lspd"
      :rules="rules"
      label-width="250px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="硅钢片型号"
        prop="stack_type"
      >
        <el-input
          v-model="lspd.stack_type"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="铁心磁密（T）"
        prop="core_flux_density"
      >
        <el-input-number
          v-model="lspd.core_flux_density"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="单位铁损（W/kg）"
        prop="unit_iron_loss"
      >
        <el-input-number
          v-model="lspd.unit_iron_loss"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="单位质量磁化容量（VA/kg）"
        prop="unit_mass_magnet_capacity"
      >
        <el-input-number
          v-model="lspd.unit_mass_magnet_capacity"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="单位面积接缝伏安值（VA/cm2）"
        prop="unit_area_seam_va"
      >
        <el-input-number
          v-model="lspd.unit_area_seam_va"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLspd"
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
import { CreateStdLspdLibraries } from '@/api/standard_libraries/lspd'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LspdCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lspd-list'],
  setup(props, context) {
    const lspdCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lspdCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lspd = reactive({
      stack_type: '',
      core_flux_density: 0,
      unit_iron_loss: 0,
      unit_mass_magnet_capacity: 0,
      unit_area_seam_va: 0
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!lspd.stack_type) {
        return callback(new Error('请输入硅钢片型号：'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lspd.core_flux_density == null) {
        return callback(new Error('请输入铁心磁密（T）：'))
      }
      if (lspd.core_flux_density < 0) {
        return callback(new Error('铁心磁密不能小于0！'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lspd.unit_iron_loss == null) {
        return callback(new Error('请输入单位铁损（W/kg）：'))
      }
      if (lspd.unit_iron_loss < 0) {
        return callback(new Error('单位铁损不能小于0！'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lspd.unit_mass_magnet_capacity == null) {
        return callback(new Error('请输入单位质量磁化容量（VA/kg）：'))
      }
      if (lspd.unit_mass_magnet_capacity < 0) {
        return callback(new Error('单位质量磁化容量不能小于0！'))
      }
      callback()
    }

    function validateFactor_5(rule: any, value: any, callback: any) {
      if (lspd.unit_area_seam_va == null) {
        return callback(new Error('请输入单位面积接缝伏安值（VA/cm^2）：'))
      }
      if (lspd.unit_area_seam_va < 0) {
        return callback(new Error('单位面积接缝伏安值不能小于0！'))
      }
      callback()
    }

    const rules = {
      stack_type: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      core_flux_density: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      unit_iron_loss: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      unit_mass_magnet_capacity: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ],
      unit_area_seam_va: [
        { required: true, validator: validateFactor_5, trigger: 'blur' }
      ]

    }
    function createLspd() {
      (lspdCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLspdLibraries(lspd)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lspd-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLspd() {
    //   const response = await CreateStdLspdLibraries(lspd)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lspd-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lspdCreateForm,
      show,
      lspd,
      rules,
      createLspd
    }
  }
})
</script>
