import { GetRoles, getPermissions } from '@/api/permission'
export let roleList = null
export async function initRoles() {
  const response = await GetRoles()
  roleList = new Map()
  if (response.status === 200) {
    response.data.spec.forEach(role => {
      roleList.set(role, true)
    })
  }
}
export function resetRoles() {
  roleList = null
}
export function checkRole(roles) {
  return roles.some(role => {
    return roleList.get(role)
  })
}
const permissionList = new Map()
export async function initPermissions() {
  const response = await getPermissions()
  if (response.status === 200 && response.data.spec) {
    response.data.spec.forEach(permission => {
      permissionList.set({ resource: permission[1], action: permission[2] }, true)
    })
  }
}
export function resetPermissions() {
  permissionList.clear()
}
export function checkPermission(resource, action) {
  if (roleList.get('root')) {
    return true
  }
  return permissionList.get({ resource, action })
}
// # sourceMappingURL=permission.js.map
