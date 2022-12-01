import { AxiosPromise } from 'axios'
import request from '@/utils/request'
import { Response } from '@/model'
import { Department, DepartmentQuery } from '@/model/department'

export function GetDepartments(param: DepartmentQuery): AxiosPromise<Response<Department[]>> {
  return request({
    url: '/department/',
    method: 'GET',
    params: param
  })
}

export function CreateDepartment(params: Department): AxiosPromise<Response<any>> {
  return request({
    url: '/department/',
    method: 'POST',
    data: params
  })
}

export function updateDepartment(params: Department): AxiosPromise<Response<any>> {
  return request({
    url: '/department/' + params.id,
    method: 'PATCH',
    data: params
  })
}

export function deleteDepartment(id : number) : AxiosPromise<Response<any>> {
  return request({
    url: '/department/' + id,
    method: 'DELETE'
  })
}
export function ExportDepartment() {
  return request({
    url: '/department/export',
    method: 'POST',
    responseType: 'blob'
  })
}
