<template>
  <el-dialog
    v-model="show"
    title="修改箔绕低压模具台账标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lflvmlUpdateForm"
      :model="tmpLflvml"
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
          v-model="tmpLflvml.mod_size"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="平台高"
        prop="platform_height"
        label-width=""
      >
        <el-input-number
          v-model="tmpLflvml.platform_height"
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
          v-model="tmpLflvml.mod_height"
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
          v-model="tmpLflvml.mod_pic"
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
          v-model="tmpLflvml.mod_num"
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
          v-model="tmpLflvml.mod_amount"
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
          v-model="tmpLflvml.mod_suit"
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
          v-model="tmpLflvml.remark"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLflvml"
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
import { Lflvml, UpdateStdlflvmlLibraries } from '@/api/standard_libraries/lflvml'

export default defineComponent({
  name: 'LflvmlUpdate',
  props: {
    modelValue: Boolean,
    lflvml: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lflvml-list'],
  setup (props, context) {
    const lflvmlUpdateForm = ref(null)
    const lflvml: Lflvml = {
      id: 0,
      mod_size: '',
      platform_height: null,
      mod_height: null,
      mod_pic: '',
      mod_num: '',
      mod_amount: null,
      mod_suit: '',
      remark: ''
    }
    const tmpLflvml = ref(lflvml)
    watch(props, () => {
      tmpLflvml.value.id = props.lflvml.id!
      tmpLflvml.value.mod_size = props.lflvml.mod_size
      tmpLflvml.value.platform_height = props.lflvml.platform_height
      tmpLflvml.value.mod_height = props.lflvml.mod_height
      tmpLflvml.value.mod_pic = props.lflvml.mod_pic
      tmpLflvml.value.mod_num = props.lflvml.mod_num
      tmpLflvml.value.mod_amount = props.lflvml.mod_amount
      tmpLflvml.value.mod_suit = props.lflvml.mod_suit
      tmpLflvml.value.remark = props.lflvml.remark
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lflvmlUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lflvml.mod_height == null) {
        return (new Error('请输入模具高度（mm）：'))
      }
      if (lflvml.mod_height <= 0) {
        return callback(new Error('模具高度必须大于0！'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lflvml.platform_height == null) {
        return (new Error('请输入平台高（mm）：'))
      }
      if (lflvml.platform_height < 0) {
        return callback(new Error('平台高必须大于0！'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lflvml.mod_size == null) {
        return (new Error('请输入模具尺寸（mm）：'))
      }
      callback()
    }

    const rules = {
      mod_height: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      platform_height: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      mod_size: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ]
    }

    function updateLflvml() {
      (lflvmlUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlflvmlLibraries(tmpLflvml.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lflvml-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLflvml() {
    //   const response = await UpdateStdlflvmlLibraries(tmpLflvml.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lflvml-list')
    //     show.value = false
    //   }
    // }
    return {
      lflvmlUpdateForm,
      show,
      tmpLflvml,
      rules,
      updateLflvml
    }
  }
})
</script>
