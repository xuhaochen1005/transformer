import Layout from '@/layout/index.vue';
const StandardRouter = [
    {
        path: '/standard',
        component: Layout,
        // redirect: '/profile/index',
        meta: {
            name: 'user',
            title: '标准库管理',
            icon: 'standard'
        },
        children: [
            /* {
              path: 'search',
              component: () => import(/!* webpackChunkName: "profile" *!/ '@/views/profile/index.vue'),
              name: 'search',
              meta: {
                title: '标准查询',
                //  icon: 'user',
                noCache: true
              }
            }, */
            {
                path: 'libraries',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/Index.vue'),
                name: 'libraries',
                meta: {
                    title: '标准库',
                    //  icon: 'user',
                    noCache: true
                }
            },
            {
                path: 'caculation',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/calculation/Index.vue'),
                name: 'caculation',
                meta: {
                    title: '计算公式',
                    //  icon: 'user',
                    noCache: true
                }
            },
            {
                path: 'files',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/profile/index.vue'),
                name: 'files',
                meta: {
                    title: '标准文件',
                    //   icon: 'user',
                    noCache: true
                }
            },
            {
                path: 'ladbl',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/ladbl/index.vue'),
                name: 'ladbl'
            },
            {
                path: 'lcc',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lcc/index.vue'),
                name: 'lcc'
            },
            {
                path: 'lcfiy',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lcfiy/index.vue'),
                name: 'lcfiy'
            },
            {
                path: 'lcs',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lcs/index.vue'),
                name: 'lcs'
            },
            {
                path: 'lcwt',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lcwt/index.vue'),
                name: 'lcwt'
            },
            {
                path: 'lct',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lct/index.vue'),
                name: 'lct'
            },
            {
                path: 'lck',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lck/index.vue'),
                name: 'lck'
            },
            {
                path: 'lcd',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lcd/index.vue'),
                name: 'lcd'
            },
            {
                path: 'lelf',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lelf/index.vue'),
                name: 'lelf'
            },
            {
                path: 'lfws',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lfws/index.vue'),
                name: 'lfws'
            },
            {
                path: 'lflvml',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lflvml/index.vue'),
                name: 'lflvml'
            },
            {
                path: 'lfoml',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lfoml/index.vue'),
                name: 'lfoml'
            },
            {
                path: 'lfp',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lfp/index.vue'),
                name: 'lfp'
            },
            {
                path: 'lhvoml',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lhvoml/index.vue'),
                name: 'lhvoml'
            },
            {
                path: 'lifhv',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lifhv/index.vue'),
                name: 'lifhv'
            },
            {
                path: 'licfih',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/licfih/index.vue'),
                name: 'licfih'
            },
            {
                path: 'lico',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lico/index.vue'),
                name: 'lico'
            },
            {
                path: 'licr',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/licr/index.vue'),
                name: 'licr'
            },
            {
                path: 'lli',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lli/index.vue'),
                name: 'lli'
            },
            {
                path: 'lll',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lcc/index.vue'),
                name: 'lll'
            },
            {
                path: 'll',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/ll/index.vue'),
                name: 'll'
            },
            {
                path: 'llvoml',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/llvoml/index.vue'),
                name: 'llvoml'
            },
            {
                path: 'lmd',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lmd/index.vue'),
                name: 'lmd'
            },
            {
                path: 'lmad',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lmad/index.vue'),
                name: 'lmad'
            },
            {
                path: 'lnp',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lnp/index.vue'),
                name: 'lnp'
            },
            {
                path: 'lpn',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lpn/index.vue'),
                name: 'lpn'
            },
            {
                path: 'lrf',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lrf/index.vue'),
                name: 'lrf'
            },
            {
                path: 'lrr',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lrr/index.vue'),
                name: 'lrr'
            },
            {
                path: 'lrw',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lrw/index.vue'),
                name: 'lrw'
            },
            {
                path: 'lrws',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lrws/index.vue'),
                name: 'lrws'
            },
            {
                path: 'lshv',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lshv/index.vue'),
                name: 'lshv'
            },
            {
                path: 'lsf',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lsf/index.vue'),
                name: 'lsf'
            },
            {
                path: 'lspd',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lspd/index.vue'),
                name: 'lspd'
            },
            {
                path: 'ltf',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/ltf/index.vue'),
                name: 'ltf'
            },
            {
                path: 'ltr',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/ltr/index.vue'),
                name: 'ltr'
            },
            {
                path: 'lti',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lti/index.vue'),
                name: 'lti'
            },
            {
                path: 'lus',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lus/index.vue'),
                name: 'lus'
            },
            {
                path: 'lwiml',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lwiml/index.vue'),
                name: 'lwiml'
            },
            {
                path: 'lwim',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lwim/index.vue'),
                name: 'lwim'
            },
            {
                path: 'lcadi',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/standard/libraries/lcadi/index.vue'),
                name: 'lcadi'
            }
            //
        ]
    }
];
export default StandardRouter;
//# sourceMappingURL=standard.js.map