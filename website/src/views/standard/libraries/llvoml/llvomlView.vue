<template>
  <el-dialog
    v-model="show"
    title="低压外模台账"
    show-close
    width="25%"
  >
    <el-form
      label-width="100px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="模具直径（mm）"
        label-width="150px"
      >
        {{ llvoml.mod_diameter }}
      </el-form-item>
      <el-form-item
        label="模具类型（mm）"
        label-width="150px"
      >
        {{ llvoml.mod_type }}
      </el-form-item>
      <el-form-item
        label="模具高度（mm）"
        label-width="150px"
      >
        {{ llvoml.mod_height }}
      </el-form-item>
      <el-form-item
        label="模具数量（个）"
        label-width="150px"
      >
        {{ llvoml.mod_amount }}
      </el-form-item>
      <el-form-item
        label="模具型号"
        label-width="150px"
      >
        {{ llvoml.mod_num }}
      </el-form-item>
      <el-form-item
        label="适用产品"
        label-width="150px"
      >
        {{ llvoml.mod_suit }}
      </el-form-item>
      <el-form-item
        label="备注"
        label-width="150px"
      >
        {{ llvoml.remark }}
      </el-form-item>
      <el-form-item
        label="创建时间"
        label-width="150px"
      >
        {{ llvoml.created_at }}
      </el-form-item>
      <el-form-item
        label="更新时间"
        label-width="150px"
      >
        {{ llvoml.updated_at }}
      </el-form-item>
      <el-form-item
        label="创建者"
        label-width=""
      >
        {{ llvoml.creator_user?.user_name_zh }}
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
  name: 'LlvomlView',
  props: {
    modelValue: Boolean,
    llvoml: {
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
