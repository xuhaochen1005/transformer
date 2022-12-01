import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
import Lstack from '@/views/standard/libraries/lstack/index.vue'
export type Lstack = {
  id?: number
  type: string | null
  density: number | null
  price: number | null
  lstack_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export interface ListLstackQuery extends BaseQuery{
  field_eq_type?: string | null
  field_eq_density?: number | null
  field_eq_price?: number | null
  field_eq_lstack_creator?: number
}
export function GetStdlstackLibraries(query:ListLstackQuery): AxiosPromise<Response<Lstack[]>> {
  return request({
    url: '/std/lstack/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlstackLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lstack/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlstackLibraries(lstack: Lstack): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lstack/' + lstack.id,
    method: 'PATCH',
    data: lstack
  })
}
export function CreateStdlstackLibraries(lstack: Lstack): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lstack',
    method: 'POST',
    data: lstack
  })
}
export function ExportStdlstackLibraries() {
  return request({
    url: '/std/lstack/export',
    method: 'POST',
    responseType: 'blob'
  })
}
