<template>
  <el-dialog
    v-model="show"
    title="新增铁心表(长圆形)标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="licoCreateForm"
      :rules="rules"
      :model="lico"
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
          v-model="lico.core_diameter"
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
          v-model="lico.pull_plate"
          style="width: 100%;"
          placeholder="有无拉板"
          clearable
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
          v-model="lico.section_area"
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
          v-model="lico.iron_corner_weight"
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
          v-model="lico.stack_width_0"
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
          v-model="lico.stack_width_1"
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
          v-model="lico.stack_width_1"
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
          v-model="lico.stack_width_2"
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
          v-model="lico.stack_width_3"
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
          v-model="lico.stack_width_4"
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
          v-model="lico.stack_width_5"
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
          v-model="lico.stack_width_6"
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
          v-model="lico.stack_width_7"
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
          v-model="lico.stack_width_8"
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
          v-model="lico.stack_thick_0"
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
          v-model="lico.stack_thick_1"
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
          v-model="lico.stack_thick_2"
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
          v-model="lico.stack_thick_3"
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
          v-model="lico.stack_thick_4"
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
          v-model="lico.stack_thick_5"
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
          v-model="lico.stack_thick_6"
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
          v-model="lico.stack_thick_7"
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
          v-model="lico.stack_thick_8"
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
          v-model="lico.stack_thick_sum"
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
          v-model="lico.main_level_real_stack_thick"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLico"
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
import { CreateStdlicoLibraries, LicoPullPlate } from '@/api/standard_libraries/lico'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LicoCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lico-list'],
  setup(props, context) {
    const pullplateOptions = [
      { label: '有', value: LicoPullPlate.Yes },
      { label: '无', value: LicoPullPlate.No }
    ]
    const licoCreateForm = ref(null)
    const lico = reactive({
      id: 0,
      core_diameter: 0,
      pull_plate: LicoPullPlate.Yes,
      section_area: 0,
      iron_corner_weight: 0,
      stack_width_0: 0,
      stack_width_1: 0,
      stack_width_2: 0,
      stack_width_3: 0,
      stack_width_4: 0,
      stack_width_5: 0,
      stack_width_6: 0,
      stack_width_7: 0,
      stack_width_8: 0,
      stack_thick_0: 0,
      stack_thick_1: 0,
      stack_thick_2: 0,
      stack_thick_3: 0,
      stack_thick_4: 0,
      stack_thick_5: 0,
      stack_thick_6: 0,
      stack_thick_7: 0,
      stack_thick_8: 0,
      stack_thick_sum: 0,
      main_level_real_stack_thick: 0
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (licoCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
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
      if (!lico.pull_plate) {
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
    function createLico() {
      (licoCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlicoLibraries(lico)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lico-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLico() {
    //   const response = await CreateStdlicoLibraries(lico)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lico-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      pullplateOptions,
      licoCreateForm,
      show,
      lico,
      rules,
      createLico
    }
  }
})
</script>
