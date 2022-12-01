<template>
  <el-dialog
    v-model="show"
    title="修改硅钢片性能"
    show-close
    width="30%"
  >
    <el-form
      ref="lspdUpdateForm"
      :model="tmpLspd"
      :rules="rules"
      label-width="250px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="硅钢片型号"
        prop="stack_type"
        label-width=""
      >
        <el-input
          v-model="tmpLspd.stack_type"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="铁心磁密（T）"
        prop="core_flux_density"
        label-width=""
      >
        <el-input-number
          v-model="tmpLspd.core_flux_density"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="单位铁损（W/kg）"
        prop="unit_iron_loss"
        label-width=""
      >
        <el-input-number
          v-model="tmpLspd.unit_iron_loss"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="单位质量磁化容量（VA/kg）"
        prop="unit_mass_magnet_capacity"
        label-width=""
      >
        <el-input-number
          v-model="tmpLspd.unit_mass_magnet_capacity"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="单位面积接缝伏安值（VA/cm2）"
        prop="unit_area_seam_va"
        label-width=""
      >
        <el-input-number
          v-model="tmpLspd.unit_area_seam_va"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLspd"
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
import { Lspd, UpdateStdLspdLibraries } from '@/api/standard_libraries/lspd'

export default defineComponent({
  name: 'LspdUpdate',
  props: {
    modelValue: Boolean,
    lspd: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lspd-list'],
  setup(props, context) {
    const lspdUpdateForm = ref(null)
    const lspd: Lspd = {
      id: 0,
      stack_type: '',
      core_flux_density: 0,
      unit_iron_loss: 0,
      unit_mass_magnet_capacity: 0,
      unit_area_seam_va: 0
    }
    const tmpLspd = ref(lspd)
    watch(props, () => {
      tmpLspd.value.id = props.lspd.id
      tmpLspd.value.stack_type = props.lspd.stack_type
      tmpLspd.value.core_flux_density = props.lspd.core_flux_density
      tmpLspd.value.unit_mass_magnet_capacity = props.lspd.unit_mass_magnet_capacity
      tmpLspd.value.unit_area_seam_va = props.lspd.unit_area_seam_va
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lspdUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
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
    function updateLspd() {
      (lspdUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLspdLibraries(tmpLspd.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lspd-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLspd() {
    //   const response = await UpdateStdLspdLibraries(tmpLspd.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lspd-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lspdUpdateForm,
      show,
      tmpLspd,
      rules,
      updateLspd
    }
  }
})
</script>
