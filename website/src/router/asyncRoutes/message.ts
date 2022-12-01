import { RouteRecordRaw } from 'vue-router'
import Layout from '@/layout/index.vue'
import { ActionList, ResourceList } from '@/model/permission'

const MessageRouter: Array<RouteRecordRaw> = [
  {
    path: '/message',
    component: Layout,
    redirect: '/profile/index',
    meta: {
      name: 'message',
      title: '消息管理',
      icon: 'message'
    },
    children: [
      {
        path: 'index',
        component: () => import(/* webpackChunkName: "message" */ '@/views/message/Index.vue'),
        name: 'center',
        meta: {
          icon: 'message',
          title: '消息中心',
          noCache: true
        }
      },
      {
        path: 'management',
        component: () => import(/* webpackChunkName: "message" */ '@/views/message/admin/Index.vue'),
        name: 'management',
        meta: {
          title: '消息管理',
          icon: 'message',
          noCache: true,
          permission: {
            resource: ResourceList.Message,
            action: ActionList.Write
          }
        }
      }
    ]
  }
]

export default MessageRouter
