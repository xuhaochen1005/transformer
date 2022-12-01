<template>
  <el-dialog
    v-model="show"
    title="新增工艺系数"
    show-close
    width="30%"
  >
    <el-form
      ref="ltfCreateForm"
      :model="ltf"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="硅钢片片数（片）"
        prop="stack_amount"
        label-width=""
      >
        <el-input-number
          v-model="ltf.stack_amount"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="工艺系数"
        prop="tech_factor"
        label-width=""
      >
        <el-input-number
          v-model="ltf.tech_factor"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLtf"
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
import { CreateStdLtfLibraries } from '@/api/standard_libraries/ltf'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LtfCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-ltf-list'],
  setup(props, context) {
    const ltfCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (ltfCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const ltf = reactive({
      stack_amount: 0,
      tech_factor: 0
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (ltf.stack_amount == null) {
        return callback(new Error('请输入硅钢片片数：'))
      }
      if (ltf.stack_amount < 0) {
        return callback(new Error('硅钢片片数不能小于0！'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (ltf.tech_factor == null) {
        return callback(new Error('请输入迭片系数：'))
      }
      if (ltf.tech_factor < 0) {
        return callback(new Error('迭片系数不能小于0！'))
      }
      callback()
    }

    const rules = {
      stack_amount: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      tech_factor: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ]

    }
    function createLtf() {
      (ltfCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLtfLibraries(ltf)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-ltf-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLtf() {
    //   const response = await CreateStdLtfLibraries(ltf)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-ltf-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      ltfCreateForm,
      show,
      ltf,
      rules,
      createLtf
    }
  }
})
</script>
