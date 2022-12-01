import { Department } from '@/model/department'
import { RoleList } from './permission'

export interface BaseQuery {
  page: number
  limit: number
  order: string
  order_by: string
}


export type UserQuery = BaseQuery & {
  role_name?: string
  roles?: RoleList[]
  'field_eq_user.id'?: number
  'field_lk_user.user_name_zh'?: string
  field_eq_name: string
  'field_lk_user.name'?: string
  field_eq_department_id: number | null
  'field_eq_user.status': number | null
}

export type User = {
  id?: number
  user_name_zh: string
  name: string
  password?: string
  repeat_password?: string
  department_id?: number
  department_name?: string
  department?: Department
  email: string
  telephone: string
  status: number
  info?: string
  created_at?: string
  updated_at?: string
  token?: string
}

export interface ListQuery {
  page: number,
  limit: number,
  importance: undefined,
  title: undefined,
  type: undefined,
  sort: string
}
