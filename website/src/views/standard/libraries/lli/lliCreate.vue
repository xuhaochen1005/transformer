<template>
  <el-dialog
    v-model="show"
    title="新增层间绝缘距离标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lliCreateForm"
      :rules="rules"
      :model="lli"
      label-width="220px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="绕制类型"
        prop="winding_type"
        label-width=""
      >
        <el-input
          v-model="lli.winding_type"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定低压下界(kV)"
        prop="rated_low_vol_min"
        label-width=""
      >
        <el-input-number
          v-model="lli.rated_low_vol_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定低压上界(kV)"
        prop="rated_low_vol_max"
        label-width=""
      >
        <el-input-number
          v-model="lli.rated_low_vol_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="层间电压下界(V)"
        prop="layer_vol_min"
        label-width=""
      >
        <el-input-number
          v-model="lli.layer_vol_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="层间电压上界(V)"
        prop="layer_vol_max"
        label-width=""
      >
        <el-input-number
          v-model="lli.layer_vol_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="层间绝缘距离(mm)"
        prop="layer_insulate"
        label-width=""
      >
        <el-input-number
          v-model="lli.layer_insulate"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLli"
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
import { CreateStdlliLibraries } from '@/api/standard_libraries/lli'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LliCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lli-list'],
  setup(props, context) {
    const lliCreateForm = ref(null)
    const lli = reactive({
      id: 0,
      winding_type: '',
      rated_low_vol_min: 0,
      rated_low_vol_max: 0,
      layer_vol_min: 0,
      layer_vol_max: 0,
      layer_insulate: 0
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          // (lccCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lli.rated_low_vol_min == null) {
        return callback(new Error('请输入额定低压下界(kV)：'))
      }
      if (lli.rated_low_vol_min < 0) {
        return callback(new Error('额定低压下界必须大于0！'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lli.rated_low_vol_max == null) {
        return callback(new Error('请输入额定低压上界(kV)：'))
      }
      if (lli.rated_low_vol_max <= 0) {
        return callback(new Error('额定低压上界必须大于0！'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lli.layer_vol_min == null) {
        return callback(new Error('请输入层间电压下界(V)：'))
      }
      if (lli.layer_vol_min < 0) {
        return callback(new Error('层间电压下界必须大于0！'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lli.layer_vol_max == null) {
        return callback(new Error('请输入层间电压上界(V)：'))
      }
      if (lli.layer_vol_max <= 0) {
        return callback(new Error('层间电压上界必须大于0！'))
      }
      callback()
    }

    function validateFactor_5(rule: any, value: any, callback: any) {
      if (lli.layer_insulate == null) {
        return callback(new Error('请输入层间绝缘距离(mm)：'))
      }
      callback()
    }

    function validateFactor_6(rule: any, value: any, callback: any) {
      if (!lli.winding_type) {
        return callback(new Error('请输入绕制类型：'))
      }
      callback()
    }

    const rules = {
      rated_low_vol_min: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      rated_low_vol_max: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      layer_vol_min: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      layer_vol_max: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ],
      layer_insulate: [
        { required: true, validator: validateFactor_5, trigger: 'blur' }
      ],
      winding_type: [
        { required: true, validator: validateFactor_6, trigger: 'blur' }
      ]
    }
    function createLli() {
      (lliCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlliLibraries(lli)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lli-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLli() {
    //   const response = await CreateStdlliLibraries(lli)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lli-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lliCreateForm,
      show,
      lli,
      rules,
      createLli
    }
  }
})
</script>
