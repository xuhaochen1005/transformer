import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User} from "@/model/common";
export type Lflvml = {
  id?: number
  mod_size: string
  platform_height: number | null
  mod_height: number | null
  mod_pic: string
  mod_num: string
  mod_amount: number | null
  mod_suit: string
  remark: string
  lflvml_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLflvmlQuery = BaseQuery & {
  field_eq_mod_size: string
  field_eq_platform_height: number | null
  field_eq_mod_height: number | null
  field_eq_mod_pic: string
  field_eq_mod_num: string
  field_eq_mod_amount: number | null
  field_eq_mod_suit: string
  field_eq_remark: string
  field_eq_lflvml_creator?: number
}
export function GetStdlflvmlLibraries(query:ListLflvmlQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lflvml/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlflvmlLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lflvml/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlflvmlLibraries(lflvml: Lflvml): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lflvml/' + lflvml.id,
    method: 'PATCH',
    data: lflvml
  })
}
export function CreateStdlflvmlLibraries(lflvml: Lflvml): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lflvml',
    method: 'POST',
    data: lflvml
  })
}
export function ExportStdlflvmlLibraries(){
  return request({
    url: '/std/lflvml/export',
    method: 'POST',
    responseType:'blob'
  })
}
