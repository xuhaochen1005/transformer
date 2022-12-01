<template>
  <el-dialog
    v-model="show"
    title="修改硅钢片标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lstackUpdateForm"
      :model="tmpLstack"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="硅钢片型号"
        prop="type"
        label-width=""
      >
        <el-input
          v-model="tmpLstack.type"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="密度(kg/cm^3)"
        prop="density"
        label-width=""
      >
        <el-input-number
          v-model="tmpLstack.density"
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
          v-model="tmpLstack.price"
          style="width:100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLstack"
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
import { Lstack, UpdateStdlstackLibraries } from '@/api/standard_libraries/lstack'

export default defineComponent({
  name: 'LstackUpdate',
  props: {
    modelValue: Boolean,
    lstack: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lstack-list'],
  setup (props, context) {
    const lstackUpdateForm = ref(null)
    const lstack: Lstack = {
      id: 0,
      type: null,
      density: null,
      price: null
    }
    const tmpLstack = ref(lstack)
    watch(props, () => {
      tmpLstack.value.id = props.lstack.id!
      tmpLstack.value.type = props.lstack.type
      tmpLstack.value.density = props.lstack.density
      tmpLstack.value.price = props.lstack.price
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lstackUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!lstack.type) {
        return callback(new Error('请输入硅钢片型号：'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lstack.density == null) {
        return callback(new Error('请输入密度(kg/cm^3)：'))
      }
      if (lstack.density < 0) {
        return callback(new Error('密度不能小于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lstack.price == null) {
        return callback(new Error('请输入单价(元/kg)：'))
      }
      if (lstack.price < 0) {
        return callback(new Error('单价不能小于0'))
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
      ],

    }
    function updateLstack() {
      (lstackUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlstackLibraries(tmpLstack.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lstack-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLstack() {
    //   const response = await UpdateStdlstackLibraries(tmpLstack.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lstack-list')
    //     show.value = false
    //   }
    // }
    return {
      lstackUpdateForm,
      show,
      tmpLstack,
      rules,
      updateLstack
    }
  }
})
</script>
