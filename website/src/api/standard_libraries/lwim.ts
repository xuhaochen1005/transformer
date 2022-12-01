import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User} from "@/model/common";
export type Lwim = {
  id?: number
  wind_insulate_media: string | null
  wind_insulate_media_sign: string | null
  lwim_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLwimQuery = BaseQuery & {
  field_eq_wind_insulate_media: string | null
  field_eq_wind_insulate_media_sign: string | null
  field_eq_lwim_creator?: number
}
export function GetStdLwimLibraries(query:ListLwimQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lwim/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdLwimLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lwim/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdLwimLibraries(lwim: Lwim): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lwim/' + lwim.id,
    method: 'PATCH',
    data: lwim
  })
}
export function CreateStdLwimLibraries(lwim: Lwim): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lwim',
    method: 'POST',
    data: lwim
  })
}
export function ExportStdlwimLibraries(){
  return request({
    url: '/std/lwim/export',
    method: 'POST',
    responseType:'blob'
  })
}
