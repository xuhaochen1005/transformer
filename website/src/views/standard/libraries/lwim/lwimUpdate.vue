<template>
  <el-dialog
    v-model="show"
    title="修改绕组外绝缘介质"
    show-close
    width="30%"
  >
    <el-form
      ref="lwimUpdateForm"
      :model="tmpLwim"
      :rules="rules"
      label-width="250px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="绕组外绝缘介质"
        prop="wind_insulate_media"
        label-width=""
      >
        <el-input
          v-model="tmpLwim.wind_insulate_media"
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
          v-model="tmpLwim.wind_insulate_media_sign"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLwim"
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
import { computed, defineComponent, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Lwim, UpdateStdLwimLibraries } from '@/api/standard_libraries/lwim'

export default defineComponent({
  name: 'LwimUpdate',
  props: {
    modelValue: Boolean,
    lwim: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lwim-list'],
  setup(props, context) {
    const lwimUpdateForm = ref(null)
    const lwim: Lwim = {
      id: 0,
      wind_insulate_media: '',
      wind_insulate_media_sign: ''
    }
    const tmpLwim = ref(lwim)
    watch(props, () => {
      tmpLwim.value.id = props.lwim.id
      tmpLwim.value.wind_insulate_media = props.lwim.wind_insulate_media
      tmpLwim.value.wind_insulate_media_sign = props.lwim.wind_insulate_media_sign
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lwimUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor(rule: any, value: any, callback: any) {
      if (!lwim.wind_insulate_media) {
        return (new Error('请输入绕组外绝缘介质类型：'))
      }
      callback()
    }

    const rules = {
      wind_insulate_media: [
        { required: true, validator: validateFactor, trigger: 'blur' }
      ]

    }
    function updateLwim() {
      (lwimUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLwimLibraries(tmpLwim.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lwim-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLwim() {
    //   const response = await UpdateStdLwimLibraries(tmpLwim.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lwim-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lwimUpdateForm,
      show,
      tmpLwim,
      rules,
      updateLwim
    }
  }
})
</script>
