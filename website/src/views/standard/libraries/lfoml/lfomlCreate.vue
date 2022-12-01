<template>
  <el-dialog
    v-model="show"
    title="新增箔绕扁形模具台账标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lfomlCreateForm"
      :rules="rules"
      :model="lfoml"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="模具尺寸"
        prop="mod_size"
        label-width=""
      >
        <el-input
          v-model="lfoml.mod_size"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具高度(mm)"
        prop="mod_height"
        label-width=""
      >
        <el-input-number
          v-model="lfoml.mod_height"
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
          v-model="lfoml.mod_pic"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具编号"
        prop="mod_num"
        label-width=""
      >
        <el-input
          v-model="lfoml.mod_num"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具数量(个)"
        prop="mod_amount"
        label-width=""
      >
        <el-input-number
          v-model="lfoml.mod_amount"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="适用产品型号"
        prop="mod_suit"
        label-width=""
      >
        <el-input
          v-model="lfoml.mod_suit"
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
          v-model="lfoml.remark"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLfoml"
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
import { CreateStdlfomlLibraries } from '@/api/standard_libraries/foml'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LfomlCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lfoml-list'],
  setup(props, context) {
    const lfomlCreateForm = ref(null)
    const lfoml = reactive({
      id: 0,
      mod_size: '',
      mod_height: 0,
      mod_pic: '',
      mod_num: '',
      mod_amount: 0,
      mod_suit: '',
      remark: ''
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lfomlCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lfoml.mod_height == null) {
        return callback(new Error('请输入模具高度（mm）'))
      }
      if (lfoml.mod_height <= 0) {
        return callback(new Error('模具高度必须大于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lfoml.mod_size == null) {
        return callback(new Error('请输入模具尺寸（mm）'))
      }
      callback()
    }

    const rules = {
      mod_height: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      mod_size: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ]
    }
    function createLfoml() {
      (lfomlCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlfomlLibraries(lfoml)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lfoml-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLfoml() {
    //   const response = await CreateStdlfomlLibraries(lfoml)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lfoml-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lfomlCreateForm,
      show,
      lfoml,
      rules,
      createLfoml
    }
  }
})
</script>
