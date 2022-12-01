<template>
  <el-dialog
    v-model="show"
    title="导线绝缘厚度表标准库"
    show-close
    width="25%"
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
        {{ lwi.id }}
      </el-form-item>
      <el-form-item
        label="导线编号"
        label-width=""
      >
        {{ lwi.wire_id }}
      </el-form-item>
      <el-form-item
        label="线厚区间下界(mm)"
        label-width=""
      >
        {{ lwi.width_min }}
      </el-form-item>
      <el-form-item
        label="线厚区间上界(mm)"
        label-width=""
      >
        {{ lwi.width_max }}
      </el-form-item>
      <el-form-item
        label="轴向绝缘厚(mm)"
        label-width=""
      >
        {{ lwi.axial_insulation }}
      </el-form-item>
      <el-form-item
        label="辐向绝缘厚(mm)"
        label-width=""
      >
        {{ lwi.radial_insulation }}
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
  name: 'LwiView',
  props: {
    modelValue: Boolean,
    lwi: {
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
