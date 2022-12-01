<template>
  <el-dialog
    v-model="show"
    title="修改导线绝缘厚度表标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lwiUpdateForm"
      :model="tmpLwi"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="导线编号"
        prop="wire_id"
        label-width=""
      >
        <el-input
          v-model="tmpLwi.wire_id"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="线厚区间下界(mm)"
        prop="width_min"
        label-width=""
      >
        <el-input-number
          v-model="tmpLwi.width_min"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="线厚区间上界(mm)"
        prop="width_max"
        label-width=""
      >
        <el-input-number
          v-model="tmpLwi.width_max"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="轴向绝缘厚(mm)"
        prop="axial_insulation"
        label-width=""
      >
        <el-input-number
          v-model="tmpLwi.axial_insulation"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="辐向绝缘厚(mm)"
        prop="radial_insulation"
        label-width=""
      >
        <el-input-number
          v-model="tmpLwi.radial_insulation"
          style="width:100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLwi"
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
import { Lwi, UpdateStdlwiLibraries } from '@/api/standard_libraries/lwi'

export default defineComponent({
  name: 'LwiUpdate',
  props: {
    modelValue: Boolean,
    lwi: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lwi-list'],
  setup (props, context) {
    const lwiUpdateForm = ref(null)
    const lwi: Lwi = {
      id: 0,
      wire_id: null,
      width_min: null,
      width_max: null,
      axial_insulation: null,
      radial_insulation: null
    }
    const tmpLwi = ref(lwi)
    watch(props, () => {
      tmpLwi.value.id = props.lwi.id!
      tmpLwi.value.wire_id = props.lwi.wire_id
      tmpLwi.value.width_min = props.lwi.width_min
      tmpLwi.value.width_max = props.lwi.width_max
      tmpLwi.value.axial_insulation = props.lwi.axial_insulation
      tmpLwi.value.radial_insulation = props.lwi.radial_insulation
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lwiUpdateForm.value as any).clearValidate()
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
    function updateLwi() {
      (lwiUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlwiLibraries(tmpLwi.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lwi-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLwi() {
    //   const response = await UpdateStdlwiLibraries(tmpLwi.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lwi-list')
    //     show.value = false
    //   }
    // }
    return {
      lwiUpdateForm,
      show,
      tmpLwi,
      rules,
      updateLwi
    }
  }
})
</script>
