import { RouteRecordRaw } from 'vue-router'
import Layout from '@/layout/index.vue'

const PermissionRouter: Array<RouteRecordRaw> = [
  {
    path: '/permission',
    component: Layout,
    // redirect: '/profile/index',
    meta: {
      name: 'permission',
      title: '权限管理',
      icon: 'permission',
      roles: ['root']
    },
    children: [
      /* {
        path: 'index',
        component: () => import(/!* webpackChunkName: "permission" *!/ '@/views/profile/index.vue'),
        name: 'Profile',
        meta: {
          title: 'profile',
          icon: 'user',
          noCache: true
        }
      },
      {
        path: 'role',
        component: () => import(/!* webpackChunkName: "permission" *!/ '@/views/permission/role/index.vue'),
        name: 'rolePermission',
        meta: {
          title: '角色权限',
          roles: ['admin']
        }
      }, */
      {
        path: 'casbin',
        component: () => import(/* webpackChunkName: "permission" */ '@/views/permission/casbin/Index.vue'),
        name: 'casbin',
        meta: {
          title: '权限配置',
          icon: 'permission',
          noCache: true
        }
      }
    ]
  }
]

export default PermissionRouter
