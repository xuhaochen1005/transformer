import Layout from '@/layout/index.vue'
const DesignRouter = [
  {
    path: '/design',
    component: Layout,
    redirect: '/profile/index',
    meta: {
      name: 'user',
      title: '产品设计',
      icon: 'design'
    },
    children: [
      {
        path: 'design_projects',
        component: () => import(/* webpackChunkName: "design" */ '@/views/projects/admin/Index.vue'),
        name: 'design_projects',
        meta: {
          title: '任务书管理',
          //  icon: 'user',
          noCache: true
        }
      },
      {
        path: 'design_start',
        component: () => import(/* webpackChunkName: "design" */ '@/views/design/design-task.vue'),
        name: 'design_start',
        meta: {
          title: '发起设计',
          //   icon: 'user',
          noCache: true
        }
      },
      /*    {
              path: 'design_recommend',
              component: () => import(/!* webpackChunkName: "design" *!/ '@/views/profile/index.vue'),
              name: 'design_recommend',
              meta: {
                title: '方案推荐',
                //  icon: 'user',
                noCache: true
              }
            }, */
      {
        path: 'design_check',
        component: () => import(/* webpackChunkName: "design" */ '@/views/profile/index.vue'),
        name: 'design_check',
        meta: {
          title: '设计复核',
          //    icon: 'user',
          noCache: true
        }
      },
      {
        path: 'design_result',
        component: () => import(/* webpackChunkName: "profile" */ '@/views/profile/index.vue'),
        name: 'design_result',
        meta: {
          title: '设计结果',
          //   icon: 'user',
          noCache: true
        }
      }
    ]
  }
]
export default DesignRouter
// # sourceMappingURL=design.js.map
