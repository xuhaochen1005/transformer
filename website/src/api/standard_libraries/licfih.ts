import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { User } from '@/model/common'
export type Licfih = {
  id?: number
  inner_coil_form_iron_heart_min: number | null
  rated_high_vol_min: number | null
  rated_high_vol_max: number | null
  licfih_creator?: number
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
export type ListLicfihQuery = BaseQuery & {
  field_eq_inner_coil_form_iron_heart_min: number | null
  // field_eq_rated_high_vol_min: number | null
  // field_eq_rated_high_vol_max: number | null
  field_lt_rated_high_vol_min?: number
  field_ge_rated_high_vol_max?: number
  field_eq_licfih_creator?: number
}
export function GetStdlicfihLibraries(query:ListLicfihQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/licfih/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlicfihLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/licfih/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlicfihLibraries(licfih: Licfih): AxiosPromise<Response<void>> {
  return request({
    url: '/std/licfih/' + licfih.id,
    method: 'PATCH',
    data: licfih
  })
}
export function CreateStdlicfihLibraries(licfih: Licfih): AxiosPromise<Response<void>> {
  return request({
    url: '/std/licfih',
    method: 'POST',
    data: licfih
  })
}
export function ExportStdlicfihLibraries() {
  return request({
    url: '/std/licfih/export',
    method: 'POST',
    responseType: 'blob'
  })
}
