<template>
  <div class="app-container">
    <aside class="page-header">
      权限配置
    </aside>
    <div class="main-table">
      <el-tabs
        v-model="activeTab"
      >
        <el-tab-pane
          label="用户角色关联"
          name="user_role"
        >
          <div
            class="filter-container"
          >
            <div class="container-row">
              <el-input
                v-model="userRoleQuery['field_lk_user.user_name_zh']"
                placeholder="用户名"
                style="width: 200px"
                class="filter-item"
                @keyup.enter="handleUserRoleListFilter"
              />
              <RoleSearchInput
                v-model="userRoleQuery.field_lk_role_name"
                placeholder="角色名"
                class="filter-item"
                style="width: 200px"
              />
              <el-button
                v-waves
                class="filter-item"
                type="primary"
                icon="el-icon-search"
                @click="handleUserRoleListFilter"
              >
                搜索
              </el-button>
              <!--              <el-button-->
              <!--                class="filter-item"-->
              <!--                style="margin-left: 10px"-->
              <!--                type="primary"-->
              <!--                icon="el-icon-edit"-->
              <!--                @click="handleUserRoleCreate()"-->
              <!--              >-->
              <!--                创建用户角色关联-->
              <!--              </el-button>-->
              <!--              <el-button-->
              <!--                v-waves-->
              <!--                :loading="downloadLoading"-->
              <!--                class="filter-item"-->
              <!--                type="primary"-->
              <!--                icon="el-icon-download"-->
              <!--                @click="handleDownload"-->
              <!--              >-->
              <!--                导出-->
              <!--              </el-button>-->
            </div>
          </div>
          <el-empty
            v-if="!listLoading && userRoleList.length === 0"
            :image-size="200"
          />
          <div
            v-else
            v-loading="listLoading"
            class="user_role-container"
          >
            <el-card
              v-for="item in userRoleList"
              :key="item.username"
              class="user_role-card"
            >
              <template #header>
                <div class="card-header">
                  <span>{{ item.user_name_zh }}</span>
                  <el-button
                    class="button"
                    size="medium"
                    style="float: right;margin-top: -7px;"
                    type="success"
                    plain
                    @click="handleUserRoleCreate(item)"
                  >
                    添加角色绑定
                  </el-button>
                </div>
              </template>
              <div
                class="role-tag-container"
              >
                <el-tag
                  v-for="userRole in item.user_role"
                  :key="userRole.id"
                  closable
                  style="font-size: 16px;margin: 0 10px 10px 0"
                  @close="handleUserRoleDelete(userRole)"
                >
                  {{ roleLabelMap[userRole.role_name] }}
                </el-tag>
                <RoleSearchInput
                  v-if="tempUserRoleShow && currentUserRole.username === item.username"
                  ref="selectUserRoleRef"
                  v-model="tempUserRoleDetail.role_name"
                  :username="item.username"
                  size="small"
                  placeholder="角色名"
                  class="filter-item"
                  :clearable="false"
                  style="width: 150px;top: -2px;margin-left: 10px;"
                  @handleChange="confirmUserRoleCreate"
                  @handleVisible="handleUserRoleSelect"
                  @ref="$event.focus()"
                />
              </div>
            </el-card>
          </div>
        </el-tab-pane>
        <el-tab-pane
          label="角色权限"
          name="role_permission"
        >
          <div class="filter-container">
            <div class="container-row">
              <el-input
                v-model="roleQuery.field_lk_cn_name"
                placeholder="角色名"
                clearable
                style="width: 200px"
                class="filter-item"
              />
              <el-button
                v-waves
                class="filter-item"
                type="primary"
                icon="el-icon-search"
                @click="handleRoleListFilter"
              >
                搜索
              </el-button>
              <el-button
                class="filter-item"
                style="margin-left: 10px"
                type="primary"
                icon="el-icon-edit"
                @click="handleRoleCreate()"
              >
                创建角色
              </el-button>
              <!--              <el-button-->
              <!--                v-waves-->
              <!--                :loading="downloadLoading"-->
              <!--                class="filter-item"-->
              <!--                type="primary"-->
              <!--                icon="el-icon-download"-->
              <!--                @click="handleDownload"-->
              <!--              >-->
              <!--                导出-->
              <!--              </el-button>-->
            </div>
          </div>
          <el-table
            key="role_permission"
            v-loading="roleListLoading"
            :data="roles"
            border
            fit
            highlight-current-row
            style="width: 100%"
            @sort-change="sortChange"
          >
            <el-table-column
              label="ID"
              prop="id"
              sortable="custom"
              align="center"
              width="150"
            >
              <template #default="{row}">
                <span>{{ row.id }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="角色"
              prop="name"
              sortable="custom"
              align="center"
              width="300"
            >
              <template #default="{row}">
                <span>{{ row.cn_name }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="描述"
              prop="description"
              sortable="custom"
              width="400"
              align="center"
            >
              <template #default="{row}">
                <span>{{ row.description }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="操作"
              align="center"
              class-name="fixed-width"
            >
              <template #default="{row}">
                <!--                <el-button-->
                <!--                  type="primary"-->
                <!--                  size="mini"-->
                <!--                  @click="handleRolePermissionDetail(row)"-->
                <!--                >-->
                <!--                  权限详情-->
                <!--                </el-button>-->
                <el-button
                  size="mini"
                  type="success"
                  @click="handleRoleUpdate(row)"
                >
                  修改
                </el-button>
                <el-button
                  :disabled="['root'].includes(row.name)"
                  size="mini"
                  type="warning"
                  @click="handleRolePermissionUpdate(row)"
                >
                  配置权限
                </el-button>
                <el-button
                  size="mini"
                  type="warning"
                  @click="handleTemporaryUserRole(row)"
                >
                  权限指派
                </el-button>
                <el-button
                  :disabled="['root','demander','designer','reviewer','chief_engineer'].includes(row.name)"
                  type="danger"
                  size="mini"
                  @click="handleRoleDelete(row.id)"
                >
                  删除
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
    </div>
    <div
      class="pagination-footer"
      style="margin-top: 1%"
    >
      <el-pagination
        v-if="activeTab === 'user_role' && userRoleTotal > 0"
        v-model:current-page="userRoleQuery.page"
        v-model:page-size="userRoleQuery.limit"
        :total="userRoleTotal"
        :page-sizes="[12, 24, 72, 144]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="getUserRoleList"
        @current-change="getUserRoleList"
      />
      <el-pagination
        v-if="activeTab === 'role_permission' && roleTotal > 0"
        v-model:current-page="roleQuery.page"
        v-model:page-size="roleQuery.limit"
        :total="roleTotal"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleRoleListFilter"
        @current-change="handleRoleListFilter"
      />
    </div>
    <div class="dialogs-main">
      <el-dialog
        v-model="showUserRoleCreate"
        title="创建用户角色关联"
        width="800px"
      >
        <template #title>
          <div>
            <span class="el-dialog__title">创建用户角色关联</span>
          </div>
        </template>
        <el-form
          ref="userRoleCreateForm"
          :model="tempUserRoleDetail"
          :rules="userRoleFormRules"
          label-width="100px"
          label-suffix=":"
        >
          <el-form-item
            label="用户名"
            prop="username"
            label-width=""
          >
            <UserSearchInput
              v-model="tempUserRoleDetail.username"

              style="width: 100%;"

              :only-name="true"
            />
          </el-form-item>
          <el-form-item
            label="角色"
            prop="role_name"
            label-width=""
          >
            <RoleSearchInput
              v-model="tempUserRoleDetail.role_name"
              style="width: 100%"
            />
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            type="primary"
            @click="confirmUserRoleCreate()"
          >
            确定
          </el-button>
          <el-button @click="showUserRoleCreate = false">
            取消
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="showUserRoleUpdate"
        title="用户角色修改"
        width="800px"
      >
        <template #title>
          <div>
            <span class="el-dialog__title">用户角色修改</span>
          </div>
        </template>
        <el-form
          ref="userRoleUpdateForm"
          :model="tempUserRoleDetail"
          :rules="userRoleFormRules"
          label-width="100px"
          label-suffix=":"
        >
          <el-form-item
            label="ID"
            prop="id"
            label-width=""
          >
            {{ tempUserRoleDetail.id }}
          </el-form-item>
          <el-form-item
            label="用户名"
            prop="username"
            label-width=""
          >
            <UserSearchInput
              v-model="tempUserRoleDetail.username"
              style="width: 100%"
              :only-name="true"
            />
          </el-form-item>
          <el-form-item
            label="角色"
            prop="username"
            label-width=""
          >
            <RoleSearchInput
              v-model="tempUserRoleDetail.role_name"
              style="width: 100%"
            />
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            type="primary"
            @click="confirmUserRoleUpdate()"
          >
            确定
          </el-button>
          <el-button @click="showUserRoleUpdate = false">
            取消
          </el-button>
        </template>
      </el-dialog>

      <el-dialog
        v-model="showRolePermissionDialog"
        custom-class="custom-dialog"
        width="800px"
        :title="rolePermissionDialogAction==='detail'?'角色权限详情':'修改角色权限'"
      >
        <template #title>
          <div>
            <span class="el-dialog__title">角色权限详情</span>
          </div>
        </template>
        <template #footer>
          <el-button
            v-show="rolePermissionDialogAction==='update'"
            type="primary"
            @click="handleConfirmUpdateRolePermission()"
          >
            保存
          </el-button>
          <el-button
            v-show="rolePermissionDialogAction==='update'"
            @click="showRolePermissionDialog = false"
          >
            取消
          </el-button>
        </template>
        <div class="filter-container">
          <div class="container-row">
            <el-input
              v-model="pageRecourseSearchTemp"
              placeholder="资源名称"
              style="width: calc(100% - 103px);"
              class="filter-item"
              @keyup.enter="handleSearchResources()"
            />
            <el-button
              v-waves
              class="filter-item"
              type="primary"
              width="width: 93px;"
              icon="el-icon-search"
              @click="handleSearchResources()"
            >
              搜索
            </el-button>
          </div>
        </div>
        <el-table
          border
          :data="rolePermissionData"
          style="width: 100%"
        >
          <el-table-column
            label="资源"
            width="180"
          >
            <template #default="scope">
              <span>{{ resourceLabelMap[scope.row] }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="操作"
          >
            <template #default="scope">
              <el-form :disabled="rolePermissionDialogAction === 'detail'">
                <div
                  class="checkbox-container"
                >
                  <el-checkbox
                    v-for="aName in Object.keys(tempRolePermissionDetail[scope.row])"
                    :key="aName"
                    v-model="tempRolePermissionDetail[scope.row][aName]"
                    style="padding: 5px 10px;box-sizing: content-box;margin-right: 0;"
                    :label="actionLabelMap[aName]"
                    name="type"
                  />
                </div>
              </el-form>
            </template>
          </el-table-column>
        </el-table>
        <el-pagination
          v-model:current-page="rolePermissionPage"
          v-model:page-size="rolePermissionLimit"
          style="margin-top: 10px;"
          :total="rolePermissionDataTotal"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="rolePermissionLimit = $event"
          @current-change="rolePermissionPage = $event"
        />
      </el-dialog>
      <el-dialog
        v-model="showRoleCreate"
        width="600px"
        title="创建角色"
      >
        <template #title>
          <div>
            <span class="el-dialog__title">创建角色</span>
          </div>
        </template>
        <el-form
          ref="roleCreateForm"
          :model="tempRole"
          :rules="roleFormRules"
          label-width="100px"
          label-suffix=":"
        >
          <el-form-item
            label="角色标识"
            prop="name"
            label-width=""
          >
            <el-input
              v-model="tempRole.name"
              placeholder="角色标识"
              clearable
            />
          </el-form-item>
          <el-form-item
            label="角色名称"
            prop="cn_name"
            label-width=""
          >
            <el-input
              v-model="tempRole.cn_name"
              placeholder="角色名称"
              clearable
            />
          </el-form-item>
          <el-form-item
            label="角色描述"
            prop="description"
            label-width=""
          >
            <el-input
              v-model="tempRole.description"
              placeholder="角色描述"
              type="textarea"
              :rows="5"
              clearable
            />
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            type="primary"
            @click="handleConfirmRoleCreate()"
          >
            确定
          </el-button>
          <el-button @click="showRoleCreate = false">
            取消
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="showRoleUpdate"
        width="800px"
        title="修改角色"
      >
        <template #title>
          <div>
            <span class="el-dialog__title">修改角色</span>
          </div>
        </template>
        <el-form
          ref="roleUpdateForm"
          :model="tempRole"
          :rules="roleFormRules"
          label-width="100px"
          label-suffix=":"
        >
          <el-form-item
            label="角色标识"
            prop="name"
            label-width=""
          >
            <el-input
              v-model="tempRole.name"
              placeholder="角色标识"
              clearable
            />
          </el-form-item>
          <el-form-item
            label="角色名称"
            prop="cn_name"
            label-width=""
          >
            <el-input
              v-model="tempRole.cn_name"
              placeholder="角色名称"
              clearable
            />
          </el-form-item>
          <el-form-item
            label="角色描述"
            prop="description"
            label-width=""
          >
            <el-input
              v-model="tempRole.description"
              placeholder="角色描述"
              type="textarea"
              :rows="5"
              clearable
            />
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            type="primary"
            @click="handleConfirmRoleUpdate()"
          >
            确定
          </el-button>
          <el-button @click="showRoleUpdate = false">
            取消
          </el-button>
        </template>
      </el-dialog>
      <TemporaryUserRole
        v-model="showTemporaryUserRole"
        :handle-value="tempRole"
      />
    </div>
  </div>
</template>
<script lang="ts">
import { computed, ComputedRef, defineComponent, onMounted, reactive, Ref, ref, watch, watchEffect } from 'vue'
import {
  CreateRole,
  CreateRolePermissions,
  CreateUserRole, DeleteRole,
  DeleteRolePermission,
  DeleteUserRole,
  GetRoleList,
  getRolePermissions,
  GetUserRole, UpdateRole
} from '@/api/permission'
import type { RolePermissionQuery, RoleQuery, UserRoleQuery, UserRoles } from '@/model/permission'
import { ActionList, ResourceList, Role, RolePermission, UserRole } from '@/model/permission'
import { ElMessage, ElMessageBox } from 'element-plus'
import { tLoading } from '@/utils/loading'
import UserSearchInput from '@/components/user-search-input/index.vue'
import RoleSearchInput from '../components/role-search-input/index.vue'
import TemporaryUserRole from './temporary-user_role.vue'
export default defineComponent({
  components: {
    UserSearchInput,
    RoleSearchInput,
    TemporaryUserRole
  },
  setup() {
    const activeTab = ref('user_role')
    // role
    const roles : Ref<Role[]> = ref([])
    const roleQuery: RoleQuery = reactive({
      field_lk_cn_name: '',
      page: 1,
      limit: 10,
      order: '',
      order_by: ''
    })
    const roleListLoading = ref(false)
    const roleTotal = ref(0)
    async function getRoles() {
      roleListLoading.value = true
      const res = await GetRoleList(roleQuery)
      if (res.data.code === 200) {
        roleTotal.value = res.data.total
        roles.value = res.data.total ? res.data.spec : []
      }
      roleListLoading.value = false
    }
    async function getAllRoles() {
      roleQuery.limit = 9999
      roleQuery.page = 1
      const res = await GetRoleList(roleQuery)
      if (res.data.code === 200) {
        roleTotal.value = res.data.total
        allRoles.value = res.data.total ? res.data.spec : []
      }
      roleQuery.limit = 10
      roleQuery.page = 1
    }

    function handleRoleListFilter() {
      getRoles()
    }
    const resources = Array<any>()
    const resourceName = {
      [ResourceList.User]: '用户',
      [ResourceList.Department]: '部门',
      [ResourceList.Documents]: '文档',
      [ResourceList.PermissionRule]: '权限',
      [ResourceList.DesignProject]: '任务书管理',
      [ResourceList.DesignTask]: '产品设计',
      [ResourceList.Message]: '消息',
      [ResourceList.Libraries]: '标准库',
      [ResourceList.Dashboard]: '首页',
      [ResourceList.PriceManagement]: '材料价格管理'
    }
    for (const r in ResourceList) {
      const rValue = ResourceList[r as keyof typeof ResourceList]
      resources.push({
        name: resourceName[rValue],
        value: rValue
      })
    }
    const resourceOptions = computed(() => resources.map((r) => {
      return { value: r.value, label: r.name }
    }))
    const resourceLabelMap = computed(() => {
      const mapObj = Object()
      resources.map((r) => {
        mapObj[r.value] = r.name
      })
      return mapObj
    })
    const actions = Array<any>()
    const actionName = {
      [ActionList.Read]: '访问',
      [ActionList.Write]: '修改',
      [ActionList.Cancel]: '撤销',
      [ActionList.Design]: '设计',
      [ActionList.Review]: '复核',
      [ActionList.Approve]: '审核'
    }
    for (const r in ActionList) {
      const rValue = ActionList[r as keyof typeof ActionList]
      actions.push({
        name: actionName[rValue],
        value: rValue
      })
    }
    const actionLabelMap = computed(() => {
      const mapObj = Object()
      actions.map((r) => {
        mapObj[r.value] = r.name
      })
      return mapObj
    })
    const allRoles : Ref<Role[]> = ref([])
    const roleOptions = computed(() => allRoles.value.map((role) => {
      return { value: role.name, label: role.cn_name }
    }))
    const roleLabelMap = computed(() => {
      const mapObj = Object()
      allRoles.value.map((r) => {
        mapObj[r.name] = r.cn_name
      })
      return mapObj
    })

    // user_role
    // 列表
    const userRoleQuery: UserRoleQuery = reactive({
      field_lk_username: '',
      field_lk_role_name: '',
      page: 1,
      limit: 12,
      order: '',
      order_by: ''
    })
    const userRoleTotal = ref(0)
    const userRoleList: Ref<UserRoles[]> = ref([])
    const listLoading = ref(false)
    function sortChange(column: any) {
      switch (activeTab.value) {
        case 'user_role': {
          console.log(column)
          userRoleQuery.order_by = column.prop
          userRoleQuery.order = column.order === 'descending' ? 'desc' : 'asc'
          getUserRoleList()
          break
        }
        case 'role_permission': {
          roleQuery.order_by = column.prop
          roleQuery.order = column.order === 'descending' ? 'desc' : 'asc'
          getRoles()
          break
        }
      }
    }
    async function getUserRoleList() {
      listLoading.value = true
      const response = await GetUserRole(userRoleQuery)
      if (response.data.code === 200) {
        userRoleTotal.value = response.data.total
        userRoleList.value = response.data.total ? response.data.spec : []
        console.log(userRoleList)
      }
      listLoading.value = false
    }
    function handleUserRoleListFilter() {
      getUserRoleList()
    }

    const tempUserRoleDetail : Ref<UserRole> = ref({
      id: 0,
      username: '',
      user_name_zh: '',
      role_name: ''
    })
    const tempUserRoleShow = ref(false)
    // 修改
    const emptyUserRoles : UserRoles = {
      username: '',
      user_name_zh: '',
      user_role: []
    }
    const currentUserRole : Ref<UserRoles> = ref(JSON.parse(JSON.stringify(emptyUserRoles)) as UserRoles)

    function handleUserRoleCreate(row : UserRoles) {
      currentUserRole.value = JSON.parse(JSON.stringify(row)) as UserRoles
      tempUserRoleDetail.value.username = row.username
      tempUserRoleShow.value = true
    }
    async function confirmUserRoleCreate() {
      if (tempUserRoleDetail.value.role_name === '') {
        tempUserRoleShow.value = false
        return
      }
      const response = await CreateUserRole(tempUserRoleDetail.value)
      tempUserRoleDetail.value.role_name = ''
      if (response.data.code === 200) {
        ElMessage.success('用户角色更新成功')
        handleUserRoleListFilter()
      }
      tempUserRoleShow.value = false
    }
    function handleUserRoleSelect(value: Boolean) {
      if (!value) {
        tempUserRoleShow.value = false
      }
    }
    // 删除
    function handleUserRoleDelete(row: UserRole) {
      ElMessageBox.confirm('是否删除此用户与角色的关联', '确认删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteUserRole(row.id!)
            if (response.data.code === 200) {
              ElMessage.success('用户角色关联删除成功!')
              handleUserRoleListFilter()
            }
          }
        }
      })
    }
    // role_permission
    // 详情 + 修改
    // 本地分页
    const rolePermissionPage = ref(1)
    const rolePermissionLimit = ref(10)
    const pageRecourseSearch = ref('')
    const pageRecourseSearchTemp = ref('')
    function handleSearchResources() {
      pageRecourseSearch.value = pageRecourseSearchTemp.value
    }
    const rolePermissionData: ComputedRef<string[]> = computed(() => {
      return Object.keys(tempRolePermissionDetail.value).filter((v) => {
        if (pageRecourseSearch.value === '') {
          return true
        }
        return resourceLabelMap.value[v].includes(pageRecourseSearch.value)
      }).slice((rolePermissionPage.value - 1) * rolePermissionLimit.value, rolePermissionPage.value * rolePermissionLimit.value)
    })
    const rolePermissionDataTotal : ComputedRef<number> = computed(() => {
      return Object.keys(tempRolePermissionDetail.value).filter((v) => {
        if (pageRecourseSearch.value === '') {
          return true
        }
        return resourceLabelMap.value[v].includes(pageRecourseSearch.value)
      }).length
    })
    const rolePermissionQuery: RolePermissionQuery = {
      field_eq_role: '',
      field_eq_resource: '',
      field_eq_action: '',
      page: 1,
      limit: 99999,
      order: '',
      order_by: ''
    }
    const rolePermissionTotal = ref(0)
    let rolePermissionList : RolePermission[] = []
    // 初始化 权限表单
    // 格式 {resourceName: {actionName: boolValue}}
    // example: {user: {read: true,write: false,cancel: false,design: false,review: false,approve: false}
    type ActionData = {
      [propName in ActionList]?: Boolean
    }
    type ResourceData = {
      [propName in ResourceList]?: ActionData
    }
    const emptyRolePermissionData: ResourceData = {}
    const tempRolePermissionDetail: Ref<ResourceData> = ref((JSON.parse(JSON.stringify(emptyRolePermissionData)) as ResourceData))
    for (const propName in ResourceList) {
      const enumPropName = ResourceList[propName as keyof typeof ResourceList]
      emptyRolePermissionData[enumPropName] = {}
      if (enumPropName === ResourceList.DesignTask) {
        Object.assign(emptyRolePermissionData[enumPropName], {
          [ActionList.Read]: false,
          [ActionList.Write]: false,
          [ActionList.Cancel]: false,
          [ActionList.Design]: false,
          [ActionList.Review]: false,
          [ActionList.Approve]: false
        })
      } else {
        Object.assign(emptyRolePermissionData[enumPropName], {
          [ActionList.Read]: false,
          [ActionList.Write]: false
        })
      }
    }
    async function getRolePermissionList() {
      const response = await getRolePermissions(rolePermissionQuery)
      if (response.data.code === 200) {
        rolePermissionTotal.value = response.data.total
        rolePermissionList = response.data.total ? response.data.spec : []
      }
    }
    const showRolePermissionDialog = ref(false)
    watch(showRolePermissionDialog, (value) => {
      if (!value) {
        rolePermissionPage.value = 1
        rolePermissionLimit.value = 10
        pageRecourseSearch.value = ''
        pageRecourseSearchTemp.value = ''
      }
    })
    const rolePermissionDialogAction = ref('detail')
    // 设置权限
    async function setTempRolePermissionDetail(roleName: string) {
      rolePermissionQuery.field_eq_role = roleName
      tempRolePermissionDetail.value = JSON.parse(JSON.stringify(emptyRolePermissionData))
      if (roleName === 'root') {
        for (const enumMember in ResourceList) {
          const isValueProperty = parseInt(enumMember, 10) >= 0
          if (!isValueProperty) {
            const resourceKey = ResourceList[enumMember as keyof typeof ResourceList]
            for (const actionEnumMember in ActionList) {
              const isValueProperty = parseInt(actionEnumMember, 10) >= 0
              if (!isValueProperty) {
                const actionKey = ActionList[actionEnumMember as keyof typeof ActionList]
                if (tempRolePermissionDetail.value[resourceKey]![actionKey] !== undefined) {
                  tempRolePermissionDetail.value[resourceKey]![actionKey] = true
                }
              }
            }
          }
        }
        return
      }
      await getRolePermissionList()
      rolePermissionList.map((rp) => {
        let resource : ResourceList | null = null
        let action : ActionList | null = null
        // 遍历value拿key or //ts ignore
        for (const enumMember in ResourceList) {
          const isValueProperty = parseInt(enumMember, 10) >= 0
          if (!isValueProperty) {
            const key = ResourceList[enumMember as keyof typeof ResourceList]
            if (key === rp.resource) {
              resource = key
            }
          }
        }
        for (const enumMember in ActionList) {
          const isValueProperty = parseInt(enumMember, 10) >= 0
          if (!isValueProperty) {
            const key = ActionList[enumMember as keyof typeof ActionList]
            if (key === rp.action) {
              action = key
            }
          }
        }
        if (resource && action) {
          tempRolePermissionDetail.value[resource]![action] = true
        }
      })
    }
    const rolePermissionDialogLoading = ref(false)
    // 查询某个角色的详情
    tLoading(rolePermissionDialogLoading, '.custom-dialog')
    let currentRoleName = ''
    async function handleRolePermissionDetail(row: any) {
      rolePermissionDialogAction.value = 'detail'
      showRolePermissionDialog.value = true
      rolePermissionDialogLoading.value = true
      await setTempRolePermissionDetail(row.name)
      currentRoleName = row.name
      rolePermissionDialogLoading.value = false
    }
    let cloneTempRolePermissionDetail: ResourceData
    async function handleRolePermissionUpdate(row: any) {
      rolePermissionDialogAction.value = 'update'
      showRolePermissionDialog.value = true
      rolePermissionDialogLoading.value = true
      await setTempRolePermissionDetail(row.name)
      currentRoleName = row.name
      cloneTempRolePermissionDetail = JSON.parse(JSON.stringify(tempRolePermissionDetail.value))
      rolePermissionDialogLoading.value = false
    }
    const rolePermissionForm = ref(null)
    const rolePermissionRules = {}
    // diff permission
    async function handleConfirmUpdateRolePermission() {
      const createRolePermissionArr : RolePermission[] = []
      const deleteRolePermissionIDs: number[] = []
      for (const key in ResourceList) {
        const resource = key as keyof typeof ResourceList
        const resourceValue = ResourceList[resource]
        for (const kkey in ActionList) {
          const action = kkey as keyof typeof ActionList
          const actionValue = ActionList[action]
          // 结果跟原结果有差异
          if (cloneTempRolePermissionDetail[resourceValue]![actionValue] !== tempRolePermissionDetail.value[resourceValue]![actionValue]) {
            if (tempRolePermissionDetail.value[resourceValue]![actionValue]) {
              createRolePermissionArr.push({
                id: 0,
                role: currentRoleName,
                resource: resourceValue,
                action: actionValue
              })
            } else {
              // 从role_permission查看与当前资源，行为一致的,加入删除列表
              for (const rp of rolePermissionList) {
                if (rp.role === currentRoleName &&
                  rp.resource === resourceValue &&
                  rp.action === actionValue) {
                  deleteRolePermissionIDs.push(rp.id)
                }
              }
            }
          }
        }
      }
      console.log(createRolePermissionArr, deleteRolePermissionIDs)
      if (createRolePermissionArr.length > 0) {
        const createResponse = await CreateRolePermissions(createRolePermissionArr)
        if (createResponse.data.code !== 200) {
          return
        }
      }
      if (deleteRolePermissionIDs.length > 0) {
        const deleteResponse = await DeleteRolePermission(deleteRolePermissionIDs)
        if (deleteResponse.data.code === 200) {
          showRolePermissionDialog.value = false
          ElMessage.success('角色权限更新成功')
          return
        }
      }
      showRolePermissionDialog.value = false
      ElMessage.success('角色权限更新成功')
    }
    // 创建角色
    const showRoleCreate = ref(false)
    const roleCreateForm = ref(null)
    watch(showRoleCreate, function(value) {
      if (!value) {
        (roleCreateForm.value as any).clearValidate()
      }
    })
    function validateRoleName(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('请输入角色标识'))
      }
      callback()
    }
    function validateRoleCnName(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('请输入角色名'))
      }
      callback()
    }
    function validateRoleDescription(rule: any, value: any, callback: any) {
      if (!value) {
        return callback(new Error('请输入角色描述'))
      }
      callback()
    }
    const roleFormRules = {
      name: [
        { required: true, validator: validateRoleName, trigger: 'blur' }
      ],
      cn_name: [
        { required: true, validator: validateRoleCnName, trigger: 'blur' }
      ],
      description: [
        { required: true, validator: validateRoleDescription, trigger: 'blur' }
      ]
    }
    const emptyRole: Role = {
      name: '',
      cn_name: '',
      description: ''
    }
    const tempRole: Ref<Role> = ref(JSON.parse(JSON.stringify(emptyRole)) as Role)
    async function handleRoleCreate() {
      showRoleCreate.value = true
      tempRole.value = JSON.parse(JSON.stringify(emptyRole)) as Role
    }
    async function handleConfirmRoleCreate() {
      (roleCreateForm.value as any).validate(async (valid: boolean) => {
        if (!valid) {
          return
        }
        const res = await CreateRole(tempRole.value)
        if (res.data.code === 200) {
          ElMessage.success('角色创建成功！')
          getRoles()
        }
        showRoleCreate.value = false
      })
    }
    // 更新角色
    const showRoleUpdate = ref(false)
    const roleUpdateForm = ref(null)
    watch(showRoleUpdate, function(value) {
      if (!value) {
        (roleUpdateForm.value as any).clearValidate()
      }
    })
    async function handleRoleUpdate(row: Role) {
      showRoleUpdate.value = true
      tempRole.value = JSON.parse(JSON.stringify(row)) as Role
    }
    async function handleConfirmRoleUpdate() {
      (roleUpdateForm.value as any).validate(async (valid: boolean) => {
        if (!valid) {
          return
        }
        const res = await UpdateRole(tempRole.value)
        if (res.data.code === 200) {
          ElMessage.success('角色更新成功！')
          getRoles()
        }
        showRoleUpdate.value = false
      })
    }
    // 删除角色

    async function handleRoleDelete(id: number) {
      ElMessageBox.confirm('是否删除此角色，请注意此操作不可恢复！', '确认删除角色', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: async (value: string) => {
          if (value === 'confirm') {
            const response = await DeleteRole(id)
            if (response.data.code === 200) {
              ElMessage.success('角色删除成功!')
              getRoles()
            }
          }
        }
      })
    }
    // 指派权限
    const showTemporaryUserRole = ref(false)
    function handleTemporaryUserRole(row: Role) {
      tempRole.value = JSON.parse(JSON.stringify(row)) as Role
      showTemporaryUserRole.value = true
    }

    onMounted(async () => {
      await getAllRoles()
      getRoles()
      getUserRoleList()
      // getRolePermissionList()
    })
    return {
      // 用户角色
      activeTab,
      roleOptions,
      resourceOptions,
      resourceLabelMap,
      actionLabelMap,
      roleLabelMap,
      userRoleQuery,
      userRoleTotal,
      listLoading,
      userRoleList,
      tempUserRoleDetail,
      currentUserRole,
      tempUserRoleShow,
      handleUserRoleSelect,
      sortChange,
      handleRoleListFilter,
      handleUserRoleListFilter,
      getUserRoleList,
      handleUserRoleCreate,
      confirmUserRoleCreate,
      handleUserRoleDelete,
      // 角色权限
      rolePermissionPage,
      rolePermissionLimit,
      rolePermissionData,
      pageRecourseSearch,
      pageRecourseSearchTemp,
      rolePermissionDataTotal,
      handleSearchResources,
      roles,
      roleTotal,
      roleQuery,
      roleListLoading,
      rolePermissionTotal,
      showRolePermissionDialog,
      rolePermissionDialogLoading,
      rolePermissionDialogAction,
      tempRolePermissionDetail,
      rolePermissionForm,
      rolePermissionRules,
      handleRolePermissionUpdate,
      handleConfirmUpdateRolePermission,
      handleRoleDelete,
      // 创建角色
      showRoleCreate,
      tempRole,
      roleCreateForm,
      roleFormRules,
      handleRoleCreate,
      handleConfirmRoleCreate,
      // 更新角色
      showRoleUpdate,
      roleUpdateForm,
      handleRoleUpdate,
      handleConfirmRoleUpdate,
      // 权限指派
      showTemporaryUserRole,
      handleTemporaryUserRole
    }
  }
})
</script>

<style lang="scss" scoped>

.filter-container {
  .container-row:first-child {
    margin-bottom: 10px;
  }
  .filter-item:not(button) {
    margin-right: 10px;
  }
}
 :deep{
   .dialogs-main {
     .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner::after {
       border-color: #fff;
     }
     .el-checkbox__input.is-disabled+span.el-checkbox__label {
       color: unset;
     }
     .el-checkbox__input.is-checked .el-checkbox__inner, .el-checkbox__input.is-indeterminate .el-checkbox__inner {
       background-color: $--color-primary;
       border-color: $--color-primary;
     }
   }
   .dialogs-main {

     .custom-dialog {
       min-height: calc(80vh);
       margin: 0 auto !important;
       margin-top: 5vh !important;
       overflow: hidden;
       .el-dialog__body {
         height: 100%;
         overflow: auto;
       }
       .dialog__operator {
         position: absolute;
         right: 70px;
         top: 20px;
       }
       .el-dialog__headerbtn .el-dialog__close {
         color: #fff;
         font-size: 20px;
         font-weight: 700;
         &hover {
           color: #ccc;
         }
       }
       aside {
         color: $--color-primary;
         font-size: 20px;
         background: unset !important;
         padding: unset !important;
       }
     }
     .checkbox-container {
       padding:0 10px;
     }
     table {
       th {
         font-size: 18px;
         color: #000;
       }
       td {
         font-size: 16px;
         padding: 5px 0;
       }
       .el-checkbox__inner {
         width: 20px;
         height: 20px;
         &::after {
           height: 10px;
           width: 6px;
           left: 5px;
           top: 2px;
         }
       }
     }
   }
 }
 .user_role-container {
   min-height: 300px;
   display: flex;
   width: 90%;
   flex: 1;
   flex-flow: row wrap;
 }
 .user_role-card {
   margin: 0 10px 10px;
   width: calc(25% - 20px);
   min-height: 130px;
   :deep() {
     .el-card__header {
       background: $__color-primary;
       background-clip: border-box;
     }
   }
   .card-header {
     font-size: 18px;
     color: #fff;

   }
   .add-user_role {
     color: $__color-primary;
     font-size: 20px;
   }

 }

</style>
