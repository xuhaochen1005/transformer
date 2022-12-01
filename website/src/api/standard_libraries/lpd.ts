import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Lpd = {
  id?: number
  phase_dist_min: number | null
  vol_min: number | null
  vol_max: number | null
  lpd_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLpdQuery = BaseQuery & {
  field_eq_phase_dist_min: number | null
  // field_eq_vol_min: number | null
  // field_eq_vol_max: number | null
  field_lt_vol_min?: number | null
  field_ge_vol_max?: number | null
  field_eq_lpd_creator?: number| null
}
export function GetStdLpdLibraries(query:ListLpdQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lpd/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdLpdLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lpd/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdLpdLibraries(lpd: Lpd): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lpd/' + lpd.id,
    method: 'PATCH',
    data: lpd
  })
}
export function CreateStdLpdLibraries(lpd: Lpd): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lpd',
    method: 'POST',
    data: lpd
  })
}
export function ExportStdlpdLibraries() {
  return request({
    url: '/std/lpd/export',
    method: 'POST',
    responseType: 'blob'
  })
}
