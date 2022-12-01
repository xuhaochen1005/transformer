import { BaseQuery, User } from '@/model/common'
export interface Project {
  id: number
  title: string
  content: string
  send_by: number
  send_by_user?: User
  send_to: number
  send_to_user?: User
  status: number
  created_at?: string
  updated_at?: string
}

export type ProjectQuery = BaseQuery & {
  self?: number
  field_eq_send_by?: number
  field_lk_title?: string
  field_lk_content?: string
  field_eq_status?: number | null
}

export enum ProjectStatus {
  UNREAD = 1,
  HASREAD = 2
}
