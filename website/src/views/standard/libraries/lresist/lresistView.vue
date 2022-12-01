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
        {{ lresist.id }}
      </el-form-item>
      <el-form-item
        label="导线材质"
        label-width=""
      >
        {{ lresist.wire_material }}
      </el-form-item>
      <el-form-item
        label="温度(℃)"
        label-width=""
      >
        {{ lresist.temp }}
      </el-form-item>
      <el-form-item
        label="电阻率(Ω·m)"
        label-width=""
      >
        {{ lresist.resistivity }}
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
  name: 'LresistView',
  props: {
    modelValue: Boolean,
    lresist: {
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
