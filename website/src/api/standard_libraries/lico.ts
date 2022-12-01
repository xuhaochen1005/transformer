import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User} from "@/model/common";
export type Lico = {
  id?: number
  core_diameter: number | null
  pull_plate: number
  section_area: number | null
  iron_corner_weight: number | null
  stack_width_0: number | null
  stack_width_1: number | null
  stack_width_2: number | null
  stack_width_3: number | null
  stack_width_4: number | null
  stack_width_5: number | null
  stack_width_6: number | null
  stack_width_7: number | null
  stack_width_8: number | null
  stack_thick_0: number | null
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
  lico_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLicoQuery = BaseQuery & {
  field_eq_core_diameter: number | null
  field_eq_pull_plate?: number
  field_lk_section_area: number | null
  field_eq_iron_corner_weight: number | null
  field_eq_stack_width_0: number | null
  field_eq_stack_width_1: number | null
  field_eq_stack_width_2: number | null
  field_eq_stack_width_3: number | null
  field_eq_stack_width_4: number | null
  field_eq_stack_width_5: number | null
  field_eq_stack_width_6: number | null
  field_eq_stack_width_7: number | null
  field_eq_stack_width_8: number | null
  field_eq_stack_thick_0: number | null
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
  field_eq_lico_creator?: number
}
export enum LicoPullPlate {
  Yes = 1,
  No = 2
}
export function GetStdlicoLibraries(query:ListLicoQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lico/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlicoLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lico/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlicoLibraries(lico: Lico): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lico/' + lico.id,
    method: 'PATCH',
    data: lico
  })
}
export function CreateStdlicoLibraries(lico: Lico): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lico',
    method: 'POST',
    data: lico
  })
}
export function ExportStdlicoLibraries(){
  return request({
    url: '/std/lico/export',
    method: 'POST',
    responseType:'blob'
  })
}
