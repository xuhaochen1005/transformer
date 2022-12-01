import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Lcwt = {
  id?: number
  coil_wire_type: string
  coil_wire_type_sign: string
  lcwt_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export interface ListLcwtQuery extends BaseQuery {
  field_eq_coil_wire_type?: string
  field_eq_coil_wire_type_sign?: string
  field_eq_lcwt_creator?: number
}
export function GetStdlcwtLibraries(query:ListLcwtQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lcwt/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlcwtLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcwt/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlcwtLibraries(lcwt: Lcwt): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcwt/' + lcwt.id,
    method: 'PATCH',
    data: lcwt
  })
}
export function CreateStdlcwtLibraries(lcwt: Lcwt): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcwt',
    method: 'POST',
    data: lcwt
  })
}
export function ExportStdlcwtLibraries() {
  return request({
    url: '/std/lcwt/export',
    method: 'POST',
    responseType: 'blob'
  })
}
