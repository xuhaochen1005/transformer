import request from '@/utils/request'
export function GetSpecCalculation(query) {
  return request({
    url: '/std/calculation/',
    method: 'GET',
    params: query
  })
}
export function DeleteCalculation(id) {
  return request({
    url: '/std/calculation/' + id,
    method: 'DELETE'
  })
}
export function UpdateCalculation(calculation) {
  return request({
    url: '/std/calculation/' + calculation.id,
    method: 'PATCH',
    data: calculation
  })
}
export function CreateCalculation(calculation) {
  return request({
    url: '/std/calculation',
    method: 'POST',
    data: calculation
  })
}
// # sourceMappingURL=lfp.js.map
