<template>
  <el-dialog
    v-model="show"
    title="新增涡流损耗系数标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lelfCreateForm"
      :rules="rules"
      :model="lelf"
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
          v-model="lelf.rated_capacity"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="涡流损耗系数(%)"
        prop="eddy_loss_factor"
        label-width=""
      >
        <el-input-number
          v-model="lelf.eddy_loss_factor"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLelf"
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
import { CreateStdlelfLibraries } from '@/api/standard_libraries/lelf'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LelfCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lelf-list'],
  setup(props, context) {
    const lelfCreateForm = ref(null)
    const lelf = reactive({
      id: 0,
      rated_capacity: 0,
      eddy_loss_factor: 0
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lelfCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lelf.rated_capacity == null) {
        return (new Error('请输入额定容量（kVA）：'))
      }
      if (lelf.rated_capacity <= 0) {
        return callback(new Error('额定容量必须大于0！'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lelf.eddy_loss_factor == null) {
        return (new Error('请输入涡流损耗系数（%）：'))
      }
      callback()
    }

    const rules = {
      rated_capacity: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      eddy_loss_factor: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ]
    }
    function createLelf() {
      (lelfCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlelfLibraries(lelf)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lelf-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLelf() {
    //   const response = await CreateStdlelfLibraries(lelf)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lelf-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lelfCreateForm,
      show,
      lelf,
      rules,
      createLelf
    }
  }
})
</script>
