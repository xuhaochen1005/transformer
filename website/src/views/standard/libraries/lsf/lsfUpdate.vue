<template>
  <el-dialog
    v-model="show"
    title="修改迭片系数"
    show-close
    width="30%"
  >
    <el-form
      ref="lsfUpdateForm"
      :model="tmpLsf"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="迭片片厚（mm）"
        prop="stack_thick"
      >
        <el-input-number
          v-model="tmpLsf.stack_thick"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="迭片系数"
        prop="stack_factor"
      >
        <el-input-number
          v-model="tmpLsf.stack_factor"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLsf"
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
import { Lsf, UpdateStdLsfLibraries } from '@/api/standard_libraries/lsf'

export default defineComponent({
  name: 'LsfUpdate',
  props: {
    modelValue: Boolean,
    lsf: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lsf-list'],
  setup(props, context) {
    const lsfUpdateForm = ref(null)
    const lsf: Lsf = {
      id: 0,
      stack_thick: 0,
      stack_factor: 0
    }
    const tmpLsf = ref(lsf)
    watch(props, () => {
      tmpLsf.value.id = props.lsf.id
      tmpLsf.value.stack_thick = props.lsf.stack_thick
      tmpLsf.value.stack_factor = props.lsf.stack_factor
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lsfUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
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
    function updateLsf() {
      (lsfUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLsfLibraries(tmpLsf.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lsf-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLsf() {
    //   const response = await UpdateStdLsfLibraries(tmpLsf.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lsf-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lsfUpdateForm,
      show,
      tmpLsf,
      rules,
      updateLsf
    }
  }
})
</script>
