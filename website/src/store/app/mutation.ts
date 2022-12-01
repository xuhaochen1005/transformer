import { MutationTree } from 'vuex'
import { AppState, DeviceType } from '@/store/state'

export enum AppMutation {
  TOGGLE_SIDEBAR = 'TOGGLE_SIDEBAR',
  CLOSE_SIDEBAR = 'CLOSE_SIDEBAR',
  TOGGLE_DEVICE = 'TOGGLE_DEVICE',
}

export type AppMutations = {
  [AppMutation.TOGGLE_SIDEBAR](state: AppState, withoutAnimation: boolean): void
  [AppMutation.CLOSE_SIDEBAR](state: AppState, withoutAnimation: boolean): void
  [AppMutation.TOGGLE_DEVICE](state: AppState, device: DeviceType): void
}

export const appMutations: MutationTree<AppState> & AppMutations = {
  [AppMutation.TOGGLE_SIDEBAR](state: AppState, withoutAnimation: boolean) {
    state.sidebar.opened = !state.sidebar.opened
    state.sidebar.withoutAnimation = withoutAnimation
    if (state.sidebar.opened) {
      localStorage.setItem('sidebar', 'opened')
    } else {
      localStorage.setItem('sidebar', 'closed')
    }
  },
  [AppMutation.CLOSE_SIDEBAR](state: AppState, withoutAnimation: boolean) {
    state.sidebar.opened = false
    state.sidebar.withoutAnimation = withoutAnimation
    localStorage.setItem('sidebar', 'closed')
  },
  [AppMutation.TOGGLE_DEVICE](state: AppState, device: DeviceType) {
    state.device = device
  }
}
