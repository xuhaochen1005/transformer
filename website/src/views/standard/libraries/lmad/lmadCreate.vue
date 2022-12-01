<template>
  <el-dialog
    v-model="show"
    title="新增主风道初选"
    show-close
    width="30%"
  >
    <el-form
      ref="lmadCreateForm"
      :model="lmad"
      :rules="rules"
      label-width="200px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定高压下界（kV）"
        prop="rated_high_vol_min"
        label-width=""
      >
        <el-input-number
          v-model="lmad.rated_high_vol_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定高压上界（kV）"
        prop="rated_high_vol_max"
        label-width=""
      >
        <el-input-number
          v-model="lmad.rated_high_vol_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="主风道初选最小值（mm）"
        prop="main_air_duct_min"
        label-width=""
      >
        <el-input-number
          v-model="lmad.main_air_duct_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLmad"
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
import { CreateStdLmadLibraries } from '@/api/standard_libraries/lmad'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LmadCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lmad-list'],
  setup(props, context) {
    const lmadCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lmadCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lmad = reactive({
      rated_high_vol_min: 0,
      rated_high_vol_max: 0,
      main_air_duct_min: 0
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
    function createLmad() {
      (lmadCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLmadLibraries(lmad)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lmad-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLmad() {
    //   const response = await CreateStdLmadLibraries(lmad)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lmad-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lmadCreateForm,
      show,
      lmad,
      rules,
      createLmad
    }
  }
})
</script>
