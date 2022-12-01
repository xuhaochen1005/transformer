<template>
  <el-dialog
    v-model="show"
    title="新增导线表标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lwireCreateForm"
      :rules="rules"
      :model="lwire"
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
          v-model="lwire.wire_type"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="导线材质"
        prop="wire_material"
        label-width=""
      >
        <el-select
          v-model="lwire.wire_material"
          filterable
          allow-create
          size="medium"
          style="width: 100%"
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
          v-model="lwire.wire_shape"
          filterable
          allow-create
          size="medium"
          style="width: 100%"
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
        prop="wire_material"
        label-width=""
      >
        <el-input-number
          v-model="lwire.wire_density"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLwire"
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
import { CreateStdlwireLibraries } from '@/api/standard_libraries/lwire'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LwireCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lwire-list'],
  setup(props, context) {
    const lwireCreateForm = ref(null)
    const lwire = reactive({
      id: 0,
      wire_type: '',
      wire_material: '',
      wire_shape: '',
      wire_density: 0
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lwireCreateForm.value as any).clearValidate()
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
      ],
    }
    function createLwire() {
      (lwireCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlwireLibraries(lwire)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lwire-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLwire() {
    //   const response = await CreateStdlwireLibraries(lwire)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lwire-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lwireCreateForm,
      show,
      lwire,
      rules,
      createLwire
    }
  }
})
</script>
