<template>
  <el-dialog
    v-model="show"
    title="新增绕线内模台账"
    show-close
    width="30%"
  >
    <el-form
      ref="lwimlCreateForm"
      :model="lwiml"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="模具直径（mm）"
        prop="mod_diameter"
        label-width=""
      >
        <el-input-number
          v-model="lwiml.mod_diameter"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具规格"
        prop="mod_size"
        label-width=""
      >
        <el-input
          v-model="lwiml.mod_size"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具数量（个）"
        prop="mod_amount"
        label-width=""
      >
        <el-input-number
          v-model="lwiml.mod_amount"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具图号"
        prop="mod_pic"
        label-width=""
      >
        <el-input
          v-model="lwiml.mod_pic"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具型号"
        prop="mod_num"
        label-width=""
      >
        <el-input
          v-model="lwiml.mod_num"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="适用产品"
        prop="mod_suit"
        label-width=""
      >
        <el-input
          v-model="lwiml.mod_suit"
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
          v-model="lwiml.remark"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLwiml"
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
import { CreateStdLwimlLibraries } from '@/api/standard_libraries/lwiml'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LwimlCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lwiml-list'],
  setup(props, context) {
    const lwimlCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lwimlCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lwiml = reactive({
      mod_diameter: 0,
      mod_size: '',
      mod_pic: '',
      mod_amount: 0,
      mod_num: '',
      mod_suit: '',
      remark: ''
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lwiml.mod_diameter == null) {
        return callback(new Error('请输入模具直径（mm）：'))
      }
      if (lwiml.mod_diameter < 0) {
        return callback(new Error('模具直径不能小于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lwiml.mod_size == null) {
        return callback(new Error('请输入模具规格：'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lwiml.mod_amount == null) {
        return callback(new Error('请输入模具数量：'))
      }
      if (lwiml.mod_amount < 0) {
        return callback(new Error('模具数量不能小于0'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (!lwiml.mod_pic) {
        return callback(new Error('请输入模具图号：'))
      }
      callback()
    }

    function validateFactor_5(rule: any, value: any, callback: any) {
      if (!lwiml.mod_num) {
        return callback(new Error('请输入模具编号：'))
      }
      callback()
    }

    const rules = {
      mod_diameter: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      mod_size: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      mod_amount: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      mod_pic: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ],
      mod_num: [
        { required: true, validator: validateFactor_5, trigger: 'blur' }
      ],

    }
    function createLwiml() {
      (lwimlCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLwimlLibraries(lwiml)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lwiml-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLwiml() {
    //   const response = await CreateStdLwimlLibraries(lwiml)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lwiml-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lwimlCreateForm,
      show,
      lwiml,
      rules,
      createLwiml
    }
  }
})
</script>
