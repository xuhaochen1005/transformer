import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Lrws = {
  id?: number
  std_diameter: number | null
  section_area: number | null
  max_diameter: number | null
  lrws_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLrwsQuery = BaseQuery & {
  field_eq_std_diameter: number | null
  field_eq_section_area: number | null
  field_eq_max_diameter: number | null
  field_eq_lrws_creator?: number
}
export function GetStdLrwsLibraries(query:ListLrwsQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lrws/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdLrwsLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lrws/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdLrwsLibraries(lrws: Lrws): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lrws/' + lrws.id,
    method: 'PATCH',
    data: lrws
  })
}
export function CreateStdLrwsLibraries(lrws: Lrws): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lrws',
    method: 'POST',
    data: lrws
  })
}
export function ExportStdlrwsLibraries(){
  return request({
    url: '/std/lrws/export',
    method: 'POST',
    responseType:'blob'
  })
}
