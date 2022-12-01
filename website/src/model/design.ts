import { BaseQuery, User } from '@/model/common'
import Design from '@/router/asyncRoutes/design'
import { ApproveNodeInstance } from '@/model/approve'

export enum DesignProjectUsage {
  DesignProjectUsageDefault = '/',
  DesignProjectUsageZL = 'ZL',
  DesignProjectUsageZB = 'ZB'
}

export interface DesignProject {
  id: number
  project_name:string
  creator: number
  creator_user?: User
  designer?: number
  designer_signed?: number
  designer_user?: User
  reviewer?: number
  reviewer_user?: User
  reviewer_signed?: number
  review_at: number
  // approve?: number
  project_status?: number
  design_comment: string
  comment: string
  product_name: string
  product_model: string
  deliver_at: number
  design_finish_at: number
  product_code: string
  tech_requirments: string
  order_at: number
  drawing_at: number
  design_at: number
  // approve_at: number
  check_by: number
  check_user?: User
  drafting_by?: number
  drafting_user?: User
  drafting_at: number
  checked_at: number
  product_usage: string
  product_phrase: string
  product_frequency: number
  product_frequency_special: string
  product_rated_capacity: number
  product_rated_v_high: number
  product_rated_v_low: number
  product_tap_range?: string
  product_tap_range_special: string
  product_insulation_high_li: number
  product_insulation_low_li: number
  product_insulation_high_ac: number
  product_insulation_low_ac: number
  product_insulation_level: string
  product_insulation_level_special: string
  product_temp_limit: number
  product_connection_group: string
  product_ip_level?: string
  product_short_circuit_resistance: number
  product_altitude?: number
  product_cooling_method: string
  product_load_loss: number
  product_noload_loss: number
  product_total_loss: number
  product_wire_material: string
  product_industry_freq_hold_vol: string
  product_spec_shock_vol: string
  product_induct_high_vol: number
  product_wind_type: string
  product_wire_shape: string
  need_master_approve: number
  serial_code: string
  created_at?: string
  updated_at?: string
  deleted_at?: string
}
export interface TransformerTaskForm {
  limit_total_loss: number;
  limit_no_load_current: number;
  deviation_total_loss: number;
  vol_regulation_step: number;
  has_pull_plate: string;
  limit_hv_temp_rise: number;
  deviation_no_load_current: number;
  deviation_noise_prediction: number;
  eddy_current_loss_rate: number;
  deviation_hv_temp_rise: number;
  id: number;
  rated_high_vol: number;
  resin_type: [string, number];
  hv_wire_type: {round: [string, number], flat: [string, number]};
  lv_wire_type: {round: [string, number], flat: [string, number], foil: [string, number]};
  steel_type: [string, number, number][];
  temp_rise_coeff: {
    insula_tube1: number
    insula_tube2: number
    lv: [number, number, number, number]
    hv: [number, number, number, number]
  };
  radiating_surface_coeff: {
    insula_tube1: [number, number]
    insula_tube2: [number, number]
    lv: [[number, number], [number, number], [number, number], [number, number]]
    hv: [[number, number], [number, number], [number, number], [number, number]]
  }
  insulation_temp: number;
  limit_no_load_loss: number;
  deviation_no_load_loss: number;
  phase_number: number;
  limit_lv_temp_rise: number;
  boss_spec: string;
  limit_impedance_vol: number;
  deviation_lv_temp_rise: number;
  lead_loss: number;
  deviation_load_loss: number;
  insulation_tube2: string;
  ks: number;
  insulation_tube1: string;
  insulation_class: string;
  vol_regulation_max: number;
  deviation_impedance_vol: number;
  wire_material: string;
  rated_capacity: number;
  limit_load_loss: number;
  connect_type: string;
  vol_regulation_min: number;
  limit_noise_prediction: number;
  rated_low_vol: number;
  rated_frequency: number;
}
export interface TempTaskResults {
  fitnesses: any[]
  objs: any[]
  pops: any[]
}
export interface DesignTask{
  id: number;
  task_status: number,
  design_project?: DesignProject;
  // eslint-disable-next-line no-use-before-define
  final_design_results: DesignResults[];
  designer_user?: User;
  design_results?: string;
  name: string;
  design_project_id: number;
  approve_node_instance?: ApproveNodeInstance;
  approve: number;
  creator: number;
  designer: number;
  reviewer: number;
  approve_by_creator: boolean;
  input: TransformerTaskForm | string;
  result_progress: number;
  comment: string;
  created_at?: string
  updated_at?: string
  deleted_at?: string
}
export interface DecodeDesignTask extends DesignTask{
  input: TransformerTaskForm
}
export interface DesignTaskQuery extends BaseQuery {
  id?: number;
}
export type RecommendDesignTaskQuery = BaseQuery & {
  field_ge_spec_capacity?: number
  field_le_spec_capacity?: number
  field_eq_spec_phase?: string | string[]
  field_ge_spec_high_v?: number
  field_le_spec_high_v?: number
  field_ge_spec_low_v?: number
  field_le_spec_low_v?: number
  field_lk_spec_uni_type?: string
  field_eq_std_impedance_v?: number[]
  field_re_name?: string
}
export type DesignProjectQuery = BaseQuery & {
  field_lk_project_name: string
  field_lk_product_model: string
  field_lk_serial_code?: string
  field_eq_spec_uni_type?: number
  field_eq_project_status?: number | number[]
  field_ne_project_status?: number
}
export interface DesignTaskResultsQuery extends BaseQuery{
  id?: number
}

export interface DesignResults {
  impedance_result: string | Object;
  cost_result: string | Object;
  total_cost: number;
  design_task_id: number;
  project_code: string;
  core_result: string | Object;
  project_name: string;
  product_name: string;
  design_project_id: number;
  performance_result: string | Object;
  continue_vars: string | Object;
  lv_result: string | Object;
  design_task?: DesignTask | DecodeDesignTask;
  temp_result: string | Object;
  id: number;
  design_project?: DesignProject;
  hv_result: string | Object;
  created_at?: string
  updated_at?: string
  deleted_at?: string
}

export type ListDesignResultsQuery = BaseQuery & {
  'field_eq_design_results.result_status'?: number
  'field_lk_d.project_name'?: string
  'field_lk_d.product_model'?: string
  'field_lk_d.serial_code'?: string
  'field_eq_design_results.id'?: number
  field_eq_design_task_id?: number
}

export enum DesignResultStatus {
  ResultStatusUnStarted = 0,
  ResultStatusCalcing = 1,
  ResultStatusFinsihed = 2
}

// 任务状态
export enum DesignTaskStatus {
  DesignTaskCanceled = 1,
  DesignTaskCreated = 2,
  DesignTaskStarted = 3,
  DesignTaskFinished = 4,
  DesignTaskResultSelected = 5,
  DesignTaskReviewUnaccepted = 6,
  DesignTaskReviewed = 7,
  DesignTaskApproveUnaccepted = 8,
  DesignTaskApproved = 9,
  DesignTaskCheckUnaccepted = 10,
  DesignTaskChecked = 11
}

export enum DesignProjectStatus {
  DesignProjectImported = 1,
  DesignProjectCreated = 2,
  DesignProjectStarted = 3,
  DesignProjectFinished = 4,
  DesignProjectReviewUnaccepted = 5,
  DesignProjectReviewed = 6,
  DesignProjectApproveUnaccepted = 7,
  DesignProjectApproved = 8,
  DesignProjectCheckUnAccepted = 9,
  DesignProjectChecked = 10
}
// 接缝形式
export enum SeamType {
  HalfOblique = 0,
  FullOblique = 1,
}

// 接缝形式
export enum PullType {
  WithoutPull=0,
  WithPull = 1
}

// 相数
export enum PhaseType {
  Single = '单相',
  Three = '三相'
}
