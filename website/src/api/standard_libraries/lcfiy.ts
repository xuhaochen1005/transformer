import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Lcfiy = {
  id?: number
  rated_high_vol_min: number | null
  rated_high_vol_max: number | null
  rated_low_vol_min: number | null
  rated_low_vol_max: number | null
  low_vol_coil_from_iron_yoke: number | null
  high_vol_coil_from_iron_yoke: number | null
  lcfiy_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLcfiyQuery = BaseQuery & {
  field_lt_rated_high_vol_min?: number
  field_ge_rated_high_vol_max?: number
  field_lt_rated_low_vol_min?: number
  field_ge_rated_low_vol_max?: number
  field_eq_low_vol_coil_from_iron_yoke: number | null
  field_eq_high_vol_coil_from_iron_yoke: number | null
  field_eq_lcc_creator?: number
}
export function GetStdlcfiyLibraries(query:ListLcfiyQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lcfiy/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlcfiyLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcfiy/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlcfiyLibraries(lcfiy: Lcfiy): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcfiy/' + lcfiy.id,
    method: 'PATCH',
    data: lcfiy
  })
}
export function CreateStdlcfiyLibraries(lcfiy: Lcfiy): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcfiy',
    method: 'POST',
    data: lcfiy
  })
}
export function ExportStdlcfiyLibraries() {
  return request({
    url: '/std/lcfiy/export',
    method: 'POST',
    responseType: 'blob'
  })
}
