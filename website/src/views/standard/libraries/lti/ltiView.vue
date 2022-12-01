<template>
  <el-dialog
    v-model="show"
    title="端绝缘距离"
    show-close
    width="25%"
  >
    <el-form
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定高压下界（kV）"
      >
        {{ lti.rated_high_vol_min }}
      </el-form-item>
      <el-form-item
        label="额定高压上界（kV）"
      >
        {{ lti.rated_high_vol_max }}
      </el-form-item>
      <el-form-item
        label="额定低压下界（kV）"
      >
        {{ lti.rated_low_vol_min }}
      </el-form-item>
      <el-form-item
        label="额定低压上界（kV）"
      >
        {{ lti.rated_low_vol_max }}
      </el-form-item>
      <el-form-item
        label="端绝缘距离（mm）"
      >
        {{ lti.terminal_insulate }}
      </el-form-item>
      <el-form-item
        label="创建时间"
      >
        {{ lti.created_at }}
      </el-form-item>
      <el-form-item
        label="更新时间"
      >
        {{ lti.updated_at }}
      </el-form-item>
      <el-form-item
        label="创建者"
      >
        {{ lti.creator_user?.user_name_zh }}
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
  name: 'LtiView',
  props: {
    modelValue: Boolean,
    lti: {
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
