export var AppMutation;
(function (AppMutation) {
    AppMutation["TOGGLE_SIDEBAR"] = "TOGGLE_SIDEBAR";
    AppMutation["CLOSE_SIDEBAR"] = "CLOSE_SIDEBAR";
    AppMutation["TOGGLE_DEVICE"] = "TOGGLE_DEVICE";
})(AppMutation || (AppMutation = {}));
export const appMutations = {
    [AppMutation.TOGGLE_SIDEBAR](state, withoutAnimation) {
        state.sidebar.opened = !state.sidebar.opened;
        state.sidebar.withoutAnimation = withoutAnimation;
        if (state.sidebar.opened) {
            localStorage.setItem('sidebar', 'opened');
        }
        else {
            localStorage.setItem('sidebar', 'closed');
        }
    },
    [AppMutation.CLOSE_SIDEBAR](state, withoutAnimation) {
        state.sidebar.opened = false;
        state.sidebar.withoutAnimation = withoutAnimation;
        localStorage.setItem('sidebar', 'closed');
    },
    [AppMutation.TOGGLE_DEVICE](state, device) {
        state.device = device;
    }
};
//# sourceMappingURL=mutation.js.map