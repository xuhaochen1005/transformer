import { CommitOptions, createStore, DispatchOptions, Module, Store as VuexStore } from 'vuex'
import { RootState, appState, AppState, userState, UserState, tagViewState, TagViewState } from '@/store/state'
import { appActions } from '@/store/app/action'
import { UserActions, userActions } from '@/store/user/action'
import { TagViewActions, tagViewActions } from '@/store/tagsview/action'
import { AppMutations, appMutations } from '@/store/app/mutation'
import { UserMutations, userMutations } from '@/store/user/mutation'
import { TagViewMutations, tagViewMutations } from '@/store/tagsview/mutation'

type AppStore = Omit<VuexStore<Pick<RootState, 'app'>>, 'getters' | 'commit'>
  & {
  commit<K extends keyof AppMutations, P extends Parameters<AppMutations[K]>[1]>(
    key: K, payload: P, options?: CommitOptions
  ): ReturnType<AppMutations[K]>
}

type UserStore = Omit<VuexStore<Pick<RootState, 'user'>>, 'getters' | 'commit' | 'dispatch'>
& {
  commit<K extends keyof UserMutations, P extends Parameters<UserMutations[K]>[1]>(
    key: K, payload?: P, options?: CommitOptions
  ): ReturnType<UserMutations[K]>
} & {
  dispatch<K extends keyof UserActions> (
    key: K, payload?: Parameters<UserActions[K]>[1], options?: DispatchOptions
  ): ReturnType<UserActions[K]>
}

type TagViewStore = Omit<VuexStore<Pick<RootState, 'tagViews'>>, 'getters' | 'commit' | 'dispatch'>
  & {
  commit<K extends keyof TagViewMutations, P extends Parameters<TagViewMutations[K]>[1]>(
    key: K, payload?: P, options?: CommitOptions
  ): ReturnType<TagViewMutations[K]>
} & {
  dispatch<K extends keyof TagViewActions> (
    key: K, payload?: Parameters<TagViewActions[K]>[1], options?: DispatchOptions
  ): ReturnType<TagViewActions[K]>
}

export type Store = AppStore & UserStore & TagViewStore

const appModule: Module<AppState, RootState> = {
  state: appState,
  actions: appActions,
  mutations: appMutations
}

const userModule: Module<UserState, RootState> = {
  state: userState,
  actions: userActions,
  mutations: userMutations
}

const tagViewModule: Module<TagViewState, RootState> = {
  state: tagViewState,
  actions: tagViewActions,
  mutations: tagViewMutations
}

export const store = createStore<RootState>({
  modules: {
    app: appModule,
    user: userModule,
    tagViews: tagViewModule
  }
})

export function useStore(): Store {
  return store as Store
}
