import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { User } from '@/model/common'
export type Lifhv = {
  id?: number
  industry_freq_hold_vol: number | null
  rated_high_vol_min: number | null
  rated_high_vol_max: number | null
  lifhv_creator?: number
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
export interface ListLifhvQuery extends BaseQuery {
  id?: number
}
export function GetStdlifhvLibraries(query:ListLifhvQuery): AxiosPromise<Response<Lifhv[]>> {
  return request({
    url: '/std/lifhv/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlifhvLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lifhv/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlifhvLibraries(lifhv: Lifhv): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lifhv/' + lifhv.id,
    method: 'PATCH',
    data: lifhv
  })
}
export function CreateStdlifhvLibraries(lifhv: Lifhv): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lifhv',
    method: 'POST',
    data: lifhv
  })
}
export function ExportStdlifhvLibraries() {
  return request({
    url: '/std/lifhv/export',
    method: 'POST',
    responseType: 'blob'
  })
}
