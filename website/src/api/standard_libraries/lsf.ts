import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Lsf = {
  id?: number
  stack_thick: number | null
  stack_factor: number | null
  lsf_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export interface ListLsfQuery extends BaseQuery{
  field_eq_stack_thick?: number | null
  field_eq_stack_factor?: number | null
  field_eq_lsf_creator?: number
}
export function GetStdLsfLibraries(query:ListLsfQuery): AxiosPromise<Response<Lsf[]>> {
  return request({
    url: '/std/lsf/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdLsfLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lsf/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdLsfLibraries(lsf: Lsf): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lsf/' + lsf.id,
    method: 'PATCH',
    data: lsf
  })
}
export function CreateStdLsfLibraries(lsf: Lsf): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lsf',
    method: 'POST',
    data: lsf
  })
}
export function ExportStdlsfLibraries() {
  return request({
    url: '/std/lsf/export',
    method: 'POST',
    responseType: 'blob'
  })
}
