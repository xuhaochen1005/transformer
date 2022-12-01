import { BaseQuery, User } from '@/model/common'
import { DesignTask } from '@/model/design'

export const ApproveIDDesignTask = 1
export const ApproveNodeDesignTaskHeadID = 1
export interface Approve {
  id: number
  name: string
  head: number
  created_by: number
  deleted_by: number
  created_at?: string
  updated_at?: string
  deleted_at?: string
  created_by_user?: User
}

export interface ApproveNode {
  id: number
  approve_id: number
  approve?: Approve
  name: string
  api: string
  pre: number
  next: number
  created_by: number
  deleted_by: number
  created_at?: string
  updated_at?: string
  deleted_at?: string
  pre_node?: ApproveNode
  next_node?: ApproveNode
  created_by_user?: User
}

export interface ApproveNodeInstance {
  id: number
  approve_node_id: number
  approve_instance_id: number
  approval: number
  status: number
  data: string
  dataParse?: any
  comment: string
  pre: number
  next: number
  deleted_at?: string
  approve_node?: ApproveNode
  approval_user?: User
  pre_node_instance?: ApproveNodeInstance
  next_node?: ApproveNodeInstance
  created_at?: string
  updated_at?: string
}

export interface ApproveInstance {
  id: number
  approve_id: number
  data: string
  dataParse?: {
    id: number
    design_project_id: number,
    design_result_id:number
    dermander_id: number
    product_model: string
    product_name:string
    project_name:string
    serial_code:string
  }
  head: number
  status: number
  applicant: number
  participants: string
  approve?: Approve
  applicant_user?: User
  participant_users?: User[]
  approve_node_instances?: ApproveNodeInstance[]
  design_task?: DesignTask
}

// export type ApproveNodeInstanceQuery = BaseQuery & {
//   field_eq_status?: number
// }

export type ApproveInstanceQuery = BaseQuery & {
    'field_eq_approve_instance.status'?: number
    project_name?: string
    serial_code?: string,
    self?: number
}

export enum ApproveNodeStatus {
  ApproveNodeInaction = 1, // 审批节点进行中
  ApproveNodeRejected = 2, // 审批节点被拒绝
  ApproveNodeAccepted = 3, // 审批节点被通过
  ApproveNodeClosed = 4 // 审批节点被关闭
}

export enum ApproveStatus {
  ApproveNodeInaction = 1, // 审批进行中
  ApproveNodeRejected = 2, // 审批被拒绝
  ApproveNodeAccepted = 3, // 审批被通过
  ApproveNodeClosed = 4 // 审批被关闭
}
