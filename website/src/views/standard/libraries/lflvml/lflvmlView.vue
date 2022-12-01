<template>
  <el-dialog
    v-model="show"
    title="箔绕低压模具台账标准库"
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
        {{ lflvml.id }}
      </el-form-item>
      <el-form-item
        label="模具尺寸"
        label-width=""
      >
        {{ lflvml.mod_size }}
      </el-form-item>
      <el-form-item
        label="平台高"
        label-width=""
      >
        {{ lflvml.platform_height }}
      </el-form-item>
      <el-form-item
        label="模具高度(mm)"
        label-width=""
      >
        {{ lflvml.mod_height }}
      </el-form-item>
      <el-form-item
        label="模具图号"
        label-width=""
      >
        {{ lflvml.mod_pic }}
      </el-form-item>
      <el-form-item
        label="模具编号"
        label-width=""
      >
        {{ lflvml.mod_num }}
      </el-form-item>
      <el-form-item
        label="模具数量(个)"
        label-width=""
      >
        {{ lflvml.mod_amount }}
      </el-form-item>
      <el-form-item
        label="适用产品型号"
        label-width=""
      >
        {{ lflvml.mod_suit }}
      </el-form-item>
      <el-form-item
        label="备注"
        label-width=""
      >
        {{ lflvml.remark }}
      </el-form-item>
      <el-form-item
        label="创建者"
        label-width=""
      >
        {{ lflvml.creator_user?.user_name_zh }}
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
  name: 'LflvmlView',
  props: {
    modelValue: Boolean,
    lflvml: {
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
