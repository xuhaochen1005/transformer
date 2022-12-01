import { ActionTree, ActionContext } from 'vuex'
import { TagViewMutation, TagViewMutations } from '@/store/tagsview/mutation'
import { RootState, TagView, TagViewState } from '@/store/state'

export enum TagViewAction {
  ADD_VIEW = 'ADD_VIEW',
  DEL_VIEW = 'DEL_VIEW',
  DEL_OTHER_VIEW = 'DEL_OTHER_VIEW',
  DEL_OTHER_VIEWS = 'DEL_OTHER_VIEWS',
  DEL_ALL_VIEWS = 'DEL_ALL_VIEWS',
}

type TagViewActionContext = {
  commit<K extends keyof TagViewMutations>(
    key: K,
    payload: Parameters<TagViewMutations[K]>[1],
  ): ReturnType<TagViewMutations[K]>
} & Omit<ActionContext<TagViewState, RootState>, 'commit'>

type TagViewActionContextWithoutPayload = {
  commit<K extends keyof TagViewMutations>(
    key: K
  ): ReturnType<TagViewMutations[K]>
} & Omit<ActionContext<TagViewState, RootState>, 'commit'>

export interface TagViewActions {
  [TagViewAction.ADD_VIEW](
    { commit }: TagViewActionContext,
    view: TagView
  ): void
  [TagViewAction.DEL_VIEW](
    { commit }: TagViewActionContext,
    view: TagView
  ): void
  [TagViewAction.DEL_OTHER_VIEW](
    { commit }: TagViewActionContext,
    view: TagView
  ): void
  [TagViewAction.DEL_OTHER_VIEWS](
    { commit }: TagViewActionContext,
    view: TagView
  ): void
  [TagViewAction.DEL_ALL_VIEWS](
    { commit }: TagViewActionContextWithoutPayload
  ): void
}

export const tagViewActions: ActionTree<TagViewState, RootState> & TagViewActions = {
  async [TagViewAction.ADD_VIEW]({ commit }, view: TagView) {
    commit(TagViewMutation.ADD_VISITED_VIEW, view)
    commit(TagViewMutation.ADD_CACHED_VIEW, view)
  },
  [TagViewAction.DEL_VIEW]({ commit }, view: TagView) {
    commit(TagViewMutation.DEL_VISITED_VIEW, view)
    commit(TagViewMutation.DEL_CACHED_VIEW, view)
  },
  [TagViewAction.DEL_OTHER_VIEW]({ commit }, view: TagView) {
    commit(TagViewMutation.DEL_OTHERS_VISITED_VIEWS, view)
    commit(TagViewMutation.DEL_OTHERS_CACHED_VIEWS, view)
  },
  [TagViewAction.DEL_OTHER_VIEWS]({ commit }, view: TagView) {
    commit(TagViewMutation.DEL_OTHERS_VISITED_VIEWS, view)
    commit(TagViewMutation.DEL_OTHERS_CACHED_VIEWS, view)
  },
  [TagViewAction.DEL_ALL_VIEWS]({ commit }) {
    commit(TagViewMutation.DEL_ALL_VISITED_VIEWS)
    commit(TagViewMutation.DEL_ALL_CACHED_VIEWS)
  }
}
