import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { User } from '@/model/common'
export type Lresist = {
  id?: number
  wire_material: string | null
  temp: number | null
  resistivity: number | null
  lresist_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type BaseQuery = {
  page: number
  limit: number
  order: string
  order_by: string
}
export type ListLresistQuery = BaseQuery & {
  field_eq_wire_material: string | null
  field_eq_temp: number | null
  field_eq_resistivity: number | null
  field_eq_lresist_creator?: number
}
export function GetStdlresistLibraries(query:ListLresistQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lresist/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlresistLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lresist/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlresistLibraries(lresist: Lresist): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lresist/' + lresist.id,
    method: 'PATCH',
    data: lresist
  })
}
export function CreateStdlresistLibraries(lresist: Lresist): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lresist',
    method: 'POST',
    data: lresist
  })
}
export function ExportStdlresistLibraries() {
  return request({
    url: '/std/lresist/export',
    method: 'POST',
    responseType: 'blob'
  })
}
