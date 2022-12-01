import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User} from "@/model/common";
export type Ltf = {
  id?: number
  stack_amount: number | null
  tech_factor: number | null
  ltf_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLtfQuery = BaseQuery & {
  field_eq_stack_amount: number | null
  field_eq_tech_factor: number | null
  field_eq_ltf_creator?: number
}
export function GetStdLtfLibraries(query:ListLtfQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/ltf/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdLtfLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/ltf/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdLtfLibraries(ltf: Ltf): AxiosPromise<Response<void>> {
  return request({
    url: '/std/ltf/' + ltf.id,
    method: 'PATCH',
    data: ltf
  })
}
export function CreateStdLtfLibraries(ltf: Ltf): AxiosPromise<Response<void>> {
  return request({
    url: '/std/ltf',
    method: 'POST',
    data: ltf
  })
}
export function ExportStdltfLibraries(){
  return request({
    url: '/std/ltf/export',
    method: 'POST',
    responseType:'blob'
  })
}
