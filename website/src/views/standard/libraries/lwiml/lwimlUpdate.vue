<template>
  <el-dialog
    v-model="show"
    title="修改绕线内模台账"
    show-close
    width="30%"
  >
    <el-form
      ref="lwimlUpdateForm"
      :model="tmpLwiml"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="模具直径（mm）"
        prop="mod_diameter"
      >
        <el-input-number
          v-model="tmpLwiml.mod_diameter"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具规格（mm*mm）"
        prop="mod_size"
      >
        <el-input
          v-model="tmpLwiml.mod_size"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具数量（个）"
        prop="mod_amount"
      >
        <el-input-number
          v-model="tmpLwiml.mod_amount"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具图号"
        prop="mod_pic"
      >
        <el-input
          v-model="tmpLwiml.mod_pic"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具编号"
        prop="mod_num"
      >
        <el-input
          v-model="tmpLwiml.mod_num"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="适用产品"
        prop="mod_suit"
      >
        <el-input
          v-model="tmpLwiml.mod_suit"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="备注"
        prop="remark"
      >
        <el-input
          v-model="tmpLwiml.remark"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLwiml"
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
import { Lwiml, UpdateStdLwimlLibraries } from '@/api/standard_libraries/lwiml'

export default defineComponent({
  name: 'LwimlUpdate',
  props: {
    modelValue: Boolean,
    lwiml: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lwiml-list'],
  setup(props, context) {
    const lwimlUpdateForm = ref(null)
    const lwiml: Lwiml = {
      id: 0,
      mod_diameter: 0,
      mod_size: '',
      mod_pic: '',
      mod_amount: 0,
      mod_num: '',
      mod_suit: '',
      remark: ''
    }
    const tmpLwiml = ref(lwiml)
    watch(props, () => {
      tmpLwiml.value.id = props.lwiml.id
      tmpLwiml.value.mod_diameter = props.lwiml.mod_diameter
      tmpLwiml.value.mod_size = props.lwiml.mod_size
      tmpLwiml.value.mod_pic = props.lwiml.mod_pic
      tmpLwiml.value.mod_amount = props.lwiml.mod_amount
      tmpLwiml.value.mod_num = props.lwiml.mod_num
      tmpLwiml.value.mod_suit = props.lwiml.mod_suit
      tmpLwiml.value.remark = props.lwiml.remark
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lwimlUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
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
      if (!lwiml.mod_size) {
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
    function updateLwiml() {
      (lwimlUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLwimlLibraries(tmpLwiml.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lwiml-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLwiml() {
    //   const response = await UpdateStdLwimlLibraries(tmpLwiml.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lwiml-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lwimlUpdateForm,
      show,
      tmpLwiml,
      rules,
      updateLwiml
    }
  }
})
</script>
