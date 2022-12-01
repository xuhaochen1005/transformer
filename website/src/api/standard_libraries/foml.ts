import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User} from "@/model/common";
export type Lfoml = {
  id?: number
  mod_size: string
  mod_height: number | null
  mod_pic: string
  mod_num: string
  mod_amount: number | null
  mod_suit: string
  remark: string
  lfoml_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLfomlQuery = BaseQuery & {
  field_eq_mod_size: string
  field_eq_mod_height: number | null
  field_eq_mod_pic: string
  field_eq_mod_num: string
  field_eq_mod_amount: number | null
  field_eq_mod_suit: string
  field_eq_remark: string
  field_eq_lfoml_creator?: number
}
export function GetStdlfomlLibraries(query:ListLfomlQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lfoml/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlfomlLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lfoml/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlfomlLibraries(lfoml: Lfoml): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lflvml/' + lfoml.id,
    method: 'PATCH',
    data: lfoml
  })
}
export function CreateStdlfomlLibraries(lfoml: Lfoml): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lfoml',
    method: 'POST',
    data: lfoml
  })
}
export function ExportStdlfomlLibraries(){
  return request({
    url: '/std/lfoml/export',
    method: 'POST',
    responseType:'blob'
  })
}
