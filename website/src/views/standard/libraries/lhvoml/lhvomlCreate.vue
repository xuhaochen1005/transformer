<template>
  <el-dialog
    v-model="show"
    title="新增高压外模台账标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lhvomlCreateForm"
      :model="lhvoml"
      :rules="rules"
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="模具尺寸"
        prop="mod_size"
        label-width=""
      >
        <el-input
          v-model="lhvoml.mod_size"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具类型"
        prop="mod_type"
        label-width=""
      >
        <el-input
          v-model="lhvoml.mod_type"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具高度(mm)"
        prop="mod_height"
        label-width=""
      >
        <el-input-number
          v-model="lhvoml.mod_height"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具数量(个)"
        prop="mod_amount"
        label-width=""
      >
        <el-input-number
          v-model="lhvoml.mod_amount"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具图号"
        prop="mod_pic"
        label-width=""
      >
        <el-input
          v-model="lhvoml.mod_pic"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="模具编号"
        prop="mod_num"
        label-width=""
      >
        <el-input
          v-model="lhvoml.mod_num"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="适用产品"
        prop="mod_suit"
        label-width=""
      >
        <el-input
          v-model="lhvoml.mod_suit"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="凸台高度(封板尺寸)"
        prop="boss_width"
        label-width=""
      >
        <el-input
          v-model="lhvoml.boss_width"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="是否开槽"
        prop="groove"
        label-width=""
      >
        <el-select
          v-model="lhvoml.groove"
          size="medium"
          style="width: 100%"
          clearable
        >
          <el-option
            label="有"
            value="有"
          />
          <el-option
            label="否"
            value="否"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="螺母尺寸(mm)"
        prop="nut_size"
        label-width=""
      >
        <el-input-number
          v-model="lhvoml.nut_size"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="A头尺寸(mm)"
        prop="a_size"
        label-width=""
      >
        <el-input-number
          v-model="lhvoml.a_size"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="抽头孔距(mm)"
        prop="tap_hole_distance"
        label-width=""
      >
        <el-input-number
          v-model="lhvoml.tap_hole_distance"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="封板图号"
        prop="closure_pic"
        label-width=""
      >
        <el-input
          v-model="lhvoml.closure_pic"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="备注"
        prop="remark"
        label-width=""
      >
        <el-input
          v-model="lhvoml.remark"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLhvoml"
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
import { CreateStdLhvomlLibraries } from '@/api/standard_libraries/lhvoml'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'LhvomlCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-lhvoml-list'],
  setup(props, context) {
    const lhvomlCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lhvomlCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const lhvoml = reactive({
      id: 0,
      mod_size: '',
      mod_type: '',
      mod_height: 0,
      mod_amount: 0,
      mod_pic: '',
      mod_num: '',
      mod_suit: '',
      boss_width: '',
      groove: '',
      nut_size: 0,
      a_size: 0,
      tap_hole_distance: 0,
      closure_pic: '',
      remark: ''
    })
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lhvoml.mod_height == null) {
        return callback(new Error('请输入模具高度(mm)：'))
      }
      if (lhvoml.mod_height < 0) {
        return callback(new Error('模具高度必须大于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lhvoml.mod_size == null) {
        return callback(new Error('请输入模具尺寸(mm)：'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (!lhvoml.mod_type) {
        return callback(new Error('请输入模具类型：'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lhvoml.boss_width == null) {
        return callback(new Error('请输入凸台高度（封板尺寸）(mm)：'))
      }
      callback()
    }

    function validateFactor_5(rule: any, value: any, callback: any) {
      if (lhvoml.nut_size == null) {
        return callback(new Error('请输入螺母尺寸(mm)：'))
      }
      callback()
    }

    function validateFactor_6(rule: any, value: any, callback: any) {
      if (lhvoml.a_size == null) {
        return callback(new Error('请输入A头尺寸(mm)：'))
      }
      callback()
    }

    function validateFactor_7(rule: any, value: any, callback: any) {
      if (lhvoml.tap_hole_distance == null) {
        return callback(new Error('请输入抽头孔距(mm)：'))
      }
      callback()
    }

    const rules = {
      mod_height: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      mod_size: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      mod_type: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      boss_width: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ],
      nut_size: [
        { required: true, validator: validateFactor_5, trigger: 'blur' }
      ],
      a_size: [
        { required: true, validator: validateFactor_6, trigger: 'blur' }
      ],
      tap_hole_distance: [
        { required: true, validator: validateFactor_7, trigger: 'blur' }
      ]
    }
    function createLhvoml() {
      (lhvomlCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLhvomlLibraries(lhvoml)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-lhvoml-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLhvoml() {
    //   const response = await CreateStdLhvomlLibraries(lhvoml)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-lhvoml-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lhvomlCreateForm,
      show,
      lhvoml,
      rules,
      createLhvoml
    }
  }
})
</script>
