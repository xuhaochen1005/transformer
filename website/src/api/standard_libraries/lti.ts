import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'

export type Lti = {
  id?: number
  rated_high_vol_min: number
  rated_high_vol_max: number
  rated_low_vol_min: number
  rated_low_vol_max: number
  terminal_insulate: number
  lti_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLtiQuery = BaseQuery & {
  // field_eq_rated_high_vol_min: number | null
  // field_eq_rated_high_vol_max: number | null
  // field_eq_rated_low_vol_min: number | null
  // field_eq_rated_low_vol_max: number | null
  field_lt_rated_high_vol_min?: number | null
  field_ge_rated_high_vol_max?: number | null
  field_lt_rated_low_vol_min?: number | null
  field_ge_rated_low_vol_max?: number | null
  field_eq_terminal_insulate: number | null
  field_eq_lti_creator?: number
}
export function GetStdLtiLibraries(query:ListLtiQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lti',
    method: 'GET',
    params: query
  })
}

export function CreateStdLtiLibraries(lti: Lti): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lti',
    method: 'POST',
    data: lti
  })
}

export function UpdateStdLtiLibraries(lti: Lti): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lti/' + lti.id,
    method: 'PATCH',
    data: lti
  })
}

export function DeleteStdLtiLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lti/' + id,
    method: 'DELETE'
  })
}
export function ExportStdltiLibraries() {
  return request({
    url: '/std/lti/export',
    method: 'POST',
    responseType: 'blob'
  })
}
