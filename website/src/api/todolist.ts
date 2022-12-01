import type { TodoList, TodoListQuery } from '@/model/todoList'
import request from '@/utils/request'
import { AxiosPromise } from 'axios'
import { Response } from '@/model'

export function GetTodoLists(param: TodoListQuery) : AxiosPromise<Response<TodoList[]>> {
  return request({
    url: '/todolist/',
    method: 'GET'
  })
}

export function CreateTodoList(data: TodoList) : AxiosPromise<Response<any>> {
  return request({
    url: '/todolist/',
    method: 'POST',
    data
  })
}

export function UpdateTodoList(data: TodoList) : AxiosPromise<Response<any>> {
  return request({
    url: `/todolist/${data.id}`,
    method: 'PATCH',
    data
  })
}

export function DeleteTodoList(id: number) : AxiosPromise<Response<any>> {
  return request({
    url: `/todolist/${id}`,
    method: 'DELETE'
  })
}
