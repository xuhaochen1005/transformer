import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'

export type Ll = {
  id?: number
  loss_level: number | null
  regulate_way: string | null
  rated_capacity_min: number | null
  rated_capacity_max: number | null
  rated_high_vol_min: number | null
  rated_high_vol_max: number | null
  rated_low_vol_min: number | null
  rated_low_vol_max: number | null
  regulate_range_min: number | null
  regulate_range_max: number | null
  regulate_range_step: number | null
  temperature: number | null
  load_loss: number | null
  no_load_loss: number | null
  no_load_current: number | null
  short_circuit_imped: number | null
  ll_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}

export interface ListLlQuery extends BaseQuery {
  field_eq_loss_level?: number | null
  field_eq_regulate_way?: string| null
  field_lt_rated_capacity_min?: number | null
  field_ge_rated_capacity_max?: number | null
  field_lt_rated_high_vol_min?: number | null
  field_ge_rated_high_vol_max?: number | null
  field_lt_rated_low_vol_min?: number | null
  field_ge_rated_low_vol_max?: number | null
  field_eq_regulate_range_min?: number | null
  field_eq_regulate_range_max?: number | null
  field_eq_regulate_range_step?: number | null
  field_eq_temperature?: number | null
  field_eq_ll_creator?: number
}
export function GetStdLlLibraries(query:ListLlQuery): AxiosPromise<Response<Ll[]>> {
  return request({
    url: '/std/ll/',
    method: 'GET',
    params: query
  })
}

export function CreateStdLlLibraries(ll: Ll): AxiosPromise<Response<void>> {
  return request({
    url: '/std/ll',
    method: 'POST',
    data: ll
  })
}

export function UpdateStdLlLibraries(ll: Ll): AxiosPromise<Response<void>> {
  return request({
    url: '/std/ll/' + ll.id,
    method: 'PATCH',
    data: ll
  })
}

export function DeleteStdLlLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/ll/' + id,
    method: 'DELETE'
  })
}
export function ExportStdllLibraries() {
  return request({
    url: '/std/ll/export',
    method: 'POST',
    responseType: 'blob'
  })
}
