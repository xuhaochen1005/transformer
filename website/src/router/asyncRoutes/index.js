// import approve from './approve'
import department from './department';
import maintenance from './maintenance';
import message from './message';
import PermissionRouter from './permission';
// import projects from './projects'
import standard from './standard';
import user from './user';
// import document from '@/router/asyncRoutes/document'
import DesignRouter from './design';
const asyncRoutes = [
    // ...approve,
    ...DesignRouter,
    // ...projects,
    ...standard,
    // ...document,
    ...message,
    ...PermissionRouter,
    ...department,
    ...user,
    ...maintenance
];
export default asyncRoutes;
//# sourceMappingURL=index.js.map