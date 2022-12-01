<template>
  <el-dialog
    v-model="show"
    title="新增迭片系数"
    show-close
    width="30%"
  >
    <el-form
      ref="lsfCreateForm"
      :model="lsf"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="迭片片厚（mm）"
        prop="stack_thick"
        label-width=""
      >
        <el-input-number
          v-model="lsf.stack_thick"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="迭片系数"
        prop="stack_factor"
        label-width=""
      >
        <el-input-number
          v-model="lsf.stack_factor"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLsf"
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
import { CreateStdLsfLibraries } from '@/api/standard_libraries/lsf'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LsfCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lsf-list'],
  setup(props, context) {
    const lsfCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lsfCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lsf = reactive({
      stack_thick: 0,
      stack_factor: 0
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lsf.stack_thick == null) {
        return callback(new Error('请输入迭片厚度（mm）：'))
      }
      if (lsf.stack_thick < 0) {
        return callback(new Error('迭片厚度不能小于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lsf.stack_factor == null) {
        return callback(new Error('请输入迭片系数：'))
      }
      if (lsf.stack_factor < 0) {
        return callback(new Error('迭片系数不能小于0'))
      }
      callback()
    }

    const rules = {
      stack_thick: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      stack_factor: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],

    }
    function createLsf() {
      (lsfCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLsfLibraries(lsf)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lsf-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLsf() {
    //   const response = await CreateStdLsfLibraries(lsf)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lsf-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lsfCreateForm,
      show,
      lsf,
      rules,
      createLsf
    }
  }
})
</script>
