import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { User } from '@/model/common'
export type Lwi = {
  id?: number
  wire_id: number | null
  width_min: number | null
  width_max: number | null
  axial_insulation: number | null
  radial_insulation: number | null
  lwi_creator?: number
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
export type ListLwiQuery = BaseQuery & {
  field_eq_wire_id: number | null
  field_lt_width_min?: number
  field_ge_width_max?: number
  // field_eq_width_min: number | null
  // field_eq_width_max: number | null
  field_eq_axial_insulation: number | null
  field_eq_radial_insulation: number | null
  field_eq_lwi_creator?: number
}
export function GetStdlwiLibraries(query:ListLwiQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lwi/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlwiLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lwi/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlwiLibraries(lwi: Lwi): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lwi/' + lwi.id,
    method: 'PATCH',
    data: lwi
  })
}
export function CreateStdlwiLibraries(lwi: Lwi): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lwi',
    method: 'POST',
    data: lwi
  })
}
export function ExportStdlwiLibraries() {
  return request({
    url: '/std/lwi/export',
    method: 'POST',
    responseType: 'blob'
  })
}
