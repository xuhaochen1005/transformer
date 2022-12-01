import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'

export type Lrr = {
  id?: number
  regulate_range_min: number
  regulate_range_max: number
  regulate_range_step: number
  lrr_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
  deleted_at?: string
}

export type ListLrrQuery = BaseQuery & {
  field_eq_regulate_range_min: number | null
  field_eq_regulate_range_max: number | null
  field_eq_regulate_range_step: number | null
  field_eq_lrr_creator?: number
}
export function GetStdLrrLibraries(query:ListLrrQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lrr/',
    method: 'GET',
    params: query
  })
}

export function CreateStdLrrLibraries(lrr: Lrr): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lrr/',
    method: 'POST',
    data: lrr
  })
}

export function UpdateStdLrrLibraries(lrr: Lrr): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lrr/' + lrr.id,
    method: 'PATCH',
    data: lrr
  })
}

export function DeleteStdLrrLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lrr/' + id,
    method: 'DELETE'
  })
}
export function ExportStdlrrLibraries() {
  return request({
    url: '/std/lrr/export',
    method: 'POST',
    responseType: 'blob'
  })
}
