<template>
  <el-dialog
    v-model="show"
    title="层间绝缘距离标准库"
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
        {{ lli.id }}
      </el-form-item>
      <el-form-item
        label="绕制类型"
        label-width=""
      >
        {{ lli.winding_type }}
      </el-form-item>
      <el-form-item
        label="额定低压下界(kV)"
        label-width=""
      >
        {{ lli.rated_low_vol_min }}
      </el-form-item>
      <el-form-item
        label="额定低压上界(kV)"
        label-width=""
      >
        {{ lli.rated_low_vol_max }}
      </el-form-item>
      <el-form-item
        label="层间电压下界(V)"
        label-width=""
      >
        {{ lli.layer_vol_min }}
      </el-form-item>
      <el-form-item
        label="层间电压上界(V)"
        label-width=""
      >
        {{ lli.layer_vol_max }}
      </el-form-item>
      <el-form-item
        label="层间电压下界(V)"
        label-width=""
      >
        {{ lli.layer_vol_min }}
      </el-form-item>
      <el-form-item
        label="层间绝缘距离(mm)"
        label-width=""
      >
        {{ lli.layer_insulate }}
      </el-form-item>
      <el-form-item
        label="创建者"
        label-width=""
      >
        {{ lli.creator_user?.user_name_zh }}
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
  name: 'LliView',
  props: {
    modelValue: Boolean,
    lli: {
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
