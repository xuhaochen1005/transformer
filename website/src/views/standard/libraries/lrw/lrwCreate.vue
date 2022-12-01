<template>
  <el-dialog
    v-model="show"
    title="新增调压方式"
    show-close
    width="30%"
  >
    <el-form
      ref="lrwCreateForm"
      :rules="rules"
      :model="lrw"
      label-width="220px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="调压方式"
        prop="regulate_way"
        label-width=""
      >
        <el-input
          v-model="lrw.regulate_way"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="调压方式代号"
        prop="regulate_way_sign"
        label-width=""
      >
        <el-input
          v-model="lrw.regulate_way_sign"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLrw"
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
import { CreateStdLrwLibraries } from '@/api/standard_libraries/lrw'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LrwCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lrw-list'],
  setup(props, context) {
    const lrwCreateForm = ref(null)
    const lrw = reactive({
      id: 0,
      regulate_way: '',
      regulate_way_sign: ''
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lrwCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!lrw.regulate_way) {
        return callback(new Error('请输入调压方式：'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (!lrw.regulate_way) {
        return callback(new Error('请输入调压方式代号：'))
      }
      callback()
    }

    const rules = {
      regulate_way: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      regulate_way_sign: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ]
    }
    function createLrw() {
      (lrwCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLrwLibraries(lrw)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lrw-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLrw() {
    //   const response = await CreateStdLrwLibraries(lrw)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lrw-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lrwCreateForm,
      show,
      lrw,
      rules,
      createLrw
    }
  }
})
</script>
