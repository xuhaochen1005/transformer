<template>
  <el-dialog
    v-model="show"
    title="新增圆铜（铝）线规格"
    show-close
    width="30%"
  >
    <el-form
      ref="lrwsCreateForm"
      :model="lrws"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="标称直径（mm）"
        prop="std_diameter"
        label-width=""
      >
        <el-input-number
          v-model="lrws.std_diameter"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="截面积（mm^2）"
        prop="section_area"
        label-width=""
      >
        <el-input-number
          v-model="lrws.section_area"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="最大外径（mm）"
        prop="max_diameter"
        label-width=""
      >
        <el-input-number
          v-model="lrws.max_diameter"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLrws"
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
import { CreateStdLrwsLibraries } from '@/api/standard_libraries/lrws'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LrwsCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lrws-list'],
  setup(props, context) {
    const lrwsCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lrwsCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lrws = reactive({
      std_diameter: 0,
      section_area: 0,
      max_diameter: 0
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lrws.std_diameter == null) {
        return callback(new Error('请输入标称直径（mm）：'))
      }
      if (lrws.std_diameter < 0) {
        return callback(new Error('标称直径不能小于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lrws.section_area == null) {
        return callback(new Error('请输入截面积（mm^2）：'))
      }
      if (lrws.section_area < 0) {
        return callback(new Error('截面积不能小于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lrws.max_diameter == null) {
        return callback(new Error('请输入最大外径（mm）：'))
      }
      if (lrws.max_diameter < 0) {
        return callback(new Error('最大外径不能小于0'))
      }
      callback()
    }

    const rules = {
      std_diameter: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      section_area: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      max_diameter: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ]

    }
    function createLrws() {
      (lrwsCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLrwsLibraries(lrws)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lrws-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLrws() {
    //   const response = await CreateStdLrwsLibraries(lrws)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lrws-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lrwsCreateForm,
      show,
      lrws,
      rules,
      createLrws
    }
  }
})
</script>
