<template>
  <el-dialog
    v-model="show"
    title="线圈端部距铁轭距离标准库"
    show-close
    width="25%"
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
        {{ lcfiy.id }}
      </el-form-item>
      <el-form-item
        label="额定高压下界(kV)"
        label-width=""
      >
        {{ lcfiy.rated_high_vol_min }}
      </el-form-item>
      <el-form-item
        label="额定高压上界(kV)"
        label-width=""
      >
        {{ lcfiy.rated_high_vol_max }}
      </el-form-item>
      <el-form-item
        label="额定低压下界(kV)"
        label-width=""
      >
        {{ lcfiy.rated_low_vol_min }}
      </el-form-item>
      <el-form-item
        label="额定低压上界(kV)"
        label-width=""
      >
        {{ lcfiy.rated_low_vol_max }}
      </el-form-item>
      <el-form-item
        label="低压线圈端部距铁轭距离(mm)"
        label-width=""
      >
        {{ lcfiy.low_vol_coil_from_iron_yoke }}
      </el-form-item>
      <el-form-item
        label="高压线圈端部距铁轭距离(mm)"
        label-width=""
      >
        {{ lcfiy.high_vol_coil_from_iron_yoke }}
      </el-form-item>
      <el-form-item
        label="创建者"
        label-width=""
      >
        {{ lcfiy.creator_user?.user_name_zh }}
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
  name: 'LcfiyView',
  props: {
    modelValue: Boolean,
    lcfiy: {
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
