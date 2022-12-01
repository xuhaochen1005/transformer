import Cookies from 'js-cookie';
import { resetPermissions, resetRoles } from '@/utils/permission';
import { resetRouter } from '@/router';
export var UserMutation;
(function (UserMutation) {
    UserMutation["SET_TOKEN"] = "SET_TOKEN";
    UserMutation["SET_AVATAR"] = "SET_AVATAR";
    UserMutation["SET_TELEPHONE"] = "SET_TELEPHONE";
    UserMutation["SET_EMAIL"] = "SET_EMAIL";
    UserMutation["SET_INTRODUCTION"] = "SET_INTRODUCTION";
    UserMutation["SET_USERINFO"] = "SET_USERINFO";
    UserMutation["LOGOUT"] = "LOGOUT";
})(UserMutation || (UserMutation = {}));
export const userMutations = {
    [UserMutation.SET_TOKEN](state, token) {
        state.token = token;
        const expires = new Date(new Date().getTime() + 86400000);
        localStorage.setItem('token@' + document.domain, JSON.stringify({ token: token, expires: expires }));
        Cookies.set('token', token, { expires: expires, domain: document.domain });
    },
    [UserMutation.SET_AVATAR](state, avatar) {
        state.avatar = avatar;
    },
    [UserMutation.SET_TELEPHONE](state, telephone) {
        state.telephone = telephone;
    },
    [UserMutation.SET_EMAIL](state, email) {
        state.email = email;
    },
    [UserMutation.SET_INTRODUCTION](state, introduction) {
        state.introduction = introduction;
    },
    [UserMutation.SET_USERINFO](state, userinfo) {
        state.id = userinfo.id;
        state.avatar = userinfo.avatar;
        state.telephone = userinfo.telephone;
        state.email = userinfo.email;
        state.introduction = userinfo.introduction;
    },
    [UserMutation.LOGOUT](state) {
        state.token = '';
        localStorage.removeItem('token@' + document.domain);
        Cookies.remove('token', { domain: document.domain });
        resetRoles();
        resetPermissions();
        resetRouter();
    }
};
//# sourceMappingURL=mutation.js.map