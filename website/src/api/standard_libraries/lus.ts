import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Lus = {
  id?: number
  usage: string | null
  usage_sign: string | null
  lus_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLusQuery = BaseQuery & {
  field_eq_usage: string | null
  field_eq_usage_sign: string | null
  field_eq_lus_creator?: number
}
export function GetStdLusLibraries(query:ListLusQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lus/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdLusLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lus/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdLusLibraries(lus: Lus): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lus/' + lus.id,
    method: 'PATCH',
    data: lus
  })
}
export function CreateStdLusLibraries(lus: Lus): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lus',
    method: 'POST',
    data: lus
  })
}
export function ExportStdlusLibraries() {
  return request({
    url: '/std/lus/export',
    method: 'POST',
    responseType: 'blob'
  })
}
