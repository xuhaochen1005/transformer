<template>
  <el-dialog
    v-model="show"
    title="修改干式变压器铁心直径经验系数K标准库"
    show-close
    width="30%"
  >
    <el-form
      ref="lckUpdateForm"
      :model="tmpLck"
      :rules="rules"
      label-width="220px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="绕组材质"
        prop="winding_material"
        label-width=""
      >
        <el-select
          v-model="tmpLck.winding_material"
          size="medium"
          style="width: 100%"
          clearable
        >
          <el-option
            v-for="item in materialOptions"
            :key="item"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="额定容量下限(kVA)"
        prop="rated_capacity_min"
        label-width=""
      >
        <el-input-number
          v-model="tmpLck.rated_capacity_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定容量上限(kVA)"
        prop="rated_capacity_max"
        label-width=""
      >
        <el-input-number
          v-model="tmpLck.rated_capacity_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="k值下限"
        prop="k_min"
        label-width=""
      >
        <el-input-number
          v-model="tmpLck.k_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="k值上限"
        prop="k_max"
        label-width=""
      >
        <el-input-number
          v-model="tmpLck.k_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="updateLck"
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
import { Lck, UpdateStdlckLibraries } from '@/api/standard_libraries/lck'

export default defineComponent({
  name: 'LckUpdate',
  props: {
    modelValue: Boolean,
    lck: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue', 'get-lck-list'],
  setup (props, context) {
    const lckUpdateForm = ref(null)
    const lck: Lck = {
      id: 0,
      winding_material: '',
      rated_capacity_min: 0,
      rated_capacity_max: 0,
      k_min: 0,
      k_max: 0
    }
    const tmpLck = ref(lck)
    watch(props, () => {
      tmpLck.value.id = props.lck.id
      tmpLck.value.winding_material = props.lck.winding_material
      tmpLck.value.rated_capacity_min = props.lck.rated_capacity_min
      tmpLck.value.rated_capacity_max = props.lck.rated_capacity_max
      tmpLck.value.k_min = props.lck.k_min
      tmpLck.value.k_max = props.lck.k_max
    })
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (lckUpdateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })

    const materialOptions = [
      { label:'铝', value:'铝'},
      { label:'铜', value:'铜'},
    ]

    function validateFactor_1(rule: any, value: any, callback: any) {
      if (lck.rated_capacity_min == null) {
        return callback(new Error('请输入额定容量下界（kVA）'))
      }
      if (lck.rated_capacity_min < 0) {
        return callback(new Error('额定容量下界必须大于0!'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (lck.rated_capacity_max == null) {
        return callback(new Error('请输入额定容量上界（kVA）'))
      }
      if (lck.rated_capacity_max < 0) {
        return callback(new Error('额定容量上界必须大于0!'))
      }
      callback()
    }

    function validateFactor_3(rule: any, value: any, callback: any) {
      if (lck.k_min == null) {
        return callback(new Error('请输入经验系数k值下界'))
      }
      callback()
    }

    function validateFactor_4(rule: any, value: any, callback: any) {
      if (lck.k_max == null) {
        return callback(new Error('请输入经验系数k值上界'))
      }
      callback()
    }

    function validateFactor_5(rule: any, value: any, callback: any) {
      if (lck.winding_material == null) {
        return callback(new Error('请选择绕组材质'))
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
      k_min: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      k_max: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ],
      winding_material: [
        { required: true, validator: validateFactor_5, trigger: 'blur' }
      ]
    }

    function updateLck() {
      (lckUpdateForm.value as any).validate(async (valid: boolean) => {
        if (valid) {
          const response = await UpdateStdlckLibraries(tmpLck.value)
          if (response.data.code === 200) {
            ElMessage.success('更新成功')
            context.emit('get-lck-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function updateLck() {
    //   const response = await UpdateStdlckLibraries(tmpLck.value)
    //   if (response.data.code === 200) {
    //     ElMessage.success('信息更新成功')
    //     context.emit('get-lck-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    return {
      lckUpdateForm,
      show,
      tmpLck,
      rules,
      materialOptions,
      updateLck
    }
  }
})
</script>
