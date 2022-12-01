import { MutationTree } from 'vuex'
import { UserState } from '../state'
import Cookies from 'js-cookie'
import { resetPermissions, resetRoles } from '@/utils/permission'
import router, { resetRouter } from '@/router'

export enum UserMutation {
  SET_TOKEN = 'SET_TOKEN',
  SET_USERINFO = 'SET_USERINFO',
  LOGOUT = 'LOGOUT'
}

export type UserMutations = {
  [UserMutation.SET_USERINFO](state: UserState, userinfo: UserState): void
  [UserMutation.LOGOUT](state: UserState): void
}

function setUserinfo(state: UserState) {
  localStorage.setItem('userinfo@' + document.domain,
    JSON.stringify(state))
}

export const userMutations: MutationTree<UserState> & UserMutations = {
  [UserMutation.SET_TOKEN](state: UserState, token: string): void {
    state.token = token
    const expires = new Date(new Date().getTime() + 86400000)
    localStorage.setItem('token@' + document.domain,
      JSON.stringify({ token: token, expires: expires }))
    Cookies.set('token', token, { expires: expires, domain: document.domain })
  },
  [UserMutation.SET_USERINFO](state: UserState, userinfo: UserState): void {
    Object.assign(state, userinfo)
    setUserinfo(state)
  },
  [UserMutation.LOGOUT](state: UserState): void {
    state.token = ''
    localStorage.removeItem('token@' + document.domain)
    Cookies.remove('token', { domain: document.domain })
    localStorage.removeItem('userinfo@' + document.domain)
    resetRoles()
    resetPermissions()
    resetRouter()
    router.push('/login')
  }
}
