import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
import lcs from '@/views/standard/libraries/lcs/index.vue'
export type Lcs = {
  id?: number
  coil_shape: string
  lcs_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
  deleted_at?: string
}

export interface ListLcsQuery extends BaseQuery {
  field_eq_coil_shape?: string
  field_eq_lcs_creator?: number
}
export function GetStdlcsLibraries(query:ListLcsQuery): AxiosPromise<Response<Lcs[]>> {
  return request({
    url: '/std/lcs/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlcsLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcs/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlcsLibraries(lcs: Lcs): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcs/' + lcs.id,
    method: 'PATCH',
    data: lcs
  })
}
export function CreateStdlcsLibraries(lcs: Lcs): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcs',
    method: 'POST',
    data: lcs
  })
}
export function ExportStdlcsLibraries() {
  return request({
    url: '/std/lcs/export',
    method: 'POST',
    responseType: 'blob'
  })
}
