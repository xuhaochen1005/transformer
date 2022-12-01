<template>
  <el-dialog
    v-model="show"
    title="修改电流密度标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lcdUpdateForm"
      :model="tmpLcd"
      :rules="rules"
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
          v-model="tmpLcd.winding_material"
          style="width:100%"
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
        label-width="200px"
      >
        <el-input-number
          v-model="tmpLcd.current_density_min"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="电流密度上界(A/mm^2)"
        prop="current_density_max"
        label-width=""
      >
        <el-input-number
          v-model="tmpLcd.current_density_max"
          style="width:100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLcd"
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
import { Lcd, UpdateStdlcdLibraries } from '@/api/standard_libraries/lcd'

export default defineComponent({
  name: 'LcdUpdate',
  props: {
    modelValue: Boolean,
    lcd: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lcd-list'],
  setup (props, context) {
    const lcdUpdateForm = ref(null)
    const lcd: Lcd = {
      id: 0,
      winding_material: '',
      current_density_min: null,
      current_density_max: null
    }
    const tmpLcd = ref(lcd)
    watch(props, () => {
      tmpLcd.value.id = props.lcd.id!
      tmpLcd.value.winding_material = props.lcd.winding_material
      tmpLcd.value.current_density_min = props.lcd.current_density_min
      tmpLcd.value.current_density_max = props.lcd.current_density_max
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lcdUpdateForm.value as any).clearValidate()
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
      if (lcd.current_density_min <= 0) {
        return callback(new Error('电磁密度必须大于0！'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lcd.current_density_max == null) {
        return (new Error('请输入电磁密度上界：'))
      }
      if (lcd.current_density_max <= 0) {
        return callback(new Error('电磁密度必须大于0！'))
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

    function updateLcd() {
      (lcdUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlcdLibraries(tmpLcd.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lcd-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLcd() {
    //   const response = await UpdateStdlcdLibraries(tmpLcd.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lcd-list')
    //     show.value = false
    //   }
    // }
    return {
      lcdUpdateForm,
      show,
      tmpLcd,
      rules,
      materialOptions,
      updateLcd
    }
  }
})
</script>
