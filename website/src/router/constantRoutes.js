import Layout from '@/layout/index.vue';
const constantRoutes = [
    {
        path: '/',
        component: Layout,
        redirect: '/dashboard',
        children: [
            {
                path: 'dashboard',
                component: () => import(/* webpackChunkName: "dashboard" */ '@/views/dashboard/index.vue'),
                name: '主页',
                meta: {
                    title: '首页',
                    icon: 'user',
                    affix: true
                }
            }
        ]
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import(/* webpackChunkName: "userManager" */ '@/views/user/login.vue')
    },
    {
        path: '/redirect',
        component: Layout,
        meta: { hidden: true },
        children: [
            {
                path: '/redirect/:path(.*)',
                component: () => import(/* webpackChunkName: "redirect" */ '@/views/redirect/index.vue')
            }
        ]
    }
];
export default constantRoutes;
//# sourceMappingURL=constantRoutes.js.map