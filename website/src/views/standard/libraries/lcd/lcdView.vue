<template>
  <el-dialog
    v-model="show"
    title="电流密度标准库"
    show-close
  >
    <el-form
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="编号"
        label-width=""
      >
        {{ lcd.id }}
      </el-form-item>
      <el-form-item
        label="导线材料"
        label-width=""
      >
        {{ lcd.winding_material }}
      </el-form-item>
      <el-form-item
        label="电流密度下界(A/mm^2)"
        label-width=""
      >
        {{ lcd.current_density_max }}
      </el-form-item>
      <el-form-item
        label="电流密度上界(A/mm^2)"
        label-width=""
      >
        {{ lcd.current_density_max }}
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
  name: 'LcdView',
  props: {
    modelValue: Boolean,
    lcd: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue'],
  setup(props, context) {
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        context.emit('update:modelValue', value)
      }
    })
    return {
      show
    }
  }
})
</script>
