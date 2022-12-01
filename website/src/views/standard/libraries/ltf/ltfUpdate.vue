<template>
  <el-dialog
    v-model="show"
    title="修改工艺系数"
    show-close
    width="30%"
  >
    <el-form
      ref="ltfUpdateForm"
      :model="tmpLtf"
      :rules="rules"
      label-width="150px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="硅钢片片数（片）"
        prop="stack_amount"
        label-width=""
      >
        <el-input-number
          v-model="tmpLtf.stack_amount"
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
          v-model="tmpLtf.tech_factor"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLtf"
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
import { Ltf, UpdateStdLtfLibraries } from '@/api/standard_libraries/ltf'

export default defineComponent({
  name: 'LtfUpdate',
  props: {
    modelValue: Boolean,
    ltf: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-ltf-list'],
  setup(props, context) {
    const ltfUpdateForm = ref(null)
    const ltf: Ltf = {
      id: 0,
      stack_amount: 0,
      tech_factor: 0
    }
    const tmpLtf = ref(ltf)
    watch(props, () => {
      tmpLtf.value.id = props.ltf.id
      tmpLtf.value.stack_amount = props.ltf.stack_amount
      tmpLtf.value.tech_factor = props.ltf.tech_factor
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (ltfUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
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
    function updateLtf() {
      (ltfUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLtfLibraries(tmpLtf.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-ltf-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLtf() {
    //   const response = await UpdateStdLtfLibraries(tmpLtf.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-ltf-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      ltfUpdateForm,
      show,
      tmpLtf,
      rules,
      updateLtf
    }
  }
})
</script>
