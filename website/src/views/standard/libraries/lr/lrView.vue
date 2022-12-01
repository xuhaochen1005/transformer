<template>
  <el-dialog
    v-model="show"
    title="树脂规格表标准库"
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
        {{ lr.id }}
      </el-form-item>
      <el-form-item
        label="树脂型号"
        label-width=""
      >
        {{ lr.type }}
      </el-form-item>
      <el-form-item
        label="树脂密度(kg/m^3)"
        label-width=""
      >
        {{ lr.density }}
      </el-form-item>
      <el-form-item
        label="单价"
        label-width=""
      >
        {{ lr.price }}
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
  name: 'LrView',
  props: {
    modelValue: Boolean,
    lr: {
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
