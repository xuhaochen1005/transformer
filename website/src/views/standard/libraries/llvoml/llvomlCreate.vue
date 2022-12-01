<template>
  <el-dialog
    v-model="show"
    title="新增低压外模台账"
    show-close
    width="30%"
  >
    <el-form
      ref="llvomlCreateForm"
      :model="llvoml"
      :rules="rules"
      label-width="170px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="模具类型"
        prop="mod_type"
        label-width=""
      >
        <el-input
          v-model="llvoml.mod_type"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="低压外模外径（mm）"
        prop="mod_diameter"
        label-width=""
      >
        <el-input-number
          v-model="llvoml.mod_diameter"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具高度（mm）"
        prop="mod_height"
        label-width=""
      >
        <el-input-number
          v-model="llvoml.mod_height"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具数量（个）"
        prop="mod_amount"
        label-width=""
      >
        <el-input-number
          v-model="llvoml.mod_amount"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具型号"
        prop="mod_num"
        label-width=""
      >
        <el-input
          v-model="llvoml.mod_num"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="适用产品"
        prop="mod_suit"
        label-width=""
      >
        <el-input
          v-model="llvoml.mod_suit"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="备注"
        prop="remark"
        label-width=""
      >
        <el-input
          v-model="llvoml.remark"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLlvoml"
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
import { CreateStdLlvomlLibraries } from '@/api/standard_libraries/llvoml'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LlvomlCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-llvoml-list'],
  setup(props, context) {
    const llvomlCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (llvomlCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const llvoml = reactive({
      mod_diameter: 0,
      mod_type: '',
      mod_height: 0,
      mod_amount: 0,
      mod_num: '',
      mod_suit: '',
      remark: ''
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (llvoml.mod_diameter == null) {
        return callback(new Error('请输入低压外模外径（mm）：'))
      }
      if (llvoml.mod_diameter < 0) {
        return callback(new Error('低压外模外径不能小于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (!llvoml.mod_type) {
        return callback(new Error('请输入模具类型：'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (llvoml.mod_height == null) {
        return callback(new Error('请输入模具高度（mm）：'))
      }
      if (llvoml.mod_height < 0) {
        return callback(new Error('模具高度不能小于0'))
      }
      callback()
    }

    const rules = {
      mod_diameter: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      mod_type: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      mod_height: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ]
    }
    function createLlvoml() {
      (llvomlCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLlvomlLibraries(llvoml)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-llvoml-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLlvoml() {
    //   const response = await CreateStdLlvomlLibraries(llvoml)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-llvoml-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      llvomlCreateForm,
      show,
      llvoml,
      rules,
      createLlvoml
    }
  }
})
</script>
