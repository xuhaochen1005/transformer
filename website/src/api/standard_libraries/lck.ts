import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Lck = {
  id?: number
  winding_material: string
  rated_capacity_min: number | null
  rated_capacity_max: number | null
  k_min: number | null
  k_max: number | null
  lck_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLckQuery = BaseQuery & {
  field_eq_winding_material: string
  // field_eq_rated_capacity_min: number | null
  // field_eq_rated_capacity_max: number | null
  // field_eq_k_min: number | null
  // field_eq_k_max: number | null
  field_lt_rated_capacity_min?: number,
  field_ge_rated_capacity_max?: number,
  field_lt_k_min?: number,
  field_ge_k_max?: number,
  field_eq_lck_creator?: number
  }
export function GetStdlckLibraries(query:ListLckQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lck/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlckLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lck/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlckLibraries(lck: Lck): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lck/' + lck.id,
    method: 'PATCH',
    data: lck
  })
}
export function CreateStdlckLibraries(lck: Lck): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lck',
    method: 'POST',
    data: lck
  })
}
export function ExportStdlckLibraries() {
  return request({
    url: '/std/lck/export',
    method: 'POST',
    responseType: 'blob'
  })
}
