import { BaseQuery } from '@/model/common'
import { AxiosPromise } from 'axios'
import { Feedback } from '@/model/feedback'
import request from '@/utils/request'
import { Response } from '@/model'

export interface FeedbackListQuery extends BaseQuery {
  id?: number
}

export function GetFeedbackList(query: FeedbackListQuery): AxiosPromise<Response<Feedback[]>> {
  return request({
    url: '/feedback',
    method: 'get',
    params: query
  })
}

export function GetSpecFeedback(id: number): AxiosPromise<Response<Feedback>> {
  return request({
    url: `/feedback/${id}`,
    method: 'get'
  })
}

export function CreateFeedback(data: Feedback): AxiosPromise<Response<any>> {
  return request({
    url: '/feedback',
    method: 'post',
    data
  })
}
export function DeleteFeedback(id: number): AxiosPromise<Response<any>> {
  return request({
    url: `/feedback/${id}`,
    method: 'delete'
  })
}
