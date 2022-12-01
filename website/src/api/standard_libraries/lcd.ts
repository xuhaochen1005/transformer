import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
import { BaseQuery, User } from '@/model/common'
export type Lcd = {
  id?: number
  winding_material: string
  current_density_min: number | null
  current_density_max: number | null
  lcd_creator?: number
  creator_user?: User
  created_at?: string
  updated_at?: string
}
export type ListLcdQuery = BaseQuery & {
  field_eq_winding_material: string
  field_lt_current_density_min?: number
  field_ge_current_density_max?: number
  field_eq_lcd_creator?: number
  }
export function GetStdlcdLibraries(query:ListLcdQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/lcd/',
    method: 'GET',
    params: query
  })
}
export function DeleteStdlcdLibraries(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcd/' + id,
    method: 'DELETE'
  })
}
export function UpdateStdlcdLibraries(lcd: Lcd): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcd/' + lcd.id,
    method: 'PATCH',
    data: lcd
  })
}
export function CreateStdlcdLibraries(lcd: Lcd): AxiosPromise<Response<void>> {
  return request({
    url: '/std/lcd',
    method: 'POST',
    data: lcd
  })
}
export function ExportStdlcdLibraries() {
  return request({
    url: '/std/lcd/export',
    method: 'POST',
    responseType: 'blob'
  })
}
