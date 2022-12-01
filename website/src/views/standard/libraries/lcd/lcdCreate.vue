<template>
  <el-dialog
    v-model="show"
    title="新增电流密度标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lcdCreateForm"
      :rules="rules"
      :model="lcd"
      label-width="200px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="导线材料"
        prop="winding_material"
        label-width=""
      >
        <el-select
          v-model="lcd.winding_material"
          filterable
          allow-create
          size="medium"
          style="width: 100%"
          placeholder=""
          clearable
        >
          <el-option
            v-for="item in materialOptions"
            :key="item"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="电流密度下界(A/mm^2)"
        prop="current_density_min"
        label-width=""
      >
        <el-input-number
          v-model="lcd.current_density_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="电流密度上界(A/mm^2)"
        prop="current_density_max"
        label-width=""
      >
        <el-input-number
          v-model="lcd.current_density_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLcd"
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
import { CreateStdlcdLibraries } from '@/api/standard_libraries/lcd'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LcdCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lcd-list'],
  setup(props, context) {
    const lcdCreateForm = ref(null)
    const lcd = reactive({
      id: 0,
      winding_material: '',
      current_density_min: 0,
      current_density_max: 0
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lcdCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    const materialOptions = [
      { label:'铝', value:'铝'},
      { label:'铜', value:'铜'},
    ]

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lcd.current_density_min == null) {
        return (new Error('请输入电磁密度下界：'))
      }
      if (lcd.current_density_min < 0) {
        return callback(new Error('电磁密度不能小于0！'))
      }
      callback()
    }
    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lcd.current_density_max == null) {
        return (new Error('请输入电磁密度上界：'))
      }
      if (lcd.current_density_max < 0) {
        return callback(new Error('电磁密度不能小于0！'))
      }
      callback()
    }
    function validateFactor_3(rule: any, value: any, callback: any) {
      if (!lcd.winding_material) {
        return (new Error('请选择导线材料：'))
      }
      callback()
    }

    const rules = {
      current_density_min: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      current_density_max: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      winding_material: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ]
    }
    function createLcd() {
      (lcdCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlcdLibraries(lcd)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lcd-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLcd() {
    //   const response = await CreateStdlcdLibraries(lcd)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lcd-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lcdCreateForm,
      show,
      lcd,
      rules,
      materialOptions,
      createLcd
    }
  }
})
</script>
