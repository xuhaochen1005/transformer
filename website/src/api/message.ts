import { AxiosPromise } from 'axios'
import request from '@/utils/request'
import { Response } from '@/model'
import type { Message, MessageQuery, MessageTemplate, MessageTemplateQuery } from '@/model/message'

export function GetUserMessages(params:MessageQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/notify/message',
    method: 'GET',
    params: params
  })
}

export function UpdateUserMessage(params: Message): AxiosPromise<Response<any>> {
  return request({
    url: '/notify/message/' + params.id,
    method: 'PATCH',
    data: params
  })
}

export function UpdateAllMessageRead(): AxiosPromise<Response<any>> {
  return request({
    url: '/notify/message/UpdateAllMessageRead',
    method: 'POST'
  })
}

export function MessageNotify(params: Message): AxiosPromise<Response<any>> {
  return request({
    url: '/notify/notify',
    method: 'POST',
    data: params
  })
}

export function DeleteMessage(id: number): AxiosPromise<Response<any>> {
  return request({
    url: '/notify/message/' + id,
    method: 'DELETE',
    data: {
      id
    }
  })
}

export function GetMessageTemplates(params: MessageTemplateQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/notify/message_template',
    method: 'GET',
    params: params
  })
}

export function CreateMessageTemplate(params: MessageTemplate): AxiosPromise<Response<any>> {
  return request({
    url: '/notify/message_template',
    method: 'POST',
    data: params
  })
}

export function UpdateMessageTemplate(params: MessageTemplate): AxiosPromise<Response<any>> {
  return request({
    url: '/notify/message_template/' + params.id,
    method: 'PATCH',
    data: params
  })
}

export function DeleteMessageTemplate(id: number): AxiosPromise<Response<any>> {
  return request({
    url: '/notify/message_template/' + id,
    method: 'DELETE',
    data: {
      id
    }
  })
}
