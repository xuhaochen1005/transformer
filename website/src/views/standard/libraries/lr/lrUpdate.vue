<template>
  <el-dialog
    v-model="show"
    title="修改树脂规格表标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lrUpdateForm"
      :model="tmpLr"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="树脂型号"
        prop="type"
        label-width=""
      >
        <el-input
          v-model="tmpLr.type"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="树脂密度(kg/m^3)"
        prop="density"
        label-width=""
      >
        <el-input-number
          v-model="tmpLr.density"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="单价(元/kg)"
        prop="price"
        label-width=""
      >
        <el-input-number
          v-model="tmpLr.price"
          style="width:100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLr"
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
import { Lr, UpdateStdlrLibraries } from '@/api/standard_libraries/lr'

export default defineComponent({
  name: 'LrUpdate',
  props: {
    modelValue: Boolean,
    lr: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lr-list'],
  setup (props, context) {
    const lrUpdateForm = ref(null)
    const lr: Lr = {
      id: 0,
      type: null,
      density: null,
      price: null
    }
    const tmpLr = ref(lr)
    watch(props, () => {
      tmpLr.value.id = props.lr.id!
      tmpLr.value.type = props.lr.type
      tmpLr.value.density = props.lr.density
      tmpLr.value.price = props.lr.price
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lrUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!lr.type) {
        return callback(new Error('请输入树脂型号：'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lr.density == null) {
        return callback(new Error('请输入树脂密度（kg/m^3）：'))
      }
      if (lr.density < 0) {
        return callback(new Error('树脂密度不能小于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lr.price == null) {
        return callback(new Error('请输入树脂单价（元/kg）：'))
      }
      if (lr.price < 0) {
        return callback(new Error('树脂单价不能小于0'))
      }
      callback()
    }

    const rules = {
      type: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      density: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      price: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ]
    }
    function updateLr() {
      (lrUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlrLibraries(tmpLr.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lr-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLr() {
    //   const response = await UpdateStdlrLibraries(tmpLr.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lr-list')
    //     show.value = false
    //   }
    // }
    return {
      lrUpdateForm,
      show,
      tmpLr,
      rules,
      updateLr
    }
  }
})
</script>
