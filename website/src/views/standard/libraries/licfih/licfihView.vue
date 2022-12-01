<template>
  <el-dialog
    v-model="show"
    title="内线圈至铁芯距离标准库"
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
        {{ licfih.id }}
      </el-form-item>
      <el-form-item
        label="额定高压下界(kV)"
        label-width=""
      >
        {{ licfih.rated_high_vol_min }}
      </el-form-item>
      <el-form-item
        label="额定高压上界(kV)"
        label-width=""
      >
        {{ licfih.rated_high_vol_max }}
      </el-form-item>
      <el-form-item
        label="内线圈至铁芯距离最小值(mm)"
        label-width=""
      >
        {{ licfih.inner_coil_form_iron_heart_min }}
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
  name: 'LicfihView',
  props: {
    modelValue: Boolean,
    licfih: {
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
