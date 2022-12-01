import { AxiosPromise } from 'axios'
import request from '@/utils/request'
import { Response } from '@/model'

import { ListQuery } from '@/model/common'

export function GetStdAdblLibraries(param:ListQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/adbl',
    method: 'GET',
    data: param
  })
}
