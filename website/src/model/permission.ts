import { BaseQuery } from '@/model/common'

export enum RoleList {
  Root = 'root',
  Demander = 'demander',
  Designer = 'designer',
  Reviewer = 'reviewer',
  ChiefEngineer = 'chief_engineer'
}

export enum ResourceList {
  User = 'user', // 用户
  Department = 'department', // 部门
  Documents = 'documents', // 文档
  PermissionRule = 'permission_rule', // 权限配置
  DesignProject = 'design_project', // 任务书管理
  DesignTask = 'design_task', // 设计任务
  Message = 'message', // 消息
  Libraries = 'libraries', // 标准库
  Dashboard = 'dashboard', // 首页
  PriceManagement = 'price_management' // 价格管理
}

export enum ActionList {
  Read = 'read',
  Write = 'write',
  Cancel = 'cancel',
  Design = 'design',
  Review = 'review',
  Approve = 'approve'
}

export interface Permission {
  resource: string
  action: string
}

export type PermissionType = String

export interface UserRole {
  id?: number,
  username: string,
  role_name: string,
  temporary?: number
}

export interface UserRoles {
  id?: number
  username: string
  user_name_zh: string
  user_role: UserRole[]
}

export type UserRoleQuery = BaseQuery & {
  field_eq_temporary?: number
  field_lk_username: string
  'field_lk_user.user_name_zh'?: string
  field_lk_role_name: string
  field_eq_role_name?: string
}

export interface Role {
  id?: number,
  name: string,
  cn_name: string,
  description: string
}

export type RoleQuery = BaseQuery & {
  username?: string
  field_eq_name?: string
  field_lk_cn_name?: string
}

export interface RolePermission {
  id: number,
  role: string,
  resource: string,
  action: string
}

export type RolePermissionQuery = BaseQuery & {
  field_eq_role: string
  field_eq_resource: string
  field_eq_action: string
}
