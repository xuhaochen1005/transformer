<template>
  <el-dialog
    v-model="show"
    title="修改额定频率"
    show-close
    width="30%"
  >
    <el-form
      ref="lrfUpdateForm"
      :model="tmpLrf"
      :rules="rules"
      label-width="150px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定频率（Hz）"
        prop="rated_freq"
        label-width=""
      >
        <el-input-number
          v-model="tmpLrf.rated_freq"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLrf"
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
import { Lrf, UpdateStdLrfLibraries } from '@/api/standard_libraries/lrf'

export default defineComponent({
  name: 'LrfUpdate',
  props: {
    modelValue: Boolean,
    lrf: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lrf-list'],
  setup(props, context) {
    const lrfUpdateForm = ref(null)
    const lrf: Lrf = {
      id: 0,
      rated_freq: 0
    }
    const tmpLrf = ref(lrf)
    watch(props, () => {
      tmpLrf.value.id = props.lrf.id
      tmpLrf.value.rated_freq = props.lrf.rated_freq
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lrfUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor(rule: any, value: any, callback: any) {
      if (!lrf.rated_freq) {
        return callback(new Error('请输入额定频率：'))
      }
      if (lrf.rated_freq < 0) {
        return callback(new Error('额定频率必须大于0！'))
      }
      callback()
    }

    const rules = {
      rated_freq: [
        { required: true, validator: validateFactor, trigger: 'blur' }
      ]

    }
    function updateLrf() {
      (lrfUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLrfLibraries(tmpLrf.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lrf-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLrf() {
    //   const response = await UpdateStdLrfLibraries(tmpLrf.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lrf-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lrfUpdateForm,
      show,
      tmpLrf,
      rules,
      updateLrf
    }
  }
})
</script>
