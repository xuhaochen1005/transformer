<template>
  <el-dialog
    v-model="show"
    title="硅钢片性能"
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
        {{ lspd.id }}
      </el-form-item>
      <el-form-item
        label="硅钢片性能"
        label-width=""
      >
        {{ lspd.stack_type }}
      </el-form-item>
      <el-form-item
        label="铁心磁密（T）"
        label-width=""
      >
        {{ lspd.core_flux_density }}
      </el-form-item>
      <el-form-item
        label="单位铁损（W/kg）"
        label-width=""
      >
        {{ lspd.unit_iron_loss }}
      </el-form-item>
      <el-form-item
        label="单位质量磁化容量（VA/kg）"
        label-width=""
      >
        {{ lspd.unit_mass_magnet_capacity }}
      </el-form-item>
      <el-form-item
        label="单位面积接缝伏安值（VA/cm2）"
        label-width=""
      >
        {{ lspd.unit_area_seam_va }}
      </el-form-item>
      <el-form-item
        label="创建者"
        label-width=""
      >
        {{ lspd.creator_user?.user_name_zh }}
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
  name: 'LspdView',
  props: {
    modelValue: Boolean,
    lspd: {
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
