import { RouteRecordRaw } from 'vue-router'
import Layout from '@/layout/index.vue'
import { ActionList, ResourceList } from '@/model/permission'

const StandardRouter: Array<RouteRecordRaw> = [
  {
    path: '/standard',
    component: Layout,
    // redirect: '/standard/index',
    meta: {
      name: 'standard',
      title: '标准库管理',
      icon: 'standard',
      permission: {
        resource: ResourceList.Libraries,
        action: ActionList.Read
      }
    },
    children: [
      /* {
        path: 'search',
        component: () => import(/!* webpackChunkName: "standard" *!/ '@/views/standard/index.vue'),
        name: 'search',
        meta: {
          title: '标准查询',
          //  icon: 'user',
          noCache: true
        }
      }, */
      {
        path: 'libraries',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/Index.vue'),
        name: 'libraries',
        meta: {
          title: '标准库',
          icon: 'standard',
          noCache: false
        }
      },
      {
        path: 'calculation',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/calculation/index.vue'),
        name: 'calculation',
        meta: {
          title: '计算公式',
          icon: 'standard',
          noCache: true
        }
      },
      {
        path: 'ladbl',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/ladbl/index.vue'),
        name: 'ladbl'
      },
      {
        path: 'lcc',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lcc/index.vue'),
        name: 'lcc'
      },
      {
        path: 'lcfiy',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lcfiy/index.vue'),
        name: 'lcfiy'
      },
      {
        path: 'lcs',
        component: () => import(/* webpackChunkName: "standard" */'@/views/standard/libraries/lcs/index.vue'),
        name: 'lcs'
      },
      {
        path: 'lcwt',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lcwt/index.vue'),
        name: 'lcwt'
      },
      {
        path: 'lct',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lct/index.vue'),
        name: 'lct'
      },
      {
        path: 'lck',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lck/index.vue'),
        name: 'lck'
      },
      {
        path: 'lcd',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lcd/index.vue'),
        name: 'lcd'
      },
      {
        path: 'lelf',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lelf/index.vue'),
        name: 'lelf'
      },
      {
        path: 'lfws',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lfws/index.vue'),
        name: 'lfws'
      },
      {
        path: 'lflvml',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lflvml/index.vue'),
        name: 'lflvml'
      },
      {
        path: 'lfoml',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lfoml/index.vue'),
        name: 'lfoml'
      },
      // {
      //   path: 'lfp',
      //   component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lfp/index.vue'),
      //   name: 'lfp'
      // },
      {
        path: 'lhvoml',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lhvoml/index.vue'),
        name: 'lhvoml'
      },
      {
        path: 'lifhv',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lifhv/index.vue'),
        name: 'lifhv'
      },
      {
        path: 'licfih',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/licfih/index.vue'),
        name: 'licfih'
      },
      {
        path: 'lico',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lico/index.vue'),
        name: 'lico'
      },
      {
        path: 'licr',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/licr/index.vue'),
        name: 'licr'
      },
      {
        path: 'lli',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lli/index.vue'),
        name: 'lli'
      },
      {
        path: 'lls',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lls/index.vue'),
        name: 'lls'
      },
      {
        path: 'll',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/ll/index.vue'),
        name: 'll'
      },
      {
        path: 'llvoml',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/llvoml/index.vue'),
        name: 'llvoml'
      },
      {
        path: 'lmd',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lmd/index.vue'),
        name: 'lmd'
      },
      {
        path: 'lmad',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lmad/index.vue'),
        name: 'lmad'
      },
      {
        path: 'lnp',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lnp/index.vue'),
        name: 'lnp'
      },
      {
        path: 'lpd',
        component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lpd/index.vue'),
        name: 'lpd'
      },
      {
        path: 'lpn',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lpn/index.vue'),
        name: 'lpn'
      },
      {
        path: 'lr',
        component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lr/index.vue'),
        name: 'lr'
      },
      {
        path: 'lresist',
        component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lresist/index.vue'),
        name: 'lresist'
      },
      {
        path: 'lrf',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lrf/index.vue'),
        name: 'lrf'
      },
      {
        path: 'lrr',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lrr/index.vue'),
        name: 'lrr'
      },
      {
        path: 'lrw',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lrw/index.vue'),
        name: 'lrw'
      },
      {
        path: 'lrws',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lrws/index.vue'),
        name: 'lrws'
      },
      {
        path: 'lstack',
        component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lstack/index.vue'),
        name: 'lstack'
      },
      {
        path: 'lshv',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lshv/index.vue'),
        name: 'lshv'
      },
      {
        path: 'lsf',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lsf/index.vue'),
        name: 'lsf'
      },
      {
        path: 'lspd',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lspd/index.vue'),
        name: 'lspd'
      },
      {
        path: 'ltf',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/ltf/index.vue'),
        name: 'ltf'
      },
      {
        path: 'ltr',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/ltr/index.vue'),
        name: 'ltr'
      },
      {
        path: 'lti',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lti/index.vue'),
        name: 'lti'
      },
      {
        path: 'lus',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lus/index.vue'),
        name: 'lus'
      },
      {
        path: 'lwire',
        component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lwire/index.vue'),
        name: 'lwire'
      },
      {
        path: 'lwi',
        component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lwi/index.vue'),
        name: 'lwi'
      },
      {
        path: 'lwiml',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lwiml/index.vue'),
        name: 'lwiml'
      },
      {
        path: 'lwim',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lwim/index.vue'),
        name: 'lwim'
      },
      {
        path: 'lcadi',
        component: () => import(/* webpackChunkName: "standard" */ '@/views/standard/libraries/lcadi/index.vue'),
        name: 'lcadi'
      }
      // 测试
    ]
  }
]

export default StandardRouter
