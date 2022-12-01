import { TagViewMutation } from '@/store/tagsview/mutation';
export var TagViewAction;
(function (TagViewAction) {
    TagViewAction["ADD_VIEW"] = "ADD_VIEW";
    TagViewAction["DEL_VIEW"] = "DEL_VIEW";
    TagViewAction["DEL_OTHER_VIEW"] = "DEL_OTHER_VIEW";
    TagViewAction["DEL_OTHER_VIEWS"] = "DEL_OTHER_VIEWS";
    TagViewAction["DEL_ALL_VIEWS"] = "DEL_ALL_VIEWS";
})(TagViewAction || (TagViewAction = {}));
export const tagViewActions = {
    async [TagViewAction.ADD_VIEW]({ commit }, view) {
        commit(TagViewMutation.ADD_VISITED_VIEW, view);
        commit(TagViewMutation.ADD_CACHED_VIEW, view);
    },
    [TagViewAction.DEL_VIEW]({ commit }, view) {
        commit(TagViewMutation.DEL_VISITED_VIEW, view);
        commit(TagViewMutation.DEL_CACHED_VIEW, view);
    },
    [TagViewAction.DEL_OTHER_VIEW]({ commit }, view) {
        commit(TagViewMutation.DEL_OTHERS_VISITED_VIEWS, view);
        commit(TagViewMutation.DEL_OTHERS_CACHED_VIEWS, view);
    },
    [TagViewAction.DEL_OTHER_VIEWS]({ commit }, view) {
        commit(TagViewMutation.DEL_OTHERS_VISITED_VIEWS, view);
        commit(TagViewMutation.DEL_OTHERS_CACHED_VIEWS, view);
    },
    [TagViewAction.DEL_ALL_VIEWS]({ commit }) {
        commit(TagViewMutation.DEL_ALL_VISITED_VIEWS);
        commit(TagViewMutation.DEL_ALL_CACHED_VIEWS);
    }
};
//# sourceMappingURL=action.js.map