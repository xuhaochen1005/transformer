<template>
  <el-dialog
    v-model="show"
    title="新增端绝缘距离"
    show-close
    width="30%"
  >
    <el-form
      ref="ltiCreateForm"
      :model="lti"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定高压下界（kV）"
        prop="rated_high_vol_min"
        label-width=""
      >
        <el-input-number
          v-model="lti.rated_high_vol_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定高压上界（kV）"
        prop="rated_high_vol_max"
        label-width=""
      >
        <el-input-number
          v-model="lti.rated_high_vol_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定低压下界（kV）"
        prop="rated_low_vol_min"
        label-width=""
      >
        <el-input-number
          v-model="lti.rated_low_vol_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定低压上界（kV）"
        prop="rated_low_vol_max"
        label-width=""
      >
        <el-input-number
          v-model="lti.rated_low_vol_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="端绝缘距离（mm）"
        prop="terminal_insulate"
        label-width=""
      >
        <el-input-number
          v-model="lti.terminal_insulate"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLti"
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
import { CreateStdLtiLibraries } from '@/api/standard_libraries/lti'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LtiCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lti-list'],
  setup(props, context) {
    const ltiCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (ltiCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lti = reactive({
      rated_high_vol_min: 0,
      rated_high_vol_max: 0,
      rated_low_vol_min: 0,
      rated_low_vol_max: 0,
      terminal_insulate: 0
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lti.rated_high_vol_min == null) {
        return callback(new Error('请输入额定高压下界（kV）'))
      }
      if (lti.rated_high_vol_min < 0) {
        return callback(new Error('额定高压下界不能小于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lti.rated_high_vol_max == null) {
        return callback(new Error('请输入额定高压上界（kV）'))
      }
      if (lti.rated_high_vol_max < 0) {
        return callback(new Error('额定高压上界不能小于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lti.rated_low_vol_min == null) {
        return callback(new Error('请输入额定低压下界（kV）'))
      }
      if (lti.rated_low_vol_min < 0) {
        return callback(new Error('额定低压下界不能小于0'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lti.rated_low_vol_max == null) {
        return callback(new Error('请输入额定低压上界（kV）'))
      }
      if (lti.rated_low_vol_max < 0) {
        return callback(new Error('额定低压上界不能小于0'))
      }
      callback()
    }

    function validateFactor_5(rule: any, value: any, callback: any) {
      if (lti.terminal_insulate == null) {
        return callback(new Error('请输入端绝缘距离（mm）'))
      }
      if (lti.terminal_insulate < 0) {
        return callback(new Error('端绝缘距离不能小于0'))
      }
      callback()
    }

    const rules = {
      rated_high_vol_min: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      rated_high_vol_max: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      rated_low_vol_min: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      rated_low_vol_max: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ],
      terminal_insulate: [
        { required: true, validator: validateFactor_5, trigger: 'blur' }
      ],

    }
    function createLti() {
      (ltiCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLtiLibraries(lti)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lti-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLti() {
    //   const response = await CreateStdLtiLibraries(lti)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lti-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      ltiCreateForm,
      show,
      lti,
      rules,
      createLti
    }
  }
})
</script>
