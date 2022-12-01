<template>
  <el-dialog
    v-model="show"
    title="线圈和风道绝缘"
    show-close
    width="25%"
  >
    <el-form
      label-width="250px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="编号"
        label-width=""
      >
        {{ lcadi.id }}
      </el-form-item>
      <el-form-item
        label="绕制类型"
        label-width=""
      >
        {{ lcadi.wind_type }}
      </el-form-item>
      <el-form-item
        label="线圈内层绝缘（mm）"
        label-width=""
      >
        {{ lcadi.coil_inner_insulate }}
      </el-form-item>
      <el-form-item
        label="线圈外层绝缘（mm）"
        label-width=""
      >
        {{ lcadi.coil_outer_insulate }}
      </el-form-item>
      <el-form-item
        label="风道厚度可选值（mm）"
        label-width=""
      >
        {{ lcadi.air_duct_thick }}
      </el-form-item>
      <el-form-item
        label="风道内外层绝缘（mm）"
        label-width=""
      >
        {{ lcadi.air_duct_in_out_insulate }}
      </el-form-item>
      <el-form-item
        label="创建者"
        label-width=""
      >
        {{ lcadi.creator_user?.user_name_zh }}
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
  name: 'LcadiView',
  props: {
    modelValue: Boolean,
    lcadi: {
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
