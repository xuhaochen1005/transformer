import { AxiosPromise } from 'axios'
import request from '@/utils/request'
import { Response } from '@/model'
import { UserInfo } from '@/model/user'
import { DesignProject, DesignProjectQuery } from '@/model/design'

export function GetDesignProjects(param:DesignProjectQuery): AxiosPromise<Response<DesignProject[]>> {
  return request({
    url: '/design/projects',
    method: 'GET',
    params: param
  })
}

export function GetSpecDesignProject(id: number): AxiosPromise<Response<DesignProject>> {
  return request({
    url: `/design/project/${id}`,
    method: 'GET'
  })
}

export function CreateDesignProject(param: DesignProject): AxiosPromise<Response<any>> {
  return request({
    url: '/design/project',
    method: 'POST',
    data: param
  })
}

export function UpdateDesignProject(param: DesignProject): AxiosPromise<Response<any>> {
  return request({
    url: `/design/project/${param.id}`,
    method: 'PATCH',
    data: param
  })
}

export function DeleteDesignProject(id : number) : AxiosPromise<Response<any>> {
  return request({
    url: `/design/project/${id}`,
    method: 'DELETE'
  })
}
export function ExportDesignProject(id : number){
  return request({
    url: `/design/project_export/${id}`,
    method: 'POST',
    responseType: 'blob'
  })
}
