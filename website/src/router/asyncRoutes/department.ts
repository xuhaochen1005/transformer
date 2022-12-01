import { RouteRecordRaw } from 'vue-router'
import Layout from '@/layout/index.vue'
import { ActionList, ResourceList } from '@/model/permission'

const DepartmentRouter: Array<RouteRecordRaw> = [
  {
    path: '/department',
    component: Layout,
    // redirect: '/department/index',
    meta: {
      name: 'department',
      title: '产品设计',
      icon: 'department',
      permission: {
        resource: ResourceList.Department,
        action: ActionList.Write
      }
    },
    children: [
      {
        path: 'index',
        component: () => import(/* webpackChunkName: "department" */ '@/views/department/admin/Index.vue'),
        name: 'department',
        meta: {
          name: 'department',
          title: '部门管理',
          icon: 'department'
        }
      }
    ]
  }
]

export default DepartmentRouter
