<template>
  <el-dialog
    v-model="show"
    title="修改线圈形状标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lcsUpdateForm"
      :model="tmpLcs"
      :rules="rules"
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
          v-model="tmpLcs.coil_shape"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLcs"
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
import { Lcs, UpdateStdlcsLibraries } from '@/api/standard_libraries/lcs'

export default defineComponent({
  name: 'LcsUpdate',
  props: {
    modelValue: Boolean,
    lcs: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lcs-list'],
  setup (props, context) {
    const lcsUpdateForm = ref(null)
    const lcs: Lcs = {
      id: 0,
      coil_shape: ''
    }
    const tmpLcs = ref(lcs)
    watch(props, () => {
      tmpLcs.value.id = props.lcs.id!
      tmpLcs.value.coil_shape = props.lcs.coil_shape
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lcsUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor(rule: any, value: any, callback: any) {
      if (!lcs.coil_shape) {
        return (new Error('请输入线圈形状：'))
      }
      callback()
    }

    const rules = {
      coil_shape: [
        { required: true, validator: validateFactor, trigger: 'blur' }
      ]
    }

    function updateLcs() {
      (lcsUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlcsLibraries(tmpLcs.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lcs-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLcs() {
    //   const response = await UpdateStdlcsLibraries(tmpLcs.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lcs-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lcsUpdateForm,
      show,
      tmpLcs,
      rules,
      updateLcs
    }
  }
})
</script>
