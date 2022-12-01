import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User} from "@/model/common";

export type Lrf = {
  id?: number
  rated_freq: number
  lrf_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
  deleted_at?: string
}

export type ListLrfQuery = BaseQuery & {
  field_eq_rated_freq: number | null
  field_eq_lrf_creator?: number
}
export function GetStdLrfLibraries(query:ListLrfQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lrf',
    method: 'GET',
    params: query
  })
}

export function CreateStdLrfLibraries(lrf: Lrf): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lrf',
    method: 'POST',
    data: lrf
  })
}

export function UpdateStdLrfLibraries(lrf: Lrf): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lrf/' + lrf.id,
    method: 'PATCH',
    data: lrf
  })
}

export function DeleteStdLrfLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lrf/' + id,
    method: 'DELETE'
  })
}
export function ExportStdlrfLibraries(){
  return request({
    url: '/std/lrf/export',
    method: 'POST',
    responseType:'blob'
  })
}
