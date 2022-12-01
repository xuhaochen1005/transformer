import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User} from "@/model/common";
export type Lcadi = {
  id?: number
  wind_type: string | null
  coil_inner_insulate: number | null
  coil_outer_insulate: number | null
  air_duct_thick: string | null
  air_duct_in_out_insulate: number | null
  lcadi_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}

export type ListLcadiQuery = BaseQuery & {
  field_eq_wind_type: string | null
  field_lk_coil_inner_insulate: number | null
  field_lk_coil_outer_insulate: number | null
  field_lk_air_duct_thick: string | null
  field_lk_air_duct_in_out_insulate: number | null
  field_eq_lcadi_creator?: number
}
export function GetStdLcadiLibraries(query:ListLcadiQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lcadi/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdLcadiLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcadi/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdLcadiLibraries(lcadi: Lcadi): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcadi/' + lcadi.id,
    method: 'PATCH',
    data: lcadi
  })
}
export function CreateStdLcadiLibraries(lcadi: Lcadi): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcadi',
    method: 'POST',
    data: lcadi
  })
}
export function ExportStdlcadiLibraries(){
  return request({
    url: '/std/lcadi/export',
    method: 'POST',
    responseType:'blob'
  })
}
