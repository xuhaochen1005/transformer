<template>
  <el-dialog
    v-model="show"
    title="新增电阻率标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lresistCreateForm"
      :rules="rules"
      :model="lresist"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="导线材质"
        prop="wire_material"
        label-width=""
      >
        <el-select
          v-model="lresist.wire_material"
          filterable
          allow-create
          size="medium"
          style="width: 100%"
          clearable
        >
          <el-option
            label="铜线"
            value="铜线"
          />
          <el-option
            label="铝线"
            value="铝线"
          />
          <el-option
            label="铜箔"
            value="铜箔"
          />
          <el-option
            label="铝箔"
            value="铝箔"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="温度(℃)"
        prop="temp"
        label-width=""
      >
        <el-input-number
          v-model="lresist.temp"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="电阻率(Ω·m)"
        prop="resistivity"
        label-width=""
      >
        <el-input-number
          v-model="lresist.resistivity"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLresist"
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
import { CreateStdlresistLibraries } from '@/api/standard_libraries/lresist'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LresistCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lresist-list'],
  setup(props, context) {
    const lresistCreateForm = ref(null)
    const lresist = reactive({
      id: 0,
      wire_material: '',
      temp: 0,
      resistivity: 0
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lresistCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!lresist.wire_material) {
        return (new Error('请输入导线材质：'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lresist.temp == null) {
        return (new Error('请输入温度(℃)：'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lresist.resistivity == null) {
        return (new Error('请输入电阻率(Ω·m)：'))
      }
      if (lresist.resistivity < 0) {
        return callback(new Error('电阻率不能小于0！'))
      }
      callback()
    }

    const rules = {
      wire_material: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      temp: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      resistivity: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
    }
    function createLresist() {
      (lresistCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdlresistLibraries(lresist)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lresist-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLresist() {
    //   const response = await CreateStdlresistLibraries(lresist)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lresist-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lresistCreateForm,
      show,
      lresist,
      rules,
      createLresist
    }
  }
})
</script>
