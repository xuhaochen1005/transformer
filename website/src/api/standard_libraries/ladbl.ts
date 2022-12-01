import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Ladbl = {
  id?: number
  air_duct_bar_thickness: number | null
  air_duct_bar_width: number | null
  air_duct_bar_length: number | null
  air_duct_bar_num: number
  ladbl_creator: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLadblQuery = BaseQuery & {
  field_lk_air_duct_bar_thickness: number | null
  field_lk_air_duct_bar_width: number | null
  field_lk_air_duct_bar_length: number | null
  field_lk_air_duct_bar_num: number | null
  field_eq_ladbl_creator?: number
}
export function GetStdladblLibraries(query:ListLadblQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/ladbl/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdladblLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/ladbl/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdladblLibraries(ladbl: Ladbl): AxiosPromise<Response<void>> {
  return request({
    url: '/std/ladbl/' + ladbl.id,
    method: 'PATCH',
    data: ladbl
  })
}
export function CreateStdladblLibraries(ladbl: Ladbl): AxiosPromise<Response<void>> {
  return request({
    url: '/std/ladbl',
    method: 'POST',
    data: ladbl
  })
}
export function ExportStdladblLibraries() {
  return request({
    url: '/std/ladblexport',
    method: 'POST',
    responseType: 'blob'
  })
}
