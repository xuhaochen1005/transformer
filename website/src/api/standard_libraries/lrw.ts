import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User} from "@/model/common";
export type Lrw = {
  id?: number
  regulate_way: string | null
  regulate_way_sign: string | null
  lrw_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLrwQuery = BaseQuery & {
  field_eq_regulate_way: string | null
  field_eq_regulate_way_sign: string | null
  field_eq_lrw_creator?: number
}
export function GetStdLrwLibraries(query:ListLrwQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lrw/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdLrwLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lrw/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdLrwLibraries(lrw: Lrw): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lrw/' + lrw.id,
    method: 'PATCH',
    data: lrw
  })
}
export function CreateStdLrwLibraries(lrw: Lrw): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lrw',
    method: 'POST',
    data: lrw
  })
}
export function ExportStdlrwLibraries(){
  return request({
    url: '/std/lrw/export',
    method: 'POST',
    responseType:'blob'
  })
}
