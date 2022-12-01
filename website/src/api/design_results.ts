import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { DesignResults } from '@/model/design'

export type BaseQuery = {
  page: number
  limit: number
  order: string
  order_by: string
}

export function DeleteDesignResults(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/design/design_results/' + id,
    method: 'DELETE'
  })
}
export function ExportDesign(id: number) {
  return request({
    url: '/design/design_result/' + id,
    method: 'POST',
    responseType: 'blob'
  })
}

export function CreateDesignResults(DesignResults: DesignResults[]): AxiosPromise<Response<null>> {
  return request({
    url: '/design/design_results',
    method: 'POST',
    data: DesignResults
  })
}
