<template>
  <el-dialog
    v-model="show"
    title="箔绕扁形模具台账标准库"
    show-close
    width="30%"
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
        {{ lfoml.id }}
      </el-form-item>
      <el-form-item
        label="模具尺寸(mm)"
        label-width=""
      >
        {{ lfoml.mod_size }}
      </el-form-item>
      <el-form-item
        label="模具高度(mm)"
        label-width=""
      >
        {{ lfoml.mod_height }}
      </el-form-item>
      <el-form-item
        label="模具图号"
        label-width=""
      >
        {{ lfoml.mod_pic }}
      </el-form-item>
      <el-form-item
        label="模具编号"
        label-width=""
      >
        {{ lfoml.mod_num }}
      </el-form-item>
      <el-form-item
        label="模具数量(个)"
        label-width=""
      >
        {{ lfoml.mod_amount }}
      </el-form-item>
      <el-form-item
        label="适用产品型号"
        label-width=""
      >
        {{ lfoml.mod_suit }}
      </el-form-item>
      <el-form-item
        label="备注"
        label-width=""
      >
        {{ lfoml.remark }}
      </el-form-item>
      <el-form-item
        label="创建者"
        label-width=""
      >
        {{ lfoml.creator_user?.user_name_zh }}
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
  name: 'LfomlView',
  props: {
    modelValue: Boolean,
    lfoml: {
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
