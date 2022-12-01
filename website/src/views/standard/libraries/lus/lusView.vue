<template>
  <el-dialog
    v-model="show"
    title="变压器用途"
    show-close
  >
    <el-form
      label-width="250px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="编号"
        label-width=""
      >
        {{ lus.id }}
      </el-form-item>
      <el-form-item
        label="变压器用途"
        label-width=""
      >
        {{ lus.usage }}
      </el-form-item>
      <el-form-item
        label="变压器用途代号"
        label-width=""
      >
        {{ lus.usage_sign }}
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
  name: 'LusView',
  props: {
    modelValue: Boolean,
    lus: {
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
