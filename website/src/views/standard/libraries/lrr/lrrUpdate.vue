<template>
  <el-dialog
    v-model="show"
    title="修改调压范围"
    show-close
    width="30%"
  >
    <el-form
      ref="lrrUpdateForm"
      :model="tmpLrr"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="调压范围下界（%）"
        prop="regulate_range_min"
      >
        <el-input-number
          v-model="tmpLrr.regulate_range_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="调压范围上界（%）"
        prop="regulate_range_max"
      >
        <el-input-number
          v-model="tmpLrr.regulate_range_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="调压范围步长（%）"
        prop="regulate_range_step"
      >
        <el-input-number
          v-model="tmpLrr.regulate_range_step"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLrr"
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
import { Lrr, UpdateStdLrrLibraries } from '@/api/standard_libraries/lrr'

export default defineComponent({
  name: 'LrrUpdate',
  props: {
    modelValue: Boolean,
    lrr: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lrr-list'],
  setup(props, context) {
    const lrrUpdateForm = ref(null)
    const lrr: Lrr = {
      id: 0,
      regulate_range_min: 0,
      regulate_range_max: 0,
      regulate_range_step: 0
    }
    const tmpLrr = ref(lrr)
    watch(props, () => {
      tmpLrr.value.id = props.lrr.id
      tmpLrr.value.regulate_range_min = props.lrr.regulate_range_min
      tmpLrr.value.regulate_range_max = props.lrr.regulate_range_max
      tmpLrr.value.regulate_range_step = props.lrr.regulate_range_step
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lrrUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
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
    function updateLrr() {
      (lrrUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLrrLibraries(tmpLrr.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lrr-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLrr() {
    //   const response = await UpdateStdLrrLibraries(tmpLrr.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lrr-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lrrUpdateForm,
      show,
      tmpLrr,
      rules,
      updateLrr
    }
  }
})
</script>
