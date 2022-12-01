<template>
  <el-dialog
    v-model="show"
    title="修改调压方式"
    show-close
    width="30%"
  >
    <el-form
      ref="lrwUpdateForm"
      :model="tmpLrw"
      :rules="rules"
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
          v-model="tmpLrw.regulate_way"
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
          v-model="tmpLrw.regulate_way_sign"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLrw"
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
import { Lrw, UpdateStdLrwLibraries } from '@/api/standard_libraries/lrw'

export default defineComponent({
  name: 'LrwUpdate',
  props: {
    modelValue: Boolean,
    lrw: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lrw-list'],
  setup (props, context) {
    const lrwUpdateForm = ref(null)
    const lrw: Lrw = {
      id: 0,
      regulate_way: null,
      regulate_way_sign: null
    }
    const tmpLrw = ref(lrw)
    watch(props, () => {
      tmpLrw.value.id = props.lrw.id!
      tmpLrw.value.regulate_way = props.lrw.regulate_way
      tmpLrw.value.regulate_way_sign = props.lrw.regulate_way_sign
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lrwUpdateForm.value as any).clearValidate()
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
    function updateLrw() {
      (lrwUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLrwLibraries(tmpLrw.value)
          if (response.data.code === 200) {
            ElMessage.success('信息更新成功')
            context.emit('get-lrw-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLrw() {
    //   const response = await UpdateStdLrwLibraries(tmpLrw.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lrw-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lrwUpdateForm,
      show,
      tmpLrw,
      rules,
      updateLrw
    }
  }
})
</script>
