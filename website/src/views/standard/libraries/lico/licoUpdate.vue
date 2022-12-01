<template>
  <el-dialog
    v-model="show"
    title="修改铁心表(长圆形)标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="licoUpdateForm"
      :model="tmpLico"
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
          v-model="tmpLico.core_diameter"
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
          v-model="tmpLico.pull_plate"
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
          v-model="tmpLico.section_area"
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
          v-model="tmpLico.iron_corner_weight"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="片宽(0级)(mm)"
        prop="stack_width_0"
        label-width=""
      >
        <el-input-number
          v-model="tmpLico.stack_width_0"
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
          v-model="tmpLico.stack_width_1"
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
          v-model="tmpLico.stack_width_1"
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
          v-model="tmpLico.stack_width_2"
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
          v-model="tmpLico.stack_width_3"
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
          v-model="tmpLico.stack_width_4"
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
          v-model="tmpLico.stack_width_5"
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
          v-model="tmpLico.stack_width_6"
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
          v-model="tmpLico.stack_width_7"
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
          v-model="tmpLico.stack_width_8"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="迭厚(0级)(mm)"
        prop="stack_thick_0"
        label-width=""
      >
        <el-input-number
          v-model="tmpLico.stack_thick_0"
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
          v-model="tmpLico.stack_thick_1"
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
          v-model="tmpLico.stack_thick_2"
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
          v-model="tmpLico.stack_thick_3"
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
          v-model="tmpLico.stack_thick_4"
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
          v-model="tmpLico.stack_thick_5"
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
          v-model="tmpLico.stack_thick_6"
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
          v-model="tmpLico.stack_thick_7"
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
          v-model="tmpLico.stack_thick_8"
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
          v-model="tmpLico.stack_thick_sum"
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
          v-model="tmpLico.main_level_real_stack_thick"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLico"
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
import { Lico, LicoPullPlate, UpdateStdlicoLibraries } from '@/api/standard_libraries/lico'

export default defineComponent({
  name: 'LicoUpdate',
  props: {
    modelValue: Boolean,
    lico: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lico-list'],
  setup (props, context) {
    const pullplateOptions = [
      { label: '有', value: LicoPullPlate.Yes },
      { label: '无', value: LicoPullPlate.No }
    ]
    const licoUpdateForm = ref(null)
    const lico: Lico = {
      id: 0,
      core_diameter: null,
      pull_plate: LicoPullPlate.Yes,
      section_area: null,
      iron_corner_weight: null,
      stack_width_0: null,
      stack_width_1: null,
      stack_width_2: null,
      stack_width_3: null,
      stack_width_4: null,
      stack_width_5: null,
      stack_width_6: null,
      stack_width_7: null,
      stack_width_8: null,
      stack_thick_0: null,
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
    const tmpLico = ref(lico)
    watch(props, () => {
      tmpLico.value.id = props.lico.id!
      tmpLico.value.core_diameter = props.lico.core_diameter
      tmpLico.value.pull_plate = props.lico.pull_plate
      tmpLico.value.section_area = props.lico.section_area
      tmpLico.value.iron_corner_weight = props.lico.iron_corner_weight
      tmpLico.value.stack_width_0 = props.lico.stack_width_0
      tmpLico.value.stack_width_1 = props.lico.stack_width_1
      tmpLico.value.stack_width_2 = props.lico.stack_width_2
      tmpLico.value.stack_width_3 = props.lico.stack_width_3
      tmpLico.value.stack_width_4 = props.lico.stack_width_4
      tmpLico.value.stack_width_5 = props.lico.stack_width_5
      tmpLico.value.stack_width_6 = props.lico.stack_width_6
      tmpLico.value.stack_width_7 = props.lico.stack_width_7
      tmpLico.value.stack_width_8 = props.lico.stack_width_8
      tmpLico.value.stack_thick_0 = props.lico.stack_thick_0
      tmpLico.value.stack_thick_1 = props.lico.stack_thick_1
      tmpLico.value.stack_thick_2 = props.lico.stack_thick_2
      tmpLico.value.stack_thick_3 = props.lico.stack_thick_3
      tmpLico.value.stack_thick_4 = props.lico.stack_thick_4
      tmpLico.value.stack_thick_5 = props.lico.stack_thick_5
      tmpLico.value.stack_thick_6 = props.lico.stack_thick_6
      tmpLico.value.stack_thick_7 = props.lico.stack_thick_7
      tmpLico.value.stack_thick_8 = props.lico.stack_thick_8
      tmpLico.value.stack_thick_sum = props.lico.stack_thick_sum
      tmpLico.value.main_level_real_stack_thick = props.lico.main_level_real_stack_thick
    })
    const show = computed({
      get: () => props.modelValue,
      set: (v) => context.emit('update:modelValue', v)
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lico.core_diameter == null) {
        return callback(new Error('请输入铁心直径(mm)：'))
      }
      if (lico.core_diameter <= 0) {
        return callback(new Error('铁心直径必须大于0！'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lico.pull_plate == null) {
        return callback(new Error('请选择有无拉板：'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lico.section_area == null) {
        return callback(new Error('请输入铁心截面积(mm)：'))
      }
      if (lico.section_area <= 0) {
        return callback(new Error('铁心截面积必须大于0！'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lico.iron_corner_weight == null) {
        return callback(new Error('请输入角重(kg)：'))
      }
      if (lico.iron_corner_weight <= 0) {
        return callback(new Error('角重必须大于0！'))
      }
      callback()
    }

    const rules = {
      core_diameter: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      pull_plate: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      section_area: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      iron_corner_weight: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ]
    }
    function updateLico() {
      (licoUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlicoLibraries(tmpLico.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lico-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLico() {
    //   const response = await UpdateStdlicoLibraries(tmpLico.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lico-list')
    //     show.value = false
    //   }
    // }
    return {
      pullplateOptions,
      licoUpdateForm,
      show,
      tmpLico,
      rules,
      updateLico
    }
  }
})
</script>
