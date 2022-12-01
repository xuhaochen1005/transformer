import { ActionContext, ActionTree } from 'vuex'
import { login } from '@/api/user'
import { UserInfo } from '@/model/user'
import { UserMutations, UserMutation } from './mutation'
import { UserState, RootState } from '../state'

export enum UserAction {
  LOGIN = 'LOGIN'
}

type UserActionContext = {
  commit<K extends keyof UserMutations>(
    key: K, payload: Parameters<UserMutations[K]>[1]
  ): ReturnType<UserMutations[K]>
} & Omit<ActionContext<UserState, RootState>, 'commit'>

export interface UserActions {
  [UserAction.LOGIN]({ commit }: UserActionContext, userInfo: UserInfo): Promise<void>
}

export const userActions: ActionTree<UserState, RootState> & UserActions = {
  async [UserAction.LOGIN]({ commit }: UserActionContext, userInfo: UserInfo): Promise<void> {
    userInfo.username = userInfo.username.trim()
    const res = await login(userInfo)
    if (res.status === 200) {
      // commit(UserMutation.SET_TOKEN, res.data.spec.token)
      commit(UserMutation.SET_USERINFO, res.data.spec)
    }
  }
}
