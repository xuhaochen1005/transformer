import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Lcc = {
  id?: number
  low_vol_coil_connect: string
  high_vol_coil_connect: string
  lcc_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}

export type ListLccQuery = BaseQuery & {
  field_eq_low_vol_coil_connect: string
  field_eq_high_vol_coil_connect: string
  field_eq_lcc_creator?: number
}
export function GetStdlccLibraries(query:ListLccQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lcc/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlccLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcc/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlccLibraries(lcc: Lcc): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcc/' + lcc.id,
    method: 'PATCH',
    data: lcc
  })
}
export function CreateStdlccLibraries(lcc: Lcc): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcc',
    method: 'POST',
    data: lcc
  })
}
export function ExportStdlccLibraries(){
  return request({
    url: '/std/lcc/export',
    method: 'POST',
    responseType:'blob'
  })
}
