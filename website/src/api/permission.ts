import { AxiosPromise } from 'axios'
import request from '@/utils/request'
import { Response } from '@/model'
import {
  Role,
  RolePermission,
  RolePermissionQuery,
  RoleQuery,
  UserRole,
  UserRoleQuery,
  UserRoles
} from '@/model/permission'

export function GetRoles(): AxiosPromise<Response<string[]>> {
  return request({
    url: '/user/roles',
    method: 'GET'
  })
}

export function GetRoleList(params: RoleQuery): AxiosPromise<Response<Role[]>> {
  return request({
    url: '/permission/roles',
    method: 'GET',
    params: params
  })
}

export function CreateRole(params: Role): AxiosPromise<Response<any>> {
  return request({
    url: '/permission/role',
    method: 'POST',
    data: params
  })
}

export function DeleteRole(id: number): AxiosPromise<Response<any>> {
  return request({
    url: '/permission/role/' + id,
    method: 'DELETE'
  })
}

export function UpdateRole(params: Role): AxiosPromise<Response<any>> {
  return request({
    url: '/permission/role/' + params.id,
    method: 'PATCH',
    data: params
  })
}

export function getPermissions(): AxiosPromise<Response<string[][]>> {
  return request({
    url: '/user/permissions',
    method: 'GET'
  })
}

export function getRolePermissions(params: RolePermissionQuery): AxiosPromise<Response<RolePermission[]>> {
  return request({
    url: '/permission/role_permission',
    method: 'GET',
    params: params
  })
}

export function GetUserRole(params: UserRoleQuery): AxiosPromise<Response<UserRoles[]>> {
  return request({
    url: '/permission/user_role',
    method: 'GET',
    params: params
  })
}

export function UpdateUserRole(params: UserRole): AxiosPromise<Response<any>> {
  return request({
    url: '/permission/user_role/' + params.id,
    method: 'PATCH',
    data: params
  })
}

export function CreateUserRole(params: UserRole): AxiosPromise<Response<any>> {
  return request({
    url: '/permission/user_role',
    method: 'POST',
    data: params
  })
}

export function DeleteUserRole(id : number): AxiosPromise<Response<any>> {
  return request({
    url: '/permission/user_role/' + id,
    method: 'DELETE'
  })
}

export function CreateRolePermissions(params: RolePermission[]): AxiosPromise<Response<RolePermission>> {
  return request({
    url: '/permission/role_permission',
    method: 'POST',
    data: {
      role_permissions: params
    }
  })
}

export function DeleteRolePermission(params: number[] | number) : AxiosPromise<Response<any>> {
  return request({
    url: '/permission/role_permission/' + [params].flatMap(i => i).join(','),
    method: 'DELETE'
  })
}
