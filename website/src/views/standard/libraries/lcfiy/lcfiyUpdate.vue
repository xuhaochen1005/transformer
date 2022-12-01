<template>
  <el-dialog
    v-model="show"
    title="修改线圈端部距铁轭距离标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lcfiyUpdateForm"
      :model="tmpLcfiy"
      :rules="rules"
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
          v-model="tmpLcfiy.rated_high_vol_min"
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
          v-model="tmpLcfiy.rated_high_vol_max"
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
          v-model="tmpLcfiy.rated_low_vol_min"
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
          v-model="tmpLcfiy.rated_low_vol_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="低压线圈端部距铁轭距离(mm)"
        prop="low_vol_coil_from_iron_yoke"
        label-width=""
      >
        <el-input-number
          v-model="tmpLcfiy.low_vol_coil_from_iron_yoke"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="高压线圈端部距铁轭距离(mm)"
        prop="high_vol_coil_from_iron_yoke"
        label-width=""
      >
        <el-input-number
          v-model="tmpLcfiy.high_vol_coil_from_iron_yoke"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLcfiy"
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
import { Lcfiy, UpdateStdlcfiyLibraries } from '@/api/standard_libraries/lcfiy'
export default defineComponent({
  name: 'LcfiyUpdate',
  props: {
    modelValue: Boolean,
    lcfiy: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lcfiy-list'],
  setup (props, context) {
    const lcfiyUpdateForm = ref(null)
    const lcfiy: Lcfiy = {
      id: 0,
      rated_high_vol_min: null,
      rated_high_vol_max: null,
      rated_low_vol_min: null,
      rated_low_vol_max: null,
      low_vol_coil_from_iron_yoke: null,
      high_vol_coil_from_iron_yoke: null
    }
    const tmpLcfiy = ref(lcfiy)
    watch(props, () => {
      tmpLcfiy.value.id = props.lcfiy.id!
      tmpLcfiy.value.rated_high_vol_min = props.lcfiy.rated_high_vol_min
      tmpLcfiy.value.rated_high_vol_max = props.lcfiy.rated_high_vol_max
      tmpLcfiy.value.rated_low_vol_min = props.lcfiy.rated_low_vol_min
      tmpLcfiy.value.rated_low_vol_max = props.lcfiy.rated_low_vol_max
      tmpLcfiy.value.low_vol_coil_from_iron_yoke = props.lcfiy.low_vol_coil_from_iron_yoke
      tmpLcfiy.value.high_vol_coil_from_iron_yoke = props.lcfiy.high_vol_coil_from_iron_yoke
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lcfiyUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lcfiy.rated_high_vol_min === null) {
        return (new Error('请输入额定高压下界（kV）：'))
      }
      if (lcfiy.rated_high_vol_min < 0) {
        return callback(new Error('额定高压下界必须大于0！'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lcfiy.rated_low_vol_min === null) {
        return (new Error('请输入额定低压下界（kV）：'))
      }
      if (lcfiy.rated_low_vol_min < 0) {
        return callback(new Error('额定低压下界必须大于0！'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lcfiy.rated_high_vol_max === null) {
        return (new Error('请输入额定高压上界（kV）：'))
      }
      if (lcfiy.rated_high_vol_max < 0) {
        return callback(new Error('额定高压上界必须大于0！'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lcfiy.rated_low_vol_max === null) {
        return (new Error('请输入额定低压上界（kV）：'))
      }
      if (lcfiy.rated_low_vol_max < 0) {
        return callback(new Error('额定低压上界必须大于0！'))
      }
      callback()
    }

    function validateFactor_5(rule: any, value: any, callback: any) {
      if (lcfiy.low_vol_coil_from_iron_yoke === null) {
        return (new Error('请输入低压线圈端部距铁轭距离（mm）：'))
      }
      if (lcfiy.low_vol_coil_from_iron_yoke < 0) {
        return callback(new Error('低压线圈端部距铁轭距离必须大于0！'))
      }
      callback()
    }

    function validateFactor_6(rule: any, value: any, callback: any) {
      if (lcfiy.high_vol_coil_from_iron_yoke === null) {
        return (new Error('请输入高压线圈端部距铁轭距离（mm）：'))
      }
      if (lcfiy.high_vol_coil_from_iron_yoke < 0) {
        return callback(new Error('高压线圈端部距铁轭距离必须大于0！'))
      }
      callback()
    }

    const rules = {
      rated_high_vol_min: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      rated_low_vol_min: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      rated_high_vol_max: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      rated_low_vol_max: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ],
      low_vol_coil_from_iron_yoke: [
        { required: true, validator: validateFactor_5, trigger: 'blur' }
      ],
      high_vol_coil_from_iron_yoke: [
        { required: true, validator: validateFactor_6, trigger: 'blur' }
      ]
    }
    function updateLcfiy() {
      (lcfiyUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlcfiyLibraries(tmpLcfiy.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lcfiy-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    return {
      lcfiyUpdateForm,
      show,
      tmpLcfiy,
      rules,
      updateLcfiy
    }
  }
})
</script>
