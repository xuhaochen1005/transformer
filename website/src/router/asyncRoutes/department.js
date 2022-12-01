import Layout from '@/layout/index.vue';
const DepartmentRouter = [
    {
        path: '/department',
        component: Layout,
        // redirect: '/department/index',
        meta: {
            name: 'department',
            title: '产品设计',
            icon: 'department'
        },
        children: [
            {
                path: 'index',
                component: () => import(/* webpackChunkName: "profile" */ '@/views/department/admin/Index.vue'),
                name: 'department',
                meta: {
                    name: 'department',
                    title: '部门管理',
                    icon: 'department'
                }
            }
        ]
    }
];
export default DepartmentRouter;
//# sourceMappingURL=department.js.map