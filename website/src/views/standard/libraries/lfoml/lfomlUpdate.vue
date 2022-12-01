<template>
  <el-dialog
    v-model="show"
    title="修改箔绕扁形模具台账标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lfomlUpdateForm"
      :model="tmpLfoml"
      :rules="rules"
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
          v-model="tmpLfoml.mod_size"
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
          v-model="tmpLfoml.mod_height"
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
          v-model="tmpLfoml.mod_pic"
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
          v-model="tmpLfoml.mod_num"
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
          v-model="tmpLfoml.mod_amount"
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
          v-model="tmpLfoml.mod_suit"
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
          v-model="tmpLfoml.remark"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLfoml"
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
import { computed, defineComponent, reactive, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Lfoml, UpdateStdlfomlLibraries } from '@/api/standard_libraries/foml'

export default defineComponent({
  name: 'LfomlUpdate',
  props: {
    modelValue: Boolean,
    lfoml: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lfoml-list'],
  setup (props, context) {
    const lfomlUpdateForm = ref(null)
    const lfoml: Lfoml = {
      id: 0,
      mod_size: '',
      mod_height: 0,
      mod_pic: '',
      mod_num: '',
      mod_amount: 0,
      mod_suit: '',
      remark: ''
    }
    const tmpLfoml = ref(lfoml)
    watch(props, () => {
      tmpLfoml.value.id = props.lfoml.id!
      tmpLfoml.value.mod_size = props.lfoml.mod_size
      tmpLfoml.value.mod_height = props.lfoml.mod_height
      tmpLfoml.value.mod_pic = props.lfoml.mod_pic
      tmpLfoml.value.mod_num = props.lfoml.mod_num
      tmpLfoml.value.mod_amount = props.lfoml.mod_amount
      tmpLfoml.value.mod_suit = props.lfoml.mod_suit
      tmpLfoml.value.remark = props.lfoml.remark
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lfomlUpdateForm.value as any).clearValidate()
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

    function updateLfoml() {
      (lfomlUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlfomlLibraries(tmpLfoml.value)
          if (response.data.code === 200) {
            ElMessage.success('更新成功')
            context.emit('get-lfoml-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLfoml() {
    //   const response = await UpdateStdlfomlLibraries(tmpLfoml.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lfoml-list')
    //     show.value = false
    //   }
    // }
    return {
      lfomlUpdateForm,
      show,
      tmpLfoml,
      rules,
      updateLfoml
    }
  }
})
</script>
