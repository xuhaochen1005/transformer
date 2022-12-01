<template>
  <el-dialog
    v-model="show"
    title="修改圆铜（铝）线规格"
    show-close
    width="30%"
  >
    <el-form
      ref="lrwsUpdateForm"
      :model="tmpLrws"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="标称直径（mm）"
        prop="std_diameter"
      >
        <el-input-number
          v-model="tmpLrws.std_diameter"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="截面积（mm^2）"
        prop="section_area"
      >
        <el-input-number
          v-model="tmpLrws.section_area"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="最大外径（mm）"
        prop="max_diameter"
      >
        <el-input-number
          v-model="tmpLrws.max_diameter"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLrws"
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
import { Lrws, UpdateStdLrwsLibraries } from '@/api/standard_libraries/lrws'

export default defineComponent({
  name: 'LrwsUpdate',
  props: {
    modelValue: Boolean,
    lrws: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lrws-list'],
  setup(props, context) {
    const lrwsUpdateForm = ref(null)
    const lrws: Lrws = {
      id: 0,
      std_diameter: 0,
      section_area: 0,
      max_diameter: 0
    }
    const tmpLrws = ref(lrws)
    watch(props, () => {
      tmpLrws.value.id = props.lrws.id
      tmpLrws.value.std_diameter = props.lrws.std_diameter
      tmpLrws.value.section_area = props.lrws.section_area
      tmpLrws.value.max_diameter = props.lrws.max_diameter
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lrwsUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
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
    function updateLrws() {
      (lrwsUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLrwsLibraries(tmpLrws.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lrws-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLrws() {
    //   const response = await UpdateStdLrwsLibraries(tmpLrws.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lrws-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lrwsUpdateForm,
      show,
      tmpLrws,
      rules,
      updateLrws
    }
  }
})
</script>
