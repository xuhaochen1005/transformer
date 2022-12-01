<template>
  <el-dialog
    v-model="show"
    title="高压外模台账标准库"
    show-close
    width="40%"
  >
    <el-form
      label-width="100px"
      label-position="left"
      label-suffix=":"
    >
      <el-form-item
        label="模具尺寸"
        label-width="150px"
      >
        {{ lhvoml.mod_size }}
      </el-form-item>
      <el-form-item
        label="模具类型"
        label-width="150px"
      >
        {{ lhvoml.mod_type }}
      </el-form-item>
      <el-form-item
        label="模具高度(mm)"
        label-width="150px"
      >
        {{ lhvoml.mod_height }}
      </el-form-item>
      <el-form-item
        label="模具数量(个)"
        label-width="150px"
      >
        {{ lhvoml.mod_amount }}
      </el-form-item>
      <el-form-item
        label="模具图号"
        label-width="150px"
      >
        {{ lhvoml.mod_pic }}
      </el-form-item>
      <el-form-item
        label="模具编号"
        label-width="150px"
      >
        {{ lhvoml.mod_num }}
      </el-form-item>
      <el-form-item
        label="适用产品"
        label-width="150px"
      >
        {{ lhvoml.mod_suit }}
      </el-form-item>
      <el-form-item
        label="凸台高度(封板尺寸)"
        label-width="150px"
      >
        {{ lhvoml.boss_width }}
      </el-form-item>
      <el-form-item
        label="是否开槽"
        label-width="150px"
      >
        {{ lhvoml.groove }}
      </el-form-item>
      <el-form-item
        label="螺母尺寸(mm)"
        label-width="150px"
      >
        {{ lhvoml.nut_size }}
      </el-form-item>
      <el-form-item
        label="A头尺寸(mm)"
        label-width="150px"
      >
        {{ lhvoml.a_size }}
      </el-form-item>
      <el-form-item
        label="抽头孔距(mm)"
        label-width="150px"
      >
        {{ lhvoml.tap_hole_distance }}
      </el-form-item>
      <el-form-item
        label="封板图号"
        label-width="150px"
      >
        {{ lhvoml.closure_pic }}
      </el-form-item>
      <el-form-item
        label="备注"
        label-width="150px"
      >
        {{ lhvoml.remark }}
      </el-form-item>
      <el-form-item
        label="创建时间"
        label-width="150px"
      >
        {{ lhvoml.created_at }}
      </el-form-item>
      <el-form-item
        label="更新时间"
        label-width="150px"
      >
        {{ lhvoml.updated_at }}
      </el-form-item>
      <el-form-item
        label="删除时间"
        label-width="150px"
      >
        {{ lhvoml.deleted_at }}
      </el-form-item>
      <el-form-item
        label="创建者"
        label-width=""
      >
        {{ lhvoml.creator_user?.user_name_zh }}
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
  name: 'LhvomlView',
  props: {
    modelValue: Boolean,
    lhvoml: {
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
