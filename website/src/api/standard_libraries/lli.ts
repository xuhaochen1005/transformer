import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Lli = {
  id?: number
  winding_type: string
  rated_low_vol_min: number | null
  rated_low_vol_max: number | null
  layer_vol_min: number | null
  layer_vol_max: number | null
  layer_insulate: number | null
  lli_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLliQuery = BaseQuery & {
  field_eq_winding_type: string
  // field_eq_rated_low_vol_min: number | null
  // field_eq_rated_low_vol_max: number | null
  // field_eq_layer_vol_min: number | null
  // field_eq_layer_vol_max: number | null
  field_lt_rated_low_vol_min?: number | null
  field_ge_rated_low_vol_max?: number | null
  field_lt_layer_vol_min?: number | null
  field_ge_layer_vol_max?: number | null
  field_eq_layer_insulate: number | null
  field_eq_lli_creator?: number
  }
export function GetStdlliLibraries(query:ListLliQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lli/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlliLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lli/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlliLibraries(lli: Lli): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lli/' + lli.id,
    method: 'PATCH',
    data: lli
  })
}
export function CreateStdlliLibraries(lli: Lli): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lli',
    method: 'POST',
    data: lli
  })
}
export function ExportStdlliLibraries() {
  return request({
    url: '/std/lli/export',
    method: 'POST',
    responseType: 'blob'
  })
}
