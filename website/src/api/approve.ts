import { AxiosPromise } from 'axios'
import request from '@/utils/request'
import { Response } from '@/model'
import type {
  Approve,
  ApproveInstance,
  ApproveInstanceQuery, ApproveNodeInstance
  // ApproveNodeInstance,
  // ApproveNodeInstanceQuery
} from '@/model/approve'

// export function GetApproveNodeInstances(param: ApproveNodeInstanceQuery): AxiosPromise<Response<ApproveNodeInstance[]>> {
//   return request({
//     url: '/approve/approve_node_instances/',
//     method: 'GET',
//     params: param
//   })
// }

export function GetApproveInstances(param: ApproveInstanceQuery): AxiosPromise<Response<ApproveInstance[]>> {
  return request({
    url: '/approve/approve_instances',
    method: 'GET',
    params: param
  })
}

export function UpdateApproveNodeInstance(param: ApproveNodeInstance): AxiosPromise<Response<any>> {
  return request({
    url: `/approve/approve_node_instance/${param.id}`,
    method: 'PATCH',
    data: param,
    timeout: 15000
  })
}
