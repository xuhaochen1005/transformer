<template>
  <el-dialog
    v-model="show"
    title="噪声预测"
    show-close
    width="25%"
  >
    <el-form
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定容量下界（kVA）"
      >
        {{ lnp.rated_capacity_min }}
      </el-form-item>
      <el-form-item
        label="额定容量上界（kVA）"
      >
        {{ lnp.rated_capacity_max }}
      </el-form-item>
      <el-form-item
        label="额定高压下界（kV）"
      >
        {{ lnp.rated_high_vol_min }}
      </el-form-item>
      <el-form-item
        label="额定高压上界（kV）"
      >
        {{ lnp.rated_high_vol_max }}
      </el-form-item>
      <el-form-item
        label="冷却方式"
      >
        {{ lnp.cool_type }}
      </el-form-item>
      <el-form-item
        label="噪声预测（dB）"
      >
        {{ lnp.noise_predict }}
      </el-form-item>
      <el-form-item
        label="创建时间"
      >
        {{ lnp.created_at }}
      </el-form-item>
      <el-form-item
        label="更新时间"
      >
        {{ lnp.updated_at }}
      </el-form-item>
      <el-form-item
        label="创建者"
      >
        {{ lnp.creator_user?.user_name_zh }}
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
  name: 'LnpView',
  props: {
    modelValue: Boolean,
    lnp: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue'],
  setup(props, context) {
    const status = new Map([[1, '正常'], [2, '停用']])
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        context.emit('update:modelValue', value)
      }
    })
    return {
      status,
      show
    }
  }
})
</script>
