<template>
  <el-dialog
    v-model="show"
    title="新增树脂规格表标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lrCreateForm"
      :rules="rules"
      :model="lr"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="树脂型号"
        prop="type"
        label-width=""
      >
        <el-input
          v-model="lr.type"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="树脂密度(kg/m^3)"
        prop="density"
        label-width=""
      >
        <el-input-number
          v-model="lr.density"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="单价"
        prop="price"
        label-width=""
      >
        <el-input-number
          v-model="lr.price"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLr"
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
import { CreateStdlrLibraries } from '@/api/standard_libraries/lr'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LrCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lr-list'],
  setup(props, context) {
    const lrCreateForm = ref(null)
    const lr = reactive({
      id: 0,
      type: '',
      density: 0,
      price: 0
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lrCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!lr.type) {
        return callback(new Error('请输入树脂型号：'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lr.density == null) {
        return callback(new Error('请输入树脂密度（kg/m^3）：'))
      }
      if (lr.density < 0) {
        return callback(new Error('树脂密度不能小于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lr.price == null) {
        return callback(new Error('请输入树脂单价（元/kg）：'))
      }
      if (lr.price < 0) {
        return callback(new Error('树脂单价不能小于0'))
      }
      callback()
    }

    const rules = {
      type: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      density: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      price: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
    }
    function createLr() {
      (lrCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlrLibraries(lr)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lr-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLr() {
    //   const response = await CreateStdlrLibraries(lr)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lr-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lrCreateForm,
      show,
      lr,
      rules,
      createLr
    }
  }
})
</script>
