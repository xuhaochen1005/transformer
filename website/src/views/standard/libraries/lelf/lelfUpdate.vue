<template>
  <el-dialog
    v-model="show"
    title="修改涡流损耗系数标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lelfUpdateForm"
      :model="tmpLelf"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定容量(kVA)"
        prop="rated_capacity"
        label-width=""
      >
        <el-input-number
          v-model="tmpLelf.rated_capacity"
          style="width:100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="涡流损耗系数(%)"
        prop="eddy_loss_factor"
        label-width=""
      >
        <el-input-number
          v-model="tmpLelf.eddy_loss_factor"
          style="width:100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLelf"
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
import { Lelf, UpdateStdlelfLibraries } from '@/api/standard_libraries/lelf'

export default defineComponent({
  name: 'LelfUpdate',
  props: {
    modelValue: Boolean,
    lelf: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lelf-list'],
  setup (props, context) {
    const lelfUpdateForm = ref(null)
    const lelf: Lelf = {
      id: 0,
      rated_capacity: null,
      eddy_loss_factor: null
    }
    const tmpLelf = ref(lelf)
    watch(props, () => {
      tmpLelf.value.id = props.lelf.id!
      tmpLelf.value.rated_capacity = props.lelf.rated_capacity
      tmpLelf.value.eddy_loss_factor = props.lelf.eddy_loss_factor
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lelfUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lelf.rated_capacity === null) {
        return (new Error('请输入额定容量（kVA）：'))
      }
      if (lelf.rated_capacity <= 0) {
        return callback(new Error('额定容量必须大于0！'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lelf.eddy_loss_factor === null) {
        return (new Error('请输入涡流损耗系数（%）：'))
      }
      callback()
    }

    const rules = {
      rated_capacity: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      eddy_loss_factor: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ]
    }
    function updateLelf() {
      (lelfUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlelfLibraries(tmpLelf.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lelf-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLelf() {
    //   const response = await UpdateStdlelfLibraries(tmpLelf.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lelf-list')
    //     show.value = false
    //   }
    // }
    return {
      lelfUpdateForm,
      show,
      tmpLelf,
      rules,
      updateLelf
    }
  }
})
</script>
