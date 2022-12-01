import { createRouter, createWebHashHistory } from 'vue-router'
import constantRoutes from './constantRoutes'
import asyncRoutes from './asyncRoutes'
import { checkPermission, checkRole } from '@/utils/permission'
const router = createRouter({
  history: createWebHashHistory(),
  routes: constantRoutes
})
export function resetRouter() {
  const newRouter = createRouter({
    history: createWebHashHistory(),
    routes: constantRoutes
  })
  router.matcher = newRouter.matcher
}
let permissionRoutes
function filterAsyncRoutes(routes) {
  const result = []
  routes.forEach(route => {
    const r = { ...route }
    if (r.meta === undefined || (r.meta.roles === undefined && r.meta.permission === undefined) ||
            ((r.meta.roles === undefined || checkRole(r.meta.roles)) &&
                (r.meta.permission === undefined || checkPermission(r.meta.permission.resource, r.meta.permission.action)))) {
      if (r.children) {
        r.children = filterAsyncRoutes(r.children)
      }
      result.push(r)
    }
  })
  return result
}
export function addPermissionRoutes() {
  permissionRoutes = filterAsyncRoutes(asyncRoutes)
  permissionRoutes.forEach(route => {
    router.addRoute(route)
  })
}
export default router
export { permissionRoutes }
// # sourceMappingURL=index.js.map
