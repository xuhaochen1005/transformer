import { AxiosPromise } from 'axios'
import request from '@/utils/request'
import { Response } from '@/model'
import { UserInfo } from '@/model/user'
import { ListQuery, User, UserQuery } from '@/model/common'
import type { Library, LibraryQuery } from '@/model/library'

export function GetLibraries(param: LibraryQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/libraries',
    method: 'GET',
    params: param
  })
}
export function CreateLibrary(param: Library): AxiosPromise<Response<any>> {
  return request({
    url: '/std/library',
    method: 'POST',
    data: param
  })
}

export function GetFormulaParams(param:ListQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lfp',
    method: 'GET',
    data: param
  })
}
