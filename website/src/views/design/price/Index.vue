<template>
  <div>
    <div class="app-container">
      <aside class="page-header">
        材料价格管理
      </aside>
      <div
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
              :class="['grid-content bg-purple']"
              style="margin-top: 2%"
            >
              <el-card
                class="box-card"
                shadow="hover"
                style="cursor: pointer"
                @click="showStandardlibrariesDetails(libraries.libraries_short)"
              >
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
                  </el-descriptions>
                </div>
              </el-card>
            </div>
          </el-col>
        </el-row>
        <el-empty
          v-if="!listLoading && listData.length <= 0"
          :image-size="200"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, Ref } from 'vue'
import { Library, LibraryQuery } from '@/model/library'
import { GetLibraries } from '@/api/standard'
import { useRouter } from 'vue-router'

export default defineComponent({
  setup() {
    const router = useRouter()
    const listLoading = ref(false)
    const listData = ref<Library[]>([])
    async function getList() {
      listLoading.value = true
      const res = await GetLibraries({
        limit: 10,
        order: '',
        order_by: '',
        page: 1,
        field_eq_libraries_name: [
          'lib_wire', 'lib_stack', 'lib_resin'
        ]
      } as LibraryQuery)
      if (res.data.code === 200) {
        listData.value = res.data.spec
      }
      listLoading.value = false
    }
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
    onMounted(() => {
      getList()
    })
    return {
      listData,
      listLoading,
      showStandardlibrariesDetails
    }
  }
})
</script>

<style scoped lang="scss">
:deep(.libraries-list .el-card__header) {
  background: $green;
  background-clip: padding-box;
  color: #fff;
  font-size: 18px;
}
:deep(.libraries-list .el-card__body) {
  min-height: 80px;
}

</style>
