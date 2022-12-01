import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User} from "@/model/common";
export type Licr = {
  id?: number
  core_diameter: number | null
  pull_plate: number
  section_area: number | null
  iron_corner_weight: number | null
  stack_width_1: number | null
  stack_width_2: number | null
  stack_width_3: number | null
  stack_width_4: number | null
  stack_width_5: number | null
  stack_width_6: number | null
  stack_width_7: number | null
  stack_width_8: number | null
  stack_thick_1: number | null
  stack_thick_2: number | null
  stack_thick_3: number | null
  stack_thick_4: number | null
  stack_thick_5: number | null
  stack_thick_6: number | null
  stack_thick_7: number | null
  stack_thick_8: number | null
  stack_thick_sum: number | null
  main_level_real_stack_thick: number | null
  licr_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLicrQuery = BaseQuery & {
  field_eq_core_diameter: number | null
  field_eq_pull_plate?: number
  field_eq_section_area: number | null
  field_eq_iron_corner_weight: number | null
  field_eq_stack_width_1: number | null
  field_eq_stack_width_2: number | null
  field_eq_stack_width_3: number | null
  field_eq_stack_width_4: number | null
  field_eq_stack_width_5: number | null
  field_eq_stack_width_6: number | null
  field_eq_stack_width_7: number | null
  field_eq_stack_width_8: number | null
  field_eq_stack_thick_1: number | null
  field_eq_stack_thick_2: number | null
  field_eq_stack_thick_3: number | null
  field_eq_stack_thick_4: number | null
  field_eq_stack_thick_5: number | null
  field_eq_stack_thick_6: number | null
  field_eq_stack_thick_7: number | null
  field_eq_stack_thick_8: number | null
  field_eq_stack_thick_sum: number | null
  field_eq_main_level_real_stack_thick: number | null
  field_eq_licr_creator?: number
}
export enum LicrPullPlate {
  Yes = 1,
  No = 2
}
export function GetStdlicrLibraries(query:ListLicrQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/licr/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlicrLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/licr/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlicrLibraries(licr: Licr): AxiosPromise<Response<void>> {
  return request({
    url: '/std/licr/' + licr.id,
    method: 'PATCH',
    data: licr
  })
}
export function CreateStdlicrLibraries(licr: Licr): AxiosPromise<Response<void>> {
  return request({
    url: '/std/licr',
    method: 'POST',
    data: licr
  })
}
export function ExportStdlicrLibraries(){
  return request({
    url: '/std/licr/export',
    method: 'POST',
    responseType:'blob'
  })
}
