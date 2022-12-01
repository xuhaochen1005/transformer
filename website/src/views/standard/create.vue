<template>
  <el-dialog
    v-model="show"
    class="custom-dialog"
    width="800px"
  >
    <template #title>
      <div>
        <span class="el-dialog__title">创建标准库</span>
      </div>
    </template>
    <el-form
      ref="CreateForm"
      :model="tempDetail"
      :rules="FormRules"
      label-width="100px"
      label-suffix=":"
    >
      <el-form-item
        label="&emsp;标准库标识"
        prop="libraries_name"
        label-width=""
      >
        <el-input
          v-model="tempDetail.libraries_name"
          placeholder="标准库标识"
        />
      </el-form-item>
      <el-form-item
        label="&emsp;标准库缩写"
        prop="libraries_short"
        label-width=""
      >
        <el-input
          v-model="tempDetail.libraries_short"
          placeholder="标准库缩写"
        />
      </el-form-item>
      <el-form-item
        label="&emsp;标准库名称"
        prop="libraries_description"
        label-width=""
      >
        <el-input
          v-model="tempDetail.libraries_description"
          placeholder="标准库名称"
        />
      </el-form-item>
      <el-form-item
        label="&emsp;标准库来源"
        prop="libraries_source"
        label-width=""
      >
        <el-input
          v-model="tempDetail.libraries_source"
          placeholder="标准库来源"
        />
      </el-form-item>
      <el-form-item
        label="标准库创建者"
        prop="libraries_creator"
        label-width=""
      >
        <UserSearchInput
          v-model="tempDetail.libraries_creator"
          style="width: 100%;"
          :users="tempDetail.creator"
        />
      </el-form-item>
      <el-form-item
        label="&emsp;标准库状态"
        prop="libraries_status"
        label-width=""
      >
        <el-select
          style="width: 100%;"
          v-model="tempDetail.libraries_status"
        >
          <el-option
            v-for="item in libraryStatusOptions"
            :key="item"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="confirmCreate()"
      >
        确认
      </el-button>
      <el-button
        @click="show = false"
      >
        取消
      </el-button>
    </template>
  </el-dialog>
</template>

<script lang='ts'>
import {
  computed,
  defineComponent, onMounted,
  ref
} from 'vue'

import UserSearchInput from '@/components/user-search-input/index.vue'
import type { Library } from '@/model/library'
import { LibraryStatus } from '@/model/library'
import { CreateLibrary } from '@/api/standard'
import { ElMessage } from 'element-plus'
export default defineComponent({
  components: {
    UserSearchInput
  },
  props: {
    modelValue: Boolean
  },
  setup(props, ctx) {
    const CreateForm = ref(null)

    const show = computed({
      get: () => props.modelValue,
      set: (v) => {
        if (!v) {
          (CreateForm.value as any).clearValidate()
        }
        ctx.emit('update:modelValue', v)
      }
    })
    const empty : Library = {
      libraries_name: '',
      libraries_short: '',
      libraries_description: '',
      libraries_source: '',
      libraries_status: LibraryStatus.Normal
    }

    const libraryStatusOptions = [
      { label: '启用', value: LibraryStatus.Normal },
      { label: '禁用', value: LibraryStatus.Disabled }
    ]

    const tempDetail = ref(JSON.parse(JSON.stringify(empty)) as Library)
    // 消息通用验证表单
    function validateLibrariesName(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('标准库标识不可为空'))
      }
      if (value.length > 255) {
        return callback(new Error('标准库标识长度不能大于255'))
      }
      callback()
    }
    function validateLibrariesShort(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('标准库缩写不可为空'))
      }
      if (value.length > 255) {
        return callback(new Error('标准库缩写长度不能大于255'))
      }
      callback()
    }
    function validateLibrariesDescription(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('标准库名称不可为空'))
      }
      if (value.length > 255) {
        return callback(new Error('标准库名称长度不能大于255'))
      }
      callback()
    }
    function validateLibrariesSource(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('标准库来源不可为空'))
      }
      if (value.length > 255) {
        return callback(new Error('标准库来源长度不能大于255'))
      }
      callback()
    }
    function validateLibrariesCreator(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('标准库创建者不可为空'))
      }
      callback()
    }
    function validateLibrariesStatus(rule: any, value: any, callback: any) {
      if (!(value in LibraryStatus)) {
        return callback(new Error('标准库状态不能为空'))
      }
      callback()
    }

    const FormRules = {
      librariesName: [
        { required: true, validator: validateLibrariesName, trigger: 'blur' }
      ],
      librariesShort: [
        { required: true, validator: validateLibrariesShort, trigger: 'blur' }
      ],
      librariesDescription: [
        { required: true, validator: validateLibrariesDescription, trigger: 'blur' }
      ],
      librariesCreator: [
        { required: true, validator: validateLibrariesCreator, trigger: 'blur' }
      ],
      librariesSource: [
        { required: true, validator: validateLibrariesSource, trigger: 'blur' }
      ],
      librariesStatus: [
        { required: true, validator: validateLibrariesStatus, trigger: 'blur' }
      ]
    }

    // 创建
    let libraryCreate: Library = {
      libraries_name: '',
      libraries_short: '',
      libraries_description: '',
      libraries_source: '',
      libraries_status: LibraryStatus.Normal
    }
    function confirmCreate() {
      (CreateForm.value as any).validate(async(valid: boolean) => {
        if (valid) {
          libraryCreate = JSON.parse(JSON.stringify(tempDetail.value)) as Library
          const response = await CreateLibrary(libraryCreate)
          if (response.data.code === 200) {
            ElMessage.success('消息发送成功')
            show.value = false
          }
        }
      })
    }
    return {
      show,
      tempDetail,
      FormRules,
      CreateForm,
      libraryStatusOptions,
      confirmCreate
    }
  }
})
</script>

<style lang="scss" scoped>
.form-icon {
  font-size: 18px;
  &.warning {
    color: $yellow;
  }
}
</style>
