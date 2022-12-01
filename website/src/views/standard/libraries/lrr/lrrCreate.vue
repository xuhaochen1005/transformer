<template>
  <el-dialog
    v-model="show"
    title="新增铁心磁密Bm初选值"
    show-close
    width="30%"
  >
    <el-form
      ref="lrrCreateForm"
      :model="lrr"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="调压范围下界（%）"
        prop="regulate_range_min"
        label-width=""
      >
        <el-input-number
          v-model="lrr.regulate_range_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="调压范围上界（%）"
        prop="regulate_range_max"
        label-width=""
      >
        <el-input-number
          v-model="lrr.regulate_range_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="调压范围步长（%）"
        prop="regulate_range_step"
        label-width=""
      >
        <el-input-number
          v-model="lrr.regulate_range_step"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="铁心磁密上界（T）"
        prop="mod_amount"
        label-width=""
      >
        <el-input-number
          v-model="lrr.magnet_density_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLrr"
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
import { CreateStdLrrLibraries } from '@/api/standard_libraries/lrr'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LrrCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lrr-list'],
  setup(props, context) {
    const lrrCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lrrCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lrr = reactive({
      regulate_range_min: 0,
      regulate_range_max: 0,
      regulate_range_step: 0
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lrr.regulate_range_min == null) {
        return callback(new Error('请输入调压范围下界（%）：'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lrr.regulate_range_max == null) {
        return callback(new Error('请输入调压范围上界（%）：'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lrr.regulate_range_step == null) {
        return callback(new Error('请输入调压范围步长（%）：'))
      }
      callback()
    }

    const rules = {
      regulate_range_min: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      regulate_range_max: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      regulate_range_step: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ]

    }
    function createLrr() {
      (lrrCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLrrLibraries(lrr)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lrr-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLrr() {
    //   const response = await CreateStdLrrLibraries(lrr)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lrr-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lrrCreateForm,
      show,
      lrr,
      rules,
      createLrr
    }
  }
})
</script>
