import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery,User} from "@/model/common";
export type Lfws = {
  id?: number
  std_length: number | null
  std_width: number | null
  section_area: number | null
  round_corner: number | null
  remark: string
  lfws_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLfwsQuery = BaseQuery & {
  field_lk_std_length: number | null
  field_lk_std_width: number | null
  field_lk_section_area: number | null
  field_lk_round_corner: number | null
  field_lk_remark: string
  field_eq_lfws_creator?: number
}
export function GetStdlfwsLibraries(query:ListLfwsQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lfws/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlfwsLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lfws/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlfwsLibraries(lfws: Lfws): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lfws/' + lfws.id,
    method: 'PATCH',
    data: lfws
  })
}
export function CreateStdlfwsLibraries(lfws: Lfws): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lfws',
    method: 'POST',
    data: lfws
  })
}
export function ExportStdlfwsLibraries(){
  return request({
    url: '/std/lfws/export',
    method: 'POST',
    responseType:'blob'
  })
}
