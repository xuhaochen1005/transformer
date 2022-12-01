<template>
  <el-dialog
    v-model="show"
    title="修改冷却方式标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lctUpdateForm"
      :model="tmpLct"
      :rules="rules"
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
          v-model="tmpLct.cool_type"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLct"
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
import { computed, defineComponent, reactive, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Lct, UpdateStdlctLibraries } from '@/api/standard_libraries/lct'

export default defineComponent({
  name: 'LctUpdate',
  props: {
    modelValue: Boolean,
    lct: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lct-list'],
  setup (props, context) {
    const lctUpdateForm = ref(null)
    const lct: Lct = {
      id: 0,
      cool_type: ''
    }
    const tmpLct = ref(lct)
    watch(props, () => {
      tmpLct.value.id = props.lct.id!
      tmpLct.value.cool_type = props.lct.cool_type
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lctUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor(rule: any, value: any, callback: any) {
      if (!lct.cool_type) {
        return (new Error('请输入冷却类型：'))
      }
      callback()
    }

    const rules = {
      cool_type: [
        { required: true, validator: validateFactor, trigger: 'blur' }
      ]
    }

    function updateLct() {
      (lctUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlctLibraries(tmpLct.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lct-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLct() {
    //   const response = await UpdateStdlctLibraries(tmpLct.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lct-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lctUpdateForm,
      show,
      tmpLct,
      rules,
      updateLct
    }
  }
})
</script>
