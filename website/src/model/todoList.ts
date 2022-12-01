import { BaseQuery } from '@/model/common'

export interface TodoList {
  id: number
  title: string
  category: string
  created_by?: number
  checked_by?: number
  checked_at?: string
  created_at?: string
  updated_at?: string
  deleted_at?: string
}

export type TodoListQuery = BaseQuery & {
  field_eq_category?: string
}

export enum TodoListCategory {
  FINISHED = '已完成',
  UN_FINISHED = '待完成',
  ALL = '全部'
}
