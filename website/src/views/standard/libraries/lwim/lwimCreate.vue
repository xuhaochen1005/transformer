<template>
  <el-dialog
    v-model="show"
    title="新增绕组外绝缘介质"
    show-close
    width="30%"
  >
    <el-form
      ref="lwimCreateForm"
      :model="lwim"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="绕组外绝缘介质"
        prop="wind_insulate_media"
        label-width=""
      >
        <el-input
          v-model="lwim.wind_insulate_media"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="绕组外绝缘介质代号"
        prop="wind_insulate_media_sign"
        label-width=""
      >
        <el-input
          v-model="lwim.wind_insulate_media_sign"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLwim"
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
import { CreateStdLwimLibraries } from '@/api/standard_libraries/lwim'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LwimCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lwim-list'],
  setup(props, context) {
    const lwimCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lwimCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lwim = reactive({
      wind_insulate_media: '',
      wind_insulate_media_sign: ''
    })
    function validateFactor(rule: any, value: any, callback: any) {
      if (!lwim.wind_insulate_media) {
        return callback(new Error('请输入绕组外绝缘介质：'))
      }
      callback()
    }

    const rules = {
      wind_insulate_media: [
        { required: true, validator: validateFactor, trigger: 'blur' }
      ]

    }
    function createLwim() {
      (lwimCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLwimLibraries(lwim)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lwim-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLwim() {
    //   const response = await CreateStdLwimLibraries(lwim)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lwim-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lwimCreateForm,
      show,
      lwim,
      rules,
      createLwim
    }
  }
})
</script>
