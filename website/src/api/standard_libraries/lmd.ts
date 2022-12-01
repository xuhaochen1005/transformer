import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'

export type Lmd = {
  id?: number
  rated_capacity_min: number
  rated_capacity_max: number
  magnet_density_min: number
  magnet_density_max: number
  lmd_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLmdQuery = BaseQuery & {
  // field_eq_rated_capacity_min: number | null
  // field_eq_rated_capacity_max: number | null
  // field_eq_magnet_density_min: number | null
  // field_eq_magnet_density_max: number | null
  field_lt_rated_capacity_min?: number | null
  field_ge_rated_capacity_max?: number | null
  field_le_magnet_density_min?: number | null
  field_ge_magnet_density_max?: number | null
  field_eq_lmd_creator?: number

}
export function GetStdLmdLibraries(query:ListLmdQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lmd',
    method: 'GET',
    params: query
  })
}

export function CreateStdLmdLibraries(lmd: Lmd): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lmd',
    method: 'POST',
    data: lmd
  })
}

export function UpdateStdLmdLibraries(lmd: Lmd): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lmd/' + lmd.id,
    method: 'PATCH',
    data: lmd
  })
}

export function DeleteStdLmdLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lmd/' + id,
    method: 'DELETE'
  })
}
export function ExportStdlmdLibraries() {
  return request({
    url: '/std/lmd/export',
    method: 'POST',
    responseType: 'blob'
  })
}
