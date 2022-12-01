<template>
  <el-dialog
    v-model="show"
    title="修改变压器用途"
    show-close
    width="30%"
  >
    <el-form
      ref="lusUpdateForm"
      :model="tmpLus"
      :rules="rules"
      label-width="250px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="变压器用途"
        prop="usage"
        label-width=""
      >
        <el-input
          v-model="tmpLus.usage"
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
          v-model="tmpLus.usage_sign"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLus"
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
import { Lus, UpdateStdLusLibraries } from '@/api/standard_libraries/lus'

export default defineComponent({
  name: 'LusUpdate',
  props: {
    modelValue: Boolean,
    lus: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lus-list'],
  setup(props, context) {
    const lusUpdateForm = ref(null)
    const lus: Lus = {
      id: 0,
      usage: '',
      usage_sign: ''
    }
    const tmpLus = ref(lus)
    watch(props, () => {
      tmpLus.value.id = props.lus.id
      tmpLus.value.usage = props.lus.usage
      tmpLus.value.usage_sign = props.lus.usage_sign
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lusUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
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
    function updateLus() {
      (lusUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLusLibraries(tmpLus.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lus-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLus() {
    //   const response = await UpdateStdLusLibraries(tmpLus.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lus-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lusUpdateForm,
      show,
      tmpLus,
      rules,
      updateLus
    }
  }
})
</script>
