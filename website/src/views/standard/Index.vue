<template>
  <div class="app-container">
    <aside class="page-header">
      标准库管理
    </aside>
    <div class="filter-container">
      <el-input
        v-model="listQuery.field_lk_libraries_description"
        placeholder="标注库名"
        style="width: 200px"
        class="filter-item"
        @keyup.enter="handleFilter"
      />
      <el-input
        v-model="listQuery.field_lk_libraries_source"
        placeholder="原始文件"
        clearable
        style="width: 200px"
        class="filter-item"
      />
      <UserSearchInput
        v-model="listQuery.field_eq_libraries_creator"
        style="width: 200px"
        class="filter-item"
        :placeholder="'创建者'"
        clearable
      />
      <el-select
        v-model="listQuery.field_eq_libraries_status"
        placeholder="标准库状态"
        clearable
        class="filter-item"
        style="width: 200px"
      >
        <el-option
          v-for="item in statusOptions"
          :key="item"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
      <el-button
        v-waves
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="handleFilter"
      >
        搜索
      </el-button>
      <!--      <el-button-->
      <!--        class="filter-item"-->
      <!--        style="margin-left: 10px"-->
      <!--        type="primary"-->
      <!--        icon="el-icon-edit"-->
      <!--        @click="handleCreate"-->
      <!--      >-->
      <!--        添加-->
      <!--      </el-button>-->
    </div>
    <div
      v-loading="listLoading"
      class="main-table libraries-list"
    >
      <el-row
        :gutter="20"
      >
        <el-col
          v-for="libraries in listData"
          :key="libraries"
          :span="6"
        >
          <div
            :class="['grid-content bg-purple',{'status-disabled': libraries.libraries_status === 2}]"
            style="margin-top: 2%"
          >
            <el-card class="box-card">
              <template #header>
                <div class="card-header">
                  <span>{{ libraries.libraries_description }}</span>
                </div>
              </template>
              <div>
                <el-descriptions

                  :column="2"
                >
                  <el-descriptions-item
                    label="创建者:"
                  >
                    {{ libraries.creator?.user_name_zh }}
                  </el-descriptions-item>
                  <el-descriptions-item
                    label="标准库状态:"
                    label-align="right"
                    align="center"
                  >
                    <el-tag :type="libraries.libraries_status===1?'success':'danger'">
                      {{ statusOptions.find(i => i.value === libraries.libraries_status)?.label }}
                    </el-tag>
                  </el-descriptions-item>
                </el-descriptions>

                <el-descriptions
                  column="1"
                >
                  <el-descriptions-item label="原始文件&emsp;：">
                    <el-tag size="small">
                      {{ libraries.libraries_source }}<br>
                    </el-tag>
                  </el-descriptions-item>
                  <el-descriptions-item label="创建时间&emsp;：">
                    {{ UnixTime2HumanWithStr(libraries.created_at) }}<br>
                  </el-descriptions-item>
                  <el-descriptions-item label="更新时间&emsp;：">
                    {{ UnixTime2HumanWithStr(libraries.updated_at) }}<br>
                  </el-descriptions-item>
                </el-descriptions>
              </div>
              <el-link
                class="linker"
                type="success"
                @click.prevent="showStandardlibrariesDetails(libraries.libraries_short)"
              >
                点击查看详情
              </el-link>
            </el-card>
          </div>
        </el-col>
      </el-row>
      <el-empty
        v-if="!listLoading && listData.length <= 0"
        :image-size="200"
      />
    </div>
    <div
      class="pagination-footer"
      style="margin-top: 1%"
    >
      <el-pagination
        v-show="total > 0"
        v-model:current-page="listQuery.page"
        v-model:page-size="listQuery.limit"
        :total="total"
        :page-sizes="[8, 16, 48, 96]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleFilter"
        @current-change="handleFilter"
      />
    </div>
    <div classs="dialogs-main">
      <standard-create v-model="showCreate" />
    </div>
  </div>
</template>
<script lang="ts">
import {
  defineComponent,
  ref,
  Ref,
  reactive, onMounted
} from 'vue'
import { GetLibraries } from '@/api/standard'
import { UnixTime2HumanWithStr } from '@/utils/timeutils'
import type { Library, LibraryQuery } from '@/model/library'
import UserSearchInput from '@/components/user-search-input/index.vue'
import { useRouter } from 'vue-router'
import { UnwrapNestedRefs } from '@vue/reactivity'
import standardCreate from './create.vue'
export default defineComponent({
  components: {
    UserSearchInput,
    standardCreate
  },
  setup() {
    const listQuery : UnwrapNestedRefs<LibraryQuery> = reactive({
      field_lk_libraries_description: '',
      field_lk_libraries_name: '',
      field_lk_libraries_source: '',
      limit: 8,
      order: '',
      order_by: '',
      page: 1
    })
    const router = useRouter()
    const listLoading = ref(false)
    const total = ref(0)
    const listData : Ref<Library[]> = ref([])
    async function getList() {
      listLoading.value = true
      const res = await GetLibraries(listQuery)
      if (res.data.code === 200) {
        total.value = res.data.total
        listData.value = res.data.spec
      }
      listLoading.value = false
    }
    function handleFilter() {
      getList()
    }
    const statusOptions = [
      { label: '启用', value: 1 },
      { label: '禁用', value: 2 }
    ]
    function showStandardlibrariesDetails(key:string) {
      router.push({
        name: key,
        params: {
          libraries: key
        }
      }
      ).catch(err => {
        console.warn(err)
      })
      /* if (router.noGoBack) {

       } else {
       router.go(-1)
       } */
    }
    const showCreate = ref(false)
    function handleCreate() {
      showCreate.value = true
    }
    onMounted(() => {
      getList()
    })
    return {
      UnixTime2HumanWithStr,
      listQuery,
      listLoading,
      total,
      listData,
      getList,
      handleFilter,
      statusOptions,
      showStandardlibrariesDetails,
      showCreate,
      handleCreate
    }
  }
})
</script>

<style lang="scss">
.filter-container {
  .container-row:first-child {
    margin-bottom: 10px;
  }
  .filter-item:not(button) {
    margin-right: 10px;
  }
}
.libraries-list .el-card__header {
  background: $__color-primary;
  background-clip: padding-box;
  color: #fff;
  font-size: 18px;
}
.libraries-list {
  .status-disabled {
    .el-card__header {
      background: $__color-danger
    }
  }
  .el-card__body {
    min-height: 250px;
    position: relative;
  }
  .linker {
    position: absolute;
    bottom: 10px;
  }
}

</style>
<style lang="scss" scoped>

  .filter-container {
    .container-row:first-child {
      margin-bottom: 10px;
    }
    .filter-item:not(button) {
      margin-right: 10px;
    }
  }
  .form-icon {
    font-size: 18px;
    &.warning {
      color: $yellow;
    }
  }
  .custom-table-popper {
    max-width: 500px;
  }
  </style>
