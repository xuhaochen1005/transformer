import { AxiosPromise } from 'axios'
import request from '@/utils/request'
import { Response } from '@/model'
import { BaseQuery,User} from "@/model/common";

export type Documents = {
  id?: number
  name: string
  uuid: string
  location: string
  docs_creator?: number
  creator_user?: User
  file_size: number | null
  document_type: string
  category: string
  created_at?: string
  updated_at?: string
}

export type ListDocumentsQuery = BaseQuery & {
  field_lk_name: string
  field_eq_category: string
  field_eq_document_type: string
  field_eq_docs_creator?: number
}

export function GetDocuments(query:ListDocumentsQuery): AxiosPromise<Response<any>> {
  return request({
    url: '/documents/',
    method: 'GET',
    params: query
  })
}
export function UploadDocuments(data: FormData): AxiosPromise<Response<any>> {
  return request({
    url: '/documents/upload',
    method: 'POST',
    data: data,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

export function UpdateDocuments(documents: Documents): AxiosPromise<Response<void>> {
  return request({
    url: '/documents/' + documents.id,
    method: 'PATCH',
    data: documents
  })
}
export function DeleteDocuments(id: number): AxiosPromise<Response<void>> {
  return request({
    url: '/documents/' + id,
    method: 'DELETE'
  })
}
export function DownloadDocuments(id: number) {
  return request({
    url: '/documents/download/' + id,
    method: 'GET',
    responseType: 'blob'
  })
}
export function UDocuments(documents: Documents,data: FormData): AxiosPromise<Response<any>> {
  return request({
    url: '/documents/upload',
    method: 'POST',
    data: data,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}
