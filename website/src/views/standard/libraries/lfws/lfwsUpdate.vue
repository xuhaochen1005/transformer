<template>
  <el-dialog
    v-model="show"
    title="修改扁铜(铝)线规格标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lfwsUpdateForm"
      :model="tmpLfws"
      :rules="rules"
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
          v-model="tmpLfws.std_length"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="标称宽(mm)"
        prop="std_width"
        label-width=""
      >
        <el-input-number
          v-model="tmpLfws.std_width"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="截面积(mm^2)"
        prop="section_area"
        label-width=""
      >
        <el-input-number
          v-model="tmpLfws.section_area"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="圆角半径(mm)"
        prop="round_corner"
        label-width=""
      >
        <el-input-number
          v-model="tmpLfws.round_corner"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="备注"
        prop="remark"
        label-width=""
      >
        <el-input
          v-model="tmpLfws.remark"
          style="width:100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLfws"
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
import {computed, defineComponent, PropType, reactive, ref, watch} from 'vue'
import { ElMessage } from 'element-plus'
import { Lfws, UpdateStdlfwsLibraries } from '@/api/standard_libraries/lfws'

export default defineComponent({
  name: 'LfwsUpdate',
  props: {
    modelValue: Boolean,
    lfws: {
      type: Object as PropType<Lfws>,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lfws-list'],
  setup (props, context) {
    const lfwsUpdateForm = ref(null)
    const lfws: Lfws = {
      id: 0,
      std_length: 0,
      std_width: 0,
      section_area: 0,
      round_corner: 0,
      remark: ''
    }
    const tmpLfws = ref(lfws)
    watch(() => props.lfws, () => {
      tmpLfws.value.id = props.lfws.id!
      tmpLfws.value.std_length = props.lfws.std_length
      tmpLfws.value.std_width = props.lfws.std_width
      tmpLfws.value.section_area = props.lfws.section_area
      tmpLfws.value.round_corner = props.lfws.round_corner
      tmpLfws.value.remark = props.lfws.remark
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lfwsUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lfws.std_length == null) {
        return callback(new Error('请输入标称直径（mm）：'))
      }
      if (lfws.std_length < 0) {
        return callback(new Error('标称直径不能小于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lfws.std_width == null) {
        return callback(new Error('请输入标称宽（mm）：'))
      }
      if (lfws.std_width < 0) {
        return callback(new Error('标称直径不能小于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lfws.section_area == null) {
        return callback(new Error('请输入截面积（mm）：'))
      }
      if (lfws.section_area < 0) {
        return callback(new Error('截面积不能小于0'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lfws.round_corner === null) {
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
    function updateLfws() {
      (lfwsUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlfwsLibraries(tmpLfws.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lfws-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLfws() {
    //   const response = await UpdateStdlfwsLibraries(tmpLfws.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lfws-list')
    //     show.value = false
    //   }
    // }
    return {
      lfwsUpdateForm,
      show,
      tmpLfws,
      rules,
      updateLfws
    }
  }
})
</script>
