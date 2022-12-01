<template>
  <el-dialog
    v-model="show"
    title="相数"
    show-close
  >
    <el-form
      label-width="220px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="编号"
        label-width=""
      >
        {{ lpn.id }}
      </el-form-item>
      <el-form-item
        label="相数"
        label-width=""
      >
        {{ lpn.phase_num }}
      </el-form-item>
      <el-form-item
        label="相数代号"
        label-width=""
      >
        {{ lpn.phase_num_sign }}
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
  name: 'LpnView',
  props: {
    modelValue: Boolean,
    lpn: {
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
