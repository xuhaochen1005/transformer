import { RouteLocationNormalized } from 'vue-router'
import { User } from '../model/common'

export enum DeviceType {
  Mobile,
  Desktop,
}

export interface AppState {
  device: DeviceType
  sidebar: {
    opened: boolean
    withoutAnimation: boolean
  }
}

export type UserState = User

export interface TagView extends Partial<RouteLocationNormalized> {
  title?: string
}

export interface TagViewState {
  visitedViews: TagView[]
  cachedViews: (string | undefined)[]
}

export interface RootState {
  app: AppState
  user: UserState
  tagViews: TagViewState
}

export const appState: AppState = {
  device: DeviceType.Desktop,
  sidebar: {
    opened: localStorage.getItem('sidebar') !== 'closed',
    withoutAnimation: false
  }
}

function getToken(): string {
  const tokenStr = localStorage.getItem('token@' + document.domain)
  if (tokenStr) {
    const tokenObj = JSON.parse(tokenStr)
    if (new Date(tokenObj.expires).getTime() < new Date().getTime()) {
      localStorage.removeItem('token@' + document.domain)
      return ''
    } else {
      return tokenObj.token
    }
  } else {
    return ''
  }
}
function getUserinfo() {
  const user : UserState = {
    token: '',
    user_name_zh: '',
    name: '',
    password: '',
    repeat_password: '',
    department_id: 0,
    department_name: '',
    email: '',
    telephone: '',
    status: 0,
    info: '',
    created_at: '',
    updated_at: ''
  }
  const storage = localStorage.getItem('userinfo@' + document.domain)
  if (storage) {
    Object.assign(user, JSON.parse(storage))
  }
  return user
}

export const userState: UserState = {
  ...getUserinfo()
}

export const tagViewState: TagViewState = {
  visitedViews: [],
  cachedViews: []
}
