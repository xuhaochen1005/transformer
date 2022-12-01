<template>
  <el-dialog
    v-model="show"
    title="新增引线损耗"
    show-close
    width="30%"
  >
    <el-form
      ref="llsCreateForm"
      :rules="rules"
      :model="lls"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定容量(kVA)"
        prop="rated_capacity"
        label-width=""
      >
        <el-input-number
          v-model="lls.rated_capacity"
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
          v-model="lls.lead_loss"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLls"
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
import { CreateStdLlsLibraries } from '@/api/standard_libraries/lls'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LlsCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lls-list'],
  setup(props, context) {
    const llsCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (llsCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lls = reactive({
      rated_capacity: 0,
      lead_loss: 0
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
    function createLls() {
      (llsCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLlsLibraries(lls)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lls-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLls() {
    //   const response = await CreateStdLlsLibraries(lls)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lls-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      llsCreateForm,
      show,
      lls,
      rules,
      createLls
    }
  }
})
</script>
