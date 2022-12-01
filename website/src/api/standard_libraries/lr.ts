import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { User } from '@/model/common'
export type Lr = {
  id?: number
  type: string | null
  density: number | null
  price: number | null
  lr_creator?: number
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
export interface ListLrQuery extends BaseQuery {
  field_eq_type: string | null | undefined
  field_eq_density: number | null | undefined
  field_eq_price: number | null | undefined
  field_eq_lr_creator?: number
}
export function GetStdlrLibraries(query:ListLrQuery): AxiosPromise<Response<Lr[]>> {
  return request({
    url: '/std/lr/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlrLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lr/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlrLibraries(lr: Lr): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lr/' + lr.id,
    method: 'PATCH',
    data: lr
  })
}
export function CreateStdlrLibraries(lr: Lr): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lr',
    method: 'POST',
    data: lr
  })
}
export function ExportStdlrLibraries() {
  return request({
    url: '/std/lr/export',
    method: 'POST',
    responseType: 'blob'
  })
}
