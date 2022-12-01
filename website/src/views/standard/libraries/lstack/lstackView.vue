<template>
  <el-dialog
    v-model="show"
    title="硅钢片标准库"
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
        {{ lstack.id }}
      </el-form-item>
      <el-form-item
        label="硅钢片型号"
        label-width=""
      >
        {{ lstack.type }}
      </el-form-item>
      <el-form-item
        label="密度(kg/cm^3)"
        label-width=""
      >
        {{ lstack.density }}
      </el-form-item>
      <el-form-item
        label="单价"
        label-width=""
      >
        {{ lstack.price }}
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
  name: 'LstackView',
  props: {
    modelValue: Boolean,
    lstack: {
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
