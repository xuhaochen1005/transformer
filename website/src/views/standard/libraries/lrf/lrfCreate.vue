<template>
  <el-dialog
    v-model="show"
    title="额定频率"
    show-close
    width="30%"
  >
    <el-form
      ref="lrfCreateForm"
      :model="lrf"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定频率（Hz）"
        prop="rated_freq"
        label-width=""
      >
        <el-input-number
          v-model="lrf.rated_freq"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLrf"
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
import { CreateStdLrfLibraries } from '@/api/standard_libraries/lrf'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LrfCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lrf-list'],
  setup(props, context) {
    const lrfCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          // (lrfCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lrf = reactive({
      rated_freq: 0
    })
    function validateFactor(rule: any, value: any, callback: any) {
      if (lrf.rated_freq == null) {
        return callback(new Error('请输入额定频率：'))
      }
      if (lrf.rated_freq < 0) {
        return callback(new Error('额定频率不能小于0！'))
      }
      callback()
    }

    const rules = {
      rated_freq: [
        { required: true, validator: validateFactor, trigger: 'blur' }
      ]

    }
    function createLrf() {
      (lrfCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLrfLibraries(lrf)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lrf-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLrf() {
    //   const response = await CreateStdLrfLibraries(lrf)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lrf-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lrfCreateForm,
      show,
      lrf,
      rules,
      createLrf
    }
  }
})
</script>
