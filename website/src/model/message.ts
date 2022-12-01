import { BaseQuery, User } from '@/model/common'
export interface Message {
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

export type MessageQuery = BaseQuery & {
  self?: number
  field_eq_send_by?: number
  field_lk_title?: string
  field_lk_content?: string
  field_eq_status?: number | null
}

export enum MessageStatus {
  UNREAD = 1,
  HASREAD = 2
}

export interface MessageTemplate {
  id: number
  name: string
  title: string
  content: string
  create_by: number
  create_by_user?: User
  created_at?: string
  updated_at?: string
}

export type MessageTemplateQuery = BaseQuery & {
  field_lk_name: string
  field_lk_title: string
  field_lk_content: string
}
