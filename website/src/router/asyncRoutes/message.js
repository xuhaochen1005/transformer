import Layout from '@/layout/index.vue';
const MessageRouter = [
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
                    title: '消息中心',
                    //   icon: 'user',
                    noCache: true
                }
            },
            {
                path: 'start',
                component: () => import(/* webpackChunkName: "message" */ '@/views/message/Create.vue'),
                name: 'start',
                meta: {
                    title: '发起消息',
                    // icon: 'user',
                    noCache: true
                }
            },
            {
                path: 'management',
                component: () => import(/* webpackChunkName: "message" */ '@/views/message/admin/Index.vue'),
                name: 'management',
                meta: {
                    title: '消息管理',
                    // icon: 'user',
                    noCache: true
                }
            }
        ]
    }
];
export default MessageRouter;
//# sourceMappingURL=message.js.map