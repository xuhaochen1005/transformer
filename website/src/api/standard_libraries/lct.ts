import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery,User} from "@/model/common";
export type Lct = {
  id?: number
  cool_type: string
  lct_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLctQuery = BaseQuery & {
  field_eq_cool_type: string
  field_eq_lct_creator?: number
}
export function GetStdlctLibraries(query:ListLctQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lct/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlctLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lct/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlctLibraries(lct: Lct): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lct/' + lct.id,
    method: 'PATCH',
    data: lct
  })
}
export function CreateStdlctLibraries(lct: Lct): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lct',
    method: 'POST',
    data: lct
  })
}
export function ExportStdlctLibraries(){
  return request({
    url: '/std/lct/export',
    method: 'POST',
    responseType:'blob'
  })
}
