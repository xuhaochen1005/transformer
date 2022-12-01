import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { StatisticsData } from '@/model/statistics'

export function GetStatistics() : AxiosPromise<Response<StatisticsData>> {
  return request({
    url: '/statistics/',
    method: 'GET'
  })
}
