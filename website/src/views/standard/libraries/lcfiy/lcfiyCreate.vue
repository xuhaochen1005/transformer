<template>
  <el-dialog
    v-model="show"
    title="新增线圈端部距铁轭距离标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lcfiyCreateForm"
      :rules="rules"
      :model="lcfiy"
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
          v-model="lcfiy.rated_high_vol_min"
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
          v-model="lcfiy.rated_high_vol_max"
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
          v-model="lcfiy.rated_low_vol_min"
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
          v-model="lcfiy.rated_low_vol_max"
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
          v-model="lcfiy.low_vol_coil_from_iron_yoke"
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
          v-model="lcfiy.high_vol_coil_from_iron_yoke"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLcfiy"
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
import { CreateStdlcfiyLibraries } from '@/api/standard_libraries/lcfiy'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LcfiyCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lcfiy-list'],
  setup(props, context) {
    const lcfiyCreateForm = ref(null)
    const lcfiy = reactive({
      id: 0,
      rated_high_vol_min: 0,
      rated_high_vol_max: 0,
      rated_low_vol_min: 0,
      rated_low_vol_max: 0,
      low_vol_coil_from_iron_yoke: 0,
      high_vol_coil_from_iron_yoke: 0
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lcfiyCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lcfiy.rated_high_vol_min == null) {
        return (new Error('请输入额定高压下界（kV）：'))
      }
      if (lcfiy.rated_high_vol_min < 0) {
        return callback(new Error('额定高压下界必须大于0！'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lcfiy.rated_low_vol_min == null) {
        return (new Error('请输入额定低压下界（kV）：'))
      }
      if (lcfiy.rated_low_vol_min < 0) {
        return callback(new Error('额定低压下界必须大于0！'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lcfiy.rated_high_vol_max == null) {
        return (new Error('请输入额定高压上界（kV）：'))
      }
      if (lcfiy.rated_high_vol_max < 0) {
        return callback(new Error('额定高压上界必须大于0！'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lcfiy.rated_low_vol_max == null) {
        return (new Error('请输入额定低压上界（kV）：'))
      }
      if (lcfiy.rated_low_vol_max < 0) {
        return callback(new Error('额定低压上界必须大于0！'))
      }
      callback()
    }

    function validateFactor_5(rule: any, value: any, callback: any) {
      if (lcfiy.low_vol_coil_from_iron_yoke == null) {
        return (new Error('请输入低压线圈端部距铁轭距离（mm）：'))
      }
      if (lcfiy.low_vol_coil_from_iron_yoke < 0) {
        return callback(new Error('低压线圈端部距铁轭距离必须大于0！'))
      }
      callback()
    }

    function validateFactor_6(rule: any, value: any, callback: any) {
      if (lcfiy.high_vol_coil_from_iron_yoke == null) {
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
    function createLcfiy() {
      (lcfiyCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlcfiyLibraries(lcfiy)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lcfiy-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLcfiy() {
    //   const response = await CreateStdlcfiyLibraries(lcfiy)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lcfiy-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lcfiyCreateForm,
      show,
      lcfiy,
      rules,
      createLcfiy
    }
  }
})
</script>
