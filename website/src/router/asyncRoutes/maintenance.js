import Layout from '@/layout/index.vue';
const MaintenanceRouter = [
    {
        path: '/maintainence',
        component: Layout,
        redirect: '/profile/index',
        meta: {
            // name: 'user',
            title: '系统维护',
            icon: 'maintenance'
        },
        children: [
            {
                path: 'index',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/maintenance/Index.vue'),
                name: 'index',
                meta: {
                    name: 'user',
                    title: '系统维护',
                    icon: 'maintenance'
                }
            }
        ]
    }
];
export default MaintenanceRouter;
//# sourceMappingURL=maintenance.js.map