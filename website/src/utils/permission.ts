import { GetRoles, getPermissions } from '@/api/permission'
import { Permission, PermissionType } from '@/model/permission'

export let roleList: Map<string, boolean> | null = null

export async function initRoles(): Promise<void> {
  const response = await GetRoles()
  roleList = new Map<string, boolean>()
  if (response.status === 200) {
    response.data.spec.forEach(role => {
      roleList!.set(role, true)
    })
  }
}

export function resetRoles(): void {
  roleList = null
}

export function checkRole(roles: string[]): boolean | undefined {
  return roles.some(role => {
    return roleList!.get(role)
  })
}

export function getPermissionTypeString(permission: Permission): PermissionType {
  return `${permission.resource}_${permission.action}`
}

const permissionList = new Map<PermissionType, Permission>()

export async function initPermissions(): Promise<void> {
  const response = await getPermissions()
  if (response.status === 200 && response.data.spec) {
    response.data.spec.forEach(permission => {
      const permissionObj = { resource: permission[1], action: permission[2] } as Permission
      permissionList.set(getPermissionTypeString(permissionObj), permissionObj)
    })
  }
}

export function resetPermissions(): void {
  permissionList.clear()
}

export function checkPermission(resource: string, action: string): boolean | undefined {
  if (roleList!.get('root')) {
    return true
  }
  return !!permissionList.get(getPermissionTypeString({ resource, action } as Permission))
}
