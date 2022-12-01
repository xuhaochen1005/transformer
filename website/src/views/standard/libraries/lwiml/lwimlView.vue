<template>
  <el-dialog
    v-model="show"
    title="绕线内模台账"
    show-close
    width="25%"
  >
    <el-form
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="模具直径（mm）"
      >
        {{ lwiml.mod_diameter }}
      </el-form-item>
      <el-form-item
        label="模具规格（mm*mm）"
      >
        {{ lwiml.mod_size }}
      </el-form-item>
      <el-form-item
        label="模具数量（个）"
      >
        {{ lwiml.mod_amount }}
      </el-form-item>
      <el-form-item
        label="模具图号"
      >
        {{ lwiml.mod_pic }}
      </el-form-item>
      <el-form-item
        label="模具型号"
      >
        {{ lwiml.mod_num }}
      </el-form-item>
      <el-form-item
        label="适用产品"
      >
        {{ lwiml.mod_suit }}
      </el-form-item>
      <el-form-item
        label="备注"
      >
        {{ lwiml.remark }}
      </el-form-item>
      <el-form-item
        label="创建者"
      >
        {{ lwiml.creator_user?.user_name_zh }}
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
  name: 'LwimlView',
  props: {
    modelValue: Boolean,
    lwiml: {
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
