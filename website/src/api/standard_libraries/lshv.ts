import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Lshv = {
  id?: number
  rated_high_vol_min: number | null
  rated_high_vol_max: number | null
  shock_hold_vol: number | null
  lshv_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export interface ListLshvQuery extends BaseQuery {
  id?: number
}
export function GetStdLshvLibraries(query:ListLshvQuery): AxiosPromise<Response<Lshv[]>> {
  return request({
    url: '/std/lshv/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdLshvLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lshv/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdLshvLibraries(lshv: Lshv): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lshv/' + lshv.id,
    method: 'PATCH',
    data: lshv
  })
}
export function CreateStdLshvLibraries(lshv: Lshv): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lshv',
    method: 'POST',
    data: lshv
  })
}
export function ExportStdlshvLibraries() {
  return request({
    url: '/std/lshv/export',
    method: 'POST',
    responseType: 'blob'
  })
}
