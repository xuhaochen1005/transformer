<template>
  <el-dialog
    v-model="show"
    title="扁铜(铝)线规格标准库"
    show-close
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
        {{ lfws.id }}
      </el-form-item>
      <el-form-item
        label="标称长(mm)"
        label-width=""
      >
        {{ lfws.std_length }}
      </el-form-item>
      <el-form-item
        label="标称宽(mm)"
        label-width=""
      >
        {{ lfws.std_width }}
      </el-form-item>
      <el-form-item
        label="截面积(mm^2)"
        label-width=""
      >
        {{ lfws.section_area }}
      </el-form-item>
      <el-form-item
        label="圆角半径(mm)"
        label-width=""
      >
        {{ lfws.round_corner }}
      </el-form-item>
      <el-form-item
        label="备注"
        label-width=""
      >
        {{ lfws.remark }}
      </el-form-item>
      <el-form-item
        label="创建者"
        label-width=""
      >
        {{ lfws.creator_user?.user_name_zh }}
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
  name: 'LfwsView',
  props: {
    modelValue: Boolean,
    lfws: {
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
