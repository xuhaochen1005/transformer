import { MutationTree } from 'vuex'
import { TagView, TagViewState } from '@/store/state'

export enum TagViewMutation {
  ADD_VISITED_VIEW = 'ADD_VISITED_VIEW',
  ADD_CACHED_VIEW = 'ADD_CACHED_VIEW',
  DEL_VISITED_VIEW = 'DEL_VISITED_VIEW',
  DEL_CACHED_VIEW = 'DEL_CACHED_VIEW',
  DEL_OTHERS_VISITED_VIEWS = 'DEL_OTHERS_VISITED_VIEWS',
  DEL_OTHERS_CACHED_VIEWS = 'DEL_OTHERS_CACHED_VIEWS',
  DEL_ALL_VISITED_VIEWS = 'DEL_ALL_VISITED_VIEWS',
  DEL_ALL_CACHED_VIEWS = 'DEL_ALL_CACHED_VIEWS',
  UPDATE_VISITED_VIEW = 'UPDATE_VISITED_VIEW',
}

export type TagViewMutations = {
  [TagViewMutation.ADD_VISITED_VIEW](state: TagViewState, view: TagView): void
  [TagViewMutation.ADD_CACHED_VIEW](state: TagViewState, view: TagView): void
  [TagViewMutation.DEL_VISITED_VIEW](state: TagViewState, view: TagView): void
  [TagViewMutation.DEL_CACHED_VIEW](state: TagViewState, view: TagView): void
  [TagViewMutation.DEL_OTHERS_VISITED_VIEWS](state: TagViewState, view: TagView): void
  [TagViewMutation.DEL_OTHERS_CACHED_VIEWS](state: TagViewState, view: TagView): void
  [TagViewMutation.DEL_ALL_VISITED_VIEWS](state: TagViewState): void
  [TagViewMutation.DEL_ALL_CACHED_VIEWS](state: TagViewState): void
  [TagViewMutation.UPDATE_VISITED_VIEW](state: TagViewState, view: TagView): void
}

export const tagViewMutations: MutationTree<TagViewState> & TagViewMutations = {
  [TagViewMutation.ADD_VISITED_VIEW](state: TagViewState, view: TagView) {
    if (state.visitedViews.some(v => v.path === view.path)) return
    state.visitedViews.push(
      Object.assign({}, view, {
        title: view.meta?.title || 'no-name'
      })
    )
  },
  [TagViewMutation.ADD_CACHED_VIEW](state: TagViewState, view: TagView) {
    if (view.name === null) return
    if (state.cachedViews.includes(view.name?.toString())) return
    if (!view.meta?.noCache) {
      state.cachedViews.push(view.name?.toString())
    }
  },
  [TagViewMutation.DEL_VISITED_VIEW](state: TagViewState, view: TagView) {
    for (const [i, v] of state.visitedViews.entries()) {
      if (v.path === view.path) {
        state.visitedViews.splice(i, 1)
        break
      }
    }
  },
  [TagViewMutation.DEL_CACHED_VIEW](state: TagViewState, view: TagView) {
    if (view.name === null) return
    const index = state.cachedViews.indexOf(view.name?.toString())
    index > -1 && state.cachedViews.splice(index, 1)
  },
  [TagViewMutation.DEL_OTHERS_VISITED_VIEWS](state: TagViewState, view: TagView) {
    state.visitedViews = state.visitedViews.filter(v => {
      return v.meta?.affix || v.path === view.path
    })
  },
  [TagViewMutation.DEL_OTHERS_CACHED_VIEWS](state: TagViewState, view: TagView) {
    if (view.name === null) return
    const index = state.cachedViews.indexOf(view.name?.toString())
    if (index > -1) {
      state.cachedViews = state.cachedViews.slice(index, index + 1)
    } else {
      state.cachedViews = []
    }
  },
  [TagViewMutation.DEL_ALL_VISITED_VIEWS](state: TagViewState) {
    state.visitedViews = state.visitedViews.filter(tag => tag.meta?.affix)
  },
  [TagViewMutation.DEL_ALL_CACHED_VIEWS](state: TagViewState) {
    state.cachedViews = []
  },
  [TagViewMutation.UPDATE_VISITED_VIEW](state: TagViewState, view: TagView) {
    for (let v of state.visitedViews) {
      if (v.path === view.path) {
        v = Object.assign(v, view)
        break
      }
    }
  }
}
