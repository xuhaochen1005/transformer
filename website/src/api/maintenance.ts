import { AxiosPromise } from 'axios'
import request from '@/utils/request'
import { Response } from '@/model'

import { ListQuery } from '@/model/common'

export function GetSysHost(param:ListQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/sys/host',
    method: 'GET',
    data: param
  })
}
export function GetSysCPU(param:ListQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/sys/cpu',
    method: 'GET',
    data: param
  })
}
export function GetSysNet(param:ListQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/sys/net',
    method: 'GET',
    data: param
  })
}
export function GetSysProcess(param:ListQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/sys/process',
    method: 'GET',
    data: param
  })
}
export function GetSysMemory(param:ListQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/sys/mem',
    method: 'GET',
    data: param
  })
}
export function GetSysBasic(param:ListQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/sys/basic',
    method: 'GET',
    data: param
  })
}
export function GetSysDisk(param:ListQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/sys/disk',
    method: 'GET',
    data: param
  })
}
