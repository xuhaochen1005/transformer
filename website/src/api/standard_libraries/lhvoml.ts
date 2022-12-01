import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import {User} from "@/model/common";

export type Lhvoml = {
  id?: number
  mod_size: string
  mod_type: string
  mod_height: number | null
  mod_amount: number | null
  mod_pic: string
  mod_num: string
  mod_suit: string
  boss_width: string
  groove: string
  nut_size: number | null
  a_size: number | null
  tap_hole_distance: number | null
  closure_pic: string
  remark: string
  lhvoml_creator?: number
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

export type ListLhvomlQuery = BaseQuery & {
  field_eq_mod_size: string
  field_eq_mod_type: string
  field_eq_mod_height: number | null
  field_eq_mod_amount: number | null
  field_eq_mod_pic: string
  field_eq_mod_num: string
  field_eq_boss_width: string
  field_eq_groove: string
  field_eq_nut_size: number | null
  field_eq_a_size: number | null
  field_eq_tap_hole_distance: number | null
  field_eq_closure_pic: string
  field_eq_mod_suit: string
  field_eq_remark: string
  field_eq_lhvoml_creator?: number
}
export function GetStdLhvomlLibraries(query:ListLhvomlQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lhvoml/',
    method: 'GET',
    params: query
  })
}

export function CreateStdLhvomlLibraries(lhvoml: Lhvoml): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lhvoml',
    method: 'POST',
    data: lhvoml
  })
}

export function UpdateStdLhvomlLibraries(lhvoml: Lhvoml): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lhvoml/' + lhvoml.id,
    method: 'PATCH',
    data: lhvoml
  })
}

export function DeleteStdLhvomlLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lhvoml/' + id,
    method: 'DELETE'
  })
}
export function ExportStdlhvomlLibraries(){
  return request({
    url: '/std/lhvoml/export',
    method: 'POST',
    responseType:'blob'
  })
}
