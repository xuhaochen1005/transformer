import { login } from '@/api/user';
import { UserMutation } from './mutation';
export var UserAction;
(function (UserAction) {
    UserAction["LOGIN"] = "LOGIN";
})(UserAction || (UserAction = {}));
export const userActions = {
    async [UserAction.LOGIN]({ commit }, userInfo) {
        userInfo.username = userInfo.username.trim();
        const response = await login(userInfo);
        if (response.status === 200) {
            commit(UserMutation.SET_TOKEN, response.data.spec);
        }
    }
};
//# sourceMappingURL=action.js.map