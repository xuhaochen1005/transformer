<template>
  <el-dialog
    v-model="show"
    title="绕组外绝缘介质"
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
        {{ lwim.id }}
      </el-form-item>
      <el-form-item
        label="绕组外绝缘介质"
        label-width=""
      >
        {{ lwim.wind_insulate_media }}
      </el-form-item>
      <el-form-item
        label="绕组外绝缘介质代号"
        label-width=""
      >
        {{ lwim.wind_insulate_media_sign }}
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
  name: 'LwimView',
  props: {
    modelValue: Boolean,
    lwim: {
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
