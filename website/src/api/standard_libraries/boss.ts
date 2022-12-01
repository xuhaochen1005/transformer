import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery } from '@/model/common'

export interface Boss {
  boss_spec: string;
  c: number;
  id?: number;
}

export interface ListBossQuery extends BaseQuery {
  id?: number;
}

export function GetStdBossLibraries(query: ListBossQuery): AxiosPromise<Response<Boss[]>> {
  return request({
    url: '/std/boss',
    method: 'get',
    params: query
  })
}

export function GetSpecStdBossLibraries(id: number): AxiosPromise<Response<any>> {
  return request({
    url: '/std/boss/' + id,
    method: 'GET'
  })
}
export function DeleteStdBossLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/boss/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdBossLibraries(Boss: Boss): AxiosPromise<Response<void>> {
  return request({
    url: '/std/boss/' + Boss.id,
    method: 'PATCH',
    data: Boss
  })
}
export function CreateStdBossLibraries(Boss: Boss): AxiosPromise<Response<void>> {
  return request({
    url: '/std/boss',
    method: 'POST',
    data: Boss
  })
}
export function ExportStdBossLibraries() {
  return request({
    url: '/std/boss/export',
    method: 'POST',
    responseType: 'blob'
  })
}
