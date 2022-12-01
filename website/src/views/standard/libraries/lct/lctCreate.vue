<template>
  <el-dialog
    v-model="show"
    title="新增冷却方式标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lctCreateForm"
      :rules="rules"
      :model="lct"
      label-width="120px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="冷却方式"
        prop="cool_type"
        label-width=""
      >
        <el-input
          v-model="lct.cool_type"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLct"
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
import { CreateStdlctLibraries } from '@/api/standard_libraries/lct'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LctCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lct-list'],
  setup(props, context) {
    const lctCreateForm = ref(null)
    const lct = reactive({
      id: 0,
      cool_type: ''
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lctCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor(rule: any, value: any, callback: any) {
      if (!lct.cool_type) {
        return callback(new Error('请输入冷却方式'))
      }
      callback()
    }

    const rules = {
      cool_type: [
        { required: true, validator: validateFactor, trigger: 'blur' }
      ]
    }
    function createLct() {
      (lctCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlctLibraries(lct)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lct-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLct() {
    //   const response = await CreateStdlctLibraries(lct)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lct-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lctCreateForm,
      show,
      lct,
      rules,
      createLct
    }
  }
})
</script>
