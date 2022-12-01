<template>
  <el-dialog
    v-model="show"
    title="修改风道条台账标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="ladblUpdateForm"
      :model="tmpLadbl"
      :rules="rules"
      label-width="160px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="风道条厚度(mm)"
        prop="air_duct_bar_thickness"
        label-width=""
      >
        <el-input-number
          v-model="tmpLadbl.air_duct_bar_thickness"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="风道条宽度(mm)"
        prop="air_duct_bar_width"
        label-width=""
      >
        <el-input-number
          v-model="tmpLadbl.air_duct_bar_width"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="风道条长度(mm)"
        prop="air_duct_bar_length"
        label-width=""
      >
        <el-input-number
          v-model="tmpLadbl.air_duct_bar_length"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="风道条数量(个)"
        prop="air_duct_bar_num"
        label-width=""
      >
        <el-input-number
          v-model="tmpLadbl.air_duct_bar_num"
          style="width:100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLadbl"
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
import { Ladbl, UpdateStdladblLibraries } from '@/api/standard_libraries/ladbl'

export default defineComponent({
  name: 'LadblUpdate',
  props: {
    modelValue: Boolean,
    ladbl: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-ladbl-list'],
  setup (props, context) {
    const ladblUpdateForm = ref(null)
    const ladbl: Ladbl = {
      id: 0,
      air_duct_bar_thickness: 0,
      air_duct_bar_width: 0,
      air_duct_bar_length: 0,
      air_duct_bar_num: 0,
      ladbl_creator: 0
    }
    const tmpLadbl = ref(ladbl)
    watch(props, () => {
      tmpLadbl.value.id = props.ladbl.id!
      tmpLadbl.value.air_duct_bar_thickness = props.ladbl.air_duct_bar_thickness
      tmpLadbl.value.air_duct_bar_width = props.ladbl.air_duct_bar_width
      tmpLadbl.value.air_duct_bar_length = props.ladbl.air_duct_bar_length
      tmpLadbl.value.air_duct_bar_num = props.ladbl.air_duct_bar_num
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (ladblUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (ladbl.air_duct_bar_thickness === null) {
        return (new Error('请输入风道条厚度（mm）：'))
      }
      if (ladbl.air_duct_bar_thickness < 0) {
        return callback(new Error('风道条厚度必须大于0'))
      }
      callback()
    }
    function validateFactor_2(rule: any, value: any, callback: any) {
      if (ladbl.air_duct_bar_width === null) {
        return (new Error('请输入风道条宽度（mm）：'))
      }
      if (ladbl.air_duct_bar_width < 0) {
        return callback(new Error('风道条宽度必须大于0'))
      }
      callback()
    }
    function validateFactor_3(rule: any, value: any, callback: any) {
      if (ladbl.air_duct_bar_length === null) {
        return (new Error('请输入风道条长度（mm）：'))
      }
      if (ladbl.air_duct_bar_length < 0) {
        return callback(new Error('风道条长度必须大于0'))
      }
      callback()
    }

    const rules = {
      air_duct_bar_thickness: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      air_duct_bar_width: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      air_duct_bar_length: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ]
    }

    function updateLadbl() {
      (ladblUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdladblLibraries(tmpLadbl.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-ladbl-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLadbl() {
    //   const response = await UpdateStdladblLibraries(tmpLadbl.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-ladbl-list')
    //     show.value = false
    //   }
    // }
    return {
      ladblUpdateForm,
      show,
      tmpLadbl,
      rules,
      updateLadbl
    }
  }
})
</script>
