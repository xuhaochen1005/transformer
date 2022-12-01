import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Lwire = {
  id?: number
  wire_type: string | null
  wire_price?: number
  wire_material: string | null
  wire_shape?: string | null
  wire_density?: number | null
  lwire_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export interface ListLwireQuery extends BaseQuery {
  field_eq_wire_type?: string | null
  field_eq_wire_material?: string | null
  field_eq_wire_shape?: string | null
  field_eq_wire_density?: number | null
  field_eq_lwire_creator?: number
}
export function GetStdlwireLibraries(query:ListLwireQuery): AxiosPromise<Response<Lwire[]>> {
  return request({
    url: '/std/lwire/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlwireLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lwire/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlwireLibraries(lwire: Lwire): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lwire/' + lwire.id,
    method: 'PATCH',
    data: lwire
  })
}
export function CreateStdlwireLibraries(lwire: Lwire): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lwire',
    method: 'POST',
    data: lwire
  })
}
export function ExportStdlwireLibraries() {
  return request({
    url: '/std/lwire/export',
    method: 'POST',
    responseType: 'blob'
  })
}
