import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import {User} from "@/model/common";

export type Llvoml = {
  id?: number
  mod_diameter: number
  mod_type: string
  mod_height: number
  mod_amount: number
  mod_num: string| null
  mod_suit: string| null
  remark: string| null
  llvoml_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
  deleted_at?: string
}

export type BaseQuery = {
  page: number
  limit: number
  order: string
  order_by: string
}

export type ListLlvomlQuery = BaseQuery & {
  field_eq_mod_diameter: number | null
  field_eq_mod_type: string | null
  field_eq_mod_height: number | null
  // field_eq_mod_amount: number | null
  field_eq_mod_num: string | null
  // field_eq_mod_suit: string | null
  field_eq_remark: string | null
  field_eq_llvoml_creator?: number
}
export function GetStdLlvomlLibraries(query:ListLlvomlQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/llvoml',
    method: 'GET',
    params: query
  })
}

export function CreateStdLlvomlLibraries(llvoml: Llvoml): AxiosPromise<Response<void>> {
  return request({
    url: '/std/llvoml',
    method: 'POST',
    data: llvoml
  })
}

export function UpdateStdLlvomlLibraries(llvoml: Llvoml): AxiosPromise<Response<void>> {
  return request({
    url: '/std/llvoml/' + llvoml.id,
    method: 'PATCH',
    data: llvoml
  })
}

export function DeleteStdLlvomlLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/llvoml/' + id,
    method: 'DELETE'
  })
}
export function ExportStdllvomlLibraries(){
  return request({
    url: '/std/llvoml/export',
    method: 'POST',
    responseType:'blob'
  })
}
