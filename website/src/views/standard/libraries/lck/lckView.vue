<template>
  <el-dialog
    v-model="show"
    title="铁心直径经验系数K"
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
        {{ lck.id }}
      </el-form-item>
      <el-form-item
        label="绕组材质"
        label-width=""
      >
        {{ lck.winding_material }}
      </el-form-item>
      <el-form-item
        label="额定容量下限(kVA)"
        label-width=""
      >
        {{ lck.rated_capacity_min }}
      </el-form-item>
      <el-form-item
        label="额定容量上限(kVA)"
        label-width=""
      >
        {{ lck.rated_capacity_max }}
      </el-form-item>
      <el-form-item
        label="k值下限"
        label-width=""
      >
        {{ lck.k_min }}
      </el-form-item>
      <el-form-item
        label="k值上限"
        label-width=""
      >
        {{ lck.k_max }}
      </el-form-item>
      <el-form-item
        label="创建者"
        label-width=""
      >
        {{ lck.creator_user?.user_name_zh }}
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
  name: 'LckView',
  props: {
    modelValue: Boolean,
    lck: {
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
