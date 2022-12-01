import { AxiosPromise } from 'axios'
import request from '@/utils/request'
import { Response } from '@/model'
import { UserInfo } from '@/model/user'
import type { UserQuery, User } from '@/model/common'

export function login(userInfo: UserInfo): AxiosPromise<Response<User>> {
  return request({
    url: '/user/token',
    method: 'POST',
    data: userInfo
  })
}

export function CreateUser(user: User): AxiosPromise<Response<void>> {
  return request({
    url: '/user/user',
    method: 'POST',
    data: user
  })
}

export function UpdateUser(user: User): AxiosPromise<Response<void>> {
  return request({
    url: '/user/' + user.id,
    method: 'PATCH',
    data: user
  })
}

export function GetUsers(query: UserQuery): AxiosPromise<Response<User[]>> {
  return request({
    url: '/user/',
    method: 'GET',
    params: query
  })
}

export function DeleteUser(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/user/' + id,
    method: 'DELETE'
  })
}

export function GetUserBasic(): AxiosPromise<Response<User>> {
  return request({
    url: '/user/basic',
    method: 'GET'
    //  params: query
  })
}
