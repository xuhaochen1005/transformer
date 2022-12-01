export interface Response<T> {
  code: number
  error: string
  total: number
  spec: T
}
