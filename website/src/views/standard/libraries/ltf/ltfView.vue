<template>
  <el-dialog
    v-model="show"
    title="工艺系数"
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
        {{ ltf.id }}
      </el-form-item>
      <el-form-item
        label="硅钢片片数（片）"
        label-width=""
      >
        {{ ltf.stack_amount }}
      </el-form-item>
      <el-form-item
        label="工艺系数"
        label-width=""
      >
        {{ ltf.tech_factor }}
      </el-form-item>
      <el-form-item
        label="创建者"
        label-width=""
      >
        {{ ltf.creator_user?.user_name_zh }}
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
  name: 'LtfView',
  props: {
    modelValue: Boolean,
    ltf: {
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
