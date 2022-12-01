<template>
  <el-dialog
    v-model="show"
    title="工频耐压标准库"
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
        {{ lifhv.id }}
      </el-form-item>
      <el-form-item
        label="额定高压下界(kV)"
        label-width=""
      >
        {{ lifhv.rated_high_vol_min }}
      </el-form-item>
      <el-form-item
        label="额定高压上界(kV)"
        label-width=""
      >
        {{ lifhv.rated_high_vol_max }}
      </el-form-item>
      <el-form-item
        label="工频耐压(kV)"
        label-width=""
      >
        {{ lifhv.industry_freq_hold_vol }}
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
  name: 'LifhvView',
  props: {
    modelValue: Boolean,
    lifhv: {
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
