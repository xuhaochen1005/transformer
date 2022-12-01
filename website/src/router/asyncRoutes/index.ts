import { RouteRecordRaw } from 'vue-router'
// import approve from './approve'
import department from './department'

import maintenance from './maintenance'
import message from './message'
import PermissionRouter from './permission'

import standard from './standard'
import user from './user'
import DesignRouter from './design'
import documents from './document'

const asyncRoutes: Array<RouteRecordRaw> = [
  // ...approve,

  ...DesignRouter,
  // ...projects,
  ...standard,
  ...documents,

  ...message,
  ...PermissionRouter,
  ...department,
  ...user,
  ...maintenance
]

export default asyncRoutes
