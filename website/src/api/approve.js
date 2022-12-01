import request from '@/utils/request';
export function login(userInfo) {
    return request({
        url: '/user/token',
        method: 'POST',
        data: userInfo
    });
}
export function GetUsers(param) {
    return request({
        url: '/user/',
        method: 'GET',
        data: param
    });
}
//# sourceMappingURL=approve.js.map