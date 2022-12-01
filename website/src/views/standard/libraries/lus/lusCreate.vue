<template>
  <el-dialog
    v-model="show"
    title="新增变压器用途"
    show-close
    width="30%"
  >
    <el-form
      ref="lusCreateForm"
      :model="lus"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="变压器用途"
        prop="usage"
        label-width=""
      >
        <el-input
          v-model="lus.usage"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="变压器用途代号"
        prop="usage_sign"
        label-width=""
      >
        <el-input
          v-model="lus.usage_sign"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLus"
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
import { CreateStdLusLibraries } from '@/api/standard_libraries/lus'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LusCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lus-list'],
  setup(props, context) {
    const lusCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lusCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lus = reactive({
      usage: '',
      usage_sign: ''
    })
    function validateFactor(rule: any, value: any, callback: any) {
      if (!lus.usage) {
        return callback(new Error('请输入用途：'))
      }
      callback()
    }

    const rules = {
      usage: [
        { required: true, validator: validateFactor, trigger: 'blur' }
      ]

    }
    function createLus() {
      (lusCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLusLibraries(lus)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lus-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLus() {
    //   const response = await CreateStdLusLibraries(lus)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lus-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lusCreateForm,
      show,
      lus,
      rules,
      createLus
    }
  }
})
</script>
