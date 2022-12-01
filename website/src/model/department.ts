import type { BaseQuery } from '@/model/common'

export interface Department {
  id?: number
  name: string
  status: number
  created_at?: string
  updated_at?: string
}

export type DepartmentQuery = BaseQuery & {
  field_lk_name: string,
  field_eq_status?: number
}

export enum DepartmentStatus {
  Normal = 1,
  Disabled = 2
}
