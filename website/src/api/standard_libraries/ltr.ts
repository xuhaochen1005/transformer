import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Ltr = {
  id?: number
  temp: number
  temp_sign: string
  low_vol_temp_rise: number
  high_vol_temp_rise: number
  ltr_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export interface ListLtrQuery extends BaseQuery{
  field_eq_temp?: number | null
  field_eq_temp_sign?: string | null
  field_eq_low_vol_temp_rise?: number | null
  field_eq_high_vol_temp_rise?: number | null
  field_eq_ltr_creator?: number
}
export function GetStdLtrLibraries(query:ListLtrQuery): AxiosPromise<Response<Ltr[]>> {
  return request({
    url: '/std/ltr/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdLtrLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/ltr/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdLtrLibraries(ltr: Ltr): AxiosPromise<Response<void>> {
  return request({
    url: '/std/ltr/' + ltr.id,
    method: 'PATCH',
    data: ltr
  })
}
export function CreateStdLtrLibraries(ltr: Ltr): AxiosPromise<Response<void>> {
  return request({
    url: '/std/ltr',
    method: 'POST',
    data: ltr
  })
}
export function ExportStdltrLibraries() {
  return request({
    url: '/std/ltr/export',
    method: 'POST',
    responseType: 'blob'
  })
}
