import { RouteRecordRaw } from 'vue-router'
import Layout from '@/layout/index.vue'
import { RoleList } from '@/model/permission'

const MaintenanceRouter: Array<RouteRecordRaw> = [
  {
    path: '/maintainence',
    component: Layout,
    redirect: '/profile/index',
    meta: {
      // name: 'user',
      title: '系统维护',
      icon: 'maintenance',
      roles: [RoleList.Root]
    },
    children: [
      {
        path: 'index',
        component: () => import(/* webpackChunkName: "maintainence" */ '@/views/maintenance/Index.vue'),
        name: 'index',
        meta: {
          name: 'user',
          title: '系统维护',
          icon: 'maintenance'
        }
      }
    ]
  }
]

export default MaintenanceRouter
