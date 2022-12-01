<template>
  <el-dialog
    v-model="show"
    title="空载损耗，负载损耗，空载电流，短路阻抗"
    show-close
    width="25%"
  >
    <el-form
      label-width="100px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="损耗水平代号"
        label-width="170px"
      >
        {{ ll.loss_level }}
      </el-form-item>
      <el-form-item
        label="调压方式"
        label-width="170px"
      >
        {{ ll.regulate_way }}
      </el-form-item>
      <el-form-item
        label="额定容量下界（kVA）"
        label-width="170px"
      >
        {{ ll.rated_capacity_min }}
      </el-form-item>
      <el-form-item
        label="额定容量上界（kVA）"
        label-width="170px"
      >
        {{ ll.rated_capacity_max }}
      </el-form-item>
      <el-form-item
        label="额定高压下界(kV)"
        label-width="150px"
      >
        {{ ll.rated_high_vol_min }}
      </el-form-item>
      <el-form-item
        label="额定高压上界(kV)"
        label-width="150px"
      >
        {{ ll.rated_high_vol_max }}
      </el-form-item>
      <el-form-item
        label="额定低压下界(kV)"
        label-width="150px"
      >
        {{ ll.rated_low_vol_min }}
      </el-form-item>
      <el-form-item
        label="额定低压上界(kV)"
        label-width="150px"
      >
        {{ ll.rated_low_vol_max }}
      </el-form-item>
      <el-form-item
        label="调压范围下界（%）"
        label-width="150px"
      >
        {{ ll.regulate_range_min }}
      </el-form-item>
      <el-form-item
        label="调压范围上界（%）"
        label-width="150px"
      >
        {{ ll.regulate_range_max }}
      </el-form-item>
      <el-form-item
        label="调压范围步长（%）"
        label-width="150px"
      >
        {{ ll.regulate_range_step }}
      </el-form-item>
      <el-form-item
        label="绝缘系统温度（K）"
        label-width="150px"
      >
        {{ ll.temperature }}
      </el-form-item>
      <el-form-item
        label="负载损耗(kW)"
        label-width="150px"
      >
        {{ ll.load_loss }}
      </el-form-item>
      <el-form-item
        label="空载损耗（kW）"
        label-width="150px"
      >
        {{ ll.no_load_loss }}
      </el-form-item>
      <el-form-item
        label="空载电流（%）"
        label-width="150px"
      >
        {{ ll.no_load_current }}
      </el-form-item>
      <el-form-item
        label="短路阻抗（%）"
        label-width="150px"
      >
        {{ ll.short_circuit_imped }}
      </el-form-item>
      <el-form-item
        label="创建者"
        label-width=""
      >
        {{ ll.creator_user?.user_name_zh }}
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="show = false">
        关闭
      </el-button>
    </template>
  </el-dialog>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue'

export default defineComponent({
  name: 'LlView',
  props: {
    modelValue: Boolean,
    ll: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue'],
  setup(props, context) {
    const status = new Map([[1, '正常'], [2, '停用']])
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        context.emit('update:modelValue', value)
      }
    })
    return {
      status,
      show
    }
  }
})
</script>
