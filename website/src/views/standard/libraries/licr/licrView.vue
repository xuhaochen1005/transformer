<template>
  <el-dialog
    v-model="show"
    title="铁心表(圆形)标准库"
    show-close
    width="25%"
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
        {{ licr.id }}
      </el-form-item>
      <el-form-item
        label="铁心直径"
        label-width=""
      >
        {{ licr.core_diameter }}
      </el-form-item>
      <el-form-item
        label="有无拉板"
        label-width=""
      >
        {{ pullplate.get(licr.pull_plate) }}
      </el-form-item>
      <el-form-item
        label="截面积(mm)"
        label-width=""
      >
        {{ licr.section_area }}
      </el-form-item>
      <el-form-item
        label="截面积(mm)"
        label-width=""
      >
        {{ licr.section_area }}
      </el-form-item>
      <el-form-item
        label="角重(mm)"
        label-width=""
      >
        {{ licr.iron_corner_weight }}
      </el-form-item>
      <el-form-item
        label="片宽(1级)(mm)"
        label-width=""
      >
        {{ licr.stack_width_1 }}
      </el-form-item>
      <el-form-item
        label="片宽(2级)(mm)"
        label-width=""
      >
        {{ licr.stack_width_2 }}
      </el-form-item>
      <el-form-item
        label="片宽(3级)(mm)"
        label-width=""
      >
        {{ licr.stack_width_3 }}
      </el-form-item>
      <el-form-item
        label="片宽(4级)(mm)"
        label-width=""
      >
        {{ licr.stack_width_4 }}
      </el-form-item>
      <el-form-item
        label="片宽(5级)(mm)"
        label-width=""
      >
        {{ licr.stack_width_5 }}
      </el-form-item>
      <el-form-item
        label="片宽(6级)(mm)"
        label-width=""
      >
        {{ licr.stack_width_6 }}
      </el-form-item>
      <el-form-item
        label="片宽(7级)(mm)"
        label-width=""
      >
        {{ licr.stack_width_7 }}
      </el-form-item>
      <el-form-item
        label="片宽(8级)(mm)"
        label-width=""
      >
        {{ licr.stack_width_8 }}
      </el-form-item>
      <el-form-item
        label="迭厚(1级)(mm)"
        label-width=""
      >
        {{ licr.stack_thick_1 }}
      </el-form-item>
      <el-form-item
        label="迭厚(2级)(mm)"
        label-width=""
      >
        {{ licr.stack_thick_2 }}
      </el-form-item>
      <el-form-item
        label="迭厚(3级)(mm)"
        label-width=""
      >
        {{ licr.stack_thick_3 }}
      </el-form-item>
      <el-form-item
        label="迭厚(4级)(mm)"
        label-width=""
      >
        {{ licr.stack_thick_4 }}
      </el-form-item>
      <el-form-item
        label="迭厚(5级)(mm)"
        label-width=""
      >
        {{ licr.stack_thick_5 }}
      </el-form-item>
      <el-form-item
        label="迭厚(6级)(mm)"
        label-width=""
      >
        {{ licr.stack_thick_6 }}
      </el-form-item>
      <el-form-item
        label="迭厚(7级)(mm)"
        label-width=""
      >
        {{ licr.stack_thick_7 }}
      </el-form-item>
      <el-form-item
        label="迭厚(8级)(mm)"
        label-width=""
      >
        {{ licr.stack_thick_8 }}
      </el-form-item>
      <el-form-item
        label="总迭厚(mm)"
        label-width=""
      >
        {{ licr.stack_thick_sum }}
      </el-form-item>
      <el-form-item
        label="主级实迭厚迭厚(mm)"
        label-width=""
      >
        {{ licr.main_level_real_stack_thick }}
      </el-form-item>
      <el-form-item
        label="创建者"
        label-width=""
      >
        {{ licr.creator_user?.user_name_zh }}
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
  name: 'LicrView',
  props: {
    modelValue: Boolean,
    licr: {
      type: Object,
      required: true
    }
  },
  emits: ['update:modelValue'],
  setup(props, context) {
    const pullplate = new Map([[2, '无'], [1, '有']])
    const show = computed({
      get() {
        return props.modelValue
      },
      set(value) {
        context.emit('update:modelValue', value)
      }
    })
    return {
      pullplate,
      show
    }
  }
})
</script>
