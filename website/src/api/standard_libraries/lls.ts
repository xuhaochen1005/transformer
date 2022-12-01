import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'

export type Lls = {
  id?: number
  rated_capacity: number
  lead_loss: number
  lls_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
  deleted_at?: string
}

export interface ListLlsQuery extends BaseQuery{
  field_eq_rated_capacity?: number | null
  field_eq_lead_loss?: number | null
  field_eq_lls_creator?: number
}
export function GetStdLlsLibraries(query:ListLlsQuery): AxiosPromise<Response<Lls[]>> {
  return request({
    url: '/std/lls',
    method: 'GET',
    params: query
  })
}

export function CreateStdLlsLibraries(lls: Lls): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lls',
    method: 'POST',
    data: lls
  })
}

export function UpdateStdLlsLibraries(lls: Lls): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lls/' + lls.id,
    method: 'PATCH',
    data: lls
  })
}

export function DeleteStdLlsLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lls/' + id,
    method: 'DELETE'
  })
}
export function ExportStdllsLibraries() {
  return request({
    url: '/std/lls/export',
    method: 'POST',
    responseType: 'blob'
  })
}
