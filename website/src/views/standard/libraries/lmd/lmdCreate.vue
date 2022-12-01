<template>
  <el-dialog
    v-model="show"
    title="新增铁心磁密Bm初选值"
    show-close
    width="30%"
  >
    <el-form
      ref="lmdCreateForm"
      :model="lmd"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定容量下界（kVA）"
        prop="rated_capacity_min"
        label-width=""
      >
        <el-input-number
          v-model="lmd.rated_capacity_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定容量上界（kVA）"
        prop="rated_capacity_max"
        label-width=""
      >
        <el-input-number
          v-model="lmd.rated_capacity_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="铁心磁密下界（T）"
        prop="magnet_density_min"
        label-width=""
      >
        <el-input-number
          v-model="lmd.magnet_density_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="铁心磁密上界（T）"
        prop="magnet_density_max"
        label-width=""
      >
        <el-input-number
          v-model="lmd.magnet_density_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLmd"
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
import { CreateStdLmdLibraries } from '@/api/standard_libraries/lmd'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LmdCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lmd-list'],
  setup(props, context) {
    const lmdCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lmdCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lmd = reactive({
      rated_capacity_min: 0,
      rated_capacity_max: 0,
      magnet_density_min: 0,
      magnet_density_max: 0
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
    function createLmd() {
      (lmdCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLmdLibraries(lmd)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lmd-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLmd() {
    //   const response = await CreateStdLmdLibraries(lmd)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lmd-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lmdCreateForm,
      show,
      lmd,
      rules,
      createLmd
    }
  }
})
</script>
