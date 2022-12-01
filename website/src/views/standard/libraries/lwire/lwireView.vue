<template>
  <el-dialog
    v-model="show"
    title="导线表标准库"
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
        {{ lwire.id }}
      </el-form-item>
      <el-form-item
        label="导线型号"
        label-width=""
      >
        {{ lwire.wire_type }}
      </el-form-item>
      <el-form-item
        label="导线材质"
        label-width=""
      >
        {{ lwire.wire_material }}
      </el-form-item>
      <el-form-item
        label="导线形状"
        label-width=""
      >
        {{ lwire.wire_shape }}
      </el-form-item>
      <el-form-item
        label="导线密度(kg/m^3)"
        label-width=""
      >
        {{ lwire.wire_density }}kg/m^3
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
  name: 'LwireView',
  props: {
    modelValue: Boolean,
    lwire: {
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
