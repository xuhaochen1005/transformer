import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'

export type Lnp = {
  id?: number
  rated_capacity_min: number
  rated_capacity_max: number
  rated_high_vol_min: number
  rated_high_vol_max: number
  cool_type: string | null
  noise_predict: number | null
  lnp_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
  deleted_at?: string
}
export interface ListLnpQuery extends BaseQuery{
  field_lt_rated_capacity_min?: number | null
  field_ge_rated_capacity_max?: number | null
  field_lt_rated_high_vol_min?: number | null
  field_ge_rated_high_vol_max?: number | null
  field_eq_cool_type?: string | null
  field_eq_noise_predict?: string | null
  field_eq_lnp_creator?: number
}

export function GetStdLnpLibraries(query: ListLnpQuery): AxiosPromise<Response<Lnp[]>> {
  return request({
    url: '/std/lnp',
    method: 'GET',
    params: query
  })
}

export function CreateStdLnpLibraries(lnp: Lnp): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lnp',
    method: 'POST',
    data: lnp
  })
}

export function UpdateStdLnpLibraries(lnp: Lnp): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lnp/' + lnp.id,
    method: 'PATCH',
    data: lnp
  })
}

export function DeleteStdLnpLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lnp/' + id,
    method: 'DELETE'
  })
}
export function ExportStdlnpLibraries() {
  return request({
    url: '/std/lnp/export',
    method: 'POST',
    responseType: 'blob'
  })
}
