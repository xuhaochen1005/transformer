import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { User } from '@/model/common'
export type Lelf = {
  id?: number
  rated_capacity: number | null
  eddy_loss_factor: number | null
  lelf_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type BaseQuery = {
  page: number
  limit: number
  order: string
  order_by: string
}
export interface ListLelfQuery extends BaseQuery {
  field_lk_rated_capacity?: number | null
  field_lk_eddy_loss_factor?: number | null
  field_eq_lelf_creator?: number
  }
export function GetStdlelfLibraries(query:ListLelfQuery): AxiosPromise<Response<Lelf[]>> {
  return request({
    url: '/std/lelf/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlelfLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lelf/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlelfLibraries(lelf: Lelf): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lelf/' + lelf.id,
    method: 'PATCH',
    data: lelf
  })
}
export function CreateStdlelfLibraries(lelf: Lelf): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lelf',
    method: 'POST',
    data: lelf
  })
}
export function ExportStdlelfLibraries() {
  return request({
    url: '/std/lelf/export',
    method: 'POST',
    responseType: 'blob'
  })
}
