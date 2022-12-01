<template>
  <el-dialog
    v-model="show"
    title="铁心磁密Bm初选值"
    show-close
    width="40%"
  >
    <el-form
      label-width="180px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="额定容量下界（kVA）"
      >
        {{ lmd.rated_capacity_min }}
      </el-form-item>
      <el-form-item
        label="额定容量上界（kVA）"
      >
        {{ lmd.rated_capacity_max }}
      </el-form-item>
      <el-form-item
        label="铁心磁密下界(kV)"
      >
        {{ lmd.magnet_density_min }}
      </el-form-item>
      <el-form-item
        label="铁心磁密上界(kV)"
      >
        {{ lmd.magnet_density_max }}
      </el-form-item>
      <el-form-item
        label="创建时间"
      >
        {{ lmd.created_at }}
      </el-form-item>
      <el-form-item
        label="更新时间"
      >
        {{ lmd.updated_at }}
      </el-form-item>
      <el-form-item
        label="创建者"
        label-width=""
      >
        {{ lmd.creator_user?.user_name_zh }}
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
  name: 'LmdView',
  props: {
    modelValue: Boolean,
    lmd: {
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
