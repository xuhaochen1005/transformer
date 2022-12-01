<template>
  <el-dialog
    v-model="show"
    title="修改线圈接法标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lccUpdateForm"
      :model="tmpLcc"
      :rules="rules"
      label-width="150px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="低压线圈接法"
        prop="low_vol_coil_connect"
        label-width=""
      >
        <el-input
          v-model="tmpLcc.low_vol_coil_connect"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="高压线圈接法"
        prop="high_vol_coil_connect"
        label-width=""
      >
        <el-input
          v-model="tmpLcc.high_vol_coil_connect"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLcc"
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
import { Lcc, UpdateStdlccLibraries } from '@/api/standard_libraries/lcc'

export default defineComponent({
  name: 'LccUpdate',
  props: {
    modelValue: Boolean,
    lcc: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lcc-list'],
  setup: function (props, context) {
    const lccUpdateForm = ref(null)
    const lcc: Lcc = {
      id: 0,
      low_vol_coil_connect: '',
      high_vol_coil_connect: ''
    }
    const tmpLcc = ref(lcc)
    watch(props, () => {
      tmpLcc.value.id = props.lcc.id!
      tmpLcc.value.low_vol_coil_connect = props.lcc.low_vol_coil_connect
      tmpLcc.value.high_vol_coil_connect = props.lcc.high_vol_coil_connect
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lccUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!lcc.low_vol_coil_connect) {
        return callback(new Error('请输入低压线圈连接方式'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (!lcc.high_vol_coil_connect == null) {
        return callback(new Error('请输入高压线圈连接方式'))
      }
      callback()
    }

    const rules = {
      low_vol_coil_connect: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      high_vol_coil_connect: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ]
    }

    function updateLcc() {
      (lccUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlccLibraries(tmpLcc.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lcc-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }

    // async function updateLcc() {
    //   const response = await UpdateStdlccLibraries(tmpLcc.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lcc-list')
    //     show.value = false
    //   }
    // }
    return {
      lccUpdateForm,
      show,
      tmpLcc,
      rules,
      updateLcc
    }
  }
})
</script>
