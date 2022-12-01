import Layout from '@/layout/index.vue';
const ApproveRouter = [
    {
        path: '/approve',
        component: Layout,
        // redirect: '/approve/index',
        meta: {
            name: 'approve',
            title: '审批管理',
            icon: 'user'
        },
        children: [
            {
                path: 'start_approve',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/approve/index.vue'),
                name: 'start_approve',
                meta: {
                    title: '发起审批',
                    icon: 'user'
                    // noCache: true
                }
            },
            {
                path: 'manager_approve',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/approve/admin/Index.vue'),
                name: 'manager_approve',
                meta: {
                    title: '审批管理',
                    icon: 'user',
                    noCache: true
                }
            }
        ]
    }
];
export default ApproveRouter;
//# sourceMappingURL=approve.js.map