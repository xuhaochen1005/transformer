import { AxiosPromise } from 'axios'
import request from '@/utils/request'
import { Response } from '@/model'
import type { DesignResults, DesignTask, DesignTaskQuery, ListDesignResultsQuery } from '@/model/design'
import { RecommendDesignTaskQuery } from '@/model/design'

export function getDesignTaskList(params: DesignTaskQuery) : AxiosPromise<Response<DesignTask[]>> {
  return request({
    url: '/design/tasks',
    method: 'GET',
    params: params
  })
}

export function getDesignTaskResultsList(params: DesignTaskQuery) : AxiosPromise<Response<DesignTask[]>> {
  return request({
    url: '/design/tasks/results',
    method: 'GET',
    params: params
  })
}

export function getSpecDesignTask(id: number): AxiosPromise<Response<DesignTask>> {
  return request({
    url: `/design/task/${id}`,
    method: 'GET'
  })
}

export function createDesignTask(params: DesignTask) : AxiosPromise<Response<any>> {
  return request({
    url: '/design/task',
    method: 'POST',
    data: params,
    timeout: 150000
  })
}

export function updateDesignTask(params: DesignTask) : AxiosPromise<Response<any>> {
  return request({
    url: '/design/task/' + params.id,
    method: 'PATCH',
    data: params
  })
}

export function deleteDesignTask(id: number) : AxiosPromise<Response<any>> {
  return request({
    url: '/design/task/' + id,
    method: 'DELETE'
  })
}

export function getRecommendDesignTaskList(params : RecommendDesignTaskQuery) : AxiosPromise<Response<DesignTask[]>> {
  return request({
    url: '/design/recommand',
    method: 'GET',
    params
  })
}

export function GetDesignResults(params : ListDesignResultsQuery): AxiosPromise<Response<DesignResults[]>> {
  return request({
    url: '/design/design_results',
    method: 'GET',
    params
  })
}

export function GetSpecDesignResults(id: number): AxiosPromise<Response<DesignResults>> {
  return request({
    url: `/design/design_result/${id}`,
    method: 'GET'
  })
}

export function UpdateDesignResults(data: DesignResults): AxiosPromise<Response<any>> {
  return request({
    url: `/design/design_result/${data.id}`,
    method: 'PATCH',
    data
  })
}

export function DeleteDesignResults(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/design/design_results/' + id,
    method: 'DELETE'
  })
}

export function ExportDesignTask(id : number) {
  return request({
    url: `/design/task/export/${id}`,
    method: 'GET',
    responseType: 'blob'
  })
}
