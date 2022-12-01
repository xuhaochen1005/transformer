import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import router, { addPermissionRoutes } from '@/router'
import { RouteLocationNormalized } from 'vue-router'
import { useStore } from './store'
import { roleList, initRoles, initPermissions } from '@/utils/permission'

NProgress.configure({ showSpinner: false })

const store = useStore()

router.beforeEach(async(to: RouteLocationNormalized, _: RouteLocationNormalized, next: any) => {
  NProgress.start()
  if (store.state.user.token) {
    if (to.path === '/login') {
      next({ path: '/' })
    } else {
      if (roleList === null) {
        await initRoles()
        await initPermissions()
        addPermissionRoutes()
        next({ ...to, replace: true })
      } else {
        next()
      }
    }
  } else {
    if (to.path === '/login') {
      next()
    } else {
      next(`/login?redirect=${to.path}`)
    }
  }
  NProgress.done()
})
