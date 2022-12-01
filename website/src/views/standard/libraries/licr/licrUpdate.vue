<template>
  <el-dialog
    v-model="show"
    title="修改铁心表(圆形)标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="licrUpdateForm"
      :model="tmpLicr"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="铁心直径"
        prop="core_diameter"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.core_diameter"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="有无拉板"
        prop="pull_plate"
        label-width=""
      >
        <el-select
          v-model="tmpLicr.pull_plate"
          style="width: 100%"
        >
          <el-option
            v-for="item in pullplateOptions"
            :key="item"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="截面积(mm)"
        prop="section_area"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.section_area"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="角重(mm)"
        prop="iron_corner_weight"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.iron_corner_weight"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="片宽(1级)(mm)"
        prop="stack_width_1"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_width_1"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="片宽(1级)(mm)"
        prop="stack_width_1"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_width_1"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="片宽(2级)(mm)"
        prop="stack_width_2"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_width_2"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="片宽(3级)(mm)"
        prop="stack_width_3"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_width_3"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="片宽(4级)(mm)"
        prop="stack_width_4"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_width_4"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="片宽(5级)(mm)"
        prop="stack_width_5"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_width_5"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="片宽(6级)(mm)"
        prop="stack_width_6"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_width_6"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="片宽(7级)(mm)"
        prop="stack_width_7"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_width_7"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="片宽(8级)(mm)"
        prop="stack_width_8"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_width_8"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="迭厚(1级)(mm)"
        prop="stack_thick_1"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_thick_1"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="迭厚(2级)(mm)"
        prop="stack_thick_2"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_thick_2"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="迭厚3级)(mm)"
        prop="stack_thick_3"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_thick_3"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="迭厚(4级)(mm)"
        prop="stack_thick_4"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_thick_4"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="迭厚(5级)(mm)"
        prop="stack_thick_5"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_thick_5"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="迭厚(6级)(mm)"
        prop="stack_thick_6"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_thick_6"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="迭厚(7级)(mm)"
        prop="stack_thick_7"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_thick_7"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="迭厚(8级)(mm)"
        prop="stack_thick_8"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_thick_8"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="总迭厚(mm)"
        prop="stack_thick_sum"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.stack_thick_sum"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="主级实迭厚迭厚(mm)"
        prop="main_level_real_stack_thick"
        label-width=""
      >
        <el-input-number
          v-model="tmpLicr.main_level_real_stack_thick"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLicr"
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
import { computed, defineComponent, reactive, ref, watch, Ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Licr, LicrPullPlate, UpdateStdlicrLibraries } from '@/api/standard_libraries/licr'

export default defineComponent({
  name: 'LicrUpdate',
  props: {
    modelValue: Boolean,
    licr: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-licr-list'],
  setup (props, context) {
    const pullplateOptions = [
      { label: '有', value: LicrPullPlate.Yes },
      { label: '无', value: LicrPullPlate.No }
    ]
    const licrUpdateForm = ref(null)
    const licr: Licr = {
      id: 0,
      core_diameter: null,
      pull_plate: LicrPullPlate.Yes,
      section_area: null,
      iron_corner_weight: null,
      stack_width_1: null,
      stack_width_2: null,
      stack_width_3: null,
      stack_width_4: null,
      stack_width_5: null,
      stack_width_6: null,
      stack_width_7: null,
      stack_width_8: null,
      stack_thick_1: null,
      stack_thick_2: null,
      stack_thick_3: null,
      stack_thick_4: null,
      stack_thick_5: null,
      stack_thick_6: null,
      stack_thick_7: null,
      stack_thick_8: null,
      stack_thick_sum: null,
      main_level_real_stack_thick: null
    }
    const tmpLicr = ref(licr)
    watch(props, () => {
      tmpLicr.value.id = props.licr.id!
      tmpLicr.value.core_diameter = props.licr.core_diameter
      tmpLicr.value.pull_plate = props.licr.pull_plate
      tmpLicr.value.section_area = props.licr.section_area
      tmpLicr.value.iron_corner_weight = props.licr.iron_corner_weight
      tmpLicr.value.stack_width_1 = props.licr.stack_width_1
      tmpLicr.value.stack_width_2 = props.licr.stack_width_2
      tmpLicr.value.stack_width_3 = props.licr.stack_width_3
      tmpLicr.value.stack_width_4 = props.licr.stack_width_4
      tmpLicr.value.stack_width_5 = props.licr.stack_width_5
      tmpLicr.value.stack_width_6 = props.licr.stack_width_6
      tmpLicr.value.stack_width_7 = props.licr.stack_width_7
      tmpLicr.value.stack_width_8 = props.licr.stack_width_8
      tmpLicr.value.stack_thick_1 = props.licr.stack_thick_1
      tmpLicr.value.stack_thick_2 = props.licr.stack_thick_2
      tmpLicr.value.stack_thick_3 = props.licr.stack_thick_3
      tmpLicr.value.stack_thick_4 = props.licr.stack_thick_4
      tmpLicr.value.stack_thick_5 = props.licr.stack_thick_5
      tmpLicr.value.stack_thick_6 = props.licr.stack_thick_6
      tmpLicr.value.stack_thick_7 = props.licr.stack_thick_7
      tmpLicr.value.stack_thick_8 = props.licr.stack_thick_8
      tmpLicr.value.stack_thick_sum = props.licr.stack_thick_sum
      tmpLicr.value.main_level_real_stack_thick = props.licr.main_level_real_stack_thick
    })
    const show = computed({
      get: () => props.modelValue,
      set: (v) => context.emit('update:modelValue', v)
    })

    function validateFactor(rule: any, value: any, callback: any) {
      if (licr.core_diameter === null) {
        return callback(new Error('请输入铁心直径(mm)：'))
      }
      if (licr.core_diameter <= 0) {
        return callback(new Error('铁心直径必须大于0！'))
      }
      if (licr.pull_plate === null) {
        return callback(new Error('请选择有无拉板：'))
      }
      if (licr.section_area === null) {
        return callback(new Error('请输入铁心截面积(mm)：'))
      }
      if (licr.section_area <= 0) {
        return callback(new Error('铁心截面积必须大于0！'))
      }
      if (licr.iron_corner_weight === null) {
        return callback(new Error('请输入角重(kg)：'))
      }
      if (licr.iron_corner_weight <= 0) {
        return callback(new Error('角重必须大于0！'))
      }
      callback()
    }

    const rules = {
      core_diameter: [
        { required: true, validator: validateFactor, trigger: 'blur' }
      ],
      pull_plate: [
        { required: true, validator: validateFactor, trigger: 'blur' }
      ],
      section_area: [
        { required: true, validator: validateFactor, trigger: 'blur' }
      ],
      iron_corner_weight: [
        { required: true, validator: validateFactor, trigger: 'blur' }
      ]
    }
    function updateLicr() {
      (licrUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          console.log(tmpLicr)
          const response = await UpdateStdlicrLibraries(tmpLicr.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-licr-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLicr() {
    //   const response = await UpdateStdlicrLibraries(tmpLicr.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-licr-list')
    //     show.value = false
    //   }
    // }
    return {
      pullplateOptions,
      licrUpdateForm,
      show,
      tmpLicr,
      rules,
      updateLicr
    }
  }
})
</script>
