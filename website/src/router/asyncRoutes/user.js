import Layout from '@/layout/index.vue';
const UserRouter = [
    {
        path: '/user',
        component: Layout,
        // redirect: '/profile/index',
        meta: {
            name: 'user',
            title: '用户中心',
            icon: 'user'
        },
        children: [
            {
                path: 'profile',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/user/profile.vue'),
                name: 'profile',
                meta: {
                    title: '个人信息',
                    icon: 'user',
                    noCache: true
                }
            },
            {
                path: 'manage',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/user/admin/Index.vue'),
                name: 'manage',
                meta: {
                    title: '用户管理',
                    roles: ['root'],
                    icon: 'user',
                    noCache: true
                }
            }
        ]
    }
];
export default UserRouter;
//# sourceMappingURL=user.js.map