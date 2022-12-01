<template>
  <el-dialog
    v-model="show"
    title="修改低压外模台账"
    show-close
    width="30%"
  >
    <el-form
      ref="lmadUpdateForm"
      :model="tmpLmad"
      :rules="rules"
      label-width="100px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定高压下界（kV）"
        prop="rated_high_vol_min"
        label-width=""
      >
        <el-input-number
          v-model="tmpLmad.rated_high_vol_min"
          size="medium"
          style="width: 50%"
          placeholder="额定高压下界（kV）"
          clearable
        >
        </el-input-number>
      </el-form-item>
      <el-form-item
        label="额定高压上界（kV）"
        prop="rated_high_vol_max"
        label-width=""
      >
        <el-input-number
          v-model="tmpLmad.rated_high_vol_max"
          size="medium"
          style="width: 50%"
          placeholder="额定高压上界（kV）"
          clearable
        >
        </el-input-number>
      </el-form-item>
      <el-form-item
        label="主风道初选最小值（mm）"
        prop="main_air_duct_min"
        label-width=""
      >
        <el-input-number
          v-model="tmpLmad.main_air_duct_min"
          size="medium"
          style="width: 50%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLmad"
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
import { Lmad, UpdateStdLmadLibraries } from '@/api/standard_libraries/lmad'

export default defineComponent({
  name: 'LmadUpdate',
  props: {
    modelValue: Boolean,
    lmad: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lmad-list'],
  setup(props, context) {
    const lmadUpdateForm = ref(null)
    const lmad: Lmad = {
      id: 0,
      rated_high_vol_min: 0,
      rated_high_vol_max: 0,
      main_air_duct_min: 0
    }
    const tmpLmad = ref(lmad)
    watch(props, () => {
      tmpLmad.value.id = props.lmad.id
      tmpLmad.value.rated_high_vol_min = props.lmad.rated_high_vol_min
      tmpLmad.value.rated_high_vol_max = props.lmad.rated_high_vol_max
      tmpLmad.value.main_air_duct_min = props.lmad.main_air_duct_min
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lmadUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lmad.rated_high_vol_min == null) {
        return callback(new Error('请输入额定容量下界（kVA）'))
      }
      if (lmad.rated_high_vol_min < 0) {
        return callback(new Error('额定容量下界不能小于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lmad.rated_high_vol_max == null) {
        return callback(new Error('请输入额定容量上界（kVA）'))
      }
      if (lmad.rated_high_vol_max < 0) {
        return callback(new Error('额定容量上界不能小于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lmad.main_air_duct_min == null) {
        return callback(new Error('请输入主风道初选最小值（mm）'))
      }
      if (lmad.main_air_duct_min < 0) {
        return callback(new Error('主风道初选最小值不能小于0'))
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
      main_air_duct_min: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],

    }
    function updateLmad() {
      (lmadUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLmadLibraries(tmpLmad.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lmad-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLmad() {
    //   const response = await UpdateStdLmadLibraries(tmpLmad.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lmad-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lmadUpdateForm,
      show,
      tmpLmad,
      rules,
      updateLmad
    }
  }
})
</script>
