<template>
  <el-dialog
    v-model="show"
    title="新增扁铜(铝)线规格标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lfwsCreateForm"
      :rules="rules"
      :model="lfws"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="标称长(mm)"
        prop="std_length"
        label-width=""
      >
        <el-input-number
          v-model="lfws.std_length"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="标称宽(mm)"
        prop="std_width"
        label-width=""
      >
        <el-input-number
          v-model="lfws.std_width"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="截面积(mm^2)"
        prop="section_area"
        label-width=""
      >
        <el-input-number
          v-model="lfws.section_area"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="圆角半径(mm)"
        prop="round_corner"
        label-width=""
      >
        <el-input-number
          v-model="lfws.round_corner"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="备注"
        prop="remark"
        label-width=""
      >
        <el-input
          v-model="lfws.remark"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLfws"
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
import { CreateStdlfwsLibraries } from '@/api/standard_libraries/lfws'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LfwsCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lfws-list'],
  setup(props, context) {
    const lfwsCreateForm = ref(null)
    const lfws = reactive({
      id: 0,
      std_length: 0,
      std_width: 0,
      section_area: 0,
      round_corner: 0,
      remark: ''
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lfwsCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lfws.std_length == null) {
        return callback(new Error('请输入标称直径（mm）：'))
      }
      if (lfws.std_length <= 0) {
        return callback(new Error('标称直径必须大于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lfws.std_width == null) {
        return callback(new Error('请输入标称宽（mm）：'))
      }
      if (lfws.std_width <= 0) {
        return callback(new Error('标称直径必须大于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lfws.section_area == null) {
        return callback(new Error('请输入截面积（mm）：'))
      }
      if (lfws.section_area <= 0) {
        return callback(new Error('截面积必须大于0'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lfws.round_corner == null) {
        return callback(new Error('请输入圆角半径（mm）：'))
      }
      callback()
    }

    const rules = {
      std_length: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      std_width: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      section_area: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      round_corner: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ]
    }
    function createLfws() {
      (lfwsCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlfwsLibraries(lfws)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lfws-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLfws() {
    //   const response = await CreateStdlfwsLibraries(lfws)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lfws-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lfwsCreateForm,
      show,
      lfws,
      rules,
      createLfws
    }
  }
})
</script>
