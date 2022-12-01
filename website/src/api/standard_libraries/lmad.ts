import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { User } from '@/model/common'

export type Lmad = {
  id?: number
  rated_high_vol_min: number
  rated_high_vol_max: number
  main_air_duct_min: number
  lmad_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
  deleted_at?: string
}

export type BaseQuery = {
  page: number
  limit: number
  order: string
  order_by: string
}

export type ListLmadQuery = BaseQuery & {
  // field_eq_rated_high_vol_min: number | null
  // field_eq_rated_high_vol_max: number | null
  field_lt_rated_high_vol_min?: number | null
  field_ge_rated_high_vol_max?: number | null
  field_eq_main_air_duct_min: number | null
  field_eq_lcc_creator?: number
}
export function GetStdLmadLibraries(query:ListLmadQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lmad',
    method: 'GET',
    params: query
  })
}

export function CreateStdLmadLibraries(lmad: Lmad): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lmad',
    method: 'POST',
    data: lmad
  })
}

export function UpdateStdLmadLibraries(lmad: Lmad): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lmad/' + lmad.id,
    method: 'PATCH',
    data: lmad
  })
}

export function DeleteStdLmadLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lmad/' + id,
    method: 'DELETE'
  })
}
export function ExportStdlmadLibraries() {
  return request({
    url: '/std/lmad/export',
    method: 'POST',
    responseType: 'blob'
  })
}
