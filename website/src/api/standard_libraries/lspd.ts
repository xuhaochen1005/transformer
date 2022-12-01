import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User} from "@/model/common";
export type Lspd = {
  id?: number
  stack_type: string | null
  core_flux_density: number | null
  unit_iron_loss: number | null
  unit_mass_magnet_capacity: number | null
  unit_area_seam_va: number | null
  lspd_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLspdQuery = BaseQuery & {
  field_eq_stack_type: string | null
  field_eq_core_flux_density: number | null
  field_eq_unit_iron_loss: number | null
  field_eq_unit_mass_magnet_capacity: number | null
  field_eq_unit_area_seam_va: number | null
  field_eq_lspd_creator?: number
}
export function GetStdLspdLibraries(query:ListLspdQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lspd/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdLspdLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lspd/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdLspdLibraries(lspd: Lspd): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lspd/' + lspd.id,
    method: 'PATCH',
    data: lspd
  })
}
export function CreateStdLspdLibraries(lspd: Lspd): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lspd',
    method: 'POST',
    data: lspd
  })
}
export function ExportStdlspdLibraries(){
  return request({
    url: '/std/lspd/export',
    method: 'POST',
    responseType:'blob'
  })
}
