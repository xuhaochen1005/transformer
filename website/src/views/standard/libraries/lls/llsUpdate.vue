<template>
  <el-dialog
    v-model="show"
    title="修改引线损耗"
    show-close
    width="30%"
  >
    <el-form
      ref="llsUpdateForm"
      :model="tmpLls"
      :rules="rules"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定容量(kVA)"
        prop="wind_type"
        label-width=""
      >
        <el-input-number
          v-model="tmpLls.rated_capacity"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="引线损耗(W)"
        prop="lead_loss"
        label-width=""
      >
        <el-input-number
          v-model="tmpLls.lead_loss"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLls"
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
import { Lls, UpdateStdLlsLibraries } from '@/api/standard_libraries/lls'

export default defineComponent({
  name: 'LlsUpdate',
  props: {
    modelValue: Boolean,
    lls: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lls-list'],
  setup(props, context) {
    const llsUpdateForm = ref(null)
    const lls: Lls = {
      id: 0,
      rated_capacity: 0,
      lead_loss: 0
    }
    const tmpLls = ref(lls)
    watch(props, () => {
      tmpLls.value.id = props.lls.id
      tmpLls.value.rated_capacity = props.lls.rated_capacity
      tmpLls.value.lead_loss = props.lls.lead_loss
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (llsUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lls.rated_capacity == null) {
        return callback(new Error('请输入额定容量（kVA）：'))
      }
      if (lls.rated_capacity <= 0) {
        return callback(new Error('额定容量必须大于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lls.lead_loss == null) {
        return callback(new Error('请输入引线损耗（W）：'))
      }
      if (lls.lead_loss <= 0) {
        return callback(new Error('引线损耗必须大于0'))
      }
      callback()
    }

    const rules = {
      rated_capacity: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      lead_loss: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ]
    }

    function updateLls() {
      (llsUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLlsLibraries(tmpLls.value)
          if (response.data.code === 200) {
            ElMessage.success('更新成功')
            context.emit('get-lls-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLls() {
    //   const response = await UpdateStdLlsLibraries(tmpLls.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lls-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      llsUpdateForm,
      show,
      tmpLls,
      rules,
      updateLls
    }
  }
})
</script>
