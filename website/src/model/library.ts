import { BaseQuery, User } from './common'
export interface Library {
  id?: number,
  libraries_name: string,
  libraries_short: string,
  libraries_description: string,
  libraries_creator?: number,
  creator?: User,
  libraries_source: string,
  created_at?: string,
  updated_at?: string,
  libraries_status: number
}
export interface LibraryQuery extends BaseQuery{
  field_lk_libraries_name?: string
  field_lk_libraries_description?: string
  field_lk_libraries_source?: string
  field_eq_libraries_status?: number
  field_eq_libraries_creator?: number
}

export enum LibraryStatus {
  Normal = 1,
  Disabled = 2
}
