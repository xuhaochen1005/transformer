<template>
  <el-dialog
    v-model="show"
    title="修改导线表标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lwireUpdateForm"
      :model="tmpLwire"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="导线型号"
        prop="wire_type"
        label-width=""
      >
        <el-input
          v-model="tmpLwire.wire_type"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="导线材质"
        prop="wire_material"
        label-width=""
      >
        <el-select
          v-model="tmpLwire.wire_material"
          filterable
          allow-create
          style="width:100%"
          clearable
        >
          <el-option
            label="铜线"
            value="铜线"
          />
          <el-option
            label="铜箔"
            value="铜箔"
          />
          <el-option
            label="铝线"
            value="铝线"
          />
          <el-option
            label="铝箔"
            value="铝箔"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="导线形状"
        prop="wire_shape"
        label-width=""
      >
        <el-select
          v-model="tmpLwire.wire_shape"
          filterable
          allow-create
          style="width:100%"
          clearable
        >
          <el-option
            label="圆"
            value="圆"
          />
          <el-option
            label="扁"
            value="扁"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="导线密度(kg/m^3)"
        prop="wire_density"
        label-width=""
      >
        <el-input-number
          v-model="tmpLwire.wire_density"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="价格(元/m)"
        prop="wire_price"
        label-width=""
      >
        <el-input-number
          v-model="tmpLwire.wire_price"
          :precision="2"
          :step="0.1"
          :min="0"
          style="width:100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLwire"
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
import { Lwire, UpdateStdlwireLibraries } from '@/api/standard_libraries/lwire'

export default defineComponent({
  name: 'LwireUpdate',
  props: {
    modelValue: Boolean,
    lwire: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lwire-list'],
  setup (props, context) {
    const lwireUpdateForm = ref(null)
    const lwire: Lwire = {
      id: 0,
      wire_type: null,
      wire_material: null,
      wire_shape: null,
      wire_density: null
    }
    const tmpLwire = ref(lwire)
    watch(props, () => {
      tmpLwire.value.id = props.lwire.id!
      tmpLwire.value.wire_type = props.lwire.wire_type
      tmpLwire.value.wire_material = props.lwire.wire_material
      tmpLwire.value.wire_shape = props.lwire.wire_shape
      tmpLwire.value.wire_density = props.lwire.wire_density
      tmpLwire.value.wire_price = props.lwire.wire_price
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lwireUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!lwire.wire_type) {
        return callback(new Error('请输入导线型号：'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (!lwire.wire_material) {
        return callback(new Error('请输入导线材质：'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (!lwire.wire_shape) {
        return callback(new Error('请输入导线形状：'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lwire.wire_density == null) {
        return callback(new Error('请输入导线密度(kg/m^3)：'))
      }
      if (lwire.wire_density < 0) {
        return callback(new Error('导线密度不能小于0'))
      }
      callback()
    }

    const rules = {
      wire_type: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      wire_material: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      wire_shape: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      wire_density: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ]
    }
    function updateLwire() {
      (lwireUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlwireLibraries(tmpLwire.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lwire-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLwire() {
    //   const response = await UpdateStdlwireLibraries(tmpLwire.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lwire-list')
    //     show.value = false
    //   }
    // }
    return {
      lwireUpdateForm,
      show,
      tmpLwire,
      rules,
      updateLwire
    }
  }
})
</script>
