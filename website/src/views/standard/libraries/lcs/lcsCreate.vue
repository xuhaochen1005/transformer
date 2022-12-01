<template>
  <el-dialog
    v-model="show"
    title="新增线圈形状标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lcsCreateForm"
      :rules="rules"
      :model="lcs"
      label-width="100px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="线圈形状"
        prop="coil_shape"
        label-width=""
      >
        <el-input
          v-model="lcs.coil_shape"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLcs"
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
import { CreateStdlcsLibraries } from '@/api/standard_libraries/lcs'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LcsCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lcs-list'],
  setup(props, context) {
    const lcsCreateForm = ref(null)
    const lcs = reactive({
      id: 0,
      coil_shape: ''
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lcsCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor(rule: any, value: any, callback: any) {
      if (!lcs.coil_shape) {
        return callback(new Error('请输入线圈形状'))
      }
      callback()
    }

    const rules = {
      coil_shape: [
        { required: true, validator: validateFactor, trigger: 'blur' }
      ]
    }
    function createLcs() {
      (lcsCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlcsLibraries(lcs)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lcs-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLcs() {
    //   const response = await CreateStdlcsLibraries(lcs)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lcs-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lcsCreateForm,
      show,
      lcs,
      rules,
      createLcs
    }
  }
})
</script>
