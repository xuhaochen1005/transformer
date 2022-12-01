<template>
  <el-dialog
    v-model="show"
    title="涡流损耗系数标准库"
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
        {{ lelf.id }}
      </el-form-item>
      <el-form-item
        label="额定容量(kVA)"
        label-width=""
      >
        {{ lelf.rated_capacity }}
      </el-form-item>
      <el-form-item
        label="涡流损耗系数(%)"
        label-width=""
      >
        {{ lelf.eddy_loss_factor }}
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
  name: 'LelfView',
  props: {
    modelValue: Boolean,
    lelf: {
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
