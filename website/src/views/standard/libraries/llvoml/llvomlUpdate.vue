<template>
  <el-dialog
    v-model="show"
    title="修改低压外模台账"
    show-close
    width="30%"
  >
    <el-form
      ref="llvomlUpdateForm"
      :model="tmpLlvoml"
      :rules="rules"
      label-width="170px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="模具类型"
        prop="mod_type"
        label-width=""
      >
        <el-select
          v-model="tmpLlvoml.mod_type"
          placeholder="模具类型"
          filterable
          allow-create
          clearable
        >
          <el-option
            label="硬膜"
            value="硬膜"
          />
          <el-option
            label="软膜"
            value="软膜"
          />
          <el-option
            label="箔式浇注外模"
            value="箔式浇注外模"
          />
          <el-option
            label="水冷变浇注外模"
            value="水冷变浇注外模"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="模具直径（mm）"
        prop="mod_diameter"
        label-width=""
      >
        <el-input-number
          v-model="tmpLlvoml.mod_diameter"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具高度（mm）"
        prop="mod_height"
        label-width=""
      >
        <el-input-number
          v-model="tmpLlvoml.mod_height"
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
          v-model="tmpLlvoml.mod_amount"
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
          v-model="tmpLlvoml.mod_num"
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
          v-model="tmpLlvoml.mod_suit"
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
          v-model="tmpLlvoml.remark"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLlvoml"
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
import { Llvoml, UpdateStdLlvomlLibraries } from '@/api/standard_libraries/llvoml'

export default defineComponent({
  name: 'LlvomlUpdate',
  props: {
    modelValue: Boolean,
    llvoml: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-llvoml-list'],
  setup(props, context) {
    const llvomlUpdateForm = ref(null)
    const llvoml: Llvoml = {
      id: 0,
      mod_diameter: 0,
      mod_type: '',
      mod_height: 0,
      mod_amount: 0,
      mod_num: '',
      mod_suit: '',
      remark: ''
    }
    const tmpLlvoml = ref(llvoml)
    watch(props, () => {
      tmpLlvoml.value.id = props.llvoml.id
      tmpLlvoml.value.mod_diameter = props.llvoml.mod_diameter
      tmpLlvoml.value.mod_type = props.llvoml.mod_type
      tmpLlvoml.value.mod_height = props.llvoml.mod_height
      tmpLlvoml.value.mod_amount = props.llvoml.mod_amount
      tmpLlvoml.value.mod_num = props.llvoml.mod_num
      tmpLlvoml.value.mod_suit = props.llvoml.mod_suit
      tmpLlvoml.value.remark = props.llvoml.remark
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (llvomlUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (llvoml.mod_diameter == null) {
        return callback(new Error('请输入低压外模外径（mm）：'))
      }
      if (llvoml.mod_diameter < 0) {
        return callback(new Error('低压外模外径不能小于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (!llvoml.mod_type) {
        return callback(new Error('请输入模具类型：'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (llvoml.mod_height == null) {
        return callback(new Error('请输入模具高度（mm）：'))
      }
      if (llvoml.mod_height < 0) {
        return callback(new Error('模具高度不能小于0'))
      }
      callback()
    }

    const rules = {
      mod_diameter: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      mod_type: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      mod_height: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ]
    }
    function updateLlvoml() {
      (llvomlUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLlvomlLibraries(tmpLlvoml.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-llvoml-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLlvoml() {
    //   const response = await UpdateStdLlvomlLibraries(tmpLlvoml.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-llvoml-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      llvomlUpdateForm,
      show,
      tmpLlvoml,
      rules,
      updateLlvoml
    }
  }
})
</script>
