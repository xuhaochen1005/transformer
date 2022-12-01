<template>
  <el-dialog
    v-model="show"
    title="新增硅钢片标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lstackCreateForm"
      :rules="rules"
      :model="lstack"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="硅钢片型号"
        prop="type"
        label-width=""
      >
        <el-input
          v-model="lstack.type"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="密度(kg/cm^3)"
        prop="density"
        label-width=""
      >
        <el-input-number
          v-model="lstack.density"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="单价(元/kg)"
        prop="price"
        label-width=""
      >
        <el-input-number
          v-model="lstack.price"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLstack"
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
import { CreateStdlstackLibraries } from '@/api/standard_libraries/lstack'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LstackCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lstack-list'],
  setup(props, context) {
    const lstackCreateForm = ref(null)
    const lstack = reactive({
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
          (lstackCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!lstack.type) {
        return callback(new Error('请输入硅钢片型号：'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lstack.density == null) {
        return callback(new Error('请输入密度(kg/cm^3)：'))
      }
      if (lstack.density < 0) {
        return callback(new Error('密度不能小于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lstack.price == null) {
        return callback(new Error('请输入单价(元/kg)：'))
      }
      if (lstack.price < 0) {
        return callback(new Error('单价不能小于0'))
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
    function createLstack() {
      (lstackCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlstackLibraries(lstack)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lstack-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLstack() {
    //   const response = await CreateStdlstackLibraries(lstack)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lstack-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lstackCreateForm,
      show,
      lstack,
      rules,
      createLstack
    }
  }
})
</script>
