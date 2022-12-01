<template>
  <el-dialog
    v-model="show"
    title="修改铁心磁密Bm初选值"
    show-close
    width="30%"
  >
    <el-form
      ref="lmdUpdateForm"
      :model="tmpLmd"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定容量下界（kVA）"
        prop="rated_capacity_min"
      >
        <el-input-number
          v-model="tmpLmd.rated_capacity_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定容量上界（kVA）"
        prop="rated_capacity_max"
      >
        <el-input-number
          v-model="tmpLmd.rated_capacity_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="磁通密度下界（T）"
        prop="magnet_density_min"
      >
        <el-input-number
          v-model="tmpLmd.magnet_density_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="磁通密度上界（T）"
        prop="magnet_density_max"
      >
        <el-input-number
          v-model="tmpLmd.magnet_density_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLmd"
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
import { Lmd, UpdateStdLmdLibraries } from '@/api/standard_libraries/lmd'

export default defineComponent({
  name: 'LmdUpdate',
  props: {
    modelValue: Boolean,
    lmd: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lmd-list'],
  setup(props, context) {
    const lmdUpdateForm = ref(null)
    const lmd: Lmd = {
      id: 0,
      rated_capacity_min: 0,
      rated_capacity_max: 0,
      magnet_density_min: 0,
      magnet_density_max: 0
    }
    const tmpLmd = ref(lmd)
    watch(props, () => {
      tmpLmd.value.id = props.lmd.id
      tmpLmd.value.rated_capacity_min = props.lmd.rated_capacity_min
      tmpLmd.value.rated_capacity_max = props.lmd.rated_capacity_max
      tmpLmd.value.magnet_density_min = props.lmd.magnet_density_min
      tmpLmd.value.magnet_density_max = props.lmd.magnet_density_max
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lmdUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lmd.rated_capacity_min == null) {
        return callback(new Error('请输入额定高压下界'))
      }
      if (lmd.rated_capacity_min < 0) {
        return callback(new Error('额定高压下界不能小于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lmd.rated_capacity_max == null) {
        return callback(new Error('请输入额定高压上界'))
      }
      if (lmd.rated_capacity_max < 0) {
        return callback(new Error('额定高压上界不能小于0'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lmd.magnet_density_min == null) {
        return callback(new Error('请输入铁心磁密下界'))
      }
      if (lmd.magnet_density_min < 0) {
        return callback(new Error('铁心磁密下界不能小于0'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lmd.magnet_density_max == null) {
        return callback(new Error('请输入铁心磁密上界'))
      }
      if (lmd.magnet_density_max < 0) {
        return callback(new Error('铁心磁密上界不能小于0'))
      }
      callback()
    }

    const rules = {
      rated_capacity_min: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      rated_capacity_max: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      magnet_density_min: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      magnet_density_max: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ]

    }
    function updateLmd() {
      (lmdUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLmdLibraries(tmpLmd.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lmd-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLmd() {
    //   const response = await UpdateStdLmdLibraries(tmpLmd.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lmd-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lmdUpdateForm,
      show,
      tmpLmd,
      rules,
      updateLmd
    }
  }
})
</script>
