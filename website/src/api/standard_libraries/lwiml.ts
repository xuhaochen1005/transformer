import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User} from "@/model/common";

export type Lwiml = {
  id?: number
  mod_diameter: number
  mod_size: string
  mod_amount: number
  mod_pic: string
  mod_num: string| null
  mod_suit: string| null
  remark: string| null
  lwiml_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLwimlQuery = BaseQuery & {
  field_eq_mod_diameter: number | null
  field_eq_mod_size: string | null
  field_eq_mod_amount: number | null
  field_eq_mod_pic: string | null
  field_eq_mod_num: string | null
  field_eq_mod_suit: string | null
  field_eq_remark: string | null
  field_eq_lwiml_creator?: number
}
export function GetStdLwimlLibraries(query:ListLwimlQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lwiml',
    method: 'GET',
    params: query
  })
}

export function CreateStdLwimlLibraries(lwiml: Lwiml): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lwiml',
    method: 'POST',
    data: lwiml
  })
}

export function UpdateStdLwimlLibraries(lwiml: Lwiml): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lwiml/' + lwiml.id,
    method: 'PATCH',
    data: lwiml
  })
}

export function DeleteStdLwimlLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lwiml/' + id,
    method: 'DELETE'
  })
}
export function ExportStdlwimlLibraries() {
  return request({
    url: '/std/lwiml/export',
    method: 'POST',
    responseType: 'blob'
  })
}
