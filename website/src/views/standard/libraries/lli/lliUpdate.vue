<template>
  <el-dialog
    v-model="show"
    title="修改层间绝缘距离标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lliUpdateForm"
      :model="tmpLli"
      :rules="rules"
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
          v-model="tmpLli.winding_type"
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
          v-model="tmpLli.rated_low_vol_min"
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
          v-model="tmpLli.rated_low_vol_max"
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
          v-model="tmpLli.layer_vol_min"
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
          v-model="tmpLli.layer_vol_max"
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
          v-model="tmpLli.layer_insulate"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLli"
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
import { Lli, UpdateStdlliLibraries } from '@/api/standard_libraries/lli'

export default defineComponent({
  name: 'LliUpdate',
  props: {
    modelValue: Boolean,
    lli: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lli-list'],
  setup (props, context) {
    const lliUpdateForm = ref(null)
    const lli: Lli = {
      id: 0,
      winding_type: '',
      rated_low_vol_min: null,
      rated_low_vol_max: null,
      layer_vol_min: null,
      layer_vol_max: null,
      layer_insulate: null
    }
    const tmpLli = ref(lli)
    watch(props, () => {
      tmpLli.value.id = props.lli.id!
      tmpLli.value.winding_type = props.lli.winding_type
      tmpLli.value.rated_low_vol_min = props.lli.rated_low_vol_min
      tmpLli.value.rated_low_vol_max = props.lli.rated_low_vol_max
      tmpLli.value.layer_vol_min = props.lli.layer_vol_min
      tmpLli.value.layer_vol_max = props.lli.layer_vol_max
      tmpLli.value.layer_insulate = props.lli.layer_insulate
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lliUpdateForm.value as any).clearValidate()
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
    function updateLli() {
      (lliUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlliLibraries(tmpLli.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lli-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLli() {
    //   const response = await UpdateStdlliLibraries(tmpLli.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lli-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lliUpdateForm,
      show,
      tmpLli,
      rules,
      updateLli
    }
  }
})
</script>
