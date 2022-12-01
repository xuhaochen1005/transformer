import { RouteRecordRaw } from 'vue-router'
import Layout from '@/layout/index.vue'
import { ActionList, ResourceList } from '@/model/permission'

const DocumentsRouter: Array<RouteRecordRaw> = [
  {
    path: '/documents',
    component: Layout,
    meta: {
      name: 'Documents',
      title: '文档管理',
      icon: 'documents',
      permission: {
        action: ActionList.Read, resource: ResourceList.Documents
      }
    },
    children: [
      {
        path: 'design',
        component: () => import(/* webpackChunkName: "document" */ '@/views/documents/design.vue'),
        name: 'documents-design',
        meta: {
          title: '设计文档',
          icon: 'documents',
          permission: {
            action: ActionList.Read, resource: ResourceList.Documents
          }
        }
      },
      {
        path: 'training',
        component: () => import(/* webpackChunkName: "document" */ '@/views/documents/training.vue'),
        name: 'documents-training',
        meta: {
          title: '培训文档',
          icon: 'documents'
        }
      },
      {
        path: 'system',
        component: () => import(/* webpackChunkName: "document" */ '@/views/documents/system.vue'),
        name: 'documents-system',
        meta: {
          title: '系统文档',
          icon: 'documents'
        }
      },
      {
        path: 'feedback',
        component: () => import(/* webpackChunkName: "document" */ '@/views/documents/feedback.vue'),
        name: 'documents-feedback',
        meta: {
          title: '反馈意见',
          icon: 'documents'
        }
      }
    ]
  }
]

export default DocumentsRouter
