import { RouteRecordRaw } from 'vue-router'
import Layout from '@/layout/index.vue'
import { ActionList, ResourceList, RoleList } from '@/model/permission'

const DesignRouter: Array<RouteRecordRaw> = [
  {
    path: '/design',
    component: Layout,
    redirect: '/profile/index',
    meta: {
      name: 'user',
      title: '产品设计',
      permission: {
        action: ActionList.Read, resource: ResourceList.DesignProject
      },
      icon: 'design'
    },
    children: [
      {
        path: 'design_projects',
        component: () => import(/* webpackChunkName: "design_result" */ '@/views/design/projects/Index.vue'),
        name: 'design_projects',
        meta: {
          title: '任务书管理',
          icon: 'design',
          permission: {
            action: ActionList.Read, resource: ResourceList.DesignProject
          },
          noCache: true
        }
      },
      // {
      //   path: 'design_task',
      //   component: () => import(/* webpackChunkName: "design_result" */ '@/views/design/design-task/index.vue'),
      //   name: 'design_task',
      //   meta: {
      //     title: '产品设计管理',
      //     //   icon: 'user',
      //     noCache: true
      //   }
      // },
      {
        path: 'design_create',
        component: () => import(/* webpackChunkName: "design_result" */ '@/views/design/design-create.vue'),
        name: 'design_create',
        meta: {
          title: '新建设计',
          icon: 'design',
          permission: {
            action: ActionList.Write, resource: ResourceList.DesignTask
          },
          noCache: true
        }
      },
      /*    {
        path: 'design_recommend',
        component: () => import(/!* webpackChunkName: "design_result" *!/ '@/views/profile/index.vue'),
        name: 'design_recommend',
        meta: {
          title: '方案推荐',
          //  icon: 'user',
          noCache: true
        }
      }, */
      {
        path: 'design_check',
        component: () => import(/* webpackChunkName: "design_result" */ '@/views/design/design-check.vue'),
        name: 'design_check',
        meta: {
          title: '设计复核',
          icon: 'design',
          roles: [RoleList.Designer, RoleList.Reviewer, RoleList.ChiefEngineer, RoleList.Demander],
          noCache: true
        }
      },
      {
        path: 'design_result',
        component: () => import(/* webpackChunkName: "design_result" */ '@/views/design/design-results.vue'),
        name: 'design_result',
        meta: {
          title: '设计结果',
          icon: 'design',
          permission: {
            action: ActionList.Write, resource: ResourceList.DesignTask
          },
          noCache: true
        }
      },
      {
        path: 'design_result_detail',
        component: () => import(/* webpackChunkName: "design_result" */ '@/views/design/design-results-detail.vue'),
        name: 'design_result_detail'
      },
      {
        path: 'price_management',
        component: () => import(/* webpackChunkName: "design_result" */ '@/views/design/price/Index.vue'),
        name: 'price_management',
        meta: {
          title: '价格管理',
          icon: 'design',
          permission: {
            action: ActionList.Read, resource: ResourceList.PriceManagement
          },
          noCache: true
        }
      }
    ]
  }
]

export default DesignRouter
