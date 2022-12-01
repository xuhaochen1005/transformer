import { AxiosPromise } from 'axios'
import { Response } from '@/model'
import request from '@/utils/request'
export type Calculation = {
  id?: number
  formula_name_zh:string
  formula_express:string
  formula_name: string
  formula_description: string
  formula_param: string
  remark: string
  created_at?: string
  updated_at?: string
}
export type BaseQuery = {
  page: number
  limit: number
  order: string
  order_by: string
}
export type ListCalculationQuery= BaseQuery & {
  field_eq_formula_name_zh: string
  field_eq_formula_name: string
  field_eq_formula_express: string
  field_eq_formula_description: string
  field_eq_formula_param: string
  field_eq_remark: string
}
export function GetSpecCalculation(query:ListCalculationQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/std/calculation/',
    method: 'GET',
    params: query
  })
}
export function DeleteCalculation(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/std/calculation/' + id,
    method: 'DELETE'
  })
}
export function UpdateCalculation(calculation: Calculation): AxiosPromise<Response<void>> {
  return request({
    url: '/std/calculation/' + calculation.id,
    method: 'PATCH',
    data: calculation
  })
}
export function CreateCalculation(calculation: Calculation): AxiosPromise<Response<void>> {
  return request({
    url: '/std/calculation',
    method: 'POST',
    data: calculation
  })
}
