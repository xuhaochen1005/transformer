<template>
  <el-dialog
    v-model="show"
    title="新增空载损耗，负载损耗，空载电流，短路阻抗"
    show-close
    width="30%"
  >
    <el-form
      ref="llCreateForm"
      :model="ll"
      :rules="rules"
      label-width="200px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="损耗水平代号"
        prop="regulate_way"
        label-width=""
      >
        <el-input-number
          v-model="ll.loss_level"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="调压方式"
        prop="regulate_way"
        label-width=""
      >
        <el-select
          v-model="ll.regulate_way"
          placeholder="调压方式"
        >
          <el-option
            label="无励磁调压"
            value="无励磁调压"
          />
          <el-option
            label="有载调压"
            value="有载调压"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="额定容量下界（kVA）"
        prop="rated_capacity_min"
        label-width=""
      >
        <el-input-number
          v-model="ll.rated_capacity_min"
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
          v-model="ll.rated_capacity_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定高压下界（kV）"
        prop="rated_high_vol_min"
        label-width=""
      >
        <el-input-number
          v-model="ll.rated_high_vol_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定高压上界（kV）"
        prop="rated_high_vol_max"
        label-width=""
      >
        <el-input-number
          v-model="ll.rated_high_vol_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定低压下界（kV）"
        prop="rated_low_vol_min"
        label-width=""
      >
        <el-input-number
          v-model="ll.rated_low_vol_min"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="额定低压上界（kV）"
        prop="rated_low_vol_max"
        label-width=""
      >
        <el-input-number
          v-model="ll.rated_low_vol_max"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="调压范围下界（%）"
        prop="regulate_range_min"
      >
        <el-select
          v-model="ll.regulate_range_min"
          filterable
          allow-create
          placeholder="调压范围下界（%）"
          style="width: 100%"
        >
          <el-option
            v-for="item in regulateMinOptions"
            :key="item"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="调压范围上界（%）"
        prop="regulate_range_max"
        label-width=""
        style="width: 100%"
      >
        <el-select
          v-model="ll.regulate_range_max"
          filterable
          allow-create
          placeholder="调压范围上界（%）"
          style="width: 100%"
        >
          <el-option
            v-for="item in regulateMaxOptions"
            :key="item"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="调压范围步长（%）"
        prop="regulate_range_step"
        label-width=""
      >
        <el-select
          v-model="ll.regulate_range_step"
          filterable
          allow-create
          placeholder="调压范围步长（%）"
          style="width: 100%"
        >
          <el-option
            v-for="item in regulateStepOptions"
            :key="item"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="绝缘系统温度（K）"
        prop="temperature"
        label-width=""
      >
        <el-input-number
          v-model="ll.temperature"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="负载损耗（W）"
        prop="load_loss"
        label-width=""
      >
        <el-input-number
          v-model="ll.load_loss"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="空载损耗（W）"
        prop="load_loss"
        label-width=""
      >
        <el-input-number
          v-model="ll.no_load_loss"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="空载电流（%）"
        prop="load_loss"
        label-width=""
      >
        <el-input-number
          v-model="ll.no_load_current"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
      <el-form-item
        label="短路阻抗（%）"
        prop="short_circuit_imped"
        label-width=""
      >
        <el-input-number
          v-model="ll.short_circuit_imped"
          size="medium"
          style="width: 100%"
          clearable
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="createLl"
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
import {defineComponent, computed, reactive, ref, Ref, onMounted} from 'vue'
import { CreateStdLlLibraries } from '@/api/standard_libraries/ll'
import { ElMessage } from 'element-plus'
import {GetStdLrrLibraries, ListLrrQuery, Lrr} from "@/api/standard_libraries/lrr";

export default defineComponent({
  name: 'LlCreate',
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue', 'get-ll-list'],
  setup(props, context) {
    const llCreateForm = ref(null)
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        if (!value) {
          (llCreateForm.value as any).clearValidate()
        }
        context.emit('update:modelValue', value)
      }
    })
    const ll = reactive({
      loss_level: 0,
      regulate_way: '',
      rated_capacity_min: 0,
      rated_capacity_max: 0,
      rated_high_vol_min: 0,
      rated_high_vol_max: 0,
      rated_low_vol_min: 0,
      rated_low_vol_max: 0,
      regulate_range_min: 0,
      regulate_range_max: 0,
      regulate_range_step: 0,
      temperature: 0,
      load_loss: 0,
      no_load_loss: 0,
      no_load_current: 0,
      short_circuit_imped: 0
    })
    const regulateStepOptions : Ref<{label: string, value: string}[]> = ref([])
    const regulateMinOptions : Ref<{label: string, value: string}[]> = ref([])
    const regulateMaxOptions : Ref<{label: string, value: string}[]> = ref([])
    const getStdLrrLibraries = async () => {
      const res = await GetStdLrrLibraries({
        field_eq_regulate_range_step: 0,
        field_eq_regulate_range_min: 0,
        field_eq_regulate_range_max: 0,
        limit: 9999,
        page: 1,
        order: '',
        order_by: ''
      })
      if (res.data.code === 200) {
        regulateStepOptions.value = res.data.spec.map((i: Lrr) => {
          return {
            label: i.regulate_range_step + '%',
            value: i.regulate_range_step
          }
        })
        regulateMinOptions.value = res.data.spec.map((j: Lrr) => {
          return {
            label: j.regulate_range_min + '%',
            value: j.regulate_range_min
          }
        })
        regulateMaxOptions.value = res.data.spec.map((k: Lrr) => {
          return {
            label: k.regulate_range_max + '%',
            value: k.regulate_range_max
          }
        })
      }
    }
    function validateFactor_1(rule: any, value: any, callback: any) {
      if (!ll.regulate_way) {
        return callback(new Error('请输入调压方式'))
      }
      callback()
    }

    function validateFactor_2(rule: any, value: any, callback: any) {
      if (ll.rated_capacity_min == null) {
        return callback(new Error('请输入额定容量下界（kVA）'))
      }
      if (ll.rated_capacity_min < 0) {
        return callback(new Error('额定容量下界不能小于0'))
      }
      callback()
    }
    function validateFactor_3(rule: any, value: any, callback: any) {
      if (ll.rated_capacity_max == null) {
        return callback(new Error('请输入额定容量上界（kVA）'))
      }
      if (ll.rated_capacity_max < 0) {
        return callback(new Error('额定容量上界不能小于0'))
      }
      callback()
    }
    function validateFactor_4(rule: any, value: any, callback: any) {
      if (ll.rated_high_vol_min == null) {
        return callback(new Error('请输入额定高压下界（kV）'))
      }
      if (ll.rated_high_vol_min < 0) {
        return callback(new Error('额定高压下界不能小于0'))
      }
      callback()
    }
    function validateFactor_5(rule: any, value: any, callback: any) {
      if (ll.rated_high_vol_max == null) {
        return callback(new Error('请输入额定高压上界（kVA）'))
      }
      if (ll.rated_high_vol_max < 0) {
        return callback(new Error('额定高压上界不能小于0'))
      }
      callback()
    }
    function validateFactor_6(rule: any, value: any, callback: any) {
      if (ll.rated_low_vol_min == null) {
        return callback(new Error('请输入额定低压下界（kV）'))
      }
      if (ll.rated_low_vol_min < 0) {
        return callback(new Error('额定低压下界不能小于0'))
      }
      callback()
    }
    function validateFactor_7(rule: any, value: any, callback: any) {
      if (ll.rated_low_vol_max == null) {
        return callback(new Error('请输入额定低压上界（kV）'))
      }
      if (ll.rated_low_vol_max < 0) {
        return callback(new Error('额定低压上界不能小于0'))
      }
      callback()
    }
    function validateFactor_8(rule: any, value: any, callback: any) {
      if (ll.regulate_range_min === null) {
        return callback(new Error('请输入调压范围下界（%）'))
      }
      callback()
    }
    function validateFactor_9(rule: any, value: any, callback: any) {
      if (ll.regulate_range_max === null) {
        return callback(new Error('请输入调压范围上界（%）'))
      }
      callback()
    }
    function validateFactor_10(rule: any, value: any, callback: any) {
      if (ll.regulate_range_step === null) {
        return callback(new Error('请输入调压范围步长（%）'))
      }
      callback()
    }
    function validateFactor_11(rule: any, value: any, callback: any) {
      if (ll.temperature == null) {
        return callback(new Error('请输入绝缘系统温度（K）'))
      }
      callback()
    }
    function validateFactor_12(rule: any, value: any, callback: any) {
      if (ll.load_loss == null) {
        return callback(new Error('请输入负载损耗（W）'))
      }
      callback()
    }
    function validateFactor_13(rule: any, value: any, callback: any) {
      if (ll.no_load_loss == null) {
        return callback(new Error('请输入空载损耗（W）'))
      }
      callback()
    }
    function validateFactor_14(rule: any, value: any, callback: any) {
      if (ll.no_load_current == null) {
        return callback(new Error('请输入空载电流（%）'))
      }
      callback()
    }
    function validateFactor_15(rule: any, value: any, callback: any) {
      if (ll.short_circuit_imped == null) {
        return callback(new Error('请输入短路阻抗（%）'))
      }
      callback()
    }

    const rules = {
      regulate_way: [
        { required: true, validator: validateFactor_1, trigger: 'blur' }
      ],
      rated_capacity_min: [
        { required: true, validator: validateFactor_2, trigger: 'blur' }
      ],
      rated_capacity_max: [
        { required: true, validator: validateFactor_3, trigger: 'blur' }
      ],
      rated_high_vol_min: [
        { required: true, validator: validateFactor_4, trigger: 'blur' }
      ],
      rated_high_vol_max: [
        { required: true, validator: validateFactor_5, trigger: 'blur' }
      ],
      rated_low_vol_min: [
        { required: true, validator: validateFactor_6, trigger: 'blur' }
      ],
      rated_low_vol_max: [
        { required: true, validator: validateFactor_7, trigger: 'blur' }
      ],
      regulate_range_min: [
        { required: true, validator: validateFactor_8, trigger: 'blur' }
      ],
      regulate_range_max: [
        { required: true, validator: validateFactor_9, trigger: 'blur' }
      ],
      regulate_range_step: [
        { required: true, validator: validateFactor_10, trigger: 'blur' }
      ],
      temperature: [
        { required: true, validator: validateFactor_11, trigger: 'blur' }
      ],
      load_loss: [
        { required: true, validator: validateFactor_12, trigger: 'blur' }
      ],
      no_load_loss: [
        { required: true, validator: validateFactor_13, trigger: 'blur' }
      ],
      no_load_current: [
        { required: true, validator: validateFactor_14, trigger: 'blur' }
      ],
      short_circuit_imped: [
        { required: true, validator: validateFactor_15, trigger: 'blur' }
      ],
    }
    function createLl() {
      (llCreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          const response = await CreateStdLlLibraries(ll)
          if (response.data.code === 200) {
            ElMessage.success('创建成功')
            context.emit('get-ll-list')
            show.value = false
          }
        } else {
          ElMessage.error('请提供正确的信息')
        }
      })
    }
    // async function createLl() {
    //   const response = await CreateStdLlLibraries(ll)
    //   if (response.data.code === 200) {
    //     ElMessage.success('创建成功')
    //     context.emit('get-ll-list')
    //     show.value = false
    //   } else {
    //     ElMessage.error('请提供正确的信息')
    //   }
    // }
    const paramLoading = ref(true)
    onMounted(async () => {
      await Promise.all([
        // 分接范围
        getStdLrrLibraries()
      ])
      paramLoading.value = false
    })
    return {
      llCreateForm,
      show,
      ll,
      rules,
      regulateStepOptions,
      regulateMinOptions,
      regulateMaxOptions,
      createLl
    }
  }
})

</script>
