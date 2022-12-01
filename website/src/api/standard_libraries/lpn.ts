import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Lpn = {
  id?: number
  phase: number
  phase_num: string | null
  phase_num_sign: string | null
  lpn_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLpnQuery = BaseQuery & {
  field_eq_phase_num: string | null
  field_eq_phase_num_sign: string | null
  field_eq_lpn_creator?: number
}
export function GetStdLpnLibraries(query:ListLpnQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lpn/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdLpnLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lpn/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdLpnLibraries(lpn: Lpn): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lpn/' + lpn.id,
    method: 'PATCH',
    data: lpn
  })
}
export function CreateStdLpnLibraries(lpn: Lpn): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lpn',
    method: 'POST',
    data: lpn
  })
}
export function ExportStdlpnLibraries() {
  return request({
    url: '/std/lpn/export',
    method: 'POST',
    responseType: 'blob'
  })
}
