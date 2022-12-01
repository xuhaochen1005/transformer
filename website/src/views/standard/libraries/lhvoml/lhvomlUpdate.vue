<template>
  <el-dialog
    v-model="show"
    title="修改高压外模台账标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lhvomlUpdateForm"
      :model="tmpLhvoml"
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
          v-model="tmpLhvoml.mod_size"
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
          v-model="tmpLhvoml.mod_type"
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
          v-model="tmpLhvoml.mod_height"
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
          v-model="tmpLhvoml.mod_amount"
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
          v-model="tmpLhvoml.mod_pic"
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
          v-model="tmpLhvoml.mod_num"
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
          v-model="tmpLhvoml.mod_suit"
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
          v-model="tmpLhvoml.boss_width"
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
        <el-input
          v-model="tmpLhvoml.groove"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="螺母尺寸(mm)"
        prop="nut_size"
        label-width=""
      >
        <el-input-number
          v-model="tmpLhvoml.nut_size"
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
          v-model="tmpLhvoml.a_size"
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
          v-model="tmpLhvoml.tap_hole_distance"
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
          v-model="tmpLhvoml.closure_pic"
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
          v-model="tmpLhvoml.remark"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLhvoml"
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
import { Lhvoml, UpdateStdLhvomlLibraries } from '@/api/standard_libraries/lhvoml'

export default defineComponent({
  name: 'LhvomlUpdate',
  props: {
    modelValue: Boolean,
    lhvoml: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lhvoml-list'],
  setup(props, context) {
    const lhvomlUpdateForm = ref(null)
    const lhvoml: Lhvoml = {
      id: 0,
      mod_size: '',
      mod_type: '',
      mod_height: null,
      mod_amount: null,
      mod_pic: '',
      mod_num: '',
      mod_suit: '',
      boss_width: '',
      groove: '',
      nut_size: null,
      a_size: null,
      tap_hole_distance: null,
      closure_pic: '',
      remark: ''
    }
    const tmpLhvoml = ref(lhvoml)
    watch(props, () => {
      tmpLhvoml.value.id = props.lhvoml.id
      tmpLhvoml.value.mod_size = props.lhvoml.mod_size
      tmpLhvoml.value.mod_type = props.lhvoml.mod_type
      tmpLhvoml.value.mod_height = props.lhvoml.mod_height
      tmpLhvoml.value.mod_amount = props.lhvoml.mod_amount
      tmpLhvoml.value.mod_pic = props.lhvoml.mod_pic
      tmpLhvoml.value.mod_num = props.lhvoml.mod_num
      tmpLhvoml.value.mod_suit = props.lhvoml.mod_suit
      tmpLhvoml.value.boss_width = props.lhvoml.boss_width
      tmpLhvoml.value.groove = props.lhvoml.groove
      tmpLhvoml.value.nut_size = props.lhvoml.nut_size
      tmpLhvoml.value.a_size = props.lhvoml.a_size
      tmpLhvoml.value.tap_hole_distance = props.lhvoml.tap_hole_distance
      tmpLhvoml.value.closure_pic = props.lhvoml.closure_pic
      tmpLhvoml.value.remark = props.lhvoml.remark
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lhvomlUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lhvoml.mod_height == null) {
        return callback(new Error('请输入模具高度(mm)：'))
      }
      if (lhvoml.mod_height <= 0) {
        return callback(new Error('模具高度必须大于0'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (!lhvoml.mod_size) {
        return callback(new Error('请输入模具尺寸(mm)：'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lhvoml.mod_type == null) {
        return callback(new Error('请输入模具类型：'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (!lhvoml.boss_width) {
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
    function updateLhvoml() {
      (lhvomlUpdateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await UpdateStdLhvomlLibraries(lhvoml)
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
    // async function updateLhvoml() {
    //   const response = await UpdateStdLhvomlLibraries(tmpLhvoml.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lhvoml-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lhvomlUpdateForm,
      show,
      tmpLhvoml,
      rules,
      updateLhvoml
    }
  }
})
</script>
