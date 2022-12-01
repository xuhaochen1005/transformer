import { createStore } from 'vuex';
import { appState, userState, tagViewState } from '@/store/state';
import { appActions } from '@/store/app/action';
import { userActions } from '@/store/user/action';
import { tagViewActions } from '@/store/tagsview/action';
import { appMutations } from '@/store/app/mutation';
import { userMutations } from '@/store/user/mutation';
import { tagViewMutations } from '@/store/tagsview/mutation';
const appModule = {
    state: appState,
    actions: appActions,
    mutations: appMutations
};
const userModule = {
    state: userState,
    actions: userActions,
    mutations: userMutations
};
const tagViewModule = {
    state: tagViewState,
    actions: tagViewActions,
    mutations: tagViewMutations
};
export const store = createStore({
    modules: {
        app: appModule,
        user: userModule,
        tagViews: tagViewModule
    }
});
export function useStore() {
    return store;
}
//# sourceMappingURL=index.js.map